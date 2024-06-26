// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InviteStatusRejectedType string

const (
	InviteStatusRejectedTypeRejected InviteStatusRejectedType = "rejected"
)

func (e InviteStatusRejectedType) ToPointer() *InviteStatusRejectedType {
	return &e
}
func (e *InviteStatusRejectedType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rejected":
		*e = InviteStatusRejectedType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteStatusRejectedType: %v", v)
	}
}

type InviteStatusRejected struct {
	// System generated unique identifier for a user. Not guaranteed to have a specific format.
	UserID string                   `json:"userId"`
	Type   InviteStatusRejectedType `json:"type"`
}

func (o *InviteStatusRejected) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *InviteStatusRejected) GetType() InviteStatusRejectedType {
	if o == nil {
		return InviteStatusRejectedType("")
	}
	return o.Type
}
