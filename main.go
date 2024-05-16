package main

import (
	"log/slog"
	"net/http"

	"github.com/sebar-io/go-websocket/pkg/ws"
)

func main() {
	s := ws.NewServer()
	http.Handle("/", s.NewServeMux())

	slog.Info("go-websocket started at :8080")
	http.ListenAndServe(":8080", nil)
}
