package main

import (
	"database/sql"
	"flag"
	"fmt"

	"log/slog"
	"net/http"
	"os"

	"github.com/RupenderSinghRathore/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   slog.Logger
	snippets *models.SnippetModel
}

func main() {
	port := flag.String("p", ":8080", "Port to connect to")
	dsn := flag.String("dsn", "furry:touka@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := dataBaseConn(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	app := application{
		logger:   *logger,
		snippets: &models.SnippetModel{DB: db},
	}
	app.logger.Info("Server starting..", "port", *port)
	if err := http.ListenAndServe(*port, app.serverRouter()); err != nil {
		app.logger.Error(err.Error())
	}
}

func dataBaseConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, nil
}
