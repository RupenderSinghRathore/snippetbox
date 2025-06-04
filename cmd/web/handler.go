package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/RupenderSinghRathore/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData()
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.Error(w, "Snippet not found", http.StatusBadRequest)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData()
	data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := 7

		// Pass the data to the SnippetModel.Insert() method, receiving the
		// ID of the new record back.
		id, err := app.snippets.Insert(title, content, expires)
		if err != nil {
			app.serverError(w, err)
			return
		}

		// Redirect the user to the relevant page for the snippet.
		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
func (app *application) serverError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	trace := string(debug.Stack())
	app.logger.Error(err.Error(), "trace", trace)
}
