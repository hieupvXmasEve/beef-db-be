package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/utils"
)

type HealthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{
		pool: pool,
	}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	// Check database connection
	if err := h.pool.Ping(r.Context()); err != nil {
		utils.SendResponse(w, http.StatusServiceUnavailable,
			model.NewErrorResponse("Database connection failed", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Service is healthy", nil))
}
