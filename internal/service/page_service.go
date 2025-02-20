package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type PageService struct {
	queries *repository.Queries
	pool    *pgxpool.Pool
}

func NewPageService(dbPool *pgxpool.Pool) *PageService {
	return &PageService{
		queries: repository.New(dbPool),
		pool:    dbPool,
	}
}

func (s *PageService) CreatePage(ctx context.Context, req model.CreatePageRequest) (*model.Page, error) {
	page, err := s.queries.CreatePage(ctx, repository.CreatePageParams{
		Title:   req.Title,
		Slug:    req.Slug,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &model.Page{
		ID:        int64(page.ID),
		Title:     page.Title,
		Slug:      page.Slug,
		Content:   page.Content,
		CreatedAt: page.CreatedAt.Time,
	}, nil
}

func (s *PageService) GetPage(ctx context.Context, id int32) (*model.Page, error) {
	page, err := s.queries.GetPage(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &model.Page{
		ID:        int64(page.ID),
		Title:     page.Title,
		Slug:      page.Slug,
		Content:   page.Content,
		CreatedAt: page.CreatedAt.Time,
	}, nil
}

func (s *PageService) GetPageBySlug(ctx context.Context, slug string) (*model.Page, error) {
	page, err := s.queries.GetPageBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &model.Page{
		ID:        int64(page.ID),
		Title:     page.Title,
		Slug:      page.Slug,
		Content:   page.Content,
		CreatedAt: page.CreatedAt.Time,
	}, nil
}

func (s *PageService) ListPages(ctx context.Context, pagination model.Pagination) ([]model.Page, int64, error) {
	totalCount, err := s.queries.GetTotalPages(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}
	pages, err := s.queries.ListPages(ctx, repository.ListPagesParams{
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	})
	if err != nil {
		return nil, 0, err
	}

	result := make([]model.Page, len(pages))
	for i, page := range pages {
		result[i] = model.Page{
			ID:        int64(page.ID),
			Title:     page.Title,
			Slug:      page.Slug,
			Content:   page.Content,
			CreatedAt: page.CreatedAt.Time,
		}
	}
	return result, totalCount, nil
}

func (s *PageService) UpdatePage(ctx context.Context, id int32, req model.UpdatePageRequest) error {
	err := s.queries.UpdatePage(ctx, repository.UpdatePageParams{
		ID:      id,
		Title:   req.Title,
		Slug:    req.Slug,
		Content: req.Content,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (s *PageService) DeletePage(ctx context.Context, id int32) error {
	err := s.queries.DeletePage(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return nil
}
