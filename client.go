package agora_chat

import (
	"errors"
	"net/http"
	"time"

	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
)

type Client struct {
	http           *http.Client `json:"-"`
	baseURL        string
	appID          string
	appCertificate string
	appToken       string
}

func New(appID, appCertificate string, baseURL string) (*Client, error) {
	switch {
	case appID == "":
		return nil, errors.New("app ID is empty")
	case appCertificate == "":
		return nil, errors.New("app Certificate  is empty")
	case baseURL == "":
		return nil, errors.New("chat BaseUrl is empty")
	}
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second
	tr.ExpectContinueTimeout = 2 * time.Second
	client := &Client{
		appID:          appID,
		appCertificate: appCertificate,
		baseURL:        baseURL,
		http: &http.Client{
			Timeout:   6 * time.Second,
			Transport: tr,
		},
	}
	token, err := client.createAppToken(uint32(time.Now().Unix()) + 7200) // 2h
	if err != nil {
		return nil, err
	}
	client.appToken = token
	return client, nil
}

func (c *Client) createAppToken(expire uint32) (string, error) {
	return chatTokenBuilder.BuildChatAppToken(c.appID, c.appCertificate, expire)
}

func (c *Client) createUserToken(Uuid string, expire uint32) (string, error) {
	return chatTokenBuilder.BuildChatUserToken(c.appID, c.appCertificate, Uuid, expire)
}
