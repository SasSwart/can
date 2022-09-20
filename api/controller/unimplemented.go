// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.com/sasswart/gin-in-a-can//api/models"
)

type UnimplementedServer struct{}


func (UnimplementedServer) UserDelete(c *gin.Context, requestModel *models.UserDeleteRequest) models.UserDeleteResponse {
  return models.UserDelete400Response{}
}
func (UnimplementedServer) UserGet(c *gin.Context, requestModel *models.UserGetRequest) models.UserGetResponse {
  return models.UserGet400Response{}
}
func (UnimplementedServer) UserPatch(c *gin.Context, requestModel *models.UserPatchRequest) models.UserPatchResponse {
  return models.UserPatch400Response{}
}
func (UnimplementedServer) UserPost(c *gin.Context, requestModel *models.UserPostRequest) models.UserPostResponse {
  return models.UserPost400Response{}
}
func (UnimplementedServer) ProjectDelete(c *gin.Context, requestModel *models.ProjectDeleteRequest) models.ProjectDeleteResponse {
  return models.ProjectDelete400Response{}
}
func (UnimplementedServer) ProjectGet(c *gin.Context, requestModel *models.ProjectGetRequest) models.ProjectGetResponse {
  return models.ProjectGet400Response{}
}
func (UnimplementedServer) ProjectPatch(c *gin.Context, requestModel *models.ProjectPatchRequest) models.ProjectPatchResponse {
  return models.ProjectPatch400Response{}
}
func (UnimplementedServer) ProjectPost(c *gin.Context, requestModel *models.ProjectPostRequest) models.ProjectPostResponse {
  return models.ProjectPost400Response{}
}