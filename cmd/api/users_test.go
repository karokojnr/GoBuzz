package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {

	app := newTestApplication(t)
	mux := app.mount()

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
	})
}
