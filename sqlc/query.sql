-- name: CreateUser :execresult
INSERT INTO
    users (email, password)
VALUES
    (?, ?);

-- name: GetUser :one
SELECT
    id,
    email,
    password,
    role,
    created_at,
    updated_at
FROM
    users
WHERE
    id = ?;

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    password,
    role,
    created_at,
    updated_at
FROM
    users
WHERE
    email = ?;

-- name: ListUsers :many
SELECT
    id,
    email,
    role,
    created_at,
    updated_at
FROM
    users;

-- name: UpdateUser :exec
UPDATE users
SET
    email = ?
WHERE
    id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE
    id = ?;

-- name: GetUserCount :one
SELECT
    COUNT(*)
FROM
    users;

-- name: CreateCategory :execresult
INSERT INTO
    categories (name, slug, description, image_url)
VALUES
    (?, ?, ?, ?);

-- name: GetCategory :one
SELECT
    *
FROM
    categories
WHERE
    id = ?;

-- name: GetCategoryBySlug :one
SELECT
    *
FROM
    categories
WHERE
    slug = ?;

-- name: ListCategories :many
SELECT
    *
FROM
    categories
ORDER BY
    created_at DESC;

-- name: UpdateCategory :exec
UPDATE categories
SET
    name = ?,
    slug = ?,
    description = ?,
    image_url = ?
WHERE
    id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE
    id = ?;

-- name: CreateProduct :execresult
INSERT INTO
    products (
        category_id,
        name,
        slug,
        description,
        price,
        image_url,
        thumb_url
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: GetProduct :one
SELECT
    p.*,
    c.name as category_name,
    c.slug as category_slug
FROM
    products p
    JOIN categories c ON p.category_id = c.id
WHERE
    p.id = ?;

-- name: GetProductBySlug :one
SELECT
    p.*,
    c.name as category_name,
    c.slug as category_slug
FROM
    products p
    JOIN categories c ON p.category_id = c.id
WHERE
    p.slug = ?;

-- name: ListProducts :many
SELECT
    p.*,
    c.name as category_name,
    c.slug as category_slug
FROM
    products p
    JOIN categories c ON p.category_id = c.id
ORDER BY
    p.created_at DESC;

-- name: ListProductsByCategory :many
SELECT
    p.*,
    c.name as category_name,
    c.slug as category_slug
FROM
    products p
    JOIN categories c ON p.category_id = c.id
WHERE
    c.id = ?
    OR c.slug = ?
ORDER BY
    p.created_at DESC;

-- name: UpdateProduct :exec
UPDATE products
SET
    category_id = ?,
    name = ?,
    slug = ?,
    description = ?,
    price = ?,
    image_url = ?,
    thumb_url = ?
WHERE
    id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE
    id = ?;