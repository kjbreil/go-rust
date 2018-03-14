package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kjbreil/wsrcon"
)

var addr = flag.String("addr", "127.0.0.1:28016", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	ss := wsrcon.Settings{Host: "127.0.0.1", Port: 28016, Password: "docker"}

	// send := make(chan string)

	// chat := make(chan string)

	// generic := make(chan string)

	// go func() {
	// 	send <- "{\n  \"Identifier\": -1,\n  \"Message\": \"say HELLO\",\n  \"Name\": \"WebRcon\"}"
	// }()

	// ConnectOld(ss, send, generic, chat)

	rcon := wsrcon.Connect(&ss)

	rcon.AddChatHandler(basicChatHandler)

	rcon.Start()
}

func basicChatHandler(msg string) {
	fmt.Printf("BASIC CHAT: %s", msg)
}
