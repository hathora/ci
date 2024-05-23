// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type RescindInviteRequest struct {
	OrgID             string                   `pathParam:"style=simple,explode=false,name=orgId"`
	RescindUserInvite shared.RescindUserInvite `request:"mediaType=application/json"`
}

func (o *RescindInviteRequest) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *RescindInviteRequest) GetRescindUserInvite() shared.RescindUserInvite {
	if o == nil {
		return shared.RescindUserInvite{}
	}
	return o.RescindUserInvite
}

type RescindInviteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RescindInviteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RescindInviteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RescindInviteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}