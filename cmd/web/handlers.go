package main

import (
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create a new snippet!"))
	if err != nil {
		return
	}
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create a new snippet!"))
}
