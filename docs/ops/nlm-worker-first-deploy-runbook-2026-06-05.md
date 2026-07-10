# NLM Worker first deploy runbook

日期：2026-06-05
状态：候选施工 / 首次部署手册
位置：docs/ops

## 0. 一句话

首次部署只把 NLM Worker MCP facade 放上 Cloudflare，验收 `/healthz`、`/mcp initialize`、`tools/list` 与 manifest 只读工具；不配置 runner，不碰 NotebookLM，不写任何 Google cookie。

## 1. 前置条件

必须先在仓库根目录通过或人工确认：

```bash
bash scripts/nlm-worker-deploy-preflight.sh
```

并确保以下脚本在受控环境里可用：

```plain text
scripts/nlm-worker-local-smoke.mjs
scripts/nlm-worker-manifest-dry-run.mjs
scripts/nlm-worker-wrangler-smoke.sh
scripts/nlm-worker-mcp-http-smoke.mjs
scripts/nlm-worker-bearer-smoke.mjs
scripts/nlm-worker-deploy-preflight.sh
```

## 2. 首次部署范围

允许部署的只有：

```plain text
workers/nlm-worker-mcp/
```

首次部署不接 runner，因此不设置：

```plain text
NLM_RUNNER_URL
NLM_RUNNER_TOKEN
```

首次部署不做：

- 不调用 NotebookLM；
- 不处理 Google cookie；
- 不执行 source add / sync / delete / share / deep research；
- 不让 Worker 写 GitHub；
- 不把 NotebookLM 回答写回沃壤。

## 3. 准备 wrangler.toml

在有 Cloudflare 权限的环境执行：

```bash
cd workers/nlm-worker-mcp
cp wrangler.toml.example wrangler.toml
```

确认 `wrangler.toml` 中 vars 只包含第一阶段变量：

```toml
[vars]
GITHUB_REPO = "dongxi-heji/worang"
GITHUB_REF = "main"
NLM_MANIFEST_PATH = "wo/nlm/manifest.json"
```

如果需要让现有 becoming pack 显示 notebook 绑定，可加：

```toml
NLM_BECOMING = "<actual notebook id>"
```

注意：`wrangler.toml` 若包含真实 notebook id，不应提交回 GitHub；当前仓库只提交 `wrangler.toml.example`。

完成后回到仓库根目录运行 smoke 脚本：

```bash
cd ../..
```

## 4. 设置 secret

必需：

```bash
wrangler secret put MCP_BEARER_TOKEN
```

可选：

```bash
wrangler secret put GITHUB_TOKEN
```

`GITHUB_TOKEN` 只应用 fine-grained token，权限为 Contents read-only；公开 repo 可先不设。

暂不设置：

```bash
wrangler secret put NLM_RUNNER_TOKEN
```

## 5. 本地 dev 最小验收

以下命令默认从仓库根目录运行。

先跑 Wrangler health smoke：

```bash
MCP_BEARER_TOKEN="<local-token>" \
NLM_WRANGLER_SMOKE_MODE=dev \
bash scripts/nlm-worker-wrangler-smoke.sh
```

再跑 MCP HTTP smoke：

```bash
MCP_BEARER_TOKEN="<local-token>" \
bash scripts/nlm-worker-wrangler-mcp-smoke.sh
```

再跑 bearer auth smoke：

```bash
MCP_BEARER_TOKEN="<local-token>" \
bash scripts/nlm-worker-wrangler-auth-smoke.sh
```

## 6. 部署

只有上一步通过后才执行：

```bash
cd workers/nlm-worker-mcp
wrangler deploy
```

部署后记录：

```plain text
worker name
worker URL
deployment time
wrangler version
commit sha
```

## 7. 部署后验收

设定：

```bash
WORKER_URL="https://<worker-url>"
TOKEN="<MCP_BEARER_TOKEN>"
```

### 7.1 healthz

```bash
curl -sS "$WORKER_URL/healthz"
```

预期：

```json
{
  "ok": true,
  "service": "nlm-worker-mcp"
}
```

### 7.2 tools/list

```bash
curl -sS -X POST "$WORKER_URL/mcp" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}'
```

预期看到：

```plain text
nlm_manifest_list_packs
nlm_manifest_get_pack
nlm_check_pack_freshness
nlm_list_notebooks
nlm_list_sources
nlm_query_grains
```

### 7.3 manifest list

```bash
curl -sS -X POST "$WORKER_URL/mcp" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"nlm_manifest_list_packs","arguments":{"kind":"source_pack"}}}'
```

预期：返回 pack count，不泄露真实 notebook/source id。

### 7.4 query_grains 关门

```bash
curl -sS -X POST "$WORKER_URL/mcp" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"nlm_query_grains","arguments":{"source_pack_key":"wo/nlm/becoming-projects-overview.md","question":"沃壤里 NotebookLM 应该扮演什么角色？"}}}'
```

预期：

```json
{
  "ok": false,
  "error": "runner_not_configured"
}
```

这是第一阶段成功条件。

## 8. 失败即停条件

出现以下任一情况，停止部署后续动作：

- `/mcp` 无 token 不返回 401；
- 错 token 不返回 401；
- `tools/list` 返回非预期工具；
- manifest tool 泄露真实 id；
- `nlm_query_grains` 在无 runner 时没有关门；
- Worker 试图调用 runner 或 NotebookLM；
- Worker 出现写 GitHub 行为。

## 9. 部署后仍不做

首次部署完成后仍不配置：

```plain text
NLM_RUNNER_URL
NLM_RUNNER_TOKEN
```

下一阶段另起小刀处理 runner 绑定。
