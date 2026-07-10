# WQB 50+ 赞精品源整理与明早接续地图 — 2026-06-23

## 0. 本文件的目的

用户要求：睡前不再继续 simulation / submit，只把当前已经抓到的 50+ 赞精品材料处理到足以让明早可以设计一个更好的系统。

本文件做三件事：

1. 对本次 full discovery 的覆盖质量做统计判断。
2. 对 50+ 赞 must-read 源做分层、聚类、可信度和用途划分。
3. 把这些源转成明早可直接使用的系统设计输入，而不是继续线性“派发-回收-打分”。

安全边界：

- 不存论坛正文。
- 不存 cookie / token / password / session。
- 不存 private alpha expression。
- 不复制表达式。
- 不做 WQB submit。
- 不派发 simulation。

---

## 1. 发现与读取状态

来源：

- `runs/wqb-community/premium-source-full-discovery-latest.json`
- `runs/wqb-community/premium-source-full-discovery-queue.json`
- `runs/wqb-community/premium-source-full-discovery-claims.json`
- `runs/wqb-current/forum-claims-integration-summary.json`

本次可达面发现：

| 指标 | 值 | 解释 |
|---|---:|---|
| unique_posts_discovered | 418 | CNHKMCP search fallback 可达帖子数 |
| must_read_count | 52 | vote_count >= 50 的 P0 帖子 |
| high_value_count | 108 | vote_count 20–49 的 P1 帖子 |
| claim extraction attempted | 100 | 已读取 queue 前 100 条，覆盖全部 52 个 P0 |
| read_status=read | 100 | 全部读取成功；但不代表正文都被 extractor 有效提炼 |
| usable_claim_count | 95 | integrator 接收的 claim 数 |
| blocked_or_deferred_claim_count | 5 | 弱/空/低置信 claim |

覆盖限制：

- `list_forum_posts_in_topic` 被 Support auth / Cloudflare 403 阻断。
- 因此这不是数学意义上的全论坛历史全量枚举。
- 但 search fallback 覆盖了 418 条可达结果，且把当前可达面内 52 个 50+ 赞帖子全部纳入了读取队列。

准确表述：

> 当前 CNHKMCP search fallback 可达面内，50+ 赞规则已闭环；但因 topic list route 被 Cloudflare 403 阻断，不能声称全 Zendesk 历史的数学完备覆盖。

---

## 2. Claim 集成统计

集成后的 failure_mode_counts：

| failure mode | count | 解释 |
|---|---:|---|
| SIMULATION_BUDGET_WASTE | 7 | 社区高票源强烈支持反无效回测、反盲目批量 |
| LOW_SUB_UNIVERSE_SHARPE | 4 | 子宇宙/稳定性是核心系统约束 |
| OVERFITTING | 4 | 需要稳健性和 regime 诊断 |
| HIGH_TURNOVER | 3 | turnover 修复需结构化 entry/exit，不是 decay-only |
| LOW_FITNESS | 3 | Fitness 不能只看 Sharpe；要拆收益/turnover/coverage |
| SOURCE_DRIFT | 3 | 工具/流程帖必须和 alpha-design 分离 |
| CONCENTRATED_WEIGHT | 3 | coverage / breadth / max-weight 需要拆解 |
| PROD_CORRELATION | 1 | search/claim 层低估了 prod_corr；但 RA 检测帖证据很强 |
| SELF_CORRELATION | 1 | 同上 |
| LOW_2Y_SHARPE | 1 | 子宇宙和过拟合源间接覆盖更多 |

解释：

- 这些 counts 不是“社区真正关注度”的完整统计；它们受 deterministic extractor 模板影响。
- 真正有用的是：高票源把系统设计方向从“局部修表达式”推向“反浪费流程、源分层、稳健性、turnover 结构、数据集评估、表达式防复制”。

---

## 3. 50+ 赞源分层总览

### A. Alpha-design 强相关源：明早优先消化

这些源最可能直接改变 Batch F 的设计。

