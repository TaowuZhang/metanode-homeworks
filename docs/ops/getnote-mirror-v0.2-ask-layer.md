# GetNote mirror v0.2：好问层与 Cloudflare MCP

日期：2026-06-04
状态：候选施工 / Worker MCP v0.2
位置：`workers/getnote-mirror-mcp/` 与 `资源/得到/_getnote/`

## 目标

把 GetNote / 得到大脑相关操作继续收束到 Cloudflare Worker MCP：

1. 克可以通过 MCP 稳定读取 GetNote GitHub mirror。
2. v0.2 不扩权限，先把“好问 recent-window”的能力放进 Worker。
3. 工具输出能和资源总登记联动，但只产出候选，不自动沉积、不改标签。

本轮不是退缩到“只敢设计”。Worker v0.1 已经存在，v0.2 直接进入代码候选：新增只读问法工具、边界 flags、候选登记 payload 与基础防越界检查。

## 现有地基

已有 Cloudflare Worker MCP：

```plain text
workers/getnote-mirror-mcp/
  worker.js
  README.md
  wrangler.toml.example
```

已有只读工具：

```plain text
getnote_mirror_health()
getnote_mirror_observability(format?)
getnote_recent_notes(limit?, source?, content_kind?)
getnote_search_notes(query, limit?)
getnote_read_note(id)
getnote_list_by_source(source, limit?)
getnote_list_by_topic(topic, limit?)
getnote_list_by_tag(tag, limit?)
getnote_list_by_content_kind(content_kind, limit?)
```

已有观测面显示 recent-window 当前健康：`written_count = 20`、`manifest_row_count = 20`、`recent_md_count = 20`、`search_row_count = 20`，并保留 ref 哨兵。

## v0.2 新增工具

### `getnote_digest_recent`

用途：把 recent-window 的 metadata 组织成近期主题、代表候选和资源登记候选。

输入：

```json
{
  "limit": 20,
  "query": "optional keyword",
  "source": "optional exact source",
  "content_kind": "optional exact content_kind",
  "max_notes_read": 10
}
```

输出重点：

- `note_count_seen`
- `one_line`
- `themes`
- `registry_candidates`
- `warnings`

边界：只 digest recent-window，不读全量得到大脑。

### `getnote_search_by_time`

用途：按时间窗口问 recent-window，避免一问时间就扩成全量。

输入：

```json
{
  "field": "created_at",
  "since": "2026-05-29",
  "until": "2026-06-04",
  "timezone": "Asia/Shanghai",
  "query": "optional keyword",
  "source": "optional exact source",
  "content_kind": "optional exact content_kind",
  "limit": 20
}
```

输出重点：

- `time`
- `count`
- `rows`
- `warnings`

边界：命中为 0 就返回 0，不自动扩大窗口。

### `getnote_note_brief`

用途：把单条 mirrored note 转成可登记 brief。

输入：

```json
{
  "id": "note_id",
  "brief_style": "registry_candidate",
  "include_markdown_excerpt": false,
  "max_chars": 1200
}
```

输出重点：

- `note`
- `brief.what_it_is`
- `brief.core_points`
- `brief.possible_use`
- `brief.registry_suggestion`

边界：brief 是候选，不是沉积；markdown excerpt 默认不开。

### `getnote_find_related_notes`

用途：在 recent-window 内为一条 note 或一个 query 找近邻。

输入：

```json
{
  "id": "optional note id",
  "query": "optional keyword",
  "relation_types": ["same_source", "same_topic", "time_neighbor", "text_similarity"],
  "limit": 10
}
```

输出重点：

- `anchor`
- `related`
- `relation_type`
- `reasons`
- `confidence`

边界：relation 是候选关系，不写回、不改标签。

### `getnote_cluster_recent`

用途：把 recent-window 临时分簇，给克下一步提问入口。

输入：

```json
{
  "limit": 100,
  "cluster_by": ["source", "topic", "tag", "content_kind", "registry_candidate"],
  "min_cluster_size": 2,
  "max_clusters": 12,
  "include_orphans": true
}
```

