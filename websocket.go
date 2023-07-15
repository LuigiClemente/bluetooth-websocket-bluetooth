package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		var msg []byte
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		// Handle your buffer here
		// send it to bluetooth
		err = sendBluetooth(msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
