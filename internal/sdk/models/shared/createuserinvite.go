// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateUserInvite struct {
	// A user's email.
	UserEmail string `json:"userEmail"`
}

func (o *CreateUserInvite) GetUserEmail() string {
	if o == nil {
		return ""
	}
	return o.UserEmail
}
