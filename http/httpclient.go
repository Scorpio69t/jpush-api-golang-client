package httplib

import (
	"strconv"
	"time"
)

const (
	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	CONTENT_TYPE_FORM          = "application/x-www-form-urlencoded"
	DEFAULT_CONNECT_TIMEOUT    = 60 // Connect timeout in seconds
	DEFAULT_READ_WRITE_TIMEOUT = 60 // Read and write timeout in seconds
)

const (
	cid_url = "https://api.jpush.cn/v3/push/cid"
)

type JPushClient struct {
	appKey    string
	appSecret string
}

// NewJPushClient returns a new *JPushClient
func NewJPushClient(appKey string, appSecret string) *JPushClient {
	return &JPushClient{appKey: appKey, appSecret: appSecret}
}

// GetAuthorization returns the authorization string
func (c *JPushClient) GetAuthorization() string {
	return c.appKey + ":" + c.appSecret
}

// SendPostString sends a post request and returns the response body as string
func (c *JPushClient) SendPostString(url string, content string) (string, error) {
	req := Post(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(c.appKey, c.appSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetBody(content)

	return req.String()
}

func (c *JPushClient) SendPostBytes(url string, content []byte) (string, error) {
	req := Post(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(c.appKey, c.appSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetBody(content)

	return req.String()
}

func (c *JPushClient) SendGet(url string) (string, error) {
	req := Get(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(c.appKey, c.appSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")

	return req.String()
}

func (c *JPushClient) GetCid(count int, push_type string) ([]byte, error) {
	req := Get(cid_url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(c.appKey, c.appSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetQueryParam("count", strconv.Itoa(count))
	req.SetQueryParam("type", push_type)

	return req.Bytes()
}
