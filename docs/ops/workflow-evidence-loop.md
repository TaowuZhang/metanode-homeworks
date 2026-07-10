# Workflow evidence loop

状态：v0，先在 GetNote mirror 上试跑。

来源问题：论文《Investigating Autonomous Agent Contributions in the Wild: Activity Patterns and Code Change over Time》（arXiv:2604.00917）能观察代码在合并后是否继续变化，却无法知道代码是否真正运行、后续修改是否关闭了真实问题。沃壤不再另建逐 PR 台账，而用现有 workflow 补上这段时间证据。

## 一、基本判断

PR、测试、workflow 与真实使用回答不同问题：

| 证据 | 最多能说明什么 |
|---|---|
| PR 合并 | 当前变化被接受进入仓库 |
| CI 通过 | 已知内部检查在该 SHA 上成立 |
| dry-run 通过 | 排练路径成立，未必接触真实外部输入 |
| 真实 workflow 通过 | 某次真实输入与边界检查成立 |
| 真实 workflow 失败 | 出现反例，或出现执行/基础设施缺口 |
| 修复后同一路径重新通过 | 完成一次可观察的纠正闭环 |

没有修改不等于有效；没有运行只表示尚未暴露。高 churn 也不自动等于低质量，必须看变化是否对应真实反例、环境变化或计划内演化。

## 二、两个时间锚点

### Last attempt

最近一次 workflow 尝试。它可以成功、失败、取消、未启动或仍在运行。

它回答：**最近真实发生了什么？**

### Last healthy

最近一次通过验收并被允许进入 `main` 的真实运行证据。

它回答：**当前仓库最后接受的是哪一次现实结果？**

失败、降级或未真正执行的 attempt 不覆盖 last healthy。

## 三、GetNote 第一条闭环

GetNote 是确定性 recent-window 镜像，正常运行无需每次开 PR：

```text
真实 GetNote 读取
→ 生成镜像与 observability
→ evaluate-getnote-attempt
→ healthy：直接更新 main，刷新 last-attempt 与 last-healthy
→ degraded：不写 main，只上传短期诊断 artifact，last-healthy 保持不动
```

验收分成硬失败与 advisory：

- 数量不一致、坏 ID、空无效记录、路径碰撞、已出现 sentinel 却结构不匹配：硬失败；
- 固定 sentinel 不在滚动 recent window：advisory；
- 有身份和有效内容的无标题记录：advisory，不单独判 parser 失败。

## 四、修复什么时候算闭环

一次后续修改只有满足下面三步，才可说关闭了真实问题：

1. 修改前有可观察的失败或降级；
2. 修改针对该失败面，而不是只在附近改代码；
3. 修改后同一条真实 workflow 路径重新通过。

修复已合并但尚未重新运行，状态是“待复验”，不是“已有效”。

## 五、每周只浮异常

`weekly-governance-report.yml` 调用 `scripts/workflow-evidence-inspector.mjs`，比较 GitHub Actions 的 last attempt 与仓库中的 last healthy。

当前状态包括：

- `healthy-current`：最新尝试就是最后健康证据；
- `counterexample-after-last-healthy`：最后健康之后出现真实失败；
- `execution-gap-after-last-healthy`：startup failure、取消或超时，没有形成有效执行证据；
- `successful-attempt-not-recorded-as-healthy`：workflow 成功，但证据没有进入 main；
- `unexposed`：看不到真实运行。

正常状态保持安静；治理报告只负责显影证据缺口，不给代码或代理打质量分。

## 六、边界

- 不为历史三百多个 PR 回填台账；从仍在运行的 workflow 开始。
- 不把 workflow success 当成永久正确。
- 不把 infra/startup failure 误记成代码反例。
- 不自动判定责任归属；机器整理证据，人判断必要演化、环境变化或可避免返工。
- 不让监测系统反过来成为新的登记机关；新 workflow 只有在存在真实运行路径时才接入。
