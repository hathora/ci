// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Region string

const (
	RegionSeattle      Region = "Seattle"
	RegionLosAngeles   Region = "Los_Angeles"
	RegionWashingtonDc Region = "Washington_DC"
	RegionChicago      Region = "Chicago"
	RegionLondon       Region = "London"
	RegionFrankfurt    Region = "Frankfurt"
	RegionMumbai       Region = "Mumbai"
	RegionSingapore    Region = "Singapore"
	RegionTokyo        Region = "Tokyo"
	RegionSydney       Region = "Sydney"
	RegionSaoPaulo     Region = "Sao_Paulo"
	RegionDallas       Region = "Dallas"
)

func (e Region) ToPointer() *Region {
	return &e
}
func (e *Region) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Seattle":
		fallthrough
	case "Los_Angeles":
		fallthrough
	case "Washington_DC":
		fallthrough
	case "Chicago":
		fallthrough
	case "London":
		fallthrough
	case "Frankfurt":
		fallthrough
	case "Mumbai":
		fallthrough
	case "Singapore":
		fallthrough
	case "Tokyo":
		fallthrough
	case "Sydney":
		fallthrough
	case "Sao_Paulo":
		fallthrough
	case "Dallas":
		*e = Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Region: %v", v)
	}
}