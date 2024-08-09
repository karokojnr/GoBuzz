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
	Version   int       `json:"version"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int `json:"comments_count"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, p *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
	SELECT id, content, title, user_id, tags, created_at, updated_at, version
	FROM posts
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var p Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Content,
		&p.Title,
		&p.UserID,
		pq.Array(&p.Tags),
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Version,
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

func (s *PostStore) GetUserFeed(ctx context.Context, id int64, pfq PaginatedFeedQuery) ([]PostWithMetadata, error) {
	query := `
	SELECT p.id, p.user_id, p.title, p.content, p.created_at, p.version, p.tags, u.username, COUNT(c.id) AS comments_count
	FROM posts p
	LEFT JOIN comments c ON c.post_id = p.id
	LEFT JOIN users u ON p.user_id = u.id
	JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
	WHERE f.user_id = $1 OR p.user_id = $1
	GROUP BY p.id, u.username
	ORDER BY p.created_at ` + pfq.Sort + `
	LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, id, pfq.Limit, pfq.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feed []PostWithMetadata
	for rows.Next() {
		var p PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			&p.Version,
			pq.Array(&p.Tags),
			&p.User.Username,
			&p.CommentsCount,
		)
		if err != nil {
			return nil, err
		}
		feed = append(feed, p)
	}

	return feed, nil
}

func (s *PostStore) Update(ctx context.Context, p *Post) error {
	query := `
	UPDATE posts
	SET title = $1, content = $2, version = version + 1
	WHERE id = $3 AND version = $4
	RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		p.Title,
		p.Content,
		p.ID,
		p.Version,
	).Scan(&p.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil

}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM posts
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
