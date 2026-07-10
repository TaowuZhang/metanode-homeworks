# docs/legacy · 退役参考层

`docs/legacy/` 保存已经退役、被替换或只剩历史解释价值的仓库侧文档。

它不是冷启动入口，不是当前方案库，也不作为新施工的默认参考。读取这里时，应先确认当前入口：`README.md`、`AGENTS.md`、`CLAUDE.md`、`docs/ops/worang-index/current.md`。

## 当前文件

| 文件 | 读法 | 后续动作 |
|---|---|---|
| `notion-local-github-architecture-v1.md` | 旧 Notion / GitHub 架构解释 | 仅作历史参考；若当前入口已吸收，可摘要后瘦身 |
| `notion-structure-cleanup-2026-06-20.md` | 一次性清理记录 | 可并入 legacy index 后删除候选 |
| `terminal-bridge-ops.md` | 旧终端桥操作层 | 若已被 `wo/` 或 `docs/ops` 新入口替代，可摘要后瘦身 |
| `自动继续值守机.md` | 退役 watcher / 值守构想 | 若不再运行，保留摘要即可 |

## 放什么

- 已退役但解释当前结构来历的文档。
- 一次性迁移 / 清理记录，且尚未被更高入口吸收。
- 旧方案墓碑，但必须能说明“为什么不再用”。

## 不放什么

- 当前操作指南。
- 新的 ideas。
- 长期判断正文。
- 未定归位材料。

## 退场口

每个 legacy 文件最终只保留三种结果之一：

1. 被当前入口吸收后删除。
2. 合并为本 README 的一行历史索引后删除原文。
3. 因仍需完整回访而保留，并在表中写明理由。
