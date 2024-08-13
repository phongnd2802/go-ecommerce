package services

import (
	"encoding/json"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"strconv"
)

type productFactory struct {
	productRepo  repositories.IProductRepository
	ProductTypes map[string]IProduct
}

// GetAllPublishedForShop implements IProductService.
func (pf *productFactory) GetAllPublishedForShop(productShop string, options ...int) ([]dtos.ProductResponse, int) {
	limit, skip := 50, 0
	if len(options) > 0 {
		limit = options[0]
		if len(options) > 1 {
			skip = options[1]
		}
	}
	isDraft, isPublished := false, true
	products, err := pf.productRepo.QueryProductForShop(productShop, isDraft, isPublished, limit, skip)
	if err != nil {
		return nil, response.ErrCodeFailedQueryDB
	}

	productResponses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		productPrice, _ := strconv.ParseFloat(product.ProductPrice, 64)
		var productAttributes map[string]any
		_ = json.Unmarshal(product.ProductAttributes, &productAttributes)

		var productVariations []string
		_ = json.Unmarshal(product.ProductVariations, &productVariations)

		productResponses[i] = dtos.ProductResponse{
			ID:                   product.ID,
			ProductName:          product.ProductName,
			ProductThumb:         product.ProductThumb,
			ProductDescription:   &product.ProductDescription.String,
			ProductPrice:         float32(productPrice),
			ProductQuantity:      int(product.ProductQuantity),
			ProductType:          string(product.ProductType),
			ProductShop:          product.ProductShop,
			ProductAttributes:    productAttributes,
			ProductVariations:    productVariations,
			ProductRatingAverage: product.ProductRatingaverage.String,
			CreatedAt:            product.CreatedAt.Time,
			UpdatedAt:            product.UpdatedAt.Time,
		}
	}
	return productResponses, response.CodeSuccess
}

// GetAllDraftsForShop implements IProductService.
func (pf *productFactory) GetAllDraftsForShop(productShop string, options ...int) ([]dtos.ProductResponse, int) {
	limit, skip := 50, 0
	if len(options) > 0 {
		limit = options[0]
		if len(options) > 1 {
			skip = options[1]
		}
	}
	isDraft, isPublished := true, false
	products, err := pf.productRepo.QueryProductForShop(productShop, isDraft, isPublished, limit, skip)
	if err != nil {
		return nil, response.ErrCodeFailedQueryDB
	}

	productResponses := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		productPrice, _ := strconv.ParseFloat(product.ProductPrice, 64)
		var productAttributes map[string]any
		_ = json.Unmarshal(product.ProductAttributes, &productAttributes)

		var productVariations []string
		_ = json.Unmarshal(product.ProductVariations, &productVariations)

		productResponses[i] = dtos.ProductResponse{
			ID:                   product.ID,
			ProductName:          product.ProductName,
			ProductThumb:         product.ProductThumb,
			ProductDescription:   &product.ProductDescription.String,
			ProductPrice:         float32(productPrice),
			ProductQuantity:      int(product.ProductQuantity),
			ProductType:          string(product.ProductType),
			ProductShop:          product.ProductShop,
			ProductAttributes:    productAttributes,
			ProductVariations:    productVariations,
			ProductRatingAverage: product.ProductRatingaverage.String,
			CreatedAt:            product.CreatedAt.Time,
			UpdatedAt:            product.UpdatedAt.Time,
		}
	}
	return productResponses, response.CodeSuccess
}

// PublishProductByShop implements IProductService.
func (pf *productFactory) PublishProductByShop(productShop string, productID string) (*dtos.ProductUpdateResponse, int) {
	foundProduct, err := pf.productRepo.GetProductByShopAndID(productShop, productID)
	if err != nil {
		return nil, response.ErrCodeNotFoundProduct
	}
	isDraft, isPublished := false, true
	_ = pf.productRepo.UpdatedStatusProductByShop(foundProduct.ID, isDraft, isPublished)
	return &dtos.ProductUpdateResponse{ID: foundProduct.ID}, response.CodePublishProductSuccess
}

// UnPublishProductByShop implements IProductService.
func (pf *productFactory) UnPublishProductByShop(productShop string, productID string) (*dtos.ProductUpdateResponse, int) {
	foundProduct, err := pf.productRepo.GetProductByShopAndID(productShop, productID)
	if err != nil {
		return nil, response.ErrCodeNotFoundProduct
	}
	isDraft, isPublished := true, false
	_ = pf.productRepo.UpdatedStatusProductByShop(foundProduct.ID, isDraft, isPublished)
	return &dtos.ProductUpdateResponse{ID: foundProduct.ID}, response.CodeUnPublishProductSuccess
}

// CreateProduct implements IProductService.
func (pf *productFactory) CreateProduct(payload dtos.ProductCreateRequest, productType string, productShop string) (*dtos.ProductCreateResponse, int) {
	productTypeRef, ok := pf.ProductTypes[productType]
	if !ok {
		return nil, response.ErrCodeInvalidProductType
	}

	result, err := productTypeRef.CreateProduct(payload, productShop)
	if err != nil {
		return nil, response.ErrCodeFailedInsertDB
	}

	productPrice, _ := strconv.ParseFloat(result.ProductPrice, 64)
	var productAttributes map[string]any
	_ = json.Unmarshal(result.ProductAttributes, &productAttributes)

	var productVariations []string
	_ = json.Unmarshal(result.ProductVariations, &productVariations)
	return &dtos.ProductCreateResponse{
		ProductResponse: dtos.ProductResponse{
			ID:                   result.ID,
			ProductName:          result.ProductName,
			ProductThumb:         result.ProductThumb,
			ProductDescription:   &result.ProductDescription.String,
			ProductPrice:         float32(productPrice),
			ProductQuantity:      int(result.ProductQuantity),
			ProductType:          string(result.ProductType),
			ProductShop:          result.ProductShop,
			ProductAttributes:    productAttributes,
			ProductVariations:    productVariations,
			ProductRatingAverage: result.ProductRatingaverage.String,
			ProductSlug:          result.ProductSlug.String,
			CreatedAt:            result.CreatedAt.Time,
			UpdatedAt:            result.UpdatedAt.Time,
		},
	}, response.CodeCreated

}

func NewProductFactory(
	productRepo repositories.IProductRepository,
	clothingRepo repositories.IClothingRepository,
	electronicRepo repositories.IElectronicsRepository,
	furnitureRepo repositories.IFurnitureRepository,
) IProductService {
	productTypes := make(map[string]IProduct)
	product := NewProduct(productRepo)

	// Register Product Type
	RegisterProductType(productTypes, "Clothing", NewClothing(product, clothingRepo))
	RegisterProductType(productTypes, "Electronics", NewElectronic(product, electronicRepo))
	RegisterProductType(productTypes, "Furniture", NewFurniture(product, furnitureRepo))
	return &productFactory{
		productRepo:  productRepo,
		ProductTypes: productTypes,
	}
}

func RegisterProductType(productTypes map[string]IProduct, productTypeString string, productTypeRef IProduct) {
	productTypes[productTypeString] = productTypeRef
}
