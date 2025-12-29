# jpush-api-golang-client

> 提示：当前正在进行重构以现代化 API 和结构，近期可能有变更，请关注版本与变更日志。

英文为默认说明，中文说明在此文件。如需英文请查看 `README.md`。

## 概述
极光推送 JPush REST API 的 Go 语言 SDK（推送、定时、报表、CID、短信）。现有代码基于官方文档并提供少量 HTTP 辅助封装与 payload 构建器。

当前支持：
- ✅ Push API v3
- ✅ Report API v3
- ✅ Schedule API v3
- ✅ SMS API v1（模板短信发送）
- ⏳ 尚未实现：Device API v3、File API v3、Image API v3、Admin API v3

## 安装
```bash
go get github.com/Scorpio69t/jpush-api-golang-client
```

## 推送快速开始
1) 平台
```go
var pf jpush.Platform
pf.Add(jpush.ANDROID)
pf.Add(jpush.IOS)
pf.Add(jpush.WINPHONE)
// pf.All()
```

2) 接收目标
```go
var at jpush.Audience
at.SetTag([]string{"tag1", "tag2"})
at.SetID([]string{"1", "2"})
// at.All()
```

3) 通知或自定义消息
```go
var n jpush.Notification
n.SetAlert("alert")
n.SetAndroid(&jpush.AndroidNotification{Alert: "alert", Title: "title"})
n.SetIos(&jpush.IosNotification{Alert: "alert", Badge: 1})
n.SetWinPhone(&jpush.WinPhoneNotification{Alert: "alert"})

var m jpush.Message
m.MsgContent = "This is a message"
m.Title = "Hello"
```

4) 负载
```go
payload := jpush.NewPayLoad()
payload.SetPlatform(&pf)
payload.SetAudience(&at)
payload.SetNotification(&n)
payload.SetMessage(&m)
```

5) 发送
```go
c := jpush.NewJPushClient("appKey", "masterSecret") // 前往 https://www.jiguang.cn/ 获取
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

完整推送与 CID 样例见 `examples/`。

## 现代化用法（推荐，支持 context）
```go
ctx := context.Background()
c, err := jpush.NewClient("appKey", "masterSecret")
if err != nil {
    panic(err)
}

// 构建 payload 同上
resp, err := c.Push(ctx, payload)
if err != nil {
    panic(err)
}
fmt.Println(resp.MsgID)
```

旧接口（如 `Push([]byte)`) 仍可用但已进入弃用通道，请尽快迁移到 `Client`/context API。

## 迁移指引（从旧接口）
- 将 `NewJPushClient` + `Push([]byte)` 替换为 `NewClient` + `Push(ctx, *PayLoad)`。
- 如果需要链路追踪/Mock，可通过 `WithHTTPClient` 注入自定义 `http.Client`，超时用 `WithTimeout` 配置。
- 使用结构化响应（如 `PushResponse`、`SMSResponse`），不再依赖字符串解析。

## 短信
已支持模板短信发送，示例见 `sms_test.go`。运行前请填写自己的 `appKey`/`masterSecret` 和模板参数。
