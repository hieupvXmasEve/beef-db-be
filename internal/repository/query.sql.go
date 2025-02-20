// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBlogPost = `-- name: CreateBlogPost :one
INSERT INTO blog_posts (
    title,
    description,
    content,
    slug,
    image_url,
    created_at
)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
RETURNING id, title, slug, description, content, image_url, created_at, updated_at
`

type CreateBlogPostParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Slug        string `json:"slug"`
	ImageUrl    string `json:"image_url"`
}

// Blog Post Queries
func (q *Queries) CreateBlogPost(ctx context.Context, arg CreateBlogPostParams) (BlogPost, error) {
	row := q.db.QueryRow(ctx, createBlogPost,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.Slug,
		arg.ImageUrl,
	)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Description,
		&i.Content,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name, slug, description, image_url)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateCategoryParams struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Description pgtype.Text `json:"description"`
	ImageUrl    pgtype.Text `json:"image_url"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, createCategory,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.ImageUrl,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createPage = `-- name: CreatePage :one
INSERT INTO pages (
    slug,
    title,
    content,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, slug, title, description, content, created_at, updated_at
`

type CreatePageParams struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Pages Queries
func (q *Queries) CreatePage(ctx context.Context, arg CreatePageParams) (Page, error) {
	row := q.db.QueryRow(ctx, createPage, arg.Slug, arg.Title, arg.Content)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    category_id,
    name,
    slug,
    description,
    price,
    price_sale,
    unit_of_measurement,
    image_url,
    thumb_url
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id
`

type CreateProductParams struct {
	CategoryID        int32   `json:"category_id"`
	Name              string  `json:"name"`
	Slug              string  `json:"slug"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	PriceSale         float64 `json:"price_sale"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
	ImageUrl          string  `json:"image_url"`
	ThumbUrl          string  `json:"thumb_url"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (int32, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.CategoryID,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.Price,
		arg.PriceSale,
		arg.UnitOfMeasurement,
		arg.ImageUrl,
		arg.ThumbUrl,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Password)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createWebsiteSetting = `-- name: CreateWebsiteSetting :one
INSERT INTO website_settings (name, value)
VALUES ($1, $2)
RETURNING id
`

type CreateWebsiteSettingParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (q *Queries) CreateWebsiteSetting(ctx context.Context, arg CreateWebsiteSettingParams) (int32, error) {
	row := q.db.QueryRow(ctx, createWebsiteSetting, arg.Name, arg.Value)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteBlogPost = `-- name: DeleteBlogPost :exec
DELETE FROM blog_posts
WHERE id = $1
`

func (q *Queries) DeleteBlogPost(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteBlogPost, id)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const deletePage = `-- name: DeletePage :exec
DELETE FROM pages
WHERE id = $1
`

func (q *Queries) DeletePage(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePage, id)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const deleteWebsiteSetting = `-- name: DeleteWebsiteSetting :exec
DELETE FROM website_settings
WHERE id = $1
`

func (q *Queries) DeleteWebsiteSetting(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteWebsiteSetting, id)
	return err
}

const getBlogPost = `-- name: GetBlogPost :one
SELECT id, title, slug, description, content, image_url, created_at, updated_at
FROM blog_posts
WHERE id = $1
`

func (q *Queries) GetBlogPost(ctx context.Context, id int32) (BlogPost, error) {
	row := q.db.QueryRow(ctx, getBlogPost, id)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Description,
		&i.Content,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBlogPostBySlug = `-- name: GetBlogPostBySlug :one
SELECT id, title, slug, description, content, image_url, created_at, updated_at
FROM blog_posts
WHERE slug = $1
`

func (q *Queries) GetBlogPostBySlug(ctx context.Context, slug string) (BlogPost, error) {
	row := q.db.QueryRow(ctx, getBlogPostBySlug, slug)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Description,
		&i.Content,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, slug, description, image_url, created_at
FROM categories
WHERE id = $1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRow(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getCategoryBySlug = `-- name: GetCategoryBySlug :one
SELECT id, name, slug, description, image_url, created_at
FROM categories
WHERE slug = $1
`

func (q *Queries) GetCategoryBySlug(ctx context.Context, slug string) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryBySlug, slug)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getPage = `-- name: GetPage :one
SELECT id, slug, title, description, content, created_at, updated_at
FROM pages
WHERE id = $1
`

func (q *Queries) GetPage(ctx context.Context, id int32) (Page, error) {
	row := q.db.QueryRow(ctx, getPage, id)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPageBySlug = `-- name: GetPageBySlug :one
SELECT id, slug, title, description, content, created_at, updated_at
FROM pages
WHERE slug = $1
`

func (q *Queries) GetPageBySlug(ctx context.Context, slug string) (Page, error) {
	row := q.db.QueryRow(ctx, getPageBySlug, slug)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.id = $1
`

type GetProductRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
}

func (q *Queries) GetProduct(ctx context.Context, id int32) (GetProductRow, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i GetProductRow
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.Price,
		&i.PriceSale,
		&i.UnitOfMeasurement,
		&i.ImageUrl,
		&i.ThumbUrl,
		&i.CreatedAt,
		&i.CategoryName,
		&i.CategorySlug,
	)
	return i, err
}

const getProductBySlug = `-- name: GetProductBySlug :one
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.slug = $1
`

type GetProductBySlugRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
}

func (q *Queries) GetProductBySlug(ctx context.Context, slug string) (GetProductBySlugRow, error) {
	row := q.db.QueryRow(ctx, getProductBySlug, slug)
	var i GetProductBySlugRow
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.Price,
		&i.PriceSale,
		&i.UnitOfMeasurement,
		&i.ImageUrl,
		&i.ThumbUrl,
		&i.CreatedAt,
		&i.CategoryName,
		&i.CategorySlug,
	)
	return i, err
}

const getTotalBlogPosts = `-- name: GetTotalBlogPosts :one
SELECT COUNT(*) as total_count
FROM blog_posts
`

func (q *Queries) GetTotalBlogPosts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalBlogPosts)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getTotalProducts = `-- name: GetTotalProducts :one
SELECT COUNT(*) AS total_count FROM products
`

func (q *Queries) GetTotalProducts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalProducts)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getTotalProductsByCategoryID = `-- name: GetTotalProductsByCategoryID :one
SELECT COUNT(*) AS total_count
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE (c.id = $1)
`

func (q *Queries) GetTotalProductsByCategoryID(ctx context.Context, id int32) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalProductsByCategoryID, id)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getTotalProductsByCategorySlug = `-- name: GetTotalProductsByCategorySlug :one
SELECT COUNT(*) AS total_count
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE (c.slug = $1)
`

func (q *Queries) GetTotalProductsByCategorySlug(ctx context.Context, slug string) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalProductsByCategorySlug, slug)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password, role, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, role, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserCount = `-- name: GetUserCount :one
SELECT COUNT(*)
FROM users
`

func (q *Queries) GetUserCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getUserCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getWebsiteSetting = `-- name: GetWebsiteSetting :one
SELECT id, name, value
FROM website_settings
WHERE id = $1
`

func (q *Queries) GetWebsiteSetting(ctx context.Context, id int32) (WebsiteSetting, error) {
	row := q.db.QueryRow(ctx, getWebsiteSetting, id)
	var i WebsiteSetting
	err := row.Scan(&i.ID, &i.Name, &i.Value)
	return i, err
}

const getWebsiteSettingByName = `-- name: GetWebsiteSettingByName :one
SELECT id, name, value
FROM website_settings
WHERE name = $1
`

func (q *Queries) GetWebsiteSettingByName(ctx context.Context, name string) (WebsiteSetting, error) {
	row := q.db.QueryRow(ctx, getWebsiteSettingByName, name)
	var i WebsiteSetting
	err := row.Scan(&i.ID, &i.Name, &i.Value)
	return i, err
}

const listBlogPosts = `-- name: ListBlogPosts :many
SELECT id, title, slug, description, content, image_url, created_at, updated_at
FROM blog_posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListBlogPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlogPosts(ctx context.Context, arg ListBlogPostsParams) ([]BlogPost, error) {
	rows, err := q.db.Query(ctx, listBlogPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BlogPost{}
	for rows.Next() {
		var i BlogPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Slug,
			&i.Description,
			&i.Content,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, slug, description, image_url, created_at
FROM categories
ORDER BY created_at DESC
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.ImageUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPages = `-- name: ListPages :many
SELECT id, slug, title, description, content, created_at, updated_at
FROM pages
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListPagesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPages(ctx context.Context, arg ListPagesParams) ([]Page, error) {
	rows, err := q.db.Query(ctx, listPages, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Page{}
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.ID,
			&i.Slug,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListProductsRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]ListProductsRow, error) {
	rows, err := q.db.Query(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsRow{}
	for rows.Next() {
		var i ListProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.PriceSale,
			&i.UnitOfMeasurement,
			&i.ImageUrl,
			&i.ThumbUrl,
			&i.CreatedAt,
			&i.CategoryName,
			&i.CategorySlug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductsByCategory = `-- name: ListProductsByCategory :many
WITH total AS (
    SELECT COUNT(*) as count
    FROM products p
    JOIN categories c ON p.category_id = c.id
    WHERE c.id = $1 OR c.slug = $2
)
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug,
    total.count as total_count
FROM products p
JOIN categories c ON p.category_id = c.id
CROSS JOIN total
WHERE c.id = $1 OR c.slug = $2
ORDER BY p.created_at DESC
LIMIT $3 OFFSET $4
`

type ListProductsByCategoryParams struct {
	ID     int32  `json:"id"`
	Slug   string `json:"slug"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type ListProductsByCategoryRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
	TotalCount        int64            `json:"total_count"`
}

func (q *Queries) ListProductsByCategory(ctx context.Context, arg ListProductsByCategoryParams) ([]ListProductsByCategoryRow, error) {
	rows, err := q.db.Query(ctx, listProductsByCategory,
		arg.ID,
		arg.Slug,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsByCategoryRow{}
	for rows.Next() {
		var i ListProductsByCategoryRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.PriceSale,
			&i.UnitOfMeasurement,
			&i.ImageUrl,
			&i.ThumbUrl,
			&i.CreatedAt,
			&i.CategoryName,
			&i.CategorySlug,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductsByCategoryID = `-- name: ListProductsByCategoryID :many
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE c.id = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3
`

type ListProductsByCategoryIDParams struct {
	ID     int32 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListProductsByCategoryIDRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
}

func (q *Queries) ListProductsByCategoryID(ctx context.Context, arg ListProductsByCategoryIDParams) ([]ListProductsByCategoryIDRow, error) {
	rows, err := q.db.Query(ctx, listProductsByCategoryID, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsByCategoryIDRow{}
	for rows.Next() {
		var i ListProductsByCategoryIDRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.PriceSale,
			&i.UnitOfMeasurement,
			&i.ImageUrl,
			&i.ThumbUrl,
			&i.CreatedAt,
			&i.CategoryName,
			&i.CategorySlug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductsByCategorySlug = `-- name: ListProductsByCategorySlug :many
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE c.slug = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3
`

type ListProductsByCategorySlugParams struct {
	Slug   string `json:"slug"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type ListProductsByCategorySlugRow struct {
	ID                int32            `json:"id"`
	CategoryID        int32            `json:"category_id"`
	Name              string           `json:"name"`
	Slug              string           `json:"slug"`
	Description       string           `json:"description"`
	Price             float64          `json:"price"`
	PriceSale         float64          `json:"price_sale"`
	UnitOfMeasurement string           `json:"unit_of_measurement"`
	ImageUrl          string           `json:"image_url"`
	ThumbUrl          string           `json:"thumb_url"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	CategoryName      string           `json:"category_name"`
	CategorySlug      string           `json:"category_slug"`
}

func (q *Queries) ListProductsByCategorySlug(ctx context.Context, arg ListProductsByCategorySlugParams) ([]ListProductsByCategorySlugRow, error) {
	rows, err := q.db.Query(ctx, listProductsByCategorySlug, arg.Slug, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsByCategorySlugRow{}
	for rows.Next() {
		var i ListProductsByCategorySlugRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.PriceSale,
			&i.UnitOfMeasurement,
			&i.ImageUrl,
			&i.ThumbUrl,
			&i.CreatedAt,
			&i.CategoryName,
			&i.CategorySlug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, role, created_at, updated_at
FROM users
`

type ListUsersRow struct {
	ID        int64            `json:"id"`
	Email     string           `json:"email"`
	Role      string           `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersRow{}
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listWebsiteSettings = `-- name: ListWebsiteSettings :many
SELECT id, name, value
FROM website_settings
`

func (q *Queries) ListWebsiteSettings(ctx context.Context) ([]WebsiteSetting, error) {
	rows, err := q.db.Query(ctx, listWebsiteSettings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WebsiteSetting{}
	for rows.Next() {
		var i WebsiteSetting
		if err := rows.Scan(&i.ID, &i.Name, &i.Value); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchBlogPosts = `-- name: SearchBlogPosts :many
SELECT id, title, slug, description, content, image_url, created_at, updated_at
FROM blog_posts
WHERE 
    title ILIKE '%' || $1 || '%' OR
    content ILIKE '%' || $1 || '%'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type SearchBlogPostsParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) SearchBlogPosts(ctx context.Context, arg SearchBlogPostsParams) ([]BlogPost, error) {
	rows, err := q.db.Query(ctx, searchBlogPosts, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BlogPost{}
	for rows.Next() {
		var i BlogPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Slug,
			&i.Description,
			&i.Content,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlogPost = `-- name: UpdateBlogPost :exec
UPDATE blog_posts
SET 
    title = $1,
    description = $2,
    content = $3,
    image_url = $4,
    slug = $5
WHERE id = $6
`

type UpdateBlogPostParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	ImageUrl    string `json:"image_url"`
	Slug        string `json:"slug"`
	ID          int32  `json:"id"`
}

func (q *Queries) UpdateBlogPost(ctx context.Context, arg UpdateBlogPostParams) error {
	_, err := q.db.Exec(ctx, updateBlogPost,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.ImageUrl,
		arg.Slug,
		arg.ID,
	)
	return err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET name = $1, slug = $2, description = $3, image_url = $4
WHERE id = $5
`

type UpdateCategoryParams struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Description pgtype.Text `json:"description"`
	ImageUrl    pgtype.Text `json:"image_url"`
	ID          int32       `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.Exec(ctx, updateCategory,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.ImageUrl,
		arg.ID,
	)
	return err
}

const updatePage = `-- name: UpdatePage :exec
UPDATE pages
SET 
    slug = $1,
    title = $2,
    content = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $4
`

type UpdatePageParams struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      int32  `json:"id"`
}

func (q *Queries) UpdatePage(ctx context.Context, arg UpdatePageParams) error {
	_, err := q.db.Exec(ctx, updatePage,
		arg.Slug,
		arg.Title,
		arg.Content,
		arg.ID,
	)
	return err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET
    category_id = $1,
    name = $2,
    slug = $3,
    description = $4,
    price = $5,
    price_sale = $6,
    unit_of_measurement = $7,
    image_url = $8,
    thumb_url = $9
WHERE id = $10
`

type UpdateProductParams struct {
	CategoryID        int32   `json:"category_id"`
	Name              string  `json:"name"`
	Slug              string  `json:"slug"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	PriceSale         float64 `json:"price_sale"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
	ImageUrl          string  `json:"image_url"`
	ThumbUrl          string  `json:"thumb_url"`
	ID                int32   `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.CategoryID,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.Price,
		arg.PriceSale,
		arg.UnitOfMeasurement,
		arg.ImageUrl,
		arg.ThumbUrl,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET email = $1
WHERE id = $2
`

type UpdateUserParams struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Email, arg.ID)
	return err
}

const updateWebsiteSetting = `-- name: UpdateWebsiteSetting :exec
UPDATE website_settings
SET value = $1
WHERE name = $2
`

type UpdateWebsiteSettingParams struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (q *Queries) UpdateWebsiteSetting(ctx context.Context, arg UpdateWebsiteSettingParams) error {
	_, err := q.db.Exec(ctx, updateWebsiteSetting, arg.Value, arg.Name)
	return err
}
