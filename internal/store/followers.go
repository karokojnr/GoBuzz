package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowersStore struct {
	db *sql.DB
}

func (fs *FollowersStore) Follow(ctx context.Context, followerId, userId int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, `
		INSERT INTO followers (user_id, follower_id)
		VALUES ($1, $2)
	`, followerId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FollowersStore) Unfollow(ctx context.Context, followerId, userId int64) error {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, `
		DELETE FROM followers
		WHERE user_id = $1 AND follower_id = $2
	`, userId, followerId)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}

	return nil
}
