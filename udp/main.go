package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	address := ":8888"
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	buf := make([]byte, 512)
	length, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
	}

	message := string(buf)[0 : length-1] // remove last "\r", "\n"
	log.Printf("received length = %d, message = %s", length, message)
	datetime := time.Now().Format("2006/1/2 15:04:05")
	response := fmt.Sprintf("[%s] your message is %s\n", datetime, message)
	conn.WriteToUDP([]byte(response), addr)
}
