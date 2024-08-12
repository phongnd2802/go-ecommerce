package services

import (
	"encoding/json"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"strconv"
)

type productFactory struct {
	ProductTypes map[string]IProduct
}

// CreateProduct implements IProductService.
func (p *productFactory) CreateProduct(payload dtos.ProductCreateRequest, productType string) (*dtos.ProductCreateResponse, int) {
	productTypeRef, ok := p.ProductTypes[productType]
	if !ok {
		return nil, response.ErrCodeInvalidProductType
	}

	result, err := productTypeRef.CreateProduct(payload)
	if err != nil {
		return nil, response.ErrCodeFailedInsertDB
	}

	productPrice, _ := strconv.ParseFloat(result.ProductPrice, 64)
	var productAttributes map[string]any
	_ = json.Unmarshal(result.ProductAttributes, &productAttributes)
	return &dtos.ProductCreateResponse{
		ProductResponse: dtos.ProductResponse{
			ID:                   result.ID,
			ProductName:          result.ProductName,
			ProductThumb:         result.ProductThumb,
			ProductDescription:   result.ProductDescription.String,
			ProductPrice:         float32(productPrice),
			ProductQuantity:      int(result.ProductQuantity),
			ProductType:          string(result.ProductType),
			ProductShop:          result.ProductShop,
			ProductAttributes:    productAttributes,
			ProductRatingAverage: result.ProductRatingaverage.String,
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
		ProductTypes: productTypes,
	}
}

func RegisterProductType(productTypes map[string]IProduct, productTypeString string, productTypeRef IProduct) {
	productTypes[productTypeString] = productTypeRef	
}


