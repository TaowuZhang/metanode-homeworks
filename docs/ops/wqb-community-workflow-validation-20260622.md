# WQB community workflow validation — 2026-06-22

## Result

The local forum refresh workflow is validated end-to-end.

Validated path:

```text
local CNHKMCP authenticated read
→ metadata-only forum index
→ metadata distillation
→ local sensitive-word scan
→ GitHub commit/push
→ Notion AI reads GitHub artifacts directly
```

## User local run

Local commit pushed by user:

```text
175f973 chore(wqb): refresh forum community artifacts
```

Observed refresh output:

- `call_ok=true`
- `post_count=50`
- raw top-level keys: `success`, `topic_id`, `posts`, `total_found`
- wrote:
  - `runs/wqb-community/forum-post-index-latest.json`
  - `runs/wqb-community/forum-post-index-20260622T152946Z.json`
  - `runs/wqb-community/forum-post-index-local.json`

Observed distillation output:

- `post_count=50`
- `read_queue_size=12`
- clusters:
  - `platform_webinars_competitions`: 4
  - `research_paper_hypotheses`: 3
  - `ai_research_workflow_and_infra`: 13
  - `python_alpha_and_signal_processing`: 10
  - `correlation_prod_corr_and_robustness`: 7
  - `gate_failure_repair_methods`: 13
  - `dataset_guides_and_alpha_cases`: 9
  - `uncategorized`: 7

The user then committed and pushed 9 artifact files to GitHub. Notion AI successfully read the pushed artifacts through GitHub MCP.

## Top read queue after validation

Current read queue is metadata-selected and must be used for claim-level summaries only, with no full post body storage.

Top items:

1. `41285657542295` — Super Alpha Combo / low-correlation SA — `PROD_CORRELATION`
2. `41314164291863` — PAC 2026 Simulation Optimization Webinar — `SIMULATION_BUDGET_WASTE`
3. `15465058281495` — Valuation ratios / surprises / sentiment / earnings announcements — `LOW_FITNESS`
4. `41148978492183` — Lower Alpha turnover — `HIGH_TURNOVER`
5. `41141847644311` — `trade_when(x, y, z)` interpretation — `HIGH_TURNOVER`
6. `41162636720919` — Sub-universe Sharpe repair experience — `LOW_SUB_UNIVERSE_SHARPE`, `LOW_FITNESS`
7. `41143142740631` — prod_corr breakpoint continuation and result persistence — `PROD_CORRELATION`
8. `41197045436951` — Concentrated Weight repair manual — `CONCENTRATED_WEIGHT`
9. `41163824201623` — Codex flow failed after hundreds of thousands of backtests — `SIMULATION_BUDGET_WASTE`
10. `41279111112855` — prod_corr failed, virgin dataset passed — `PROD_CORRELATION`
11. `41292678360727` — frequency-aware seed, CW pre-diagnosis, multi-GEM decorrelation — `PROD_CORRELATION`, `CONCENTRATED_WEIGHT`
12. `41313321869463` — anti-invalid-backtest small-batch method — `SIMULATION_BUDGET_WASTE`

## Scheduler consequences

This validates the workflow but does not yet authorize simulation.

Immediate consequences:

- Keep Batch F undispatched until explicit per-batch authorization.
- Before Batch F dispatch, use this read queue to extract claim-level rules for:
  - prod_corr / low-correlation / virgin dataset pivots,
  - simulation-budget discipline,
  - turnover / `trade_when`,
  - sub-universe Sharpe,
  - concentrated weight,
  - anti-invalid-backtest methodology.
- Treat community evidence as a scheduler prior and failure-mode map, not copyable alpha expressions.
- Keep Cloudflare as optional read-only artifact facade; do not move WQB credentials to Cloudflare in the first version.

## Current state

The community ingestion layer has moved from one-off manual extraction to a reusable local-refresh / GitHub-artifact workflow.

The next productive work is not another simulation dispatch; it is to convert the top read queue into claim-level summaries and update the star-map / Batch F weighting accordingly.
