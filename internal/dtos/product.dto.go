package dtos

import "time"

type (
	ProductCreateRequest struct {
		ProductName        string         `json:"product_name"`
		ProductThumb       string         `json:"product_thumb"`
		ProductDescription string         `json:"product_description"`
		ProductPrice       float32        `json:"product_price"`
		ProductQuantity    int            `json:"product_quantity"`
		ProductType        string         `json:"product_type"`
		ProductShop        string         `json:"product_shop"`
		ProductAttributes  map[string]any `json:"product_attributes"`
	}

	ProductResponse struct {
		ID                   string         `json:"id"`
		ProductName          string         `json:"product_name"`
		ProductThumb         string         `json:"product_thumb"`
		ProductDescription   string         `json:"product_description"`
		ProductPrice         float32        `json:"product_price"`
		ProductQuantity      int            `json:"product_quantity"`
		ProductType          string         `json:"product_type"`
		ProductShop          string         `json:"product_shop"`
		ProductAttributes    map[string]any `json:"product_attributes"`
		ProductRatingAverage string         `json:"product_rating_avg"`
		CreatedAt            time.Time      `json:"created_at"`
		UpdatedAt            time.Time      `json:"updated_at"`
	}

	ProductCreateResponse struct {
		ProductResponse `json:"product"`
	}
)
