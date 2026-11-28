package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rudimuliawan/snippetbox/internal/models"
	"github.com/rudimuliawan/snippetbox/internal/validators"
)

type SnippetCreateForm struct {
	Title                string `form:"title"`
	Content              string `form:"content"`
	Expires              int    `form:"expires"`
	validators.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets
	err = app.render(w, r, http.StatusOK, "home.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			http.NotFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet
	err = app.render(w, r, http.StatusOK, "view.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = SnippetCreateForm{
		Expires: 365,
	}

	err := app.render(w, r, http.StatusOK, "create.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form SnippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validators.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validators.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters")
	form.CheckField(validators.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validators.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.IsValid() {
		data := app.newTemplateData(r)
		data.Form = form

		err = app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
