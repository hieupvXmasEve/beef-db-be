package main

import (
	"fmt"
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
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local" // Mặc định là local nếu không có biến môi trường
	}
	if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
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
	websiteSettingService := service.NewWebsiteSettingService(pool)
	healthHandler := handler.NewHealthHandler(pool)
	pageService := service.NewPageService(pool)
	blogPostService := service.NewBlogPostService(pool)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService, websiteSettingService, categoryService)
	websiteSettingHandler := handler.NewWebsiteSettingHandler(websiteSettingService)
	pageHandler := handler.NewPageHandler(pageService)
	blogPostHandler := handler.NewBlogPostHandler(blogPostService)

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
		r.Get("/products/by-setting-categories", productHandler.ListProductsBySettingCategories)
		r.Get("/categories/{categoryId}/products", productHandler.ListProductsByCategoryByID)
		r.Get("/categories/slug/{categorySlug}/products", productHandler.ListProductsByCategoryBySlug)

		// Public page routes
		r.Get("/pages", pageHandler.ListPages)
		r.Get("/pages/{id}", pageHandler.GetPage)
		r.Get("/pages/slug/{slug}", pageHandler.GetPageBySlug)

		// Public blog post routes
		r.Get("/blog-posts", blogPostHandler.List)
		r.Get("/blog-posts/{id}", blogPostHandler.GetByID)
		r.Get("/blog-posts/slug/{slug}", blogPostHandler.GetBySlug)

		// Public website settings routes
		r.Get("/settings", websiteSettingHandler.List)
		r.Get("/settings/{id}", websiteSettingHandler.Get)
		r.Get("/settings/name/{name}", websiteSettingHandler.GetByName)

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

			// Website settings management
			r.Post("/settings", websiteSettingHandler.Create)
			r.Put("/settings/name/{name}", websiteSettingHandler.Update)
			r.Delete("/settings/{id}", websiteSettingHandler.Delete)

			// Page management
			r.Post("/pages", pageHandler.CreatePage)
			r.Put("/pages/{id}", pageHandler.UpdatePage)
			r.Delete("/pages/{id}", pageHandler.DeletePage)

			// Blog post management
			r.Post("/blog-posts", blogPostHandler.Create)
			r.Put("/blog-posts/{id}", blogPostHandler.Update)
			r.Delete("/blog-posts/{id}", blogPostHandler.Delete)
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
