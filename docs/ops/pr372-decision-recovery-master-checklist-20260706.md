# PR #372 判断链追回总控清单（2026-07-06）

## 0. 总目标

把 PR #372 从“结果、闭环、摘要和操作残留”整理成可审查的判断资产。

最终应能回答：

- TopHub 里为什么是这些节点；
- Folo 里为什么是这些来源；
- 哪些只放沃壤、浏览器、链接、语境或按需；
- 为什么同页 / 同类 / 相近来源没有进入；
- 哪些判断已经追回，哪些还没追回；
- 哪些文件只是操作流水，不能当判断证据。

未打勾的部分，不得在 PR 描述中宣称完成。

---

## 1. 判断链最低格式

```yaml
- source_id:
  current_route:
  decision_status:
  selected_reason:
  rejected_alternatives:
  route_reason:
  evidence_locator:
  next_action:
```

不合格条目：只有结论、只有数量、只有“重复 / 噪声 / 停更”、没有拒绝项、没有 evidence locator、用当前模型补编旧判断。

---

## 2. 已落入 PR 的控制 / ledger 文件

- [x] `docs/ops/pr372-decision-recovery-todo-20260706.md`
- [x] `docs/ops/pr372-decision-recovery-master-checklist-20260706.md`
- [x] `docs/ops/pr372-decision-recovery-coverage-audit-20260706.md`
- [x] `docs/ops/pr372-decision-recovery-source-index-20260706.md`
  - TopHub 80 项、Folo 可见 20 项、Folo 省略 7 项、Folo 替换 / 移出来源已登记。
  - 它是母表，不是判断账本。
- [x] `docs/ops/pr372-decision-recovery-ledger-draft-20260706.md`
  - AI / 设计 / 购物 / 财经 / 报刊 / 娱乐部分来源。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-20260706.md`
  - 综合 closeout 三项真实动作、Folo / 追踪器 / 通知机器人分工、知乎 8 节点。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-hotsearch-rolling-20260706.md`
  - 综合热搜、滚动新闻。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-platform-attention-20260706.md`
  - 微信、微博、今日头条。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-portal-aggregators-20260706.md`
  - 百度、腾讯、网易、搜狐、新浪、ZAKER。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-institutional-media-20260706.md`
  - 人民网、新华社、南方周末、半月谈、CCTV。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-comprehensive-action-sources-20260706.md`
  - 健康、房产、本地宝、腾讯日报。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-technology-20260706.md`
  - 科技 final closeout 第一版。
  - 状态：partial。
- [x] `docs/ops/pr372-decision-recovery-ledger-community-20260706.md`
  - 社区 pending actions 第一版。
  - 状态：partial。

---

## 3. 当前综合大类已经追回的复用规则

- 热搜类：不叠加同类热搜榜；搜索引擎热搜按需；平台热搜不重复；热搜不进 Folo、不建宽泛追踪器。
- 滚动新闻类：不把快讯流搬进 Folo；高重叠滚动源按需；地方、赛事、标签不稳来源按任务处理或排除。
- 平台注意力类：平台切片不等于独立来源；领域标签不能替代证据等级；真实公众号整理另开；平台榜单不进 Folo、不建宽泛追踪器。
- 门户聚合类：点击、评论、滚动、视频、热榜、频道标签通常只是同一内容池的不同切片；独立栏目需看更新证据、原始来源、真实阅读价值与负荷。
- 官方与调查媒体类：保留主入口和互补功能；政策、地区、人事、知识产权、法治等按具体问题核验；调查型节目只有补出新结构功能时才扩容。
- 行动型来源类：健康、住房、城市服务和日报摘要只在具体任务成立时有价值；不自动变成持续阅读关系。

---

## 4. 当前待办队列

按顺序执行，不跳。

1. [x] 生成 source index。
2. [ ] 继续补综合剩余同族 closure 与边界文件。
3. [ ] 继续补科技同族 closure。
4. [ ] 继续补社区同族 closure。
5. [ ] 抽开发目录 25 页判断链。
6. [ ] 补 AI reassessment 六批的让位关系。
7. [ ] 抽政务 / 校务 snapshot 判断链。
8. [ ] 抽专栏 94 页候选与拒绝组。
9. [ ] 抽浏览器 / 链接 / 平台分工判断链。
10. [ ] 回旧 ChatGPT 对话追索所有 partial 条目。
11. [ ] 更新 coverage audit。
12. [ ] 改 PR 描述，明确“已追回 / 未追回 / 不得冒充”。

综合剩余重点：directory review / rolling review 边界、部分候选的 Folo 复查条件，以及若仍有不适合在总控展开的敏感类别，用低展开度方式记录边界。

---

## 5. 给本地 Codex 的最小交接边界

如果需要本地 Codex 接手，它只做文件操作，不重新理解上下文：

1. 读取本文件。
2. 读取 `coverage audit`。
3. 读取 `source index`。
4. 按第 4 节顺序，一个模块一个 ledger。
5. 每个 ledger 只从 PR #372 现有文件抽链，不许重审网页，不许新编理由。
6. 每完成一个 ledger，更新本文件和 coverage audit。
7. 不改真实订阅账本，不合并 PR，不 Ready for Review。
