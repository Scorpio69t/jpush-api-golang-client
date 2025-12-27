# Project Context

## Purpose
JPush 推送/短信的 Go 语言服务器端 SDK，提供对 JPush REST API（推送、定时、报表、CID、短信等）的轻量封装，补齐官方示例缺失的能力，方便在后端服务中直接集成。

## Tech Stack
- Go 1.22 模块 `github.com/Scorpio69t/jpush-api-golang-client`
- 仅使用标准库（net/http、encoding/json、time/strconv 等），无三方依赖
- 自带简化的 HTTP 辅助封装（`httplib.go`、`httpclient.go`），用于构造带 Basic Auth 的请求
- 提供 `examples/` 示例演示推送与 CID 获取流程

## Project Conventions

### Code Style
- 使用 gofmt/idiomatic Go，导出类型/字段采用 PascalCase，常量使用 SNAKE_CASE
- 通过 builder/Setter 方式构建 payload（`Platform`、`Audience`、`Notification`、`Message`、`Options` 等），最终用 `PayLoad.Bytes()` 生成 JSON
- API 返回 `(string, error)` 或 `([]byte, error)`，调用处需显式处理 error；默认网络超时 60s（连接/读写）
- JSON 字段通过 struct tag 固定，尽量保持与 JPush REST 协议兼容，避免破坏已有字段命名

### Architecture Patterns
- 单包 `jpush` 暴露客户端 `JPushClient`，内部常量集中管理各 REST 入口（Push/Schedule/Report/CID/SMS）
- HTTP 适配层封装在 `httplib.go`/`httpclient.go`，提供基础的 GET/POST 方法与认证/头部设置
- 数据建模与序列化集中在独立文件：`platform.go`、`audience.go`、`notification*.go`、`message.go`、`options.go`、`payload.go`、`smspayload.go`
- 示例与用法文档位于 `README.md` 与 `examples/`

### Testing Strategy
- 当前仅有少量集成型测试示例（如 `sms_test.go` 需要真实 appKey/masterSecret），默认 `go test ./...` 会尝试真实请求
- 贡献新功能时优先添加可离线跑的单元测试；若依赖外部接口，优先通过接口抽象/HTTP mock 降低对真实服务的依赖
- 示例代码用于手动验证典型调用路径（推送、CID、短信）

### Git Workflow
- 默认 main 干线开发；新特性/修复建议走 feature 分支并通过 PR 合并
- 提交前保持 gofmt，通过 go test ./...（若需要外部凭证的测试请标注或跳过）
- 暂无强制的 commit message 规范，鼓励清晰的前缀（feat/fix/doc/test/chore）

## Domain Context
- 服务面向极光推送（JPush）生态：Push v3、Schedule v3、Report v3、CID v3、SMS v1；Device/File/Image/Admin 目前未实现
- 认证采用 appKey/masterSecret 的 Basic Auth；部分接口对 apns_production、cid、目标 audience 等字段敏感
- Payload 结构遵循 JPush REST 协议，需保持字段命名与类型与官方文档一致以避免推送/短信失败

## Important Constraints
- 保持与官方 REST API 的兼容性为最高优先级，避免随意修改已发布字段或默认值
- 不要在代码库中提交真实 appKey/masterSecret；集成测试需显式标注依赖外部资源
- 默认超时为 60s；若调整需评估对现有调用方的影响
- 目标是零三方依赖的轻量 SDK，引入新依赖需有充分理由（安全/性能/维护成本）

## External Dependencies
- 极光推送 REST 服务：`https://api.jpush.cn/v3/push`、`/v3/schedules`、`https://report.jpush.cn/v3/received`、`/v3/push/cid`
- 极光短信 REST 服务：`https://api.sms.jpush.cn/v1/messages`
- 其他 JPush 相关接口（Device/File/Image/Admin）暂未接入，未来扩展需遵循相同认证/JSON 规范
