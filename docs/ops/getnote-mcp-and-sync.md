# GetNote 接入：recent-window 镜像与证据分流

GetNote 通过只读 CLI / OpenAPI，把最近一小窗笔记镜像到 `资源/得到/_getnote/`。它是机器镜像区，不替代旧的 `资源/得到/`，也不成为第二套知识母本。

## 当前入口

- 同步：`scripts/sync-getnote.mjs`
- 尝试验收：`scripts/evaluate-getnote-attempt.mjs`
- 定时入口：`.github/workflows/sync-getnote.yml`
- 周度证据检查：`scripts/workflow-evidence-inspector.mjs`

## 同步边界

只读取最近笔记，默认 20 条；只生成 manifest、recent Markdown、轻量索引和 observability；只修改 `_getnote/`。不做全量补齐，不写回或删除外部内容，不修改旧 `资源/得到/`。

## 数据形态

| `content_kind` | 含义 | 处理 |
|---|---|---|
| `content` | 主正文存在 | 正常镜像 |
| `ref_content` | 引用正文存在 | 正常镜像 |
| `metadata_only` | 无正文但身份与 metadata 有效 | 保留 |
| `empty_invalid` | 缺有效身份且无内容或 metadata | 跳过并判异常 |

无标题不单独等于 parser 失败。只要 ID、内容形态和路径成立，它先作为 advisory。

## Rolling-window sentinel

固定笔记可能自然离开最近窗口：

- sentinel 不在窗口：advisory；
- sentinel 在窗口但 title 或 `content_kind` 不匹配：硬失败。

稳定 contract probe 应另行按 ID 只读探测，不能要求固定旧笔记永远留在滚动窗口。

## 健康快车道与异常慢车道

```text
sync-getnote.mjs
→ 生成本次 attempt
→ evaluate-getnote-attempt.mjs
```

### Healthy

所有硬门通过且修改范围正确时，本次镜像进入仓库，并刷新 `last-healthy.json`。正常确定性镜像不需要每次开 PR。

### Degraded

异常尝试不更新仓库中的健康镜像；本次诊断材料只进入短期 artifact，`last-healthy.json` 保持不动，并由 weekly governance 浮出异常。

## 硬门与 advisory

硬门：

- listed / written / manifest / recent 数量不一致；
- `id=0`、缺 ID、`empty_invalid`；
- manifest path 冲突；
- sentinel 已出现但结构不符；
- 修改越出 `_getnote/`。

Advisory：

- 上游窗口少于请求上限，但 listed 与 written 一致；
- sentinel 不在滚动窗口；
- 有有效身份和内容的 Untitled 记录。

## 两个时间锚点

- **Last attempt**：全局最新尝试以 GitHub Actions run 为准。每次运行也会生成 `last-attempt.json` receipt；healthy 时它随镜像进入仓库，degraded 时只留在短期 artifact。
- **Last healthy**：`last-healthy.json`，保存最近一次被允许进入仓库的真实运行结果。

因此仓库中的 `last-attempt.json` 不能单独证明全局没有更新的失败尝试；weekly governance 会读取 Actions 的最新 run，再与 `last-healthy` 比较。

## 修复闭环

一次后续修改只有满足三步，才算关闭真实问题：

1. 修改前出现可观察失败或降级；
2. 修改针对该失败面；
3. 修改后同一条真实 workflow 路径重新通过。

修复已合并但尚未重新运行，只能标记为待复验。

## 本地验证

```bash
node --test scripts/evaluate-getnote-attempt.test.mjs scripts/workflow-evidence-inspector.test.mjs
```

真实外部调用与凭据撤销仍沿用原有最小权限和只读边界。完整跨 workflow 口径见 `docs/ops/workflow-evidence-loop.md`。
