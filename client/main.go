package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

// Message represents the structure of the websocket message
type Message struct {
	MessageType int
	Data        []byte
}

func main() {
	//connect with remote ws
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3000",
		Path:   "/ws",
	}
	fmt.Printf("Connecting to %s \n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer conn.Close()

	//Channels for managing messages
	send := make(chan Message)
	done := make(chan struct{})

	//Goroutine for read messages
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			fmt.Printf("Received: %s \n", message)
		}
	}()

	//Goroutine for write messages
	go func() {
		for {
			select {
			case msg := <-send:
				//write that to the websocket connection
				err := conn.WriteMessage(msg.MessageType, msg.Data)
				if err != nil {
					log.Println("write: ", err)
					return
				}
			case <-done:
				return
			}
		}
	}()

	//Read the input from the terminal and sed it to the web socket server
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type something...")
	for scanner.Scan() {
		text := scanner.Text()
		//Send the text to the channel
		send <- Message{websocket.TextMessage, []byte(text)}
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scanner err: ", err)
	}
}
