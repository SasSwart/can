// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.com/sasswart/gin-in-a-can/cmd/test_api/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  StudentGet(*gin.Context, *models.StudentGetRequest) models.StudentGetResponse
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.GET("/student", func(c *gin.Context) {
		request := &models.StudentGetRequest{}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.StudentGet(c, request)
		c.JSON(response.GetStatus(), response)
	})
}

