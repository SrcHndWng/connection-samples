package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

const redisURL = "redis://localhost:6379"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("args invalid. please input your name.")
		os.Exit(1)
	}
	userName := os.Args[1]

	conn, err := redisurl.ConnectToURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userKey := fmt.Sprintf("online.%s", userName)
	val, err := conn.Do("SET", userKey, userName, "NX", "EX", "120")
	if err != nil {
		log.Fatal(err)
	}
	if val == nil {
		fmt.Println("User already online")
		os.Exit(1)
	}

	val, err = conn.Do("SADD", "users", userName)
	if err != nil {
		log.Fatal(err)
	}

	subChan := make(chan string)
	go subscribe(subChan)

	inputChan := make(chan string)
	go stdin(inputChan, userName)

	conn.Do("PUBLISH", "messages", fmt.Sprintf("%s has joined", userName))

	tickerChan := time.NewTicker(time.Second * 60).C
	isOnline := true

	for isOnline {
		select {
		case msg := <-subChan:
			fmt.Println(msg)
		case input := <-inputChan:
			switch input {
			case "/exit":
				isOnline = false
			case "/who":
				names, err := redis.Strings(conn.Do("SMEMBERS", "users"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(fmt.Sprintf("Online members are %s.", strings.Join(names, ",")))
			default:
				conn.Do("PUBLISH", "messages", fmt.Sprintf("%s:%s", userName, input))
			}
		case <-tickerChan:
			val, err = conn.Do("SET", userKey, userName, "XX", "EX", "120")
			if err != nil || val == nil {
				log.Println("Heartbeat set failed")
				isOnline = false
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	conn.Do("DEL", userKey)
	conn.Do("SREM", "users", userName)
	conn.Do("PUBLISH", "messages", fmt.Sprintf("%s has left", userName))
}

func subscribe(subChan chan string) {
	conn, err := redisurl.ConnectToURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("messages")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			subChan <- string(v.Data)
		case redis.Subscription:
			break
		case error:
			return
		}
	}
}

func stdin(inputChan chan string, userName string) {
	prompt := fmt.Sprintf("%s>", userName)
	bio := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _, err := bio.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		inputChan <- string(line)
	}
}
