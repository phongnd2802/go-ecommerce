package dtos

import "time"

type (
	ProductCreateRequest struct {
		ProductName        string         `json:"product_name" binding:"required"`
		ProductThumb       string         `json:"product_thumb" binding:"required"`
		ProductDescription string         `json:"product_description" binding:"required"`
		ProductPrice       float32        `json:"product_price" binding:"required"`
		ProductQuantity    int            `json:"product_quantity" binding:"required"`
		ProductType        string         `json:"product_type" binding:"required"`
		ProductAttributes  map[string]any `json:"product_attributes" binding:"required"`
	}

	ProductResponse struct {
		ID                   string         `json:"id"`
		ProductName          string         `json:"product_name"`
		ProductThumb         string         `json:"product_thumb"`
		ProductDescription   *string        `json:"product_description"`
		ProductPrice         float32        `json:"product_price"`
		ProductQuantity      int            `json:"product_quantity"`
		ProductType          string         `json:"product_type"`
		ProductShop          string         `json:"product_shop"`
		ProductAttributes    map[string]any `json:"product_attributes"`
		ProductRatingAverage string         `json:"product_rating_avg"`
		ProductVariations    []string       `json:"product_variations"`
		ProductSlug          string         `json:"product_slug"`
		CreatedAt            time.Time      `json:"created_at"`
		UpdatedAt            time.Time      `json:"updated_at"`
	}

	ProductCreateResponse struct {
		ProductResponse `json:"product"`
	}
)
