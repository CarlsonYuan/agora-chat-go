package agora_chat

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type MessageBody struct {
	Action string `json:"action"`
}
type Message struct {
	Type string       `json:"type"`
	To   []string     `json:"to"`
	Body *MessageBody `json:"body"`
	From string       `json:"from"`
}

// SendUploadLogsCommand sends a cmd message by user IDS
// Once the users online, Chat SDK logs will be automatically uploaded.
func (c *Client) SendUploadLogsCommand(ctx context.Context, userIDs ...string) error {
	cmd := Message{
		From: "admin",
		Type: "users",
		To:   userIDs,
		Body: &MessageBody{Action: "em_upload_log"},
	}
	err := c.makeRequest(context.Background(), http.MethodPost, "messages/users", nil, cmd, nil)
	return err
}

type DeviceLogsListResponse struct {
	Response
	LogFile []LogFile `json:"entities"`
}
type LogFile struct {
	OsVersion     string `json:"os_version,omitempty"`
	LogfileUUID   string `json:"logfile_uuid,omitempty"`
	LoginUsername string `json:"login_username,omitempty"`
	SdkVersion    string `json:"sdk_version,omitempty"`
	UploadDate    string `json:"uploadDate,omitempty"`
	Created       int64  `json:"created,omitempty"`
	Modified      int64  `json:"modified,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

// ListDeviceLogs lists the details of the log files on device
func (c *Client) ListDeviceLogs(ctx context.Context, uID string) (*DeviceLogsListResponse, error) {
	var devicelogs DeviceLogsListResponse
	p := path.Join("users", url.PathEscape(uID), "devicelogs")
	err := c.makeRequest(context.Background(), http.MethodGet, p, nil, nil, &devicelogs)
	return &devicelogs, err
}
