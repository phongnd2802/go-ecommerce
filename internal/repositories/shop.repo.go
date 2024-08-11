package repositories

import (
	"context"

	"github.com/google/uuid"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type IShopRepository interface {
	GetShopByEmail(email string) (*database.Shop, error)
	CreateShop(shopName, email, password string) (*database.Shop, error)
}

type shopRepository struct {
	db *database.Store
}

// CreateShop implements IShopRepository.
func (sr *shopRepository) CreateShop(shopName, email, password string) (*database.Shop, error) {
	var result database.Shop
	id := uuid.New().String()
	err := sr.db.CreateShop(context.Background(), database.CreateShopParams{
		ID: id,
		ShopName: shopName,
		Email: email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	result, err = sr.db.GetShopByID(context.Background(), id)
	if err != nil {
		return nil, err
	}


	return &result, err

}

// GetShopByEmail implements IShopRepository.
func (sr *shopRepository) GetShopByEmail(email string) (*database.Shop, error) {
	result, err := sr.db.GetShopByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewShopRepository(db *database.Store) IShopRepository {
	return &shopRepository{db: db}
}
