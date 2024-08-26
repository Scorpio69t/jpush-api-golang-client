# jpush-api-golang-client



## 概述
JPush's Golang client library for accessing JPush APIs. 极光推送的 Golang 版本服务器端 SDK。
该项目参考[ylywyn](https://github.com/ylywyn/jpush-api-go-client)结合极光推送官方文档而来。(原项目年久失修，有很多新特性都没有提供，本项目旨在将其完善，方便大家使用，后续会持续更新，不足之处欢迎大家指正，谢谢~)

[参考REST API文档](https://docs.jiguang.cn/jpush/server/push/server_overview/)

### 极光短信

[参考短信 REST API 文档](https://docs.jiguang.cn/jsms/server/rest_api_summary)

短信这边暂时只实现了发送单条模板短信

**现已支持以下内容**

- [x] Push API v3
- [x] Report API v3
- [ ] Device API v3
- [x] Schedule API v3
- [ ] File API v3
- [ ] Image API v3
- [ ] Admin API v3
- [x] SMS API v1

## 使用
`go get github.com/Scorpio69t/jpush-api-golang-client`

## 推送流程



### 1.构建要推送的平台：jpush.Platform
```go
// Platform: all
var pf jpush.Platform
pf.Add(jpush.ANDROID)
pf.Add(jpush.IOS)
pf.Add(jpush.WINPHONE)
// pf.All()
```



### 2.构建接收目标：jpush.Audience

```go
// Audience: tag
var at jpush.Audience
s := []string{"tag1", "tag2"}
at.SetTag(s)
id := []string{"1", "2"}
at.SetID(id)
// at.All()
```



### 3.构建通知：jpush.Notification 或者消息：jpush.Message

```go
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
```



### 4.构建消息负载：jpush.PayLoad

```go
// PayLoad
payload := jpush.NewPayLoad()
payload.SetPlatform(&pf)
payload.SetAudience(&at)
payload.SetNotification(&n)
payload.SetMessage(&m)
```



### 5.构建JPushClient，发送推送

```go
// Send
c := jpush.NewJPushClient("appKey", "masterSecret") // appKey and masterSecret can be gotten from https://www.jiguang.cn/
data, err := payload.Bytes()
if err != nil {
	panic(err)
}
res, err := c.Push(data)
if err != nil {
	fmt.Printf("%+v\n", err)
} else {
	fmt.Printf("ok: %v\n", res)
}
```

### 6.详细例子见examples

## 发送短信参见sms_test中的代码

