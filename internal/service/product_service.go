package service

import (
	"context"
	"database/sql"
	"errors"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type ProductService struct {
	queries *repository.Queries
	db      *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		queries: repository.New(db),
		db:      db,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	result, err := s.queries.CreateProduct(ctx, repository.CreateProductParams{
		CategoryID:  int32(req.CategoryID),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Price:       req.Price,
		ImageUrl:    sql.NullString{String: req.ImageURL, Valid: req.ImageURL != ""},
		ThumbUrl:    sql.NullString{String: req.ThumbURL, Valid: req.ThumbURL != ""},
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.GetProduct(ctx, int(id))
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	product, err := s.queries.GetProduct(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &model.Product{
		ID:           int(product.ID),
		CategoryID:   int(product.CategoryID),
		Name:         product.Name,
		Slug:         product.Slug,
		Description:  product.Description.String,
		Price:        product.Price,
		ImageURL:     product.ImageUrl.String,
		ThumbURL:     product.ThumbUrl.String,
		CreatedAt:    product.CreatedAt.Time,
		CategoryName: product.CategoryName,
		CategorySlug: product.CategorySlug,
	}, nil
}

func (s *ProductService) GetProductBySlug(ctx context.Context, slug string) (*model.Product, error) {
	product, err := s.queries.GetProductBySlug(ctx, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &model.Product{
		ID:           int(product.ID),
		CategoryID:   int(product.CategoryID),
		Name:         product.Name,
		Slug:         product.Slug,
		Description:  product.Description.String,
		Price:        product.Price,
		ImageURL:     product.ImageUrl.String,
		ThumbURL:     product.ThumbUrl.String,
		CreatedAt:    product.CreatedAt.Time,
		CategoryName: product.CategoryName,
		CategorySlug: product.CategorySlug,
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context) ([]model.Product, error) {
	products, err := s.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:           int(p.ID),
			CategoryID:   int(p.CategoryID),
			Name:         p.Name,
			Slug:         p.Slug,
			Description:  p.Description.String,
			Price:        p.Price,
			ImageURL:     p.ImageUrl.String,
			ThumbURL:     p.ThumbUrl.String,
			CreatedAt:    p.CreatedAt.Time,
			CategoryName: p.CategoryName,
			CategorySlug: p.CategorySlug,
		}
	}

	return result, nil
}

func (s *ProductService) ListProductsByCategory(ctx context.Context, categoryID int, categorySlug string) ([]model.Product, error) {
	products, err := s.queries.ListProductsByCategory(ctx, repository.ListProductsByCategoryParams{
		ID:   int32(categoryID),
		Slug: categorySlug,
	})
	if err != nil {
		return nil, err
	}

	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:           int(p.ID),
			CategoryID:   int(p.CategoryID),
			Name:         p.Name,
			Slug:         p.Slug,
			Description:  p.Description.String,
			Price:        p.Price,
			ImageURL:     p.ImageUrl.String,
			ThumbURL:     p.ThumbUrl.String,
			CreatedAt:    p.CreatedAt.Time,
			CategoryName: p.CategoryName,
			CategorySlug: p.CategorySlug,
		}
	}

	return result, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, req model.UpdateProductRequest) (*model.Product, error) {
	err := s.queries.UpdateProduct(ctx, repository.UpdateProductParams{
		ID:          int32(id),
		CategoryID:  int32(req.CategoryID),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Price:       req.Price,
		ImageUrl:    sql.NullString{String: req.ImageURL, Valid: req.ImageURL != ""},
		ThumbUrl:    sql.NullString{String: req.ThumbURL, Valid: req.ThumbURL != ""},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return s.GetProduct(ctx, id)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	err := s.queries.DeleteProduct(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}
