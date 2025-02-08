package model

import "time"

// Category represents a product category
type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateCategoryRequest represents the request to create a category
type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// UpdateCategoryRequest represents the request to update a category
type UpdateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// Product represents a product
type Product struct {
	ID           int       `json:"id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Description  string    `json:"description,omitempty"`
	Price        float64   `json:"price"`
	ImageURL     string    `json:"image_url,omitempty"`
	ThumbURL     string    `json:"thumb_url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
}

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	CategoryID  int     `json:"category_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Slug        string  `json:"slug" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	ImageURL    string  `json:"image_url"`
	ThumbURL    string  `json:"thumb_url"`
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	CategoryID  int     `json:"category_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Slug        string  `json:"slug" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	ImageURL    string  `json:"image_url"`
	ThumbURL    string  `json:"thumb_url"`
}
