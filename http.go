package agora_chat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Error struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ExceptionFields map[string]string `json:"exception_fields,omitempty"`
	StatusCode      int               `json:"StatusCode"`
	Duration        string            `json:"duration"`
	MoreInfo        string            `json:"more_info"`
}

func (e Error) Error() string {
	return e.Message
}

type Response struct {
	Path      string `json:"path"`
	Uri       string `json:"uri"`
	Timestamp int64  `json:"timestamp"`
	Count     int    `json:"count"`
	Action    string `json:"action"`
	Duration  int    `json:"duration"`
}

func (c *Client) parseResponse(resp *http.Response, result interface{}) error {
	if resp.Body == nil {
		return errors.New("http body is nil")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response: %w", err)
	}

	if resp.StatusCode >= 399 {
		var apiErr Error
		err := json.Unmarshal(b, &apiErr)
		if err != nil {
			// IP rate limit errors sent by our Edge infrastructure are not JSON encoded.
			// If decode fails here, we need to handle this manually.
			apiErr.Message = string(b)
			apiErr.StatusCode = resp.StatusCode
			return apiErr
		}
		return apiErr
	}

	if _, ok := result.(*Response); !ok {
		// Unmarshal the body only when it is expected.
		err = json.Unmarshal(b, result)
		if err != nil {
			return fmt.Errorf("cannot unmarshal body: %w", err)
		}
	}

	return nil
}

func (c *Client) requestURL(path string, values url.Values) (string, error) {
	u, err := url.Parse(c.BaseURL + "/" + path)
	if err != nil {
		return "", errors.New("url.Parse: " + err.Error())
	}

	if values == nil {
		values = make(url.Values)
	}
	return u.String(), nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, params url.Values, data interface{}) (*http.Request, error) {
	u, err := c.requestURL(path, params)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, u, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(r)
	switch t := data.(type) {
	case nil:
		r.Body = nil

	case io.ReadCloser:
		r.Body = t

	case io.Reader:
		r.Body = io.NopCloser(t)

	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(b))
	}

	return r, nil
}

func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+c.authToken)
}

func (c *Client) makeRequest(ctx context.Context, method, path string, params url.Values, data, result interface{}) error {
	r, err := c.newRequest(ctx, method, path, params, data)
	if err != nil {
		return err
	}
	resp, err := c.HTTP.Do(r)
	if err != nil {
		select {
		case <-ctx.Done():
			// If we got an error, and the context has been canceled,
			// return context's error which is more useful.
			return ctx.Err()
		default:
		}
		return err
	}
	return c.parseResponse(resp, result)
}
