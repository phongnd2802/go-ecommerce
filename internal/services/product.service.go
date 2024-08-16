package services

import (
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
)

type IProductService interface {
	CreateProduct(payload dtos.ProductCreateRequest, productType string, productShop string) (*dtos.ProductCreateResponse, int)
	GetAllDraftsForShop(productShop string, options ...int) ([]dtos.ProductResponse, int)
	GetAllPublishedForShop(productShop string, options ...int) ([]dtos.ProductResponse, int)
	GetProducByIDForShop(productShop, productID string) (*dtos.ProductResponse, int)
	PublishProductByShop(productShop, productID string) (*dtos.ProductUpdateResponse, int)
	UnPublishProductByShop(productShop, productID string) (*dtos.ProductUpdateResponse, int)
	UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productType string, productID string) (*dtos.ProductUpdateResponse, int)
}



