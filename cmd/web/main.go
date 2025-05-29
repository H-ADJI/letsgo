package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type app struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":3030", "Http network address to listen to")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	app := &app{logger: logger}

	logger.Info("Server listening", "port", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
