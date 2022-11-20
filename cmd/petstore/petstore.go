// GENERATED CODE. DO NOT EDIT

package main

import "github.com/gin-gonic/gin"

type Server interface {
  Pets
  PetsPetid
}

func RegisterServer(e *gin.Engine, srv Server) {
  RegisterPetsPath(e, srv)
  RegisterPetsPetidPath(e, srv)
}

type UnimplementedServer struct {
  UnimplementedPets
  UnimplementedPetsPetid
}

// GENERATED INTERFACE. DO NOT EDIT

type Pets interface {
  Petsget(*gin.Context, *PetsgetParameters, *Petsgetrequestbody) PetsgetResponse
  Petspost(*gin.Context, *PetspostParameters, *Petspostrequestbody) PetspostResponse
}

type UnimplementedPets struct {}
func (u UnimplementedPets) Petsget(*gin.Context, *PetsgetParameters, *Petsgetrequestbody) PetsgetResponse {
	return Petsget405Response{}
}
func (u UnimplementedPets) Petspost(*gin.Context, *PetspostParameters, *Petspostrequestbody) PetspostResponse {
	return Petspost405Response{}
}

func RegisterPetsPath(e *gin.Engine, srv Pets) {
  e.GET("Petsget", func(c *gin.Context) {
  	params := &PetsgetParameters{}
  	body := &Petsgetrequestbody{}
    response := srv.Petsget(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
  e.POST("Petspost", func(c *gin.Context) {
  	params := &PetspostParameters{}
  	body := &Petspostrequestbody{}
    response := srv.Petspost(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
}
type PetsgetResponse interface {
	GetStatus() int
}

type Petsget200Response struct {
	Petsget200responsemodel
}
func (Petsget200Response) GetStatus() int {
	return 200
}

type PetsgetdefaultResponse struct {
	Petsgetdefaultresponsemodel
}
func (PetsgetdefaultResponse) GetStatus() int {
	return default
}
type Petsget405Response struct {}
func (Petsget405Response) GetStatus() int {
	return 405
}

type PetsgetParameters struct {}
type Petsgetrequestbody struct {}
type PetspostResponse interface {
	GetStatus() int
}


type PetspostdefaultResponse struct {
	Petspostdefaultresponsemodel
}
func (PetspostdefaultResponse) GetStatus() int {
	return default
}
type Petspost405Response struct {}
func (Petspost405Response) GetStatus() int {
	return 405
}

type PetspostParameters struct {}
type Petspostrequestbody struct {}


// GENERATED MODEL. DO NOT EDIT

type Petsgetdefaultresponsemodel struct {
  code Petsgetdefaultresponsemodelcode
  message Petsgetdefaultresponsemodelmessage
}


// GENERATED MODEL. DO NOT EDIT

type Petsgetdefaultresponsemodelcode int


// GENERATED MODEL. DO NOT EDIT

type Petsgetdefaultresponsemodelmessage string


// GENERATED MODEL. DO NOT EDIT

type Petsget200responsemodel []Petsget200responsemodelitem


// GENERATED MODEL. DO NOT EDIT

type Petsget200responsemodelitem struct {
  id Petsget200responsemodelitemid
  name Petsget200responsemodelitemname
  tag Petsget200responsemodelitemtag
}


// GENERATED MODEL. DO NOT EDIT

type Petsget200responsemodelitemid int


// GENERATED MODEL. DO NOT EDIT

type Petsget200responsemodelitemname string


// GENERATED MODEL. DO NOT EDIT

type Petsget200responsemodelitemtag string


// GENERATED MODEL. DO NOT EDIT

type Petspostdefaultresponsemodel struct {
  code Petspostdefaultresponsemodelcode
  message Petspostdefaultresponsemodelmessage
}


// GENERATED MODEL. DO NOT EDIT

type Petspostdefaultresponsemodelcode int


// GENERATED MODEL. DO NOT EDIT

type Petspostdefaultresponsemodelmessage string


// GENERATED INTERFACE. DO NOT EDIT

type PetsPetid interface {
  Petspetidget(*gin.Context, *PetspetidgetParameters, *Petspetidgetrequestbody) PetspetidgetResponse
}

type UnimplementedPetsPetid struct {}
func (u UnimplementedPetsPetid) Petspetidget(*gin.Context, *PetspetidgetParameters, *Petspetidgetrequestbody) PetspetidgetResponse {
	return Petspetidget405Response{}
}

func RegisterPetsPetidPath(e *gin.Engine, srv PetsPetid) {
  e.GET("Petspetidget", func(c *gin.Context) {
  	params := &PetspetidgetParameters{}
  	body := &Petspetidgetrequestbody{}
    response := srv.Petspetidget(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
}
type PetspetidgetResponse interface {
	GetStatus() int
}

type Petspetidget200Response struct {
	Petspetidget200responsemodel
}
func (Petspetidget200Response) GetStatus() int {
	return 200
}

type PetspetidgetdefaultResponse struct {
	Petspetidgetdefaultresponsemodel
}
func (PetspetidgetdefaultResponse) GetStatus() int {
	return default
}
type Petspetidget405Response struct {}
func (Petspetidget405Response) GetStatus() int {
	return 405
}

type PetspetidgetParameters struct {}
type Petspetidgetrequestbody struct {}


// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodel struct {
  code Petspetidgetdefaultresponsemodelcode
  message Petspetidgetdefaultresponsemodelmessage
}


// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodelmessage string


// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodelcode int


// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodel struct {
  id Petspetidget200responsemodelid
  name Petspetidget200responsemodelname
  tag Petspetidget200responsemodeltag
}


// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodelid int


// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodelname string


// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodeltag string

