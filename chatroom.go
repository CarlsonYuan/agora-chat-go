package agora_chat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Chatroom struct {
	Members           []string `json:"members,omitempty"`
	Owner             string   `json:"owner,omitempty"`
	ID                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Description       string   `json:"description,omitempty"`
	Membersonly       bool     `json:"membersonly,omitempty"`
	Allowinvites      bool     `json:"allowinvites,omitempty"`
	InviteNeedConfirm bool     `json:"invite_need_confirm,omitempty"`
	Maxusers          int      `json:"maxusers,omitempty"`
	Created           int64    `json:"created,omitempty"`
	Custom            string   `json:"custom,omitempty"`
	Mute              bool     `json:"mute,omitempty"`
	Scale             string   `json:"scale,omitempty"`
	AffiliationsCount int      `json:"affiliations_count,omitempty"`
	Disabled          bool     `json:"disabled,omitempty"`
	Affiliations      []struct {
		Member string `json:"member,omitempty"`
		Owner  string `json:"owner,omitempty"`
	} `json:"affiliations,omitempty"`
	Public bool `json:"public,omitempty"`
}
type RestHandleResult struct {
	ID     string `json:"id,omitempty"`
	Result bool   `json:"result,omitempty"`
}

type ChatroomResponse struct {
	Data RestHandleResult `json:"data"`
	Response
}

// CreateChatroom creates a chat room.
func (c *Client) CreateChatroom(ctx context.Context, chatroom *Chatroom) (*ChatroomResponse, error) {
	req := chatroom
	var resp ChatroomResponse
	err := c.makeRequest(ctx, http.MethodPost, "chatrooms", nil, req, &resp)
	return &resp, err
}

type QueryChatroomsRequest struct {
	Pagination *PaginationParamsRequest
}

type QueryChatroomsResponseItem struct {
	AffiliationsCount int         `json:"affiliations_count"`
	Disabled          interface{} `json:"disabled"`
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Owner             string      `json:"owner"`
}
type QueryChatroomsResponse struct {
	Data []*QueryChatroomsResponseItem `json:"data"`
	Response
}

// QueryChatrooms retrieves the basic information of all chat rooms under the app by page.
func (c *Client) QueryChatrooms(ctx context.Context, q *QueryChatroomsRequest) (*QueryChatroomsResponse, error) {

	values := url.Values{}
	values.Add("limit", fmt.Sprint(q.Pagination.Limit))
	values.Add("cursor", q.Pagination.Cursor)

	var resp QueryChatroomsResponse
	err := c.makeRequest(ctx, http.MethodGet, "chatrooms", values, nil, &resp)
	return &resp, err
}

// QueryUserJoinedChatrooms retrieves all the chat rooms that a user joins.
func (c *Client) QueryUserJoinedChatrooms(ctx context.Context, userID string, q *QueryChatroomsRequest) (*QueryChatroomsResponse, error) {

	values := url.Values{}
	values.Add("limit", fmt.Sprint(q.Pagination.Limit))
	values.Add("cursor", q.Pagination.Cursor)

	var resp QueryChatroomsResponse
	p := path.Join("users", url.PathEscape(userID), "joined_chatrooms")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

type QuerySpecifiedChatroomsResponse struct {
	Data []*Chatroom `json:"data"`
	Response
}

// QuerySpecifiedChatrooms retrieves the detailed information of one or more specified chat rooms.
func (c *Client) QuerySpecifiedChatrooms(ctx context.Context, chatroomIDs ...string) (*QuerySpecifiedChatroomsResponse, error) {
	var resp QuerySpecifiedChatroomsResponse
	chatroomIDsString := strings.Join(chatroomIDs, ",")
	p := path.Join("chatrooms", url.PathEscape(chatroomIDsString))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

type UpdateChatroomRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Maxusers    int    `json:"maxusers"`
}
type UpdateChatroomResponse struct {
	Data *map[string]interface{} `json:"data"`
	Response
}

// UpdateChatroom modifies the information of the specified chat room.
func (c *Client) UpdateChatroom(ctx context.Context, chatroomID string, req *UpdateChatroomRequest) (*UpdateChatroomResponse, error) {
	var resp UpdateChatroomResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, req, &resp)
	return &resp, err
}

// DeleteChatroom deletes the specified chat room
func (c *Client) DeleteChatroom(ctx context.Context, chatroomID string) (*interface{}, error) {
	var resp interface{}
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

type QueryChatroomAnnouncementResponse struct {
	Data *map[string]interface{} `json:"data"`
	Response
}

// QueryChatroomAnnouncement retrieves the announcement text for the specified chat room.
func (c *Client) QueryChatroomAnnouncement(ctx context.Context, chatroomID string) (*QueryChatroomAnnouncementResponse, error) {
	var resp QueryChatroomAnnouncementResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

type UpdateChatroomAnnouncementRequest struct {
	Announcement string `json:"announcement"`
}

type UpdateChatroomAnnouncementResponse struct {
	Data *RestHandleResult `json:"data"`
	Response
}

// UpdateChatroomAnnouncement modifies the announcement text of the specified chat room.
func (c *Client) UpdateChatroomAnnouncement(ctx context.Context, chatroomID string, req *UpdateChatroomAnnouncementRequest) (*UpdateChatroomAnnouncementResponse, error) {
	var resp UpdateChatroomAnnouncementResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, req, &resp)
	return &resp, err
}
