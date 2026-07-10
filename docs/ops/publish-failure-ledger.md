# 发布失败台账

用途：记录发布管道的失败、风控、限流、凭证、平台适配断裂。它不保存 token，不保存 cookie，不保存完整敏感响应。

## 记录格式

| 时间 | event_id | 平台 | 模式 | 阶段 | 失败类型 | 摘要 | 下一步 |
|---|---|---|---|---|---|---|---|
| 2026-06-21T18:28:00+08:00 | prism-entry-v1 | simulated-broken | dry_run | adapter | simulated | v0 烟测模拟失败，确认失败台账入口可用 | 保留为样例 |

## 失败类型

| 类型 | 含义 | 默认动作 |
|---|---|---|
| `schema` | 发布包字段不合格 | 回 prepare / validate |
| `risk` | 隐私、版权、署名、公共后果未过 | 回临发 / 人工判断 |
| `credential` | 缺 token、权限、频道 ID | 不在 Notion 补凭证，去安全介质配置 |
| `rate_limit` | 平台限流 | 延后，不重试刷屏 |
| `platform_changed` | UI / API 变更 | 降级 last-click，登记适配缺口 |
| `policy_risk` | 可能触发平台警告 / 风控 | 停 live，只保留草稿 |
| `network` | 网络、代理、DNS | 进入可达性检查 |
| `simulated` | 烟测模拟失败 | 只验证台账链路 |

## 禁止记录

- token / cookie / session。
- 私密原文。
- 未脱敏个人信息。
- 平台完整敏感响应体。
