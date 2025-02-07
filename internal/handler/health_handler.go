package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"beef-db-be/internal/config"
)

type HealthHandler struct {
	db *sql.DB
}

type HealthResponse struct {
	Status   string           `json:"status"`
	Database *config.DBStatus `json:"database"`
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	dbStatus, err := config.CheckDBConnection(h.db)
	
	response := HealthResponse{
		Status:   "healthy",
		Database: dbStatus,
	}

	if err != nil {
		response.Status = "unhealthy"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
} 