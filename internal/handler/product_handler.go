package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"beef-db-be/internal/model"
	"beef-db-be/internal/service"
	"beef-db-be/internal/utils"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct handles product creation
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req model.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	product, err := h.productService.CreateProduct(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Failed to create product", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated,
		model.NewSuccessResponse("Product created successfully", product))
}

// GetProduct retrieves a product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid product ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	product, err := h.productService.GetProduct(r.Context(), id)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound,
			model.NewErrorResponse("Product not found", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Product retrieved successfully", product))
}

// GetProductBySlug retrieves a product by slug
func (h *ProductHandler) GetProductBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	product, err := h.productService.GetProductBySlug(r.Context(), slug)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound,
			model.NewErrorResponse("Product not found", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Product retrieved successfully", product))
}

// ListProducts retrieves all products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.ListProducts(r.Context())
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve products", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", products))
}

// ListProductsByCategory retrieves products by category
func (h *ProductHandler) ListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := chi.URLParam(r, "categoryId")
	categoryID, _ := strconv.Atoi(categoryIDStr) // It's ok if this fails, we'll try slug
	categorySlug := chi.URLParam(r, "categorySlug")

	if categoryID == 0 && categorySlug == "" {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid category", "Must provide either category ID or slug"))
		return
	}

	products, err := h.productService.ListProductsByCategory(r.Context(), categoryID, categorySlug)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve products", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", products))
}

// UpdateProduct updates a product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid product ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	var req model.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	product, err := h.productService.UpdateProduct(r.Context(), id, req)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Failed to update product", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Product updated successfully", product))
}

// DeleteProduct deletes a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid product ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	if err := h.productService.DeleteProduct(r.Context(), id); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Failed to delete product", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Product deleted successfully", nil))
}
