// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteAppGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteAppGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DeleteAppRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteAppRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}
