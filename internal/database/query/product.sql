-- name: CreateProduct :exec
INSERT INTO products(
    id, product_name, product_thumb, product_description,
    product_price, product_quantity, product_type, product_shop,
    product_slug, product_attributes
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);


-- name: GetProductByID :one
SELECT * FROM products
WHERE id = ?;

-- name: QueryProrductForShop :many
SELECT * FROM products
WHERE isDraft = ? 
AND isPublished = ?
AND product_shop = ?
ORDER BY updated_at DESC
LIMIT ?
OFFSET ?;

-- name: GetProductByShopAndID :one
SELECT * FROM products
WHERE id = ? AND product_shop = ?;

-- name: UpdateStatusProductByShop :exec
UPDATE products 
SET isPublished = ?, isDraft = ? 
WHERE id = ?;
