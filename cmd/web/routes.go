package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)
	mux.HandleFunc("GET /snippet/create", app.snippetCreatePost)

	return mux
}
