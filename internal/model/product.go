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

// Product represents a product in the system
type Product struct {
	ID                int       `json:"id"`
	CategoryID        int       `json:"category_id"`
	Name              string    `json:"name"`
	Slug              string    `json:"slug"`
	Description       string    `json:"description"`
	Price             float64   `json:"price"`
	PriceSale         float64   `json:"price_sale,omitempty"`
	ImageURL          string    `json:"image_url"`
	ThumbURL          string    `json:"thumb_url"`
	CreatedAt         time.Time `json:"created_at"`
	CategoryName      string    `json:"category_name"`
	CategorySlug      string    `json:"category_slug"`
	UnitOfMeasurement string    `json:"unit_of_measurement"`
}

// CreateProductRequest represents the request body for product creation
type CreateProductRequest struct {
	CategoryID        int     `json:"category_id" validate:"required"`
	Name              string  `json:"name" validate:"required"`
	Slug              string  `json:"slug" validate:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price" validate:"required,gt=0"`
	PriceSale         float64 `json:"price_sale,omitempty" validate:"omitempty,gtefield=0"`
	ImageURL          string  `json:"image_url"`
	ThumbURL          string  `json:"thumb_url"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
}

// UpdateProductRequest represents the request body for product update
type UpdateProductRequest struct {
	CategoryID        int     `json:"category_id" validate:"required"`
	Name              string  `json:"name" validate:"required"`
	Slug              string  `json:"slug" validate:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price" validate:"required,gt=0"`
	PriceSale         float64 `json:"price_sale,omitempty" validate:"omitempty,gtefield=0"`
	ImageURL          string  `json:"image_url"`
	ThumbURL          string  `json:"thumb_url"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
}

// CategoryProductsResponse represents a category with its products
type CategoryProductsResponse struct {
	Name     string    `json:"name"`
	ImageURL string    `json:"image_url"`
	Slug     string    `json:"slug"`
	Products []Product `json:"products"`
}

// CategoryProductsListResponse represents a list of categories with their products
type CategoryProductsListResponse struct {
	Categories []CategoryProductsResponse `json:"categories"`
}

// CategoryWithProductsResponse represents a category with its paginated products
type CategoryWithProductsResponse struct {
	Category   Category          `json:"category"`
	Products   []Product         `json:"products"`
	Pagination PaginatedResponse `json:"pagination"`
}
