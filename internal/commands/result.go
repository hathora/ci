package commands

import (
	"encoding/json"
)

type DefaultResult struct {
	Success bool   `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *DefaultResult) MarshalJSON() ([]byte, error) {
	status := "success"
	if !r.Success {
		status = "failure"
	}

	alias := struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Code    int    `json:"code"`
	}{
		Message: r.Message,
		Status:  status,
		Code:    r.Code,
	}

	return json.Marshal(alias)
}
