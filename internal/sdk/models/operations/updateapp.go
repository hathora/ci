// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type UpdateAppGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *UpdateAppGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type UpdateAppRequest struct {
	AppConfig shared.AppConfig `request:"mediaType=application/json"`
	AppID     *string          `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *UpdateAppRequest) GetAppConfig() shared.AppConfig {
	if o == nil {
		return shared.AppConfig{}
	}
	return o.AppConfig
}

func (o *UpdateAppRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type UpdateAppResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Application *shared.Application
}

func (o *UpdateAppResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAppResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAppResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAppResponse) GetApplication() *shared.Application {
	if o == nil {
		return nil
	}
	return o.Application
}