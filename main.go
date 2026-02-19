package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) error {
	templ.Handler(views.Index(models.GetEmptyChar())).ServeHTTP(w, r)
	return nil
}

func rollStats(w http.ResponseWriter, r *http.Request) error {
	templ.Handler(views.Stats(models.RollStats())).ServeHTTP(w, r)
	return nil
}

func generateDescription(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return &appError{err, "failed to parse form", http.StatusBadRequest}
	}

	maxStat, err := strconv.Atoi(r.FormValue("max_stat"))
	if err != nil {
		return &appError{err, "invalid max_stat", http.StatusBadRequest}
	}

	hp, err := strconv.Atoi(r.FormValue("hp"))
	if err != nil {
		return &appError{err, "invalid hp", http.StatusBadRequest}
	}

	templ.Handler(views.Description(models.GenerateStarter(maxStat, hp))).ServeHTTP(w, r)
	return nil
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", appHandler(home))
	mux.Handle("GET /rollstats", appHandler(rollStats))
	mux.Handle("POST /generatedescription", appHandler(generateDescription))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	slog.Info("server starting", "addr", "http://localhost:42069")
	if err := http.ListenAndServe(":42069", middlewareLog(mux)); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
