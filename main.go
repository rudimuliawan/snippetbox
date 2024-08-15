package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	_, err = w.Write([]byte(msg))
	if err != nil {
		return
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Snippet Create!"))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Listening on port 8080")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
