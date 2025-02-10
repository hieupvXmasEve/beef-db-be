package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type CategoryService struct {
	queries *repository.Queries
	pool    *pgxpool.Pool
}

func NewCategoryService(pool *pgxpool.Pool) *CategoryService {
	return &CategoryService{
		queries: repository.New(pool),
		pool:    pool,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req model.CreateCategoryRequest) (*model.Category, error) {
	result, err := s.queries.CreateCategory(ctx, repository.CreateCategoryParams{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
		ImageUrl:    pgtype.Text{String: req.ImageURL, Valid: req.ImageURL != ""},
	})
	if err != nil {
		return nil, err
	}

	return s.GetCategory(ctx, int(result))
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (*model.Category, error) {
	category, err := s.queries.GetCategory(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &model.Category{
		ID:          int(category.ID),
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description.String,
		ImageURL:    category.ImageUrl.String,
		CreatedAt:   category.CreatedAt.Time,
	}, nil
}

func (s *CategoryService) GetCategoryBySlug(ctx context.Context, slug string) (*model.Category, error) {
	category, err := s.queries.GetCategoryBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &model.Category{
		ID:          int(category.ID),
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description.String,
		ImageURL:    category.ImageUrl.String,
		CreatedAt:   category.CreatedAt.Time,
	}, nil
}

func (s *CategoryService) ListCategories(ctx context.Context) ([]model.Category, error) {
	categories, err := s.queries.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.Category, len(categories))
	for i, category := range categories {
		result[i] = model.Category{
			ID:          int(category.ID),
			Name:        category.Name,
			Slug:        category.Slug,
			Description: category.Description.String,
			ImageURL:    category.ImageUrl.String,
			CreatedAt:   category.CreatedAt.Time,
		}
	}

	return result, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id int, req model.UpdateCategoryRequest) (*model.Category, error) {
	err := s.queries.UpdateCategory(ctx, repository.UpdateCategoryParams{
		ID:          int32(id),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
		ImageUrl:    pgtype.Text{String: req.ImageURL, Valid: req.ImageURL != ""},
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return s.GetCategory(ctx, id)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {
	err := s.queries.DeleteCategory(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("category not found")
		}
		return err
	}
	return nil
}
