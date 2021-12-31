package jpush

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

type JPushClient struct {
	AppKey       string // app key
	MasterSecret string // master secret
}

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_PUSH     = "https://api.jpush.cn/v3/push"
	HOST_SCHEDULE = "https://api.jpush.cn/v3/schedules"
	HOST_REPORT   = "https://report.jpush.cn/v3/received"
	HOST_CID      = "https://api.jpush.cn/v3/push/cid"
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
	req := Get(HOST_CID)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(j.AppKey, j.MasterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetQueryParam("count", strconv.Itoa(count))
	req.SetQueryParam("type", push_type)

	return req.Bytes()
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
	ret, err := SendPostString(HOST_PUSH, content, j.AppKey, j.MasterSecret)
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
	ret, err := SendPostBytes2(HOST_PUSH, content, j.AppKey, j.MasterSecret)
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
	ret, err := SendPostBytes2(HOST_SCHEDULE, content, j.AppKey, j.MasterSecret)
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
	req := Get(HOST_SCHEDULE)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(j.AppKey, j.MasterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetQueryParam("schedule_id", schedule_id)

	return req.String()
}

// SendDeleteScheduleRequest sends a delete schedule request and returns the response body as string
func (j *JPushClient) sendDeleteScheduleRequest(schedule_id string) (string, error) {
	req := Delete(HOST_SCHEDULE)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(j.AppKey, j.MasterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetQueryParam("schedule_id", schedule_id)

	return req.String()
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
