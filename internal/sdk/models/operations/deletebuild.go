// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteBuildGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteBuildGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DeleteBuildRequest struct {
	BuildID int     `pathParam:"style=simple,explode=false,name=buildId"`
	AppID   *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteBuildRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *DeleteBuildRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DeleteBuildResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteBuildResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteBuildResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteBuildResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}