package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//upgrade the incoming GET request into a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//loop to read the messages that the clients send and write them on the server
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received Message: %s", message)

		res := fmt.Sprintf("%s , at %s", string(message), time.Now().String())

		if err := conn.WriteMessage(messageType, []byte(res)); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", Handler)

	fmt.Println("Starting server on: 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
