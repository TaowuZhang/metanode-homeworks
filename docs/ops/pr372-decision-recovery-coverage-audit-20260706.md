# PR #372 判断链覆盖审计（2026-07-06）

## 0. 当前结论

PR #372 的判断链追回已经开始，但仍未完成。

当前已经写入：

- `docs/ops/pr372-decision-recovery-todo-20260706.md`
- `docs/ops/pr372-decision-recovery-master-checklist-20260706.md`
- `docs/ops/pr372-decision-recovery-coverage-audit-20260706.md`
- `docs/ops/pr372-decision-recovery-source-index-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-draft-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-hotsearch-rolling-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-platform-attention-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-portal-aggregators-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-institutional-media-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-action-sources-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-technology-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-community-20260706.md`

这些文件不能被描述为“完整判断账本”。

当前只能描述为：

> 已建立 PR #372 判断链追回控制层和 source index，并从 PR 文件中追回了第一批判断链。综合已补入多批平台源、门户聚合源、官方与调查媒体源、行动型来源；科技、社区已有第一版 ledger，但都仍是 partial。

禁止描述为：

- TopHub 80 + Folo 27 已经逐源补齐；
- source index 已经等于判断账本；
- 旧 ChatGPT 对话已经迁移完成；
- PR #372 可以直接合并；
- 综合 / 科技 / 社区已经全量闭环。

---

## 1. 覆盖状态表

| 模块 | 状态 | 已写入 | 当前缺口 | 下一步 |
|---|---|---|---|---|
| 控制层 | established | todo / master checklist / coverage audit | 无，但需持续维护 | 每新增 ledger 同步更新 |
| source index | established_not_judgment | TopHub 80、Folo 可见 20、省略 7、替换 / 移出来源 | 它是母表，不是判断链 | 按 needs_recovery / partial 回 ledger 补链 |
| 第一批草案 | partial | AI、设计、购物、财经、报刊、娱乐部分来源 | 不是 TopHub 80 + Folo 27 完整账本 | 拆分或继续细化 |
| 综合 | partial_recovered_from_pr | 已覆盖 closeout、知乎、热搜、滚动新闻、平台注意力源、门户聚合源、官方与调查媒体源、行动型来源 | directory review / rolling review 边界、部分候选复查条件仍需补 | 继续综合边界收口 |
| 科技 | partial_recovered_from_pr | final closeout 第一版 | 未覆盖科技全部同族 closure；多个替换理由仍需细化 | 继续 technology closure 文件族 |
| 社区 | partial_recovered_from_pr | pending actions 第一版 | 未覆盖社区全部 closure；社区热门第 2—20 页仍非全量闭环 | 继续 community closure 文件族 |
| AI reassessment | not_done | 无 | 大量让位理由未追回 | 读 6 个 reassessment |
| 开发目录 | not_done | 无 | 25 页未追回 | 生成 development ledger |
| 财经逐页 | not_done | final closure 已有部分 | 逐页未展开 | 读 page-by-page / reassessment |
| 购物逐页 | not_done | final closure 已有部分 | 逐页未展开 | 读 shopping ledger 与 reassessment |
| 政务 / 校务 | not_done | 无 | snapshot closure 未追回 | 单独抽链 |
| 专栏 94 页 | not_done | 无 | 专栏候选 / 拒绝组未追回 | 单独抽链 |
| 浏览器 / 链接 / 语境 | not_done | 无 | 为什么不进 TopHub / Folo / GitHub 尚未完整追回 | 单独抽链 |
| 娱乐 Folo 候选 | paused_by_user | draft 中 partial | 用户要求暂缓单项 closure | 等用户恢复允许 |

---

## 2. 当前比较合格的综合判断

- 主入口优先：同一媒体如果已有横向主入口，不再叠加大量栏目切片。
- 平台注意力源只显影平台如何组织注意力，不替代事实、质量或行动判断。
- 门户聚合源通常只是同一内容池按不同指标或标签反复切片，不进 Folo。
- 官方和调查媒体源只在补出明确互补功能时常驻；否则按任务使用。
- 行动型来源只在具体问题成立时使用：健康回权威证据，住房回具体城市和官方部门，城市服务按现实城市关系，摘要产品按需看一个总入口。
- 任务型入口不制造未读：政策、地区、领域、节目、人物、作品、公司或事件明确时再打开。
- 宽泛平台词、媒体词、栏目词不进追踪器；追踪器只收具体对象。

---

## 3. 当前仍不合格或未完成

- source index 只登记对象，不说明全部判断理由。
- 综合剩余边界文件与候选复查条件未完全追回。
- 科技同族 closure 未追回。
- 社区同族 closure 未追回，且社区热门第 2—20 页仍不能写成全量闭环。
- AI reassessment 六批未追回。
- 开发目录 25 页未追回。
- 财经逐页、购物逐页未追回。
- 政务、校务、专栏、浏览器 / 链接未追回。
- 娱乐 Folo 候选按用户要求暂缓。
- 旧 ChatGPT 对话尚未系统追索。

---

## 4. 合并前禁止说法

合并说明不得写：

- “逐源判断已补齐”；
- “source index 已经补齐判断链”；
- “PR #372 已经完整沉积选择理由”；
- “TopHub 80 + Folo 27 均有完整判断链”；
- “旧 ChatGPT 对话过程已完成迁移”；
- “综合大类已经全量追回”；
- “科技大类已经全量追回”；
- “社区大类已经全量追回”。

当前只能写：

> 已新增 PR #372 判断链追回工作台、总控清单、source index 和覆盖审计，并开始从 PR 文件中追回判断链。当前已覆盖综合多批同族文件，以及科技、社区的第一版 ledger。source index 已登记 TopHub 80 与 Folo 可见来源，但不是判断账本。综合 / 科技 / 社区仍为 partial；开发、专栏、浏览器 / 链接、旧 ChatGPT 对话、娱乐 Folo 候选和多个逐页文件族仍待继续追回。
