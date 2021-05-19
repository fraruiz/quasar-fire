package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Message string `json:"message" binding:"required"`
}

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, HealthCheck{
			Message: "everything is ok!",
		})
	}
}
