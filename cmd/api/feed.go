package main

import (
	"net/http"

	"github.com/karokojnr/GoBuzz/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	pfq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	pfq, err := pfq.Parse(r)
	if err != nil {
		app.badRequestRepsonse(w, r, err)
		return
	}

	if err := Validate.Struct(pfq); err != nil {
		app.badRequestRepsonse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(508), pfq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
