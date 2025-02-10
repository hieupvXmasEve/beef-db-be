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

// Pagination represents pagination parameters
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"total_items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(items interface{}, totalItems int64, page, pageSize int) PaginatedResponse {
	totalPages := (int(totalItems) + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return PaginatedResponse{
		Items:      items,
		TotalItems: totalItems,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// GetOffset returns the offset for SQL queries based on page and page size
func (p *Pagination) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10 // Default page size
	}
	return (p.Page - 1) * p.PageSize
}

// GetLimit returns the limit for SQL queries
func (p *Pagination) GetLimit() int {
	if p.PageSize < 1 {
		p.PageSize = 10 // Default page size
	}
	if p.PageSize > 100 {
		p.PageSize = 100 // Maximum page size
	}
	return p.PageSize
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
