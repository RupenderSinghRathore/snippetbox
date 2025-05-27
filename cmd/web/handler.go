package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Server", "Go")
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.internalSeverError(w, err)
		return
	}
	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		app.internalSeverError(w, err)
		return
	}
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.Error(w, "Snippet not found", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Showing snippet no. %v", id)
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
}
func (app *application) internalSeverError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	app.logger.Error(err.Error())
}
