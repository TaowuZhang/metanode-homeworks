# 发布包规格

发布包是原始沃壤事件与平台外手之间的托盘。

它不是新真相源，不是内容库，不是运营看板。它只负责把「要排出去的东西」整理到平台可以消费、人可以检查、失败可以定位的形状。

## 路径

建议生成：

- `ai-browse/out/packages/<event-id>.md`
- `ai-browse/out/packages/<event-id>.json`

Markdown 给人检查，JSON 给外手读取。机器契约见：`schemas/publish-package.schema.json`。

## 最小字段 v0

| 字段 | 含义 | 要求 |
|---|---|---|
| `source_id` | 上游事件、文件、页面或材料 ID | 必填；能回到主源 |
| `source_path` | 上游路径或 URL | 必填；Git / Notion / 本地路径均可 |
| `canonical_title` | 唯一主标题 | 必填 |
| `canonical_body` | 母稿正文 | 必填；不得包含未筛隐私 |
| `summary_short` | 140–300 字短摘要 | 必填；用于轻平台 |
| `platform_variants` | 各平台改写版本 | 可空对象；缺省由适配器生成草案 |
| `assets` | 图片、截图、附件、封面等 | 可空数组 |
| `targets` | 本次目标平台 | 必填数组 |
| `publish_mode` | `dry_run` / `last_click` / `live` | 必填；默认 `dry_run` |
| `risk_notes` | 隐私 / 版权 / 身份 / 公共反作用检查 | 必填 |
| `callback_sink` | 回声回收位置 | 必填 |
| `created_at` | 生成时间 | 必填 |
| `human_check_required` | 是否需要谷人工确认 | 必填；公共发布默认 `true` |
| `status` | 当前状态 | `draft` / `ready_for_review` / `dry_run_passed` / `sent` / `failed` |

## 命令

```bash
wo excrete --json
wo excrete last --platforms xiaohongshu,x-twitter,linkedin,maimai --mode dry_run --prepare --json
node scripts/publish/prepare-package.mjs --source prism-entry-v1 --targets web,telegram,rss --mode dry_run
node scripts/publish/validate-package.mjs ai-browse/out/packages/prism-entry-v1.json
node scripts/publish/run-adapters.mjs --package ai-browse/out/packages/prism-entry-v1.json --targets web,telegram,rss --mode dry_run
```

`wo excrete --prepare` 与 `scripts/publish/*` 使用同一 schema。旧 `wo.publishing-package.v1` 形状已退役，不再由 `scripts/wo-excrete.mjs` 生成。

平台外手不在沃壤中重写：X 交给 `xurl`，X/LinkedIn/YouTube/TikTok 草稿交给 Postiz，国内视频/图文 uploader 计划交给 `sau`。固定版本与 doctor 见 `platform-cli-selection-2026-06-28.md`。

## 状态流转

```plain text
raw material → package draft → validate → dry-run → human check → last-click/live → collect result
```

- 生成发布包 ≠ 真正发布。
- 平台适配 ≠ 改写主源。
- 外发结果只写外手日志或指定回声入口，不反向覆盖沃壤事实。

## 与其他文档关系

- 分发总图：`distribution-map.md`
- 发布通道铺路：`distribution-road-v0.md`
- 外手登记：`distribution-cli-registry.v0.md`
- 发布候选放行卡：`distribution-publish-candidate-gate.md`
- 失败台账：`publish-failure-ledger.md`
