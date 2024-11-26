package store

import (
	"context"
	"database/sql"

	"github.com/Gulshan256/social/models"
	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {

	query := `INSERT INTO posts (title, content ,user_id, tag, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	error := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.UserID,
		pq.Array(post.Tags),
		post.CreatedAt,
	).Scan(&post.ID, &post.CreatedAt)

	if error != nil {
		return error
	}

	return nil
}
