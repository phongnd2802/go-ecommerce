package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IProductRepository interface {
	CreateProduct(
		productName string, productThumb string, productDescription *string,
		productPrice float32, productQuantity int, productType string, productShop string,
		productSlug string, productAttributes map[string]any,
	) (*database.Product, error)
}

type productRepository struct {
	db *database.Store
	clothing IClothingRepository
}

// CreateProduct implements IProductRepository.
func (pr *productRepository) CreateProduct(
	productName string, productThumb string, productDescription *string, 
	productPrice float32, productQuantity int, productType string, 
	productShop string, productSlug string, productAttributes map[string]any,
) (*database.Product, error) {
	productAttributesJSON, err := json.Marshal(productAttributes)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	err = pr.db.CreateProduct(context.Background(), database.CreateProductParams{
		ID: id,
		ProductName: productName,
		ProductThumb: productThumb,
		ProductDescription: sql.NullString{
			String: *productDescription,
			Valid: true,
		},
		ProductPrice: fmt.Sprintf("%.2f", productPrice),
		ProductQuantity: int32(productQuantity),
		ProductType: database.ProductsProductType(productType),
		ProductShop: productShop,
		ProductSlug: sql.NullString{
			String: productSlug,
			Valid: true,
		},
		ProductAttributes: productAttributesJSON,
	})

	if err != nil {
		return nil, err
	}

	result, err := pr.db.Queries.GetProductByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewProductReposiroty(db *database.Store, clothing IClothingRepository) IProductRepository {
	return &productRepository{
		db: db,
		clothing: clothing,
	}
}
