-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id;

-- name: GetUser :one
SELECT id, email, password, role, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, password, role, created_at, updated_at
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT id, email, role, created_at, updated_at
FROM users;

-- name: UpdateUser :exec
UPDATE users
SET email = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserCount :one
SELECT COUNT(*)
FROM users;

-- name: CreateCategory :one
INSERT INTO categories (name, slug, description, image_url)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetCategory :one
SELECT *
FROM categories
WHERE id = $1;

-- name: GetCategoryBySlug :one
SELECT *
FROM categories
WHERE slug = $1;

-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY created_at DESC;

-- name: UpdateCategory :exec
UPDATE categories
SET name = $1, slug = $2, description = $3, image_url = $4
WHERE id = $5;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (
    category_id,
    name,
    slug,
    description,
    price,
    price_sale,
    image_url,
    thumb_url
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: GetProduct :one
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.id = $1;

-- name: GetProductBySlug :one
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.slug = $1;

-- name: ListProducts :many
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListProductsByCategory :many
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
LIMIT $3 OFFSET $4;

-- name: UpdateProduct :exec
UPDATE products
SET
    category_id = $1,
    name = $2,
    slug = $3,
    description = $4,
    price = $5,
    price_sale = $6,
    image_url = $7,
    thumb_url = $8
WHERE id = $9;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: GetTotalProducts :one
SELECT COUNT(*) AS total_count FROM products;

-- name: GetTotalProductsByCategory :one
SELECT COUNT(*) AS total_count 
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE (SQLC_OMIT_IF_NULL(@id::int4) IS NULL OR c.id = @id)
  AND (SQLC_OMIT_IF_NULL(@slug::text) IS NULL OR c.slug = @slug);