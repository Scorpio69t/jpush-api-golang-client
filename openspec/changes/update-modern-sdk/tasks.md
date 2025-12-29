## 1. Implementation
- [ ] 1.1 新增 `Client` + Option 构造，提供 context 优先的 Push/Schedule/CID/SMS 方法，并保留旧方法包装（标记 Deprecated）
- [ ] 1.2 抽象统一 HTTP do 层（认证/超时/JSON 编解码/User-Agent），支持注入 http.Client/RoundTripper
- [ ] 1.3 增强 payload 构建与校验：默认值（如 APNs）、platform/audience/SMS 参数校验，确保 JSON 兼容（golden 测试）
- [ ] 1.4 定义结构化响应与错误模型（区分校验/网络/API/解码），更新调用返回值，旧接口转发
- [ ] 1.5 单元测试：mock RoundTripper 覆盖 push/schedule/report/CID/SMS 成功与错误路径；校验测试
- [ ] 1.6 集成测试占位（`-tags=integration`），需真实凭证；默认跳过
- [ ] 1.7 更新示例（context API）、README 中英、迁移说明与 CHANGELOG；确保 gofmt & go test ./...
