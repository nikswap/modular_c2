package server

import (
	"fmt"
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

func main() {
	nu := time.Now()
	fmt.Println(nu)
}