| vote | post_id | title | 初步用途 | 明早动作 |
|---:|---|---|---|---|
| 313 | 32226888249239 | 【新人指南】到底要交什么样的Alpha？ | submission target / alpha quality standard | 提炼“什么值得交”的 gate-aware objective |
| 233 | 26054361848343 | 可以尝试使用的Alpha模板 | 模板/表达式风险极高，但可抽象 operator grammar | 只抽象结构，不复制表达式 |
| 212 | 37266277327767 | 【柯楠】基于操作符的因子构建实战与说明（1） | operator grammar / factor construction | 变成 Experiment Factory grammar |
| 210 | 34696235484567 | 【柯楠】一文带你了解alpha表达式的经济学含义-第二期 | economic meaning / hypothesis discipline | 引入“经济含义必须先于算子堆叠” |
| 167 | 30927669645207 | 如何拯救高turnover因子？ | turnover repair | 设计 entry / hold / exit / trade_when grammar |
| 136 | 19137620283415 | ⭐【Alpha灵感】股票收益是球队还是硬币？ | return model metaphor / cross-sectional logic | 作为 price-volume / relative-strength scout prior |
| 125 | 36947868698519 | RA 的 Prod 检测 24h 可检测600个 | prod_corr / RA 检测 / workflow | 只抽象检测与预筛思想；不做 bulk mining |
| 122 | 18600255773719 | 【Alpha灵感】A股换手率类因子 | turnover signal family | 作为 turnover-signal scout，不直接套 A 股假设 |
| 119 | 32034293019671 | 组sa时如何选取高质量因子？ | SA/portfolio selection | 转成 candidate pool quality filters |
| 108 | 35928620962839 | Combined Alpha Performance 计算原理 | portfolio / combined score | 接入 gate-aware portfolio scheduler |
| 101 | 37483774843671 | 混信号 / Single Data 底层逻辑 | source purity / single-data discipline | 防止无动机混信号过拟合 |
| 99 | 35377811169175 | SuperAlpha Not Own Selection 分层框架 | SA selection / hierarchy | method-only, no expression copy |
| 91 | 24897131810839 | 论文合集 | research prior library | 建立 research-paper scout queue |
| 89 | 28050210195863 | GLOBAL 数据处理方法 | global data processing | 仅作 cross-region/data-processing method prior |
| 87 | 29327820720151 | A股恐慌度因子在美股的应用 | sentiment / fear transfer | 作为 cross-market transfer scout，需强降权 |
| 67 | 34949059814679 | MCP Workflow turnover 优化 | workflow + turnover | process-only + turnover grammar |
| 62 | 30413891877655 | 豆包推荐用户数据集1 | dataset scout | 数据集评估，不直接仿写 |
| 61 | 30668648522519 | 存货周转率疑义探讨 | fundamental turnover ambiguity | field meaning / accounting sanity check |
| 59 | 29085671898775 | Operator 大师 | operator composition | grammar library / complexity penalty |
| 56 | 19353819839639 | 新闻动量反转策略 | news / momentum / reversal | sentiment-news hypothesis prior |
| 55 | 23840950088855 | 基本面量化之回归算子 | regression/fundamental | regression operator family |
| 51 | 19139902140439 | 基于量价信息的可靠投资信号 | price-volume signal | price-volume scout prior |

### B. 流程 / 反浪费 / 工具源：系统基础设施，不直接产 alpha

| vote | post_id | title | 用途 | 明早动作 |
|---:|---|---|---|---|
| 402 | 22863075241623 | BRAIN 平台自学路径图推荐 | 学习路径/知识结构 | 转成 curriculum graph，不进 Batch F |
| 181 | 30446782096407 | Python 小白搞懂挖掘过程 | 本地流程认知 | 提炼流程，不替代 WQB policy |
| 160 | 39200810917015 | IQC Course 1 作业 | 教学/作业 | 只作学习材料，不进 simulation |
| 125 | 36947868698519 | RA Prod 检测 | 预筛 / workflow | 允许反浪费思想；禁止 bulk mining |
| 107 | 35220740797975 | 远程控制电脑 | 本地执行环境 | 工程参考，不进 alpha design |
| 88 | 39467969845783 | 新手课代码分享 | 代码流程 | 不复制，抽象流程 |
| 67 | 34949059814679 | MCP Workflow turnover 优化 | workflow + turnover | process-only, guardrails |
| 54 | 25326918974871 | 批量 Alpha 挖掘思路分享 | batch workflow | 强 anti-bulk-mining 约束 |

### C. Policy / consultant / boundary 源：约束系统，不驱动表达式

| vote | post_id | title | 用途 |
|---:|---|---|---|
| 116 | 28584425090455 | BRAIN 量化研究顾问申请流程 | consultant / application context |
| 107 | 28611992884631 | 顾问收入计算 | incentive / value-factor context |
| 84 | 40373406402455 | Consultant 项目定位、预期与行为边界 | 行为边界 / policy |
| 78 | 28857003916055 | Reflecting on 2024 and Welcoming 2025 | community reflection |
| 66 | 39645863500055 | IQC FAQ 与进阶指南 | competition context |
| 61 | 35133725170071 | IQC 转正式顾问后想法 | transition context |
| 55 | 33785558468887 | 写给还在朝 consultant 奋斗的新人 | process / career |
| 53 | 33037839973783 | Path to Genius 大纲 | learning path / program structure |

### D. 学习 / evergreen / 高票但不直接转策略

