package main

import (
	"fmt"

	"github.com/Scorpio69t/jpush-api-golang-client"
)

func main() {
	// Platform: all
	var pf jpush.Platform
	pf.Add(jpush.ANDROID)
	pf.Add(jpush.IOS)
	pf.Add(jpush.WINPHONE)
	// pf.All()

	// Audience: tag
	var at jpush.Audience
	s := []string{"tag1", "tag2"}
	at.SetTag(s)
	id := []string{"1", "2"}
	at.SetID(id)
	// at.All()

	// Notification
	var n jpush.Notification
	n.SetAlert("alert")
	n.SetAndroid(&jpush.AndroidNotification{Alert: "alert", Title: "title"})
	n.SetIos(&jpush.IosNotification{Alert: "alert", Badge: 1})
	n.SetWinPhone(&jpush.WinPhoneNotification{Alert: "alert"})

	// Message
	var m jpush.Message
	m.MsgContent = "This is a message"
	m.Title = "Hello"

	// PayLoad
	payload := jpush.NewPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&at)
	payload.SetNotification(&n)
	payload.SetMessage(&m)

	// Send
	c := jpush.NewJPushClient("appKey", "masterSecret") // appKey and masterSecret can be gotten from https://www.jiguang.cn/
	data, err := payload.Bytes()
	fmt.Printf("%s\n", string(data))
	if err != nil {
		panic(err)
	}
	res, err := c.Push(data)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("ok: %v\n", res)
	}
}
