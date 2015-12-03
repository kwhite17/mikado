package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./app")))
	log.Fatalf("%s", http.ListenAndServe("localhost:3000", nil))
}
