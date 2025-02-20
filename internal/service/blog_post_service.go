package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

type BlogPostService struct {
	queries *repository.Queries
	pool    *pgxpool.Pool
}

func NewBlogPostService(pool *pgxpool.Pool) *BlogPostService {
	return &BlogPostService{
		queries: repository.New(pool),
		pool:    pool,
	}
}

func (s *BlogPostService) Create(ctx context.Context, req model.CreateBlogPostRequest) (*model.BlogPost, error) {
	result, err := s.queries.CreateBlogPost(ctx, repository.CreateBlogPostParams{
		Title:    req.Title,
		Slug:     req.Slug,
		Content:  req.Content,
		ImageUrl: req.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	return &model.BlogPost{
		ID:        int64(result.ID),
		Title:     result.Title,
		Slug:      result.Slug,
		Content:   result.Content,
		ImageURL:  result.ImageUrl,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (s *BlogPostService) GetByID(ctx context.Context, id int64) (*model.BlogPost, error) {
	post, err := s.queries.GetBlogPost(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &model.BlogPost{
		ID:          int64(post.ID),
		Title:       post.Title,
		Slug:        post.Slug,
		Description: post.Description,
		Content:     post.Content,
		ImageURL:    post.ImageUrl,
		CreatedAt:   post.CreatedAt.Time,
	}, nil
}

func (s *BlogPostService) GetBySlug(ctx context.Context, slug string) (*model.BlogPost, error) {
	post, err := s.queries.GetBlogPostBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &model.BlogPost{
		ID:        int64(post.ID),
		Title:     post.Title,
		Slug:      post.Slug,
		Content:   post.Content,
		ImageURL:  post.ImageUrl,
		CreatedAt: post.CreatedAt.Time,
	}, nil
}

func (s *BlogPostService) List(ctx context.Context, limit, offset int) ([]model.BlogPost, int64, error) {
	// Get total count first
	totalCount, err := s.queries.GetTotalBlogPosts(ctx)
	if err != nil {
		return nil, 0, err
	}

	posts, err := s.queries.ListBlogPosts(ctx, repository.ListBlogPostsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}

	result := make([]model.BlogPost, len(posts))
	for i, post := range posts {
		result[i] = model.BlogPost{
			ID:        int64(post.ID),
			Title:     post.Title,
			Slug:      post.Slug,
			Content:   post.Content,
			ImageURL:  post.ImageUrl,
			CreatedAt: post.CreatedAt.Time,
		}
	}

	return result, totalCount, nil
}

func (s *BlogPostService) Update(ctx context.Context, id int64, req model.UpdateBlogPostRequest) error {
	err := s.queries.UpdateBlogPost(ctx, repository.UpdateBlogPostParams{
		ID:       int32(id),
		Title:    req.Title,
		Slug:     req.Slug,
		Content:  req.Content,
		ImageUrl: req.ImageURL,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (s *BlogPostService) Delete(ctx context.Context, id int64) error {
	err := s.queries.DeleteBlogPost(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
