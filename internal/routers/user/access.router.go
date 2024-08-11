package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/phongnd2802/go-ecommerce/internal/middlewares"
	"github.com/phongnd2802/go-ecommerce/internal/wire"
)

type AccessRouter struct{}

func (ar *AccessRouter) InitAccessRouter(Router *gin.RouterGroup) {
	accessController, _ := wire.InitAccessRouterHandler(global.Db)

	// Public Router
	accessRouterPublic := Router.Group("/shop")
	{
		accessRouterPublic.POST("/signup", accessController.SignUp)
		accessRouterPublic.POST("/login", accessController.Login)
	}

	// Private Router
	accessRouterPrivate := Router.Group("/shop")
	accessRouterPrivate.Use(middlewares.Authentication())
	{
		accessRouterPrivate.POST("/logout", accessController.Logout)
	}
}