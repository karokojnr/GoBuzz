package cache

import (
	"context"

	"github.com/karokojnr/GoBuzz/internal/store"
)

func NewMockCache() Cache {
	return Cache{
		Users: &MockUserCacheStore{},
	}
}

type MockUserCacheStore struct{}

func (m *MockUserCacheStore) Get(ctx context.Context, id int64) (*store.User, error) {
	return nil, nil
}

func (m *MockUserCacheStore) Set(ctx context.Context, u *store.User) error {
	return nil
}

func (m *MockUserCacheStore) Delete(ctx context.Context, id int64) error {
	return nil
}
