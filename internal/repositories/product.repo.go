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
	GetAllDraftsForShop(productShop string, limit int, skip int) ([]database.Product, error)
}

type productRepository struct {
	db       *database.Store
}

// GetAllDraftsForShop implements IProductRepository.
func (pr *productRepository) GetAllDraftsForShop(productShop string, limit int, skip int) ([]database.Product, error) {
	products, err := pr.db.Queries.GetAllDraftsForShop(context.Background(), database.GetAllDraftsForShopParams{
		ProductShop: productShop,
		Limit: int32(limit),
		Offset: int32(skip),
	})
	if err != nil {
		return nil, err
	}
	return products, nil
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
		ID:           id,
		ProductName:  productName,
		ProductThumb: productThumb,
		ProductDescription: sql.NullString{
			String: *productDescription,
			Valid:  true,
		},
		ProductPrice:    fmt.Sprintf("%.2f", productPrice),
		ProductQuantity: int32(productQuantity),
		ProductType:     database.ProductsProductType(productType),
		ProductShop:     productShop,
		ProductSlug: sql.NullString{
			String: productSlug,
			Valid:  true,
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

func NewProductReposiroty(db *database.Store) IProductRepository {
	return &productRepository{
		db:       db,
	}
}
