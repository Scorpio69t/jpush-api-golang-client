package jpush

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	CONTENT_TYPE_FORM          = "application/x-www-form-urlencoded"
	DEFAULT_CONNECT_TIMEOUT    = 60 // Connect timeout in seconds
	DEFAULT_READ_WRITE_TIMEOUT = 60 // Read and write timeout in seconds
)

// SendPostString sends a post request and returns the response body as string
func SendPostString(url string, content, appKey, masterSecret string) (string, error) {
	req := Post(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(appKey, masterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetBody(content)

	return req.String()
}

// SendPostBytes sends a post request and returns the response body as bytes
func SendPostBytes(url string, content []byte, appKey, masterSecret string) (string, error) {
	req := Post(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(appKey, masterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetBody(content)

	return req.String()
}

// SendPostBytes2 sends a post request and returns the response body as bytes
func SendPostBytes2(url string, data []byte, appKey, masterSecret string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Add("Charset", CHARSET)
	req.SetBasicAuth(appKey, masterSecret)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)
	resp, err := client.Do(req)
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		return "", err
	}

	if resp == nil {
		return "", errors.New("response is nil")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SendGet sends a get request and returns the response body as string
func SendGet(url, appKey, masterSecret string) (string, error) {
	req := Get(url)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(appKey, masterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")

	return req.String()
}
