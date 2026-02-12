package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func home(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index(models.GetEmptyChar())).ServeHTTP(w, r)
}

func rollStats(w http.ResponseWriter, r *http.Request) {
	stats := models.RollStats()
	templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
}

func generateDescription(w http.ResponseWriter, r *http.Request) {
	starter := models.GenerateStarter()
	templ.Handler(views.Description(starter)).ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /rollstats", rollStats)
	mux.HandleFunc("GET /generatedescription", generateDescription)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	log.Println("http://localhost:42069 running...")
	err := http.ListenAndServe(":42069", mux)
	log.Fatal(err)
}
