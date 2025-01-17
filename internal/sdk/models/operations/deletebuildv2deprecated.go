// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteBuildV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteBuildV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DeleteBuildV2DeprecatedRequest struct {
	AppID   *string `pathParam:"style=simple,explode=false,name=appId"`
	BuildID int     `pathParam:"style=simple,explode=false,name=buildId"`
}

func (o *DeleteBuildV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *DeleteBuildV2DeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}
