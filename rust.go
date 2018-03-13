package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type serverSettings struct {
	host     string
	port     int
	password string
}

// Connect connects to a rust server given the input settings and channels to put data into
func Connect(settings serverSettings, send chan string, generic chan string, chat chan string) {

	// creat interrupt channel to feed into when an os.Interrupt is triggered (ctrl-c)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// combine port with address
	host := settings.host + ":" + strconv.Itoa(settings.port)

	// create URL scheme
	u := url.URL{Scheme: "ws", Host: host, Path: settings.password}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	var data rcv

	go func() {
		defer close(done)
		for {

			// _, message, err := c.ReadMessage()
			err = c.ReadJSON(&data)

			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s, msg: %s", data.Type, data.Message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case t := <-send:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
