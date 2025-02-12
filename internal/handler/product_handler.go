package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"beef-db-be/internal/model"
	"beef-db-be/internal/service"
	"beef-db-be/internal/utils"
)

// ProductHandler handles HTTP requests related to products
type ProductHandler struct {
	productService *service.ProductService
	websiteService *service.WebsiteSettingService
}

// NewProductHandler creates a new ProductHandler instance
func NewProductHandler(productService *service.ProductService, websiteService *service.WebsiteSettingService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		websiteService: websiteService,
	}
}

// getPaginationFromRequest extracts pagination parameters from request query
func getPaginationFromRequest(r *http.Request) model.Pagination {
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
	pagination := getPaginationFromRequest(r)

	products, totalCount, err := h.productService.ListProducts(r.Context(), pagination)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve products", err.Error()))
		return
	}

	fmt.Println("totalCount", products[0])
	paginatedResp := model.NewPaginatedResponse(products, totalCount, pagination.Page, pagination.PageSize)
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", paginatedResp))
}

// ListProductsByCategory retrieves products by category
func (h *ProductHandler) ListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	pagination := getPaginationFromRequest(r)

	var categoryID int64
	var categorySlug string

	// Try to get category ID from URL
	if idStr := chi.URLParam(r, "categoryId"); idStr != "" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest,
				model.NewErrorResponse("Invalid category ID", err.Error()))
			return
		}
		categoryID = id
	} else {
		// If no ID, try to get slug
		categorySlug = chi.URLParam(r, "categorySlug")
		if categorySlug == "" {
			utils.SendResponse(w, http.StatusBadRequest,
				model.NewErrorResponse("Category ID or slug is required", nil))
			return
		}
	}

	products, totalCount, err := h.productService.ListProductsByCategory(r.Context(), categoryID, categorySlug, pagination)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve products", err.Error()))
		return
	}

	paginatedResp := model.NewPaginatedResponse(products, totalCount, pagination.Page, pagination.PageSize)
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", paginatedResp))
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

// ListProductsBySettingCategories handles the request to get products by category IDs from website settings
func (h *ProductHandler) ListProductsBySettingCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the website setting for show_product_category
	setting, err := h.websiteService.GetByName(ctx, "show_product_category")

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get website settings", err.Error()))
		return
	}

	// Parse category IDs from the setting value
	var categoryIDs []int
	if err := json.Unmarshal([]byte(setting.Value), &categoryIDs); err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Invalid category IDs in settings", err.Error()))
		return
	}

	// Get products by category IDs
	categories, err := h.productService.GetProductsByCategoryIDs(ctx, categoryIDs)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get products", err.Error()))
		return
	}
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", categories))
}
