package main

import (
	"fmt"
	"net"
	"os"
)

func echo(c net.Conn)  {
	for {
		buf := make([]byte, 128)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Failed to read data:", err)
			return
		}

		data := buf[0:n]
		fmt.Print("Server got:", string(data))
		// Write data back to the client (echo)
		_, err = c.Write(data)
		if err != nil {
			fmt.Println("Failed to write data:", err)
			return
		}
	}
}

func main()  {
	if len(os.Args) == 1 {
		fmt.Println("Please provide socket path")
		return
	}

	socketPath := os.Args[1]

	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing socket file")
		err = os.Remove(socketPath)
		if err != nil {
			fmt.Println("Failed to delete existing socket file:", err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Failed to create listener:", err)
		return
	}

	defer l.Close()
	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			return
		}

		go echo(fd)
	}
}