package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type wsHandler struct {
	connections map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (handler wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	handler.connections[conn] = true
	sendMessage(handler.connections, conn)
}

func sendMessage(conns map[*websocket.Conn]bool, source *websocket.Conn) {
	for {
		messageType, message, err := source.ReadMessage()
		if err != nil {
			source.Close()
			log.Println(err)
			return
		}
		for k := range conns {
			if source != k {
				k.WriteMessage(messageType, message)
			}
		}
	}
}

func main() {
	var connections = make(map[*websocket.Conn]bool)
	handler := wsHandler{connections: connections}
	http.Handle("/", http.FileServer(http.Dir("./app")))
	http.Handle("/socket", handler)
	s := &http.Server{Addr: ":" + os.Getenv("PORT")}
	log.Printf("Address: %v", s.Addr)
	log.Fatalf("%s", s.ListenAndServe())
}
