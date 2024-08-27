package store

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// model
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  password `json:"-"`
	CreatedAt string `json:"created_at"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, u *User) error {
	query := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		u.Username,
		u.Password,
		u.Email,
	).Scan(
		&u.ID,
		&u.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
	SELECT id, username, email, password, created_at
	FROM users
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	u := new(User)
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return u, nil
}