输出重点：

- `clusters`
- `representative_notes`
- `suggested_question`
- `registry_hint`
- `orphans`

边界：cluster 是临时阅读视图，不生成永久分类。

## 与 Resource Registry 的接口

v0.2 统一输出 `registry_candidate`，不直接写资源总登记。

最小字段：

```json
{
  "source_system": "getnote_mirror",
  "source_scope": "recent-window",
  "source_path": "资源/得到/_getnote/recent/example.md",
  "source_id": "note_id",
  "title": "note title",
  "source_name": "web | dedao | getnote_bu | unknown",
  "captured_at": "created_at or updated_at",
  "content_kind": "content | ref_content | metadata_only | unknown",
  "suggested_registry": "资源/书籍 | 资源/BibiGPT | 资源/语境 | 资源/得到 | 领域/代码 | 成为/项目 | only_recent",
  "relation_type": "excerpt_of | same_media_source | question_frame | domain_seed | project_material | orphan",
  "maturity": "raw | candidate | domain_seed | project_seed",
  "confidence": "low | medium | high",
  "human_confirm_required": true,
  "boundary_flags": {}
}
```

### 分流口径

| 去处 | v0.2 判定 |
|---|---|
| `only_recent` | 信号不足、低熟化、单条临时材料 |
| `资源/得到` | 人工确认后的长期 GetNote / 得到 source record |
| `资源/书籍` | 明确是书籍、书摘、阅读材料 |
| `资源/BibiGPT` | 明确对应视频 / 音频 / 转录 / 课程材料 |
| `资源/语境` | 是提问、写作、审查、操作口径 |
| `领域/` | 已进入长期能力域或方法判断 |
| `成为/项目` | 会改变当前项目下一刀 |

## 边界 flags

所有 v0.2 派生工具默认返回：

```json
{
  "source_scope": "recent-window",
  "is_full_brain": false,
  "no_backfill": true,
  "no_writeback": true,
  "no_tag_change": true,
  "relations_are_candidates": true
}
```

## 防越界检查

v0.2 Worker 对工具参数做最小拒绝：

- `backfill`
- `--all`
- `writeback`
- `delete`
- `share`
- `update_tag`
- `tag mutation`

若出现这些词，直接拒绝该调用。

这不是最终安全系统，只是第一道闸。后续还应加 smoke / golden output。

## smoke 候选

部署前先做三类 smoke：

1. **语法 smoke**
   - `node --check workers/getnote-mirror-mcp/worker.js`

2. **MCP tool list smoke**
   - `tools/list` 中应出现 v0.1 九个工具 + v0.2 五个工具。
   - `serverInfo.version = 0.2.0`。

3. **recent-window 行为 smoke**
   - `getnote_digest_recent({ limit: 5 })` 返回 scope flags。
   - `getnote_note_brief({ id })` 返回 registry candidate。
   - `getnote_search_by_time({ since, until })` 不越窗。
   - `getnote_find_related_notes({ id })` 返回 candidate relation。
   - `getnote_cluster_recent({ limit: 20 })` 返回 temporary clusters。

## observability 后续应加

v0.2 先把 flags 放进工具输出。下一刀再把以下检查写成自动 smoke：

1. scope guard：所有派生工具输出 `source_scope = recent-window`。
2. no-backfill guard：参数含 `backfill` / `--all` 必须拒绝。
3. no-writeback guard：不存在写回、删除、分享工具。
4. candidate relation guard：related / cluster 输出不得写成 confirmed。
5. registry payload guard：候选登记字段齐全。
6. empty-shell guard：brief 不得只有标题。
7. path guard：代表 note 必须带 id 与 path。
8. recent misread guard：输出不得声称“全量得到大脑”。

## 阶段命名

v0.1：Worker 只读 GitHub mirror。

v0.2：Worker 内置好问层，克能问 recent-window，并产出 Resource Registry 候选。

v0.3：再接 direct GetNote API / CLI hosted MCP，仍默认只读。

v0.4：另开 backfill 设计，必须 dry-run first。
