package api

import "github.com/gin-gonic/gin"

// GENERATED INTERFACE. DO NOT EDIT

type SwaggerPetstorePets interface {
  SwaggerPetstorePetsGet(*gin.Context, *SwaggerPetstorePetsGetParameters, *SwaggerPetstorePetsGetRequestBody) SwaggerPetstorePetsGetResponse
  SwaggerPetstorePetsPost(*gin.Context, *SwaggerPetstorePetsPostParameters, *SwaggerPetstorePetsPostRequestBody) SwaggerPetstorePetsPostResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedSwaggerPetstorePets struct {}
func (u UnimplementedSwaggerPetstorePets) SwaggerPetstorePetsGet(*gin.Context, *SwaggerPetstorePetsGetParameters, *SwaggerPetstorePetsGetRequestBody) SwaggerPetstorePetsGetResponse {
	return SwaggerPetstorePetsGet405Response{}
}
func (u UnimplementedSwaggerPetstorePets) SwaggerPetstorePetsPost(*gin.Context, *SwaggerPetstorePetsPostParameters, *SwaggerPetstorePetsPostRequestBody) SwaggerPetstorePetsPostResponse {
	return SwaggerPetstorePetsPost405Response{}
}
func (u UnimplementedSwaggerPetstorePets) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

func RegisterSwaggerPetstorePetsPath(e *gin.Engine, srv SwaggerPetstorePets) {
	
  e.GET("/pets", func(c *gin.Context) {
  	params := &SwaggerPetstorePetsGetParameters{}
  	body := &SwaggerPetstorePetsGetRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
    response := srv.SwaggerPetstorePetsGet(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
	
  e.POST("/pets", func(c *gin.Context) {
  	params := &SwaggerPetstorePetsPostParameters{}
  	body := &SwaggerPetstorePetsPostRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
    response := srv.SwaggerPetstorePetsPost(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
}
type SwaggerPetstorePetsGetResponse interface {
	GetStatus() int
}

type SwaggerPetstorePetsGet200Response SwaggerPetstorePetsGet200ResponseModel

func (SwaggerPetstorePetsGet200Response) GetStatus() int {
	return 200
}

type SwaggerPetstorePetsGet500Response SwaggerPetstorePetsGet500ResponseModel

func (SwaggerPetstorePetsGet500Response) GetStatus() int {
	return 500
}
type SwaggerPetstorePetsGet405Response struct {}
func (SwaggerPetstorePetsGet405Response) GetStatus() int {
	return 405
}
type SwaggerPetstorePetsPostResponse interface {
	GetStatus() int
}


type SwaggerPetstorePetsPost500Response SwaggerPetstorePetsPost500ResponseModel

func (SwaggerPetstorePetsPost500Response) GetStatus() int {
	return 500
}
type SwaggerPetstorePetsPost405Response struct {}
func (SwaggerPetstorePetsPost405Response) GetStatus() int {
	return 405
}
