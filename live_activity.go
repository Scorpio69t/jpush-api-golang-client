package jpush

type LiveActivity struct {
	Ios *IosLiveActivity `json:"ios,omitempty"`
}

type IosLiveActivity struct {
	Event          string                `json:"event"`                     // 开始：“start”，更新：“update”，结束："end"。
	ContentState   interface{}           `json:"content-state"`             // 需与客户端 SDK 值匹配（对应 Apple 官方的 content-state 字段）。
	AttributesType string                `json:"attributes-type"`           // 创建实时活动事件必填（更新或者结束实时活动无需传递），字段规则：由数字，字母，下划线组成，但是不能以数字开头；对应 Apple 官方的 attributes-type 字段
	Attributes     interface{}           `json:"attributes"`                // 创建实时活动事件必填（更新或者结束实时活动无需传递）；对应 Apple 官方的 attributes 字段
	RelevanceScore int                   `json:"relevance-score,omitempty"` // 对应 Apple 官方的 relevance-score 字段
	StaleDate      int                   `json:"stale-date,omitempty"`      // 对应 Apple 官方的 stale-date 字段
	Alert          *IosLiveActivityAlert `json:"alert,omitempty"`           // 通知内容
	DismissalDate  int                   `json:"dismissal-date,omitempty"`  // 实时活动结束展示时间。
}

type IosLiveActivityAlert struct {
	Title string `json:"title,omitempty"` // 标题
	Body  string `json:"body,omitempty"`  // 内容
	Sound string `json:"sound,omitempty"` // 声音
}
