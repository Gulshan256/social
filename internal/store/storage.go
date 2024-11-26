package store

import (
	"context"
	"database/sql"

	"github.com/Gulshan256/social/models"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
	}

	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
