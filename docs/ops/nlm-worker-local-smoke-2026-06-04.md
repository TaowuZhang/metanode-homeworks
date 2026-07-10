# NLM Worker local smoke

日期：2026-06-04
状态：候选施工 / Worker 本地 smoke
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-worker-local-smoke.mjs` 是部署前的 Worker MCP 本地检查：直接 import `workers/nlm-worker-mcp/worker.js`，用本地 `wo/nlm/manifest.json` 和本地 pack 文件模拟 GitHub contents API，验证 MCP facade 能启动、列工具、读 manifest、查 pack、判 freshness，并在没有 runner 时安全返回 `runner_not_configured`。

它不部署 Worker，不访问 GitHub，不调用 runner，不调用 NotebookLM。

## 1. 本刀新增

- `scripts/nlm-worker-local-smoke.mjs`

## 2. 验证内容

脚本依次调用 Worker：

```plain text
initialize
tools/list
nlm_manifest_list_packs
nlm_manifest_get_pack
nlm_check_pack_freshness
nlm_query_grains
```

其中 `nlm_query_grains` 在没有 `NLM_RUNNER_URL` 时，预期返回：

```json
{
  "ok": false,
  "error": "runner_not_configured"
}
```

这表示 Worker 能识别 pack 绑定，但没有 runner 时不会误触 NotebookLM。

## 3. 使用方式

默认 pack：

```plain text
wo/nlm/becoming-projects-overview.md
```

执行：

```bash
node scripts/nlm-worker-local-smoke.mjs
```

指定 pack：

```bash
NLM_SMOKE_PACK_KEY="wo/nlm/becoming-coordinates-index.md" \
node scripts/nlm-worker-local-smoke.mjs
```

如果 pack 使用 `upload.notebook_id_env = NLM_BECOMING`，脚本默认给一个本地 placeholder，只用于确认绑定路径，不是真实 notebook id。

## 4. 输出

输出 JSON 包含：

- manifest hash；
- 每个 MCP 调用的状态；
- tools 名称；
- pack 绑定摘要；
- freshness 结果；
- query 无 runner 时的关门失败。

## 5. 边界

本脚本不做：

- 不请求 GitHub API；
- 不启动 Cloudflare Worker；
- 不使用 wrangler；
- 不调用 runner；
- 不调用 NotebookLM；
- 不读取或写入 Google cookie；
- 不写回 manifest。

## 6. 下一步

若本地 smoke 通过，再进入：

1. Wrangler / Miniflare 启动 Worker；
2. `/healthz` 检查；
3. `/mcp` tools/list；
4. 仍不接 runner；
5. 最后再配置 `NLM_RUNNER_URL`。
