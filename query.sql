-- name: CreateUser :execresult
INSERT INTO users (
    name,
    email,
    password_hash
) VALUES (
    ?,
    ?,
    ?
);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT ? OFFSET ?;

-- name: UpdateUser :execresult
UPDATE users
SET 
    name = ?,
    email = ?,
    password_hash = ?
WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = ?;

-- name: GetUserCount :one
SELECT COUNT(*) FROM users;
