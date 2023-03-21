package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareAbortResponse(context *gin.Context, response Response) {
	context.JSON(http.StatusBadRequest, response)
	context.Abort()
}
