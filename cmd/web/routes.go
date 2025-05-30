package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.IndexHandler)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.SnippetViewHandler)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.SnippetCreateViewHandler)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.SnippetCreateHandler)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
