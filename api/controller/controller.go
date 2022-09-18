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
  NetworkDelete(*models.NetworkDeleteRequest) (models.NetworkDeleteResponse, error)
  NetworkGet(*models.NetworkGetRequest) (models.NetworkGetResponse, error)
  NetworkPatch(*models.NetworkPatchRequest) (models.NetworkPatchResponse, error)
  NetworkPost(*models.NetworkPostRequest) (models.NetworkPostResponse, error)
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user", func(c *gin.Context) {
		request := &models.UserDeleteRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.GET("/user", func(c *gin.Context) {
		request := &models.UserGetRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.PATCH("/user", func(c *gin.Context) {
		request := &models.UserPatchRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.POST("/user", func(c *gin.Context) {
		request := &models.UserPostRequest{
		}
		if err := request.IsValid(); err != nil {
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
	e.PATCH("/project", func(c *gin.Context) {
		request := &models.ProjectPatchRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.POST("/project", func(c *gin.Context) {
		request := &models.ProjectPostRequest{
		}
		if err := request.IsValid(); err != nil {
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
	e.DELETE("/project", func(c *gin.Context) {
		request := &models.ProjectDeleteRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.GET("/project", func(c *gin.Context) {
		request := &models.ProjectGetRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
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
	e.DELETE("/network", func(c *gin.Context) {
		request := &models.NetworkDeleteRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.NetworkDelete(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.GET("/network", func(c *gin.Context) {
		request := &models.NetworkGetRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.NetworkGet(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.PATCH("/network", func(c *gin.Context) {
		request := &models.NetworkPatchRequest{
			Id: c.Query("id"),
		}
		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.NetworkPatch(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
	e.POST("/network", func(c *gin.Context) {
		request := &models.NetworkPostRequest{
		}
		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := srv.NetworkPost(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	})
}

