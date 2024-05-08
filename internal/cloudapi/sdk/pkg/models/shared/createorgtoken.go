// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateOrgToken struct {
	// Readable name for a token. Must be unique within an organization.
	Name string `json:"name"`
}

func (o *CreateOrgToken) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
