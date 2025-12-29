package jpush

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"testing"
)

type mockRoundTripper struct {
	resp     *http.Response
	err      error
	requests []*http.Request
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.requests = append(m.requests, req)
	if m.err != nil {
		return nil, m.err
	}
	return m.resp, nil
}

func newMockClient(t *testing.T, respBody string, status int) (*Client, *mockRoundTripper) {
	t.Helper()
	rt := &mockRoundTripper{
		resp: &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(bytes.NewBufferString(respBody)),
			Header:     make(http.Header),
		},
	}
	httpClient := &http.Client{Transport: rt}
	c, err := NewClient("appKey", "masterSecret", WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	return c, rt
}

func validPushPayload(t *testing.T) *PayLoad {
	t.Helper()
	var pf Platform
	_ = pf.Add(ANDROID)
	var aud Audience
	aud.SetID([]string{"id1"})
	pl := NewPayLoad()
	pl.SetPlatform(&pf)
	pl.SetAudience(&aud)
	return pl
}

func TestPushValidation(t *testing.T) {
	c, _ := newMockClient(t, `{"msg_id":1}`, http.StatusOK)
	if _, err := c.Push(context.Background(), nil); err == nil {
		t.Fatalf("expected validation error")
	}

	pl := validPushPayload(t)
	// remove audience to trigger error
	pl.SetAudience(nil)
	if _, err := c.Push(context.Background(), pl); err == nil {
		t.Fatalf("expected validation error for audience")
	}
}

func TestPushSuccess(t *testing.T) {
	c, rt := newMockClient(t, `{"msg_id":123,"sendno":456}`, http.StatusOK)
	resp, err := c.Push(context.Background(), validPushPayload(t))
	if err != nil {
		t.Fatalf("push error: %v", err)
	}
	if resp.MsgID != 123 {
		t.Fatalf("unexpected msg_id: %d", resp.MsgID)
	}
	if len(rt.requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(rt.requests))
	}
	auth := rt.requests[0].Header.Get("Authorization")
	expected := "Basic " + base64.StdEncoding.EncodeToString([]byte("appKey:masterSecret"))
	if auth != expected {
		t.Fatalf("unexpected auth header: %s", auth)
	}
}

func TestPushAPIError(t *testing.T) {
	c, _ := newMockClient(t, `{"error":{"code":1011,"message":"invalid payload"}}`, http.StatusBadRequest)
	_, err := c.Push(context.Background(), validPushPayload(t))
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.Code != 1011 {
		t.Fatalf("expected code 1011, got %d", apiErr.Code)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", apiErr.StatusCode)
	}
}

func TestSMSValidation(t *testing.T) {
	c, _ := newMockClient(t, `{"msg_id":"1"}`, http.StatusOK)
	_, err := c.SendSMS(context.Background(), &SmsPayLoad{})
	if err == nil {
		t.Fatalf("expected validation error for sms")
	}
}

func TestCIDSuccess(t *testing.T) {
	body := `{"cidlist":["a","b"]}`
	c, _ := newMockClient(t, body, http.StatusOK)
	resp, err := c.CID(context.Background(), 2, "push")
	if err != nil {
		t.Fatalf("cid error: %v", err)
	}
	if len(resp.CIDList) != 2 {
		t.Fatalf("unexpected cid list length: %d", len(resp.CIDList))
	}
}

func TestDeprecatedWrapperPush(t *testing.T) {
	rt := &mockRoundTripper{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"msg_id":123}`)),
			Header:     make(http.Header),
		},
	}
	legacy := &JPushClient{AppKey: "k", MasterSecret: "s"}
	// inject custom client for deterministic test
	SetLegacyClientOptions(WithHTTPClient(&http.Client{Transport: rt}))
	defer SetLegacyClientOptions()
	_, err := legacy.Push([]byte(`{"x":1}`))
	if err != nil {
		t.Fatalf("expected wrapper to succeed: %v", err)
	}
}
