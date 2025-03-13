package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
)

const (
	MIN_RANDOM = 1
	MAX_RANDOM = 100
)

var activeConnections = 0

func handleConnection(c net.Conn) {
	fmt.Printf("New client connected. Active connections: %d\n", activeConnections)
	
	// exit notifier
	defer func() {
		c.Close()
		activeConnections--
		fmt.Printf("Client disconnected. Active connections: %d\n", activeConnections)
	}()


	// message receiver
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Client requested to stop")
			return
		}

		randomNum := rand.Intn(MAX_RANDOM-MIN_RANDOM+1) + MIN_RANDOM
		response := fmt.Sprintf("Random number between %d-%d: %d\n",
			MIN_RANDOM, MAX_RANDOM, randomNum)

		_, err = c.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer l.Close()

	fmt.Printf("Random number generator server listening on port %s\n", arguments[1])

	// connection handler
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		activeConnections++
		go handleConnection(c)
	}
}
