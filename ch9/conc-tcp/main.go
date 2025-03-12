package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConn(c)
		count++
	}
}
