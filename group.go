package agora_chat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

// GroupsClient is a specialized client used to interact with the Groups endpoints.
type GroupsClient struct {
	client *Client
}

type Group struct {
	Groupname string   `json:"groupname"`
	Groupid   string   `json:"groupid"`
	Desc      string   `json:"desc"`
	Public    bool     `json:"public"`
	Maxusers  int      `json:"maxusers"`
	Owner     string   `json:"owner"`
	Members   []string `json:"members"`
}

type GroupResponse struct {
	Group *Group `json:"data"`
	Response
}
type groupForJSON Group

// UnmarshalJSON implements json.Unmarshaler.
func (g *Group) UnmarshalJSON(data []byte) error {
	var g2 groupForJSON
	if err := json.Unmarshal(data, &g2); err != nil {
		return err
	}
	*g = Group(g2)
	return nil
}

// CreateGroups creates the given group.
func (g *GroupsClient) CreateGroups(ctx context.Context, group *Group) (*GroupResponse, error) {
	req := group
	var resp GroupResponse
	err := g.client.makeRequest(ctx, http.MethodPost, "chatgroups", nil, req, &resp)
	return &resp, err
}

// DeleteGroup deletes the group with the given groupID
func (g *GroupsClient) DeleteGroup(ctx context.Context, groupID string) (*GroupResponse, error) {
	if groupID == "" {
		return nil, errors.New("group ID is empty")
	}
	values := url.Values{}
	p := path.Join("chatgroups", url.PathEscape(groupID))
	var resp GroupResponse
	err := g.client.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}
