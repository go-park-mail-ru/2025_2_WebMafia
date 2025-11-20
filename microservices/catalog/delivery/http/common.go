package http

import (
	"net/http"
	"strconv"
)

const (
	defaultLimit       = 100
	defaultOffset      = 0
	maxLimit           = 1000
	queryParamLimit    = "limit"
	queryParamOffset   = "offset"
	queryParamSearch   = "q"
	defaultSearchLimit = 50
)

func parsePagination(r *http.Request) (uint64, uint64) {
	query := r.URL.Query()
	limitStr := query.Get(queryParamLimit)
	offsetStr := query.Get(queryParamOffset)

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		offset = defaultOffset
	}
	return limit, offset
}
