// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"cloudapi/pkg/utils"
	"time"
)

// LobbyV3 - A lobby object allows you to store and manage metadata for your rooms.
type LobbyV3 struct {
	// User-defined identifier for a lobby.
	ShortCode string `json:"shortCode"`
	// When the lobby was created.
	CreatedAt time.Time `json:"createdAt"`
	// UserId or email address for the user that created the lobby.
	CreatedBy  string  `json:"createdBy"`
	RoomConfig *string `json:"roomConfig"`
	// Types of lobbies a player can create.
	//
	// `private`: the player who created the room must share the roomId with their friends
	//
	// `public`: visible in the public lobby list, anyone can join
	//
	// `local`: for testing with a server running locally
	Visibility LobbyVisibility `json:"visibility"`
	Region     Region          `json:"region"`
	// Unique identifier to a game session or match. Use the default system generated ID or overwrite it with your own.
	// Note: error will be returned if `roomId` is not globally unique.
	RoomID string `json:"roomId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (l LobbyV3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LobbyV3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LobbyV3) GetShortCode() string {
	if o == nil {
		return ""
	}
	return o.ShortCode
}

func (o *LobbyV3) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *LobbyV3) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *LobbyV3) GetRoomConfig() *string {
	if o == nil {
		return nil
	}
	return o.RoomConfig
}

func (o *LobbyV3) GetVisibility() LobbyVisibility {
	if o == nil {
		return LobbyVisibility("")
	}
	return o.Visibility
}

func (o *LobbyV3) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *LobbyV3) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *LobbyV3) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
