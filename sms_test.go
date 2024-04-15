package jpush

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {

	var mobile string
	mobile = "xxxx"
	var tempid int64
	tempid = 1
	const code = "141238"
	tempara := TempPara{Code: json.RawMessage(code)}
	fmt.Printf("mobile is %s\n", string(mobile))

	smspayload := NewSmsPayLoad()
	smspayload.setMobile(&mobile)
	smspayload.SetTempId(&tempid)
	smspayload.SetTempara(&tempara)
	// Send
	c := NewJPushClient("appKey", "masterSecret") // appKey and masterSecret can be gotten from https://www.jiguang.cn/
	data, err := smspayload.Bytes()
	fmt.Printf("sms is %s\n", string(data))
	if err != nil {
		panic(err)
	}
	res, err := c.SendSms(data)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("ok: %v\n", res)
	}
}
