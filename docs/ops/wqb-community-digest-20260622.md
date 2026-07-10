# WQB community digest — 20260622

Source: `runs/wqb-community/forum-post-index-latest.json`

Mode: metadata-only, read-only. No full post bodies, [REDACTED]s, [REDACTED]s, comments, votes, or forum writes are stored.

## Summary

- Posts indexed: 50
- Active clusters: 8
- Read queue size: 12

## Cluster counts

- **platform_webinars_competitions**: 4
- **research_paper_hypotheses**: 3
- **ai_research_workflow_and_infra**: 13
- **python_alpha_and_signal_processing**: 10
- **correlation_prod_corr_and_robustness**: 7
- **gate_failure_repair_methods**: 13
- **dataset_guides_and_alpha_cases**: 9
- **uncategorized**: 7

## Top read queue

- 41285657542295 — 分享几个 Super Alpha 的 Combo 表达式：给“三五”低相关 SA 多一点试法 (PROD_CORRELATION, score=0.8492)
- 41314164291863 — PAC 2026 – Simulation Optimization Webinar recording (11 Jun) (SIMULATION_BUDGET_WASTE, score=0.7877)
- 15465058281495 — Research Paper 13: Valuation Ratios, Surprises, Uncertainty or Sentiment: How Does Financial Machine Learning Predict Returns From Earnings Announcements? (LOW_FITNESS, score=0.7284)
- 41148978492183 — 降低 Alpha 周转率的核心策略 (HIGH_TURNOVER, score=0.7255)
- 41141847644311 — 换个角度，理解一下trade_when(x, y, z) 方法 (HIGH_TURNOVER, score=0.7157)
- 41162636720919 — 提高MEA区的Sub-universe Sharpe of 1.11 is below cutoff of 1.14.问题一点经验 (LOW_SUB_UNIVERSE_SHARPE, LOW_FITNESS, score=0.7156)
- 41143142740631 — 【IND/D1】批量 prod_corr 断点续跑与结果落盘 — partial CSV + pickle 备份实战 (PROD_CORRELATION, score=0.7152)
- 41197045436951 — CONCENTRATED WEIGHT 系统性排查与修复手册-2（基于实战经验的 CW 攻防全指南_进阶篇） (CONCENTRATED_WEIGHT, score=0.7054)

## Scheduler implications

- Do not dispatch Batch F before selected claim-level full reads for CW, turnover/trade_when, prod_corr staging, and anti-invalid-backtest methodology.
- Treat Python/Fourier/automation posts as tooling or diagnostic priors, not copyable expressions.
- Treat research-paper posts as scout priors for analyst/sentiment/earnings-announcement directions.

