// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RoomConnectionData - Connection information for the default and additional ports.
type RoomConnectionData struct {
	AdditionalExposedPorts []ExposedPort `json:"additionalExposedPorts"`
	// Connection details for an active process.
	ExposedPort *ExposedPort    `json:"exposedPort,omitempty"`
	Status      RoomReadyStatus `json:"status"`
	// Unique identifier to a game session or match. Use the default system generated ID or overwrite it with your own.
	// Note: error will be returned if `roomId` is not globally unique.
	RoomID string `json:"roomId"`
	// System generated unique identifier to a runtime instance of your game server.
	ProcessID string `json:"processId"`
}

func (o *RoomConnectionData) GetAdditionalExposedPorts() []ExposedPort {
	if o == nil {
		return []ExposedPort{}
	}
	return o.AdditionalExposedPorts
}

func (o *RoomConnectionData) GetExposedPort() *ExposedPort {
	if o == nil {
		return nil
	}
	return o.ExposedPort
}

func (o *RoomConnectionData) GetStatus() RoomReadyStatus {
	if o == nil {
		return RoomReadyStatus("")
	}
	return o.Status
}

func (o *RoomConnectionData) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *RoomConnectionData) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}