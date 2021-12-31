package jpush

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type HttpRequest struct {
	url              string
	req              *http.Request
	params           map[string]string
	connectTimeout   time.Duration
	readWriteTimeout time.Duration
	tlsConfig        *tls.Config
	proxy            func(*http.Request) (*url.URL, error)
	transport        http.RoundTripper
}

// Get returns *HttpRequest with GET method.
func Get(url string) *HttpRequest {
	var req http.Request
	req.Method = "GET"
	req.Header = make(http.Header)

	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// Post returns *HttpRequest with POST method.
func Post(url string) *HttpRequest {
	var req http.Request
	req.Method = "POST"
	req.Header = make(http.Header)

	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// Delete returns *HttpRequest with DELETE method.
func Delete(url string) *HttpRequest {
	var req http.Request
	req.Method = "DELETE"
	req.Header = make(http.Header)

	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// Put returns *HttpRequest with PUT method.
func Put(url string) *HttpRequest {
	var req http.Request
	req.Method = "PUT"
	req.Header = make(http.Header)

	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// SetQueryParam replaces the request query values.
func (h *HttpRequest) SetQueryParam(key, value string) *HttpRequest {
	if h.req.URL == nil {
		return h
	}
	q := h.req.URL.Query()
	q.Add(key, value)
	h.req.URL.RawQuery = q.Encode()
	return h
}

// SetTimeout sets connect time out and read-write time out for Request.
func (h *HttpRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *HttpRequest {
	h.connectTimeout = connectTimeout
	h.readWriteTimeout = readWriteTimeout
	return h
}

// SetTLSConfig sets tls connection configurations if visiting https url.
func (h *HttpRequest) SetTLSConfig(config *tls.Config) *HttpRequest {
	h.tlsConfig = config
	return h
}

// SetBasicAuth sets the basic authentication header
func (h *HttpRequest) SetBasicAuth(username, password string) *HttpRequest {
	h.req.SetBasicAuth(username, password)
	return h
}

// SetUserAgent sets User-Agent header field
func (h *HttpRequest) SetUserAgent(useragent string) *HttpRequest {
	h.req.Header.Set("User-Agent", useragent)
	return h
}

// SetHeader sets header field
func (h *HttpRequest) SetHeader(key, value string) *HttpRequest {
	h.req.Header.Set(key, value)
	return h
}

// SetProtocolVersion sets the protocol version for the request.
func (h *HttpRequest) SetProtocolVersion(vers string) *HttpRequest {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}

	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		h.req.Proto = vers
		h.req.ProtoMajor = major
		h.req.ProtoMinor = minor
	}

	return h
}

// SetCookie add cookie into request.
func (h *HttpRequest) SetCookie(cookie *http.Cookie) *HttpRequest {
	h.req.Header.Add("Cookie", cookie.String())
	return h
}

// SetTransport sets Transport to HttpClient.
func (h *HttpRequest) SetTransport(transport http.RoundTripper) *HttpRequest {
	h.transport = transport
	return h
}

// SetProxy sets proxy for HttpClient.
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (h *HttpRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *HttpRequest {
	h.proxy = proxy
	return h
}

// SetParam adds query param in to request.
func (h *HttpRequest) SetParam(key, value string) *HttpRequest {
	h.params[key] = value
	return h
}

// SetBody sets request body.
// It supports string, []byte, url.Values, map[string]interface{} and io.Reader.
func (h *HttpRequest) SetBody(body interface{}) *HttpRequest {
	if h.req.Body == nil {
		h.req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("")))
	}
	switch b := body.(type) {
	case []byte:
		h.req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		h.req.ContentLength = int64(len(b))
	case string:
		h.req.Body = ioutil.NopCloser(bytes.NewBufferString(b))
		h.req.ContentLength = int64(len(b))
	case io.Reader:
		h.req.Body = ioutil.NopCloser(b)
		if h.req.ContentLength == -1 {
			if rc, ok := b.(io.ReadCloser); ok {
				h.req.Body = rc
			}
		}
	case url.Values:
		h.req.Body = ioutil.NopCloser(bytes.NewBufferString(b.Encode()))
		h.req.ContentLength = int64(len(b.Encode()))
	case map[string]interface{}:
		v := url.Values{}
		for k, val := range b {
			switch val.(type) {
			case string:
				v.Set(k, val.(string))
			case bool:
				v.Set(k, strconv.FormatBool(val.(bool)))
			case int, int8, int16, int32, int64:
				v.Set(k, strconv.FormatInt(int64(val.(int)), 10))
			case uint, uint8, uint16, uint32, uint64:
				v.Set(k, strconv.FormatUint(uint64(val.(uint)), 10))
			case float32:
				v.Set(k, strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32))
			case float64:
				v.Set(k, strconv.FormatFloat(val.(float64), 'f', -1, 64))
			}
		}
		h.req.Body = ioutil.NopCloser(bytes.NewBufferString(v.Encode()))
		h.req.ContentLength = int64(len(v.Encode()))
	default:
		if v, ok := body.(io.ReadCloser); ok {
			h.req.Body = v
		} else {
			h.req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", body))))
		}
	}

	return h
}

// GetHeader returns header data.
func (h *HttpRequest) GetHeader() http.Header {
	return h.req.Header
}

// GetCookie returns the request cookies.
func (h *HttpRequest) GetCookie(key string) string {
	for _, cookie := range h.req.Cookies() {
		if cookie.Name == key {
			return cookie.Value
		}
	}
	return ""
}

// GetParam returns the request parameter value according to its key.
func (h *HttpRequest) GetParam(key string) string {
	return h.params[key]
}

// getResponse executes the request and returns the response.
func (h *HttpRequest) getResponse() (*http.Response, error) {
	var paramBody string
	if len(h.params) > 0 {
		v := url.Values{}
		for k, val := range h.params {
			v.Set(k, val)
		}
		paramBody = v.Encode()
	}

	if h.req.Body != nil {
		if strings.Contains(h.req.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			b, _ := ioutil.ReadAll(h.req.Body)
			paramBody = string(b) + "&" + paramBody
		} else {
			return nil, errors.New("please use SetBody method instead")
		}
	}

	if h.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Contains(h.url, "?") {
			h.url += "&" + paramBody
		} else {
			h.url += "?" + paramBody
		}
	} else if h.req.Method == "POST" && h.req.Body == nil && len(paramBody) > 0 {
		h.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		h.SetBody(paramBody)
	}

	url, err := url.Parse(h.url)
	if err != nil {
		return nil, err
	}

	if url.Scheme == "" {
		h.url = "http://" + h.url
		url, err = url.Parse(h.url)
		if err != nil {
			return nil, err
		}
	}

	h.req.URL = url
	trans := h.transport

	if trans == nil {
		trans = &http.Transport{
			TLSClientConfig: h.tlsConfig,
			Proxy:           h.proxy,
			Dial:            TimeoutDialer(h.connectTimeout, h.readWriteTimeout),
		}
	} else {
		if t, ok := trans.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = h.tlsConfig
			}
			if t.Proxy == nil {
				t.Proxy = h.proxy
			}
			if t.Dial == nil {
				t.Dial = TimeoutDialer(h.connectTimeout, h.readWriteTimeout)
			}
		}
	}

	client := &http.Client{
		Transport: trans,
	}

	resp, err := client.Do(h.req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(connectTimeout time.Duration, readWriteTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, connectTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(readWriteTimeout))
		return conn, nil
	}
}

// Response executes request client gets response in the return.
func (h *HttpRequest) Response() (*http.Response, error) {
	return h.getResponse()
}

// Bytes executes request client gets response body in bytes.
func (h *HttpRequest) Bytes() ([]byte, error) {
	resp, err := h.getResponse()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// String returns the body string in response.
// it calls Response internally.
func (h *HttpRequest) String() (string, error) {
	data, err := h.Bytes()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Status returns the response status.
func (h *HttpRequest) Status() int {
	resp, _ := h.Response()
	return resp.StatusCode
}

// ToFile saves the body data in response to one file.
func (h *HttpRequest) ToFile(file string) error {
	resp, err := h.Response()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// ToJson returns the map that converted from response body bytes as json in response .
func (h *HttpRequest) ToJson(v interface{}) error {
	data, err := h.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// ToXml returns the map that converted from response body bytes as xml in response .
func (h *HttpRequest) ToXml(v interface{}) error {
	data, err := h.Bytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}
