package agora_chat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

// UsersClient is a specialized client used to interact with the Users endpoints.
type UsersClient struct {
	client *Client
}

type User struct {
	Created   int64  `json:"created"`
	Nickname  string `json:"nickname"`
	Modified  int64  `json:"modified"`
	Type      string `json:"type"`
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	Activated bool   `json:"activated"`
	Password  string `json:"password"`
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

// CreateUsers creates the given users.
func (c *UsersClient) CreateUsers(ctx context.Context, users ...*User) (*UsersResponse, error) {
	if len(users) == 0 {
		return nil, errors.New("users are not set")
	}

	req := users

	var resp UsersResponse
	err := c.client.makeRequest(ctx, http.MethodPost, "users", nil, req, &resp)
	return &resp, err
}

// DeleteUser deletes the user with the given userID(username).
func (c *UsersClient) DeleteUser(ctx context.Context, userID string) (*UsersResponse, error) {
	if userID == "" {
		return nil, errors.New("user ID is empty")
	}
	values := url.Values{}
	p := path.Join("users", url.PathEscape(userID))
	var resp UsersResponse
	err := c.client.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}
