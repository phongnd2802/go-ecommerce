package services

import (
	"database/sql"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"github.com/phongnd2802/go-ecommerce/pkg/utils"
)

type IAccessService interface {
	SignUp(email string, password string) (*dtos.ShopResponse, int)
	Login(email string, password string) (*dtos.ShopResponse, int)
	Logout(keyStoreID string) int
	HandleRefreshToken(shop jwt.MapClaims, keyStore database.Token, refreshToken string) (*dtos.ShopResponse, int)
}

type accessService struct {
	shopRepo  repositories.IShopRepository
	tokenRepo repositories.ITokenRepository
}

// HandleRefreshToken implements IAccessService.
func (as *accessService) HandleRefreshToken(shop jwt.MapClaims, keyStore database.Token, refreshToken string) (*dtos.ShopResponse, int) {
	shopID := shop["sub"].(string)
	email := shop["email"].(string)
	if keyStore.RefreshTokenUsed.String == refreshToken {
		err := as.tokenRepo.DeleteTokenByID(keyStore.ID)
		if err != nil {
			return nil, response.ErrCodeInternalServer
		}
		return nil, response.ErrCodeRefreshTokenUsed
	} 

	if keyStore.RefreshToken != refreshToken {
		return nil, response.ErrCodeShopNotExist
	}

	foundShop, err := as.shopRepo.GetShopByEmail(email)
	if err != nil {
		return nil, response.ErrCodeShopNotExist
	}

	payload := map[string]any{
		"id": shopID,
		"email": email,
	}
	privateKey, publicKey, _ := utils.GenerateRSAKeyPair(2048)
	newAccessToken, _ := utils.CreateAccessToken(payload, privateKey)
	newRefreshToken, _ := utils.CreateRefreshToken(payload, privateKey)

	err = as.tokenRepo.UpdateTokenByID(newRefreshToken, refreshToken, publicKey, keyStore.ID)
	if err != nil {
		return nil, response.ErrCodeInternalServer
	}

	return &dtos.ShopResponse{
		ID: foundShop.ID,
		Name: foundShop.ShopName,
		Email: foundShop.Email,
		IsActive: foundShop.IsActive.Bool,
		CreatedAt: foundShop.CreatedAt.Time,
		UpdatedAt: foundShop.UpdatedAt.Time,
		Tokens: dtos.TokenReponse{
			AccessToken: newAccessToken,
			RefreshToken: newRefreshToken,
		},
	}, response.CodeSuccess


}

// Logout implements IAccessService.
func (as *accessService) Logout(keyStoreID string) int {
	err := as.tokenRepo.DeleteTokenByID(keyStoreID)
	if err != nil {
		return response.ErrCodeInternalServer
	}
	return response.CodeSuccess
}

// Login implements IAccessService.
func (as *accessService) Login(email string, password string) (*dtos.ShopResponse, int) {
	foundShop, err := as.shopRepo.GetShopByEmail(email)
	if err == sql.ErrNoRows {
		return nil, response.ErrCodeEmailOrPasswordIncorrect
	}

	match := utils.CheckPasswordHash(password, foundShop.Password)
	if !match {
		return nil, response.ErrCodeEmailOrPasswordIncorrect
	}
	privateKey, publicKey, _ := utils.GenerateRSAKeyPair(2048)

	payload := map[string]any{
		"id":    foundShop.ID,
		"email": foundShop.Email,
	}
	accessToken, _ := utils.CreateAccessToken(payload, privateKey)
	refreshToken, _ := utils.CreateRefreshToken(payload, privateKey)

	_, err = as.tokenRepo.CreateKeyToken(publicKey, refreshToken, foundShop.ID)
	if err != nil {
		return nil, response.ErrCodeInternalServer
	}

	return &dtos.ShopResponse{
		ID:        foundShop.ID,
		Name:      foundShop.ShopName,
		Email:     foundShop.Email,
		IsActive:  foundShop.IsActive.Bool,
		CreatedAt: foundShop.CreatedAt.Time,
		UpdatedAt: foundShop.UpdatedAt.Time,
		Tokens: dtos.TokenReponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, response.CodeSuccess
}

// SignUp implements IAccessService.
func (as *accessService) SignUp(email string, password string) (*dtos.ShopResponse, int) {
	_, err := as.shopRepo.GetShopByEmail(email)
	if err == sql.ErrNoRows {
		emailSplitted := strings.Split(email, "@")
		shopName := emailSplitted[0]
		passwordHash, _ := utils.HashPassword(password)
		newShop, err := as.shopRepo.CreateShop(shopName, email, passwordHash)
		if err != nil {
			return nil, response.ErrCodeInternalServer
		}

		privateKey, publicKey, _ := utils.GenerateRSAKeyPair(2048)

		payload := map[string]any{
			"id":    newShop.ID,
			"email": newShop.Email,
		}
		accessToken, _ := utils.CreateAccessToken(payload, privateKey)
		refreshToken, _ := utils.CreateRefreshToken(payload, privateKey)

		_, err = as.tokenRepo.CreateKeyToken(publicKey, refreshToken, newShop.ID)
		if err != nil {
			return nil, response.ErrCodeInternalServer
		}

		return &dtos.ShopResponse{
			ID:        newShop.ID,
			Name:      newShop.ShopName,
			Email:     newShop.Email,
			IsActive:  newShop.IsActive.Bool,
			CreatedAt: newShop.CreatedAt.Time,
			UpdatedAt: newShop.UpdatedAt.Time,
			Tokens: dtos.TokenReponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}, response.CodeSuccess
	}

	return nil, response.ErrCodeShopExist
}

func NewAccessService(
	shopRepo repositories.IShopRepository,
	tokenRepo repositories.ITokenRepository,
) IAccessService {
	return &accessService{
		tokenRepo: tokenRepo,
		shopRepo:  shopRepo,
	}
}
