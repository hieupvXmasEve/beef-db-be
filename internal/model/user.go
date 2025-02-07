package model

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// User represents a user in the system
type User struct {
	GVA_MODEL
	Email    string `json:"email"`
	Role     Role   `json:"role"`
}

// SignUpRequest represents the request body for user registration
type SignUpRequest struct {
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents the response for a successful login
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
} 