package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/services"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
)

type ProductController struct {
	productFactory services.IProductService
}


func NewProductController(productFactory services.IProductService) *ProductController {
	return &ProductController{
		productFactory: productFactory,
	}
}


func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var payload dtos.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ValidatorErrorResponse(ctx, response.ErrCodeBadRequest)
		return
	}

	data, code := pc.productFactory.CreateProduct(payload, payload.ProductType)
	response.SuccessResponse(ctx, code, data)
}