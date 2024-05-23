// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NicknameObject struct {
	// An alias to represent a player.
	Nickname string `json:"nickname"`
}

func (o *NicknameObject) GetNickname() string {
	if o == nil {
		return ""
	}
	return o.Nickname
}