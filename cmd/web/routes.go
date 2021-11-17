package main

import (
	"net/http"

	"github.com/direwolf707/go-web-app/pkg/config"
	"github.com/direwolf707/go-web-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// returns mux
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//urls
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	//static file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
