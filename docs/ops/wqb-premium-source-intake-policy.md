# WQB premium-source intake policy

## User correction

The user explicitly corrected the workflow priority:

> 把真正重要的精品文档搞进来，然后再说之后再怎么推进。一定要把流程设计的非常好，才能再做。如果帖子超过 50 个人点赞，一定要看。

This means the research system must pause premature Batch F execution and first build a high-quality source layer.

## Core rule

A source can enter the premium intake queue if it satisfies at least one of:

1. **High vote threshold**: forum `vote_count >= 50`.
2. **Official evergreen**: official WQB docs or official-linked community posts about gates, submissions, correlation, overfitting, turnover, weight coverage, simulation limits, or dataset evaluation.
3. **Current blocker match**: directly addresses current dominant blockers:
   - `PROD_CORRELATION`
   - `SIMULATION_BUDGET_WASTE`
   - `CONCENTRATED_WEIGHT`
   - `HIGH_TURNOVER`
   - `LOW_2Y_SHARPE`
   - `LOW_SUB_UNIVERSE_SHARPE`
   - `OVERFITTING`
4. **Architecture method**: improves evidence ledger, source claim extraction, duplicate checking, recovery, or experiment design.
5. **Dataset scout source**: provides structured dataset/field evaluation guidance before scout slots.

## Priority order

The intake priority is:

```text
P0: vote_count >= 50
P1: official evergreen + current blocker
P2: current blocker + strong title metadata
P3: architecture / workflow method
P4: dataset guides and scout priors
P5: expression-sharing posts, method-only and deferred
```

## Hard gates before simulation

Before any new simulation batch:

1. All available `P0` sources must be at least metadata-indexed and queued for claim extraction.
2. Any `P0` source directly relevant to current blockers must be claim-extracted or explicitly marked blocked/deferred with reason.
3. Batch F cannot be described as community-informed until P0/P1 source claims are integrated.
4. Expression-sharing posts cannot be copied; they may only produce method-level claims.
5. High-vote tool/automation posts do not authorize autonomous bulk mining.

## Output standard

Premium sources must be stored as claims, not bodies:

```json
{
  "claim_id": "...",
  "source_id": "...",
  "title": "...",
  "vote_count": 91,
  "claim": "short paraphrased claim",
  "failure_modes": ["..."],
  "design_effect": "...",
  "confidence": "full_read_claim | metadata_only | blocked",
  "do_not_copy_expression": true
}
```

## Consequence for current run

The current run should move from:

```text
ready to dispatch Batch F after shallow priors
```

to:

```text
premium-source intake first
→ high-vote / official evergreen claims
→ star-map and Batch F redesign
→ then consider simulation authorization
```