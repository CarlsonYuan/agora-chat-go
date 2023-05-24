package agora_chat

import (
	"errors"
	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
)

type Client struct {
	BaseURL        string
	appID          string
	appCertificate string
	authToken      string
}

type ClientOption func(c *Client)

func NewClient(appID, appCertificate string, baseURL string, options ...ClientOption) (*Client, error) {
	switch {
	case appID == "":
		return nil, errors.New("app ID is empty")
	case appCertificate == "":
		return nil, errors.New("app Certificate  is empty")
	}

	client := &Client{
		appID:          appID,
		appCertificate: appCertificate,
		BaseURL:        baseURL,
	}
	for _, fn := range options {
		fn(client)
	}
	token, err := client.CreateAppToken(appID, appCertificate, uint32(7200))
	if err != nil {
		return nil, err
	}
	client.authToken = token
	return client, nil
}

func (c *Client) CreateAppToken(appID, appCertificate string, expire uint32) (string, error) {
	return chatTokenBuilder.BuildChatAppToken(appID, appCertificate, expire)
}
