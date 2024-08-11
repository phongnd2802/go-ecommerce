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