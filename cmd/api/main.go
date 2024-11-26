package main

import (
	"log"

	"github.com/Gulshan256/social/internal/db"
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

	db, err := db.NewDB(
		cnf.db.dsn,
		cnf.db.ssl,
		cnf.db.maxOpenConns,
		cnf.db.maxIdleConns,
		cnf.db.maxConnLifetime,
		cnf.db.maxIdleTime,
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Database connection successful")

	store := store.NewStorage(db)

	app := &appplication{
		config: cnf,
		store:  store,
	}

	mux := app.mount()

	if err := app.Start(mux); err != nil {
		panic(err)
	}

}
