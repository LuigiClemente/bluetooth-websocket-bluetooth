package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// replace this stub with your actual implementation
func sendBluetooth(data []byte) error {
	// your Bluetooth sending logic goes here
	return nil
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}

		// Handle your buffer here
		// send it to Bluetooth
		err = sendBluetooth(msg)
		if err != nil {
			log.Printf("error sending data to Bluetooth: %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("Starting WebSocket server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
