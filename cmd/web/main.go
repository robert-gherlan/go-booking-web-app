package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/robert-gherlan/go-booking-web-app/internal/config"
	"github.com/robert-gherlan/go-booking-web-app/internal/handlers"
	"github.com/robert-gherlan/go-booking-web-app/internal/helpers"
	"github.com/robert-gherlan/go-booking-web-app/internal/models"
	"github.com/robert-gherlan/go-booking-web-app/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the entry point that starts the web server on 8080 port.
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Start the web server
	log.Printf("Starting the web server on %s port.", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	// Setup loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Set session configuration
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// cookie persist after the browser window is closed by the end user
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
