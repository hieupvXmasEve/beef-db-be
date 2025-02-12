package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"beef-db-be/internal/model"
	"beef-db-be/internal/repository"
)

// WebsiteSettingService handles business logic for website settings
type WebsiteSettingService struct {
	db      *pgxpool.Pool
	queries *repository.Queries
}

// NewWebsiteSettingService creates a new website setting service
func NewWebsiteSettingService(db *pgxpool.Pool) *WebsiteSettingService {
	return &WebsiteSettingService{
		db:      db,
		queries: repository.New(db),
	}
}

// Create creates a new website setting
func (s *WebsiteSettingService) Create(ctx context.Context, req model.CreateWebsiteSettingRequest) (*model.WebsiteSettingResponse, error) {
	// Check if setting with same name already exists
	_, err := s.queries.GetWebsiteSettingByName(ctx, req.Name)
	if err == nil {
		return nil, errors.New("setting with this name already exists")
	}

	// Create new setting
	id, err := s.queries.CreateWebsiteSetting(ctx, repository.CreateWebsiteSettingParams{
		Name:  req.Name,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}

	// Get created setting
	setting, err := s.queries.GetWebsiteSetting(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.WebsiteSettingResponse{
		ID:    int(setting.ID),
		Name:  setting.Name,
		Value: setting.Value,
	}, nil
}

// Get retrieves a website setting by ID
func (s *WebsiteSettingService) Get(ctx context.Context, id int32) (*model.WebsiteSettingResponse, error) {
	setting, err := s.queries.GetWebsiteSetting(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.WebsiteSettingResponse{
		ID:    int(setting.ID),
		Name:  setting.Name,
		Value: setting.Value,
	}, nil
}

// GetByName retrieves a website setting by name
func (s *WebsiteSettingService) GetByName(ctx context.Context, name string) (*model.WebsiteSettingResponse, error) {
	setting, err := s.queries.GetWebsiteSettingByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.WebsiteSettingResponse{
		ID:    int(setting.ID),
		Name:  setting.Name,
		Value: setting.Value,
	}, nil
}

// List retrieves all website settings
func (s *WebsiteSettingService) List(ctx context.Context) (*model.WebsiteSettingsResponse, error) {
	settings, err := s.queries.ListWebsiteSettings(ctx)
	if err != nil {
		return nil, err
	}

	response := &model.WebsiteSettingsResponse{
		Settings: make([]model.WebsiteSettingResponse, len(settings)),
	}

	for i, setting := range settings {
		response.Settings[i] = model.WebsiteSettingResponse{
			ID:    int(setting.ID),
			Name:  setting.Name,
			Value: setting.Value,
		}
	}

	return response, nil
}

// Update updates a website setting
func (s *WebsiteSettingService) Update(ctx context.Context, name string, req model.UpdateWebsiteSettingRequest) error {
	return s.queries.UpdateWebsiteSetting(ctx, repository.UpdateWebsiteSettingParams{
		Value: req.Value,
		Name:  name,
	})
}

// Delete deletes a website setting
func (s *WebsiteSettingService) Delete(ctx context.Context, id int32) error {
	return s.queries.DeleteWebsiteSetting(ctx, id)
}
