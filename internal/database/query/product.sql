-- name: CreateProduct :exec
INSERT INTO products(
    id, product_name, product_thumb, product_description,
    product_price, product_quantity, product_type, product_shop,
    product_slug, product_attributes
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);


-- name: GetProductByID :one
SELECT * FROM products
WHERE id = ?;