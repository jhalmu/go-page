package main

import (
	"log/slog"
	"net/http"
	"os"
	"os/exec"

	"github.com/a-h/templ"
	"github.com/jhalmu/go-page/internal/components"
	"github.com/jhalmu/go-page/internal/templates"
)

var Environment = "development"

func init() {
	os.Setenv("env", Environment)
	// run generate script
	exec.Command("make", "tailwind-build").Run()
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	router.Handle("/404", templ.Handler(templates.NotFound(), templ.WithStatus(http.StatusNotFound)))
	router.Handle("/500", templ.Handler(templates.InternalServerError(), templ.WithStatus(http.StatusInternalServerError)))
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		templates.Layout(components.Home("Jeremy Reindeer"), "Kotosivu").Render(r.Context(), w)

	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on port :8080")

	srv.ListenAndServe()

}
