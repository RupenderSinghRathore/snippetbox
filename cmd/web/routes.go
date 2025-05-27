package main

import "net/http"

func (app *application) serverRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.view)
	mux.HandleFunc("GET /snippet/create", app.create)
	mux.HandleFunc("POST /snippet/create", app.createPost)

	return mux
}
