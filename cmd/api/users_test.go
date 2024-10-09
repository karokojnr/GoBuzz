package main

import (
	"net/http"
	"testing"

	"github.com/karokojnr/GoBuzz/internal/store/cache"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {

	withRedis := config{
		redisCfg: redisConfig{
			enabled: true,
		},
	}

	app := newTestApplication(t, withRedis)
	mux := app.mount()

	mockToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		// set up
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		// action
		rr := executeRequest(req, mux)

		// assert
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)

	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		// set up
		mockCacheStore := app.cacheStore.Users.(*cache.MockUserCacheStore)

		mockCacheStore.On("Get", int64(1)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		// action
		req.Header.Set("Authorization", "Bearer "+mockToken)

		rr := executeRequest(req, mux)

		// assert
		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("should hit the cache first and if not exists it sets the user in the cache", func(t *testing.T) {
		mockCacheStore := app.cacheStore.Users.(*cache.MockUserCacheStore)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+mockToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)
		mockCacheStore.Calls = nil // Rest the mock expectations
	})

	t.Run("should NOT hit the cache if it is not enabled", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisConfig{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()

		mockCacheStore := app.cacheStore.Users.(*cache.MockUserCacheStore)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+mockToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNotCalled(t, "Get")

		mockCacheStore.Calls = nil
	})
}
