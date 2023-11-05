package agora_chat

import (
	"errors"
	"net/http"
	"os"
	"regexp"
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

func NewClientFromEnvVars() (*Client, error) {
	appID := os.Getenv("AGORA_APP_ID")
	appCertificate := os.Getenv("AGORA_APP_CERTIFICATE")
	baseURL := os.Getenv("AGORA_CHAT_BASEURL")
	return New(appID, appCertificate, baseURL)
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
	match, _ := regexp.MatchString("#", appID)

	if match { // if you are using AppKey + AppTken
		client.appToken = appCertificate
		return client, nil
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
