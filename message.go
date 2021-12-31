package jpush

type Message struct {
	MsgContent  string                 `json:"msg_content"`            // 消息内容本身
	Title       string                 `json:"title,omitempty"`        // 消息标题
	ContentType string                 `json:"content_type,omitempty"` // 消息内容类型
	Extras      map[string]interface{} `json:"extras,omitempty"`       // JSON 格式的可选参数
}

// inapp_message 面向于通知栏消息类型，对于通知权限关闭的用户可设置启用此功能。此功能启用后，当用户前台运行APP时，会通过应用内消息的方式展示通知栏消息内容。
type InappMessage struct {
	InAppMessage bool `json:"inapp_message"`
}

type SmsMessage struct {
	DelayTime    int         `json:"delay_time"`          // 单位为秒，不能超过24小时。设置为0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
	Signid       int         `json:"signid,omitempty"`    // 签名ID，该字段为空则使用应用默认签名。
	TempId       int64       `json:"temp_id,omitempty"`   // 短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara     interface{} `json:"temp_para,omitempty"` // 短信模板中的参数。
	ActiveFilter bool        `json:"active_filter"`       // active_filter 字段用来控制是否对补发短信的用户进行活跃过滤，默认为 true ，做活跃过滤；为 false，则不做活跃过滤；
}

// SetContent 设置消息内容
func (m *Message) SetContent(content string) {
	m.MsgContent = content
}

// SetTitle 设置消息标题
func (m *Message) SetTitle(title string) {
	m.Title = title
}

// SetContentType 设置消息内容类型
func (m *Message) SetContentType(contentType string) {
	m.ContentType = contentType
}

// SetExtras 设置消息扩展内容
func (m *Message) SetExtras(extras map[string]interface{}) {
	m.Extras = extras
}

// AddExtras 添加消息扩展内容
func (m *Message) AddExtras(key string, value interface{}) {
	if m.Extras == nil {
		m.Extras = make(map[string]interface{})
	}
	m.Extras[key] = value
}

// SetInAppMessage 设置是否启用应用内消息
func (i *InappMessage) SetInAppMessage(inAppMessage bool) {
	i.InAppMessage = inAppMessage
}

// SetDelayTime 设置短信发送延时时间
func (s *SmsMessage) SetDelayTime(delayTime int) {
	s.DelayTime = delayTime
}

// SetSignid 设置短信签名ID
func (s *SmsMessage) SetSignid(signid int) {
	s.Signid = signid
}

// SetTempId 设置短信模板ID
func (s *SmsMessage) SetTempId(tempId int64) {
	s.TempId = tempId
}

// SetTempPara 设置短信模板参数
func (s *SmsMessage) SetTempPara(tempPara interface{}) {
	s.TempPara = tempPara
}

// SetActiveFilter 设置是否对补发短信的用户进行活跃过滤
func (s *SmsMessage) SetActiveFilter(activeFilter bool) {
	s.ActiveFilter = activeFilter
}
