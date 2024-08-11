package initializes

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/phongnd2802/go-ecommerce/internal/routers"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
)

func initRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		r = gin.New()
	}

	// Middlewares


	// routers
	managerRouter := routers.RouterApp.Manage
	userRouter := routers.RouterApp.User
	MainGroup := r.Group("/api/v1") 
	{
		MainGroup.GET("/monitor", func(ctx *gin.Context) {
			response.SuccessResponse(ctx, response.CodeSuccess, "OK")
		})
	}
	{
		userRouter.InitAccessRouter(MainGroup)
	}
	{
		managerRouter.InitAccessRouter(MainGroup)
	}
	return r
}