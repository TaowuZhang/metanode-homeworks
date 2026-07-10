# WQB community forum index analysis — 2026-06-22

Source: `runs/wqb-community/forum-post-index-local.json`

Mode: metadata-only, read-only. The local export used authenticated CNHKMCP forum list tooling, but retained only post-level metadata: ids, titles, votes, comments, timestamps, topic ids, and URLs. No full post body, credentials, cookies, session data, forum comments, forum votes, or forum posts are stored.

## Result

The minimum community visibility goal is reached:

- `call_ok = true`
- `post_count = 50`
- top-level response keys: `success`, `topic_id`, `posts`, `total_found`

This means the workflow can see the live WQB community list layer: what posts exist, titles, ids, activity timestamps, votes, comment counts, topic ids, and URLs.

## Title clusters

### 1. Platform webinars / competitions

Representative titles:

- `PAC 2026 - Advanced Operator Development Webinar Recording (17 June)`
- `PAC 2026 – Simulation Optimization Webinar recording (11 Jun)`
- `IQC 2026 Webinar Recording (11th June)`
- `PowerPool Eligibility Question`

Implication: recent official/semi-official webinar posts should become high-priority full-read candidates before changing operator or simulation-optimization assumptions.

### 2. Research paper hypotheses

Representative titles:

- `Research Paper: Rational Inattention and Analyst Forecast Accuracy`
- `Research Paper 13: Valuation Ratios, Surprises, Uncertainty or Sentiment: How Does Financial Machine Learning Predict Returns From Earnings Announcements?`
- `Research Paper 12: Stock Recommendations from Stochastic Discounted Cash Flows`

Implication: analyst forecast accuracy, earnings-announcement sentiment, valuation/surprise, and DCF-related ideas are active hypothesis priors for scout slots.

### 3. AI research workflow and infrastructure

Representative titles:

- `【方法论】AI时代Alpha研究全景图：从数据验证到知识沉淀的系统化工作流与实验账本`
- `让 AI 少做无效回测：Alpha 研究里的“反证优先”小批量测试法`
- `基于MCP + claude(codex)的全自动alpha挖掘系统:30分钟Regular alpha直出`
- `把 AI Alpha 研究从“会写公式”推进到“会留证据”：一个轻量实验账本模板`
- `codex流是不是不行，回测几十万都没找到`

Implication: the community is actively discussing ledgers, anti-invalid-backtest discipline, MCP workflows, and the risks of brute-force automation. This supports the nonlinear star-map, duplicate-hash, and recovery-controller direction.

### 4. Python alpha / signal processing

Representative titles:

- `[Python Alpha 挑战赛]使用蒙特卡洛来实现python alpha`
- `【Python Alpha】把alpha丢进傅里叶变换后指标真给整过去了`
- `傅里叶变换在 Alpha 开发中的应用：频率域视角下的信号分解与策略构造`
- `[Python Alpha 挑战赛] Python Alpha 回测太慢？聊聊我的性能优化三板斧`

Implication: Python and signal-processing posts are useful for diagnostics/tooling, but must not turn into autonomous bulk mining.

### 5. Correlation / prod_corr / robustness

Representative titles:

- `频率感知的 Seed 构造 × CW 前置分层诊断 × 多 GEM 降相关：让自动化 Alpha 挖掘在"源头"而非"后处理"解决问题`
- `CHN Python Alpha 踩坑实录：risk72 死磕 prod_corr 无果，换 virgin 数据集才过关`
- `【IND/D1】两段式批量筛选：IS 鲁棒性过滤 → prod_corr 续跑 — 从 self鲁棒性 到 pc回测 的衔接`
- `【IND/D1】批量 prod_corr 断点续跑与结果落盘 — partial CSV + pickle 备份实战`

Implication: prod-correlation is treated as a mechanism/data-family problem, not a small parameter repair. Batch F should favor mechanism/data pivots and robustness staging.

### 6. Gate failure repair methods

Representative titles:

- `筛选可优化RA进阶版`
- `CONCENTRATED WEIGHT 系统性排查与修复手册-2（基于实战经验的 CW 攻防全指南_进阶篇）`
- `降低 Alpha 周转率的核心策略`
- `换个角度，理解一下trade_when(x, y, z) 方法`
- `提高MEA区的Sub-universe Sharpe of 1.11 is below cutoff of 1.14.问题一点经验`

Implication: the next full-post reads should prioritize CW, turnover/trade_when, RA, sub-universe repair, and prod_corr staging.

### 7. Dataset guides and alpha cases

Representative titles:

- `Sentiment1 数据集完全指南：Research Sentiment Data`
- `Model77 数据集完全指南：Analysts' Factor Model`
- `【方法论】BRAIN 数据集系统分析工作流 — 从字段分类到 Alpha 构造方向的标准化路径`
- `跳出1、2、3阶常规，usa_ILLIQUID_MINVOL1M_risk59数据集还能怎么做`
- `Alpha 因子案例（波动率套利）`
- `Alpha 因子案例（现金流估值 + 高估个股识别）`

Implication: dataset-guide posts should influence scout slots and field-family priors, not immediate exploit unless a source claim is extracted from a full read.

## Scheduler effect before full-post reads

Metadata alone is enough to change priorities cautiously:

1. Do not dispatch Batch F yet; first read a small, selected set of community posts that directly map to current blockers.
2. Prioritize full reads for concentrated weight, turnover/trade_when, prod_corr staging, and anti-invalid-backtest methodology.
3. Treat Python/Fourier/automation posts as diagnostic/tooling priors, not as copyable alpha expressions.
4. Treat research-paper posts as hypothesis priors for future scout slots, especially analyst/sentiment/earnings-announcement fields.

## Recommended next read set

Metadata-selected post ids:

- `41197045436951` — concentrated weight repair
- `41148978492183` — turnover repair
- `41141847644311` — trade_when
- `41279111112855` — prod_corr failure and virgin dataset pivot
- `41223887764247` — IS robustness to prod_corr staging
- `41313321869463` — anti-invalid-backtest small-batch methodology

Full reads must remain read-only, sanitized, and claim-level. Do not store full private post bodies in Git unless explicitly approved and audience-safe; prefer extracting short claims with post id/title/source metadata.
