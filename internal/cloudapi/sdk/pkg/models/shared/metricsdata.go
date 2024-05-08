// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MetricsData - Construct a type with a set of properties K of type T
type MetricsData struct {
	CPU               []MetricValue `json:"cpu,omitempty"`
	Memory            []MetricValue `json:"memory,omitempty"`
	RateEgress        []MetricValue `json:"rate_egress,omitempty"`
	TotalEgress       []MetricValue `json:"total_egress,omitempty"`
	ActiveConnections []MetricValue `json:"active_connections,omitempty"`
}

func (o *MetricsData) GetCPU() []MetricValue {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *MetricsData) GetMemory() []MetricValue {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *MetricsData) GetRateEgress() []MetricValue {
	if o == nil {
		return nil
	}
	return o.RateEgress
}

func (o *MetricsData) GetTotalEgress() []MetricValue {
	if o == nil {
		return nil
	}
	return o.TotalEgress
}

func (o *MetricsData) GetActiveConnections() []MetricValue {
	if o == nil {
		return nil
	}
	return o.ActiveConnections
}
