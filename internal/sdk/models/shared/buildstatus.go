// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BuildStatus string

const (
	BuildStatusCreated   BuildStatus = "created"
	BuildStatusRunning   BuildStatus = "running"
	BuildStatusSucceeded BuildStatus = "succeeded"
	BuildStatusFailed    BuildStatus = "failed"
)

func (e BuildStatus) ToPointer() *BuildStatus {
	return &e
}
