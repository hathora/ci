// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"time"
)

// ProcessWithRoomsExposedPort - Connection details for an active process.
type ProcessWithRoomsExposedPort struct {
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	Port          int           `json:"port"`
	Host          string        `json:"host"`
	Name          string        `json:"name"`
}

func (o *ProcessWithRoomsExposedPort) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *ProcessWithRoomsExposedPort) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ProcessWithRoomsExposedPort) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ProcessWithRoomsExposedPort) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// ProcessWithRooms - A process object represents a runtime instance of your game server and its metadata.
type ProcessWithRooms struct {
	// Measures network traffic leaving the process in bytes.
	EgressedBytes int `json:"egressedBytes"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	IdleSince *time.Time `json:"idleSince"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ActiveConnectionsUpdatedAt time.Time `json:"activeConnectionsUpdatedAt"`
	// Tracks the number of active connections to a process.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ActiveConnections       int       `json:"activeConnections"`
	RoomsAllocatedUpdatedAt time.Time `json:"roomsAllocatedUpdatedAt"`
	// Tracks the number of rooms that have been allocated to the process.
	RoomsAllocated int `json:"roomsAllocated"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RoomSlotsAvailableUpdatedAt time.Time `json:"roomSlotsAvailableUpdatedAt"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RoomSlotsAvailable float64 `json:"roomSlotsAvailable"`
	// Process in drain will not accept any new rooms.
	Draining bool `json:"draining"`
	// When the process has been terminated.
	TerminatedAt *time.Time `json:"terminatedAt"`
	// When the process is issued to stop. We use this to determine when we should stop billing.
	StoppingAt *time.Time `json:"stoppingAt"`
	// When the process bound to the specified port. We use this to determine when we should start billing.
	StartedAt *time.Time `json:"startedAt"`
	// When the process started being provisioned.
	StartingAt time.Time `json:"startingAt"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess        int                          `json:"roomsPerProcess"`
	AdditionalExposedPorts []ExposedPort                `json:"additionalExposedPorts"`
	ExposedPort            *ProcessWithRoomsExposedPort `json:"exposedPort"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Port float64 `json:"port"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Host   string `json:"host"`
	Region Region `json:"region"`
	// System generated unique identifier to a runtime instance of your game server.
	ProcessID string `json:"processId"`
	// System generated id for a deployment. Increments by 1.
	DeploymentID int `json:"deploymentId"`
	// System generated unique identifier for an application.
	AppID      string                   `json:"appId"`
	Rooms      []RoomWithoutAllocations `json:"rooms"`
	TotalRooms int                      `json:"totalRooms"`
}

func (p ProcessWithRooms) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProcessWithRooms) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProcessWithRooms) GetEgressedBytes() int {
	if o == nil {
		return 0
	}
	return o.EgressedBytes
}

func (o *ProcessWithRooms) GetIdleSince() *time.Time {
	if o == nil {
		return nil
	}
	return o.IdleSince
}

func (o *ProcessWithRooms) GetActiveConnectionsUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ActiveConnectionsUpdatedAt
}

func (o *ProcessWithRooms) GetActiveConnections() int {
	if o == nil {
		return 0
	}
	return o.ActiveConnections
}

func (o *ProcessWithRooms) GetRoomsAllocatedUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.RoomsAllocatedUpdatedAt
}

func (o *ProcessWithRooms) GetRoomsAllocated() int {
	if o == nil {
		return 0
	}
	return o.RoomsAllocated
}

func (o *ProcessWithRooms) GetRoomSlotsAvailableUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.RoomSlotsAvailableUpdatedAt
}

func (o *ProcessWithRooms) GetRoomSlotsAvailable() float64 {
	if o == nil {
		return 0.0
	}
	return o.RoomSlotsAvailable
}

func (o *ProcessWithRooms) GetDraining() bool {
	if o == nil {
		return false
	}
	return o.Draining
}

func (o *ProcessWithRooms) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}

func (o *ProcessWithRooms) GetStoppingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StoppingAt
}

func (o *ProcessWithRooms) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *ProcessWithRooms) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *ProcessWithRooms) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *ProcessWithRooms) GetAdditionalExposedPorts() []ExposedPort {
	if o == nil {
		return []ExposedPort{}
	}
	return o.AdditionalExposedPorts
}

func (o *ProcessWithRooms) GetExposedPort() *ProcessWithRoomsExposedPort {
	if o == nil {
		return nil
	}
	return o.ExposedPort
}

func (o *ProcessWithRooms) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *ProcessWithRooms) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ProcessWithRooms) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *ProcessWithRooms) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *ProcessWithRooms) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *ProcessWithRooms) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *ProcessWithRooms) GetRooms() []RoomWithoutAllocations {
	if o == nil {
		return []RoomWithoutAllocations{}
	}
	return o.Rooms
}

func (o *ProcessWithRooms) GetTotalRooms() int {
	if o == nil {
		return 0
	}
	return o.TotalRooms
}