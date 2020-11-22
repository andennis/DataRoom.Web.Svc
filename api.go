package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func nodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Node")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/node", nodeHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
