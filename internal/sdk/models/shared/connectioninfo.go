// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/hathora/ci/internal/sdk/internal/utils"
)

type ConnectionInfoType string

const (
	ConnectionInfoTypeStartingConnectionInfo ConnectionInfoType = "StartingConnectionInfo"
	ConnectionInfoTypeActiveConnectionInfo   ConnectionInfoType = "ActiveConnectionInfo"
)

// ConnectionInfo - Connection information to the default port.
type ConnectionInfo struct {
	StartingConnectionInfo *StartingConnectionInfo
	ActiveConnectionInfo   *ActiveConnectionInfo

	Type ConnectionInfoType
}

func CreateConnectionInfoStartingConnectionInfo(startingConnectionInfo StartingConnectionInfo) ConnectionInfo {
	typ := ConnectionInfoTypeStartingConnectionInfo

	return ConnectionInfo{
		StartingConnectionInfo: &startingConnectionInfo,
		Type:                   typ,
	}
}

func CreateConnectionInfoActiveConnectionInfo(activeConnectionInfo ActiveConnectionInfo) ConnectionInfo {
	typ := ConnectionInfoTypeActiveConnectionInfo

	return ConnectionInfo{
		ActiveConnectionInfo: &activeConnectionInfo,
		Type:                 typ,
	}
}

func (u *ConnectionInfo) UnmarshalJSON(data []byte) error {

	startingConnectionInfo := StartingConnectionInfo{}
	if err := utils.UnmarshalJSON(data, &startingConnectionInfo, "", true, true); err == nil {
		u.StartingConnectionInfo = &startingConnectionInfo
		u.Type = ConnectionInfoTypeStartingConnectionInfo
		return nil
	}

	activeConnectionInfo := ActiveConnectionInfo{}
	if err := utils.UnmarshalJSON(data, &activeConnectionInfo, "", true, true); err == nil {
		u.ActiveConnectionInfo = &activeConnectionInfo
		u.Type = ConnectionInfoTypeActiveConnectionInfo
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectionInfo) MarshalJSON() ([]byte, error) {
	if u.StartingConnectionInfo != nil {
		return utils.MarshalJSON(u.StartingConnectionInfo, "", true)
	}

	if u.ActiveConnectionInfo != nil {
		return utils.MarshalJSON(u.ActiveConnectionInfo, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
