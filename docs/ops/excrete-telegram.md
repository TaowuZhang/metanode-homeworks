# Telegram 排泄管

Telegram 当前是发布包 payload 与受控通道候选，不再由 `scripts/wo-excrete.mjs` 默认直接发送。

## 安全 dry-run

```bash
wo excrete last --platforms telegram --mode dry_run --prepare --json
```

生成物：

- `ai-browse/out/packages/<event-id>.json`
- `ai-browse/out/packages/<event-id>.md`
- `ai-browse/out/publish-runs/<event-id>/telegram-payload.json`
- `ai-browse/out/publish-runs/<event-id>/summary.json`

这一步不读取 token，不访问 Telegram，也不表示已经发布。

## Live 边界

- `publish_mode=live` 仍会先经过 `config/publish-targets.v0.json` 的 `live_allowed` 闸门；Telegram 当前为 `false`。
- 真发送接回前，必须固定 Bot API adapter、timeout、明确 chat 白名单、receipt 与幂等策略。
- token / chat id 只能放系统 Keychain、环境 secret 或 GitHub environment secret，不进 Notion、Git、日志或发布包。
- 旧 `scripts/lib/pipes/telegram.mjs` 是历史真通道实现，已从默认 runner 断开；不得绕过发布包与人工检查直接重新挂回。
