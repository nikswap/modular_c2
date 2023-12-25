package main

import (
	"os/exec"
)

// Doit can take arguments as well
func DoIt() (string, error) {
	out, err := exec.Command("whoami").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// type speaker struct {
// }

// func (s *speaker) Speak() string {
//     return "hello"
// }

// // Exported
// var Speaker speaker
// var SpeakerName = "Alice"

// Compile with:
// go build -buildmode=plugin -o <name>.so <name>.go.
