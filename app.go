package agora_chat

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type APNSConfig struct {
	TeamId string `json:"teamId"`
	KeyId  string `json:"keyId"`
}
type HuaweiConfig struct {
	Category      string `json:"category"`
	ActivityClass string `json:"activityClass"`
}

type HonorConfig struct {
	ActivityClass string `json:"activityClass"`
}

type XiaoMiConfig struct {
	ChannelId string `json:"channelId"`
}
type VivoConfig struct {
	Category string `json:"category"`
}
type OppoConfig struct {
	Category string `json:"category"`
}

const (
	PushProviderAPNS   = PushProviderType("APNS")
	PushProviderHuaWei = PushProviderType("HUAWEIPUSH")
	PushProviderXiaoMi = PushProviderType("XIAOMIPUSH")
	PushProviderVivo   = PushProviderType("VIVOPUSH")
	PushProviderHonor  = PushProviderType("HONOR")
	PushProviderOppo   = PushProviderType("OPPOPUSH")
	PushProviderMeiZu  = PushProviderType("MEIZUPUSH")
)

type PushProviderType = string

const (
	EnvProduct     = EnvironmentType("PRODUCTION")
	EnvDevelopment = EnvironmentType("DEVELOPMENT")
)

type EnvironmentType = string

type PushProvider struct {
	Type               PushProviderType `json:"provider"`
	Name               string           `json:"name"`
	Env                EnvironmentType  `json:"environment,omitempty"`
	Certificate        string           `json:"certificate,omitempty"`
	PackageName        string           `json:"packageName,omitempty"`
	ApnsPushSettings   *APNSConfig      `json:"apnsPushSettings,omitempty"`
	HuaweiPushSettings *HuaweiConfig    `json:"huaweiPushSettings,omitempty"`
	XiaomiPushSetings  *XiaoMiConfig    `json:"xiaomiPushSetings,omitempty"`
	VivoPushSettings   *XiaoMiConfig    `json:"vivoPushSettings,omitempty"`
	HonorPushSettings  *VivoConfig      `json:"honorPushSettings,omitempty"`
	OppoPushSettings   *OppoConfig      `json:"oppoPushSettings,omitempty"`

	UUID     string `json:"uuid,omitempty"`
	Created  int64  `json:"created,omitempty"`
	Modified int64  `json:"modified,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

// InsertPushProvider inserts a push provider.
func (c *Client) InsertPushProvider(ctx context.Context, provider *PushProvider) (*PushProviderListResponse, error) {
	body := provider
	var resp PushProviderListResponse
	err := c.makeRequest(ctx, http.MethodPost, "notifiers", nil, body, &resp)
	return &resp, err
}

// DeletePushProvider deletes a push provider by uuid.
func (c *Client) DeletePushProvider(ctx context.Context, uuid string) (*PushProviderListResponse, error) {
	var resp PushProviderListResponse
	p := path.Join("notifiers", url.PathEscape(uuid))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

type PushProviderListResponse struct {
	Response
	PushProviders []PushProvider `json:"entities"`
}

// ListPushProviders returns the list of push providers.
func (c *Client) ListPushProviders(ctx context.Context) (*PushProviderListResponse, error) {
	var providers PushProviderListResponse
	err := c.makeRequest(ctx, http.MethodGet, "notifiers", nil, nil, &providers)
	return &providers, err
}

/*

{
    "name": "{{APP_ID}}",
    "provider": "HUAWEIPUSH",
    "environment": "PRODUCTION",
    "certificate": "{{APP_SECRET}}",
    "packageName": "io.github.wooEnrico.myapplication",
    "huaweiPushSettings": {
        "category": "IM",
        "activityClass": "io.github.wooEnrico.myapplication.activity.SplashActivity"
    }
}
{
    "name": "{{app_id}}",
    "provider": "XIAOMIPUSH",
    "environment": "PRODUCTION",
    "certificate": "{{app_secret}}",
    "packageName": "{{packageName}}",
    "xiaomiPushSetings": {
        "channelId": "{{channel_id}}"
    }
}
{
    "name": "{{app_id}}#{{app_key}}",
    "provider": "VIVOPUSH",
    "environment": "PRODUCTION",
    "certificate": "{{app_secret}}",
    "packageName": "{{packageName}}",
    "vivoPushSettings": {
        "category": "IM"
    }
}

{
    "name": "{{app_id}}",
    "provider": "HONOR",
    "environment": "PRODUCTION",
    "certificate": "{\"client_id\":\"{{client_id}}\",\"client_secret\":\"{{client_secret}}\"}",
    "honorPushSettings": {
        "activityClass": "{{badge_class}}"
    }
}

{
    "name": "{{app_key}}",
    "provider": "OPPOPUSH",
    "environment": "PRODUCTION",
    "certificate": "{{master_secret}}",
    "packageName": "{{packageName}}",
    "oppoPushSettings": {
        "channelId": "{{channel_id}}"
    }
}

{
    "name": "{{app_id}}",
    "provider": "MEIZUPUSH",
    "environment": "PRODUCTION",
    "certificate": "{{app_secret}}",
    "packageName": "{{packageName}}"
}
*/
