// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetOrgTokensRequest struct {
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

func (o *GetOrgTokensRequest) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

type GetOrgTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	ListOrgTokens *shared.ListOrgTokens
}

func (o *GetOrgTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOrgTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOrgTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOrgTokensResponse) GetListOrgTokens() *shared.ListOrgTokens {
	if o == nil {
		return nil
	}
	return o.ListOrgTokens
}
