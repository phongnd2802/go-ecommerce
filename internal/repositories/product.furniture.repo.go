package repositories

import (
	"context"

	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IFurnitureRepository interface {
	CreateFurniture(id, brand, size, material, productShop string) (*database.Furniture, error)
}

type furnitureRepo struct {
	db *database.Store
}

// CreateFurniture implements IFurnitureRepository.
func (fr *furnitureRepo) CreateFurniture(id string, brand string, size string, material string, productShop string) (*database.Furniture, error) {
	err := fr.db.CreateFurniture(context.Background(), database.CreateFurnitureParams{
		ID: id,
		Brand: brand,
		Size: size,
		Material: material,
		ProductShop: productShop,
	})
	if err != nil {
		return nil, err
	}

	result, _ := fr.db.Queries.GetFurnitureByID(context.Background(), id)
	return &result, nil
}

func NewFurnitureRepository(db *database.Store) IFurnitureRepository {
	return &furnitureRepo{
		db: db,
	}
}
