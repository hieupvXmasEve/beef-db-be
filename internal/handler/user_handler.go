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

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	fmt.Println("Login error:")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, 
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	resp, err := h.userService.Login(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, 
			model.NewErrorResponse("Login failed", err.Error()))
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(resp.User.ID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, 
			model.NewErrorResponse("Failed to generate token", err.Error()))
		return
	}

	// Set JWT cookie
	utils.SetJWTCookie(w, token)

	utils.SendResponse(w, http.StatusOK, 
		model.NewSuccessResponse("Login successful", resp))
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the JWT cookie
	utils.ClearJWTCookie(w)
	
	utils.SendResponse(w, http.StatusOK, 
		model.NewSuccessResponse("Successfully logged out", nil))
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUser")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, 
			model.NewErrorResponse("Invalid user ID", []model.ValidationError{
				model.NewValidationError("id", "Must be a valid number"),
			}))
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, 
			model.NewErrorResponse("User not found", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK, 
		model.NewSuccessResponse("User retrieved successfully", user))
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListUsers(r.Context())
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, 
			model.NewErrorResponse("Failed to retrieve users", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusOK, 
		model.NewSuccessResponse("Users retrieved successfully", users))
}

// SignUp handles user registration
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req model.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, 
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	user, err := h.userService.SignUp(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, 
			model.NewErrorResponse("Sign up failed", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated, 
		model.NewSuccessResponse("User created successfully", user))
} 