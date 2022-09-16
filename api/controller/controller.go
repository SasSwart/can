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
  ProjectPatch(*models.ProjectPatchRequest) (models.ProjectPatchResponse, error)
  ProjectPost(*models.ProjectPostRequest) (models.ProjectPostResponse, error)
  ProjectDelete(*models.ProjectDeleteRequest) (models.ProjectDeleteResponse, error)
  ProjectGet(*models.ProjectGetRequest) (models.ProjectGetResponse, error)
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user/", func(c *gin.Context) {
		request := &models.UserDeleteRequest{
			Name: c.Query("name"),
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
	e.GET("/user/", func(c *gin.Context) {
		request := &models.UserGetRequest{
			Name: c.Query("name"),
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
	e.PATCH("/user/", func(c *gin.Context) {
		request := &models.UserPatchRequest{
			Name: c.Query("name"),
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
	e.POST("/user/", func(c *gin.Context) {
		request := &models.UserPostRequest{
			Name: c.Query("name"),
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
	e.PATCH("/project/:id", func(c *gin.Context) {
		request := &models.ProjectPatchRequest{
			Id: c.Param("id"),
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
	e.POST("/project/:id", func(c *gin.Context) {
		request := &models.ProjectPostRequest{
			Id: c.Param("id"),
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
	e.DELETE("/project/:id", func(c *gin.Context) {
		request := &models.ProjectDeleteRequest{
			Id: c.Param("id"),
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
	e.GET("/project/:id", func(c *gin.Context) {
		request := &models.ProjectGetRequest{
			Id: c.Param("id"),
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
}

