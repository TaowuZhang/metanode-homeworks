# NLM Worker Wrangler smoke

日期：2026-06-04
状态：候选施工 / Wrangler 启动前门
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-worker-wrangler-smoke.sh` 是 Worker 部署前的 Wrangler 检查前门：默认只做 Node 语法检查和 Wrangler 存在性检查；只有显式设置 `NLM_WRANGLER_SMOKE_MODE=dev` 时，才启动本地 Worker dev server 并检查 `/healthz`。

它不部署 Worker，不配置 runner，不调用 NotebookLM。

## 1. 本刀新增

- `scripts/nlm-worker-wrangler-smoke.sh`

## 2. 默认行为

```bash
bash scripts/nlm-worker-wrangler-smoke.sh
```

默认 `NLM_WRANGLER_SMOKE_MODE=check`。

会执行：

1. 检查 `workers/nlm-worker-mcp/` 是否存在；
2. 检查 `worker.js` 是否存在；
3. 执行：

```bash
node --check workers/nlm-worker-mcp/worker.js
```

4. 检查 `wrangler` 是否存在；
5. 若无 `wrangler`，返回 `missing_wrangler`，安全关门。

## 3. 本地 dev smoke

在已安装 Wrangler 的受控环境中执行：

```bash
NLM_WRANGLER_SMOKE_MODE=dev \
bash scripts/nlm-worker-wrangler-smoke.sh
```

脚本会：

1. 进入 `workers/nlm-worker-mcp/`；
2. 启动：

```bash
wrangler dev --local --ip 127.0.0.1 --port 8787
```

3. 请求：

```plain text
http://127.0.0.1:8787/healthz
```

4. 只验收 HTTP 200 与 health body。

## 4. 可配置变量

```plain text
NLM_WORKER_DIR=workers/nlm-worker-mcp
NLM_WRANGLER_SMOKE_MODE=check|dry-run|dev
NLM_WRANGLER_HOST=127.0.0.1
NLM_WRANGLER_PORT=8787
NLM_WRANGLER_SMOKE_OUT_DIR=target/nlm-worker-wrangler-smoke
```

## 5. 边界

本脚本不做：

- 不 `wrangler deploy`；
- 不写 Cloudflare secret；
- 不配置 `NLM_RUNNER_URL`；
- 不调用 runner；
- 不调用 NotebookLM；
- 不读取 Google cookie；
- 不做 `/mcp` 真实远程连接。

## 6. 下一步

若 `dev_health_ok`：

1. 再做 `/mcp` initialize / tools/list 的本地 HTTP 检查；
2. 仍不接 runner；
3. 确认 bearer token 行为；
4. 最后才进入 Cloudflare Worker 部署候选。
