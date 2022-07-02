package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexResponse struct {
	Body string `json:"body"`
}

func index(w http.ResponseWriter, r *http.Request) {
	response := IndexResponse{Body: "Hello world!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Starting server on 0.0.0.0:4000, url: http://localhost:4000")

	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}
