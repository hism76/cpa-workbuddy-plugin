# CPA WorkBuddy / CodeBuddy Plugin

CLIProxyAPI (CPA) 的 Tencent WorkBuddy / CodeBuddy 原生 Provider 插件，支持国内版与国际版 (codebuddy.ai)。

## 功能特性

- 双区域支持: 支持国内版 (copilot.tencent.com / codebuddy.cn) 与国际版 (codebuddy.ai)，根据账号域自动分流网关与请求。
- OAuth 2.0 认证: 支持国内版微信/手机登录，以及国际版 Google/GitHub 登录。
- JWT 自动识别: 导入 Token 时自动解析提取 UID、用户昵称、有效期和所属区域，免手动填写。
- 动态模型发现: 按账号权限拉取可用模型列表，国际版原生支持 gpt-5 系列、gemini-3.5-flash、kimi-k3、deep-model 等。
- 标准协议转换: 兼容 OpenAI `/v1/chat/completions` 与 Claude `/v1/messages` 接口格式。
- 流式与工具调用: 完整支持 SSE 流式传输、OpenAI Function Calling (tools / tool_choice) 以及 reasoning_effort 思考链透传。
- 账号生命周期管理: 支持自动签到、额度监控与可视化管理面板。

## 构建指南

环境要求: Go 1.23+ 与 CGO 编译环境。

### 1. 构建国内版 (WorkBuddy)

```bash
CGO_ENABLED=1 go build -trimpath -buildmode=c-shared -o workbuddy.so .
```

### 2. 构建国际版 (CodeBuddy)

```bash
CGO_ENABLED=1 go build -trimpath -buildmode=c-shared -ldflags "-X main.providerName=codebuddy -X main.defaultRegion=global" -o codebuddy.so .
```

两个插件可以同时放入 CPA 的插件目录，提供独立的国内与国际版登录入口。

## 配置与部署

将生成的 `.so` 文件放置在 CPA 的插件目录中（如 `data/cpa/plugins/`），并在 CPA 的 `config.yaml` 中启用:

```yaml
plugins:
  enabled: true
  dir: "/app/data/plugins"
  configs:
    workbuddy:
      enabled: true
      checkin_auto: true
      lifecycle_auto: true
      token_keepalive: true
    codebuddy:
      enabled: true
      extra_models:
        - deepseek-v4.1-flash
        - hy4-preview
        - hy3
```

### 网络代理配置建议

国际版接口访问如需经过代理，请在环境或 docker-compose 中配置，但注意对腾讯边缘域名设置直连或按需分流:

```yaml
environment:
  HTTP_PROXY: "http://127.0.0.1:7890"
  HTTPS_PROXY: "http://127.0.0.1:7890"
  NO_PROXY: "127.0.0.1,localhost,codebuddy.ai,.codebuddy.ai,copilot.tencent.com,codebuddy.cn,.codebuddy.cn"
```

## 账号添加方式

1. 网页授权登录: 在 CPA 管理面板点击对应 Provider（WorkBuddy 或 CodeBuddy）即可唤起官方授权，登录成功后自动捕获并落盘。
2. 凭证导入: 支持直接粘贴官方 CLI (`@tencent-ai/codebuddy-code`) 或网页控制台中的 Access Token，系统将自动识别并补齐配置。

## 注意事项

- 国际版账号首次在浏览器端授权时，请确保客户端处于海外网络环境并建议选用 Google / GitHub 授权，避免因境内 IP 访问触发服务商安全策略限制。
