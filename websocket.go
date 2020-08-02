package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{ /* set fields as needed */ }
	ws, _, err := dialer.Dial("ws://127.0.0.1:5998", nil)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		fmt.Println(err)
		// handle error
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	fmt.Println("the message is %s", p)
}
