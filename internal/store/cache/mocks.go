package cache

import (
	"context"

	"github.com/karokojnr/GoBuzz/internal/store"
	"github.com/stretchr/testify/mock"
)

func NewMockCache() Cache {
	return Cache{
		Users: &MockUserCacheStore{},
	}
}

type MockUserCacheStore struct {
	mock.Mock
}

func (m *MockUserCacheStore) Get(ctx context.Context, id int64) (*store.User, error) {
	args := m.Called(id)
	return nil, args.Error(1)
}

func (m *MockUserCacheStore) Set(ctx context.Context, u *store.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserCacheStore) Delete(ctx context.Context, id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
