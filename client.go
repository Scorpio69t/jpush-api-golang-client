package jpush

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultTimeout   = 60 * time.Second
	defaultUserAgent = "jpush-api-golang-client"
)

// Client is the modern context-first client with injectable transport.
type Client struct {
	appKey       string
	masterSecret string
	httpClient   *http.Client
	endpoints    EndpointSet
	userAgent    string
}

// EndpointSet holds the JPush endpoints used by the client.
type EndpointSet struct {
	Push     string
	Schedule string
	Report   string
	CID      string
	SMS      string
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient injects a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithTimeout sets the default http.Client timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = timeout
	}
}

// WithBaseURLs overrides default endpoints.
func WithBaseURLs(set EndpointSet) Option {
	return func(c *Client) {
		c.endpoints = set
	}
}

// WithUserAgent overrides the default User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		if ua != "" {
			c.userAgent = ua
		}
	}
}

// NewClient creates a context-aware Client with optional configuration.
func NewClient(appKey, masterSecret string, opts ...Option) (*Client, error) {
	if appKey == "" || masterSecret == "" {
		return nil, &ValidationError{Reason: "appKey/masterSecret is required"}
	}

	client := &Client{
		appKey:       appKey,
		masterSecret: masterSecret,
		httpClient:   &http.Client{Timeout: defaultTimeout},
		endpoints: EndpointSet{
			Push:     HOST_PUSH,
			Schedule: HOST_SCHEDULE,
			Report:   HOST_REPORT,
			CID:      HOST_CID,
			SMS:      SMS,
		},
		userAgent: defaultUserAgent,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{Timeout: defaultTimeout}
	}

	return client, nil
}

// Push sends a push payload using the JPush Push API v3.
func (c *Client) Push(ctx context.Context, payload *PayLoad) (*PushResponse, error) {
	if err := ValidatePushPayload(payload); err != nil {
		return nil, err
	}

	body, err := payload.Bytes()
	if err != nil {
		return nil, fmt.Errorf("marshal push payload: %w", err)
	}

	var resp PushResponse
	if err := c.do(ctx, http.MethodPost, c.endpoints.Push, body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ScheduleCreate creates a schedule push task.
func (c *Client) ScheduleCreate(ctx context.Context, schedule *Schecule) (*ScheduleResponse, error) {
	if err := ValidateSchedule(schedule); err != nil {
		return nil, err
	}

	body, err := schedule.Bytes()
	if err != nil {
		return nil, fmt.Errorf("marshal schedule payload: %w", err)
	}

	var resp ScheduleResponse
	if err := c.do(ctx, http.MethodPost, c.endpoints.Schedule, body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ScheduleGet retrieves a schedule by id.
func (c *Client) ScheduleGet(ctx context.Context, id string) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Reason: "schedule id is required"}
	}
	url := c.endpoints.Schedule + "?schedule_id=" + id
	var resp ScheduleResponse
	if err := c.do(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ScheduleDelete deletes a schedule by id.
func (c *Client) ScheduleDelete(ctx context.Context, id string) error {
	if id == "" {
		return &ValidationError{Reason: "schedule id is required"}
	}
	url := c.endpoints.Schedule + "?schedule_id=" + id
	return c.do(ctx, http.MethodDelete, url, nil, nil)
}

// CID fetches push cid list.
func (c *Client) CID(ctx context.Context, count int, pushType string) (*CIDResponse, error) {
	if count <= 0 {
		return nil, &ValidationError{Reason: "count must be greater than zero"}
	}
	url := fmt.Sprintf("%s?count=%d", c.endpoints.CID, count)
	if pushType != "" {
		url += "&type=" + pushType
	}

	var resp CIDResponse
	if err := c.do(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendSMS sends SMS via SMS API v1.
func (c *Client) SendSMS(ctx context.Context, payload *SmsPayLoad) (*SMSResponse, error) {
	if err := ValidateSMSPayload(payload); err != nil {
		return nil, err
	}

	body, err := payload.Bytes()
	if err != nil {
		return nil, fmt.Errorf("marshal sms payload: %w", err)
	}

	var resp SMSResponse
	if err := c.do(ctx, http.MethodPost, c.endpoints.SMS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// do executes an HTTP request with JSON handling and Basic Auth.
func (c *Client) do(ctx context.Context, method, url string, body []byte, out interface{}) error {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	req.SetBasicAuth(c.appKey, c.masterSecret)
	req.Header.Set("Charset", CHARSET)
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", CONTENT_TYPE_JSON)
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return parseAPIError(resp.StatusCode, respBody)
	}

	if out == nil || len(respBody) == 0 {
		return nil
	}

	if err := json.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

// doRaw is used by deprecated wrappers to preserve string return types.
func (c *Client) doRaw(ctx context.Context, method, url string, body []byte) (string, error) {
	var sink map[string]interface{}
	err := c.do(ctx, method, url, body, &sink)
	if apiErr, ok := err.(*APIError); ok {
		return string(apiErr.RawBody), apiErr
	}
	if err != nil {
		return "", err
	}
	if sink == nil {
		return "", nil
	}
	out, marshalErr := json.Marshal(sink)
	if marshalErr != nil {
		return "", marshalErr
	}
	return string(out), nil
}
