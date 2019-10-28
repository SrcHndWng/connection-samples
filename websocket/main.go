package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// Echo ...
func Echo(ws *websocket.Conn) {
	for {
		var received string

		err := websocket.Message.Receive(ws, &received)
		if err != nil {
			log.Fatal("reveive error.", err)
			break
		}

		log.Println("Server received : " + received)

		datetime := time.Now().Format("2006/1/2 15:04:05")
		response := fmt.Sprintf("[%s] Server received : %s\n", datetime, received)

		err = websocket.Message.Send(ws, response)
		if err != nil {
			log.Fatal("send error.", err)
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
