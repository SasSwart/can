package api

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	UserDelete(*gin.Context, *UserDeleteRequest)
	UserGet(*gin.Context, *UserGetRequest)
	UserPatch(*gin.Context, *UserPatchRequest)
	UserPost(*gin.Context, *UserPostRequest)
	ProjectGet(*gin.Context, *ProjectGetRequest)
	ProjectPatch(*gin.Context, *ProjectPatchRequest)
	ProjectPost(*gin.Context, *ProjectPostRequest)
	ProjectDelete(*gin.Context, *ProjectDeleteRequest)
}

type UnimplementedServer struct{}

func (UnimplementedServer) UserDelete(*gin.Context, *UserDeleteRequest)       {}
func (UnimplementedServer) UserGet(*gin.Context, *UserGetRequest)             {}
func (UnimplementedServer) UserPatch(*gin.Context, *UserPatchRequest)         {}
func (UnimplementedServer) UserPost(*gin.Context, *UserPostRequest)           {}
func (UnimplementedServer) ProjectGet(*gin.Context, *ProjectGetRequest)       {}
func (UnimplementedServer) ProjectPatch(*gin.Context, *ProjectPatchRequest)   {}
func (UnimplementedServer) ProjectPost(*gin.Context, *ProjectPostRequest)     {}
func (UnimplementedServer) ProjectDelete(*gin.Context, *ProjectDeleteRequest) {}

type UserDeleteRequest struct{}
type UserGetRequest struct{}
type UserPatchRequest struct{}
type UserPostRequest struct{}
type ProjectGetRequest struct{}
type ProjectPatchRequest struct{}
type ProjectPostRequest struct{}
type ProjectDeleteRequest struct{}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user/:name", func(c *gin.Context) {
		srv.UserDelete(c, &UserDeleteRequest{})
	})
	e.GET("/user/:name", func(c *gin.Context) {
		srv.UserGet(c, &UserGetRequest{})
	})
	e.PATCH("/user/:name", func(c *gin.Context) {
		srv.UserPatch(c, &UserPatchRequest{})
	})
	e.POST("/user/:name", func(c *gin.Context) {
		srv.UserPost(c, &UserPostRequest{})
	})
	e.GET("/project/:id", func(c *gin.Context) {
		srv.ProjectGet(c, &ProjectGetRequest{})
	})
	e.PATCH("/project/:id", func(c *gin.Context) {
		srv.ProjectPatch(c, &ProjectPatchRequest{})
	})
	e.POST("/project/:id", func(c *gin.Context) {
		srv.ProjectPost(c, &ProjectPostRequest{})
	})
	e.DELETE("/project/:id", func(c *gin.Context) {
		srv.ProjectDelete(c, &ProjectDeleteRequest{})
	})
}
