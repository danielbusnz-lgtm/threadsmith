package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	addr := ":8080"
	log.Info("threadsmith gateway listening", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Error("server failed", "err", err)
		os.Exit(1)
	}
}
