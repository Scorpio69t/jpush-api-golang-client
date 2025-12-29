# Changelog

## Unreleased
- 开始现代化重构：新增 context-first `Client`，可注入 http.Client/Transport，结构化响应与错误模型，发送前校验。
- 旧接口（如 `Push([]byte)`）保持兼容，计划后续版本弃用，推荐迁移到新 API。
