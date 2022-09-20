// GENERATED CODE. DO NOT EDIT

package controller

import (
    "github.com/sasswart/gin-in-a-can//api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
  UserDelete(*gin.Context, *models.UserDeleteRequest) models.UserDeleteResponse
  UserGet(*gin.Context, *models.UserGetRequest) models.UserGetResponse
  UserPatch(*gin.Context, *models.UserPatchRequest) models.UserPatchResponse
  UserPost(*gin.Context, *models.UserPostRequest) models.UserPostResponse
  ProjectDelete(*gin.Context, *models.ProjectDeleteRequest) models.ProjectDeleteResponse
  ProjectGet(*gin.Context, *models.ProjectGetRequest) models.ProjectGetResponse
  ProjectPatch(*gin.Context, *models.ProjectPatchRequest) models.ProjectPatchResponse
  ProjectPost(*gin.Context, *models.ProjectPostRequest) models.ProjectPostResponse
}

func RegisterServer(e *gin.Engine, srv Server) {
	e.DELETE("/user", func(c *gin.Context) {
		request := &models.UserDeleteRequest{
			Name: c.Query("name"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.UserDelete(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.GET("/user", func(c *gin.Context) {
		request := &models.UserGetRequest{
			Name: c.Query("name"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.UserGet(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.PATCH("/user", func(c *gin.Context) {
		request := &models.UserPatchRequest{
			Name: c.Query("name"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.UserPatch(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.POST("/user", func(c *gin.Context) {
		request := &models.UserPostRequest{
			Name: c.Query("name"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.UserPost(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.DELETE("/project/:id", func(c *gin.Context) {
		request := &models.ProjectDeleteRequest{
			Id: c.Param("id"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.ProjectDelete(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.GET("/project/:id", func(c *gin.Context) {
		request := &models.ProjectGetRequest{
			Id: c.Param("id"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.ProjectGet(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.PATCH("/project/:id", func(c *gin.Context) {
		request := &models.ProjectPatchRequest{
			Id: c.Param("id"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.ProjectPatch(c, request)
		c.JSON(response.GetStatus(), response)
	})
	e.POST("/project/:id", func(c *gin.Context) {
		request := &models.ProjectPostRequest{
			Id: c.Param("id"),
		}

		if err := c.ShouldBindJSON(&request.Body); err != nil {
		c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := request.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := srv.ProjectPost(c, request)
		c.JSON(response.GetStatus(), response)
	})
}

