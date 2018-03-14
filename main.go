package main

import (
	"fmt"

	"github.com/kjbreil/wsrcon"
)

func main() {

	ss := wsrcon.Settings{Host: "127.0.0.1", Port: 28016, Password: "docker"}

	rcon := wsrcon.Connect(&ss)

	rcon.AddChatHandler(basicChatHandler)

	rcon.Start()
}

func basicChatHandler(msg string) {
	fmt.Printf("BASIC CHAT: %s", msg)
}
