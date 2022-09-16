// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.gom/sasswart/gin-in-a-can/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  UserDelete(*models.UserDeleteRequest) (models.UserDeleteResponse, error)
  UserGet(*models.UserGetRequest) (models.UserGetResponse, error)
  UserPatch(*models.UserPatchRequest) (models.UserPatchResponse, error)
  UserPost(*models.UserPostRequest) (models.UserPostResponse, error)
  ProjectDelete(*models.ProjectDeleteRequest) (models.ProjectDeleteResponse, error)
  ProjectGet(*models.ProjectGetRequest) (models.ProjectGetResponse, error)
  ProjectPatch(*models.ProjectPatchRequest) (models.ProjectPatchResponse, error)
  ProjectPost(*models.ProjectPostRequest) (models.ProjectPostResponse, error)
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user/:name", func(c *gin.Context) {
		request := &models.UserDeleteRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserDelete(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.GET("/user/:name", func(c *gin.Context) {
		request := &models.UserGetRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserGet(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/user/:name", func(c *gin.Context) {
		request := &models.UserPatchRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserPatch(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.POST("/user/:name", func(c *gin.Context) {
		request := &models.UserPostRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.UserPost(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.DELETE("/project/:name", func(c *gin.Context) {
		request := &models.ProjectDeleteRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectDelete(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.GET("/project/:name", func(c *gin.Context) {
		request := &models.ProjectGetRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectGet(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/project/:name", func(c *gin.Context) {
		request := &models.ProjectPatchRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectPatch(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.POST("/project/:name", func(c *gin.Context) {
		request := &models.ProjectPostRequest{
			NAME: c.Param("name"),
		}
		if err := request.Valid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.ProjectPost(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
}

