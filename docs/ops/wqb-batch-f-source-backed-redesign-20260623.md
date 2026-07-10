# WQB Batch F source-backed redesign — 2026-06-23

## Status

No simulation dispatched.

User completed local premium-source extraction and repaired integration at commit `dbaa3db7886f3c3e52c9eb90e1dbb10ab00f15e2`.

The old Batch F packet is stale. This redesign uses the premium official/community claim layer rather than the previous local repair loop.

## Inputs

- `runs/wqb-current/source-claim-registry.json`
- `runs/wqb-current/starmap-weight-model.json` v5
- `runs/wqb-current/forum-claims-integration-summary.json`
- `runs/wqb-community/premium-source-claims.json`
- `docs/ops/wqb-premium-forum-claims-review-20260623.md`

Forum claim emphasis after repair:

```json
{
  "SIMULATION_BUDGET_WASTE": 5,
  "LOW_SUB_UNIVERSE_SHARPE": 3,
  "PROD_CORRELATION": 2,
  "SELF_CORRELATION": 2,
  "OVERFITTING": 2,
  "CONCENTRATED_WEIGHT": 2,
  "LOW_FITNESS": 2,
  "HIGH_TURNOVER": 1,
  "LOW_SHARPE": 1,
  "LOW_2Y_SHARPE": 1,
  "SOURCE_DRIFT": 2
}
```

## Design change from old Batch F

Old Batch F over-allocated to local repair of known parents. Premium sources now imply:

1. **Anti-waste first**: no candidate without expected information gain, duplicate precheck, and kill rule.
2. **Correlation originality**: after repeated prod_corr failures, prefer mechanism/data-family/robustness/originality pivots over local parameter tweaks.
3. **Dataset scout discipline**: scout slots need field profiling / coverage / frequency / bounds / distribution claims before expensive simulation.
4. **Overfitting discipline**: 2Y/sub-universe failures require regime/liquidity diagnostics, not blind parameter expansion.
5. **Turnover grammar**: turnover repair requires entry/hold/exit or trade_when-style structure, not decay-only mutation.
6. **Expression firewall**: Super Alpha combo posts are method-only; no copied expressions.

## New slot allocation

Safe batch size remains 6, but Batch F should be treated as **proposal-only** until validators pass and exact authorization is supplied.

| Slot | Role | Source-backed purpose | Primary claims |
|---|---|---|---|
| F01 | correlation mechanism pivot | Test non-local behavioral originality using a mechanism/data-family pivot rather than D02 micro-repair | reduce-correlation, Super Alpha method-only |
| F02 | dataset scout probe | Test price-momentum / price-volume scout only after field profile rationale | Term #14 Price Momentum, evaluate-new-dataset |
| F03 | overfitting/stability diagnostic | Test whether a candidate family survives stability/sub-universe logic before climbing | avoid-overfitting, alpha-submission |
| F04 | CW / coverage diagnostic | Diagnose breadth vs max-weight concentration before fnd65 repair | weight coverage queue + official weight-test claims |
| F05 | turnover grammar diagnostic | Test trade_when-style entry/exit structure if turnover/fitness tradeoff is relevant | trade_when evergreen, high-turnover docs |
| F06 | workflow control / anti-waste sentinel | Preserve one slot as a non-simulation or ultra-low-risk control unless all hard gates pass | automation posts process-only, simulation budget claims |

## Candidate design principles

- No full FASTEXPR in Git.
- Each slot must have `expression_ref`, not expression text.
- Each slot must have:
  - source claims
  - expected information gain
  - kill rule
  - duplicate/precheck requirement
  - no-expression-copying guard
- F06 should remain non-dispatched unless a concrete low-risk diagnostic is justified.

## Required before any dispatch

1. Retry or waive weak P0 Weight Coverage source.
2. Run/record structural duplicate precheck.
3. Create private/local expression hash precheck after materialization.
4. Validate the authorization packet.
5. Reconfirm exact user authorization: `授权 Batch F，confirmSimulation=true`.

## Current recommendation

Do **not** dispatch yet. The correct next step is to convert this redesign into a validated source-backed authorization packet and run non-simulation validators.
