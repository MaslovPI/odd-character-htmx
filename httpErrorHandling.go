package main

import (
	"errors"
	"log/slog"
	"net/http"
)

type appError struct {
	cause   error
	message string
	code    int
}

func (e *appError) Error() string { return e.message }

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		var ae *appError
		if errors.As(err, &ae) {
			slog.Error(ae.message, "err", ae.cause)
			http.Error(w, ae.message, ae.code)
		} else {
			slog.Error("internal error", "err", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
