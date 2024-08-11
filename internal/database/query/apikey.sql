-- name: GetApiKey :one
SELECT * FROM api_keys
WHERE akey = ?;


-- name: CreateApiKey :exec
INSERT INTO api_keys (
    akey,
    permissions
) VALUES (
    ?, ?
);