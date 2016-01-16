package server

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type ChatServer struct {
	Server      *http.Server
	Connections map[*websocket.Conn]bool
}

func BuildServer() http.Server {
	server := http.Server{ReadTimeout: 36000 * time.Second, WriteTimeout: 36000 * time.Second}
	return server
}
