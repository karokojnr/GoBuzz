package main

import (
	"testing"

	"github.com/karokojnr/GoBuzz/internal/store"
	"github.com/karokojnr/GoBuzz/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockCache()

	return &application{
		logger:     logger,
		store:      mockStore,
		cacheStore: mockCacheStore,
	}
}
