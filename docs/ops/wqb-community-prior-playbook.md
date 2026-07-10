# WQB Community Prior Playbook

## Purpose

This file turns WQB documentation and community-style guidance into scheduler priors. It does not copy private alphas or public examples into the candidate library. It extracts process rules.

## Source pages read through CNHK Yellow

- `Must-read posts: How to improve your Alphas`
- `Finding Consultant Alphas`
- `Consultant Dos and Don'ts`
- `Consultant Submission Tests`
- `How can you avoid duplicate simulations?`
- `Pyramid multipliers`
- `Consultant leaderboard`

## Extracted priors

### 1. Consultant production correlation is central

Source guidance: consultant production correlation compares against all existing consultant alphas. A good alpha can fail if it is not unique enough. The page advises not running correlation too often because requests may be limited.

Scheduler rule:

- Treat `PROD_CORRELATION` as a strategic blocker, not a late cosmetic check.
- Do one early prod-correlation read for a promising new idea family, then avoid repeated correlation calls until material improvements exist.
- If prod correlation fails under local parameter tweaks, pivot mechanism / dataset / region rather than continuing parameter repair.

### 2. IS Ladder / recent years matter more than global backtest comfort

Source guidance: IS Ladder checks recent 2,3,...10 year Sharpe and can fail immediately if recent 2Y Sharpe is below threshold. Low-turnover alphas get a multiplier on pass thresholds, but fail threshold remains.

Scheduler rule:

- `LOW_2Y_SHARPE` is not a small issue; it is a primary kill/repair signal.
- Deep repair must include yearly stats / PnL diagnostics when available.
- A high full-period Sharpe with bad recent years is not near-gate.

### 3. Start from familiar ideas, but do not stay there

Source guidance: consultants should first understand new tests using familiar ideas, then move to other datasets/categories and new regions.

Scheduler rule:

- Legacy fnd65 repair is a learning/control branch, not the whole strategy.
- Maintain scout slots into new categories even while repairing a known parent.

### 4. New datasets require evaluation, not brute expression generation

Source guidance: evaluate new datasets quickly, handle vector fields, check coverage issues, prefer themes/multipliers when useful.

Scheduler rule:

- Before allocating scout slots to a dataset, use `dataset_brief` / `field_brief` where possible.
- Vector fields need explicit matrix adapters.
- Coverage/field type must be stored in star nodes.

### 5. Research papers and community examples are idea priors, not expressions to copy

Source guidance: papers rarely work as-is; sometimes opposite works. They guide direction. Responsible combination of ideas can create multiple alphas.

Scheduler rule:

- A source idea becomes a hypothesis node and operator archetype, not a stored expression.
- Add inverse / diagnostic candidates when a paper/community direction fails or has negative evidence.
- Do not copy public example expressions into long-term candidates.

### 6. Dos and Don'ts penalize overfitting and noisy correlation hacks

Source guidance: use economic foundation; simplify; use different datasets/operators; test sub/super-universe and regions; restrict parameter search to reasonable values like 5,20,60,120,252; improve ideas rather than adding/fitting parameters. Do not add noise to reduce correlation; do not overfit with too many parameters/datafields/if_else filters.

Scheduler rule:

- Parameter grids are bounded to canonical windows unless there is a reason.
- Noise-based decorrelation is forbidden.
- Too many if_else/filter combinations increase risk score.
- Every candidate must have a one-minute economic explanation.

### 7. Duplicate simulations must be blocked

Source guidance: hash full alpha configuration and reuse existing alpha_id instead of rerunning duplicates.

Scheduler rule:

- Add a dispatch precheck over a full configuration hash.
- Hash must include expression/config/settings/region/universe/delay/truncation/etc.
- If duplicate, reuse previous alpha_id and do not consume simulation quota.

### 8. Leaderboard breadth prior

Observed leaderboard data: top users often have many data fields used and many submissions; examples include users with 89, 137, 1799, or 6350 dataFieldsUsed. Mean prod/self correlations vary, with successful users often still near platform thresholds.

Scheduler rule:

- Breadth matters. A narrow single-field loop is unlikely to be enough.
- Correlation near threshold is common; the goal is not zero correlation but enough uniqueness with strong signal.
- Add a field-breadth metric to progress tracking.

### 9. Pyramid multiplier prior

Observed USA D1 multipliers include roughly:

- sentiment: 1.5
- other: 1.5
- model: 1.4
- option: 1.3
- earnings: 1.3
- analyst: 1.2
- news: 1.2
- pv/fundamental/socialmedia/shortinterest: 1.1

Scheduler rule:

- Multipliers are not submission gates, but they affect strategic value.
- Add `category_multiplier_bonus` to star priority, capped so it cannot overwhelm gate probability.

## Updated star priority term

```text
priority = star_score
         + uncertainty_bonus
         + decorrelation_bonus
         + community_prior_bonus
         + category_multiplier_bonus
         + strategic_bonus
         - over_sample_penalty
         - overfit_risk_penalty
```

## Community prior bonus

```text
community_prior_bonus =
  +0.08 if candidate follows WQB recommended practice for a current blocker
  +0.06 if candidate uses a documented dataset/operator handling rule correctly
  +0.05 if candidate tests a paper/community hypothesis without copying expression
  -0.10 if candidate is parameter fitting without new hypothesis
  -0.15 if candidate decorrelates by noise or unnecessary complexity
```

## Required before next dispatch

Before the next authorized simulation batch:

1. Run CNHK Yellow documentation/brief intake for candidate categories.
2. Hash candidate configurations and check duplicate cache.
3. Write slot-level community prior rationale.
4. Write expected information gain and kill rule.
5. Only then dispatch under explicit per-batch authorization.
