# WQB premium official source claims — 2026-06-22

## Context

User instruction: first ingest truly important premium official/community documents and design the workflow well before any further simulation. User rule: forum posts with `vote_count >= 50` must be read.

What I could do directly in the current Notion/CNHK MCP environment:

- Read WQB official tutorial index via CNHK Yellow.
- Read multiple official tutorial pages directly via CNHK Yellow.
- Try public web loads for several community URLs.

What failed here and remains local/CNHKMCP-live work:

- Public web loads for protected WQB community posts timed out with `CRAWL_LIVECRAWL_TIMEOUT: 504`.
- Current Notion-side Yellow MCP does not expose live forum `read_forum_post` / `read_full_forum_post` tools.

Therefore this file extracts all premium official claims that can be produced here without asking the user to run local commands.

## Official pages read directly

- `intermediate-pack-part-2` — Intermediate Pack: Improve your Alpha
- `getting-started-high-turnover-alphas` — Getting Started with High Turnover Alphas
- `data` — Understanding Data in BRAIN: Key Concepts and Tips
- `getting-started-model77` — Model77 dataset
- `getting-started-sentiment1-dataset` — Sentiment1 dataset
- `alpha-submission` — Clear these tests before submitting an Alpha

## Claim extraction

### WQB-OFFICIAL-018 — Rank and truncation address concentration differently

Source: `intermediate-pack-part-2`

Claim: Official docs illustrate that cross-sectional `rank` can reduce extreme position concentration by transforming raw values into distributed ranks, while truncation separately caps maximum single-stock weight.

Failure modes:

- `CONCENTRATED_WEIGHT`
- `LOW_FITNESS`

Design effect:

- Candidate grammar should distinguish **signal transform** concentration control (`rank`, group/rank transforms) from **portfolio cap** control (`truncation`).
- For fnd65/CW repair, do not rely only on truncation; test whether field/operator transformation reduces concentration earlier in the pipeline.

### WQB-OFFICIAL-019 — Decay reduces turnover but can attenuate signal

Source: `intermediate-pack-part-2`

Claim: Decay averages alpha signal over a specified time window; it can reduce turnover, but values that are too large attenuate the signal.

Failure modes:

- `HIGH_TURNOVER`
- `LOW_FITNESS`
- `LOW_SHARPE`

Design effect:

- Decay is a bounded repair axis, not a monotonic improvement knob.
- Batch design should record whether a decay change is testing turnover reduction, signal attenuation, or both.
- Repeated decay tweaks without mechanism change should be penalized by the star-map.

### WQB-OFFICIAL-020 — Neutralization choice must follow alpha logic and results

Source: `intermediate-pack-part-2`

Claim: Correct neutralization depends on the logic/formula used by the alpha; results should indicate which neutralization is most effective.

Failure modes:

- `PROD_CORRELATION`
- `LOW_SUB_UNIVERSE_SHARPE`
- `CONCENTRATED_WEIGHT`

Design effect:

- Neutralization should be a diagnostic branch with rationale, not a default post-hoc knob.
- CROWDING/RAM/statistical neutralizations should be used to test a hypothesis about exposure/crowding, not as blind repairs.

### WQB-OFFICIAL-021 — High turnover should be a consequence of fast-moving information, not the target

Source: `getting-started-high-turnover-alphas`

Claim: Official docs explicitly say the most common mistake is targeting turnover directly rather than targeting a short-lived source of information; high turnover should be a consequence of the idea, not the idea itself.

Failure modes:

- `HIGH_TURNOVER`
- `SIMULATION_BUDGET_WASTE`
- `OVERFITTING`

Design effect:

- Do not create high-turnover or turnover-repair candidates by adding noise or unstable transforms.
- Candidate proposals must state the fast-moving information source: event reaction, flow/activity, microstructure-sensitive effect, short-horizon fundamental refresh, or interaction with a slower conditioning variable.

### WQB-OFFICIAL-022 — High-turnover workflow requires clean base, mechanics validation, realism stress, then expansion

