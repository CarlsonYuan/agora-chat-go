package agora_chat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type User struct {
	ID string `json:"username"`

	Created   int64  `json:"created,omitempty"`
	Modified  int64  `json:"modified,omitempty"`
	Type      string `json:"type,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
	Activated bool   `json:"activated,omitempty"`
	// UserAttributes
	Nickname  string `json:"nickname,omitempty"`
	AvatarUrl string `json:"avatarurl,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Mail      string `json:"mail,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Sign      string `json:"sign,omitempty"`
	Birth     string `json:"birth,omitempty"`
	Ext       string `json:"ext,omitempty"`
	// PushInfos
	PushInfos []*PushInfo `json:"pushInfo,omitempty"`
	Password  string      `json:"password"`
}

// NewUser create new instance of User
func NewUser(uID string) User {
	u := User{}
	u.ID = uID
	u.Password = uID
	return u
}

type PushInfo struct {
	DeviceID     string `json:"device_Id"`
	DeviceToken  string `json:"device_token"`
	NotifierName string `json:"notifier_name"`
}

type CreateUserError struct {
	ID     string `json:"username"`
	Reason string `json:"registerUserFailReason"`
}
type UsersResponse struct {
	Users []*User           `json:"entities"`
	Data  []CreateUserError `json:"data"`
	Response
}

// CreateUsers creates the given users.
func (c *Client) CreateUsers(ctx context.Context, users ...*User) (*UsersResponse, error) {
	if len(users) == 0 {
		return nil, errors.New("users are not set")
	}
	req := users
	var resp UsersResponse
	err := c.makeRequest(ctx, http.MethodPost, "users", nil, req, &resp)
	return &resp, err
}

// QueryUser returns the specified user that match userID.
func (c *Client) QueryUser(ctx context.Context, userID string) (*UsersResponse, error) {
	if userID == "" {
		return nil, errors.New("user ID is empty")
	}
	values := url.Values{}
	p := path.Join("users", url.PathEscape(userID))
	var resp UsersResponse
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// DeleteUser deletes the user with the given userID(username).
func (c *Client) DeleteUser(ctx context.Context, userID string) (*UsersResponse, error) {
	if userID == "" {
		return nil, errors.New("user ID is empty")
	}
	values := url.Values{}
	p := path.Join("users", url.PathEscape(userID))
	var resp UsersResponse
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

type PaginationParamsRequest struct {
	Limit  int    `json:"limit,omitempty"`  // pagination option: limit number of results
	Cursor string `json:"cursor,omitempty"` // pagination option: offset to return items from
}

type QueryUsersRequest struct {
	Pagination *PaginationParamsRequest
}

// QueryUsers returns list of users that match QueryUsersRequest.
func (c *Client) QueryUsers(ctx context.Context, q *QueryUsersRequest) (*UsersResponse, error) {

	values := url.Values{}
	values.Add("limit", fmt.Sprint(q.Pagination.Limit))
	values.Add("cursor", q.Pagination.Cursor)

	var resp UsersResponse
	err := c.makeRequest(ctx, http.MethodGet, "users", values, nil, &resp)
	return &resp, err
}
