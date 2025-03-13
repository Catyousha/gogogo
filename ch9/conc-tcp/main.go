package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var count = 0

func handleConn(c net.Conn) {
	fmt.Print(".")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		fmt.Println(temp)
		counter := fmt.Sprintf("Client number: %d\n", count)
		c.Write([]byte(counter))
	}

	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	// exercise 3: Add UNIX signal processing to the concurrent TCP server
	// to gracefully stop the server process when a given signal is received.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// Handle signals in a separate goroutine
	go func() {
		sig := <-sigChan
		fmt.Printf("\nReceived signal: %v\n", sig)
		fmt.Println("Shutting down server...")
		l.Close()
		done <- true
	}()

	fmt.Printf("Server listening on port %s\nPress Ctrl+C to stop\n", arguments[1])

	// Main server loop with shutdown handling
	for {
		select {
		case <-done:
			fmt.Println("Server stopped")
			return
		default:
			c, err := l.Accept()
			if err != nil {
				// Check if the error is due to server shutdown
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				}
				fmt.Println(err)
				continue
			}
			go handleConn(c)
			count++
		}
	}
}
