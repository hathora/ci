// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreateWithMultipartUploadsV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateWithMultipartUploadsV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateWithMultipartUploadsV2DeprecatedRequest struct {
	AppID                      *string                           `pathParam:"style=simple,explode=false,name=appId"`
	CreateMultipartBuildParams shared.CreateMultipartBuildParams `request:"mediaType=application/json"`
}

func (o *CreateWithMultipartUploadsV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateWithMultipartUploadsV2DeprecatedRequest) GetCreateMultipartBuildParams() shared.CreateMultipartBuildParams {
	if o == nil {
		return shared.CreateMultipartBuildParams{}
	}
	return o.CreateMultipartBuildParams
}

type CreateWithMultipartUploadsV2DeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse            *http.Response
	BuildWithMultipartUrls *shared.BuildWithMultipartUrls
}

func (o *CreateWithMultipartUploadsV2DeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWithMultipartUploadsV2DeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWithMultipartUploadsV2DeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWithMultipartUploadsV2DeprecatedResponse) GetBuildWithMultipartUrls() *shared.BuildWithMultipartUrls {
	if o == nil {
		return nil
	}
	return o.BuildWithMultipartUrls
}
