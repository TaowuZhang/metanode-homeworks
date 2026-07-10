# WQB Source Intelligence Scoring Model

Updated: 2026-06-23T13:18:56+08:00

Scope: source intelligence / system design only. This artifact does **not** dispatch WQB simulation, submit alpha, call forum write tools, store credentials/cookies/sessions, store private alpha full expressions, or copy community expressions.

## Required score fields

`source_importance`, `evidence_strength`, `execution_relevance`, `copy_risk`, `policy_risk`, `workflow_only_score`, `alpha_design_score`, `portfolio_method_score`, `research_prior_score`, `expression_scheduler_usability`, `method_scheduler_usability`, `confidence_reason`.

## Formulas

```text
source_importance =
  0.35 * log_vote_score
+ 0.20 * comment_signal
+ 0.15 * recency_score
+ 0.15 * official_or_evergreen_bonus
+ 0.15 * repeated_search_hit_bonus
```

```text
evidence_strength =
  0 if body_chars_seen_in_memory = 0
  else min(1, ln(1 + body_chars_seen_in_memory) / ln(1 + 50000))
       * confidence_multiplier
       * non_generic_claim_multiplier
```

```text
execution_relevance =
  evidence_strength
  * max(alpha_design_score, portfolio_method_score, research_prior_score)
  * gate_mapping_quality
  * allowed_use_multiplier
```

```text
expression_scheduler_usability =
  execution_relevance
  * evidence_strength
  * (1 - copy_risk)
  * (1 - policy_risk)
  * expression_allowed_multiplier
  * claim_quality_multiplier
```

```text
method_scheduler_usability =
  evidence_strength
  * max(portfolio_method_score, research_prior_score, workflow_guardrail_score, alpha_design_score * method_only_multiplier)
  * (1 - policy_risk)
  * claim_quality_multiplier
```

## Non-negotiable rules

1. High votes are importance, not execution evidence.
2. `read_status=read` is not equivalent to understood.
3. `body_chars_seen_in_memory=0` forces `evidence_strength=0`.
4. Generic claims cannot be promoted for validation.
5. Workflow / MCP / automation posts are process guardrails only; they never authorize autonomous bulk mining.
6. Template / SuperAlpha / Combo / code-sharing posts are high copy-risk and method-only.
7. Official docs and community sources merge into one claim graph but retain different priors.
8. Validation pass means design-package validity only; it does not mean Batch F ready or simulation authorized.

## Missing value policy

- Missing vote count: set to 0 and mark `source_importance_low_confidence`.
- Missing comment count: set to 0 and mark `comment_signal_low_confidence`.
- Missing body chars: treat as 0.
- Missing claim quality: treat as generic.
- Ambiguous source type: use the most conservative bucket.
- Missing allowed use: set `scheduler_eligible=false`.
- Missing copy/policy risk: infer conservatively; if unknown, set 0.50.

## Source trust

Prior initialization:

- Official docs: high gate/policy prior.
- Full-read specific community claim: medium/high method prior.
- Full-read generic community claim: social importance may remain high but evidence prior is low.
- Metadata-only community source: zero evidence trust.
- Workflow/policy source: guardrail trust only, not expression trust.

Future posterior update:

```text
posterior_trust =
  prior_trust
+ validated_extractor_gain
+ gate_prediction_accuracy
+ reproducible_batch_learning_gain
- contradiction_penalty
- copy_or_policy_risk_penalty
```

No posterior update from simulation exists in this package because no simulation was dispatched.