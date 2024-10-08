// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreatePrivateLobbyDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreatePrivateLobbyDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreatePrivateLobbyDeprecatedSecurity struct {
	PlayerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *CreatePrivateLobbyDeprecatedSecurity) GetPlayerAuth() string {
	if o == nil {
		return ""
	}
	return o.PlayerAuth
}

type CreatePrivateLobbyDeprecatedRequest struct {
	AppID  *string        `pathParam:"style=simple,explode=false,name=appId"`
	Region *shared.Region `queryParam:"style=form,explode=true,name=region"`
	Local  *bool          `default:"false" queryParam:"style=form,explode=true,name=local"`
}

func (c CreatePrivateLobbyDeprecatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePrivateLobbyDeprecatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetRegion() *shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetLocal() *bool {
	if o == nil {
		return nil
	}
	return o.Local
}

type CreatePrivateLobbyDeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	RoomID *string
}

func (o *CreatePrivateLobbyDeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePrivateLobbyDeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePrivateLobbyDeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePrivateLobbyDeprecatedResponse) GetRoomID() *string {
	if o == nil {
		return nil
	}
	return o.RoomID
}
