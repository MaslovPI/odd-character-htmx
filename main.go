package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index(models.Stats{})).ServeHTTP(w, r)
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		stats := models.RollStats()
		templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
	})

	http.Handle("/css/",
		http.StripPrefix("/css/",
			http.FileServer(http.Dir("css"))))

	log.Println("http://localhost:42069 running...")
	http.ListenAndServe(":42069", nil)
}
