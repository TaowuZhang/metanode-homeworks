# NLM Worker manifest dry-run

日期：2026-06-04
状态：候选施工 / 本地 dry-run 工具
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-worker-manifest-dry-run.mjs` 是 Worker 部署前的本地检查口：用同一套 manifest 绑定规则，先确认 source pack 能被列出、找到、判 fresh/stale、以及是否具备 query 所需绑定。

它不启动 Worker，不调用 runner，不调用 NotebookLM。

## 1. 本刀新增

- `scripts/nlm-worker-manifest-dry-run.mjs`

## 2. 支持命令

```bash
node scripts/nlm-worker-manifest-dry-run.mjs list
node scripts/nlm-worker-manifest-dry-run.mjs get "wo/nlm/becoming-projects-overview.md"
node scripts/nlm-worker-manifest-dry-run.mjs freshness "wo/nlm/becoming-projects-overview.md"
node scripts/nlm-worker-manifest-dry-run.mjs query-plan "wo/nlm/becoming-projects-overview.md"
```

## 3. 绑定规则

Notebook id 解析顺序：

```plain text
pack.nlm.notebook_id
→ pack.notebook_id
→ pack.upload.notebook_id
→ env[pack.upload.notebook_id_env]
→ env[pack.nlm.notebook_id_env]
```

Source id 解析顺序：

```plain text
pack.nlm.source_id
→ pack.source_id
→ pack.upload.source_id
```

## 4. 典型用法

若 manifest 里写的是：

```json
{
  "upload": {
    "notebook_id_env": "NLM_BECOMING",
    "source_id": "..."
  }
}
```

则 dry-run 时可以这样检查 query 绑定：

```bash
NLM_BECOMING="<actual notebook id>" \
node scripts/nlm-worker-manifest-dry-run.mjs query-plan "wo/nlm/becoming-projects-overview.md"
```

输出中若：

```json
{
  "runnable": true
}
```

说明 manifest 绑定足够进入 Worker → runner 查询链路。

## 5. 边界

本脚本不做：

- 不读取 Google cookie；
- 不调用 NotebookLM；
- 不调用 NLM runner；
- 不部署 Worker；
- 不写回 manifest；
- 不暴露真实 notebook/source id，只显示 `<bound>`。

## 6. 下一步

1. 在完整仓库环境里跑 `list`。
2. 对已上传 pack 跑 `query-plan`。
3. 对 pack 文件跑 `freshness`。
4. 再部署 Worker 或做 Miniflare / Wrangler dry-run。
