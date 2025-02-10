package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type UserService struct {
	queries *repository.Queries
	pool    *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		queries: repository.New(pool),
		pool:    pool,
	}
}

func (s *UserService) SignUp(ctx context.Context, req model.SignUpRequest) (*model.User, error) {
	// Check if user already exists
	_, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	result, err := s.queries.CreateUser(ctx, repository.CreateUserParams{
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return s.GetUser(ctx, result)
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	// Get user by email
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
		User: model.User{
			GVA_MODEL: model.GVA_MODEL{
				ID: user.ID,
			},
			Email: user.Email,
			Role:  model.Role(user.Role),
		},
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &model.User{
		GVA_MODEL: model.GVA_MODEL{
			ID: user.ID,
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
				ID: user.ID,
			},
			Email: user.Email,
			Role:  model.Role(user.Role),
		}
	}

	return result, nil
}

func generateJWT(userID int64) (string, error) {
	// Get JWT secret from environment variable
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	// Get expiry hours from environment variable
	expiryHours := 24 // Default to 24 hours
	if envExpiry := os.Getenv("JWT_EXPIRY_HOURS"); envExpiry != "" {
		if parsed, err := time.ParseDuration(envExpiry + "h"); err == nil {
			expiryHours = int(parsed.Hours())
		}
	}

	// Create claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(expiryHours) * time.Hour).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(secret))
}
