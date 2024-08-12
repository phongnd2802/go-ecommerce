package services

import "github.com/phongnd2802/go-ecommerce/internal/dtos"

type IProductService interface {
	CreateProduct(payload dtos.ProductCreateRequest, productType string) (*dtos.ProductCreateResponse, int)
}



