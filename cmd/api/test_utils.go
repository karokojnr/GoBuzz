package main

import (
	"testing"

	"github.com/karokojnr/GoBuzz/internal/store"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()

	return &application{
		logger: logger,
		store:  mockStore,
	}
}
