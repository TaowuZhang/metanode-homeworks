# GetNote Mirror MCP smoke test · 2026-06-04

状态：Notion MCP 已连通；v0.1 读入口可用；v0.1.2 smoke 工具候选。

## Scope

本记录只覆盖 GetNote GitHub mirror 的只读 MCP 入口。

不覆盖：

- full backfill；
- `--all`；
- GetNote 写回；
- 标签修改；
- 分享 / 删除；
- 自定义域名绑定。

## Connection

```plain text
Name: GetNote Mirror
Server URL: https://getnote-mirror-mcp.1171964523.workers.dev/mcp
Authentication: Bearer token
Status: ready
```

不要把 bearer token 写入 Git 或文档。

## Tools visible from Notion

```plain text
getnote_mirror_health
getnote_mirror_observability
getnote_recent_notes
getnote_search_notes
getnote_read_note
getnote_list_by_source
getnote_list_by_topic
getnote_list_by_tag
getnote_list_by_content_kind
```

v0.1.2 candidate adds:

```plain text
getnote_mirror_smoke
```

## Smoke results

### Observability

```plain text
health_status=healthy
mirror_version=v0.1
requested_limit=60
listed_count=20
effective_limit=20
written_count=20
manifest_row_count=20
recent_md_count=20
search_row_count=20
```

Content kinds:

```plain text
content=19
ref_content=1
metadata_only=0
empty_invalid=0
```

All observability gates returned `ok=true`.

### Recent notes

`getnote_recent_notes({ limit: 3 })` returned recent rows including:

```plain text
1911322061248449200 | 肉身承担的不可逆时刻
1911248294110356008 | 大禹与河流的分形物理学
1911248289816773296 | 当我们说出名字时，究竟承诺了什么
```

### Sentinel search

`getnote_search_notes({ query: "你一生的故事", limit: 5 })` returned:

```plain text
id=1909926264522813232
title=Ref · 你一生的故事（译林幻系列）
source=dedao
note_type=ref
content_kind=ref_content
path=资源/得到/_getnote/recent/1909926264522813232-Ref · 你一生的故事(译林幻系列).md
```

### Sentinel read

`getnote_read_note({ id: "1909926264522813232" })` returned the markdown note. The body includes:

```plain text
人类发展出前后连贯的意识模式，而七肢桶却发展出同步并举式的意识模式。我们依照先后顺序来感知事件，将各个事件之间的关系理解为因与果。它们则同时感知所有事件，并按所有事件均有目的的方式来理解它们，有最小目的，也有最大目的。
```

### Content kind filters

```plain text
getnote_list_by_content_kind(content_kind="content", limit=3): count=19
getnote_list_by_content_kind(content_kind="ref_content", limit=5): count=1
```

### Source filter

```plain text
getnote_list_by_source(source="dedao", limit=5): count=1
```

### Negative id

```plain text
getnote_read_note(id="does-not-exist-smoke"): found=false
```

## Known edge

A Notion UI call to the original no-argument `getnote_mirror_health` surfaced:

```plain text
payload.toolArguments should be defined, instead was `undefined`
```

This was not a mirror failure. Calls with explicit arguments succeeded, and the mirror read/search path was healthy.

v0.1.2 candidate changes:

- `getnote_mirror_health` now expects `{ check: "health" }`;
- `getnote_mirror_smoke({ mode: "basic" })` is the preferred first Notion smoke check.

## Interpretation boundary

This is a recent-window mirror. Do not call it the full GetNote brain.

If `requested_limit > listed_count` while `written_count == listed_count`, treat that as an upstream recent-window cap, not mirror data loss. Do not respond by running `--all` first.
