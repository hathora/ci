// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hathora/ci/internal/sdk/internal/utils"
)

type DeploymentConfigEnv struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (o *DeploymentConfigEnv) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *DeploymentConfigEnv) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// DeploymentConfig - User specified deployment configuration for your application at runtime.
type DeploymentConfig struct {
	// Option to shut down processes that have had no new connections or rooms
	// for five minutes.
	IdleTimeoutEnabled *bool `default:"true" json:"idleTimeoutEnabled"`
	// The environment variable that our process will have access to at runtime.
	Env []DeploymentConfigEnv `json:"env"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess int `json:"roomsPerProcess"`
	// A plan defines how much CPU and memory is required to run an instance of your game server.
	//
	// `tiny`: shared core, 1gb memory
	//
	// `small`: 1 core, 2gb memory
	//
	// `medium`: 2 core, 4gb memory
	//
	// `large`: 4 core, 8gb memory
	PlanName PlanName `json:"planName"`
	// Additional ports your server listens on.
	AdditionalContainerPorts []ContainerPort `json:"additionalContainerPorts,omitempty"`
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	// Default port the server listens on.
	ContainerPort int `json:"containerPort"`
}

func (d DeploymentConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeploymentConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeploymentConfig) GetIdleTimeoutEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IdleTimeoutEnabled
}

func (o *DeploymentConfig) GetEnv() []DeploymentConfigEnv {
	if o == nil {
		return []DeploymentConfigEnv{}
	}
	return o.Env
}

func (o *DeploymentConfig) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *DeploymentConfig) GetPlanName() PlanName {
	if o == nil {
		return PlanName("")
	}
	return o.PlanName
}

func (o *DeploymentConfig) GetAdditionalContainerPorts() []ContainerPort {
	if o == nil {
		return nil
	}
	return o.AdditionalContainerPorts
}

func (o *DeploymentConfig) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *DeploymentConfig) GetContainerPort() int {
	if o == nil {
		return 0
	}
	return o.ContainerPort
}
