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

type BlogPostHandler struct {
	service *service.BlogPostService
}

func NewBlogPostHandler(service *service.BlogPostService) *BlogPostHandler {
	return &BlogPostHandler{service: service}
}

func (h *BlogPostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBlogPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	post, err := h.service.Create(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to create blog post", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated,
		model.NewSuccessResponse("Blog post created successfully", post))
}

func (h *BlogPostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	post, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Blog post not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get blog post", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Blog post retrieved successfully", post))
}

func (h *BlogPostHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid slug", []model.ValidationError{
				model.NewValidationError("slug", "Slug is required"),
			}))
		return
	}

	post, err := h.service.GetBySlug(r.Context(), slug)
	if err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Blog post not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get blog post", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Blog post retrieved successfully", post))
}

func (h *BlogPostHandler) List(w http.ResponseWriter, r *http.Request) {
	pagination := utils.GetPaginationFromRequest(r)

	posts, totalCount, err := h.service.List(r.Context(), pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to list blog posts", err.Error()))
		return
	}

	paginatedResp := model.NewPaginatedResponse(posts, totalCount, pagination.Page, pagination.PageSize)
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Blog posts retrieved successfully", paginatedResp))
}

func (h *BlogPostHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	var req model.UpdateBlogPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Blog post not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to update blog post", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Blog post updated successfully", nil))
}

func (h *BlogPostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Blog post not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to delete blog post", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Blog post deleted successfully", nil))
}
