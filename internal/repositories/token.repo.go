package repositories

import (
	"context"

	"github.com/google/uuid"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
)

type ITokenRepository interface {
	CreateKeyToken(publicKey, refreshToken string, shopID string) (*database.Token, error)
	DeleteTokenByID(id string) error
}

type tokenRepository struct {
	db *database.Store
}

// DeleteTokenByID implements ITokenRepository.
func (tr *tokenRepository) DeleteTokenByID(id string) error {
	err := tr.db.DeleteTokenByID(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

// CreateKeyToken implements ITokenRepository.
func (tr *tokenRepository) CreateKeyToken(publicKey, refreshToken, shopID string) (*database.Token, error) {
	id := uuid.New().String()
	err := tr.db.CreateToken(context.Background(), database.CreateTokenParams{
		ID:           id,
		PublicKey:    publicKey,
		RefreshToken: refreshToken,
		ShopID:       shopID,
	})

	if err != nil {
		return nil, err
	}
	var result database.Token
	result, err = tr.db.GetTokenByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewTokenRepository(db *database.Store) ITokenRepository {
	return &tokenRepository{
		db: db,
	}
}
