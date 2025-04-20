package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/snippet/view", SnippetViewHandler)
	mux.HandleFunc("/snippet/create", SnippetCreateHandler)

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
