package push

type Message struct {
	MsgContent  string                 `json:"msg_content"`            // 消息内容本身
	Title       string                 `json:"title,omitempty"`        // 消息标题
	ContentType string                 `json:"content_type,omitempty"` // 消息内容类型
	Extras      map[string]interface{} `json:"extras,omitempty"`       // JSON 格式的可选参数
}

func (m *Message) SetContent(content string) {
	m.MsgContent = content
}

func (m *Message) SetTitle(title string) {
	m.Title = title
}

func (m *Message) SetContentType(contentType string) {
	m.ContentType = contentType
}

func (m *Message) SetExtras(extras map[string]interface{}) {
	m.Extras = extras
}

func (m *Message) AddExtras(key string, value interface{}) {
	if m.Extras == nil {
		m.Extras = make(map[string]interface{})
	}
	m.Extras[key] = value
}
