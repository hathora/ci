// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetLatestProcessesV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLatestProcessesV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLatestProcessesV2DeprecatedRequest struct {
	AppID  *string                `pathParam:"style=simple,explode=false,name=appId"`
	Status []shared.ProcessStatus `queryParam:"style=form,explode=true,name=status"`
	Region []shared.Region        `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetStatus() []shared.ProcessStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetRegion() []shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

type GetLatestProcessesV2DeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	ProcessV2s []shared.ProcessV2
}

func (o *GetLatestProcessesV2DeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLatestProcessesV2DeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLatestProcessesV2DeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLatestProcessesV2DeprecatedResponse) GetProcessV2s() []shared.ProcessV2 {
	if o == nil {
		return nil
	}
	return o.ProcessV2s
}