package server

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestBuildServerReturnsServerWithCorrectReadTimeout(t *testing.T) {
	customServer := BuildChatServer("")
	if customServer.ReadTimeout != 36000*time.Second {
		t.Fatalf("Expected Read Timeout: 36000, Actual Read Timeout: %d", customServer.ReadTimeout)
	}
}

func TestBuildServerReturnsServerWithCorrectWriteTimeout(t *testing.T) {
	customServer := BuildChatServer("")
	if customServer.WriteTimeout != 36000*time.Second {
		t.Fatalf("Expected Write Timeout: 36000, Actual Read Timeout: %d", customServer.WriteTimeout)
	}
}

func TestBuildServeMuxReturnsHandlerWithServeMuxType(t *testing.T) {
	handler := BuildServeMux()
	actual := reflect.TypeOf(handler).String()
	expected := reflect.TypeOf(&http.ServeMux{}).String()
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestBuildHandlerReturnsHandlerWithSocketHandlerType(t *testing.T) {
	handler := BuildHandler()
	actual := reflect.TypeOf(handler).String()
	expected := reflect.TypeOf(&SocketHandler{}).String()
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestBuildHandlerReturnsHandlerWithConnectionsMap(t *testing.T) {
	handler := BuildHandler()
	if handler.Connections == nil {
		t.Fatal("No map exists to manage connections!")
	}
}

func TestBuildServerReturnsServerWithCorrectHandler(t *testing.T) {
	customServer := BuildChatServer("")
	if customServer.Handler == nil {
		t.Fatal("Server is missing appropriate handler!")
	}
}

func TestBuildServerReturnsServerWithCorrectAddress(t *testing.T) {
	expected := ":4000"
	customServer := BuildChatServer(expected)
	if customServer.Addr != ":4000" {
		t.Fatalf("Wrong Port Number! Expected: %s, Actual: %s", expected, customServer.Addr)
	}
}
