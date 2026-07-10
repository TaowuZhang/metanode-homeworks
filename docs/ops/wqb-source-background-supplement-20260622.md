# WQB Source Background Supplement — 2026-06-22 22:40

## Why this supplement exists

The current work direction is source collection before redesign. This supplement adds background collected after the first source background pack, especially public notes, worldquant-skill skill docs, and negative findings around the public `cnhk/cnhkmcp` repository.

No simulation was dispatched.

## 1. Public notes: jglazar WorldQuant seminar

Source: `jglazar/notes/quant_interview/worldquant_seminar.md`, loaded through GitHub MCP.

### Operator / mechanics priors

- Cross-sectional operators: `rank`, `zscore`, `scale`, `sign`.
- Time-series operators: `ts_delta`, `ts_delay`, `ts_sum`, `ts_mean`, `ts_rank`, `ts_zscore`, `ts_std_dev`, `ts_regression`.
- Group operators are stronger cross-sectional operators and require a group axis such as sector / industry / subindustry.
- `trade_when` reduces turnover by activating between entry and exit conditions.
- `humpdecay` / hump-like decay can reduce turnover.
- `rank` can help weighting/diversity by converting raw magnitudes into percentile-like values.
- `vector_neut` and `group_vector_neut` are documented in the notes as tools to remove exposure to another vector / risk factor.

### Search priors

- Reduce turnover by increasing alpha decay, combining operators, `keep`, `trade_when`, and hump-like mechanisms.
- Combining operators can also reduce correlation.
- Same idea can be explored by different data fields and different sibling operators.
- Toggling settings, especially decay and neutralization, is a legitimate improvement axis.
- News and sentiment data require vector aggregation (`vec_avg` or similar) and usually need turnover control.
- Use Data Explorer / data descriptions for field farming.
- Reduce self-correlation by working with diverse datasets.
- Avoid overfitting; use few parameters; diversify alpha pool.
- Researcher advice emphasized:
  - review alphas and think about improvements;
  - attention to detail;
  - good performance across many settings;
  - quality through simplicity;
  - original alpha for presentation.

### Design implications

- The system should track `operator_sibling`, `field_sibling`, `settings_sibling`, and `risk_neutralization_sibling` as different mutation axes rather than treating all edits as the same.
- A good climber changes one axis at a time after a signal is found.
- Overfitting control should include:
  - few parameters;
  - simple expression structure;
  - single-dataset discipline where possible;
  - performance across settings rather than single-setting luck.
- `trade_when` and vector/risk neutralization should enter the operator grammar as purposeful tools, not random embellishment.

## 2. Public competition write-up: jglazar WQ project page

Source: `https://jglazar.github.io/projects/wq_project/`, loaded by web.

### Scoring prior

The public write-up states that fitness is:

```text
sqrt(abs(Returns) / max(turnover, 0.125)) * Sharpe
```

This is consistent with WQB-style scoring incentives: high Sharpe, high returns, and lower turnover.

### Portfolio prior

Scores from submitted alphas reward baskets of strong and uncorrelated returns.

### Example priors

- Simple price-volume reversion can have good Sharpe but high turnover.
- Fundamental/earnings-style data can reduce turnover because it changes more slowly.
- News/social-media style alphas may provide complementary / uncorrelated returns.
- Clamping / truncating denominators can avoid overweighting a few stocks.
- Industry/subindustry neutralization appears in examples where the signal is naturally group-relative.

### Design implications

- Fitness needs explicit decomposition in the gate scorer:
  - Sharpe component;
  - returns component;
  - turnover penalty via `max(turnover, 0.125)`.
- Search should not only maximize single-alpha score; it should manage pool originality and correlation.
- Fundamental / slower-update families deserve separate treatment from price-volume / news / sentiment families because their turnover dynamics differ.

## 3. worldquant-skill direct docs

Source: `GRD-Chang/worldquant-skill` README and `skills/knowledge_base_search/SKILL.md`, loaded through GitHub MCP.

### Confirmed skill architecture

The project explicitly describes three skills:

- `alpha-research-recorder`
- `factor_backtest`
- `knowledge_base_search`

`knowledge_base_search` covers:

- data field lookup;
- optimization methods;
- high-quality alpha examples;
- platform mechanisms.

It requires:

- traceable source paths;
- no fabricated content;
- relevance ranking;
- clear, actionable steps;
- direct search over `Resources/`.

The documented knowledge base structure includes:

```text
Resources/
  IND_DATASET/
  good_alpha_examples/
  How_WorldQuant_BRAIN_Backtesting_Works.md
  alpha_optimization/
  regular_operators.csv
```

### Design implications

- Our source ingestion layer should not simply grep top-level README. It should use the same structure:
  - `Resources/alpha_optimization` for gate repairs;
  - `Resources/good_alpha_examples` for design patterns, not expression copying;
  - `Resources/IND_DATASET` for field lookup;
  - platform mechanics docs for validators.
- Every extracted prior must carry source path and query.
- The output should be claim-level, not page-level.

## 4. cnhkmcp public repository negative findings

Attempts:

- PyPI links advertise `https://github.com/cnhk/cnhkmcp`.
- GitHub MCP `get_file_contents` for `cnhk/cnhkmcp` returned 404.
- GitHub code search for `repo:cnhk/cnhkmcp search_forum_posts`, `forum`, and `BrainApiClient` returned zero results.
- Web raw GitHub loads for likely paths also returned 404 or timeout.

Interpretation:

- PyPI package is real and locally installed, but the public GitHub repo advertised by PyPI is not currently accessible through available tools.
- We should treat PyPI README claims as package metadata, and rely on local installed package introspection for actual callable surface.
- The forum route remains local-first; do not build Cloudflare around credentials.

## 5. What I can still do here without user local action

I can continue to:

1. Read WQB documentation/tutorial pages through CNHK Yellow, respecting 429 cooldown.
2. Use Yellow `dataset_brief` / `field_brief` to collect field-level context for candidate families.
3. Use GitHub MCP to read public repos when accessible.
4. Search public web for non-authenticated notes and official pages, with source quality grading.
5. Build claim registries and source packs in GitHub.
6. Update the heatmap / starmap weight model from source-backed claims.
7. Update the proposal validator requirements.
8. Prepare but not dispatch Batch F unless exact per-batch authorization is present.

I should not ask the user to run local commands unless:

- external package / repo access fails in the Notion sandbox and GitHub MCP cannot fetch it;
- WQB community/forum pages require local authenticated browser/session route;
- Cloudflare login/deploy/secrets are required;
- local file contents are needed and cannot be accessed through pushed sanitized artifacts.

## 6. Next source gaps

1. Local extraction from `external/worldquant-skill/Resources/alpha_optimization`, especially overfitting, sub-universe sharpe, correlation, and weight coverage.
2. Field briefs for current Batch F fields:
   - `consecutive_return_streak_length`
   - `salience_weighted_return_score`
   - `extreme_daily_return_indicator`
   - `fnd65_us5000_cusip_d41isr`
   - `money_flow_strength_typical_price`
   - `actual_revenue_last_quarter_medium`
3. Retry `brain-api` docs after WQB cooldown.
4. Collect risk-neutralized forum pages through local CNHKMCP if available.
5. Build duplicate-hash precheck before any new simulation.