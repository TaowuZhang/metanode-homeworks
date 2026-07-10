# WQB · 下一轮投喂包 · 2026-06-21

> 角色：signal / method（Planner 自治产出，非谷拍板项）
> 取数：WQB live catalog 实时拉取（auth 201、token 在），region=USA / delay=1 / universe=TOP3000
> 不吃旧记忆：本包数据为 2026-06-21 现拉，非历史快照

## 0 · 边界与诚实标注

- **状态分析包（当前 sims / alphas / 亭子进度）不在此包**：它在本地 SQLite，属「需外部执行介质」。云端 WQB MCP 只有 `auth_status` / `get_datasets` / `get_datafields`，取不到 simulation 状态。判断与封装在场内完成，物理执行不甩本地。
- **限流缺口**：`get_datasets(search="sentiment")` 撞 429，未重试硬刷。sentiment 类数据集通过 `news` / `insider` 搜索旁路补到（见下表），但 sentiment 主题全集本轮未拉全。
- **本包只产候选与排序**：不下单 simulate、不写本地、不动 alpha 库。

## 1 · 选择判准（可证伪）

排序 = **未拥挤优先**，服务亭子阶段「抢首发 alpha」：

1. 高 `valueScore`（WQB 价值分，越高越稀），且
2. 低拥挤（`userCount` / `alphaCount` 趋近 0），且
3. `pyramidMultiplier` 高（出 alpha 的计分倍率），
4. 兼顾 `coverage`（覆盖股票比例，过低则 IS 样本薄）。

**证伪条件**：若某数据集 `userCount`/`alphaCount` 已高却仍排 Tier 1，即选择错误；若 `coverage < 0.6` 未标注样本薄风险，即标注失职。

## 2 · 投喂排序（live catalog，2026-06-21 拉取）

### Tier 1 — 未拥挤 + 高价值（优先投）

| dataset id | 名称 | valueScore | fields | alphas | users | pyramid | coverage | 备注 |
|---|---|---|---|---|---|---|---|---|
| `insider_feats` | Global Insider Transaction Features | 8 | 118 | 0 | 0 | 1.5 | 0.59 | 价值分最高、零拥挤、字段面大；**coverage 0.59 偏低→IS 样本薄，需在 sim 设置里留意** |
| `news_sentiment_nlp` | News Sentiment Signal Features | 7 | 23 | 0 | 0 | 1.5 | 0.995 | 零拥挤 + 覆盖近满；全新未验证，信号可提取性待 sim 证 |
| `insider_agg_matrix` | Smart Insider Transaction Aggregates | 7 | 17 | 0 | 0 | 1.2 | 0.77 | 零拥挤、矩阵态、字段少→表达式好上手 |

### Tier 2 — 低拥挤 + 1.5x 倍率

| dataset id | 名称 | valueScore | fields | alphas | users | pyramid | coverage |
|---|---|---|---|---|---|---|---|
| `news_sentiment_transfer` | Transferred News Sentiment Analytics | 6 | 23 | 4 | 2 | 1.5 | 0.95 |
| `insider_trx_matrix` | Global Insider Transaction Aggregates | 7 | 29 | 19 | 11 | 1.5 | 0.77 |
| `insider_matrix` | Aggregated Insider Transaction Matrix | 6 | 33 | 53 | 19 | 1.5 | 0.77 |
| `board_gov_stats` | US Board Governance & Leadership Metrics | 6 | 46 | 19 | 11 | 1.2 | 0.92 |

### Tier 3 — 宽字段面但较拥挤（晚投/作组合）

| dataset id | 名称 | valueScore | fields | alphas | users | pyramid | coverage |
|---|---|---|---|---|---|---|---|
| `news_transformer_scores` | Transformer Based News Sentiment Scores | 5 | 315 | 126 | 57 | 1.5 | 0.78 |
| `sentiment22` | News Sentiment Scores | 4 | 210 | 360 | 206 | 1.5 | 1.00 |
| `news97` | Web News Sentiment Scores | 4 | 105 | 191 | 108 | 1.2 | 0.98 |

## 3 · 下一轮具体动作

1. **字段勘探**：对 Tier 1 三只调 `get_datafields(datasetId=...)` 取真实可用字段，再拟因子表达式。（本包未拉字段，避免再撞 429。）
2. **首投**：`insider_feats` + `insider_agg_matrix`（同主题 insider，可共用表达式骨架，零拥挤），各拟 2~3 个表达式。
3. **对照**：`news_sentiment_nlp` 作 sentiment 侧首投，验证「零 alpha 数据集是否真能出信号」这一假设。
4. **组合留后**：Tier 3 宽字段面适合做 group/vector 操作，但拥挤度高，留到首发轮之后。

## 4 · 交接

- **X 自治可推进**：字段勘探（`get_datafields`）、因子表达式草拟、候选清单维护。
- **需外部执行介质（本地）**：`simulate` 下单、写 alpha 库、跑 `wqb_loop`、读本地 SQLite 当前进度。
- **需谷拍**：无——本包无不可约价值对冲。
