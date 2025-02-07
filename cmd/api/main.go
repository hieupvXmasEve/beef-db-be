package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"beef-db-be/internal/config"
	"beef-db-be/internal/handler"
	"beef-db-be/internal/middleware"
	"beef-db-be/internal/model"
	"beef-db-be/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database connection
	db, err := config.NewDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize services
	userService := service.NewUserService(db)
	healthHandler := handler.NewHealthHandler(db)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Initialize router
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(middleware.CORS)

	// Health check endpoint
	r.Get("/health", healthHandler.CheckHealth)

	// Routes
	r.Route("/api", func(r chi.Router) {
		// Public routes
		r.Post("/signup", userHandler.SignUp)
		r.Post("/login", userHandler.Login)
		r.Post("/logout", userHandler.Logout)

		// Protected routes
		r.Group(func(r chi.Router) {
			// Apply authentication middleware
			// r.Use(middleware.AuthMiddleware(userService))
			r.Use(middleware.RequireAuth)

			// Routes accessible by all authenticated users
			r.Get("/users/{id}", userHandler.GetUser)

			// Admin-only routes
			r.Group(func(r chi.Router) {
				r.Use(middleware.RequireRole(model.RoleAdmin))
				r.Get("/users", userHandler.ListUsers)
			})
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
