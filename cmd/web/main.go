package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger slog.Logger
}

func main() {
	port := flag.String("p", ":8080", "Port to connect to")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := application{logger: *logger}
	fmt.Println("Starting the server at port:", *port)
	if err := http.ListenAndServe(*port, app.serverRouter()); err != nil {
		app.logger.Error(err.Error())
	}
}
