package model

// WebsiteSetting represents a website setting in the system
type WebsiteSetting struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CreateWebsiteSettingRequest represents the request to create a website setting
type CreateWebsiteSettingRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateWebsiteSettingRequest represents the request to update a website setting
type UpdateWebsiteSettingRequest struct {
	Value string `json:"value" validate:"required"`
}

// WebsiteSettingResponse represents a website setting response
type WebsiteSettingResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// WebsiteSettingsResponse represents a list of website settings
type WebsiteSettingsResponse struct {
	Settings []WebsiteSettingResponse `json:"settings"`
}
