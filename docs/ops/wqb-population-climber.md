# WQB Population Climber

## Purpose

The population climber is the deep-search engine. It does not replace broad star-map search; it is activated only for stars with evidence.

## Candidate roles

Every climbing population must label each candidate with one of these roles:

- `control`: baseline / parent replay for comparability;
- `repair`: targets a known blocker;
- `mechanism_pivot`: changes economic or operator mechanism;
- `decorrelation`: reduces self/prod/power-pool crowding;
- `stability`: targets LOW_2Y_SHARPE or sub-universe fragility;
- `diagnostic`: tests sign, inverse, horizon, or operator ambiguity;
- `scout`: broad star-map probe.

A batch with six candidates all labeled `repair` is usually over-narrow.

## Parent eligibility

A parent may be climbed only if:

```text
(parent_sharpe >= gate_limit or close)
and at least one originality/correlation gate is pass or unknown
and dominant failure has a repair path
and the parent is not merely a one-off luck sample
```

## Climb depth budget

| State | Max extra batches | Rule |
|---|---:|---|
| weak_signal | 0-1 | mostly scout elsewhere |
| promising | 1-2 | test siblings / mechanism pivots |
| repair_seed | 1-2 | repair blockers, require improvement |
| near_gate | 2-3 | intense repair / confirmation |
| repeated_prod_corr | 0 local | mechanism pivot only |

## Improvement definition

A repair batch counts as improvement only if it does at least one of:

- turns a blocking gate from FAIL to PASS;
- materially improves a blocker without breaking another key gate;
- lowers prod/self correlation while retaining signal;
- improves 2Y or sub-universe stability while preserving Sharpe;
- reveals a sibling field or operator family with cleaner failures.

Small parameter changes that leave the same blockers are not improvement.

## Anti-bruteforce rules

- Do not enumerate horizons blindly.
- Do not keep the same field/operator and only vary decay/truncation unless near-gate.
- Do not continue a family after two no-improvement batches.
- Do not let engineering convenience choose candidates; the star map chooses candidates.
- Do not mistake completed simulations for research progress unless they update terrain.

## Output contract

Each candidate record must include:

```json
{
  "candidate_id": "F01",
  "role": "mechanism_pivot",
  "constellation": "behavioral_extrapolation_bias",
  "star_id": "...",
  "parent_alpha_id": "optional",
  "field_refs": ["..."],
  "operator_family": "...",
  "expected_information_gain": "...",
  "expected_failure_addressed": ["..."],
  "kill_rule": "...",
  "expression_ref": "redacted:F01"
}
```

The expression may be materialized only for dispatch, and private full expression text is not stored in long-term Git ledgers.
