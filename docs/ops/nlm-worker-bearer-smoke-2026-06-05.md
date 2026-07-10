# NLM Worker bearer auth smoke

日期：2026-06-05
状态：候选施工 / 鉴权反向检查
位置：scripts + docs/ops

## 0. 一句话

本刀增加 `/mcp` bearer token 反向 smoke：Worker dev server 启动后，分别用无 token、错 token、正确 token 调用 `tools/list`，验收无 token / 错 token 必须 401，正确 token 才能 200。

## 1. 本刀新增

- `scripts/nlm-worker-bearer-smoke.mjs`
- `scripts/nlm-worker-wrangler-auth-smoke.sh`

## 2. 已有 Worker dev server 时

```bash
NLM_WORKER_SMOKE_URL="http://127.0.0.1:8787" \
MCP_BEARER_TOKEN="<local-token>" \
node scripts/nlm-worker-bearer-smoke.mjs
```

检查三种情况：

```plain text
missing_token → 401
wrong_token   → 401
good_token    → 200
```

## 3. 一键启动 Wrangler dev 并检查

```bash
MCP_BEARER_TOKEN="<local-token>" \
bash scripts/nlm-worker-wrangler-auth-smoke.sh
```

脚本会：

1. `node --check workers/nlm-worker-mcp/worker.js`；
2. `wrangler dev --local --ip 127.0.0.1 --port 8787`；
3. `GET /healthz`；
4. 执行 bearer auth smoke。

## 4. 成功条件

输出 JSON 的顶层：

```json
{
  "ok": true
}
```

且每个 case 为：

```json
[
  { "name": "missing_token", "status": 401 },
  { "name": "wrong_token", "status": 401 },
  { "name": "good_token", "status": 200 }
]
```

## 5. 边界

本刀不做：

- 不部署 Worker；
- 不配置 Cloudflare secret；
- 不配置 `NLM_RUNNER_URL`；
- 不调用 runner；
- 不调用 NotebookLM；
- 不读取 Google cookie；
- 不写回 manifest。

## 6. 下一步

若 bearer auth smoke 通过：

1. 再考虑 Cloudflare Worker 部署候选；
2. 部署后先只接 `/healthz` 与 `/mcp tools/list`；
3. 仍不接 runner；
4. 通过后再配置 `NLM_RUNNER_URL`。
