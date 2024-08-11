-- name: CreateShop :exec
INSERT INTO shops (
    id,
    shop_name,
    email,
    password
) VALUES (
    ?, ?, ?, ?
);

-- name: GetShopByID :one
SELECT * FROM shops
WHERE id = ?
LIMIT 1;


-- name: GetShopByEmail :one
SELECT * FROM shops
WHERE email = ?
LIMIT 1;

