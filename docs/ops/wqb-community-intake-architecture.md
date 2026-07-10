# WQB Community Intake Architecture

## Why this exists

Official WQB documentation is not enough. The active community posts contain the living operating knowledge: how people actually reduce correlation, improve Sharpe, handle weight coverage, avoid overfitting, evaluate datasets, and work around the consultant environment without violating platform rules.

The previous scaffold used CNHK MCP mostly as an execution/recovery tool. That was insufficient. The MCP and the public/community support pages must become an intake layer that feeds the belief graph.

## Source hierarchy

1. WQB platform / CNHK Yellow docs and tools
   - `get_documentations`
   - `get_documentation_page`
   - `dataset_brief`
   - `field_brief`
   - `get_pyramid_multipliers`
   - `get_leaderboard`

2. WQB community posts
   - support.worldquantbrain.com community posts linked from official docs.
   - These are not expression libraries; they are living practice notes.

3. User/local Cloudflare runner
   - If public web fetch times out or recaptcha blocks crawler, ask the user to run local fetch commands or expose a controlled read-only Cloudflare helper.
   - Never ask for credentials or cookies in chat.

4. GitHub / local repo memory
   - Community post metadata, extracted rules, source URLs, fetch status, and derived scheduler priors are stored in Git.
   - Do not store private alpha code.

## Community intake pipeline

```text
community URL registry
  -> local/web fetch attempt
  -> raw text snapshot or fetch failure
  -> rule extraction
  -> prior classification
  -> belief graph update
  -> scheduler weight update
  -> candidate slot rationale
```

## Rule extraction schema

Each community source is converted to:

```json
{
  "source_id": "wqb-community-reduce-correlation",
  "url": "https://support.worldquantbrain.com/...",
  "fetch_status": "pending|fetched|blocked|timeout",
  "topic": "correlation",
  "claims": [
    {
      "claim": "Mechanism change is preferred over noise for correlation reduction.",
      "confidence": "medium",
      "actionability": "high",
      "maps_to_failure_modes": ["PROD_CORRELATION", "SELF_CORRELATION"],
      "scheduler_effect": "increase mechanism_pivot slots after repeated correlation fail"
    }
  ],
  "do_not": ["do not copy expressions", "do not use noise to decorrelate"],
  "candidate_implications": ["add diagnostic operator sibling", "switch dataset family"],
  "last_checked_at": "ISO8601"
}
```

## Community prior types

- `sharpe_improvement_prior`
- `returns_improvement_prior`
- `correlation_reduction_prior`
- `turnover_reduction_prior`
- `pnl_smoothing_prior`
- `overfitting_avoidance_prior`
- `dataset_evaluation_prior`
- `coverage_weight_prior`
- `neutralization_prior`
- `event_alpha_prior`

## How this changes the WQB architecture

The scheduler must not generate a batch from only our internal intuition. Before every batch it must check:

1. Does any community prior directly address the dominant blocker?
2. Is the next batch repeating a pattern the community warns against?
3. Are we missing an official/community method for the exact failure mode?
4. Should we run a diagnostic rather than another repair?
5. Is a local runner needed to fetch fresh community pages?

## Fetch failures are state, not excuses

In this session, public web loads of several support.worldquantbrain.com community posts returned crawler timeouts or cookie/recaptcha pages. That means:

- The community intake layer must record fetch failure.
- The system must prepare local fetch requests for the user/local runner.
- The system must not pretend those posts were read.
- The scheduler can use only titles/official doc references until content is fetched.

## Local runner contract

If user can run local Cloudflare / local browser / local MCP, provide a request file:

```text
runs/wqb-current/local-community-fetch-request.json
```

Expected output:

```text
runs/wqb-community/raw/<source_id>.html or .md
runs/wqb-community/extracted/<source_id>.json
```

Minimum fields:

```json
{
  "source_id": "...",
  "url": "...",
  "fetched_at": "...",
  "title": "...",
  "text": "...",
  "links": []
}
```

No cookies or credentials should be committed.

## Hard rule

Community content is a prior, not an authority. It modifies search distribution; it does not override WQB gates or platform policy.
