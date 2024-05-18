package output

import (
	"encoding/json"
	"fmt"
)

func jsonStringify[T any](value T) (string, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to json-serialize output value %w", err)
	}

	return string(jsonBytes), nil
}
