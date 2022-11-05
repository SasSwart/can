// GENERATED CODE. DO NOT EDIT

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sasswart/gin-in-a-can/cmd/test_api/api/models"
)

type UnimplementedServer struct{}

func (UnimplementedServer) HeartbeatGet(c *gin.Context, requestModel *models.HeartbeatGetRequest) models.HeartbeatGetResponse {
	return models.HeartbeatGet200Response{}
}
