# Steam 季度候选同步启用清单

## 当前状态

本页承接 PR #172。当前 PR 已经包含真实 GitHub Actions workflow 文件，但在 PR merge 前还不会进入 `main` 的常规季度节律。

已具备：

- `scripts/write-steam-quarterly-candidates.mjs`
- `.github/workflows/sync-steam-quarterly.yml`
- `docs/ops/steam-quarterly-candidate-sync-workflow-template.yml`

尚未完成：

- PR #172 merge
- GitHub Actions 首次 `workflow_dispatch` 试跑
- 首次自动生成的季度候选 Draft PR 复核

## 启用前提

GitHub repository 需要存在两个 Actions secrets：

- `STEAM_API_KEY`
- `STEAM_ID`

说明：

- 不要把真实 Steam API key 写进 Git。
- 不要把真实 Steam API key 粘到聊天里。
- 如果此前已经获取过 Steam API key，它可能在本地 `.env`、密码管理器、Steam dev 页面，或已经存在于 GitHub Actions secrets。
- 克当前不能读取 GitHub secrets 的真实值，只能按 workflow 运行结果判断 secrets 是否可用。

## 已完成的本地验证

谷已在 clean worktree `worang-steam-qsync` 完成本地端到端验证：

- 本地 `.env` 存在 `STEAM_API_KEY`。
- 本地 `.env` 存在 `STEAM_ID` / `STEAMID`。
- 未把真实 secret 粘贴到聊天或提交进 Git。
- `node --check` 通过：
  - `scripts/steam-probe.mjs`
  - `scripts/steam-review-queue.mjs`
  - `scripts/write-steam-game-facts-v0.mjs`
  - `scripts/write-steam-quarterly-candidates.mjs`
- Steam read-only probe 跑通：
  - owned games: 23
  - recent games: 0
  - candidate events: 39
  - achievements checked: 0
- Review queue 跑通：
  - games: 23
  - `existing_match_candidate`: 2
  - `new_entry_candidate`: 21
- Quarterly writer 跑通：
  - rows: 23
  - quarter: `2026-Q2`
- Generated latest CSV / MD 已本地检查后删除，未纳入本 PR。

## Workflow 文件状态

真实 workflow 已加入：

```plain text
.github/workflows/sync-steam-quarterly.yml
```

它使用 GitHub Actions secrets 引用：

```plain text
STEAM_API_KEY: GitHub Actions secret reference
STEAM_ID: GitHub Actions secret reference
```

本地已验证：

- 存在 GitHub Actions secrets 引用。
- 不存在 `__GITHUB_ACTIONS_SECRET_*` 模板占位。
- 不存在 broken `$ secrets.*` 伪语法。
- `git diff --check` 通过。

## 首次试跑

PR #172 merge 后，先用 `workflow_dispatch` 手动跑一次，不等季度 cron。

首次 workflow run 应该：

- 通过 `node --check`：
  - `scripts/steam-probe.mjs`
  - `scripts/steam-review-queue.mjs`
  - `scripts/write-steam-quarterly-candidates.mjs`
- 用 secrets 读取 Steam，只生成本地 ignored 输出：
  - `资源/_steam/steam_snapshot.json`
  - `资源/_steam/steam_candidate_events.csv`
  - `资源/_steam/steam_review_queue.csv`
- 只把候选摘要写到 tracked docs/ops：
  - `docs/ops/steam-game-resource-quarterly-candidates-latest.csv`
  - `docs/ops/steam-game-resource-quarterly-candidates-latest.md`
- 自动开 Draft PR：
  - `sync: update Steam quarterly candidates`

## 边界守卫

首次试跑和季度运行都必须保持：

- 不提交 `资源/_steam/`。
- 不修改 `资源/游戏.csv`。
- 不修改 `资源/游戏/`。
- 不处理 `wrangler.toml`。
- 不处理 `workers/getnote-mirror-mcp/.wrangler/`。
- 不自动 merge。
- 不把 Steam owned / playtime / achievements 自动等同于喜欢、重要、通关或沉积。
- 所有主表、正式旁表、价值与完成状态判断仍需谷确认。

## 回滚方式

若首次试跑出现异常：

1. 关闭或删除 `.github/workflows/sync-steam-quarterly.yml`。
2. 关闭自动生成的季度候选 Draft PR。
3. 保留 #172 中的脚本与 docs/ops 记录作为设计沉积。
4. 只在修复脚本 / secret / 权限后再重新启用。
