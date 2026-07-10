# NLM Worker MCP HTTP smoke

日期：2026-06-04
状态：候选施工 / MCP HTTP 本地检查
位置：scripts + docs/ops

## 0. 一句话

本刀增加 `/mcp` 的本地 HTTP smoke：在 Worker dev server 启动后，按 MCP JSON-RPC 调用 `initialize`、`tools/list`、manifest 工具，并确认没有 runner 时 `nlm_query_grains` 安全返回 `runner_not_configured`。

## 1. 本刀新增

- `scripts/nlm-worker-mcp-http-smoke.mjs`
- `scripts/nlm-worker-wrangler-mcp-smoke.sh`

## 2. 两种用法

### 2.1 已有 Worker dev server

如果 `wrangler dev` 已在本地运行：

```bash
NLM_WORKER_SMOKE_URL="http://127.0.0.1:8787" \
MCP_BEARER_TOKEN="<local-token>" \
node scripts/nlm-worker-mcp-http-smoke.mjs
```

它会调用：

```plain text
POST /mcp initialize
POST /mcp tools/list
POST /mcp nlm_manifest_list_packs
POST /mcp nlm_manifest_get_pack
POST /mcp nlm_query_grains
```

### 2.2 一键启动 Wrangler dev 并 smoke

在有 Wrangler 的受控环境中：

```bash
MCP_BEARER_TOKEN="<local-token>" \
bash scripts/nlm-worker-wrangler-mcp-smoke.sh
```

脚本会：

1. `node --check workers/nlm-worker-mcp/worker.js`；
2. `wrangler dev --local --ip 127.0.0.1 --port 8787`；
3. `GET /healthz`；
4. 调用 `scripts/nlm-worker-mcp-http-smoke.mjs`。

## 3. 预期结果

成功输出应显示：

- `initialize` ok；
- `tools/list` 返回工具名；
- `manifest/list` 返回 pack 数量；
- `manifest/get` 返回指定 pack 的绑定摘要；
- `query/grains-no-runner` 返回 `runner_not_configured`。

这里的 `runner_not_configured` 是预期成功条件之一：说明 Worker 能走到查询门口，但没有 runner 时不会误触 NotebookLM。

## 4. 可配置变量

```plain text
NLM_WORKER_SMOKE_URL=http://127.0.0.1:8787
NLM_SMOKE_PACK_KEY=wo/nlm/becoming-projects-overview.md
MCP_BEARER_TOKEN=<local-token>
NLM_WORKER_DIR=workers/nlm-worker-mcp
NLM_WRANGLER_HOST=127.0.0.1
NLM_WRANGLER_PORT=8787
```

## 5. 边界

本刀不做：

- 不 `wrangler deploy`；
- 不配置 Cloudflare secret；
- 不配置 `NLM_RUNNER_URL`；
- 不调用 runner；
- 不调用 NotebookLM；
- 不读取 Google cookie；
- 不写回 manifest。

## 6. 下一步

若 MCP HTTP smoke 通过：

1. 做 bearer token 反向检查：无 token / 错 token 必须 401；
2. 再考虑 Cloudflare Worker 部署候选；
3. 部署后仍先只连 manifest 工具，不接 runner；
4. 最后再配置 `NLM_RUNNER_URL`。
