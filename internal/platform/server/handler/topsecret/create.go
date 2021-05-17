package topsecret

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name     string   `json:"name" binding:"required"`
	Distance string   `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type CreateResponse struct {
	Position PositionResponse `json:"position" binding:"required"`
	Message  []string         `json:"message" binding:"required"`
}

type PositionResponse struct {
	X float64 `json:"x" binding:"required"`
	Y float64 `json:"y" binding:"required"`
}

func CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.Status(http.StatusCreated)
	}
}
