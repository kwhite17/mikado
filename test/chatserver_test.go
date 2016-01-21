package test

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/kwhite17/mikado/server"
)

func TestBuildServerReturnsServerWithCorrectReadTimeout(t *testing.T) {
	customServer := server.BuildChatServer("")
	if customServer.ReadTimeout != 36000*time.Second {
		t.Fatalf("Expected Read Timeout: 36000, Actual Read Timeout: %d", customServer.ReadTimeout)
	}
}

func TestBuildServerReturnsServerWithCorrectWriteTimeout(t *testing.T) {
	customServer := server.BuildChatServer("")
	if customServer.WriteTimeout != 36000*time.Second {
		t.Fatalf("Expected Write Timeout: 36000, Actual Read Timeout: %d", customServer.WriteTimeout)
	}
}

func TestBuildServeMuxReturnsHandlerWithServeMuxType(t *testing.T) {
	handler := server.BuildServeMux()
	actual := reflect.TypeOf(handler).String()
	expected := reflect.TypeOf(&http.ServeMux{}).String()
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestBuildHandlerReturnsHandlerWithSocketHandlerType(t *testing.T) {
	handler := server.BuildHandler()
	actual := reflect.TypeOf(handler).String()
	expected := reflect.TypeOf(&server.SocketHandler{}).String()
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestBuildHandlerReturnsHandlerWithConnectionsMap(t *testing.T) {
	handler := server.BuildHandler()
	if handler.Connections == nil {
		t.Fatal("No map exists to manage connections!")
	}
}

func TestBuildServerReturnsServerWithCorrectHandler(t *testing.T) {
	customServer := server.BuildChatServer("")
	if customServer.Handler == nil {
		t.Fatal("Server is missing appropriate handler!")
	}
}

func TestBuildServerReturnsServerWithCorrectAddress(t *testing.T) {
	expected := ":4000"
	customServer := server.BuildChatServer(expected)
	if customServer.Addr != ":4000" {
		t.Fatalf("Wrong Port Number! Expected: %s, Actual: %s", expected, customServer.Addr)
	}
}
