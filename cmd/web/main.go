package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/snippet/view", app.SnippetViewHandler)
	mux.HandleFunc("/snippet/create", app.SnippetCreateHandler)

	infoLog.Printf("Starting server on %s\n", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
