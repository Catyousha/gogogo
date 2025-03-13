package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main()  {
	if len(os.Args) == 1 {
		fmt.Println("Please provide socket path")
		return
	}

	socketPath := os.Args[1]

	c, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}

	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read from stdin:", err)
			return
		}

		_, err = c.Write([]byte(text))
		if err != nil {
			fmt.Println("Failed to write to server:", err)
			return
		}

		buf := make([]byte, 256)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Failed to read from server:", err)
			return
		}
		fmt.Print("Read:", string(buf[0:n]))

		if strings.TrimSpace(string(buf[0:n])) == "STOP" {
			fmt.Println("Exiting...")
			return
		}
		time.Sleep(5 * time.Second)
	}
}