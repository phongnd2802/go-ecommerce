package repositories

import (
	"context"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IClothingRepository interface {
	CreateClothing(id, brand, size, material, productShop string) (*database.Clothe, error)
	GetClothingByID(id string) (*database.Clothe, error)
	UpdateClothing(id, brand, size, material string) error
}

type clothingRepository struct {
	db *database.Store
}

// GetClothingByID implements IClothingRepository.
func (cr *clothingRepository) GetClothingByID(id string) (*database.Clothe, error) {
	result, err := cr.db.Queries.GetClothingByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateClothing implements IClothingRepository.
func (cr *clothingRepository) UpdateClothing(id string, brand string, size string, material string) error {
	err := cr.db.UpdateClothingByID(context.Background(), database.UpdateClothingByIDParams{
		Brand:    brand,
		Size:     size,
		Material: material,
		ID:       id,
	})
	if err != nil {
		return nil
	}

	return err
}

// CreateClothing implements IClothingRepository.
func (cr *clothingRepository) CreateClothing(id string, brand string, size string, material, productShop string) (*database.Clothe, error) {
	err := cr.db.CreateClothing(context.Background(), database.CreateClothingParams{
		ID:          id,
		Brand:       brand,
		Size:        size,
		Material: material,
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
