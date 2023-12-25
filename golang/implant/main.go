package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mod := os.Args[1]
	plugin, err := plugin.Open(mod)
	CheckError(err)

	doItSymbol, err := plugin.Lookup("DoIt")
	CheckError(err)

	doItFunc, ok := doItSymbol.(func() (string, error))
	if !ok {
		fmt.Println("Could not do the thing")
	}

	result, err := doItFunc()
	CheckError(err)

	fmt.Println("RESULT: " + result)

	// symSpeaker, err := plugin.Lookup("Speaker")
	// CheckError(err)

	// var speaker Speaker
	// speaker, ok := symSpeaker.(Speaker)
	// if !ok {
	// 	return errors.New("unexpected type from module symbol")
	// }

	// speaker.Speak()
}
