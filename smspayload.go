package jpush

import "encoding/json"

type SmsPayLoad struct {
	Mobile   string   `json:"mobile"`              // 手机号
	Signid   int      `json:"signid,omitempty"`    // 签名ID，该字段为空则使用应用默认签名。
	TempId   int64    `json:"temp_id,omitempty"`   // 短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara TempPara `json:"temp_para,omitempty"` // 短信模板中的参数。

}
type TempPara struct {
	Code json.RawMessage `json:"code,string"`
}

// NewPayLoad 创建一个新的推送对象
func NewSmsPayLoad() *SmsPayLoad {
	p := &SmsPayLoad{}
	return p
}

// SetPlatform 设置平台
func (p *SmsPayLoad) setMobile(mobile *string) {
	p.Mobile = *mobile
}
func (p *SmsPayLoad) SetSignid(signid *int) {
	p.Signid = *signid
}
func (p *SmsPayLoad) SetTempId(tempId *int64) {
	p.TempId = *tempId
}
func (p *SmsPayLoad) SetTempara(tempara *TempPara) {
	p.TempPara = *tempara
}

// Bytes 返回推送对象的json字节数组
func (p *SmsPayLoad) Bytes() ([]byte, error) {
	payload := struct {
		Mobile   string   `json:"mobile"`              // 手机号
		Signid   int      `json:"signid,omitempty"`    // 签名ID，该字段为空则使用应用默认签名。
		TempId   int64    `json:"temp_id,omitempty"`   // 短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
		TempPara TempPara `json:"temp_para,omitempty"` // 短信模板中的参数。
	}{

		Mobile:   p.Mobile,
		Signid:   p.Signid,
		TempId:   p.TempId,
		TempPara: p.TempPara,
	}
	return json.Marshal(payload)
}
