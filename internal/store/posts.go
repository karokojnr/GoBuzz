package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// model
type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, p *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		p.Content,
		p.Title,
		p.UserID,
		pq.Array(p.Tags),
	).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
	SELECT id, content, title, user_id, tags, created_at, updated_at
	FROM posts
	WHERE id = $1
	`

	var p Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Content,
		&p.Title,
		&p.UserID,
		pq.Array(&p.Tags),
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &p, nil
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM posts
	WHERE id = $1
	`

	res, err := s.db.ExecContext(ctx, query, id) // execcontext does not return rows
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