Source: `getting-started-high-turnover-alphas`

Claim: Official workflow for High TVR research: start with a simple clean base alpha, validate signal mechanics and coverage, observe turnover as an outcome, test robustness under realistic variants, and expand only after the base idea survives.

Failure modes:

- `SIMULATION_BUDGET_WASTE`
- `OVERFITTING`
- `SOURCE_DRIFT`

Design effect:

- Batch F/G proposals should not start from complex multi-transform blends.
- Each scout should first prove a clean mechanism, then add variants only after survival.
- Expected-information-gain must include what the simple base test falsifies.

### WQB-OFFICIAL-023 — High-turnover mistakes map directly to stop rules

Source: `getting-started-high-turnover-alphas`

Claim: Official docs warn against artificial turnover, single-test dependence, weak economic story, coverage fragility, and parameter mining.

Failure modes:

- `OVERFITTING`
- `SIMULATION_BUDGET_WASTE`
- `LOW_SUB_UNIVERSE_SHARPE`
- `LOW_2Y_SHARPE`

Design effect:

- Add stop-rule penalties for candidates whose performance depends on one universe, one cost assumption, one constrained setting, or small parameter flips.
- Add coverage fragility and weak-story checks before allowing population climbing.

### WQB-OFFICIAL-024 — New data fields require coverage/frequency/bounds/distribution probes

Source: `data`

Claim: Official docs recommend six basic expressions for understanding a new data field: raw field coverage, non-zero coverage, unique update frequency via `ts_std_dev`, bounds via `abs(datafield) > X`, long-window median/mean, and distribution via `scale_down` ranges.

Failure modes:

- `SIMULATION_BUDGET_WASTE`
- `CONCENTRATED_WEIGHT`
- `LOW_SUB_UNIVERSE_SHARPE`
- `OVERFITTING`

Design effect:

- Scout slots must be preceded by data-field probes or metadata claims about coverage/frequency/bounds/distribution.
- A field should not enter expensive candidate simulation solely because its name looks plausible.
- Add field-profiling step to terrain atlas before broad scout expansion.

### WQB-OFFICIAL-025 — Low coverage demands liquid-subset testing or cautious backfill

Source: `data`, `getting-started-model77`, `getting-started-sentiment1-dataset`

Claim: Official docs emphasize coverage analysis. For Model77 fields with coverage below 70%, focus on highly liquid subsets such as TOP1000/TOPSP500 or use reasonable backfill cautiously. Sentiment1 coverage is about 2000 in TOP3000; test more liquid universes while ensuring sufficient long/short counts.

Failure modes:

- `CONCENTRATED_WEIGHT`
- `LOW_SUB_UNIVERSE_SHARPE`
- `LOW_2Y_SHARPE`
- `OVERFITTING`

Design effect:

- Dataset/field scout should include coverage-aware universe diagnostics.
- Backfill should be a cautious field-quality repair, not an automatic default.
- A candidate with weak coverage should not be climbed before coverage and sub-universe behavior are understood.

### WQB-OFFICIAL-026 — Sentiment and analyst fields have different natural horizons

Source: `getting-started-sentiment1-dataset`, `getting-started-model77`

Claim: Sentiment1 combines dynamic sentiment metrics with earnings/analyst signals; sentiment scores are higher-turnover in nature, while analyst and earnings metrics are slower-moving. Long lookbacks over 63 days may lose relevance for high-frequency sentiment events. Model77 contains valuation, earnings, quality, momentum, and risk metrics with different horizons.

Failure modes:

- `LOW_SHARPE`
- `LOW_FITNESS`
- `HIGH_TURNOVER`
- `OVERFITTING`

Design effect:

- Horizon buckets must be dataset/field-type aware.
- Sentiment/attention fields should not automatically reuse fnd65 or analyst-revision horizons.
- Model77/analyst scout should split valuation, earnings surprise, quality, momentum, and risk subfamilies rather than treating dataset as one monolith.

