package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-ecommerce/global"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
)

const (
	API_KEY        = "x-api-key"
	AUTHORIZATION = "authorization"
)

const (
	OBJKEY = "objkey"
)


func ApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var db = database.NewStore(global.Db)
		key := ctx.Request.Header.Get(API_KEY)
		// db.Queries.CreateApiKey(context.Background(), database.CreateApiKeyParams{
		// 	Akey: utils.GenerateRandomSecretKeyBase64(32),
		// 	Permissions: "0000",
		// })
		if key == "" {
			response.ForbiddenError(ctx, response.ErrCodeForbidden)
			return
		}

		objKey, err := db.Queries.GetApiKey(context.Background(), key)
		if err != nil {
			response.ForbiddenError(ctx, response.ErrCodeForbidden)
			return
		}

		ctx.Set(OBJKEY, objKey)
		ctx.Next()
	}
}

func Permissions(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exist := ctx.Get(OBJKEY)
		if !exist {
			response.ForbiddenError(ctx, response.ErrCodeForbidden)
			return
		}
		objKey, ok := value.(database.ApiKey)
		if !ok {
			response.ForbiddenError(ctx, response.ErrCodeForbidden)
			return
		}

		if objKey.Permissions != permission {
			response.ForbiddenError(ctx, response.ErrCodeForbidden)
			return
		}

		ctx.Next()

	}
}