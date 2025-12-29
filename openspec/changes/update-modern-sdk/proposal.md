## Why
- 现有 SDK 代码结构与 API 设计偏旧，HTTP 封装分散，缺少上下文支持和可注入传输层，错误/响应未结构化，不利于维护与集成。
- 需要在不新增业务功能的前提下，现代化代码结构、API 体验与测试体系，提升可读性和可扩展性，并为后续特性留出演进空间。

## What Changes
- 引入统一的 `Client` 与 Option 构造，提供 context-first 方法（Push/Schedule/CID/SMS），支持可注入 `http.Client/RoundTripper` 与每请求超时。
- 统一 HTTP do 层，集中认证、JSON 编解码、User-Agent/版本头，保留默认 60s 超时，可自定义。
- 定义结构化响应与错误模型，区分校验/网络/API/解码错误；旧接口保留为薄包装并标记 Deprecated。
- 增强 Payload 构建器：默认值（如 APNs）、发送前校验（platform/audience/SMS 模板参数等），确保 JSON 线格式兼容。
- 编写单元测试（mock RoundTripper）、校验测试、集成测试占位（tag integration），并用 golden JSON 保证兼容；更新示例与文档（中英）、新增迁移说明与 CHANGELOG。

## Impact
- 受影响的 specs：核心 jpush SDK 能力（Push/Schedule/Report/CID/SMS），当前无正式 specs，后续若有规格文件需同步更新。
- 受影响代码：`client/http 封装`、`payload/*`、`response/error 模型`、`示例/文档/测试`。
- 兼容性：默认保持 JSON 线格式和旧方法签名（作为 Deprecated 包装）；未来大版本可移除旧接口。
