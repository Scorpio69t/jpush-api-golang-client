package jpush

import "encoding/json"

type PayLoad struct {
	Platform     *Platform     `json:"platform"`                // 平台
	Audience     *Audience     `json:"audience"`                // 推送目标
	Notification *Notification `json:"notification,omitempty"`  // 推送内容
	Message      *Message      `json:"message,omitempty"`       // 推送内容
	LiveActivity *LiveActivity `json:"live_activity,omitempty"` // 实时推送内容
	Options      *Options      `json:"options,omitempty"`       // 推送选项
	Cid          string        `json:"cid,omitempty"`           // 推送唯一标识符
}

// NewPayLoad 创建一个新的推送对象
func NewPayLoad() *PayLoad {
	p := &PayLoad{}
	p.Options = &Options{}
	p.Options.ApnsProduction = false
	return p
}

// SetPlatform 设置平台
func (p *PayLoad) SetPlatform(platform *Platform) {
	p.Platform = platform
}

// SetAudience 设置推送目标
func (p *PayLoad) SetAudience(audience *Audience) {
	p.Audience = audience
}

// SetNotification 设置推送内容
func (p *PayLoad) SetNotification(notification *Notification) {
	p.Notification = notification
}

// SetMessage 设置推送内容
func (p *PayLoad) SetMessage(message *Message) {
	p.Message = message
}

// SetOptions 设置推送选项
func (p *PayLoad) SetOptions(options *Options) {
	p.Options = options
}

// SetLiveActivity 设置实时推送内容
func (p *PayLoad) SetLiveActivity(liveActivity *LiveActivity) {
	p.LiveActivity = liveActivity
}

// Bytes 返回推送对象的json字节数组
func (p *PayLoad) Bytes() ([]byte, error) {
	payload := struct {
		Platform     interface{}   `json:"platform"`
		Audience     interface{}   `json:"audience"`
		Notification *Notification `json:"notification,omitempty"`
		Message      *Message      `json:"message,omitempty"`
		LiveActivity *LiveActivity `json:"live_activity,omitempty"`
		Options      *Options      `json:"options,omitempty"`
		Cid          string        `json:"cid,omitempty"`
	}{
		Platform:     p.Platform.Interface(),
		Audience:     p.Audience.Interface(),
		Notification: p.Notification,
		Message:      p.Message,
		LiveActivity: p.LiveActivity,
		Options:      p.Options,
		Cid:          p.Cid,
	}
	return json.Marshal(payload)
}
