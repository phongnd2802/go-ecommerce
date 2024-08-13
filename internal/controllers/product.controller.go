package controllers

import (
	"strconv"
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


func (pc *ProductController) GetAllPublishedForShop(ctx *gin.Context) {
	shopID := ctx.Request.Header.Get("x-client-id")
	limit := ctx.Query("limit")
	if limit == "" {
		limit = "50"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	skip := ctx.Query("skip")
	if skip == "" {
		skip = "0"
	}
	skipInt, err := strconv.Atoi(skip)
	if err != nil {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	data, code := pc.productFactory.GetAllPublishedForShop(shopID, limitInt, skipInt)
	response.SuccessResponse(ctx, code, data)
}

func (pc *ProductController) UnPublishProductByShop(ctx *gin.Context) {
	productID := ctx.Param("id")
	shopID := ctx.Request.Header.Get("x-client-id")
	data, code := pc.productFactory.UnPublishProductByShop(shopID, productID)
	response.SuccessResponse(ctx, code, data)
}

func (pc *ProductController) PublishProductByShop(ctx *gin.Context) {
	productID := ctx.Param("id")
	shopID := ctx.Request.Header.Get("x-client-id")
	data, code := pc.productFactory.PublishProductByShop(shopID, productID)
	response.SuccessResponse(ctx, code, data)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var payload dtos.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ValidatorErrorResponse(ctx, response.ErrCodeBadRequest)
		return 
	}
	shopID := ctx.Request.Header.Get("x-client-id")
	data, code := pc.productFactory.CreateProduct(payload, payload.ProductType, shopID)
	response.SuccessResponse(ctx, code, data)
}

func (pc *ProductController) GetAllDraftsForShop(ctx *gin.Context) {
	shopID := ctx.Request.Header.Get("x-client-id")
	limit := ctx.Query("limit")
	if limit == "" {
		limit = "50"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	skip := ctx.Query("skip")
	if skip == "" {
		skip = "0"
	}
	skipInt, err := strconv.Atoi(skip)
	if err != nil {
		response.ErrorResposne(ctx, response.ErrCodeBadRequest)
		return
	}
	data, code := pc.productFactory.GetAllDraftsForShop(shopID, limitInt, skipInt)
	response.SuccessResponse(ctx, code, data)
}