| vote | post_id | title | 处理方式 |
|---:|---|---|---|
| 72 | 15152019662487 | BRAIN 小贴士 | evergreen tips; 需二次提炼 |
| 73 | 30392913209367 | WorldQuant 新手攻略 | FAQ / mistakes |
| 71 | 31681360197527 | 系统性风险暴露 / Sharpe | metric education |
| 71 | 36770332492951 | 新顾问手搓模板三部曲 | template risk / method-only |
| 82 | 34836679909271 | 量化学习书籍推荐 | external reading map |
| 84 | 30684190103319 | VF 成长经验 | process / pitfalls |

### E. SA / Combine / Portfolio 源：不做表达式复制，只做组合与筛选思想

| vote | post_id | title | 用途 |
|---:|---|---|---|
| 119 | 32034293019671 | 组 SA 时如何选高质量因子 | candidate-pool quality score |
| 108 | 35928620962839 | Combined Alpha Performance | combined metric decomposition |
| 105 | 36116437735447 | 连续12天 pc<0.3 SA 体验 | low-correlation workflow / caution |
| 99 | 35377811169175 | Not Own Selection 分层框架 | layered selection framework |
| 54 | 36412349216279 | Combine 成功提升 | portfolio combine method-only |

### F. Alpha idea / research-paper 源：作为 scout priors，不直接上手抄

| vote | post_id | title | 可能映射 |
|---:|---|---|---|
| 315 | 19273239621399 | Alpha 灵感启示录合集 | idea library / high priority re-extraction |
| 136 | 19137620283415 | 股票收益是球队还是硬币 | cross-sectional return model |
| 122 | 18600255773719 | A股换手率类因子 | turnover / liquidity hypothesis |
| 91 | 24897131810839 | 论文合集 | research-paper queue |
| 87 | 29327820720151 | 恐慌度因子在美股应用 | sentiment/fear transfer |
| 62 | 30413891877655 | 用户数据集1 | dataset scout |
| 61 | 30668648522519 | 存货周转率疑义 | accounting/fundamental sanity |
| 56 | 19353819839639 | 新闻动量反转 | news momentum/reversal |
| 55 | 23840950088855 | 基本面回归算子 | regression operator family |
| 51 | 19139902140439 | 量价可靠投资信号 | price-volume scout |

---

## 4. 质量审查：哪些真的“读懂了”，哪些只是“读过了”

本轮最大收获不是 claim 数量，而是发现了 claim extractor 的瓶颈。

### 4.1 强证据源

特点：`body_chars_seen_in_memory > 0` 且 keyword hits 或 title-topic 能明确支撑设计。

代表：

- `forum:36947868698519`：Prod / correlation 相关词命中极高：prod_corr 94、correlation 101、self 56、相关 23。
- `forum:34949059814679`：MCP workflow / turnover 优化，workflow 5、mcp 55、工具 24。
- `forum:30413891877655`、`forum:30413913633687`：dataset scout / dataset evaluation。
- `forum:30668648522519`、`forum:31002256151831`：turnover repair。
- `forum:31365517513495`：sub-universe / hump 相关。

### 4.2 中证据源

特点：正文长度可观，但 deterministic extractor 给了 generic claim。

代表：

- `forum:26054361848343` Alpha 模板：31k chars，但 generic；需要 method-only 二次提炼。
- `forum:37266277327767` 操作符构建：10k chars，但 generic；应重提 operator grammar。
- `forum:34696235484567` 经济学含义：12k chars，但 generic；应重提 hypothesis discipline。
- `forum:35928620962839` Combined Alpha Performance：40k chars，但 generic；应重提 score decomposition。
- `forum:15152019662487` BRAIN 小贴士：32k chars，但 generic；应重提 evergreen rulebook。

### 4.3 弱证据 / metadata-only 源

特点：`body_chars_seen_in_memory = 0`，只能证明 metadata 和标题，不足以改变 scheduler。

代表：

- `forum:22863075241623` BRAIN 自学路径图，402 赞，但正文未提炼。
- `forum:19273239621399` Alpha 灵感启示录合集，315 赞，但正文未提炼。
- `forum:32226888249239` 到底要交什么样的 Alpha，313 赞，但正文未提炼。
- `forum:30927669645207` 高 turnover 因子，167 赞，但本次 body chars 为 0；虽然 topic strong，但需要重读。
- `forum:39200810917015` IQC Course 1 作业，160 赞，metadata-only。

明早原则：

> 不能把“read_status=read”当成“已经理解”。必须按 evidence strength 分层。

---

## 5. 明早系统设计应吸收的核心思想

### 5.1 从“表达式搜索”升级为“源证据驱动的假设图”

每个候选不再只是一个 FASTEXPR，而是：

```text
Source claim → Mechanism hypothesis → Field family → Operator grammar → Expected gate effect → Kill rule
```

