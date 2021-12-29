package jpushclient

type Options struct {
	SendNo            int           `json:"sendno,omitempty"`
	TimeToLive        int           `json:"time_to_live,omitempty"`
	OverrideMsgId     int           `json:"override_msg_id,omitempty"`
	ApnsProduction    bool          `json:"apns_production,omitempty"`
	BigPushDuration   int           `json:"big_push_duration,omitempty"`
	ThirdPartyChannel []interface{} `json:"thirdparty_channel,omitempty"`
}

// SetSendNo 设置消息的发送编号，用来覆盖推送时由 JPush 生成的编号。
func (o *Options) SetSendNo(sendNo int) {
	o.SendNo = sendNo
}

// SetTimeToLive 设置消息的有效期，单位为秒。
func (o *Options) SetTimeToLive(timeToLive int) {
	o.TimeToLive = timeToLive
}

// SetOverrideMsgId 设置覆盖推送时由 JPush 生成的消息 ID。
func (o *Options) SetOverrideMsgId(overrideMsgId int) {
	o.OverrideMsgId = overrideMsgId
}

// SetApnsProduction 设置推送时 APNs 是否生产环境。
func (o *Options) SetApnsProduction(apnsProduction bool) {
	o.ApnsProduction = apnsProduction
}

// SetBigPushDuration 设置大推送时长，单位为秒。
func (o *Options) SetBigPushDuration(bigPushDuration int) {
	o.BigPushDuration = bigPushDuration
}

// SetThirdPartyChannel 设置第三方渠道。
func (o *Options) SetThirdPartyChannel(thirdPartyChannel []interface{}) {
	o.ThirdPartyChannel = thirdPartyChannel
}
