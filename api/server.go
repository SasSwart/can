// GENERATED CODE. DO NOT EDIT

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  UserDelete(*gin.Context, *UserDeleteRequest) *UserDeleteResponse
  UserGet(*gin.Context, *UserGetRequest) *UserGetResponse
  UserPatch(*gin.Context, *UserPatchRequest) *UserPatchResponse
  UserPost(*gin.Context, *UserPostRequest) *UserPostResponse
  ProjectDelete(*gin.Context, *ProjectDeleteRequest) *ProjectDeleteResponse
  ProjectGet(*gin.Context, *ProjectGetRequest) *ProjectGetResponse
  ProjectPatch(*gin.Context, *ProjectPatchRequest) *ProjectPatchResponse
  ProjectPost(*gin.Context, *ProjectPostRequest) *ProjectPostResponse
}

type UnimplementedServer struct{}

func (UnimplementedServer) UserDelete(*gin.Context, *UserDeleteRequest) *UserDeleteResponse {
  return &UserDeleteResponse{}
}
func (UnimplementedServer) UserGet(*gin.Context, *UserGetRequest) *UserGetResponse {
  return &UserGetResponse{}
}
func (UnimplementedServer) UserPatch(*gin.Context, *UserPatchRequest) *UserPatchResponse {
  return &UserPatchResponse{}
}
func (UnimplementedServer) UserPost(*gin.Context, *UserPostRequest) *UserPostResponse {
  return &UserPostResponse{}
}
func (UnimplementedServer) ProjectDelete(*gin.Context, *ProjectDeleteRequest) *ProjectDeleteResponse {
  return &ProjectDeleteResponse{}
}
func (UnimplementedServer) ProjectGet(*gin.Context, *ProjectGetRequest) *ProjectGetResponse {
  return &ProjectGetResponse{}
}
func (UnimplementedServer) ProjectPatch(*gin.Context, *ProjectPatchRequest) *ProjectPatchResponse {
  return &ProjectPatchResponse{}
}
func (UnimplementedServer) ProjectPost(*gin.Context, *ProjectPostRequest) *ProjectPostResponse {
  return &ProjectPostResponse{}
}


type UserDeleteRequest struct {}
type UserDeleteResponse struct {}
type UserGetRequest struct {}
type UserGetResponse struct {}
type UserPatchRequest struct {}
type UserPatchResponse struct {}
type UserPostRequest struct {}
type UserPostResponse struct {}
type ProjectDeleteRequest struct {}
type ProjectDeleteResponse struct {}
type ProjectGetRequest struct {}
type ProjectGetResponse struct {}
type ProjectPatchRequest struct {}
type ProjectPatchResponse struct {}
type ProjectPostRequest struct {}
type ProjectPostResponse struct {}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user/:name", func(c *gin.Context) {
		response := srv.UserDelete(c, &UserDeleteRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.GET("/user/:name", func(c *gin.Context) {
		response := srv.UserGet(c, &UserGetRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/user/:name", func(c *gin.Context) {
		response := srv.UserPatch(c, &UserPatchRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.POST("/user/:name", func(c *gin.Context) {
		response := srv.UserPost(c, &UserPostRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.DELETE("/project/:id", func(c *gin.Context) {
		response := srv.ProjectDelete(c, &ProjectDeleteRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.GET("/project/:id", func(c *gin.Context) {
		response := srv.ProjectGet(c, &ProjectGetRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/project/:id", func(c *gin.Context) {
		response := srv.ProjectPatch(c, &ProjectPatchRequest{})
		c.JSON(http.StatusOK, response)
	})
	e.POST("/project/:id", func(c *gin.Context) {
		response := srv.ProjectPost(c, &ProjectPostRequest{})
		c.JSON(http.StatusOK, response)
	})
}

