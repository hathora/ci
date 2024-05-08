// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cloudapi/pkg/models/shared"
	"net/http"
)

type GetAppInfoGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetAppInfoGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetAppInfoRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetAppInfoRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetAppInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Application *shared.Application
}

func (o *GetAppInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAppInfoResponse) GetApplication() *shared.Application {
	if o == nil {
		return nil
	}
	return o.Application
}
