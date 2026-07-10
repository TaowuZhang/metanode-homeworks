# NLM Worker deploy preflight

日期：2026-06-05
状态：候选施工 / 部署前检查清单
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-worker-deploy-preflight.sh` 是 Cloudflare Worker 真部署前的最后一道静态检查：确认文件、脚本、语法、必需变量、必需 secret、禁止项都显影；它不执行 `wrangler deploy`。

## 1. 本刀新增

- `scripts/nlm-worker-deploy-preflight.sh`

## 2. 使用方式

```bash
bash scripts/nlm-worker-deploy-preflight.sh
```

脚本会检查：

- `workers/nlm-worker-mcp/worker.js`
- `workers/nlm-worker-mcp/README.md`
- `workers/nlm-worker-mcp/wrangler.toml.example`
- Worker local smoke 脚本
- manifest dry-run 脚本
- Wrangler smoke 脚本
- MCP HTTP smoke 脚本
- bearer auth smoke 脚本
- `node --check workers/nlm-worker-mcp/worker.js`
- `node` / `wrangler` / `curl` 是否可用

## 3. 第一阶段 Cloudflare 变量

首次部署只允许设置 manifest / facade 所需变量：

```plain text
GITHUB_REPO=dongxi-heji/worang
GITHUB_REF=main
NLM_MANIFEST_PATH=wo/nlm/manifest.json
```

如果要让现有「沃壤｜成为」pack 能显示 notebook binding，需要设置：

```plain text
NLM_BECOMING=<actual notebook id>
```

注意：`NLM_BECOMING` 是 Worker 环境变量，不写入 GitHub。

## 4. 第一阶段 Cloudflare secrets

必需：

```plain text
MCP_BEARER_TOKEN
```

可选：

```plain text
GITHUB_TOKEN
```

`GITHUB_TOKEN` 只应是 fine-grained、Contents read-only，用于提高 GitHub contents API 稳定性；公开 repo 可先不设。

暂不设置：

```plain text
NLM_RUNNER_URL
NLM_RUNNER_TOKEN
```

因为第一阶段只部署 Worker facade，不接 runner。

## 5. 首次部署后只验收这些

允许验收：

```plain text
GET /healthz
POST /mcp initialize
POST /mcp tools/list
nlm_manifest_list_packs
nlm_manifest_get_pack
nlm_check_pack_freshness
```

`nlm_query_grains` 在未配置 runner 时应返回：

```json
{
  "ok": false,
  "error": "runner_not_configured"
}
```

这不是失败，而是第一阶段的安全成功条件。

## 6. 部署前禁止项

第一阶段禁止：

- 不设置 `NLM_RUNNER_URL`；
- 不设置 NotebookLM cookie；
- 不在 Worker 保存 Google 登录态；
- 不开放 source add / sync / delete / share / deep research；
- 不让 Worker 写 GitHub；
- 不把 NotebookLM 回答写回沃壤；
- 不把外部 MCP 原始工具面暴露给克。

## 7. 进入部署的条件

只有全部满足才进入 Cloudflare 部署候选：

1. `scripts/nlm-worker-local-smoke.mjs` 可在完整仓库环境通过；
2. `scripts/nlm-worker-manifest-dry-run.mjs list` 可读 manifest；
3. `scripts/nlm-worker-wrangler-smoke.sh` 在有 Wrangler 环境下 `/healthz` 通过；
4. `scripts/nlm-worker-wrangler-mcp-smoke.sh` 的 `/mcp` smoke 通过；
5. `scripts/nlm-worker-wrangler-auth-smoke.sh` 的无 token / 错 token / 正 token 检查通过；
6. `scripts/nlm-worker-deploy-preflight.sh` 返回 `preflight_ok`。

## 8. 下一步

如果 preflight 通过，下一刀才是：

```bash
cd workers/nlm-worker-mcp
cp wrangler.toml.example wrangler.toml
wrangler secret put MCP_BEARER_TOKEN
wrangler deploy
```

部署后仍先只接 manifest tools，不接 runner。
