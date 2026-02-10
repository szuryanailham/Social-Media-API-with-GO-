package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/szuryanailham/social/internal/env/store"
)

type Application struct {
	config 	Config
	store 	store.Storage
	
}

type Config struct {
	Addr string
	db 	dbConfig
	env string
}

type dbConfig struct {
	addr string 
	maxOpenConns int
	maxIdleConns int
	maxIdleTime time.Duration
}

func (app *Application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	 r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
	r.Get("/health", app.HealthCheckHandler)
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", app.CreatePostHandler)
		r.Route("/{postID}",func(r chi.Router) {
			r.Get("/", app.GetPostHandler)
			r.Delete("/", app.deletePostHandler)
		})
	})
	})
	return r
}

func (app *Application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr : app.config.Addr,
		Handler: mux,
		WriteTimeout:time.Second*30,
		ReadTimeout: time.Second*30,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started at %s",
	app.config.Addr)

	return srv.ListenAndServe()
}

