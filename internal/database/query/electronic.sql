-- name: CreateElectronic :exec
INSERT INTO electronics (
    id, manufacturer, model, color, product_shop
) VALUES (?, ?, ?, ?, ?);


-- name: GetElectronicByID :one
SELECT * FROM electronics WHERE id = ?;