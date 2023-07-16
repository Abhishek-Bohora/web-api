-- name: CreateProduct :one
INSERT INTO products (id, created_at, updated_at, name, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;