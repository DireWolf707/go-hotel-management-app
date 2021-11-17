package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/direwolf707/go-web-app/internal/config"
	"github.com/direwolf707/go-web-app/internal/handlers"
	"github.com/direwolf707/go-web-app/internal/render"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// initializing app level config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create template cache")
	}
	// initializing sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	//
	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false
	app.Session = session
	// pointing app level config to packages
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	// server initialization
	fmt.Println(fmt.Sprintf("Listening on port  %s", portNumber))
	s := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = s.ListenAndServe()
	log.Fatal(err)
}
