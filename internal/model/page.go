package model

import "time"

type Page struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePageRequest struct {
	Title   string `json:"title" validate:"required"`
	Slug    string `json:"slug" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdatePageRequest struct {
	Title   string `json:"title" validate:"required"`
	Slug    string `json:"slug" validate:"required"`
	Content string `json:"content" validate:"required"`
}
