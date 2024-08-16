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
	QueryProductForShop(productShop string, isDraft bool, isPublished bool, limit int, skip int) ([]database.Product, error)
	UpdatedStatusProductByShop(productID string, isDraft bool, isPublished bool) error
	GetProductByShopAndID(productShop, productID string) (*database.Product, error)
	UpdateProductByID(
		productID string, productName string, productThumb string, productDescription *string,
		productPrice float32, productQuantity int, productType string, productSlug string, productAttributes map[string]any,
	) (*database.Product, error)
}

type productRepository struct {
	db *database.Store
}

// UpdateProductByID implements IProductRepository.
func (pr *productRepository) UpdateProductByID(productID string, productName string, productThumb string, productDescription *string, productPrice float32, productQuantity int, productType string, productSlug string, productAttributes map[string]any) (*database.Product, error) {
	productAttributesJSON, err := json.Marshal(productAttributes)
	if err != nil {
		return nil, err
	}
	err = pr.db.UpdateProduct(context.Background(), database.UpdateProductParams{
		ProductName: productName,
		ProductThumb: productThumb,
		ProductType: database.ProductsProductType(productType),
		ProductDescription: sql.NullString{
			String: *productDescription,
			Valid: true,
		},
		ProductSlug: sql.NullString{
			String: productSlug,
			Valid: true,
		},
		ProductPrice: fmt.Sprintf("%.2f", productPrice),
		ProductQuantity: int32(productQuantity),
		ProductAttributes: productAttributesJSON,
		ID: productID,
	})
	if err != nil {
		return nil, err
	}
	result, _ := pr.db.Queries.GetProductByID(context.Background(), productID)
	return &result, nil
}

// GetProductByShopAndID implements IProductRepository.
func (pr *productRepository) GetProductByShopAndID(productShop string, productID string) (*database.Product, error) {
	product, err := pr.db.Queries.GetProductByShopAndID(context.Background(), database.GetProductByShopAndIDParams{
		ProductShop: productShop,
		ID:          productID,
	})
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdatedStatusProductByShop implements IProductRepository.
func (pr *productRepository) UpdatedStatusProductByShop(productID string, isDraft bool, isPublished bool) error {
	err := pr.db.UpdateStatusProductByShop(context.Background(), database.UpdateStatusProductByShopParams{
		ID: productID,
		Ispublished: sql.NullBool{
			Bool:  isPublished,
			Valid: true,
		},
		Isdraft: sql.NullBool{
			Bool:  isDraft,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// QueryProductForShop implements IProductRepository.
func (pr *productRepository) QueryProductForShop(productShop string, isDraft bool, isPublished bool, limit int, skip int) ([]database.Product, error) {
	products, err := pr.db.Queries.QueryProrductForShop(context.Background(), database.QueryProrductForShopParams{
		ProductShop: productShop,
		Limit:       int32(limit),
		Offset:      int32(skip),
		Isdraft: sql.NullBool{
			Bool:  isDraft,
			Valid: true,
		},
		Ispublished: sql.NullBool{
			Bool:  isPublished,
			Valid: true,
		},
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
		db: db,
	}
}
