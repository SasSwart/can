// GENERATED CODE. DO NOT EDIT

package models

type UserDeleteRequest struct {
	Name string
	Body UserDeleteRequestBody
}

func (r UserDeleteRequest)IsValid() error {
	return nil
}

type UserDeleteRequestBody struct {
}

type UserDeleteResponse interface {
	mustImplementUserDeleteResponse()
	GetStatus() int
}


type UserDelete204Response struct {
}
func (UserDelete204Response) mustImplementUserDeleteResponse(){}

func (UserDelete204Response) GetStatus() int {
	return 204
}

type UserDelete400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserDelete400Response) mustImplementUserDeleteResponse(){}

func (UserDelete400Response) GetStatus() int {
	return 400
}

type UserDelete500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserDelete500Response) mustImplementUserDeleteResponse(){}

func (UserDelete500Response) GetStatus() int {
	return 500
}

type UserGetRequest struct {
	Name string
	Body UserGetRequestBody
}

func (r UserGetRequest)IsValid() error {
	return nil
}

type UserGetRequestBody struct {
}

type UserGetResponse interface {
	mustImplementUserGetResponse()
	GetStatus() int
}


type UserGet200Response struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}
func (UserGet200Response) mustImplementUserGetResponse(){}

func (UserGet200Response) GetStatus() int {
	return 200
}

type UserGet400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserGet400Response) mustImplementUserGetResponse(){}

func (UserGet400Response) GetStatus() int {
	return 400
}

type UserGet500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserGet500Response) mustImplementUserGetResponse(){}

func (UserGet500Response) GetStatus() int {
	return 500
}

type UserPatchRequest struct {
	Name string
	Body UserPatchRequestBody
}

func (r UserPatchRequest)IsValid() error {
	return nil
}

type UserPatchRequestBody struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}

type UserPatchResponse interface {
	mustImplementUserPatchResponse()
	GetStatus() int
}


type UserPatch201Response struct {
}
func (UserPatch201Response) mustImplementUserPatchResponse(){}

func (UserPatch201Response) GetStatus() int {
	return 201
}

type UserPatch400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPatch400Response) mustImplementUserPatchResponse(){}

func (UserPatch400Response) GetStatus() int {
	return 400
}

type UserPatch500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPatch500Response) mustImplementUserPatchResponse(){}

func (UserPatch500Response) GetStatus() int {
	return 500
}

type UserPostRequest struct {
	Name string
	Body UserPostRequestBody
}

func (r UserPostRequest)IsValid() error {
	return nil
}

type UserPostRequestBody struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}

type UserPostResponse interface {
	mustImplementUserPostResponse()
	GetStatus() int
}


type UserPost204Response struct {
}
func (UserPost204Response) mustImplementUserPostResponse(){}

func (UserPost204Response) GetStatus() int {
	return 204
}

type UserPost400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPost400Response) mustImplementUserPostResponse(){}

func (UserPost400Response) GetStatus() int {
	return 400
}

type UserPost500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPost500Response) mustImplementUserPostResponse(){}

func (UserPost500Response) GetStatus() int {
	return 500
}

type ProjectDeleteRequest struct {
	Id string
	Body ProjectDeleteRequestBody
}

func (r ProjectDeleteRequest)IsValid() error {
	return nil
}

type ProjectDeleteRequestBody struct {
}

type ProjectDeleteResponse interface {
	mustImplementProjectDeleteResponse()
	GetStatus() int
}


type ProjectDelete200Response struct {
}
func (ProjectDelete200Response) mustImplementProjectDeleteResponse(){}

func (ProjectDelete200Response) GetStatus() int {
	return 200
}

type ProjectDelete400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectDelete400Response) mustImplementProjectDeleteResponse(){}

func (ProjectDelete400Response) GetStatus() int {
	return 400
}

type ProjectDelete500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectDelete500Response) mustImplementProjectDeleteResponse(){}

func (ProjectDelete500Response) GetStatus() int {
	return 500
}

type ProjectGetRequest struct {
	Id string
	Body ProjectGetRequestBody
}

func (r ProjectGetRequest)IsValid() error {
	return nil
}

type ProjectGetRequestBody struct {
}

type ProjectGetResponse interface {
	mustImplementProjectGetResponse()
	GetStatus() int
}


type ProjectGet200Response struct {
}
func (ProjectGet200Response) mustImplementProjectGetResponse(){}

func (ProjectGet200Response) GetStatus() int {
	return 200
}

type ProjectGet400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectGet400Response) mustImplementProjectGetResponse(){}

func (ProjectGet400Response) GetStatus() int {
	return 400
}

type ProjectGet500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectGet500Response) mustImplementProjectGetResponse(){}

func (ProjectGet500Response) GetStatus() int {
	return 500
}

type ProjectPatchRequest struct {
	Id string
	Body ProjectPatchRequestBody
}

func (r ProjectPatchRequest)IsValid() error {
	return nil
}

type ProjectPatchRequestBody struct {
}

type ProjectPatchResponse interface {
	mustImplementProjectPatchResponse()
	GetStatus() int
}


type ProjectPatch200Response struct {
}
func (ProjectPatch200Response) mustImplementProjectPatchResponse(){}

func (ProjectPatch200Response) GetStatus() int {
	return 200
}

type ProjectPatch400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPatch400Response) mustImplementProjectPatchResponse(){}

func (ProjectPatch400Response) GetStatus() int {
	return 400
}

type ProjectPatch500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPatch500Response) mustImplementProjectPatchResponse(){}

func (ProjectPatch500Response) GetStatus() int {
	return 500
}

type ProjectPostRequest struct {
	Id string
	Body ProjectPostRequestBody
}

func (r ProjectPostRequest)IsValid() error {
	return nil
}

type ProjectPostRequestBody struct {
}

type ProjectPostResponse interface {
	mustImplementProjectPostResponse()
	GetStatus() int
}


type ProjectPost200Response struct {
}
func (ProjectPost200Response) mustImplementProjectPostResponse(){}

func (ProjectPost200Response) GetStatus() int {
	return 200
}

type ProjectPost400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPost400Response) mustImplementProjectPostResponse(){}

func (ProjectPost400Response) GetStatus() int {
	return 400
}

type ProjectPost500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPost500Response) mustImplementProjectPostResponse(){}

func (ProjectPost500Response) GetStatus() int {
	return 500
}

