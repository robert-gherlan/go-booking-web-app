package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/robert-gherlan/go-webapp/pkg/config"
	"github.com/robert-gherlan/go-webapp/pkg/handlers"
	"github.com/robert-gherlan/go-webapp/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the entry point that starts the web server on 8080 port.
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Init the template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Cannot create the template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	// Init the handlers repo
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

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

	// Change this to true when in production
	app.InProduction = false

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

	return nil
}
