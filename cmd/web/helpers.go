package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		uri    = r.URL.RequestURI()
		method = r.Method
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
