// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetLobbyInfoByShortCodeGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByShortCodeGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByShortCodeRequest struct {
	ShortCode string  `pathParam:"style=simple,explode=false,name=shortCode"`
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByShortCodeRequest) GetShortCode() string {
	if o == nil {
		return ""
	}
	return o.ShortCode
}

func (o *GetLobbyInfoByShortCodeRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByShortCodeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	LobbyV3 *shared.LobbyV3
}

func (o *GetLobbyInfoByShortCodeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLobbyInfoByShortCodeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLobbyInfoByShortCodeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLobbyInfoByShortCodeResponse) GetLobbyV3() *shared.LobbyV3 {
	if o == nil {
		return nil
	}
	return o.LobbyV3
}