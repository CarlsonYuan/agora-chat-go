package agora_chat

import (
	"encoding/json"
)

type User struct {
	Created   int64  `json:"created"`
	Nickname  string `json:"nickname"`
	Modified  int64  `json:"modified"`
	Type      string `json:"type"`
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	Activated bool   `json:"activated"`
}

type userForJSON User

// UnmarshalJSON implements json.Unmarshaler.
func (u *User) UnmarshalJSON(data []byte) error {
	var u2 userForJSON
	if err := json.Unmarshal(data, &u2); err != nil {
		return err
	}
	*u = User(u2)
	return nil
}
