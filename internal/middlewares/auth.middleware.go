package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/phongnd2802/go-ecommerce/global"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/pkg/response"
	"github.com/phongnd2802/go-ecommerce/pkg/utils"
)

const (
	API_KEY       = "x-api-key"
	CLIENT_ID     = "x-client-id"
	AUTHORIZATION = "authorization"
	REFRESHTOKEN  = "x-rtoken-id"
)

const (
	OBJKEY   = "objkey"
	KEYSTORE = "keystore"
	SHOP     = "shop"
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
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}

		objKey, err := db.Queries.GetApiKey(context.Background(), key)
		if err != nil {
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
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
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}
		objKey, ok := value.(database.ApiKey)
		if !ok {
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}

		if objKey.Permissions != permission {
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}

		ctx.Next()

	}
}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.NewStore(global.Db)
		shopID := ctx.Request.Header.Get(CLIENT_ID)
		if shopID == "" {
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}
		keyStore, err := db.Queries.GetTokenByShopID(context.Background(), shopID)
		if err != nil {
			response.NotFoundReponse(ctx, response.ErrCodeForbidden)
			return
		}

		refreshToken := ctx.Request.Header.Get(REFRESHTOKEN)
		if refreshToken != "" {
			decodeShop, err := utils.VerifyToken(refreshToken, keyStore.PublicKey)
			if err != nil {
				response.InternalServerReponse(ctx, response.ErrCodeFailedVerifyJWT)
				return
			}

			if claims, ok := decodeShop.Claims.(jwt.MapClaims); ok && decodeShop.Valid {
				sub := claims["sub"].(string)
				if sub != shopID {
					response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
					return
				}

				ctx.Set(KEYSTORE, keyStore)
				ctx.Set(SHOP, claims)
				ctx.Set(REFRESHTOKEN, refreshToken) 
				ctx.Next()
				return
			}

		}

		accessToken := ctx.Request.Header.Get(AUTHORIZATION)
		if accessToken == "" {
			response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
			return
		}

		verifiedToken, err := utils.VerifyToken(accessToken, keyStore.PublicKey)
		if err != nil {
			response.InternalServerReponse(ctx, response.ErrCodeFailedVerifyJWT)
			return
		}

		if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok && verifiedToken.Valid {
			sub := claims["sub"].(string)
			if sub != shopID {
				response.ForbiddenResponse(ctx, response.ErrCodeForbidden)
				return
			}

			ctx.Set(KEYSTORE, keyStore)
			ctx.Next()
		} else {
			response.InternalServerReponse(ctx, response.ErrCodeFailedVerifyJWT)
			return
		}
	}
}
