# WQB Star-map Bandit Scheduler

## Why this exists

The previous loop was too close to a brute pipeline: propose a small batch, dispatch, recover, score, repeat. That is not enough for a search space with thousands of possible stars. The scheduler must decide where to look before deciding what to simulate.

This document upgrades the scaffold from a batch ledger into a star-map search system.

## Core model

One candidate is not the unit of strategy. A star is a searchable cell:

```text
star = dataset × field_family × field_type × operator_family × horizon_bucket × neutralization × objective
```

A constellation is a group of related stars:

```text
constellation = economic_hypothesis × dataset_group × operator_archetype
```

A batch is not just six expressions. A batch is a deliberate sample from the star map.

## Scheduler loop

```text
1. update star-map from all completed gates
2. classify each constellation: unknown / weak / promising / repair_seed / near_gate / exhausted
3. allocate slots by bandit policy
4. generate population candidates only inside allocated stars
5. dispatch only after explicit per-batch authorization
6. recover, score, update terrain
```

## Slot allocation

Default six-slot batch:

| Slot type | Count | Purpose |
|---|---:|---|
| scout | 2 | broad exploration across low-sampled stars |
| climb | 2 | mutate promising parent / repair seed |
| repair | 1 | target a specific blocker |
| diagnostic | 1 | sign / inverse / operator pivot to learn terrain |

When there is a near-gate alpha:

| Slot type | Count |
|---|---:|
| confirm / repair | 3 |
| climb | 2 |
| scout | 1 |

When the current family fails two local repair batches:

| Slot type | Count |
|---|---:|
| scout | 3 |
| diagnostic | 2 |
| climb | 1 |

## Bandit scoring

Each star gets a priority score:

```text
star_priority = uncertainty_bonus
              + signal_potential
              + gate_repairability
              + decorrelation_value
              + recency_stability_value
              - over_sample_penalty
              - repeated_blocker_penalty
```

### uncertainty_bonus
High when a star has few or no completed samples. This prevents tunnel vision.

### signal_potential
Based on best Sharpe, Fitness, sub-universe Sharpe, and partial gate passes.

### gate_repairability
High when failures are actionable, e.g. turnover / concentration / horizon stability. Low when repeated prod correlation persists under mechanism changes.

### decorrelation_value
High for underexplored datasets, low userCount / alphaCount fields, different operator families, or a mechanism far from the current parent.

### over_sample_penalty
Penalizes repeating the same field/operator/horizon family without new information.

### repeated_blocker_penalty
Penalizes a family after the same blocking gates remain unchanged for multiple batches.

## Broad vs deep exploration

### Broad exploration
Broad exploration means selecting different stars, not randomly generating expressions.

Examples:
- different dataset family;
- same dataset but different economic field family;
- same field family but different operator archetype;
- same mechanism but different horizon bucket;
- low-crowding zero-user fields;
- inverse diagnostic when all direct variants are weak.

### Deep exploration
Deep exploration is allowed only after a star has evidence:

- at least one alpha passes LOW_SHARPE or is close;
- at least one non-trivial gate passes;
- failures are not all hard blockers;
- parent has a clear repair path.

Deep exploration is bounded:

- max two repair batches unless improving;
- max one local micro-mutation batch after a failed local repair;
- after repeated prod correlation, require mechanism pivot rather than parameter tuning.

## Population climber contract

A population batch must contain:

1. parent/control or near-control;
2. one primary repair mutation;
3. one mechanism mutation;
4. one horizon/stability mutation;
5. one decorrelation mutation;
6. one diagnostic / scout candidate.

If all six candidates share the same field and only differ by decay/horizon/truncation, the batch is too narrow unless the parent is near-gate.

## Stop and pivot rules

Pivot away from a family when:

- two consecutive batches improve no blocking gate;
- prod correlation worsens under neutralization pivots;
- the best candidate remains below Fitness 1 after repair attempts;
- the family consumes budget without producing a near-gate alpha;
- the batch is becoming parameter twiddling rather than hypothesis search.

Do not stop the whole system. Stop the family, update the star map, and allocate scout slots elsewhere.

## Batch F refactor

After Batch E, the behavioral D02 local repair failed. The next batch should not be a pure D02 micro-mutation. It should be a mixed star-map batch:

- 2 behavioral decorrelation / mechanism pivots;
- 2 fnd65 repair / stability probes;
- 1 ai_factor_transfer operator pivot, not the failed rank_ts_zscore template;
- 1 analyst diagnostic / inverse sample.

This preserves depth while restoring breadth.

## Platform boundary

The scheduler may prepare candidates and ledgers continuously, but it may not dispatch simulations without explicit per-batch authorization required by WQB policy. When blocked from dispatch, it continues non-simulation work instead of empty stopping.
