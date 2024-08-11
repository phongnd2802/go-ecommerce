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
	ctx.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

func ValidatorErrorResponse(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusBadRequest, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})

	defer ctx.AbortWithStatus(http.StatusBadRequest)
}





