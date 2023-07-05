package agora_chat

import (
	"context"
	"errors"
	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
	"net/http"
	"os"
	"time"
)

type Client struct {
	HTTP           *http.Client `json:"-"`
	BaseURL        string
	appID          string
	appCertificate string
	authToken      string
}

type ClientOption func(c *Client)

func NewClientFromEnvVars() (*Client, error) {
	return NewClient(os.Getenv("AGORA_ID"), os.Getenv("AGORA_APPCERTIFICATE"), os.Getenv("AGORA_BASEURL"))
}

func NewClient(appID, appCertificate string, baseURL string, options ...ClientOption) (*Client, error) {
	switch {
	case appID == "":
		return nil, errors.New("app ID is empty")
	case appCertificate == "":
		return nil, errors.New("app Certificate  is empty")
	}

	tr := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second // load balancer's idle timeout is 60 sec
	tr.ExpectContinueTimeout = 2 * time.Second
	client := &Client{
		appID:          appID,
		appCertificate: appCertificate,
		BaseURL:        baseURL,
		HTTP: &http.Client{
			Timeout:   6 * time.Second,
			Transport: tr,
		},
	}
	for _, fn := range options {
		fn(client)
	}
	token, err := client.createAppToken(24 * 60 * 60)
	if err != nil {
		return nil, err
	}
	client.authToken = token
	return client, nil
}

func (c *Client) createAppToken(expire uint32) (string, error) {
	return chatTokenBuilder.BuildChatAppToken(c.appID, c.appCertificate, expire)
}

// CreateUserToken creates a new token for user
func (c *Client) CreateUserToken(userID string, expire uint32) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is empty")
	}
	ctx := context.TODO()
	result, err := c.QueryUser(ctx, userID)
	if err != nil {
		return "", err
	}
	return c.createUserToken(result.Users[0].Uuid, expire)
}

func (c *Client) createUserToken(Uuid string, expire uint32) (string, error) {
	return chatTokenBuilder.BuildChatUserToken(c.appID, c.appCertificate, Uuid, expire)
}
