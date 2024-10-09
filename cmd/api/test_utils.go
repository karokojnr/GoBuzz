package main

import (
	"net/http"
	"net/http/httptest"
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

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}
