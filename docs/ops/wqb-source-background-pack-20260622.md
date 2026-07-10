# WQB Source Background Pack — 2026-06-22

## Purpose

This pack collects implementation-relevant background before redesigning the WQB Alpha production system. The goal is not to jump into another simulation batch, but to understand platform mechanics, CNHKMCP capabilities, WQB documentation, and existing GitHub ecosystem patterns.

Current conclusion: the architecture should be source-first and prior-aware. Simulation is only one expensive measurement step, not the whole system.

## 1. CNHKMCP / external ecosystem

### cnhkmcp package

Observed from PyPI:

- package: `cnhkmcp`
- observed version installed locally by user: `3.1.15`
- entry point from local audit: `cnhkmcp.untracked.platform_functions:main`
- PyPI description: CNHK MCP Server for quantitative trading platform integration.
- PyPI feature list includes:
  - API integration;
  - simulation management;
  - data access;
  - alpha management;
  - forum integration;
  - performance analysis;
  - competition / leaderboard support.
- PyPI README names `forum_functions.py` for glossary extraction, forum post search/reading, and support documentation access.
- Requirements include Chrome browser for forum functionality.

Implication:

- The community/forum route should use CNHKMCP capabilities locally, not public web crawling.
- Cloudflare v1 should not hold WQB credentials; it should read sanitized artifacts from GitHub.
- If a future remote CNHKMCP Worker is needed, it needs a separate security review because forum access can require browser/session state.

### worldquant-skill

Observed from GitHub and user local audit:

- repo: `GRD-Chang/worldquant-skill`
- integrates WorldQuant BRAIN + `cnhkmcp` + Claude Code / Gemini CLI / iFlow.
- core skills:
  - `alpha-research-recorder`
  - `factor_backtest`
  - `knowledge_base_search`
- knowledge base structure includes:
  - `Resources/DATASET`
  - `good_alpha_examples`
  - `How_WorldQuant_BRAIN_Backtesting_Works.md`
  - `alpha_optimization`
  - `regular_operators.csv`
- user local scan produced shallow prior counts:
  - `correlation_reduction=2`
  - `turnover_reduction=3`
  - `overfitting_avoidance=0`
  - `weight_coverage=1`
  - `dataset_evaluation=2`

Implication:

- The integration route is proven, but the current extracted priors are shallow SKILL/README hits.
- Next source work should run actual `knowledge_base_search` if available and/or inspect `Resources/alpha_optimization` locally.
- Overfitting priors are missing and must be supplemented before trusting any high-parameter repair batch.

## 2. WQB documentation priors already read

### Vector data fields

Source page: `vector-datafields`.

Key facts:

- Vector fields do not have fixed size; number of events per instrument/day varies.
- Final Alpha output must be matrix values, so vector data must be converted to matrix before normal matrix operators.
- Vector aggregation operators include:
  - `vec_avg`
  - `vec_choose`
  - `vec_count`
  - `vec_ir`
  - `vec_kurtosis`
  - `vec_max`
  - `vec_min`
  - `vec_norm`
  - `vec_percentage`
  - `vec_powersum`
  - `vec_range`
  - `vec_skewness`
  - `vec_stddev`
  - `vec_sum`
- Vector-derived features can have high raw turnover. WQB docs explicitly mention using decay operators, `ts_rank`, or `ts_decay` to reduce turnover.

System implications:

- Behavioral VECTOR fields should not default only to `vec_avg` forever.
- The star map should explicitly include vector aggregation as a search dimension.
- VECTOR candidates need a turnover prior before simulation.

### Basic neutralization

Source page: `neut-cons`.

Key facts:

- Neutralization subtracts group means and focuses on relative returns inside market/sector/industry/subindustry groups.
- `group_neutralize(x, group)` and neutralization in simulation settings are operationally related; if manually neutralizing in expression, settings can be `None` with decay/truncation controlled in expression.
- Always pick a neutralization value unless it is manually incorporated.
- Dataset category recommendations:
  - Fundamental: Industry
  - Analysts: Industry
  - Model: experiment across market/sector/industry/subindustry
  - News: Subindustry
  - Options: Market/Sector
  - Price Volume: Market/Sector, because finer neutralization may reduce generic idea performance
  - Social Media: Industry/Subindustry
  - Institutions: Sector/Industry
  - Short Interest: Industry
  - Insider: Industry/Subindustry
  - Sentiment: Industry/Subindustry
  - Earnings: Industry
  - Macro: Market/Sector/Industry

System implications:

- Batch F slot neutralization should not be arbitrary.
- fnd65/fundamental/analyst slots should prefer Industry or carefully justify alternatives.
- price-volume / ai_factor slots should not reflexively use Industry/Subindustry.

### Risk-neutralized alphas

Source page: `getting-started-risk-neutralized-alphas`.

Key facts:

- Risk-neutralized settings include Slow Factors, Fast Factors, Slow + Fast Factors, RAM, Statistical, and Crowding.
- These settings seek returns orthogonal to market, industry, and style factors.
- Slow Factors are recommended as a starting point for low-turnover signals.
- Fast or Slow + Fast may fit higher-turnover signals, but can increase turnover.
- Neutralization can improve Sharpe/drawdown but also increase turnover; choice must balance benefits and costs.
- WQB recommends trying risk neutralization on previously submitted alphas and observing impact.
- Forum pages are referenced for how to start risk neutralized research and how to choose risk factor sets.

