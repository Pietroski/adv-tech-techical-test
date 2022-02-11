-- name: GetUserByID :one
SELECT *
FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY table_id;

-- name: ListPaginatedUsers :many
SELECT *
FROM users
ORDER BY table_id LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (user_id, email, first_name, last_name)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET (email, first_name, last_name) = ($2, $3, $4)
WHERE user_id = $1 RETURNING *;

-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE user_id = $1;

-- name: DeleteUserByEmail :exec
DELETE
FROM users
WHERE email = $1;
