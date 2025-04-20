package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/snippet/view", SnippetViewHandler)
	mux.HandleFunc("/snippet/create", SnippetCreateHandler)

	infoLog.Printf("Starting server on %s\n", *addr)

	err := http.ListenAndServe(*addr, mux)
	errLog.Fatal(err)
}
