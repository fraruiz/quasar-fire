package topsecret

import (
	"errors"
	"net/http"

	core "github.com/franciscoruizar/quasar-fire/internal"
	"github.com/franciscoruizar/quasar-fire/internal/platform/server/handler"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/gin-gonic/gin"
)

type CreateRequests struct {
	Satellites []CreateRequest `json:"satellites" binding:"required"`
}

type CreateRequest struct {
	Name     string   `json:"name" binding:"required"`
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type CreateResponse struct {
	Position handler.PositionResponse `json:"position" binding:"required"`
	Message  string                   `json:"message" binding:"required"`
}

func TopSecretCreateHandler(service usecases.TopSecretCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateRequests
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		var requests []usecases.TopSecretCreatorRequest
		for i := 0; i < len(req.Satellites); i++ {
			requests = append(requests, usecases.TopSecretCreatorRequest{
				Name:      req.Satellites[i].Name,
				Dinstance: req.Satellites[i].Distance,
				Message:   req.Satellites[i].Message,
			})
		}

		response, err := service.Create(requests)

		if err != nil {
			switch {
			case errors.Is(err, core.ErrInvalidSateliteID):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		responseMap := CreateResponse{
			Message:  response.Message,
			Position: handler.NewPositionResponse(response.Position.X, response.Position.Y),
		}

		c.Status(http.StatusCreated)

		c.JSON(http.StatusCreated, responseMap)
	}
}
