package utils

import (
	"net/http"
	"strconv"

	"beef-db-be/internal/model"
)

// GetPaginationFromRequest extracts pagination parameters from request query
func GetPaginationFromRequest(r *http.Request) model.Pagination {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	return model.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}
