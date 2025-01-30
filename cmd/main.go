package main

import (
	"github.com/jhalmu/go-page/internal/components"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charaset utf-8")
		components.Home("Lucky me!").Render(r.Context(), w)
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on port :8080")

	srv.ListenAndServe()

}
