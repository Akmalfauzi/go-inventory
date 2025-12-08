-- name: CreateProduct :one
INSERT INTO products (name, price, stock)
VALUES ($1, $2, $3)
RETURNING id, name, price, stock, created_at;

-- name: ListProducts :many
SELECT id, name, price, stock, created_at
FROM products
ORDER BY created_at DESC;

-- name: GetProduct :one
SELECT id, name, price, stock, created_at
FROM products
WHERE id = $1 LIMIT 1;