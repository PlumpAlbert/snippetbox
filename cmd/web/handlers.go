package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/plumpalbert/snippetbox/internal/models"
)

func (app *application) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.renderTemplate(w, http.StatusOK, "home.html.tmpl", data)
}

func (app *application) SnippetViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.notFound(w)
			return
		}

		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.renderTemplate(w, http.StatusOK, "view.html.tmpl", data)
}

func (app *application) SnippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := app.snippets.Insert(
		"Oh snail",
		"Oh snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa",
		7,
	)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
