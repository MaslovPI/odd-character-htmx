package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/maslovpi/odd-character-htmx/models"
	"github.com/maslovpi/odd-character-htmx/providers"
	"github.com/maslovpi/odd-character-htmx/views"
)

const characterCookieName = "character"

var tenYearsInSeconds = int((10 * 365 * 24 * time.Hour).Seconds())

func saveCharacterToCookie(w http.ResponseWriter, char models.Character) error {
	data, err := json.Marshal(char)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     characterCookieName,
		Value:    url.QueryEscape(string(data)),
		Path:     "/",
		MaxAge:   tenYearsInSeconds,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}

func getCharacterFromCookie(r *http.Request) (models.Character, bool) {
	cookie, err := r.Cookie(characterCookieName)
	if err != nil {
		return models.Character{}, false
	}
	decoded, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return models.Character{}, false
	}
	var char models.Character
	if err := json.Unmarshal([]byte(decoded), &char); err != nil {
		return models.Character{}, false
	}
	return char, true
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) error {
	char, ok := getCharacterFromCookie(r)
	if !ok {
		char = models.GetEmptyChar()
	}
	templ.Handler(views.Index(char)).ServeHTTP(w, r)
	return nil
}

func rollStats(w http.ResponseWriter, r *http.Request) error {
	stats, err := models.RollStats()
	if err != nil {
		return &appError{err, "failed to roll stats", http.StatusInternalServerError}
	}
	char, _ := getCharacterFromCookie(r)
	char.Stats = stats
	char.Description = models.Description{}
	if err := saveCharacterToCookie(w, char); err != nil {
		return &appError{err, "failed to save character cookie", http.StatusInternalServerError}
	}
	templ.Handler(views.Stats(stats)).ServeHTTP(w, r)
	return nil
}

func generateDescription(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return &appError{err, "failed to parse form", http.StatusBadRequest}
	}

	char, ok := getCharacterFromCookie(r)
	if !ok {
		maxStat, err := strconv.Atoi(r.FormValue("max_stat"))
		if err != nil {
			return &appError{err, "invalid max_stat", http.StatusBadRequest}
		}
		hp, err := strconv.Atoi(r.FormValue("hp"))
		if err != nil {
			return &appError{err, "invalid hp", http.StatusBadRequest}
		}
		char.Stats = models.Stats{Max: maxStat, HitProtection: hp}
	}

	description, err := starterProvider.GenerateStarter(char.Stats.HitProtection, char.Stats.Max)
	if err != nil {
		return &appError{err, "not able to generate starter", http.StatusInternalServerError}
	}

	char.Description = description
	if err := saveCharacterToCookie(w, char); err != nil {
		return &appError{err, "failed to save character cookie", http.StatusInternalServerError}
	}

	templ.Handler(views.Description(description)).ServeHTTP(w, r)
	return nil
}

var starterProvider providers.StarterProvider

func main() {
	var err error
	starterProvider, err = providers.InitStarterProvider()
	if err != nil {
		log.Fatalf("starter provider is not initialized: %v", err)
	}

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
