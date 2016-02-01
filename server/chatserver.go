package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type SocketHandler struct {
	Connections map[*websocket.Conn]bool
}

func (handler SocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	handler.Connections[conn] = true
	sendMessage(handler.Connections, conn)
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

func BuildChatServer(address string) *http.Server {
	server := http.Server{Addr: address, Handler: BuildServeMux(), ReadTimeout: 36000 * time.Second, WriteTimeout: 36000 * time.Second}
	return &server
}

func BuildServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("./app")))
	serveMux.Handle("/socket", BuildHandler())
	return serveMux
}

func BuildHandler() *SocketHandler {
	handler := &SocketHandler{make(map[*websocket.Conn]bool)}
	return handler
}
