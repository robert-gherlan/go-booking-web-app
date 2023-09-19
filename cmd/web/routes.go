package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/robert-gherlan/go-booking-web-app/internal/config"
	"github.com/robert-gherlan/go-booking-web-app/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middleware usage
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Register the basic routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Majors)

	// Search Availability Routes
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	// Choose Room
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room/{id}", handlers.Repo.BookRoom)

	// Reservation Summary Routes
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// Make Reservation Routes
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	// Enable reading from the static folder
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
