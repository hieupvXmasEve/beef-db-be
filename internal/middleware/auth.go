package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"beef-db-be/internal/model"
	"beef-db-be/internal/service"
	"beef-db-be/internal/utils"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
	UserIDKey      contextKey = "user_id"
)

func AuthMiddleware(userService *service.UserService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID, ok := claims["user_id"].(float64)
				if !ok {
					http.Error(w, "Invalid token claims", http.StatusUnauthorized)
					return
				}

				user, err := userService.GetUser(r.Context(), int64(userID))
				if err != nil {
					http.Error(w, "User not found", http.StatusUnauthorized)
					return
				}

				ctx := context.WithValue(r.Context(), UserContextKey, user)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		})
	}
}

func RequireRole(roles ...model.Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(UserContextKey).(*model.User)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			for _, role := range roles {
				if user.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}

// RequireAuth is a middleware that checks for a valid JWT token in cookies and ensures admin role
func RequireAuth(userService *service.UserService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(utils.TokenCookieName)
			fmt.Println("err", err)
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized,
					model.NewErrorResponse("Authentication required", "No authentication token provided"))
				return
			}

			claims, err := utils.ValidateJWT(cookie.Value)
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized,
					model.NewErrorResponse("Authentication failed", "Invalid or expired token"))
				return
			}

			// Get user from database to check role
			user, err := userService.GetUser(r.Context(), claims.UserID)
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized,
					model.NewErrorResponse("Authentication failed", "User not found"))
				return
			}

			// Check if user is admin
			if user.Role != model.RoleAdmin {
				utils.SendResponse(w, http.StatusForbidden,
					model.NewErrorResponse("Access denied", "Admin role required"))
				return
			}

			// Add user ID and user object to request context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID retrieves the user ID from the request context
func GetUserID(r *http.Request) (int64, bool) {
	userID, ok := r.Context().Value(UserIDKey).(int64)
	return userID, ok
}
