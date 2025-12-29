package jpush

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type JPushClient struct {
	AppKey       string // app key
	MasterSecret string // master secret
}

const (
	SUCCESS_FLAG  = "msg_id"
	SMS           = "https://api.sms.jpush.cn/v1/messages"
	HOST_PUSH     = "https://api.jpush.cn/v3/push"
	HOST_SCHEDULE = "https://api.jpush.cn/v3/schedules"
	HOST_REPORT   = "https://report.jpush.cn/v3/received"
	HOST_CID      = "https://api.jpush.cn/v3/push/cid"
	HOST_IMAGES   = "https://api.jpush.cn/v3/images"
)

// NewJPushClient returns a new JPushClient
func NewJPushClient(appKey string, masterSecret string) *JPushClient {
	return &JPushClient{AppKey: appKey, MasterSecret: masterSecret}
}

// GetAuthorization returns the authorization string
func (j *JPushClient) GetAuthorization() string {
	return j.AppKey + ":" + j.MasterSecret
}

// GetCid returns the cid list as byte array
func (j *JPushClient) GetCid(count int, push_type string) ([]byte, error) {
	client, err := j.newModernClient()
	if err != nil {
		return nil, err
	}
	resp, err := client.CID(context.Background(), count, push_type)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

// 发送短信
func (j *JPushClient) SendSms(data []byte) (string, error) {
	return j.sendSmsBytes(data)
}
func (j *JPushClient) sendSmsBytes(content []byte) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	ret, err := client.doRaw(context.Background(), http.MethodPost, client.endpoints.SMS, content)
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, SUCCESS_FLAG) {
		return ret, nil
	}

	return "", errors.New(ret)
}

// Push 推送消息
func (j *JPushClient) Push(data []byte) (string, error) {
	return j.sendPushBytes(data)
}

// CreateSchedule 创建推送计划
func (j *JPushClient) CreateSchedule(data []byte) (string, error) {
	return j.sendScheduleBytes(data)
}

// DeleteSchedule 删除推送计划
func (j *JPushClient) DeleteSchedule(id string) (string, error) {
	return j.sendDeleteScheduleRequest(id)
}

// GetSchedule 获取推送计划
func (j *JPushClient) GetSchedule(id string) (string, error) {
	return j.sendGetScheduleRequest(id)
}

// SendPushString sends a push request and returns the response body as string
func (j *JPushClient) sendPushString(content string) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	ret, err := client.doRaw(context.Background(), http.MethodPost, client.endpoints.Push, []byte(content))
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, SUCCESS_FLAG) {
		return ret, nil
	}

	return "", errors.New(ret)
}

// SendPushBytes sends a push request and returns the response body as string
func (j *JPushClient) sendPushBytes(content []byte) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	ret, err := client.doRaw(context.Background(), http.MethodPost, client.endpoints.Push, content)
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, SUCCESS_FLAG) {
		return ret, nil
	}

	return "", errors.New(ret)
}

// SendScheduleBytes sends a schedule request and returns the response body as string
func (j *JPushClient) sendScheduleBytes(content []byte) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	ret, err := client.doRaw(context.Background(), http.MethodPost, client.endpoints.Schedule, content)
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	}

	return "", errors.New(ret)
}

// SendGetScheduleRequest sends a get schedule request and returns the response body as string
func (j *JPushClient) sendGetScheduleRequest(schedule_id string) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	url := client.endpoints.Schedule + "?schedule_id=" + schedule_id
	return client.doRaw(context.Background(), http.MethodGet, url, nil)
}

// SendDeleteScheduleRequest sends a delete schedule request and returns the response body as string
func (j *JPushClient) sendDeleteScheduleRequest(schedule_id string) (string, error) {
	client, err := j.newModernClient()
	if err != nil {
		return "", err
	}
	url := client.endpoints.Schedule + "?schedule_id=" + schedule_id
	return client.doRaw(context.Background(), http.MethodDelete, url, nil)
}

// UnmarshalResponse unmarshals the response body to the map
func UnmarshalResponse(resp string) (map[string]interface{}, error) {
	var ret map[string]interface{}

	if len(strings.TrimSpace(resp)) == 0 {
		return ret, errors.New("empty response")
	}

	err := json.Unmarshal([]byte(resp), &ret)
	if err != nil {
		return nil, err
	}

	if _, ok := ret["error"]; ok {
		return nil, errors.New(resp)
	}

	return ret, nil
}

func (j *JPushClient) newModernClient() (*Client, error) {
	opts := []Option{WithTimeout(DEFAULT_CONNECT_TIMEOUT * time.Second)}
	opts = append(opts, legacyClientOptions...)
	return NewClient(j.AppKey, j.MasterSecret, opts...)
}

// SetLegacyClientOptions injects options for legacy wrappers (primarily for testing).
func SetLegacyClientOptions(opts ...Option) {
	legacyClientOptions = opts
}

var legacyClientOptions []Option
