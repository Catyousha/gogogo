package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var SERVER = ""
var PATH = ""
var TIMESWAIT = 0
var TIMESWAITMAX = 5
var in = bufio.NewReader(os.Stdin)

func getInput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- result
}

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Usage: ws-c <server> <path>")
		return
	}

	SERVER = args[1]
	PATH = args[2]
	fmt.Println("Connecting to", SERVER, "at", PATH)

	// send interrupt to chan if reveive interrupt signal from os
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)

	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal("websocket.DefaultDialer.Dial() err:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage() err:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {

		case <-time.After(4 * time.Second):
			log.Println("Please give me input!", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				log.Println("Timeout!")
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}

		case <-done:
			return

		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("WriteMessage() err:", err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)

		case <-interrupt:
			log.Println("Caught interrupt signal - quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("WriteMessage() err:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(1 * time.Second):
			}
			return
		}
	}
}
