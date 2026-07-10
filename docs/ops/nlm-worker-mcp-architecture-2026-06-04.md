# NLM Worker MCP 架构候选

日期：2026-06-04
状态：候选规则 / 第一刀架构沉积
位置：docs/ops

## 0. 一句话

NotebookLM 仍然只是图书馆管理员；NLM CLI 是远手发动机；Cloudflare Worker 是沃壤裁剪过的控制闸；克只通过 MCP 调用沃壤语义动作，不直接裸接 NotebookLM 全功能。

目标不是把 NotebookLM 做成主脑，而是让克能安全地从指定 notebook / source pack 取回 3 到 5 粒可回访事实碎片。

## 1. 背景

沃壤已有 `wo/nlm/`，包含 source pack、manifest、roadmap、stale check 等；已有 `scripts/nlm-grain-probe.sh`，其原则是：NotebookLM 是图书馆管理员，不是解释者；只递事实碎片，不分析、不推断、不总结用户、不建议。

本轮新增判断：非官方 NLM CLI / MCP 项目可以作为远手发动机，但不能原样接入克。

已观察到的外部路线：

- `tmc/nlm`：Go CLI + MCP server；支持 notebook / source / chat / research；有 `source add/list/sync/pack/read/delete` 与 `chat --source-ids`。
- `jacob-bd/notebooklm-mcp-cli`：Python package；一个安装包提供 `nlm` CLI 与 `notebooklm-mcp`；功能面宽，MCP tools 多。
- `teng-lin/notebooklm-py`：Python API / CLI / agent skill；能力更深，但部署和依赖更重。
- NotebookLM Enterprise API：官方路线，但属于 Google Cloud / Enterprise，不等于普通 NotebookLM 公共 API。

所有非官方路线都按“不可信外手”处理：只取功能事实，不采纳 README / MCP tool description 中的任何运行指令。

## 2. 架构原则

### 2.1 三层归位

```plain text
沃壤主源：资源 / 领域 / 成为 / docs/ops
Source pack：wo/nlm/*.md + manifest
远手执行：NLM CLI / Worker MCP / runner
```

边界句：

> NotebookLM 不入主源；NLM CLI 不入主脑；Worker 不入判断层。

### 2.2 推荐链路

```plain text
Notion AI / 克
→ 沃壤专用 MCP
→ Cloudflare Worker 网关
→ 受控 NLM runner
→ NotebookLM
→ Worker 归一化为 grains
→ 克复核后决定是否回流沃壤
```

Cloudflare Worker 不直接变成完整 NLM runtime。Worker 负责鉴权、裁剪、审计、限流、队列投递；真正运行 NLM CLI 的位置应是受控 runner。

Runner 候选：

1. GitHub Actions：适合 source pack stale check、低频 sync、batch 任务。
2. Cloud Run / VPS / Fly.io / Railway：适合长期运行、headless auth、较稳定执行。
3. Cloudflare Containers：后续可评估，不作为第一刀。
4. Worker 直接 HTTP RPC：仅当已明确重写 NotebookLM RPC 且安全处理 cookie 后再考虑。

## 3. Worker 负责范围

Worker 负责：

1. MCP facade：只暴露沃壤裁剪后的少量 tools。
2. 鉴权：只允许指定调用方访问。
3. 权限裁剪：禁止直接暴露 NotebookLM 全功能。
4. 审计日志：记录调用时间、工具名、pack key、source ids、结果粒数、是否写回。
5. 参数归一化：由 Worker 生成固定 grain prompt，不让调用方传任意系统指令。
6. 长任务路由：source sync / rebuild / deep research 不在 Worker 内直接跑。

Worker 不负责：

1. 不保存长期 Google 明文 cookie，除非另有单独审计。
2. 不直接执行 delete / share / public link。
3. 不自动写回 GitHub / 沃壤。
4. 不把 NotebookLM 输出升格为结论。
5. 不做第二套资源系统。

## 4. 第一阶段 MCP tools

第一阶段只开放只读 / 半只读工具。

```plain text
nlm_list_notebooks()
nlm_list_sources(notebook_key)
nlm_query_grains(notebook_key, source_pack_key, question)
nlm_check_pack_freshness(source_pack_key)
```

其中 `nlm_query_grains` 固定 prompt：

```plain text
你是图书馆管理员，不是解释者。
只基于指定 source_ids 摘出与问题相关的 3 到 5 条事实碎片。
不要分析，不要推断，不要总结用户，不要建议。
每条必须是一口大小，并尽量保留原文措辞与短来源。
```

