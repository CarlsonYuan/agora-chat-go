package agora_chat

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
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
		Type: "cmd",
		To:   userIDs,
		Body: &MessageBody{Action: "em_upload_log"},
	}
	err := c.makeRequest(ctx, http.MethodPost, "messages/users", nil, cmd, nil)
	return err
}

type DeviceLogsListResponse struct {
	Response
	LogFiles []LogFile `json:"entities"`
}
type LogFile struct {
	AppKey        string `json:"appkey,omitempty"`
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
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &devicelogs)
	return &devicelogs, err
}

func (c *Client) DownloadFile(fileUuid string, baseDir string, fileName string) error {

	u, err := url.Parse(c.baseURL)
	if err != nil {
		panic(err)
	}
	p := path.Join("easemob/logger/chatfiles", url.PathEscape(fileUuid))
	u.Path = p
	//Get the response bytes from the url
	response, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	EnsureBaseDir(baseDir)
	//Create a empty file
	file, err := os.Create(baseDir + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func EnsureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
