package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	address := ":8888"
	tcp, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
	defer conn.Close()

	for {
		request := make([]byte, 128) // To prevent flood attack, set max bytes
		length, err := conn.Read(request)
		if err != nil {
			log.Printf("read error. err = %v\n", err)
			break
		}

		message := string(request)[0 : length-2] // remove last "\r", "\n"
		log.Printf("received length = %d, message = %s", length, message)

		if message == "quit" {
			log.Println("quit!!")
			break
		} else {
			datetime := time.Now().Format("2006/1/2 15:04:05")
			response := fmt.Sprintf("[%s] your message is %s\n", datetime, message)
			conn.Write([]byte(response))
		}
	}
}
