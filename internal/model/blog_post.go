package model

import "time"

type BlogPost struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateBlogPostRequest struct {
	Title       string `json:"title" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description" validate:"required"`
	Content     string `json:"content" validate:"required"`
	ImageURL    string `json:"image_url,omitempty"`
}

type UpdateBlogPostRequest struct {
	Title       string `json:"title" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description" validate:"required"`
	Content     string `json:"content" validate:"required"`
	ImageURL    string `json:"image_url,omitempty"`
}

type BlogPostResponse struct {
	BlogPost
}

type BlogPostsResponse struct {
	BlogPosts []BlogPost `json:"blog_posts"`
	Total     int64      `json:"total"`
}
