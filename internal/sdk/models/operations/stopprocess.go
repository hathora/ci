// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type StopProcessGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *StopProcessGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type StopProcessRequest struct {
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
	ProcessID string  `pathParam:"style=simple,explode=false,name=processId"`
}

func (o *StopProcessRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *StopProcessRequest) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}
