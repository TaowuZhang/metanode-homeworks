# WQB Field Ref Requirements Draft — 2026-06-23

Updated: 2026-06-23T14:29:36+08:00

## Boundary

This artifact is **requirements only**.

It does not perform field/datafield lookup, does not fill actual `field_refs`, does not generate alpha-language details, does not create a Batch F ready packet, does not dispatch simulation, and does not submit.

## Sources used

Only these downstream sources were used:

```text
runs/wqb-current/sanitized-slot-support-review.json
runs/wqb-current/p0a-repaired-claims-integrated-sanitized.json
```

These raw / pre-sanitized sources were not used:

```text
runs/wqb-current/local-targeted-extractor-output.json
runs/wqb-current/p0a-repaired-claims-integrated.json
```

## Slot scope

Allowed into requirements drafting:

```text
F02
F04
```

Deferred:

```text
F03
```

Blocked / not scheduler eligible sources remain blocked:

```text
SCR-003
SCR-004
SCR-016
SCR-021
SCR-013
SCR-034
```

High-copy-risk sources remain method-only / not scheduler eligible:

```text
SCR-004
SCR-013
SCR-016
SCR-021
```

---

## F02 requirements summary

```json
{
  "slot_id": "F02",
  "slot_role": "turnover_entry_hold_exit_diagnostic",
  "support_refs": ["SCR-008"],
  "field_ref_stage": "requirements_only",
  "actual_field_refs": [],
  "required_field_family_traits": [
    "can express a plausible information-arrival or state-change mechanism",
    "supports entry/hold/exit interpretation",
    "not merely a parameter-only decay/turnover fix",
    "has enough coverage in USA TOP3000 Delay1",
    "can be profiled for turnover impact before simulation"
  ],
  "preferred_field_families": [
    "turnover / trading activity / liquidity / volume-like semantic families",
    "event-state or condition-state proxies",
    "fundamental activity fields only if they have clear temporal update logic"
  ],
  "forbidden_field_families": [
    "fields chosen only because they reduce turnover mechanically",
    "fields with sparse coverage that would cause concentrated weight",
    "fields without interpretable information arrival",
    "fields copied from community examples or templates"
  ],
  "operator_requirement": [
    "must support entry-hold-exit reasoning",
    "must preserve signal meaning while reducing unnecessary trading",
    "must explicitly reject decay-only repair if no mechanism exists"
  ],
  "gate_targets": [
    "HIGH_TURNOVER",
    "LOW_FITNESS",
    "LOW_SUB_UNIVERSE_SHARPE",
    "CONCENTRATED_WEIGHT"
  ],
  "pre_lookup_questions": [
    "What is the information event?",
    "What condition opens the position?",
    "What condition holds the position?",
    "What condition exits or suppresses trading?",
    "How can turnover be reduced without destroying fitness?"
  ],
  "next_action": "field_family_search_plan"
}
```

F02 may proceed to a future **field family search plan**, not direct lookup.

---

## F04 requirements summary

```json
{
  "slot_id": "F04",
  "slot_role": "operator_grammar_economic_meaning_scout",
  "support_refs": ["SCR-005", "SCR-006", "SCR-020", "SCR-044"],
  "field_ref_stage": "requirements_only",
  "actual_field_refs": [],
  "required_field_family_traits": [
    "has a clear economic mechanism before operator construction",
    "can be reasoned about without copying community examples",
    "supports data-cleaning / invalid-value / robustness review",
    "supports either single-family discipline or coherent multi-family interaction",
    "has coverage and distribution suitable for USA TOP3000 Delay1"
  ],
  "preferred_field_families": [
    "fields with interpretable economic meaning",
    "fields whose transformation role can be stated before simulation",
    "mixed data families only if interaction has clear financial rationale",
    "operator-friendly fields where missing/invalid values can be profiled"
  ],
  "forbidden_field_families": [
    "fields selected only to mimic community operator examples",
    "fields requiring symbolic recipe reuse",
    "fields mixed only for noise injection",
    "fields with no explainable economic link",
    "fields that are likely to create invalid-value or concentration artifacts"
  ],
  "operator_requirement": [
    "operator must serve an economic or data-quality role",
    "operator complexity must be penalized unless mechanism requires it",
    "invalid/missing data handling must be specified before candidate construction",
    "mixed-signal design must explain why data families interact"
  ],
  "gate_targets": [
    "OVERFITTING",
    "LOW_FITNESS",
    "CONCENTRATED_WEIGHT",
    "LOW_SUB_UNIVERSE_SHARPE",
    "PROD_CORRELATION"
  ],
  "pre_lookup_questions": [
    "What economic mechanism is being represented?",
    "What transformation role is needed?",
    "What data-cleaning risks exist?",
    "Is this single-family, or a justified multi-family interaction?",
    "How will this avoid arbitrary operator stacking?"
  ],
  "next_action": "field_family_search_plan"
}
```

F04 may proceed to a future **field family search plan**, not direct lookup.

---

## F03 handling

```json
{
  "slot_id": "F03",
  "slot_role": "dataset_coverage_field_profile_scout",
  "support_refs": ["SCR-002"],
  "status": "deferred",
  "reason": "SCR-002 is a broad research/dataset scout prior. It needs concrete dataset questions and mechanism hypotheses before field-ref requirements.",
  "next_action": "dataset_question_decomposition_before_field_refs"
}
```

F03 remains blocked from field-ref requirements until the dataset-scout prior is decomposed into concrete dataset questions and mechanism hypotheses.

## Validation

```json
[
  {
    "check_id": "FIELD_REQUIREMENTS_ONLY_NO_LOOKUP",
    "status": "pass",
    "rule": "This artifact may define field family requirements but must not include actual field_refs or datafield IDs."
  },
  {
    "check_id": "ONLY_ACCEPTED_SUPPORT_USED",
    "status": "pass",
    "allowed_slots": ["F02", "F04"],
    "blocked_slots": ["F03"],
    "rule": "Only accepted_with_limits support from sanitized-slot-support-review may enter field-ref requirements."
  },
  {
    "check_id": "NO_BLOCKED_SOURCE_USED",
    "status": "pass",
    "blocked_sources": ["SCR-003", "SCR-004", "SCR-016", "SCR-021", "SCR-013", "SCR-034"]
  },
  {
    "check_id": "NO_SIMULATION_OR_SUBMIT",
    "status": "pass",
    "simulation_dispatched": false,
    "submit_attempted": false,
    "batch_f_ready": false
  }
]
```

## Current Batch F status

```text
batch_f_state = draft_only_not_ready
batch_f_not_ready_stage = field_ref_requirements_complete_field_lookup_pending
batch_f_ready = false
simulation_ready = false
simulation_dispatched = false
submit_attempted = false
```

## Next blockers

```text
field_refs not filled
field/datafield lookup not run
field family search plan not written
duplicate precheck not run
private/local hash precheck not built
validator not run
exact future user authorization not supplied
```

## Next step

```text
Field Family Search Plan / Datafield Lookup Plan
```

That next step should still be a plan for lookup before any actual field/datafield tool call.