-- name: GetUsers :many
SELECT * FROM users
    ORDER BY created_at ASC;

-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, email, hashed_password
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING id, email;

-- name: DeleteUsers :exec
DELETE FROM users;