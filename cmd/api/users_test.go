package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {

	app := newTestApplication(t)
	mux := app.mount()

	mockToken := "valid_token"

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

	t.Run("shoul allow authenticated requests", func(t *testing.T) {
		// set up
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		// action
		req.Header.Set("Authorization", "Bearer "+mockToken)

		rr := executeRequest(req, mux)

		// assert
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}
