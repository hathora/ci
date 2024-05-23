// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetLobbyInfoByRoomIDGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByRoomIDGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByRoomIDRequest struct {
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByRoomIDRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *GetLobbyInfoByRoomIDRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByRoomIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	LobbyV3 *shared.LobbyV3
}

func (o *GetLobbyInfoByRoomIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLobbyInfoByRoomIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLobbyInfoByRoomIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLobbyInfoByRoomIDResponse) GetLobbyV3() *shared.LobbyV3 {
	if o == nil {
		return nil
	}
	return o.LobbyV3
}