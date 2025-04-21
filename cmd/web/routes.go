package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/snippet/view", app.SnippetViewHandler)
	mux.HandleFunc("/snippet/create", app.SnippetCreateHandler)

	return mux
}
