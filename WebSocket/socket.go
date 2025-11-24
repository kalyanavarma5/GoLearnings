package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Upgrader upgrades HTTP connection to WebSocket with default buffer sizes
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (consider stricter checks in production)
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Echo loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// Log received message
		log.Printf("Received: %s", message)

		// Write back the same message
		if err = conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", echoHandler)

	log.Println("WebSocket echo server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
