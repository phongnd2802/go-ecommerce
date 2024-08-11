package services

import (
	"database/sql"
	"strings"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"github.com/phongnd2802/go-ecommerce/pkg/utils"
)

type IAccessService interface {
	SignUp(email string, password string) (*dtos.ShopResponse, int)
	Login(email string, password string) (*dtos.ShopResponse, int)
}

type accessService struct {
	shopRepo repositories.IShopRepository
	tokenRepo repositories.ITokenRepository
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
		"id": foundShop.ID,
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
			AccessToken: accessToken,
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
			"id": newShop.ID,
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
				AccessToken: accessToken,
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
		shopRepo: shopRepo,
	}
}
