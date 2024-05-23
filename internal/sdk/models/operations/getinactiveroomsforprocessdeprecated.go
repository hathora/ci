// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetInactiveRoomsForProcessDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetInactiveRoomsForProcessDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetInactiveRoomsForProcessDeprecatedRequest struct {
	ProcessID string  `pathParam:"style=simple,explode=false,name=processId"`
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetInactiveRoomsForProcessDeprecatedRequest) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *GetInactiveRoomsForProcessDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetInactiveRoomsForProcessDeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	RoomWithoutAllocations []shared.RoomWithoutAllocations
}

func (o *GetInactiveRoomsForProcessDeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInactiveRoomsForProcessDeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInactiveRoomsForProcessDeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetInactiveRoomsForProcessDeprecatedResponse) GetRoomWithoutAllocations() []shared.RoomWithoutAllocations {
	if o == nil {
		return nil
	}
	return o.RoomWithoutAllocations
}