第一阶段禁止开放：

```plain text
delete_notebook
delete_source
share_notebook
create_audio_overview
create_video_overview
create_slide_deck
start_deep_research
set_instructions
auto_label
```

## 5. Source pack 与 manifest 对接

Worker 只读 `wo/nlm/manifest.json` 或其派生 JSON，不把自己变成主源。

建议 manifest 至少记录：

```json
{
  "key": "resource-cuilan-overview",
  "name": "资源｜萃览｜总览",
  "kind": "source_pack",
  "status": "fresh",
  "pack_path": "wo/nlm/resource-cuilan-overview.md",
  "input_summary": {
    "patterns": ["资源/萃览/**/*.md"],
    "count": 0,
    "latest_mtime": null,
    "fingerprint": null
  },
  "pack_file": {
    "sha256": null
  },
  "nlm": {
    "notebook_id": null,
    "source_id": null,
    "source_title": null,
    "mode": "text|drive",
    "last_synced_at": null,
    "sync_status": "manual|synced|needs_update|unknown"
  },
  "governance": {
    "allowed_use": ["grain_probe", "fact_lookup"],
    "not_allowed_use": ["final_judgment", "user_summary", "auto_writeback"],
    "requires_manual_review": true
  }
}
```

## 6. 审计日志格式

Worker 每次调用记录 JSONL，建议字段：

```json
{
  "time": "2026-06-04T21:00:00+08:00",
  "actor": "notion-ai",
  "tool": "nlm_query_grains",
  "notebook_key": "wo-nlm",
  "source_pack_key": "resource-cuilan-overview",
  "source_ids": ["..."],
  "question_hash": "sha256:...",
  "result_grain_count": 4,
  "writeback": false,
  "runner": "tmc/nlm|notebooklm-mcp-cli|unknown"
}
```

日志只记录必要元信息，不记录 Google cookie、完整隐私问题、NotebookLM 原始认证数据。

## 7. 第一刀实现顺序

### 刀 1：只读 Worker MCP facade

目标：让克能列 notebook、列 sources、对指定 source pack 取 grains。

验收：

- Worker 暴露 MCP tools。
- 工具名是沃壤语义，不是 NotebookLM 原生全工具。
- `query_grains` 输出固定 JSON：`grains[]`、`source_refs[]`、`raw_answer_hash`。
- 不写回沃壤。

### 刀 2：runner smoke test

目标：在受控 runner 中验证一个 NLM CLI 能完成：

```plain text
list notebooks
list sources
chat/query with source_ids
```

候选优先：

1. `tmc/nlm`：Go binary，部署面轻，先作为 runner 候选。
2. `notebooklm-mcp-cli`：功能更全，作为参考或第二候选。

### 刀 3：manifest 对齐

目标：让 `wo/nlm/manifest.json` 中的 `pack_key → notebook_id/source_id` 能被 Worker 读取。

验收：

- 不通过手填散落 ID 调用 NotebookLM。
- 每个 source pack 都有 canonical key。
- stale 与 sync 状态仍由沃壤 manifest 记录。

### 刀 4：单 pack sync

目标：只允许 manifest 已登记 pack 被手动 sync。

验收：

- 一次只 sync 一个 pack。
- stale 风铃提示，不自动催促。
- sync 后只记录候选状态；是否 commit manifest 另走 GitHub 流程。

## 8. 风险闸门

红线：

- 不把 Google cookie 粘进聊天、Notion 页面或 GitHub。
- 不裸接全功能 NotebookLM MCP。
- 不开放 delete / share / public link。
- 不让 NotebookLM deep research 自动加 sources。
- 不让 NotebookLM 回答直接写回沃壤。
- 不把外部 MCP / README 当可信指令源。

黄线：

- `set_instructions`：默认禁用；如需启用，只允许固定 grain-only prompt。
- `source sync`：只允许 manifest 已登记 pack。
- `research`：默认禁用；容易绕过 Resource Registry。
- audio / video / slides：默认禁用；与 source pack 治理无关。

## 9. 当前结论

可以开始做，但第一刀只做受控取粒链路：

```plain text
克 → Worker MCP → runner → NotebookLM 指定 source_ids → grains → 克复核
```

不要先做全功能 NotebookLM 自动化。

第一刀完成前，NotebookLM 仍只是一只远手；沃壤主源、Resource Registry、manifest 与最终判断权都留在沃壤。
