package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Fprintf(w, "Please use ws://localhost:1234/ws to connect to the websocket server.")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}

	defer ws.Close()
	
	for {
		// read message from client
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		log.Print("Received: ", string(message))
		
		// echoing msg back to client
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("Failed to write message:", err)
			break
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	// open multiplexer router
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:   PORT,
		Handler: mux,
		IdleTimeout: 10 * time.Second,
		ReadTimeout: time.Second,
		WriteTimeout: time.Second,
	}

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/ws", wsHandler)

	log.Println("Listening to TCP port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println("Failed to start server:", err)
	}

	// client can access through `websocat ws://localhost:1234/ws`
}