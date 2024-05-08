// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"cloudapi/pkg/utils"
	"time"
)

type ApplicationWithLatestDeploymentAndBuildEnv struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (o *ApplicationWithLatestDeploymentAndBuildEnv) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *ApplicationWithLatestDeploymentAndBuildEnv) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ApplicationWithLatestDeploymentAndBuildDeployment struct {
	// Option to shut down processes that have had no new connections or rooms
	// for five minutes.
	IdleTimeoutEnabled bool `json:"idleTimeoutEnabled"`
	// The environment variable that our process will have access to at runtime.
	Env []ApplicationWithLatestDeploymentAndBuildEnv `json:"env"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess int `json:"roomsPerProcess"`
	// Additional ports your server listens on.
	AdditionalContainerPorts []ContainerPort `json:"additionalContainerPorts"`
	// A container port object represents the transport configruations for how your server will listen.
	DefaultContainerPort ContainerPort `json:"defaultContainerPort"`
	// When the deployment was created.
	CreatedAt time.Time `json:"createdAt"`
	// UserId or email address for the user that created the deployment.
	CreatedBy string `json:"createdBy"`
	// The amount of memory allocated to your process.
	RequestedMemoryMB float64 `json:"requestedMemoryMB"`
	// The number of cores allocated to your process.
	RequestedCPU float64 `json:"requestedCPU"`
	// System generated id for a deployment. Increments by 1.
	DeploymentID int `json:"deploymentId"`
	// System generated id for a build. Increments by 1.
	BuildID int `json:"buildId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
	// A build represents a game server artifact and its associated metadata.
	Build Build `json:"build"`
}

func (a ApplicationWithLatestDeploymentAndBuildDeployment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationWithLatestDeploymentAndBuildDeployment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetIdleTimeoutEnabled() bool {
	if o == nil {
		return false
	}
	return o.IdleTimeoutEnabled
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetEnv() []ApplicationWithLatestDeploymentAndBuildEnv {
	if o == nil {
		return []ApplicationWithLatestDeploymentAndBuildEnv{}
	}
	return o.Env
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetAdditionalContainerPorts() []ContainerPort {
	if o == nil {
		return []ContainerPort{}
	}
	return o.AdditionalContainerPorts
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetDefaultContainerPort() ContainerPort {
	if o == nil {
		return ContainerPort{}
	}
	return o.DefaultContainerPort
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetRequestedMemoryMB() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedMemoryMB
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetRequestedCPU() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedCPU
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *ApplicationWithLatestDeploymentAndBuildDeployment) GetBuild() Build {
	if o == nil {
		return Build{}
	}
	return o.Build
}

// ApplicationWithLatestDeploymentAndBuild - An application object is the top level namespace for the game server.
type ApplicationWithLatestDeploymentAndBuild struct {
	// UserId or email address for the user that deleted the application.
	DeletedBy *string `json:"deletedBy"`
	// When the application was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// When the application was created.
	CreatedAt time.Time `json:"createdAt"`
	// UserId or email address for the user that created the application.
	CreatedBy string `json:"createdBy"`
	// System generated unique identifier for an organization. Not guaranteed to have a specific format.
	OrgID string `json:"orgId"`
	// Configure [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).
	AuthConfiguration AuthConfiguration `json:"authConfiguration"`
	// Secret that is used for identity and access management.
	AppSecret string `json:"appSecret"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
	// Readable name for an application. Must be unique within an organization.
	AppName    string                                             `json:"appName"`
	Deployment *ApplicationWithLatestDeploymentAndBuildDeployment `json:"deployment,omitempty"`
}

func (a ApplicationWithLatestDeploymentAndBuild) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationWithLatestDeploymentAndBuild) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetDeletedBy() *string {
	if o == nil {
		return nil
	}
	return o.DeletedBy
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetAuthConfiguration() AuthConfiguration {
	if o == nil {
		return AuthConfiguration{}
	}
	return o.AuthConfiguration
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetAppSecret() string {
	if o == nil {
		return ""
	}
	return o.AppSecret
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}

func (o *ApplicationWithLatestDeploymentAndBuild) GetDeployment() *ApplicationWithLatestDeploymentAndBuildDeployment {
	if o == nil {
		return nil
	}
	return o.Deployment
}
