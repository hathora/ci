// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/components"
)

type UpdateFleetRegionGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *UpdateFleetRegionGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type UpdateFleetRegionRequest struct {
	FleetID           string                       `pathParam:"style=simple,explode=false,name=fleetId"`
	Region            components.Region            `pathParam:"style=simple,explode=false,name=region"`
	OrgID             *string                      `queryParam:"style=form,explode=true,name=orgId"`
	FleetRegionConfig components.FleetRegionConfig `request:"mediaType=application/json"`
}

func (o *UpdateFleetRegionRequest) GetFleetID() string {
	if o == nil {
		return ""
	}
	return o.FleetID
}

func (o *UpdateFleetRegionRequest) GetRegion() components.Region {
	if o == nil {
		return components.Region("")
	}
	return o.Region
}

func (o *UpdateFleetRegionRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

func (o *UpdateFleetRegionRequest) GetFleetRegionConfig() components.FleetRegionConfig {
	if o == nil {
		return components.FleetRegionConfig{}
	}
	return o.FleetRegionConfig
}
