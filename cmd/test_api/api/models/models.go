// GENERATED CODE. DO NOT EDIT

package models

type StudentGetRequest struct {
}

func (model StudentGetRequest) IsValid() error {

	return nil
}

type StudentGetResponse interface {
	ImplementsStudentGetResponse()
	GetStatus() int
}

type StudentGet200Response struct {
	Address    *Address `json:"address"`
	First_name *string  `json:"first_name"`
}

func (StudentGet200Response) ImplementsStudentGetResponse() {}

func (StudentGet200Response) GetStatus() int {
	return 200
}

type First_name struct {
}

func (model First_name) IsValid() error {

	return nil
}

type Address struct {
	City        *string `json:"city"`
	Postal_code *string `json:"postal_code"`
}

func (model Address) IsValid() error {

	return nil
}

type City struct {
}

func (model City) IsValid() error {

	return nil
}

type Postal_code struct {
}

func (model Postal_code) IsValid() error {

	return nil
}

type Test_api struct {
	Address    *Address `json:"address"`
	First_name *string  `json:"first_name"`
}

func (model Test_api) IsValid() error {

	return nil
}
