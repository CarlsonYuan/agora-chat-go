package agora_chat

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"path"
)

type QueryOption struct {
	Limit  int    `json:"limit,omitempty"`  // pagination option: limit number of results
	Cursor string `json:"cursor,omitempty"` // pagination option: offset to return items from
}

type UsersResponse struct {
	Users []*User `json:"entities"`
	Response
}

// QueryUser returns specified user that match userID.
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

// QueryUsers returns list of users that match QueryOption.
func (c *Client) QueryUsers(ctx context.Context, q *QueryOption) (*UsersResponse, error) {

	values := url.Values{}
	values.Add("limit", string(q.Limit))
	values.Add("cursor", q.Cursor)

	var resp UsersResponse
	err := c.makeRequest(ctx, http.MethodGet, "users", values, nil, &resp)
	return &resp, err
}
