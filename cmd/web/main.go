package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sachin-sharma-IN/bookings/pkg/config"
	"github.com/sachin-sharma-IN/bookings/pkg/handlers"
	"github.com/sachin-sharma-IN/bookings/pkg/render"
)

// const can't be changed by anything else in the application.
const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Change this to True when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
