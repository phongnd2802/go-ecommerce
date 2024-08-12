package repositories

import (
	"context"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IClothingRepository interface {
	CreateClothing(id, brand, size, material, productShop string) (*database.Clothe, error)
}

type clothingRepository struct {
	db *database.Store
}

// CreateClothing implements IClothingRepository.
func (cr *clothingRepository) CreateClothing(id string, brand string, size string, material, productShop string) (*database.Clothe, error) {
	err := cr.db.CreateClothing(context.Background(), database.CreateClothingParams{
		ID: id,
		Brand: brand,
		Size: size,
		ProductShop: productShop,
	})
	if err != nil {
		return nil, err
	}

	result, _ := cr.db.Queries.GetClothingByID(context.Background(), id)
	return &result, nil
}

func NewClothingRepository(db *database.Store) IClothingRepository {
	return &clothingRepository{
		db: db,
	}
}
