package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	stb "./Stubs"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func nodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Node")
}

func nodesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get nodes")
	json.NewEncoder(w).Encode(stb.GetStubNodes())
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/node", nodeHandler)
	http.HandleFunc("/nodes", nodesHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
