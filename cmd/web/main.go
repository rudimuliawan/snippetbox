package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	log.Printf("Starting server on %s\n", *addr)

	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
