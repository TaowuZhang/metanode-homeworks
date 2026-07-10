# Publish workflow installation record

发布包契约、目标配置、脚本适配器与端到端 dry-run 样本已经落地。

2026-06-21 首次尝试写入专用 workflow 时，GitHub integration 因缺少 `workflows` 权限拒绝。2026-06-28 已由具备 workflow scope 的本地执行位补装：

- `.github/workflows/publish-dry-run.yml`
- `.github/workflows/publish-dispatch.yml`

## 当前职责

- `publish-dry-run.yml`：PR 或手动触发；测试 package/runner，校验发布包，生成 Web、Telegram、RSS 等草稿并上传 artifact。
- `publish-dispatch.yml`：只允许手动触发；模式显式选择，`live` 还需仓库或 environment variable `PUBLISH_LIVE_ENABLED=true`。
- 原 `Worang Index` workflow 中的 publish smoke 暂时保留为兼容质量门，等专用 workflow 稳定后再单独去重。

机器运行细节以 workflow 文件本身为准，本页不复制 YAML，避免两份配置漂移。

## 边界

- dry-run workflow 不读取 live secrets。
- dispatch workflow 默认 `dry_run`。
- `live` 需要显式输入和受控环境开闸；平台配置仍可再次拒绝。
- 当前适配器只渲染本地/Actions artifact；不把浏览器模拟或逆向 API 当作低风险 live 发布。
- 国内高风险平台保持草稿 / last-click。
