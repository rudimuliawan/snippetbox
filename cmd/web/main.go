package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	mux.HandleFunc("GET /snippet/create", snippetCreatePost)

	log.Println("Listening on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
