# NLM HTTP runner scaffold

日期：2026-06-04
状态：候选施工 / 未部署
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-runner-http.mjs` 是 Worker 后面的受控 runner 雏形：它只接收 allowlist action，调用本机已配置好的 `nlm`，并把输出归一化为 JSON。它不是公开服务，也不负责 NotebookLM 登录。

## 1. 本刀新增

- `scripts/nlm-runner-http.mjs`

它服务于上一刀的 `workers/nlm-worker-mcp/worker.js`，对齐 Worker 当前发送的 payload：

```json
{
  "action": "query_grains",
  "notebook_id": "...",
  "source_ids": ["..."],
  "prompt": "..."
}
```

## 2. 支持动作

```plain text
list_notebooks
list_sources
query_grains
```

不支持：

```plain text
source add
source sync
delete
share
deep research
audio/video/slides
set instructions
任意 shell
```

## 3. 使用方式

默认只监听本机：

```bash
NLM_RUNNER_TOKEN="<runner-token>" \
node scripts/nlm-runner-http.mjs
```

默认地址：

```plain text
http://127.0.0.1:8788
```

健康检查：

```bash
curl http://127.0.0.1:8788/healthz
```

查询示例：

```bash
curl -X POST http://127.0.0.1:8788/ \
  -H "Authorization: Bearer <runner-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "action": "query_grains",
    "notebook_id": "<notebook-id>",
    "source_ids": ["<source-id>"],
    "prompt": "你是图书馆管理员，不是解释者。只摘出 3 到 5 条事实碎片。"
  }'
```

## 4. 安全边界

- runner 不保存 Google cookie；它只假设所在环境已经完成 `nlm` 登录。
- runner 默认绑定 `127.0.0.1`，不直接暴露公网。
- 如果要放到 Cloud Run / VPS，外层必须加 HTTPS、鉴权、最小网络入口。
- `NLM_RUNNER_TOKEN` 只保护 runner；不等于 Google 认证。
- 输出只返回 hash、grains、source refs、stderr tail，不写回沃壤。

## 5. 与 Worker 对接

Worker 侧设置：

```plain text
NLM_RUNNER_URL = "https://runner.example/"
NLM_RUNNER_TOKEN = "..."
```

Worker 仍负责：

- 读 manifest；
- 判断 source pack 是否允许 grain lookup；
- 生成固定 prompt；
- 不让克直接传 NotebookLM 任意工具。

Runner 负责：

- 调用本机 `nlm`；
- 处理不同 CLI 形态：`nlm query notebook` 或 `nlm chat --source-ids`；
- 把 answer 里的 `· ` grains 抽出来。

## 6. 下一步验收

1. 在一个安全环境安装 `tmc/nlm`。
2. 跑 `scripts/nlm-runner-smoke.sh` 的 help 级。
3. 启动 `scripts/nlm-runner-http.mjs`。
4. 调 `list_notebooks`。
5. 有明确 notebook/source 绑定后，再调 `query_grains`。
6. 最后把 Worker 的 `NLM_RUNNER_URL` 指向 runner。
