package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Gulshan256/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type appplication struct {
	config config
	store  store.Storage
}

type config struct {
	db   dbConfig
	addr string
}

type dbConfig struct {
	dsn             string
	ssl             bool
	maxOpenConns    int
	maxIdleConns    int
	maxConnLifetime time.Duration
	maxIdleTime     time.Duration
}

func (a *appplication) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)
	})
	return r
}

func (a *appplication) Start(mux http.Handler) error {

	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Println("API server started on %s", srv.Addr)
	fmt.Println("Press CTRL+C to stop the server")
	return srv.ListenAndServe()

}
