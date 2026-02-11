package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func home(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index(models.Stats{})).ServeHTTP(w, r)
}

func roll(w http.ResponseWriter, r *http.Request) {
	stats := models.RollStats()
	templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /roll", roll)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	log.Println("http://localhost:42069 running...")
	err := http.ListenAndServe(":42069", mux)
	log.Fatal(err)
}
