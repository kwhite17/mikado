package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/kwhite17/mikado/server"
)

func TestChatServerContainsServer(t *testing.T) {
	cs := server.ChatServer{Server: &http.Server{}}
	if cs.Server.MaxHeaderBytes != 0 {
		t.Fatalf("Expected MaxHeaderBytes: 0, Actual MaxHeaderBytes: %d", cs.Server.MaxHeaderBytes)
	}
}

func TestChatServerContainsConnectionsMap(t *testing.T) {
	cs := server.ChatServer{Server: &http.Server{}, Connections: make(map[*websocket.Conn]bool)}
	if len(cs.Connections) != 0 {
		t.Fatalf("Expected Connections: 0, Actual Connections: %d", len(cs.Connections))
	}
}

func TestBuildServerReturnsServerWithCorrectReadTimeout(t *testing.T) {
	customServer := server.BuildServer()
	if customServer.ReadTimeout != 36000*time.Second {
		t.Fatalf("Expected Read Timeout: 36000, Actual Read Timeout: %d", customServer.ReadTimeout)
	}
}

func TestBuildServerReturnsServerWithCorrectWriteTimeout(t *testing.T) {
	customServer := server.BuildServer()
	if customServer.WriteTimeout != 36000*time.Second {
		t.Fatalf("Expected Write Timeout: 36000, Actual Read Timeout: %d", customServer.WriteTimeout)
	}
}
