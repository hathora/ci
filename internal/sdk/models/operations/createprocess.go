// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreateProcessGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateProcessGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateProcessRequest struct {
	AppID  *string       `pathParam:"style=simple,explode=false,name=appId"`
	Region shared.Region `pathParam:"style=simple,explode=false,name=region"`
}

func (o *CreateProcessRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateProcessRequest) GetRegion() shared.Region {
	if o == nil {
		return shared.Region("")
	}
	return o.Region
}

type CreateProcessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	ProcessV3   *shared.ProcessV3
}

func (o *CreateProcessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateProcessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateProcessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateProcessResponse) GetProcessV3() *shared.ProcessV3 {
	if o == nil {
		return nil
	}
	return o.ProcessV3
}
