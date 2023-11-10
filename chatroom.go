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
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Maxusers    int      `json:"maxusers"`
	Owner       string   `json:"owner"`
	Members     []string `json:"members"`
}

type ChatroomResponse struct {
	Data []CreateUserError `json:"data"`
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

func (c *Client) QueryChatrooms(ctx context.Context, q *QueryUsersRequest) (*interface{}, error) {

	values := url.Values{}
	values.Add("limit", fmt.Sprint(q.Pagination.Limit))
	values.Add("cursor", q.Pagination.Cursor)

	var resp interface{}
	err := c.makeRequest(ctx, http.MethodGet, "chatrooms", values, nil, &resp)
	return &resp, err
}

func (c *Client) QueryUserJoinedChatrooms(ctx context.Context, userID string, q *QueryUsersRequest) (*interface{}, error) {

	values := url.Values{}
	values.Add("limit", fmt.Sprint(q.Pagination.Limit))
	values.Add("cursor", q.Pagination.Cursor)

	var resp interface{}
	p := path.Join("users", url.PathEscape(userID), "joined_chatrooms")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

func (c *Client) QuerySpecifiedChatrooms(ctx context.Context, chatroomIDs ...string) (*interface{}, error) {

	var resp interface{}
	chatroomIDsString := strings.Join(chatroomIDs, ",")
	p := path.Join("chatrooms", url.PathEscape(chatroomIDsString))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}
func (c *Client) UpdateChatroom(ctx context.Context, chatroom *Chatroom) (*interface{}, error) {
	req := chatroom
	var resp interface{}
	err := c.makeRequest(ctx, http.MethodPut, "chatrooms", nil, req, &resp)
	return &resp, err
}

func (c *Client) DeleteChatroom(ctx context.Context, chatroomID string) (*interface{}, error) {
	var resp interface{}
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

func (c *Client) QueryChatroomAnnouncement(ctx context.Context, chatroomID string) (*interface{}, error) {
	var resp interface{}
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

type UpdateChatroomAnnouncementRequest struct {
	Announcement string `json:"announcement"`
}

func (c *Client) UpdateChatroomAnnouncement(ctx context.Context, chatroomID string, req UpdateChatroomAnnouncementRequest) (*interface{}, error) {
	var resp interface{}
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, req, &resp)
	return &resp, err
}
