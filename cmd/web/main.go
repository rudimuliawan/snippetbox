package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/rudimuliawan/snippetbox/internal/models"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippedModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(-1)
	}

	app := &application{
		logger:        logger,
		snippets:      &models.SnippedModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("Starting server", slog.Any("address", *addr))

	err = http.ListenAndServe(*addr, app.Routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
