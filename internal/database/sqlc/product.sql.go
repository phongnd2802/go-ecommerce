// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products(
    id, product_name, product_thumb, product_description,
    product_price, product_quantity, product_type, product_shop,
    product_slug, product_attributes
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	ID                 string
	ProductName        string
	ProductThumb       string
	ProductDescription sql.NullString
	ProductPrice       string
	ProductQuantity    int32
	ProductType        ProductsProductType
	ProductShop        string
	ProductSlug        sql.NullString
	ProductAttributes  json.RawMessage
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.ProductName,
		arg.ProductThumb,
		arg.ProductDescription,
		arg.ProductPrice,
		arg.ProductQuantity,
		arg.ProductType,
		arg.ProductShop,
		arg.ProductSlug,
		arg.ProductAttributes,
	)
	return err
}

const getAllDraftsForShop = `-- name: GetAllDraftsForShop :many
SELECT id, product_name, product_thumb, product_description, product_price, product_quantity, product_type, product_shop, product_attributes, product_ratingaverage, product_variations, isdraft, ispublished, created_at, updated_at, product_slug FROM products
WHERE isDraft = 1 AND product_shop = ?
ORDER BY updated_at DESC
LIMIT ?
OFFSET ?
`

type GetAllDraftsForShopParams struct {
	ProductShop string
	Limit       int32
	Offset      int32
}

func (q *Queries) GetAllDraftsForShop(ctx context.Context, arg GetAllDraftsForShopParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getAllDraftsForShop, arg.ProductShop, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.ProductThumb,
			&i.ProductDescription,
			&i.ProductPrice,
			&i.ProductQuantity,
			&i.ProductType,
			&i.ProductShop,
			&i.ProductAttributes,
			&i.ProductRatingaverage,
			&i.ProductVariations,
			&i.Isdraft,
			&i.Ispublished,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductSlug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductByID = `-- name: GetProductByID :one
SELECT id, product_name, product_thumb, product_description, product_price, product_quantity, product_type, product_shop, product_attributes, product_ratingaverage, product_variations, isdraft, ispublished, created_at, updated_at, product_slug FROM products
WHERE id = ?
`

func (q *Queries) GetProductByID(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.ProductThumb,
		&i.ProductDescription,
		&i.ProductPrice,
		&i.ProductQuantity,
		&i.ProductType,
		&i.ProductShop,
		&i.ProductAttributes,
		&i.ProductRatingaverage,
		&i.ProductVariations,
		&i.Isdraft,
		&i.Ispublished,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductSlug,
	)
	return i, err
}
