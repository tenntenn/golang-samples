package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {

	handlers := http.NewServeMux()
	handlers.HandleFunc("/1", handler)
	server := &http.Server{Addr:":8080", Handler: handlers}

	go server.ListenAndServe()

	handlers.HandleFunc("/2", handler)

	for {}
}
