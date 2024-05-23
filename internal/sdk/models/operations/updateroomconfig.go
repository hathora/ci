// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type UpdateRoomConfigGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *UpdateRoomConfigGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type UpdateRoomConfigRequest struct {
	RoomID                 string                        `pathParam:"style=simple,explode=false,name=roomId"`
	UpdateRoomConfigParams shared.UpdateRoomConfigParams `request:"mediaType=application/json"`
	AppID                  *string                       `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *UpdateRoomConfigRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *UpdateRoomConfigRequest) GetUpdateRoomConfigParams() shared.UpdateRoomConfigParams {
	if o == nil {
		return shared.UpdateRoomConfigParams{}
	}
	return o.UpdateRoomConfigParams
}

func (o *UpdateRoomConfigRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type UpdateRoomConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateRoomConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRoomConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRoomConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}