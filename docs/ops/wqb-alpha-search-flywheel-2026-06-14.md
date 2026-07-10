# WQB Alpha Search Flywheel｜技术架构与制苗法 v0

日期：2026-06-14  
状态：本地飞轮设计稿  
来源：WQB / CNHKMCP 本地实跑现场

## 0. 一句话

WQB 本地系统不应该只是一个“跑表达式”的脚本，而应该是一个 **受限平台上的 Alpha 制苗系统**：

```text
经济假设 → 字段地形 → 候选苗圃 → 预算闸门 → 异步派发 → 结果回收 → 判决沉积 → 下一轮搜索
```

技术架构的正确位置是：它不是目标本身，而是 Alpha 制作流程的底座。它负责限额、限速、状态、回收、防重复、沉积；Alpha 本身的制作仍然由“经济机制 + 字段语义 + 简单表达 + 扰动验证 + 组合生态位”驱动。

---

## 1. 分层：技术架构放在底座，Alpha 制作放在上层

```text
L6 组合与提交判断层     Portfolio fit / originality / human review
L5 Alpha 判决层         metrics / checks / correlation / verdict
L4 Simulation 回收层    poll / Retry-After / terminal state
L3 Simulation 派发层    dispatch / budget / duplicate guard
L2 候选制苗层           seed → field family → prototype variants
L1 只读地形层           datasets / datafields / docs / operators
L0 安全与状态底座       credentials local / red tools blocked / run-state
```

技术架构主要在 L0–L4；Alpha 制作主要在 L2、L5、L6。  
如果只优化 L0–L4，会变成“更稳定地跑坏表达式”；如果只看 L2，不做 L0–L4，会在 WQB 的限额、限速、异步结果和重复派发里失控。

---

## 2. 平台约束：必须进入设计，而不是事后补丁

WQB 这类平台的真实约束包括：

- simulation 有每日 / 账户 / 活动限额。
- simulation 是异步任务，1–10+ 分钟完成都正常。
- `/simulations/{id}` 会返回 `progress`、`COMPLETE`、`ERROR` 等不同形态。
- API 会给 `Retry-After`，轮询必须尊重。
- MCP call 可能 timeout，但 simulation 已经创建成功。
- 同一 ticket 可能因 timeout 被重复 dispatch。
- 字段类型复杂：MATRIX / VECTOR / EVENT-like / GROUP-like，表达式模板必须按类型生成。
- `submit_alpha`、`set_alpha_properties`、凭证导出等属于红线，不能进入默认飞轮。

因此系统必须有：

```text
Budget Manager
State Ledger
Dispatch Queue
Poll Queue
Verdict Queue
Duplicate Guard
Circuit Breaker
```

---

## 3. 正确状态机

```text
candidate:draft
  → ticket:draft-not-authorized
  → ticket:authorized-once
  → simulation:dispatched
  → simulation:polling
  → simulation:complete | simulation:error | simulation:stalled
  → alpha:readback
  → verdict:bad | verdict:promising | verdict:needs_variant | verdict:discard
  → terrain:update
```

关键原则：**dispatch 与 poll 必须解耦**。

- `dispatch-sim`：只负责创建 simulation，拿到 simulation id / Location 后立即落盘。
- `poll-sim`：独立按 `/simulations/{id}` 回收，尊重 `Retry-After`。
- `verdict-alpha`：独立读取 alpha details / submission check，写判决。

---

## 4. 目录结构

```text
runs/wqb-.../
  run-state.json
  budget.json

  read/
    datasets.json
    datafields/
      forum_sentiment.json
      insider_feats.json

  candidates/
    A1v1-forum.raw.json
    B1v3-invert.raw.json

  tickets/
    A1v1-forum/
      simulation-ticket.json
      simulation-ticket.md
    B1v3-invert/
      simulation-ticket.json
      simulation-ticket.md

  dispatches/
    A1v1-forum.dispatch.json

  polls/
    A1v1-forum.poll.jsonl

  alphas/
    E5K2lz70.details.json
    E5K2lz70.submission-check.json

  verdicts/
    B1.json
    B1v2.json
    A1v1-forum.json

  reports/
    leaderboard.md
    daily-summary.md
```

禁止继续使用单一根目录文件作为事实源：

```text
simulation-ticket.json
get_alpha_details.json
simulation-created.json
```

这些会被覆盖，不能支撑批量或多候选。

---

## 5. Alpha 本身如何制作：六亭法嵌入飞轮

Alpha 制作不是“字段 + operator 拼装”，而是六亭制苗：

```text
Seed：经济假设
Prototype：简单表达
Backtest：初筛
Robustness：扰动
Originality：原创度 / 低相关
Portfolio Fit：组合生态位
```

### 5.1 Seed：先写经济假设

每个 candidate 必须回答：

```text
谁在什么时候，因为什么原因，没有充分定价什么信息？
这个信息预计在什么时间尺度上反映到价格？
什么结果会反证这个假设？
```

例：forum sentiment 不是“情绪字段可能有用”，而是：

```text
散户 / 论坛参与者对公司事件的买卖倾向变化，可能在短期内反映注意力与情绪冲击。
若情绪强但金融词汇权重低，可能只是噪声；若情绪被参与人数和文本质量确认，信号质量更高。
```

### 5.2 Prototype：表达式只做探针

第一版 Alpha 是探针，不是最终产品。要求：

- operator 少。
- nesting 浅。
- 一个数据族为主。
- 不让效果依赖复杂 neutralization。
- 字段类型合法。

### 5.3 Backtest：初筛不是证明

初筛只决定“是否值得继续”。它不能证明未来有效。  
判据包括：Sharpe、Fitness、Turnover、Drawdown、Margin、Sub-universe、Risk-neutralized 表现。

