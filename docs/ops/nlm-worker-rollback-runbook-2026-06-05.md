# NLM Worker rollback runbook

日期：2026-06-05
状态：候选施工 / 回滚手册
位置：docs/ops

## 0. 一句话

NLM Worker 第一阶段回滚目标很简单：立刻切断 `/mcp` 可用面，保留沃壤主源不受影响；因为第一阶段不接 runner、不写 GitHub、不碰 NotebookLM，所以回滚只需要处理 Worker 部署与 secret。

## 1. 何时回滚

出现以下任一情况，立即回滚：

- `/mcp` 鉴权异常：无 token 或错 token 可访问；
- Worker 泄露 notebook id、source id、token 或其他敏感信息；
- `nlm_query_grains` 在无 runner 时没有返回 `runner_not_configured`；
- Worker 出现非预期工具；
- Worker 访问了 runner / NotebookLM；
- Worker 写 GitHub 或尝试写 GitHub；
- Cloudflare 日志出现异常请求或外部滥用迹象。

## 2. 快速止血

### 2.1 旋转 MCP_BEARER_TOKEN

最快方式：立刻重设 secret。

```bash
cd workers/nlm-worker-mcp
wrangler secret put MCP_BEARER_TOKEN
wrangler deploy
```

然后确认旧 token 失效：

```bash
curl -sS -o /tmp/nlm-worker-old-token.json -w '%{http_code}' \
  -X POST "$WORKER_URL/mcp" \
  -H "Authorization: Bearer <old-token>" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}'
```

预期：`401`。

### 2.2 删除或禁用 Worker 路由

如果入口暴露风险更高，直接从 Cloudflare 控制台禁用 Worker route，或删除对应 route。

### 2.3 回滚代码版本

若问题来自最新代码，回滚到上一部署版本，或重新部署已知安全 commit。

记录：

```plain text
rollback time
previous deployment id
new deployment id
reason
operator
```

## 3. 验收回滚成功

回滚后验收：

```bash
curl -sS "$WORKER_URL/healthz"
```

按回滚策略，可能是：

- route 禁用：不可访问；
- token 旋转：`/healthz` 可访问但 `/mcp` 旧 token 401；
- 代码回滚：工具面恢复到上一安全版本。

重点验收 `/mcp`：

```plain text
missing token → 401
wrong token   → 401
old token     → 401
new token     → 200 only if intentionally kept online
```

## 4. 沃壤主源不回滚

第一阶段 Worker 不写沃壤，所以通常不需要回滚 GitHub 主源。

需要检查但不改动：

- `wo/nlm/manifest.json` 是否被外部改动；
- `docs/ops/` 是否只新增候选文档；
- `workers/nlm-worker-mcp/` 是否仍在候选分支；
- main 是否未被误合并。

## 5. 事后沉积

若发生回滚，新增一份 incident 文档：

```plain text
docs/ops/nlm-worker-incident-YYYY-MM-DD.md
```

最小字段：

```plain text
发生时间
触发条件
影响范围
是否泄露 secret / id
止血动作
回滚动作
后续禁止项
是否需要修改 Worker / smoke
```

## 6. 不做

回滚时不要：

- 不为了排查而打印 secret；
- 不把 Cloudflare secret 复制到聊天；
- 不把 NotebookLM cookie 放进日志；
- 不在事故中临时打开 runner；
- 不扩大到 source sync / NotebookLM query。

## 7. 下一次再部署前

先回到仓库根目录，再重新跑：

```bash
cd ../..
bash scripts/nlm-worker-deploy-preflight.sh
MCP_BEARER_TOKEN="<local-token>" bash scripts/nlm-worker-wrangler-auth-smoke.sh
```

并确认 incident 中的触发条件已被修复。
