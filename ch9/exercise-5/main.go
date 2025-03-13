package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v\n", err)
		return
	}
	defer ws.Close()

	// Read initial message containing number of random integers wanted
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Printf("Failed to read message: %v\n", err)
		return
	}

	count, err := strconv.Atoi(strings.TrimSpace(string(message)))
	if err != nil || count <= 0 {
		ws.WriteMessage(mt, []byte("Please send a valid positive number"))
		return
	}

	// Generate and send random numbers
	response := make([]string, count)
	for i := 0; i < count; i++ {
		response[i] = strconv.Itoa(rand.Intn(1000))
	}

	err = ws.WriteMessage(mt, []byte(strings.Join(response, ",")))
	if err != nil {
		log.Printf("Failed to write message: %v\n", err)
		return
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a port number")
	}

	port := os.Args[1]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WebSocket Random Number Generator\nConnect to ws://localhost:%s/ws", port)
	})
	http.HandleFunc("/ws", handleWebSocket)

	addr := ":" + port
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
