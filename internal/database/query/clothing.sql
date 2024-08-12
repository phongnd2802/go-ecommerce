-- name: CreateClothing :exec
INSERT INTO clothes (
    id, brand, size, material, product_shop
) VALUES (?, ?, ?, ?, ?);


-- name: GetClothingByID :one
SELECT * FROM clothes
WHERE id = ?;