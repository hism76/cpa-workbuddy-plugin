# CPA WorkBuddy Plugin

CLIProxyAPI (CPA) 的 Tencent WorkBuddy / CodeBuddy 原生 Provider 动态插件。

## 功能特性

- OAuth 2.0 认证与会话管理：支持多账号登录与 Token 自动保活刷新。
- 动态模型发现：按账号权限实时拉取可用模型，支持客户端模型别名映射。
- 标准协议转换：将 OpenAI `/v1/chat/completions` 与 Claude `/v1/messages` 请求转译为上游协议。
- 流式与非流式兼容：上游强制 SSE 响应，插件支持非流式自动折叠聚合。
- Agent 与工具调用：完整支持 OpenAI 标准 Function Calling (tools / tool_choice) 流式增量合并。
- 思考链透传：支持 reasoning_effort 强度控制与 reasoning_content 思考流输出。
- 生命周期管理：额度耗尽自动挂起，支持定时自动签到回血与状态恢复。
- 内置管理面板：提供可视化积分监控与凭证管理界面。

## 安装与配置

### 1. 安装插件

下载对应平台的动态库文件（如 `workbuddy.so`）并放置到 CPA 插件目录：

```bash
mkdir -p /path/to/cliproxyapi/plugins
cp workbuddy.so /path/to/cliproxyapi/plugins/
```

### 2. 启用配置

在 CPA 的 `config.yaml` 中添加：

```yaml
plugins:
  enabled: true
  dir: "plugins"
  configs:
    workbuddy:
      enabled: true
      checkin_auto: true     # 开启每日自动签到
      lifecycle_auto: true   # 开启额度耗尽自动管理
      token_keepalive: true  # 开启 Token 自动保活
```

## 构建

要求 Go 1.23+ 与 CGO 编译环境：

```bash
CGO_ENABLED=1 go build -trimpath -buildmode=c-shared -o workbuddy.so .
```
