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
	"beef-db-be/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database connection pool
	pool, err := config.NewDBPool()
	if err != nil {
		log.Fatalf("Failed to create database pool: %v", err)
	}
	defer pool.Close()

	// Initialize services
	userService := service.NewUserService(pool)
	categoryService := service.NewCategoryService(pool)
	productService := service.NewProductService(pool)
	healthHandler := handler.NewHealthHandler(pool)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

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
		r.Post("/auth/signup", userHandler.SignUp)
		r.Post("/auth/login", userHandler.Login)
		r.Post("/auth/logout", userHandler.Logout)
		r.Get("/users/me", userHandler.GetMe) // Public endpoint for getting current user

		// Public category routes
		r.Get("/categories", categoryHandler.ListCategories)
		r.Get("/categories/{id}", categoryHandler.GetCategory)
		r.Get("/categories/slug/{slug}", categoryHandler.GetCategoryBySlug)

		// Public product routes
		r.Get("/products", productHandler.ListProducts)
		r.Get("/products/{id}", productHandler.GetProduct)
		r.Get("/products/slug/{slug}", productHandler.GetProductBySlug)
		r.Get("/categories/{categoryId}/products", productHandler.ListProductsByCategory)
		r.Get("/categories/slug/{categorySlug}/products", productHandler.ListProductsByCategory)

		// Admin-only routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.RequireAuth(userService))

			// User management
			r.Get("/users/{id}", userHandler.GetUser)
			r.Get("/users", userHandler.ListUsers)

			// Category management
			r.Post("/categories", categoryHandler.CreateCategory)
			r.Put("/categories/{id}", categoryHandler.UpdateCategory)
			r.Delete("/categories/{id}", categoryHandler.DeleteCategory)

			// Product management
			r.Post("/products", productHandler.CreateProduct)
			r.Put("/products/{id}", productHandler.UpdateProduct)
			r.Delete("/products/{id}", productHandler.DeleteProduct)
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
