package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

const SOCKET_PATH = "/tmp/random.socket"

func server() {
	// Remove existing socket file if it exists
	if _, err := os.Stat(SOCKET_PATH); err == nil {
		if err := os.Remove(SOCKET_PATH); err != nil {
			log.Fatal(err)
		}
	}

	// Create listener
	l, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	fmt.Println("Server started at", SOCKET_PATH)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go func(c net.Conn) {
			defer c.Close()
			for {
				buf := make([]byte, 1024)
				n, err := c.Read(buf)
				if err != nil {
					return
				}

				cmd := strings.TrimSpace(string(buf[0:n]))
				if cmd == "STOP" {
					return
				}

				// Generate random number
				num := rand.Intn(100) + 1
				_, err = c.Write([]byte(fmt.Sprintf("%d\n", num)))
				if err != nil {
					return
				}
			}
		}(conn)
	}
}

func client() {
	conn, err := net.Dial("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Connected to server. Type anything to get a random number. Type 'STOP' to exit.")

	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if _, err = conn.Write([]byte(text)); err != nil {
			log.Fatal("write error:", err)
		}

		if text == "STOP" {
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("read error:", err)
		}

		fmt.Printf("Random number: %s", string(buf[0:n]))
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program [server|client]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		server()
	case "client":
		client()
	default:
		fmt.Println("Invalid argument. Use 'server' or 'client'")
	}
}
