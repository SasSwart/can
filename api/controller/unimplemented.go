// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.gom/sasswart/gin-in-a-can/api/models"
)

type UnimplementedServer struct{}


func (UnimplementedServer) UserDelete(requestModel *models.UserDeleteRequest) (models.UserDeleteResponse, error) {
  return models.UserDelete400Response{}, nil
}
func (UnimplementedServer) UserGet(requestModel *models.UserGetRequest) (models.UserGetResponse, error) {
  return models.UserGet400Response{}, nil
}
func (UnimplementedServer) UserPatch(requestModel *models.UserPatchRequest) (models.UserPatchResponse, error) {
  return models.UserPatch400Response{}, nil
}
func (UnimplementedServer) UserPost(requestModel *models.UserPostRequest) (models.UserPostResponse, error) {
  return models.UserPost400Response{}, nil
}
func (UnimplementedServer) ProjectDelete(requestModel *models.ProjectDeleteRequest) (models.ProjectDeleteResponse, error) {
  return models.ProjectDelete400Response{}, nil
}
func (UnimplementedServer) ProjectGet(requestModel *models.ProjectGetRequest) (models.ProjectGetResponse, error) {
  return models.ProjectGet400Response{}, nil
}
func (UnimplementedServer) ProjectPatch(requestModel *models.ProjectPatchRequest) (models.ProjectPatchResponse, error) {
  return models.ProjectPatch400Response{}, nil
}
func (UnimplementedServer) ProjectPost(requestModel *models.ProjectPostRequest) (models.ProjectPostResponse, error) {
  return models.ProjectPost400Response{}, nil
}