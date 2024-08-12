package repositories

import (
	"context"

	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IElectronicsRepository interface {
	CreateElectronic(id, manufacturer, model, color, productShop string) (*database.Electronic, error)
}

type electronicRepo struct {
	db *database.Store
}

// CreateElectronic implements IElectronicsRepository.
func (e *electronicRepo) CreateElectronic(id string, manufacturer string, model string, color string, productShop string) (*database.Electronic, error) {
	err := e.db.CreateElectronic(context.Background(), database.CreateElectronicParams{
		ID: id,
		Manufacturer: manufacturer,
		Model: model,
		Color: color,
		ProductShop: productShop,
	})
	if err != nil {
		return nil, err
	}
	result, _ := e.db.Queries.GetElectronicByID(context.Background(), id)
	return &result, err
}

func NewElectronicRepository(db *database.Store) IElectronicsRepository {
	return &electronicRepo{
		db: db,
	}
}
