// GENERATED CODE. DO NOT EDIT

package models

type UserDeleteRequest struct {
	Id string
	Body UserDeleteRequestBody
}

func (r UserDeleteRequest)IsValid() error {
	return nil
}

type UserDeleteRequestBody struct {
}

type UserDeleteResponse interface {
	mustImplementUserDeleteResponse()
}


type UserDelete204Response struct {
}
func (UserDelete204Response) mustImplementUserDeleteResponse(){}

type UserDelete400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserDelete400Response) mustImplementUserDeleteResponse(){}

type UserDelete404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserDelete404Response) mustImplementUserDeleteResponse(){}

type UserDelete500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserDelete500Response) mustImplementUserDeleteResponse(){}

type UserGetRequest struct {
	Id string
	Body UserGetRequestBody
}

func (r UserGetRequest)IsValid() error {
	return nil
}

type UserGetRequestBody struct {
}

type UserGetResponse interface {
	mustImplementUserGetResponse()
}


type UserGet200Response struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}
func (UserGet200Response) mustImplementUserGetResponse(){}

type UserGet400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserGet400Response) mustImplementUserGetResponse(){}

type UserGet404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserGet404Response) mustImplementUserGetResponse(){}

type UserGet500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserGet500Response) mustImplementUserGetResponse(){}

type UserPatchRequest struct {
	Id string
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
}


type UserPatch204Response struct {
}
func (UserPatch204Response) mustImplementUserPatchResponse(){}

type UserPatch400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserPatch400Response) mustImplementUserPatchResponse(){}

type UserPatch404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserPatch404Response) mustImplementUserPatchResponse(){}

type UserPatch500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserPatch500Response) mustImplementUserPatchResponse(){}

type UserPostRequest struct {
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
}


type UserPost201Response struct {
}
func (UserPost201Response) mustImplementUserPostResponse(){}

type UserPost400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserPost400Response) mustImplementUserPostResponse(){}

type UserPost500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (UserPost500Response) mustImplementUserPostResponse(){}

type ProjectPatchRequest struct {
	Id string
	Body ProjectPatchRequestBody
}

func (r ProjectPatchRequest)IsValid() error {
	return nil
}

type ProjectPatchRequestBody struct {
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
}

type ProjectPatchResponse interface {
	mustImplementProjectPatchResponse()
}


type ProjectPatch204Response struct {
}
func (ProjectPatch204Response) mustImplementProjectPatchResponse(){}

type ProjectPatch400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectPatch400Response) mustImplementProjectPatchResponse(){}

type ProjectPatch404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectPatch404Response) mustImplementProjectPatchResponse(){}

type ProjectPatch500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectPatch500Response) mustImplementProjectPatchResponse(){}

type ProjectPostRequest struct {
	Body ProjectPostRequestBody
}

func (r ProjectPostRequest)IsValid() error {
	return nil
}

type ProjectPostRequestBody struct {
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
}

type ProjectPostResponse interface {
	mustImplementProjectPostResponse()
}


type ProjectPost201Response struct {
}
func (ProjectPost201Response) mustImplementProjectPostResponse(){}

type ProjectPost400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectPost400Response) mustImplementProjectPostResponse(){}

type ProjectPost500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectPost500Response) mustImplementProjectPostResponse(){}

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
}


type ProjectDelete204Response struct {
}
func (ProjectDelete204Response) mustImplementProjectDeleteResponse(){}

type ProjectDelete400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectDelete400Response) mustImplementProjectDeleteResponse(){}

type ProjectDelete404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectDelete404Response) mustImplementProjectDeleteResponse(){}

type ProjectDelete500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectDelete500Response) mustImplementProjectDeleteResponse(){}

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
}


type ProjectGet200Response struct {
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
}
func (ProjectGet200Response) mustImplementProjectGetResponse(){}

type ProjectGet400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectGet400Response) mustImplementProjectGetResponse(){}

type ProjectGet404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectGet404Response) mustImplementProjectGetResponse(){}

type ProjectGet500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (ProjectGet500Response) mustImplementProjectGetResponse(){}

type NetworkDeleteRequest struct {
	Id string
	Body NetworkDeleteRequestBody
}

func (r NetworkDeleteRequest)IsValid() error {
	return nil
}

type NetworkDeleteRequestBody struct {
}

type NetworkDeleteResponse interface {
	mustImplementNetworkDeleteResponse()
}


type NetworkDelete204Response struct {
}
func (NetworkDelete204Response) mustImplementNetworkDeleteResponse(){}

type NetworkDelete400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkDelete400Response) mustImplementNetworkDeleteResponse(){}

type NetworkDelete404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkDelete404Response) mustImplementNetworkDeleteResponse(){}

type NetworkDelete500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkDelete500Response) mustImplementNetworkDeleteResponse(){}

type NetworkGetRequest struct {
	Id string
	Body NetworkGetRequestBody
}

func (r NetworkGetRequest)IsValid() error {
	return nil
}

type NetworkGetRequestBody struct {
}

type NetworkGetResponse interface {
	mustImplementNetworkGetResponse()
}


