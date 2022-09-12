// GENERATED CODE. DO NOT EDIT

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  UserDelete(*gin.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
  UserGet(*gin.Context, *UserGetRequest) (*UserGetResponse, error)
  UserPatch(*gin.Context, *UserPatchRequest) (*UserPatchResponse, error)
  UserPost(*gin.Context, *UserPostRequest) (*UserPostResponse, error)
  ProjectGet(*gin.Context, *ProjectGetRequest) (*ProjectGetResponse, error)
  ProjectPatch(*gin.Context, *ProjectPatchRequest) (*ProjectPatchResponse, error)
  ProjectPost(*gin.Context, *ProjectPostRequest) (*ProjectPostResponse, error)
  ProjectDelete(*gin.Context, *ProjectDeleteRequest) (*ProjectDeleteResponse, error)
}

type UnimplementedServer struct{}

func (UnimplementedServer) UserDelete(*gin.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
  return &UserDeleteResponse{}, nil
}
func (UnimplementedServer) UserGet(*gin.Context, *UserGetRequest) (*UserGetResponse, error) {
  return &UserGetResponse{}, nil
}
func (UnimplementedServer) UserPatch(*gin.Context, *UserPatchRequest) (*UserPatchResponse, error) {
  return &UserPatchResponse{}, nil
}
func (UnimplementedServer) UserPost(*gin.Context, *UserPostRequest) (*UserPostResponse, error) {
  return &UserPostResponse{}, nil
}
func (UnimplementedServer) ProjectGet(*gin.Context, *ProjectGetRequest) (*ProjectGetResponse, error) {
  return &ProjectGetResponse{}, nil
}
func (UnimplementedServer) ProjectPatch(*gin.Context, *ProjectPatchRequest) (*ProjectPatchResponse, error) {
  return &ProjectPatchResponse{}, nil
}
func (UnimplementedServer) ProjectPost(*gin.Context, *ProjectPostRequest) (*ProjectPostResponse, error) {
  return &ProjectPostResponse{}, nil
}
func (UnimplementedServer) ProjectDelete(*gin.Context, *ProjectDeleteRequest) (*ProjectDeleteResponse, error) {
  return &ProjectDeleteResponse{}, nil
}


type UserDeleteRequest struct {
	name string
}

func (req UserDeleteRequest)valid() error {
	return nil
}

type UserDeleteResponse struct {}
type UserGetRequest struct {
	name string
}

func (req UserGetRequest)valid() error {
	return nil
}

type UserGetResponse struct {}
type UserPatchRequest struct {
	name string
}

func (req UserPatchRequest)valid() error {
	return nil
}

type UserPatchResponse struct {}
type UserPostRequest struct {
	name string
}

func (req UserPostRequest)valid() error {
	return nil
}

type UserPostResponse struct {}
type ProjectGetRequest struct {
	name string
}

func (req ProjectGetRequest)valid() error {
	return nil
}

type ProjectGetResponse struct {}
type ProjectPatchRequest struct {
	name string
}

func (req ProjectPatchRequest)valid() error {
	return nil
}

type ProjectPatchResponse struct {}
type ProjectPostRequest struct {
	name string
}

func (req ProjectPostRequest)valid() error {
	return nil
}

type ProjectPostResponse struct {}
type ProjectDeleteRequest struct {
	name string
}

func (req ProjectDeleteRequest)valid() error {
	return nil
}

type ProjectDeleteResponse struct {}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user/:name", func(c *gin.Context) {
		request := &UserDeleteRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserDelete(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.GET("/user/:name", func(c *gin.Context) {
		request := &UserGetRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserGet(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/user/:name", func(c *gin.Context) {
		request := &UserPatchRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserPatch(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.POST("/user/:name", func(c *gin.Context) {
		request := &UserPostRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserPost(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.GET("/project/:id", func(c *gin.Context) {
		request := &ProjectGetRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectGet(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/project/:id", func(c *gin.Context) {
		request := &ProjectPatchRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectPatch(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.POST("/project/:id", func(c *gin.Context) {
		request := &ProjectPostRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectPost(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.DELETE("/project/:id", func(c *gin.Context) {
		request := &ProjectDeleteRequest{
			name: c.Param("name"),
		}
		if err := request.valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectDelete(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
}

