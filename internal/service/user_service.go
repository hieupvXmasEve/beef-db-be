package service

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type UserService struct {
	queries *repository.Queries
	db      *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		queries: repository.New(db),
		db:      db,
	}
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: tokenString,
		User: model.User{
			GVA_MODEL: model.GVA_MODEL{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Email: user.Email,
			Role:  model.Role(user.Role),
		},
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &model.User{
		GVA_MODEL: model.GVA_MODEL{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Email: user.Email,
		Role:  model.Role(user.Role),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.User, len(users))
	for i, user := range users {
		result[i] = model.User{
			GVA_MODEL: model.GVA_MODEL{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Email: user.Email,
			Role:  model.Role(user.Role),
		}
	}

	return result, nil
}

// SignUp handles user registration with default user role
func (s *UserService) SignUp(ctx context.Context, req model.SignUpRequest) (*model.User, error) {
	// Check if email already exists
	_, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user with default role
	params := repository.CreateUserParams{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	result, err := s.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly created user
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Fetch the created user
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		GVA_MODEL: model.GVA_MODEL{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Email: user.Email,
		Role:  model.Role(user.Role),
	}, nil
} 