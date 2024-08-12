package services

import (
	"encoding/json"
	"strconv"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
)

type productFactory struct {
	Clothing *clothing
	Electronic *electronic
	Furniture *furniture
}

// Create Product
func (pf *productFactory) CreateProduct(
	payload dtos.ProductCreateRequest,
	productType string,
) (*dtos.ProductCreateResponse, int) {
	if productType == "Clothing" {
		result, err := pf.Clothing.CreateProduct(payload)
		if err != nil {
			return nil, response.ErrCodeBadRequest
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
	} else if productType == "Electronics" {
			result, err := pf.Electronic.CreateProduct(payload)
			if err != nil {
				return nil, response.ErrCodeBadRequest
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
	} else if productType == "Furniture" {
		result, err := pf.Furniture.CreateProduct(payload)
		if err != nil {
			return nil, response.ErrCodeBadRequest
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
	return nil, response.ErrCodeBadRequest
}


func NewProductFactory(
	productRepo repositories.IProductRepository,
	clothingRepo repositories.IClothingRepository,
	electronicRepo repositories.IElectronicsRepository,
	furnitureRepo repositories.IFurnitureRepository,
) IProductService {
	product := NewProduct(productRepo)
	return &productFactory{
		Clothing: NewClothing(product, clothingRepo),
		Electronic: NewElectronic(product, electronicRepo),
		Furniture: NewFurniture(product, furnitureRepo),
	}
}
