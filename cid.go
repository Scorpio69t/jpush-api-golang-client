package jpush

import (
	"encoding/json"
	"fmt"
)

type CidRequest struct {
	Count int    `json:"count,omitempty"` // 数值类型，不传则默认为 1。范围为 [1, 1000]
	Type  string `json:"type,omitempty"`  // CID 类型。取值：push（默认），schedule
}

type CidResponse struct {
	CidList []string `json:"cidlist,omitempty"` // CID 列表
}

func (c *CidRequest) String() string {
	return fmt.Sprintf("CidRequest{Count: %d, Type: %s}", c.Count, c.Type)
}

func (c *CidResponse) String() string {
	return fmt.Sprintf("CidResponse{CidList: %v}", c.CidList)
}

// NewCidRequest 创建 CidRequest 对象
func NewCidRequest(count int, pushType string) *CidRequest {
	c := &CidRequest{}
	if count <= 0 || count > 1000 {
		c.Count = 1
	} else {
		c.Count = count
	}

	if pushType == "" {
		c.Type = "push"
	} else {
		c.Type = pushType
	}

	return c
}

// GetCidList 获取 CID 列表
func (c *CidRequest) GetCidList(key, secret string) (*CidResponse, error) {
	resp := &CidResponse{}
	jc := NewJPushClient(key, secret)

	data, err := jc.GetCid(c.Count, c.Type)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", string(data))

	err = json.Unmarshal(data, resp)

	return resp, err
}