### WQB-OFFICIAL-027 — Submission thresholds define final-click hard state

Source: `alpha-submission`

Claim: Delay-1 alphas require Fitness > 1, Sharpe > 1.25, Turnover between 1% and 70%, max weight < 10%, sub-universe robustness, and self-correlation < 0.7 or sufficient Sharpe improvement over correlated alphas.

Failure modes:

- `LOW_SHARPE`
- `LOW_FITNESS`
- `HIGH_TURNOVER`
- `LOW_TURNOVER`
- `CONCENTRATED_WEIGHT`
- `LOW_SUB_UNIVERSE_SHARPE`
- `SELF_CORRELATION`

Design effect:

- `final_click_candidate` remains a hard state; no candidate can be called near-final just because it has a good local metric.
- Gate scorer must preserve official thresholds and not replace them with heuristic star-map scores.

### WQB-OFFICIAL-028 — Weight test is both too-few-stocks and too-concentrated-stock exposure

Source: `alpha-submission`

Claim: Weight test can fail when too few stocks are assigned weight for a significant number of days or when alpha weight is too concentrated in one stock.

Failure modes:

- `CONCENTRATED_WEIGHT`
- `LOW_SUB_UNIVERSE_SHARPE`

Design effect:

- CW repair must separate breadth/coverage failures from max-single-stock concentration failures.
- Future claim extraction should distinguish `too_few_instruments` vs `max_weight_too_high` instead of grouping both into one undifferentiated blocker.

### WQB-OFFICIAL-029 — Sub-universe test is a robustness/liquidity test, not an arbitrary cutoff

Source: `alpha-submission`

Claim: Sub-universe test checks whether an alpha works in the next more liquid/smaller universe; failure suggests profit may come from illiquid portions and may not be robust out-of-sample.

Failure modes:

- `LOW_SUB_UNIVERSE_SHARPE`
- `LOW_2Y_SHARPE`
- `OVERFITTING`

Design effect:

- Sub-universe failure should trigger liquidity-aware diagnostics and possibly discard, not only Sharpe boosting.
- Add explicit universe/liquidity hypothesis to repair candidates.

### WQB-OFFICIAL-030 — Sub-universe repair includes avoiding size multipliers and liquidity-aware decay

Source: `alpha-submission`

Claim: Official tips include avoiding company-size multipliers that shift weights toward/away from liquidity buckets; one can decay liquid and non-liquid parts separately using liquidity proxies such as cap or volume*close; and improvements should be checked step-by-step.

Failure modes:

- `LOW_SUB_UNIVERSE_SHARPE`
- `LOW_2Y_SHARPE`
- `CONCENTRATED_WEIGHT`

Design effect:

- Add liquidity-aware diagnostic candidates only when sub-universe is a real blocker.
- Record whether a proposed repair shifts exposure across liquidity buckets.
- Do not mix multiple repairs in one candidate if it prevents attribution.

### WQB-OFFICIAL-031 — Low correlation is often more valuable than minor performance gain

Source: `alpha-submission`

Claim: Official docs advise not spending extraordinary time improving a single idea; it is generally better to try new ideas with low correlation, and low correlation is more important than a minor performance increase.

Failure modes:

- `PROD_CORRELATION`
- `SELF_CORRELATION`
- `LINEAR_SEARCH_BIAS`
- `SIMULATION_BUDGET_WASTE`

Design effect:

- Batch F should not over-allocate to local repair when prod_corr remains a blocker.
- Star-map should reward mechanism/data-family pivots after repeated correlation failures.
- Population climbing must have a strict budget and stop if correlation is not improving.

## Immediate design consequence

These official claims strengthen the user's premium-source-first correction:

- Batch F should remain blocked until premium community claims are extracted or explicitly blocked.
- Even without community body access, official docs already require stronger data-field profiling, coverage diagnostics, turnover-mechanism discipline, and low-correlation pivots.
- The current Batch F authorization packet is now stale relative to the premium-source intake policy and should be rewritten after claims are integrated.
