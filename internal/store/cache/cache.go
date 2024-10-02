package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/karokojnr/GoBuzz/internal/store"
)

type Cache struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
		Delete(context.Context, int64) error
	}
}

func NewRedisStorage(rdb *redis.Client) Cache {
	return Cache{
		Users: &UserStore{rdb: rdb},
	}
}
