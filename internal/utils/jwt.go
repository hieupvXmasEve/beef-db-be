package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// TokenCookieName is the name of the cookie that stores the JWT token
	TokenCookieName = "auth_token"
	// TokenExpiry is the duration for which the token is valid
	TokenExpiry = 24 * time.Hour
)

// Claims represents the JWT claims structure
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// isProduction returns true if the application is running in production mode
func isProduction() bool {
	return strings.Contains(os.Getenv("ALLOWED_ORIGINS"), "hieupv.site")
}

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(userID int64) (string, error) {
	// Get JWT secret from environment variable
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable is not set")
	}

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// SetJWTCookie sets the JWT token as an HTTP-only cookie
func SetJWTCookie(w http.ResponseWriter, token string) {
	isProd := isProduction()

	http.SetCookie(w, &http.Cookie{
		Name:     TokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   isProd,                // Enable Secure flag in production (HTTPS)
		SameSite: http.SameSiteNoneMode, // Use Strict in production, None in development
		Domain:   getDomain(isProd),     // Set domain in production
		MaxAge:   int(TokenExpiry.Seconds()),
	})
}

// getDomain returns the appropriate domain based on the environment
func getDomain(isProd bool) string {
	if isProd {
		return "hieupv.site" // Production domain
	}
	return "" // Empty for localhost
}

// ClearJWTCookie removes the JWT cookie
func ClearJWTCookie(w http.ResponseWriter) {
	isProd := isProduction()

	http.SetCookie(w, &http.Cookie{
		Name:     TokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteNoneMode,
		Domain:   getDomain(isProd),
		MaxAge:   -1,
	})
}

// ValidateJWT validates the JWT token and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
