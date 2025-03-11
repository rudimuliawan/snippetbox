package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	app := &application{logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}

	app.logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		app.logger.Error(err.Error())
	}
}
