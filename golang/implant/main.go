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
	if len(os.Args) != 3 {
		log.Fatal("Missing arguments. Please correct.")
	}

	host := os.Args[1]
	c2_pass := os.Args[2]

	fmt.Println("Connecting to " + host + " with username " + c2_pass)

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
