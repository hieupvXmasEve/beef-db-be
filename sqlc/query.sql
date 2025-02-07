-- name: CreateUser :execresult
INSERT INTO users (
    email, password, created_at, updated_at
) VALUES (
    ?, ?, NOW(), NOW()
);

-- name: GetUser :one
SELECT id, email, password, role, created_at, updated_at FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, password, role, created_at, updated_at FROM users
WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT id, email, role, created_at, updated_at FROM users;

-- name: UpdateUser :exec
UPDATE users
SET email = ?, updated_at = NOW()
WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = ?;

-- name: GetUserCount :one
SELECT COUNT(*) FROM users;
