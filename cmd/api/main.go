package main

import (
	"github.com/Gulshan256/social/internal/env"
	"github.com/Gulshan256/social/internal/store"
)

func main() {

	cnf := config{
		addr: env.GetString("API_ADDR", ":8080"),
		db: dbConfig{
			dsn:             env.GetString("DB_DSN", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable"),
			ssl:             env.GetInt("DB_SSL", 0) == 1,
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:     env.GetDuration("DB_MAX_IDLE_TIME", 60),
			maxConnLifetime: env.GetDuration("DB_MAX_CONN_LIFETIME", 60),
		},
	}
	store := store.NewStorage(nil)

	app := &appplication{
		config: cnf,
		store:  store,
	}

	mux := app.mount()

	if err := app.Start(mux); err != nil {
		panic(err)
	}

}
