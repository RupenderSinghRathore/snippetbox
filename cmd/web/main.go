package main

import (
	"database/sql"
	"flag"
	"html/template"
	"runtime/debug"

	"log/slog"
	"net/http"
	"os"

	"github.com/RupenderSinghRathore/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	port := flag.String("p", ":8080", "Port to connect to")
	dsn := flag.String("dsn", "furry:touka@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := dataBaseConn(*dsn)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
	defer db.Close()
	templateCache, err := newTemplateCache()
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
	app := application{
		logger:        *logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}
	app.logger.Info("Server starting..", "port", *port)
	if err := http.ListenAndServe(*port, app.serverRouter()); err != nil {
		trace := string(debug.Stack())
		app.logger.Error(err.Error(), "trace", trace)
	}
}

func dataBaseConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
