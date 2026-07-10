# GetNote mirror v0.1：可观测性与索引层

日期：2026-06-04
位置：资源/得到/_getnote/ 与 docs/ops

## 目标

把 GetNote / 得到大脑 GitHub mirror 从「日更 v0 可跑」推进到两层稳定结构：

1. 第二层：可观测性闭环，能判断每日 mirror 是否健康，以及异常最小面在哪里。
2. 第三层：索引与语义入口，让克和后续工具能从 GitHub mirror 里稳定取材，但不把 recent 误当全量得到大脑。

本轮仍守住只读 mirror 边界：不写回得到、不删改标签、不分享、不跑全量 backfill。

## 当前地基

已成立：

- workflow：`Sync GetNote`
- schedule：Asia/Shanghai 每日 04:00，即 UTC `0 20 * * *`
- 手动触发：保留 `workflow_dispatch`
- 同步区：`资源/得到/_getnote/`
- 最近窗口：默认 20 条
- 健康报告：`health.json`
- 索引：`manifest.jsonl`
- 内容镜像：`recent/*.md`

2026-06-04 验收面：

- `requested_limit = 20`
- `listed_count = 20`
- `written_count = 20`
- `skipped_count = 0`
- `id_zero_count = 0`
- `empty_invalid_count = 0`
- `content_kind_counts = { content: 19, ref_content: 1 }`
- `1909926264522813232` 保持为 `Ref · 你一生的故事（译林幻系列）` / `ref_content`

## 第二层：可观测性闭环

### 目的

可观测性层不负责同步更多内容，只负责让 mirror 自己说清楚：

- 本次是谁触发的；
- 是否是 schedule；
- 写入数量是否符合预期；
- 是否出现坏 ID、坏空壳或异常 content_kind；
- 哨兵样本是否仍然保真；
- 如果异常，应该先看哪里，而不是立刻扩 backfill。

### 建议生成面

```plain text
资源/得到/_getnote/observability/
  README.md
  latest.json
  latest.md
```

`latest.json` 是机器读面；`latest.md` 是克 / 谷快速读面。

### `latest.json` 建议字段

```json
{
  "source": "getnote",
  "mirror_version": "v0.1",
  "generated_at": "ISO-8601",
  "health_status": "healthy | degraded",
  "run_context": {
    "provider": "github_actions | local",
    "workflow": "Sync GetNote",
    "event_name": "schedule | workflow_dispatch | other | null",
    "is_schedule": true,
    "is_manual": false,
    "run_id": "...",
    "run_attempt": "...",
    "actor": "github-actions[bot]",
    "ref_name": "main",
    "sha": "...",
    "run_url": "https://github.com/.../actions/runs/..."
  },
  "counts": {
    "requested_limit": 20,
    "listed_count": 20,
    "written_count": 20,
    "skipped_count": 0,
    "manifest_row_count": 20,
    "recent_md_count": 20
  },
  "gates": [],
  "sentinels": [],
  "mirror_shape": {}
}
```

### 健康门槛

第一批门槛只判 recent-window，不碰全量：

1. `written_count == requested_limit`
2. `id_zero_count == 0`
3. `missing_id_count == 0`
4. `empty_invalid_count == 0`
5. `content + ref_content > 0`
6. manifest 行数等于 recent md 数量
7. manifest path 唯一
8. 不存在 `Untitled GetNote note`
9. ref 哨兵 `1909926264522813232` 标题和 `content_kind` 保真

### 异常处置顺序

若降级：

1. 看 `health.json`，确认数量和坏空壳字段。
2. 看 `observability/latest.json`，确认是哪一道 gate 失败。
3. 看 `manifest.jsonl` 对应行。
4. 看 `recent/<id>-*.md`。
5. 只命名最小异常面；不得第一反应跑 `--all` 或扩 backfill。

## 第三层：索引与语义入口

### 目的

第三层不解释笔记，不摘要笔记，只把 recent-window 的事实面变成可查询入口。

它要避免两个误读：

- `recent/` 不是全量得到大脑；
- 语义层不能覆盖原始 mirror。

### 建议生成面

```plain text
资源/得到/_getnote/index/
  README.md
  by-id.json
  by-content-kind.json
  by-source.json
  by-topic.json
  by-tag.json

资源/得到/_getnote/search/
  README.md
  notes.jsonl
```

### 分工

- `manifest.jsonl`：事实索引，仍是第一索引源。
- `recent/*.md`：最近窗口正文镜像。
- `index/*.json`：由 manifest 派生的 lookup，不是第二事实源。
- `search/notes.jsonl`：轻量召回面，只放标题、路径、类型、source、tag、topic、时间，不放全文和 secrets。
- `raw/`：仍只预留，不默认写入原始 API payload。

### 语义入口边界

允许：

- 按 ID 找笔记；
- 按 `content_kind` 找 ref / content；
- 按 source 找 web / dedao / getnote_bu；
- 按 tag / topic 做轻量召回；
- 从索引跳到 `recent/*.md` 读取正文。

暂不允许：

- 自动摘要全部笔记；
- 自动改写旧 `资源/得到/`；
- 把 recent 窗口包装成全量大脑；
- 写回得到；
- 公开分享；
- 扩到 `--all`。

## 实施顺序

### 第一刀：沉文档和操作口径

本文件即第一刀。它定义第二层和第三层的地形、文件面、健康门槛和异常边界。

### 第二刀：脚本生成面

在 `scripts/sync-getnote.mjs` 中新增：

- `run_context()`：从 GitHub Actions 环境变量读取触发信息。
- `health_status`：由健康 gates 生成。
- `observability/latest.json` 与 `latest.md`。
- `index/*.json`。
- `search/notes.jsonl`。

脚本仍只写 `资源/得到/_getnote/`。

### 第三刀：dry-run / 手动 run 验收

先用小窗口验证生成面：

```bash
GETNOTE_API_KEY="..." node scripts/sync-getnote.mjs --limit 3 --dry-run
```

然后在 Actions 手动触发，确认：

- 新目录生成；
- `latest.json` 能显示 `workflow_dispatch`；
- 明早 schedule 后能显示 `schedule`；
- `health_status = healthy`。

### 第四刀：再谈 backfill

只有 recent-window 的观测和索引稳定后，才设计 backfill。backfill 是另一个系统，不塞进 v0.1。

## 收束命名

当第二层和第三层都落地后，本轮可命名为：

> GetNote mirror v0.1：日更可观测，recent-window 可索引。
