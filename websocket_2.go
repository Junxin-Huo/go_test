package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func main() {
	origin := "http://localhost/"
	url := "ws://xxx"

	config, err := websocket.NewConfig(url, origin)
	if err != nil {
		log.Fatal(err)
	}
	config.Header["test_key"] = []string{"value_1", "value_2", "value_3"}

	//var msg = make([]byte, 512)
	//client, err := websocket.NewClient(config, msg)
	//if err != nil {
	//	log.Fatal(err)
	//}


	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	//ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}