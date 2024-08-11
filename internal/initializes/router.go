package initializes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/global"
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

	MainGroup := r.Group("/api/v1") 
	{
		MainGroup.GET("/monitor", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}
	return r
}