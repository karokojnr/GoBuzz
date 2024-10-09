package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, u *User, token string, invitationExp time.Duration) error {
	return nil
}

func (s *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (s *MockUserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	return nil, nil
}

func (s *MockUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {

	return nil, nil
}

func (s *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}
