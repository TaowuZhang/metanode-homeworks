# GetNote Mirror MCP Worker

日期：2026-06-04
状态：候选设计 / v0.1 读入口，v0.1.2 smoke 工具，v0.2 好问层，v0.2.1 已对齐

## 背景

GetNote → GitHub recent-window mirror v0.1 已稳定。下一步不是 backfill，而是把这个 mirror 暴露成 Notion 可调用的只读 MCP。

## 部署位

谷已有 Cloudflare Worker，因此第一版放 Worker 合适。

Worker 不直接运行 `getnote` CLI。它只读 GitHub mirror：

```plain text
Notion Agent
→ MCP over HTTP
→ Cloudflare Worker
→ GitHub contents API
→ 资源/得到/_getnote/
```

## 第一版边界

允许：

- 读 health；
- 读 observability；
- 查 `search/notes.jsonl`；
- 读 `index/*.json`；
- 按 id 读取 `recent/*.md`。

不允许：

- 写回得到；
- 改标签；
- 删除；
- 分享；
- 触发 `--all`；
- 任意 shell；
- 直接暴露 `getnote` CLI。

## Worker secrets

```plain text
GITHUB_TOKEN       # GitHub contents read-only token
MCP_BEARER_TOKEN   # Notion 连接 MCP 时使用
```

Worker vars：

```plain text
GITHUB_REPO=dongxi-heji/worang
GITHUB_REF=main
GETNOTE_MIRROR_ROOT=资源/得到/_getnote
```

## MCP tools

```plain text
getnote_mirror_health(check)
getnote_mirror_smoke(mode)
getnote_mirror_observability(format?)
getnote_recent_notes(limit?, source?, content_kind?)
getnote_search_notes(query, limit?)
getnote_read_note(id)
getnote_list_by_source(source, limit?)
getnote_list_by_topic(topic, limit?)
getnote_list_by_tag(tag, limit?)
getnote_list_by_content_kind(content_kind, limit?)
```

这些工具全部是只读。`getnote_mirror_health` 需要显式 `check: "health"` 参数，避免零参调用在部分客户端下的歧义；`getnote_mirror_smoke` 跑一组只读自检（health 状态、observability gates、哨兵 note 可按 id 与搜索命中并可读 markdown），返回 `ok` 与 boundary flags，不写回、不改标签、不触发 backfill。

## v0.2 好问层

v0.2 增加好问层：

```plain text
getnote_digest_recent(limit?, query?, source?, content_kind?, max_notes_read?)
getnote_search_by_time(field?, since?, until?, timezone?, query?, source?, content_kind?, limit?)
getnote_note_brief(id, brief_style?, include_markdown_excerpt?, max_chars?)
getnote_find_related_notes(id?, query?, relation_types?, limit?)
getnote_cluster_recent(limit?, cluster_by?, min_cluster_size?, max_clusters?, include_orphans?)
```

这些工具仍然只读，只问 recent-window，并输出 Resource Registry 候选，不自动沉积。详细设计见 `getnote-mirror-v0.2-ask-layer.md`。

## Notion 连接口径

```plain text
你可以读取 GetNote GitHub mirror。
默认只读。
不要写回得到。
不要修改标签。
不要分享。
不要触发 backfill。
不要把 recent-window 误称为全量得到大脑。
```

## 后续阶段

v0.1：Worker 只读 GitHub mirror。

v0.1.2：health 改为显式 check 参数；新增 getnote_mirror_smoke 只读自检工具；serverInfo 版本对齐到 v0.2.1（语法与本地 fixture smoke 已在沙箱验过：node --check + smoke-local.mjs 全通过）。

v0.2：增加好问层：digest_recent / find_related_notes / search_by_time / note_brief / cluster_recent，详见 `getnote-mirror-v0.2-ask-layer.md`。

v0.3：再设计 direct GetNote CLI/API hosted MCP，但仍默认只读。

v0.4：单独设计 backfill，必须 dry-run first。
