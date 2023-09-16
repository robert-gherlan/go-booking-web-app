package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/robert-gherlan/go-booking-web-app/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error("Not a *chi.Mux")
	}
}
