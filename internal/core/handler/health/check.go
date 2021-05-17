package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
