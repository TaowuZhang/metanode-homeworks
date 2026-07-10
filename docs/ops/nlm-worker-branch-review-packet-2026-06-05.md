# NLM Worker MCP branch review packet

日期：2026-06-05
状态：分支收束 / PR 前审阅包
分支：`docs/nlm-worker-mcp-2026-06-04`

## 0. 一句话

本分支把 NotebookLM 从“可讨论的远手”推进到“可审阅的受控 Worker / runner 候选链路”：先裁剪 Worker MCP facade，再准备 runner smoke、manifest 绑定、Wrangler 本地检查、鉴权反向检查、部署 preflight、首次部署与回滚 runbook。

仍然没有接 runner、没有部署 Worker、没有调用 NotebookLM、没有写 Google cookie。

## 1. 本分支目标

建立一条受控链路：

```plain text
克 / Notion AI
→ NLM Worker MCP facade
→ 受控 runner
→ NotebookLM 指定 source_ids
→ grains JSON
→ 克复核后再决定是否回流沃壤
```

第一阶段只做到 Worker facade 可审阅、可本地 smoke、可部署前检查。

## 2. 新增范围

### 2.1 Worker facade

```plain text
workers/nlm-worker-mcp/
  worker.js
  README.md
  wrangler.toml.example
```

作用：

- 暴露 MCP `/mcp`；
- 支持 `/healthz`；
- 只开放裁剪工具面；
- 读取 `wo/nlm/manifest.json`；
- 支持现有 `upload.notebook_id_env + upload.source_id` 绑定；
- 无 runner 时 `nlm_query_grains` 返回 `runner_not_configured`。

### 2.2 Runner / Worker 检查脚本

```plain text
scripts/nlm-runner-smoke.sh
scripts/nlm-runner-install-smoke.sh
scripts/nlm-runner-http.mjs
scripts/nlm-worker-manifest-dry-run.mjs
scripts/nlm-worker-local-smoke.mjs
scripts/nlm-worker-wrangler-smoke.sh
scripts/nlm-worker-mcp-http-smoke.mjs
scripts/nlm-worker-wrangler-mcp-smoke.sh
scripts/nlm-worker-bearer-smoke.mjs
scripts/nlm-worker-wrangler-auth-smoke.sh
scripts/nlm-worker-deploy-preflight.sh
```

### 2.3 Ops 文档

```plain text
docs/ops/nlm-worker-mcp-architecture-2026-06-04.md
docs/ops/nlm-runner-smoke-2026-06-04.md
docs/ops/nlm-http-runner-2026-06-04.md
docs/ops/nlm-runner-validation-2026-06-04.md
docs/ops/nlm-runner-install-smoke-2026-06-04.md
docs/ops/nlm-manifest-binding-compat-2026-06-04.md
docs/ops/nlm-worker-manifest-dry-run-2026-06-04.md
docs/ops/nlm-worker-local-smoke-2026-06-04.md
docs/ops/nlm-worker-wrangler-smoke-2026-06-04.md
docs/ops/nlm-worker-mcp-http-smoke-2026-06-04.md
docs/ops/nlm-worker-bearer-smoke-2026-06-05.md
docs/ops/nlm-worker-deploy-preflight-2026-06-05.md
docs/ops/nlm-worker-first-deploy-runbook-2026-06-05.md
docs/ops/nlm-worker-rollback-runbook-2026-06-05.md
```

## 3. 工具面

Worker 第一阶段 tools：

```plain text
nlm_manifest_list_packs
nlm_manifest_get_pack
nlm_check_pack_freshness
nlm_list_notebooks
nlm_list_sources
nlm_query_grains
```

其中 `nlm_list_notebooks`、`nlm_list_sources`、`nlm_query_grains` 都需要 runner；未配置 runner 时必须失败关门。

## 4. 明确不做

本分支不做：

- 不配置 Cloudflare secret；
- 不执行 `wrangler deploy`；
- 不设置 `NLM_RUNNER_URL`；
- 不安装 / 登录 NotebookLM；
- 不写 Google cookie；
- 不做 source add / sync；
- 不开放 delete / share / deep research / audio / video / slides；
- 不让 Worker 写 GitHub；
- 不把 NotebookLM 回答写回沃壤。

## 5. 验收顺序

建议审阅 / 执行顺序：

```bash
# 1. 静态部署前检查
bash scripts/nlm-worker-deploy-preflight.sh

# 2. manifest dry-run
node scripts/nlm-worker-manifest-dry-run.mjs list
node scripts/nlm-worker-manifest-dry-run.mjs query-plan "wo/nlm/becoming-projects-overview.md"

# 3. Worker local smoke
node scripts/nlm-worker-local-smoke.mjs

# 4. 有 Wrangler 的环境下
MCP_BEARER_TOKEN="<local-token>" NLM_WRANGLER_SMOKE_MODE=dev bash scripts/nlm-worker-wrangler-smoke.sh
MCP_BEARER_TOKEN="<local-token>" bash scripts/nlm-worker-wrangler-mcp-smoke.sh
MCP_BEARER_TOKEN="<local-token>" bash scripts/nlm-worker-wrangler-auth-smoke.sh
```

## 6. 首次部署闸门

只有以下全部成立，才允许首次部署：

- preflight 返回 `preflight_ok`；
- Worker local smoke 通过；
- Wrangler `/healthz` 通过；
- MCP HTTP smoke 通过；
- bearer auth smoke 通过；
- 已准备 `MCP_BEARER_TOKEN`；
- 未设置 `NLM_RUNNER_URL`；
- 未设置 NotebookLM cookie。

首次部署后只验：

```plain text
GET /healthz
POST /mcp initialize
POST /mcp tools/list
manifest list/get/freshness
query_grains -> runner_not_configured
```

## 7. 回滚闸门

任何以下情况出现，立即回滚或禁用 route：

- 无 token / 错 token 可访问 `/mcp`；
- manifest 工具泄露真实 notebook id / source id；
- 无 runner 时 `query_grains` 没有关门；
- 出现非预期工具；
- Worker 访问 runner / NotebookLM；
- Worker 写 GitHub。

## 8. 审阅重点

重点看：

1. `workers/nlm-worker-mcp/worker.js` 是否真的只暴露裁剪工具；
2. `resolveNotebookId` / `resolveSourceId` 是否正确兼容现有 manifest；
3. `publicPack` 是否脱敏；
4. `authorize` 是否正确使用 `MCP_BEARER_TOKEN`；
5. smoke 脚本是否默认关门、不自动安装、不自动部署、不自动 query；
6. runbook 是否明确禁止 runner / NotebookLM 第一阶段接入。

## 9. 下一阶段候选

若本分支审阅通过，下一阶段再开新刀：

1. 在 Cloudflare 权限环境部署 facade；
2. 只连 manifest tools；
3. 再在独立 runner 环境安装 `tmc/nlm`；
4. 跑 help 级 smoke；
5. 跑 query 级 smoke；
6. 最后才配置 `NLM_RUNNER_URL`。
