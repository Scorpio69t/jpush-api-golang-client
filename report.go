package jpush

import "time"

// GetReport 获取消息推送结果
func (j *JPushClient) GetReport(msg_ids string) (string, error) {
	return j.sendGetReportRequest(msg_ids)
}

// SendGetReportRequest sends a get report request and returns the response body as string
func (j *JPushClient) sendGetReportRequest(msg_ids string) (string, error) {
	req := Get(HOST_REPORT)
	req.SetTimeout(DEFAULT_CONNECT_TIMEOUT*time.Second, DEFAULT_READ_WRITE_TIMEOUT*time.Second)
	req.SetHeader("Connection", "Keep-Alive")
	req.SetHeader("Charset", CHARSET)
	req.SetBasicAuth(j.AppKey, j.MasterSecret)
	req.SetHeader("Content-Type", CONTENT_TYPE_JSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.SetQueryParam("msg_ids", msg_ids)

	return req.String()
}
