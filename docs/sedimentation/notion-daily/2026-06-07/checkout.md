# Checkout｜下一轮接续｜2026-06-07

克，checkout `docs/sedimentation/notion-daily/2026-06-07/`。

先读：

1. `index.md`
2. `decisions.md`
3. `friction.md`
4. `architecture.md`
5. `cleanup.md`

## 下一轮先做什么

1. 不要再解释“日终沉淀”概念，先验证首包结构是否够用。
2. 回到 Notion AI 共台，判断五张表是否需要补外运状态、GitHub 路径、清理动作、外运类型、外运批次。
3. 如果要继续落地 MVP，先只回写外运状态，不自动删除。
4. 第一周禁止自动清谷原声、模型原话、推论沉积、摩擦记录、全景图结构边。

## 不做什么

- 不把 GitHub 包当 Notion 现场替代品。
- 不把 raw 摘要替代谷原声。
- 不把 E0 / E1 写成规则。
- 不自动合并 PR。
- 不自动删除父节点。
- 不把 Eagle / Steam 等别线债务误当成本线程下一刀。

## 当前下一刀

落 MVP v0.1 的 Notion 侧字段 / 回写机制：

```plain text
AI 共台五张表
→ 外运状态 / GitHub 路径 / 清理动作 / 外运类型 / 外运批次
→ 先标记，不删除
```