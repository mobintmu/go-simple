-- name: CreateProduct :one
INSERT INTO products (product_name, product_description, price, is_active)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE products
SET product_name = $2, product_description = $3, price = $4, is_active = $5
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;