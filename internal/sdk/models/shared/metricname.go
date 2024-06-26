// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// MetricName - Available metrics to query over time.
type MetricName string

const (
	MetricNameCPU               MetricName = "cpu"
	MetricNameMemory            MetricName = "memory"
	MetricNameRateEgress        MetricName = "rate_egress"
	MetricNameTotalEgress       MetricName = "total_egress"
	MetricNameActiveConnections MetricName = "active_connections"
)

func (e MetricName) ToPointer() *MetricName {
	return &e
}
func (e *MetricName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cpu":
		fallthrough
	case "memory":
		fallthrough
	case "rate_egress":
		fallthrough
	case "total_egress":
		fallthrough
	case "active_connections":
		*e = MetricName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MetricName: %v", v)
	}
}