System implications:

- Risk neutralization should be a diagnostic branch, not a blind repair operator.
- Batch E showed CROWDING did not solve the D02 branch and worsened prod correlation in some variants; this should reduce blind CROWDING priority but not eliminate targeted risk-neutralized diagnostics.
- The scheduler should track `neutralization_effect` by family, not just final alpha score.

### Crowding risk-neutralized alphas

Source page: `getting-started-crowding-risk-neutralized-alphas`.

Key facts:

- Crowding risk occurs when many investors hold similar positions/trades.
- Crowded trades can unwind sharply and reduce profitability.
- CROWDING neutralization is intended to control crowding style risk.

System implications:

- CROWDING is relevant to `PROD_CORRELATION`, but not guaranteed to reduce it.
- It should be tested as a diagnostic against a control, not as a default fix.

### Simulation limits and effective simulation use

Source page: `understanding-simulation-limits`.

Key facts:

- All successful simulations count, including child simulations in multi-simulations and re-simulations of existing alphas.
- Duplicate alphas are not removed from the count.
- API simulation response headers can provide daily limit, remaining count, and reset time.
- WQB recommends planning simulations, monitoring count, and pausing/stopping near limits.
- Effective search recommendations:
  - analyze data fields: coverage, date coverage, description, alpha count, user count;
  - use BRAIN Labs for field coverage, turnover, typical values, and return correlation;
  - review risk-handled and investability-constrained performance;
  - initially restrict search space and test one element at a time;
  - pick one data field from similar groups for initial signal;
  - operator groups: aggregational, delta-based, group-based;
  - for fast signals, start with 5/10/21 day windows;
  - for slow signals, start with 63/121/252 day windows;
  - start with 10–50 simulations to test a search space, but in this system platform policy constrains per-batch dispatch to explicit safe batches.

System implications:

- Duplicate hash precheck is mandatory.
- Our safe batch size of 6 is conservative; more important is information yield per simulation.
- Every candidate must declare exactly which single element it tests unless explicitly marked as crossover.
- Field analysis should precede scout slots.

### Single dataset alphas

Source page: `single-dataset-alphas`.

Key facts:

- Single Dataset Alphas use fields from only one dataset, except grouping fields such as country/exchange/market/sector/industry/subindustry.
- Mixing multiple datasets can lead to overfitting if implemented poorly.
- Single dataset alphas are less prone to overfitting and can have more robust predictions.
- Single dataset alphas have relaxed IS testing: they do not need IS Ladder Sharpe, but must satisfy Last 2Y IS Sharpe thresholds.
- Delay-1 Last 2Y Sharpe threshold: 2.38.
- Delay-0 Last 2Y Sharpe threshold: 3.96.
- If turnover is below 30%, IS Sharpe Ladder thresholds are multiplied by 0.85.

System implications:

- Batch F should preserve single-dataset discipline where possible.
- Cross-dataset blending should be treated as overfit risk and require stronger rationale.
- `LOW_2Y_SHARPE` remains a serious blocker; single-dataset relaxation is not free.

### Simulation settings

Source page: `simulation-settings`.

Key facts:

- Delay is applied automatically in expression language.
- Decay can reduce turnover, but too-large decay attenuates signal.
- Truncation legal range is 0 to 1. Recommended setting is 0.05 to 0.1 in docs, though current system has used much smaller truncation for WQB gates.
- Pasteurization ON converts inputs outside universe to NaN; OFF can use all available inputs and may change cross-sectional/group behavior.
- NaNHandling ON can increase coverage but may introduce ambiguous information; OFF preserves NaNs and requires manual handling.
- Unit Handling can warn on incompatible units.

System implications:

- Small truncation repairs should be treated as gate-specific, not assumed globally optimal.
- NaNHandling should become a candidate dimension for coverage/weight repair, but only with careful interpretation.
- Unit warnings/type compatibility should be part of proposal validation.

## 3. Current information gaps

1. CNHKMCP local forum functions still need direct tool audit beyond PyPI description.
2. `worldquant-skill` Resources content needs deeper local extraction, especially `alpha_optimization` and overfitting-related files.
3. Need to collect forum posts referenced by risk-neutralized docs:
   - how to start risk neutralized research;
   - how to choose risk factor sets.
4. Need a duplicate simulation hash implementation before any further batch dispatch.
5. Need BRAIN API / exact simulation headers docs; an attempted `brain-api` documentation page read hit WQB 429, so this is deferred until cooldown.
6. Need field-level briefs for Batch F candidate fields using Yellow `field_brief` before dispatch.

## 4. Design consequences

The next system design should be a source-aware research machine:

```text
Source acquisition
  -> source classification
  -> claim extraction
  -> failure-mode mapping
  -> heatmap/star-map update
  -> candidate slot rationale
  -> proposal validator
  -> explicit authorized simulation
  -> recovery and scoring
  -> source feedback update
```

Key design rules:

- Do not design from intuition when relevant WQB source material is still missing.
- Do not use public crawler for authenticated community content.
- Do not treat `worldquant-skill` examples as expressions to copy.
- Prefer single-dataset experiments unless the purpose is explicitly cross-dataset diagnostic.
- Treat neutralization as a tracked experimental axis with family-specific history.
- Track simulation budget and duplicate hashes before every batch.
- Every failed simulation must update source-backed beliefs, not just a score.
