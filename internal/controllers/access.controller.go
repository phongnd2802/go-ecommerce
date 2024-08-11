package controllers

import (
	"github.com/gin-gonic/gin"
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