如果没有 source claim，不进入 Batch。

### 5.2 三类源必须分开

1. **Alpha-design 源**：可以影响 field/operator/neutralization/turnover grammar。
2. **Process/workflow 源**：只能影响 executor/recovery/anti-waste，不授权 bulk mining。
3. **Policy/community-boundary 源**：只作为 guardrail，不影响 expression。

### 5.3 高票不等于高可执行性

高票源要用两个分数：

```text
source_importance = log(1 + vote_count) + recency + comment_count_signal
execution_relevance = alpha_design_specificity * evidence_strength * safety_score
```

很多 100+ 赞学习/流程帖 source_importance 很高，但 execution_relevance 不一定高。

### 5.4 明早 Batch F 不应直接使用 95 claims

应先过滤成 12–20 个强 claim，再设计 Batch F。

优先强 claim 类型：

- Prod/correlation 检测与机制 pivot
- Turnover structural repair
- Dataset evaluation / field coverage
- Operator grammar / economic meaning
- Single Data / mixed signal discipline
- Combined / SA quality selection

---

## 6. 明早可直接用的系统设计蓝图

### 6.1 Source Intelligence Layer

输入：forum / official / docs / prior alphas。

输出每条 claim：

```json
{
  "source_id": "forum:...",
  "vote_count": 167,
  "source_type": "alpha_design | workflow | policy | learning | portfolio",
  "evidence_strength": 0.0,
  "execution_relevance": 0.0,
  "copy_risk": 0.0,
  "claim": "...",
  "design_effect": "...",
  "allowed_use": "method_only | gate_rule | scout_prior | workflow_guardrail"
}
```

### 6.2 Starmap / Terrain Layer

把 claim 映射到星图轴：

```text
field_family × operator_family × horizon × neutralization × gate_blocker × source_cluster
```

### 6.3 Population Climber Layer

不是随机变异表达式，而是：

```text
seed mechanism → small typed mutations → evidence-backed variants → gate-aware survival → family-level posterior update
```

### 6.4 Explore / Exploit Scheduler

建议明早改成：

- 2 slots：source-backed scout
- 1 slot：prod_corr / originality pivot
- 1 slot：turnover structural diagnostic
- 1 slot：dataset / coverage diagnostic
- 1 slot：holdout/control unless prechecks are complete

### 6.5 Gate-aware Stop Rules

任何 slot 必须有：

- expected_information_gain
- duplicate_hash_precheck
- expression-copying firewall
- specific kill_rule
- retry-after / recovery plan

---

## 7. 明早优先重提炼队列

这些不是让用户现在再跑，而是明早如果要继续，应优先重提取 / 重摘要。

### P0A：最值得二次提炼

1. `32226888249239` — 到底要交什么样的 Alpha？
2. `26054361848343` — 可以尝试使用的 Alpha 模板
3. `37266277327767` — 基于操作符的因子构建
4. `34696235484567` — alpha 表达式经济学含义
5. `30927669645207` — 高 turnover 因子修复
6. `35928620962839` — Combined Alpha Performance
7. `37483774843671` — 混信号 / Single Data
8. `32034293019671` — 组 SA 时如何选高质量因子
9. `35377811169175` — Not Own Selection 分层框架
10. `36947868698519` — RA Prod 检测
11. `15152019662487` — BRAIN 小贴士
12. `19273239621399` — Alpha 灵感启示录合集

### P0B：可作为 scout prior

- `19137620283415` 股票收益是球队还是硬币
- `18600255773719` A股换手率类因子
- `29327820720151` 恐慌度因子
- `30413891877655` 用户数据集1
- `30668648522519` 存货周转率
- `19353819839639` 新闻动量反转
- `23840950088855` 基本面回归算子
- `19139902140439` 量价可靠投资信号

### P0C：只进入 policy / workflow，不进 alpha design

- `22863075241623` 学习路径图
- `28584425090455` 顾问申请流程
- `28611992884631` 顾问收入计算
- `35220740797975` 远程控制电脑
- `40373406402455` 项目定位与行为边界
- `25326918974871` 批量挖掘思路分享

---

## 8. 明早接续判据

明早不应该问“要不要继续跑 Batch F”，而应该先问：

1. 50+ 源里哪些是 alpha-design 强源？
2. 哪些只是 workflow / policy？
3. 哪些 generic claim 需要 stronger extractor？
4. Batch F 的六个 slot 是否各有强 source backing？
5. 是否已经有 duplicate/hash/pre-dispatch gates？

只有当这些成立，才讨论 simulation 授权。

当前 final state：

```text
full_premium_50plus_sources_synthesized_no_simulation
final_click_candidate_found = false
Batch F = proposal-only, blocked pending strong-claim review and pre-dispatch gates
```
