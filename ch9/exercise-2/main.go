package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var activeConnections = 0

func handleConnection(c net.Conn) {
	fmt.Printf("New client connected. Active connections: %d\n", activeConnections)

	defer func() {
		c.Close()
		activeConnections--
		fmt.Printf("Client disconnected. Active connections: %d\n", activeConnections)
	}()

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		input := strings.TrimSpace(string(netData))
		if input == "STOP" {
			fmt.Println("Client requested to stop")
			return
		}

		// Parse the range values
		nums := strings.Split(input, " ")
		if len(nums) != 2 {
			c.Write([]byte("Please provide two numbers as 'min max'\n"))
			continue
		}

		min, err1 := strconv.Atoi(nums[0])
		max, err2 := strconv.Atoi(nums[1])

		if err1 != nil || err2 != nil {
			c.Write([]byte("Invalid number format\n"))
			continue
		}

		if min >= max {
			c.Write([]byte("Min must be less than max\n"))
			continue
		}

		// Generate random number in range
		randomNum := rand.Intn(max-min+1) + min
		response := fmt.Sprintf("Random number between %d-%d: %d\n",
			min, max, randomNum)

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
