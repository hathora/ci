// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type OrgsPage struct {
	Orgs []Organization `json:"orgs"`
}

func (o *OrgsPage) GetOrgs() []Organization {
	if o == nil {
		return []Organization{}
	}
	return o.Orgs
}
