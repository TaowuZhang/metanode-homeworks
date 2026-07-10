# WQB Star-map Math Model

## Status

This is the missing mathematical layer between the philosophical star-map plan and actual batch selection. The system must not rely on vague judgment such as "promising" without numeric evidence.

## Objects

### Star

```text
star = region × universe × delay × dataset × field_family × field_type × operator_family × horizon_bucket × neutralization_family
```

### Constellation

```text
constellation = dataset_group × economic_hypothesis × operator_archetype
```

### Candidate

A candidate is one sampled expression inside a star. Candidate results update the star. The star, not the individual expression, drives the scheduler.

## Gate vector

For every alpha, convert official WQB checks into a gate vector:

```json
{
  "low_sharpe": { "pass": 0|1, "value": number, "limit": number },
  "low_fitness": { "pass": 0|1, "value": number, "limit": number },
  "turnover": { "pass": 0|1, "value": number, "low_limit": 0.01, "high_limit": 0.7 },
  "concentration": { "pass": 0|1 },
  "sub_universe": { "pass": 0|1, "value": number, "limit": number },
  "self_correlation": { "pass": 0|1, "value": number, "limit": number },
  "prod_correlation": { "pass": 0|1, "value": number, "limit": number },
  "two_year_sharpe": { "pass": 0|1, "value": number, "limit": number },
  "data_diversity": { "pass": 0|1 },
  "power_pool_correlation": { "pass": 0|1 }
}
```

## Normalized component scores

### Signal score

```text
sharpe_score  = clamp(sharpe / sharpe_limit, -1, 1.5) / 1.5
fitness_score = clamp(fitness / fitness_limit, -1, 1.5) / 1.5
signal_score  = 0.65 * sharpe_score + 0.35 * fitness_score
```

Reason: Sharpe often moves first; Fitness is partly downstream of signal quality and turnover.

### Stability score

```text
sub_universe_score = clamp(sub_universe_sharpe / sub_universe_limit, -1, 1.5) / 1.5
recent_score       = clamp(two_year_sharpe / sharpe_limit, -1, 1.5) / 1.5
stability_score    = 0.45 * sub_universe_score + 0.55 * recent_score
```

Reason: `LOW_2Y_SHARPE` is currently the repeated long-term blocker, so it receives slightly higher weight.

### Originality score

Correlation is good when lower than limit. Unknown is not pass; it is uncertainty.

```text
self_corr_score = if known then clamp((limit - self_corr) / limit, -1, 1) else 0.25
prod_corr_score = if known then clamp((limit - prod_corr) / limit, -1, 1) else 0.25
power_score     = 1 if pass else 0
originality_score = 0.25*self_corr_score + 0.55*prod_corr_score + 0.20*power_score
```

Reason: prod correlation has been the primary blocker for behavioral D02 repairs.

### Tradability score

```text
turnover_band_score = 1 - normalized_distance_from_preferred_band(turnover, preferred=[0.05,0.45], hard=[0.01,0.7])
concentration_score = 1 if CONCENTRATED_WEIGHT pass else 0
tradability_score   = 0.45 * turnover_band_score + 0.55 * concentration_score
```

### Data and platform fit

```text
platform_score = mean(blocking_gate_passes excluding signal/stability/originality/tradability)
```

## Alpha gate-aware score

```text
alpha_score = 0.25 * signal_score
            + 0.22 * stability_score
            + 0.22 * originality_score
            + 0.16 * tradability_score
            + 0.10 * platform_score
            + 0.05 * simplicity_score
            - blocker_penalty
```

### Blocker penalty

```text
blocker_penalty = 0.10 * LOW_SHARPE_FAIL
                + 0.08 * LOW_FITNESS_FAIL
                + 0.10 * LOW_2Y_SHARPE_FAIL
                + 0.12 * PROD_CORRELATION_FAIL
                + 0.08 * SELF_CORRELATION_FAIL
                + 0.07 * CONCENTRATED_WEIGHT_FAIL
                + 0.05 * LOW_SUB_UNIVERSE_FAIL
                + 0.20 * EXPRESSION_OR_POLICY_FAIL
```

## Star score

A star aggregates completed candidates:

```text
star_signal_mean = mean(top_k(alpha_score, k=min(3,n)))
star_best        = max(alpha_score)
star_improvement = best_recent_alpha_score - prior_best_alpha_score
star_score       = 0.55 * star_best + 0.30 * star_signal_mean + 0.15 * max(0, star_improvement)
```

## Exploration bonus

Use UCB-like pressure so the system does not overfit to the first strong ridge:

```text
uncertainty_bonus = c * sqrt(log(total_completed_stars + 1) / (star_completed_count + 1))
```

Default `c = 0.18`.

## Over-sampling penalty

```text
over_sample_penalty = 0.04 * max(0, star_completed_count - 3)
                    + 0.08 * repeated_no_improvement_batches
                    + 0.10 * repeated_same_blocker_batches
```

## Scheduler priority

```text
priority = star_score
         + uncertainty_bonus
         + decorrelation_bonus
         + strategic_bonus
         - over_sample_penalty
```

### Strategic bonus

```text
+0.10 if star belongs to underexplored constellation
+0.12 if star can address current global blocker
+0.08 if datafield is low userCount/alphaCount
+0.15 if near-gate candidate exists in star
-0.12 if last two batches from constellation failed without information gain
```

## Batch slot optimization

For a 6-slot batch, solve a constrained greedy allocation:

```text
maximize sum(priority(candidate_slot))
subject to:
  at least 2 slots broad scout unless near_gate exists
  at most 2 slots same field unless near_gate exists
  at most 3 slots same constellation unless near_gate exists
  at least 1 diagnostic slot after repeated blocker
  no candidate without expected_information_gain
```

## Heatmap statistics required before each batch

Before dispatching any batch, write:

1. sampled stars count;
2. unsampled priority stars;
3. completed candidates per constellation;
4. dominant blocker histogram;
5. best alpha per constellation;
6. repeated blocker count;
7. slot allocation rationale.

If these are missing, the system is doing tool work, not research.

## Final-click criterion remains hard

The score only schedules search. It never overrides WQB gates. A final-click candidate requires all blocking WQB checks/correlation gates pass, with submit left manual.
