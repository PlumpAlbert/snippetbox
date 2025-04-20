package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/snippet/view", SnippetViewHandler)
	mux.HandleFunc("/snippet/create", SnippetCreateHandler)

	log.Printf("Starting server on %s\n", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
