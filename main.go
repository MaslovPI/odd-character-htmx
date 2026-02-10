package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/functions"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func rollStat() int {
	result, err := functions.RollMultipleDice(3, 6)
	check(err)
	return result
}

func rollStats() models.Stats {
	return models.Stats{
		Strength:  rollStat(),
		Dexterity: rollStat(),
		Willpower: rollStat(),
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index(models.Stats{})).ServeHTTP(w, r)
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		stats := rollStats()
		templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
	})

	http.Handle("/css/",
		http.StripPrefix("/css/",
			http.FileServer(http.Dir("css"))))

	log.Println("http://localhost:42069 running...")
	http.ListenAndServe(":42069", nil)
}
