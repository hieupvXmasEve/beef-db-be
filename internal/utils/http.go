package utils

import (
	"encoding/json"
	"net/http"

	"beef-db-be/internal/model"
)

// SendResponse is a helper function to send standardized JSON responses
func SendResponse(w http.ResponseWriter, status int, response model.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If encoding fails, try to send a simple error response
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.NewErrorResponse("Failed to encode response", err.Error()))
	}
} 