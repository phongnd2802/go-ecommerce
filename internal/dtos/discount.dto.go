package dtos

import "time"

type (
	DiscountCreateRequest struct {
		Code          string    `json:"discount_code" binding:"required"`
		StartDate     time.Time `json:"start_date" binding:"required"`
		EndDate       time.Time `json:"end_date" binding:"required"`
		IsActive      bool      `json:"is_active" binding:"required"`
		ShopID        string    `json:"shop_id" binding:"required"`
		MinOrderValue int       `json:"min_order_value" binding:"required"`
		ProductIDs    []string  `json:"product_ids" binding:"required"`
	}

	DiscountResponse struct {}
)
