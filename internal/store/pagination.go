package store

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit", validate:"required,gte=1,lte=20"`
	Offset int    `json:"offset", validate:"required,gte=0"`
	Sort   string `json:"sort", validate:"oneof=asc desc"`
}

func (pfq PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return pfq, err
		}
		pfq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return pfq, err
		}
		pfq.Offset = l
	}

	sort := qs.Get("sort")
	if sort != "" {
		pfq.Sort = sort
	}

	return pfq, nil
}
