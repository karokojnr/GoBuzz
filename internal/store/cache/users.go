package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/karokojnr/GoBuzz/internal/store"
)

type UserStore struct {
	rdb *redis.Client
}

const UserExpTime = time.Minute

func (us *UserStore) Get(ctx context.Context, id int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user:%d", id)
	data, err := us.rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	var user store.User
	if data != "" {
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return &user, nil

}

func (us *UserStore) Set(ctx context.Context, u *store.User) error {
	cacheKey := fmt.Sprintf("user:%d", u.ID)
	json, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return us.rdb.Set(ctx, cacheKey, json, UserExpTime).Err()

}
