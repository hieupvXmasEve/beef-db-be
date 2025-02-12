package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"beef-db-be/internal/model"
	"beef-db-be/internal/service"
	"beef-db-be/internal/utils"
)

// WebsiteSettingHandler handles HTTP requests for website settings
type WebsiteSettingHandler struct {
	service   *service.WebsiteSettingService
	validator *validator.Validate
}

// NewWebsiteSettingHandler creates a new website setting handler
func NewWebsiteSettingHandler(service *service.WebsiteSettingService) *WebsiteSettingHandler {
	return &WebsiteSettingHandler{
		service:   service,
		validator: validator.New(),
	}
}

// Create handles the creation of a new website setting
func (h *WebsiteSettingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateWebsiteSettingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request payload", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Validation failed", []model.ValidationError{
				model.NewValidationError("body", "Validation failed"),
			}))
		return
	}

	setting, err := h.service.Create(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to create setting", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated, model.APIResponse{
		Status:  "success",
		Message: "Setting created successfully",
		Data:    setting,
	})
}

// Get handles retrieving a website setting by ID
func (h *WebsiteSettingHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid setting ID", []model.ValidationError{
				model.NewValidationError("id", "Invalid setting ID"),
			}))
		return
	}

	setting, err := h.service.Get(r.Context(), int32(id))
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound,
			model.NewErrorResponse("Setting not found", []model.ValidationError{
				model.NewValidationError("id", "Setting not found"),
			}))
		return
	}

	utils.SendResponse(w, http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: "Setting retrieved successfully",
		Data:    setting,
	})
}

// GetByName handles retrieving a website setting by name
func (h *WebsiteSettingHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Setting name is required", []model.ValidationError{
				model.NewValidationError("name", "Setting name is required"),
			}))
		return
	}

	setting, err := h.service.GetByName(r.Context(), name)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound,
			model.NewErrorResponse("Setting not found", []model.ValidationError{
				model.NewValidationError("name", "Setting not found"),
			}))
		return
	}

	utils.SendResponse(w, http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: "Setting retrieved successfully",
		Data:    setting,
	})
}

// List handles retrieving all website settings
func (h *WebsiteSettingHandler) List(w http.ResponseWriter, r *http.Request) {
	settings, err := h.service.List(r.Context())
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve settings", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: "Settings retrieved successfully",
		Data:    settings,
	})
}

// Update handles updating a website setting
func (h *WebsiteSettingHandler) Update(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Setting name is required", []model.ValidationError{
				model.NewValidationError("name", "Setting name is required"),
			}))
		return
	}

	var req model.UpdateWebsiteSettingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request payload", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Validation failed", []model.ValidationError{
				model.NewValidationError("body", "Validation failed"),
			}))
		return
	}

	if err := h.service.Update(r.Context(), name, req); err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to update setting", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: "Setting updated successfully",
	})
}

// Delete handles deleting a website setting
func (h *WebsiteSettingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid setting ID", []model.ValidationError{
				model.NewValidationError("id", "Invalid setting ID"),
			}))
		return
	}

	if err := h.service.Delete(r.Context(), int32(id)); err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to delete setting", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: "Setting deleted successfully",
	})
}
