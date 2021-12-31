package push

import (
	"fmt"

	"github.com/Scorpio69t/jpush-api-golang-client/http/httplib"
)

const (
	base_url = "https://api.jpush.cn/v3/push/cid"
)

type CidRequest struct {
	Count int    `json:"count,omitempty"` // 数值类型，不传则默认为 1。范围为 [1, 1000]
	Type  string `json:"type,omitempty"`  // CID 类型。取值：push（默认），schedule
}

type CidResponse struct {
	CidList []string `json:"cid_list,omitempty"` // CID 列表
}

// NewCidRequest 创建 CidRequest 对象
func NewCidRequest(count int, cid_type string) *CidRequest {
	c := &CidRequest{}
	if count <= 0 || count > 1000 {
		c.Count = 1
	} else {
		c.Count = count
	}

	if cid_type == "" {
		c.Type = "push"
	} else {
		c.Type = cid_type
	}

	return c
}

// getUrl 获取请求的 url
func (c *CidRequest) getUrl() string {
	return fmt.Sprintf("%s?count=%d&type=%s", base_url, c.Count, c.Type)
}

// GetCidList 获取 CID 列表
func (c *CidRequest) GetCidList() ([]string, error) {
	var cid_list []string
	jc := httplib.NewJPushClient("", "")

	return cid_list, nil
}
