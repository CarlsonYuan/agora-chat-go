package agora_chat

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
)

type OfflineCountResponse struct {
	Data map[string]int `json:"data"`
	Response
}
type ResultResponse struct {
	Data interface{} `json:"data"`
	Response
}
type entities struct {
	Uuid        string `json:"uuid"`
	Type        string `json:"type"`
	ShareSecret string `json:"share-secret"`
}
type UploadingResponse struct {
	Entities []entities `json:"entities"`
	Response
}
type ChannelParam struct {
	Channel    string `json:"channel"`
	Type       string `json:"type"`
	DeleteRoam bool   `json:"delete_roam"`
}
type MessageDownloadParam struct {
	Dir  string `json:"dir"`
	Time string `json:"time"`
}
type MsgRecallParam struct {
	MsgId    string `json:"msg_id"`
	To       string `json:"to"`
	From     string `json:"from"`
	ChatType string `json:"chat_type"`
	Force    bool   `json:"force"`
}

// CountMissedMessages  Query the number of user offline messages
func (c *Client) CountMissedMessages(ctx context.Context, userID string) (*OfflineCountResponse, error) {
	if len(userID) == 0 {
		return nil, errors.New("userID is nil")
	}
	p := path.Join("users", url.PathEscape(userID), "offline_msg_count")

	var resp OfflineCountResponse
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// DeleteChannel The server removes the session in one direction
func (c *Client) DeleteChannel(ctx context.Context, userID string, param *ChannelParam) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "user_channel")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, param, &resp)
	return &resp, err
}

// GetHistoryAsUri Get the download address of the message history file
func (c *Client) GetHistoryAsUri(ctx context.Context, time string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatmessages", url.PathEscape(time))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// ImportChatMessage Importing messages
func (c *Client) ImportChatMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/users/import", nil, msg, &resp)
	return &resp, err
}

// ImportGroupMessage Import group messages
func (c *Client) ImportGroupMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups/import", nil, msg, &resp)
	return &resp, err
}

// IsMessageDeliveredToUser Query the status of an offline message, such as whether it has been delivered
func (c *Client) IsMessageDeliveredToUser(ctx context.Context, toUser, messageId string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(toUser), "offline_msg_status", url.PathEscape(messageId))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// RecallMsg Recall message
func (c *Client) RecallMsg(ctx context.Context, param *MsgRecallParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/msg_recall", nil, param, &resp)
	return &resp, err
}

// SendChatMessage Send chat message
func (c *Client) SendChatMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/users?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendGroupsMessage Send group chat messages
func (c *Client) SendGroupsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendRoomsMessage Send chat room messages
func (c *Client) SendRoomsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatrooms?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// UploadingChatFile upload files
func (c *Client) UploadingChatFile(ctx context.Context, filePath string) (*UploadingResponse, error) {
	var resp UploadingResponse
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	_, err = io.Copy(part, file)
	writer.Close()
	err = c.UploadingFile(ctx, http.MethodPost, "chatfiles", writer.FormDataContentType(), body, &resp)
	return &resp, err
}
