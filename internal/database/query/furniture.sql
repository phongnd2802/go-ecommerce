-- name: CreateFurniture :exec
INSERT INTO furnitures (
    id, brand, size, material, product_shop
) VALUES (?, ?, ?, ?, ?);

-- name: GetFurnitureByID :one
SELECT * FROM furnitures WHERE id = ?;