type NetworkGet200Response struct {
	Admin_state_up bool `json:"admin_state_up"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Dns_domain string `json:"dns_domain"`
	Ipv4_address_scope string `json:"ipv4_address_scope"`
	Ipv6_address_scope string `json:"ipv6_address_scope"`
	Is_default bool `json:"is_default"`
	L2_adjacency bool `json:"l2_adjacency"`
	Mtu int `json:"mtu"`
	Name string `json:"name"`
	Port_security_enabled bool `json:"port_security_enabled"`
	Project_id string `json:"project_id"`
	Provider_network_type string `json:"provider:network_type"`
	Provider_physical_network string `json:"provider:physical_network"`
	Provider_segmentation_id string `json:"provider:segmentation_id"`
	Qos_policy_id string `json:"qos_policy_id"`
	Revision_number int `json:"revision_number"`
	Router_external bool `json:"router:external"`
	Segments []struct{} `json:"segments"`
	Shared bool `json:"shared"`
	Status string `json:"status"`
	Subnets []string `json:"subnets"`
	Tags []string `json:"tags"`
	Tenant_id string `json:"tenant_id"`
	Updated_at string `json:"updated_at"`
	Vlan_transparent bool `json:"vlan_transparent"`
}
func (NetworkGet200Response) mustImplementNetworkGetResponse(){}

type NetworkGet400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkGet400Response) mustImplementNetworkGetResponse(){}

type NetworkGet404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkGet404Response) mustImplementNetworkGetResponse(){}

type NetworkGet500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkGet500Response) mustImplementNetworkGetResponse(){}

type NetworkPatchRequest struct {
	Id string
	Body NetworkPatchRequestBody
}

func (r NetworkPatchRequest)IsValid() error {
	return nil
}

type NetworkPatchRequestBody struct {
	Admin_state_up bool `json:"admin_state_up"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Dns_domain string `json:"dns_domain"`
	Ipv4_address_scope string `json:"ipv4_address_scope"`
	Ipv6_address_scope string `json:"ipv6_address_scope"`
	Is_default bool `json:"is_default"`
	L2_adjacency bool `json:"l2_adjacency"`
	Mtu int `json:"mtu"`
	Name string `json:"name"`
	Port_security_enabled bool `json:"port_security_enabled"`
	Project_id string `json:"project_id"`
	Provider_network_type string `json:"provider:network_type"`
	Provider_physical_network string `json:"provider:physical_network"`
	Provider_segmentation_id string `json:"provider:segmentation_id"`
	Qos_policy_id string `json:"qos_policy_id"`
	Revision_number int `json:"revision_number"`
	Router_external bool `json:"router:external"`
	Segments []struct{} `json:"segments"`
	Shared bool `json:"shared"`
	Status string `json:"status"`
	Subnets []string `json:"subnets"`
	Tags []string `json:"tags"`
	Tenant_id string `json:"tenant_id"`
	Updated_at string `json:"updated_at"`
	Vlan_transparent bool `json:"vlan_transparent"`
}

type NetworkPatchResponse interface {
	mustImplementNetworkPatchResponse()
}


type NetworkPatch204Response struct {
}
func (NetworkPatch204Response) mustImplementNetworkPatchResponse(){}

type NetworkPatch400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPatch400Response) mustImplementNetworkPatchResponse(){}

type NetworkPatch404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPatch404Response) mustImplementNetworkPatchResponse(){}

type NetworkPatch500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPatch500Response) mustImplementNetworkPatchResponse(){}

type NetworkPostRequest struct {
	Body NetworkPostRequestBody
}

func (r NetworkPostRequest)IsValid() error {
	return nil
}

type NetworkPostRequestBody struct {
	Admin_state_up bool `json:"admin_state_up"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Dns_domain string `json:"dns_domain"`
	Ipv4_address_scope string `json:"ipv4_address_scope"`
	Ipv6_address_scope string `json:"ipv6_address_scope"`
	Is_default bool `json:"is_default"`
	L2_adjacency bool `json:"l2_adjacency"`
	Mtu int `json:"mtu"`
	Name string `json:"name"`
	Port_security_enabled bool `json:"port_security_enabled"`
	Project_id string `json:"project_id"`
	Provider_network_type string `json:"provider:network_type"`
	Provider_physical_network string `json:"provider:physical_network"`
	Provider_segmentation_id string `json:"provider:segmentation_id"`
	Qos_policy_id string `json:"qos_policy_id"`
	Revision_number int `json:"revision_number"`
	Router_external bool `json:"router:external"`
	Segments []struct{} `json:"segments"`
	Shared bool `json:"shared"`
	Status string `json:"status"`
	Subnets []string `json:"subnets"`
	Tags []string `json:"tags"`
	Tenant_id string `json:"tenant_id"`
	Updated_at string `json:"updated_at"`
	Vlan_transparent bool `json:"vlan_transparent"`
}

type NetworkPostResponse interface {
	mustImplementNetworkPostResponse()
}


type NetworkPost204Response struct {
}
func (NetworkPost204Response) mustImplementNetworkPostResponse(){}

type NetworkPost400Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPost400Response) mustImplementNetworkPostResponse(){}

type NetworkPost404Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPost404Response) mustImplementNetworkPostResponse(){}

type NetworkPost500Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
func (NetworkPost500Response) mustImplementNetworkPostResponse(){}

