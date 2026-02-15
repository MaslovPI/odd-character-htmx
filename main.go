package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/views"
)

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index(models.GetEmptyChar())).ServeHTTP(w, r)
}

func rollStats(w http.ResponseWriter, r *http.Request) {
	stats := models.RollStats()
	templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
}

func generateDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	maxStatStr := r.FormValue("max_stat")
	hpStr := r.FormValue("hp")

	maxStat, err := strconv.Atoi(maxStatStr)
	if err != nil {
		http.Error(w, "Invalid max_stat", http.StatusBadRequest)
		return
	}

	hp, err := strconv.Atoi(hpStr)
	if err != nil {
		http.Error(w, "Invalid hp", http.StatusBadRequest)
		return
	}

	starter := models.GenerateStarter(maxStat, hp)

	templ.Handler(views.Description(starter)).ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /rollstats", rollStats)
	mux.HandleFunc("POST /generatedescription", generateDescription)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	log.Println("http://localhost:42069 running...")
	err := http.ListenAndServe(":42069", middlewareLog(mux))
	log.Fatal(err)
}
