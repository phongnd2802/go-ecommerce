package controllers

import (
	"github.com/gin-gonic/gin"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/services"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
)

type AccessController struct {
	accessService services.IAccessService
}


func NewAccessController(accessService services.IAccessService) *AccessController {
	return &AccessController{
		accessService: accessService,
	}
}

func (ac *AccessController) Logout(ctx *gin.Context) {
	value, exist := ctx.Get("keystore")
	if !exist {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	keyStore, ok := value.(database.Token)
	if !ok {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	code := ac.accessService.Logout(keyStore.ID)
	response.SuccessResponse(ctx, code, nil)
}


func (ac *AccessController) Login(ctx *gin.Context) {
	var payload dtos.ShopRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ValidatorErrorResponse(ctx, response.ErrCodeBadRequest)
		return
	}

	data, code := ac.accessService.Login(payload.Email, payload.Password)
	response.SuccessResponse(ctx, code, data)
}

func (ac *AccessController) SignUp(ctx *gin.Context) {
	var payload dtos.ShopRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ValidatorErrorResponse(ctx, response.ErrCodeBadRequest)
		return
	}

	data, code := ac.accessService.SignUp(payload.Email, payload.Password)
	response.SuccessResponse(ctx, code, data)
}
