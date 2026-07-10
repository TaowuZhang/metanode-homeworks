# Workflow evidence

这里保存 weekly governance 生成的只读异常面。

它比较：

- GitHub Actions 中最近一次真实尝试（last attempt）；
- `main` 中最近一次被接受的健康证据（last healthy）。

当前只接入 GetNote mirror，作为第一条试跑。它不统计所有 PR，不计算代码质量分，也不把“长期没改”当成有效证明。

生成器：`scripts/workflow-evidence-inspector.mjs`

操作与判断口径见：`docs/ops/workflow-evidence-loop.md`。
