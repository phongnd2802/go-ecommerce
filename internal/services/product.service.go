package services

import "github.com/phongnd2802/go-ecommerce/internal/dtos"

type IProductService interface {
	CreateProduct(payload dtos.ProductCreateRequest, productType string, productShop string) (*dtos.ProductCreateResponse, int)
	GetAllDraftsForShop(productShop string, options ...int) ([]dtos.ProductCreateResponse, int)
	PublishProductByShop(productShop, productID string)
	UnPublishProductByShop(productShop, productID string)
}



