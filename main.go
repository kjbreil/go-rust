package main

import (
	"flag"
	"log"
)

type rcv struct {
	Message    string
	Identifier int
	Type       string
	Stacktrace *string
}

var addr = flag.String("addr", "127.0.0.1:28016", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	ss := serverSettings{host: "127.0.0.1", port: 28016, password: "docker"}

	send := make(chan string)

	chat := make(chan string)

	generic := make(chan string)

	go func() {
		send <- "{\n  \"Identifier\": -1,\n  \"Message\": \"say HELLO\",\n  \"Name\": \"WebRcon\"}"
	}()

	Connect(ss, send, generic, chat)

}
