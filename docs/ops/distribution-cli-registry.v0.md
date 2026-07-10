# 发布外手 CLI 登记 v0

本登记只保留当前可执行路由。完整证据、stars 快照与选型理由见 `platform-cli-selection-2026-06-28.md`。

## 采用规则

1. 上游必须可固定版本或 commit。
2. 优先官方 API/OAuth；浏览器 uploader 不因此伪装成官方路径。
3. CLI 必须有明确参数、退出码和可审计输出。
4. 沃壤负责 plan/receipt，不复制上游协议实现。
5. 安装进入 `target/tools/`，凭证留在上游工具自己的安全位置。

## 当前登记

| 工具 | 固定版本 | 平台 | 状态 | 沃壤接法 |
|---|---|---|---|---|
| xurl | v1.1.1 | X | 试验 | `xurl-plan` → plan-id apply |
| Postiz CLI | 2.0.15 | X / LinkedIn / YouTube / TikTok | 试验 | `postiz-plan` → remote draft |
| social-auto-upload | commit `0d3f93e…` | 小红书 / 抖音 / B站 / 视频号 / YouTube | 试验 + live 闸 | `sau-plan`，不由通用 executor 自动执行 |
| Ghost Admin API | 待部署实例 | 自有站 / newsletter / paid membership | 评估 | 先保留 `ghost` target |

## 命令

```bash
scripts/publish/install-external-clis.sh
node scripts/publish/external-cli-doctor.mjs
```

机器登记：`config/external-publish-tools.v0.json`。
