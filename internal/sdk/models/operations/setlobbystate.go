// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type SetLobbyStateGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *SetLobbyStateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type SetLobbyStateRequest struct {
	AppID               *string                    `pathParam:"style=simple,explode=false,name=appId"`
	RoomID              string                     `pathParam:"style=simple,explode=false,name=roomId"`
	SetLobbyStateParams shared.SetLobbyStateParams `request:"mediaType=application/json"`
}

func (o *SetLobbyStateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *SetLobbyStateRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *SetLobbyStateRequest) GetSetLobbyStateParams() shared.SetLobbyStateParams {
	if o == nil {
		return shared.SetLobbyStateParams{}
	}
	return o.SetLobbyStateParams
}

type SetLobbyStateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Lobby *shared.Lobby
}

func (o *SetLobbyStateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetLobbyStateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetLobbyStateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetLobbyStateResponse) GetLobby() *shared.Lobby {
	if o == nil {
		return nil
	}
	return o.Lobby
}
