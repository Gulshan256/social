package store

import (
	"context"
	"database/sql"

	"github.com/Gulshan256/social/models"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	// Implement user creation logic here
	query := `INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	error := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	).Scan(&user.ID, &user.CreatedAt)

	if error != nil {
		return error
	}

	return nil

}
