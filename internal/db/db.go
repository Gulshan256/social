package db

import (
	"context"
	"database/sql"
	"time"
)

func NewDB(dsn string, ssl bool, maxOpenConns, maxIdleConns int, maxConnLifetime time.Duration, maxIdleTime time.Duration) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
