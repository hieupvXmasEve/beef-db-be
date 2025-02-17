package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type ProductService struct {
	queries *repository.Queries
	pool    *pgxpool.Pool
}

func NewProductService(pool *pgxpool.Pool) *ProductService {
	return &ProductService{
		queries: repository.New(pool),
		pool:    pool,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	result, err := s.queries.CreateProduct(ctx, repository.CreateProductParams{
		CategoryID:        int32(req.CategoryID),
		Name:              req.Name,
		Slug:              req.Slug,
		Description:       req.Description,
		Price:             req.Price,
		PriceSale:         req.PriceSale,
		ImageUrl:          req.ImageURL,
		UnitOfMeasurement: req.UnitOfMeasurement,
		ThumbUrl:          req.ThumbURL,
	})
	if err != nil {
		return nil, err
	}

	return s.GetProduct(ctx, int(result))
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	product, err := s.queries.GetProduct(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &model.Product{
		ID:                int(product.ID),
		CategoryID:        int(product.CategoryID),
		Name:              product.Name,
		Slug:              product.Slug,
		Description:       product.Description,
		Price:             product.Price,
		PriceSale:         product.PriceSale,
		ImageURL:          product.ImageUrl,
		ThumbURL:          product.ThumbUrl,
		CreatedAt:         product.CreatedAt.Time,
		CategoryName:      product.CategoryName,
		CategorySlug:      product.CategorySlug,
		UnitOfMeasurement: product.UnitOfMeasurement,
	}, nil
}

func (s *ProductService) GetProductBySlug(ctx context.Context, slug string) (*model.Product, error) {
	product, err := s.queries.GetProductBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &model.Product{
		ID:           int(product.ID),
		CategoryID:   int(product.CategoryID),
		Name:         product.Name,
		Slug:         product.Slug,
		Description:  product.Description,
		Price:        product.Price,
		PriceSale:    product.PriceSale,
		ImageURL:     product.ImageUrl,
		ThumbURL:     product.ThumbUrl,
		CreatedAt:    product.CreatedAt.Time,
		CategoryName: product.CategoryName,
		CategorySlug: product.CategorySlug,
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, pagination model.Pagination) ([]model.Product, int64, error) {
	// Get total count first
	totalCount, err := s.queries.GetTotalProducts(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}

	params := repository.ListProductsParams{
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	}

	products, err := s.queries.ListProducts(ctx, params)
	fmt.Println("UnitOfMeasurement", products[0].UnitOfMeasurement)

	if err != nil {
		return nil, 0, err
	}

	// Convert repository products to model products
	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:                int(p.ID),
			CategoryID:        int(p.CategoryID),
			Name:              p.Name,
			Slug:              p.Slug,
			Description:       p.Description,
			Price:             p.Price,
			PriceSale:         p.PriceSale,
			ImageURL:          p.ImageUrl,
			ThumbURL:          p.ThumbUrl,
			CreatedAt:         p.CreatedAt.Time,
			CategoryName:      p.CategoryName,
			CategorySlug:      p.CategorySlug,
			UnitOfMeasurement: p.UnitOfMeasurement,
		}
	}

	return result, totalCount, nil
}

// ListProductsByCategoryID retrieves products by category ID
func (s *ProductService) ListProductsByCategoryID(ctx context.Context, categoryID int64, pagination model.Pagination) ([]model.Product, int64, error) {
	// Get total count first
	totalCount, err := s.queries.GetTotalProductsByCategoryID(ctx, int32(categoryID))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}

	params := repository.ListProductsByCategoryIDParams{
		ID:     int32(categoryID),
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	}

	products, err := s.queries.ListProductsByCategoryID(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	// Convert repository products to model products
	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:                int(p.ID),
			CategoryID:        int(p.CategoryID),
			Name:              p.Name,
			Slug:              p.Slug,
			Description:       p.Description,
			Price:             p.Price,
			PriceSale:         p.PriceSale,
			ImageURL:          p.ImageUrl,
			ThumbURL:          p.ThumbUrl,
			CreatedAt:         p.CreatedAt.Time,
			CategoryName:      p.CategoryName,
			CategorySlug:      p.CategorySlug,
			UnitOfMeasurement: p.UnitOfMeasurement,
		}
	}

	return result, totalCount, nil
}

// ListProductsByCategorySlug retrieves products by category slug
func (s *ProductService) ListProductsByCategorySlug(ctx context.Context, categorySlug string, pagination model.Pagination) ([]model.Product, int64, error) {
	// Get total count first
	totalCount, err := s.queries.GetTotalProductsByCategorySlug(ctx, categorySlug)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}

	params := repository.ListProductsByCategorySlugParams{
		Slug:   categorySlug,
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	}

	products, err := s.queries.ListProductsByCategorySlug(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	// Convert repository products to model products
	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:                int(p.ID),
			CategoryID:        int(p.CategoryID),
			Name:              p.Name,
			Slug:              p.Slug,
			Description:       p.Description,
			Price:             p.Price,
			PriceSale:         p.PriceSale,
			ImageURL:          p.ImageUrl,
			ThumbURL:          p.ThumbUrl,
			CreatedAt:         p.CreatedAt.Time,
			CategoryName:      p.CategoryName,
			CategorySlug:      p.CategorySlug,
			UnitOfMeasurement: p.UnitOfMeasurement,
		}
	}

	return result, totalCount, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, req model.UpdateProductRequest) (*model.Product, error) {
	err := s.queries.UpdateProduct(ctx, repository.UpdateProductParams{
		ID:                int32(id),
		CategoryID:        int32(req.CategoryID),
		Name:              req.Name,
		Slug:              req.Slug,
		Description:       req.Description,
		UnitOfMeasurement: req.UnitOfMeasurement,
		Price:             req.Price,
		PriceSale:         req.PriceSale,
		ImageUrl:          req.ImageURL,
		ThumbUrl:          req.ThumbURL,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return s.GetProduct(ctx, id)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	err := s.queries.DeleteProduct(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}

// GetProductsByCategoryIDs retrieves products grouped by categories based on the website settings
func (s *ProductService) GetProductsByCategoryIDs(ctx context.Context, categoryIDs []int) ([]model.CategoryProductsResponse, error) {
	var result []model.CategoryProductsResponse

	// Get categories with their products
	for _, categoryID := range categoryIDs {
		// Get category details
		category, err := s.queries.GetCategory(ctx, int32(categoryID))
		if err != nil {
			if err == sql.ErrNoRows {
				continue // Skip if category not found
			}
			return nil, fmt.Errorf("error getting category: %w", err)
		}

		// Get products for the category
		products, err := s.queries.ListProductsByCategory(ctx, repository.ListProductsByCategoryParams{
			ID:     int32(categoryID),
			Slug:   category.Slug,
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			return nil, fmt.Errorf("error getting products: %w", err)
		}

		// Convert DB products to model.Product
		modelProducts := make([]model.Product, len(products))
		for i, p := range products {
			modelProducts[i] = model.Product{
				ID:                int(p.ID),
				Name:              p.Name,
				Slug:              p.Slug,
				Price:             float64(p.Price),
				PriceSale:         float64(p.PriceSale),
				ImageURL:          p.ImageUrl,
				ThumbURL:          p.ThumbUrl,
				CreatedAt:         p.CreatedAt.Time,
				UnitOfMeasurement: p.UnitOfMeasurement,
			}
		}

		// Add to result
		result = append(result, model.CategoryProductsResponse{
			Name:     category.Name,
			ImageURL: category.ImageUrl.String,
			Slug:     category.Slug,
			Products: modelProducts,
		})
	}

	return result, nil
}
