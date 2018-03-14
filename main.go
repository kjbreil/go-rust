package main

import (
	"fmt"

	"github.com/kjbreil/wsrcon"
)

func main() {
	// Connect to local docker
	ss := wsrcon.Settings{Host: "127.0.0.1", Port: 28016, Password: "docker"}

	rcon := wsrcon.Connect(&ss)

	rcon.AddGenericHandler(basicGenericHandler)
	rcon.AddChatHandler(basicChatHandler)
	rcon.Start()
}

func basicGenericHandler(msg string) {
	fmt.Printf("Generic Message: %s", msg)
}

func basicChatHandler(msg string) {
	fmt.Printf("Generic Message: %s", msg)
}
