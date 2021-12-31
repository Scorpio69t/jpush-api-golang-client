package jpush

import "errors"

type Notification3rd struct {
	Title       string                 `json:"title,omitempty"`         // 补发通知标题，如果为空则默认为应用名称
	Content     string                 `json:"content"`                 // 补发通知的内容，如果存在 notification_3rd 这个key，content 字段不能为空，且值不能为空字符串。
	ChannelId   string                 `json:"channel_id,omitempty"`    // 不超过1000字节
	UriActivity string                 `json:"uri_activity,omitempty"`  // 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值;适配华为、小米、vivo厂商通道跳转；针对 VIP 厂商通道用户使用生效。
	UriAction   string                 `json:"uri_action,omitempty"`    // 指定跳转页面；该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值;适配 oppo、fcm跳转；针对 VIP 厂商通道用户使用生效。
	BadgeAddNum string                 `json:"badge_add_num,omitempty"` // 角标数字，取值范围1-99
	BadgeClass  string                 `json:"badge_class,omitempty"`   // 桌面图标对应的应用入口Activity类， 比如“com.test.badge.MainActivity；
	Sound       string                 `json:"sound,omitempty"`         // 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀；注意：针对Android 8.0以上，当传递了channel_id 时，此属性不生效。
	Extras      map[string]interface{} `json:"extras,omitempty"`        // 扩展字段；这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
}

// SetTitle 设置标题
func (n *Notification3rd) SetTitle(title string) {
	n.Title = title
}

// SetContent 设置内容
func (n *Notification3rd) SetContent(content string) error {
	if len(content) == 0 {
		return errors.New("content is empty")
	}
	n.Content = content
	return nil
}

// SetChannelId 设置通道ID
func (n *Notification3rd) SetChannelId(channelId string) {
	n.ChannelId = channelId
}

// SetUriActivity 设置uri_activity
func (n *Notification3rd) SetUriActivity(uriActivity string) {
	n.UriActivity = uriActivity
}

// SetUriAction 设置uri_action
func (n *Notification3rd) SetUriAction(uriAction string) {
	n.UriAction = uriAction
}

// SetBadgeAddNum 设置角标数字
func (n *Notification3rd) SetBadgeAddNum(badgeAddNum string) {
	n.BadgeAddNum = badgeAddNum
}

// SetBadgeClass 设置桌面图标对应的应用入口Activity类
func (n *Notification3rd) SetBadgeClass(badgeClass string) {
	n.BadgeClass = badgeClass
}

// SetSound 设置铃声
func (n *Notification3rd) SetSound(sound string) {
	n.Sound = sound
}

// SetExtras 设置扩展字段
func (n *Notification3rd) SetExtras(extras map[string]interface{}) {
	n.Extras = extras
}
