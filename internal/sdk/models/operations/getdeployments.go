// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetDeploymentsGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentsRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	DeploymentV2s []shared.DeploymentV2
}

func (o *GetDeploymentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeploymentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeploymentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDeploymentsResponse) GetDeploymentV2s() []shared.DeploymentV2 {
	if o == nil {
		return nil
	}
	return o.DeploymentV2s
}
