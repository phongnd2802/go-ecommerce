package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data any) {
	if code >= 40000 {
		ErrorResposne(ctx, code)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})

}

func ForbiddenResponse(ctx *gin.Context, code int) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
	
}

func NotFoundReponse(ctx *gin.Context, code int) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}


func InternalServerReponse(ctx *gin.Context, code int) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}

func ErrorResposne(ctx *gin.Context, code int) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}


func ValidatorErrorResponse(ctx *gin.Context, code int) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseData{
		Code: code,
		Message: msg[code],
		Data: "Invalid Request",
	})

}