### 5.4 Robustness：扰动测试

一个 promising alpha 至少要经受：

- decay 变化。
- neutralization 变化。
- truncation 变化。
- horizon 变化。
- universe / region 变化（如果平台允许）。
- sign flip 诊断。

### 5.5 Originality：边际信息

原创不是“字段没人用”而已，还包括：

- 与自己旧 alpha 低相关。
- 与常见生产 alpha 低相关。
- 数据机制不同。
- 时间尺度不同。
- 对组合有边际贡献。

### 5.6 Portfolio Fit：生态位

最后不是单 alpha 分数，而是：

```text
这个 alpha 是否补了组合的洞？
是否只是另一个拥挤风格？
是否在风险中性后仍保留信息？
```

---

## 6. 字段类型 → 表达模板

字段类型决定表达式模板，不能硬套 `ts_mean`。

### MATRIX

```text
ts_mean(field, n)
ts_rank(field, n)
delta(field, n)
rank(field)
```

### VECTOR / EVENT-like

先聚合，再时间序列：

```text
vec_avg(field)
vec_sum(field)
vec_count(field)
ts_mean(vec_avg(field), n)
```

B1 的失败说明：

```text
ts_mean(buy_sell_ratio_all_5d_filled, 5)
```

会报：

```text
Operator ts_mean does not support event inputs.
```

修正为：

```text
ts_mean(vec_avg(buy_sell_ratio_all_5d_filled), 5)
```

### COUNT / VOLUME

```text
log(1 + x)
ts_zscore(x, n)
ts_delta(x, n)
```

### SENTIMENT

```text
positive - negative
sentiment * quality
sentiment / attention
sentiment_change
```

### FORECAST

```text
forecast - today
short_forecast - long_forecast
forecast / realized
```

---

## 7. 候选生成：按 family 而不是单表达式

每个 dataset 先拆 signal family：

### forum_sentiment / equity_forum_data

```text
family 1: buy - sell pressure
family 2: sentiment × finance-term quality
family 3: sentiment / attention volume
family 4: sentiment change
```

### insider_feats / insider_matrix

```text
family 1: short window vs long window
family 2: top insider vs all insider
family 3: transaction count vs value
family 4: sign inversion diagnostic
```

### web_traffic_engage

```text
family 1: forecast - today
family 2: 7d forecast vs 28d forecast
family 3: visit growth vs bounce deterioration
family 4: page views / session time attention quality
```

### order_flow_imb / order_book_imbalance

```text
family 1: customer bullish - bearish
family 2: broker/dealer imbalance
family 3: ask/bid pressure
family 4: auction imbalance
```

每个 family 只先跑一个代表苗。通过后才扩参数。

---

## 8. 搜索调度：有限预算下的多臂老虎机

把每个 `dataset + family` 当成一个 arm：

```text
arm = forum_sentiment / sentiment_quality
arm = insider_feats / short_vs_long
arm = web_traffic_engage / forecast_change
```

每个 arm 的分数：

```text
arm_score =
  best_sharpe
  + best_fitness
  + pass_rate
  + economic_plausibility
  - simulation_cost
  - recent_failures
  - crowding_penalty
```

调度比例：

```text
70% exploit：扩展当前最好 family
20% explore：探索新 dataset / 新 family
10% diagnostic：invert / decay / neutralization 小诊断
```

这样既不会盲目批量，也不会手工陷在一个表达式里。

---

## 9. 判决树

```text
if expression_error:
  修字段类型模板，不计入 alpha 失败

elif sharpe < -0.5 and fitness < 0:
  允许一次 invert diagnostic

elif abs(sharpe) < 0.2:
  family 降权，换 dataset

elif 0.2 <= sharpe < 0.8:
  只允许 1–2 个参数扰动

elif sharpe >= 0.8 and fitness low:
  调 turnover / decay / neutralization

elif sharpe >= 1.2 and checks mostly pass:
  做 correlation / submission-check read-only

elif passes submission check:
  标记 human-review-required
  仍不自动 submit
```

---

## 10. 今日现场沉积

已验证：

- 本地 `cnhkmcp` direct `platform_functions.py` 可用。
- `get_datasets` / `get_datafields` read-only 可用。
- raw `get_datafields` 参数是 `dataset_id`。
- `create_simulation` raw 参数应使用 `regular`，不是 `expression`。
- `/simulations/{id}` 是正确 readback endpoint。
- `Retry-After` 应驱动 polling。
- B1 报类型错误：event input 不能直接 `ts_mean`。
- B1v2 成功创建 alpha `E5K2lz70`，但指标较差：Sharpe -0.13，Fitness -0.02，风险中性 Sharpe -0.81。

结论：技术飞轮跑通；Alpha 制作层要从单表达式调参转向 “经济假设 → family → 探针 → 判决 → 地形更新”。

---

## 11. 下一步工程动作

- [ ] `ticket-sim --candidate-id`：ticket 不再覆盖。
- [ ] `dispatch-sim --candidate-id`：创建即落盘，不等完成。
- [ ] `poll-sim --candidate-id`：尊重 `Retry-After` 回收。
- [ ] `verdict-alpha --candidate-id`：统一抽 metrics / checks。
- [ ] `budget-status`：显示今日 simulation 预算。
- [ ] `queue-status`：显示 draft / ready / dispatched / polling / verdict。
- [ ] `score-datasets`：按低拥挤、高 value、字段多样性给 dataset 排序。
- [ ] `generate-candidates`：按字段类型和 signal family 出候选。

---

## 12. 边界

- 不自动 submit。
- 不自动 set alpha properties。
- 不导出凭证。
- 不把凭证 / token / cookie 写进 Git。
- 不做无授权批量 simulation。
- 所有 yellow simulation 必须有 ticket 和本地显式授权。
