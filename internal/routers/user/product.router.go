package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/phongnd2802/go-ecommerce/internal/middlewares"
	"github.com/phongnd2802/go-ecommerce/internal/wire"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productController, _ := wire.InitProductRouterHandler(global.Db)

	// public Router
	productPublicRouter := Router.Group("/product")
	{
		productPublicRouter.GET("/search")
	}

	productPrivateRouter := Router.Group("/product")
	productPrivateRouter.Use(middlewares.Authentication())
	{
		productPrivateRouter.POST("/", productController.CreateProduct)

		productPrivateRouter.PATCH("/:id", productController.UpdateProduct)
		productPrivateRouter.PATCH("/publish/:id", productController.PublishProductByShop)
		productPrivateRouter.PATCH("/unpublish/:id", productController.UnPublishProductByShop)

		productPrivateRouter.GET("/drafts/all", productController.GetAllDraftsForShop)
		productPrivateRouter.GET("/published/all", productController.GetAllPublishedForShop)
		productPrivateRouter.GET("/:id", productController.GetProductByID)
	}

}