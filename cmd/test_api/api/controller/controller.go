// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.com/sasswart/gin-in-a-can/cmd/test_api/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  HeartbeatGet(*gin.Context, *models.HeartbeatGetRequest) models.HeartbeatGetResponse
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.GET("/heartbeat", func(c *gin.Context) {
		request := &models.HeartbeatGetRequest{}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.HeartbeatGet(c, request)
		c.JSON(response.GetStatus(), response)
	})
}

