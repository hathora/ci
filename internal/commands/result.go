package commands

import "encoding/json"

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
		*DefaultResult
		Status string `json:"status"`
	}{
		DefaultResult: r,
		Status:        status,
	}

	return json.Marshal(alias)
}
