-- name: CreateClothing :exec
INSERT INTO clothes (
    id, brand, size, material, product_shop
) VALUES (?, ?, ?, ?, ?);


-- name: GetClothingByID :one
SELECT * FROM clothes
WHERE id = ?;

-- name: UpdateClothingByID :exec
UPDATE clothes 
SET brand = ?,
    size = ?,
    material = ?
WHERE id = ?;
