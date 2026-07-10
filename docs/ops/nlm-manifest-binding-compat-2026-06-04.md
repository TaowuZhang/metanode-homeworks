# NLM manifest binding compatibility

日期：2026-06-04
状态：候选施工 / Worker 兼容层
位置：workers/nlm-worker-mcp + docs/ops

## 0. 一句话

现有 `wo/nlm/manifest.json` 已经用 `upload` 字段记录部分 NotebookLM 绑定；Worker 不应强迫立刻迁移到新 `nlm` 字段，而应兼容现有 manifest。

## 1. 现有形态

当前 manifest 中已出现：

```json
{
  "upload": {
    "should_upload": true,
    "notebook": "沃壤｜成为",
    "notebook_id_env": "NLM_BECOMING",
    "source_id": "..."
  }
}
```

这意味着：

- `source_id` 已经可以直接绑定 NotebookLM source；
- `notebook_id` 不直接写入 manifest，而是通过环境变量名寻址；
- 这是合理的，因为 notebook id 可作为部署环境配置，不必硬写进每个 pack。

## 2. Worker 兼容规则

Worker 现在按以下顺序解析 notebook id：

```plain text
pack.nlm.notebook_id
→ pack.notebook_id
→ pack.upload.notebook_id
→ env[pack.upload.notebook_id_env]
→ env[pack.nlm.notebook_id_env]
```

按以下顺序解析 source id：

```plain text
pack.nlm.source_id
→ pack.source_id
→ pack.upload.source_id
```

因此，现有 `upload.notebook_id_env + upload.source_id` 可以直接服务 `nlm_query_grains`。

## 3. Worker 输出脱敏

`nlm_manifest_list_packs` 与 `nlm_manifest_get_pack` 不直接回显真实 notebook id / source id，只显示：

```plain text
<bound>
```

并返回 `binding` 摘要：

```json
{
  "notebook_id_env": "NLM_BECOMING",
  "notebook_id_env_bound": true,
  "has_upload_source_id": true
}
```

这样克能看见绑定是否成立，但不会把真实 id 当作聊天内容到处复制。

## 4. 下一步部署变量

若要让 Worker 读取现有 `沃壤｜成为` notebook，需要在 Worker 环境变量中设置：

```plain text
NLM_BECOMING=<actual notebook id>
```

同时仍需要：

```plain text
NLM_RUNNER_URL=<controlled runner url>
NLM_RUNNER_TOKEN=<runner token>
```

没有 runner 时，Worker 仍然返回 `runner_not_configured`，不触 NotebookLM。

## 5. 边界

本刀不做：

- 不改 manifest 数据；
- 不把 notebook id 写死进仓库；
- 不暴露 source id 原文；
- 不部署 Worker；
- 不调用 NotebookLM；
- 不迁移旧字段。

## 6. 后续迁移方向

长期可以把 manifest 统一成：

```json
{
  "nlm": {
    "notebook_id_env": "NLM_BECOMING",
    "source_id": "...",
    "sync_status": "manual"
  }
}
```

但这不是当前小刀。当前小刀只保证 Worker 尊重现有地形。
