package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karokojnr/GoBuzz/internal/auth"
	"github.com/karokojnr/GoBuzz/internal/store"
	"github.com/karokojnr/GoBuzz/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockCache()
	mockAuthenticator := &auth.MockAuthenticator{}

	return &application{
		logger:        logger,
		store:         mockStore,
		cacheStore:    mockCacheStore,
		authenticator: mockAuthenticator,
		config:        cfg,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected status %d, got %d", expected, actual)
	}
}
