// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrgTokenStatus string

const (
	OrgTokenStatusActive  OrgTokenStatus = "active"
	OrgTokenStatusRevoked OrgTokenStatus = "revoked"
)

func (e OrgTokenStatus) ToPointer() *OrgTokenStatus {
	return &e
}
