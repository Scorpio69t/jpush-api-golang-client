package jpush

type CallBack struct {
	Url    string                 `json:"url,omitempty"`    // 数据临时回调地址，指定后以此处指定为准，仅针对这一次推送请求生效；不指定，则以极光后台配置为准
	Params map[string]interface{} `json:"params,omitempty"` // 需要回调给用户的自定义参数
	Type   string                 `json:"type,omitempty"`   // 回调数据类型，1:送达回执, 2:点击回执, 3:送达和点击回执, 8:推送成功回执, 9:成功和送达回执, 10:成功和点击回执, 11:成功和送达以及点击回执
}

// SetUrl 设置回调地址
func (c *CallBack) SetUrl(url string) {
	c.Url = url
}

// SetParams 设置回调参数
func (c *CallBack) SetParams(params map[string]interface{}) {
	c.Params = params
}

// SetType 设置回调类型
func (c *CallBack) SetType(t string) {
	c.Type = t
}
