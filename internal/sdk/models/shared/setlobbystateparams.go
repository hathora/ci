// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SetLobbyStateParamsState - JSON blob to store metadata for a room. Must be smaller than 1MB.
type SetLobbyStateParamsState struct {
}

type SetLobbyStateParams struct {
	// JSON blob to store metadata for a room. Must be smaller than 1MB.
	State SetLobbyStateParamsState `json:"state"`
}

func (o *SetLobbyStateParams) GetState() SetLobbyStateParamsState {
	if o == nil {
		return SetLobbyStateParamsState{}
	}
	return o.State
}
