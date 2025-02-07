package model

import "time"

// GVA_MODEL contains common fields for all models
type GVA_MODEL struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIResponse represents the standard API response format
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// ValidationError represents a field-level validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, errors interface{}) APIResponse {
	return APIResponse{
		Status:  "error",
		Message: message,
		Errors:  errors,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}
