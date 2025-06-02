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
	// r.Header.Add("Server", "Go")
	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.internalSeverError(w, err)
	// 	return
	// }
	// if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
	// 	app.internalSeverError(w, err)
	// 	return
	// }
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.internalSeverError(w, err)
		return
	}
	fmt.Fprintf(w, "%+d", snippets)
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.Error(w, "Snippet not found", http.StatusBadRequest)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.internalSeverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", s)
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
		app.internalSeverError(w, err)
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := 7

		// Pass the data to the SnippetModel.Insert() method, receiving the
		// ID of the new record back.
		id, err := app.snippets.Insert(title, content, expires)
		if err != nil {
			app.internalSeverError(w, err)
			return
		}

		// Redirect the user to the relevant page for the snippet.
		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
func (app *application) internalSeverError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	trace := string(debug.Stack())
	app.logger.Error(err.Error(), "trace", trace)
}
