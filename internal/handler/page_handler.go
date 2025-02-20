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

type PageHandler struct {
	pageService *service.PageService
}

func NewPageHandler(pageService *service.PageService) *PageHandler {
	return &PageHandler{
		pageService: pageService,
	}
}

func (h *PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	var req model.CreatePageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	page, err := h.pageService.CreatePage(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to create page", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated,
		model.NewSuccessResponse("Page created successfully", page))
}

func (h *PageHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid page ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	page, err := h.pageService.GetPage(r.Context(), int32(id))
	if err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Page not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get page", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Page retrieved successfully", page))
}

func (h *PageHandler) GetPageBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	page, err := h.pageService.GetPageBySlug(r.Context(), slug)
	if err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Page not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to get page", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Page retrieved successfully", page))
}

func (h *PageHandler) ListPages(w http.ResponseWriter, r *http.Request) {
	pagination := utils.GetPaginationFromRequest(r)

	pages, err := h.pageService.ListPages(r.Context(), int32(pagination.PageSize), int32((pagination.Page-1)*pagination.PageSize))
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to list pages", err.Error()))
		return
	}

	// TODO: Get total count from service
	totalCount := int64(len(pages))
	paginatedResp := model.NewPaginatedResponse(pages, totalCount, pagination.Page, pagination.PageSize)
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Pages retrieved successfully", paginatedResp))
}

func (h *PageHandler) UpdatePage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid page ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	var req model.UpdatePageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	if err := h.pageService.UpdatePage(r.Context(), int32(id), req); err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Page not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to update page", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Page updated successfully", nil))
}

func (h *PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid page ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	if err := h.pageService.DeletePage(r.Context(), int32(id)); err != nil {
		if err == service.ErrNotFound {
			utils.SendResponse(w, http.StatusNotFound,
				model.NewErrorResponse("Page not found", err.Error()))
			return
		}
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to delete page", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Page deleted successfully", nil))
}
