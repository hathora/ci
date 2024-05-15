// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type InviteUserRequest struct {
	OrgID            string                  `pathParam:"style=simple,explode=false,name=orgId"`
	CreateUserInvite shared.CreateUserInvite `request:"mediaType=application/json"`
}

func (o *InviteUserRequest) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *InviteUserRequest) GetCreateUserInvite() shared.CreateUserInvite {
	if o == nil {
		return shared.CreateUserInvite{}
	}
	return o.CreateUserInvite
}

type InviteUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	OrgPermission *shared.OrgPermission
}

func (o *InviteUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InviteUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InviteUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InviteUserResponse) GetOrgPermission() *shared.OrgPermission {
	if o == nil {
		return nil
	}
	return o.OrgPermission
}
