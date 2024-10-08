-- name: CreateToken :exec
INSERT INTO tokens (
    id,
    public_key,
    refresh_token,
    shop_id
) VALUES (
    ?, ?, ?, ?
);


-- name: GetTokenByID :one
SELECT * FROM tokens
WHERE id = ?;

-- name: GetTokenByShopID :one
SELECT * FROM tokens
WHERE shop_id = ?;

-- name: DeleteTokenByID :exec
DELETE FROM tokens WHERE id = ?;

-- name: UpdateToken :exec
UPDATE tokens SET refresh_token = ?, refresh_token_used = ?, public_key = ? WHERE id = ?;