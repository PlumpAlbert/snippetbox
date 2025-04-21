package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	ts, err := template.ParseFiles(
		"./ui/html/base.html.tmpl",
		"./ui/html/partials/nav.html.tmpl",
		"./ui/html/pages/home.html.tmpl",
	)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) SnippetViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
}

func (app *application) SnippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
