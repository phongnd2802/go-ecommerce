package services

import (
	"database/sql"
	"strings"
	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"github.com/phongnd2802/go-ecommerce/pkg/utils"
)

type IAccessService interface {
	SignUp(email string, password string) (*dtos.ShopResponse, int)
}

type accessService struct{
	shopRepo repositories.IShopRepository
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
			global.Logger.Error(err.Error())
			return nil, response.ErrCodeInternalServer
		}

		return &dtos.ShopResponse{
			ID: newShop.ID,
			Name: newShop.ShopName,
			Email: newShop.Email,
			IsActive: newShop.IsActive.Bool,
			CreatedAt: newShop.CreatedAt.Time,
			UpdatedAt: newShop.UpdatedAt.Time,
		}, response.CodeCreated
	} 

	return nil, response.ErrCodeShopExist
}

func NewAccessService(shopRepo repositories.IShopRepository) IAccessService {
	return &accessService{
		shopRepo: shopRepo,
	}
}
