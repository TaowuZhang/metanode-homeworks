# WQB Field Family Search Plan — 2026-06-23

Updated: 2026-06-23T14:43:18+08:00

## Boundary

Read-only planning and lookup stage only. No simulation, no submit, no alpha-language generation, no Batch F ready packet, and no actual `field_refs` materialized.

## Sources

Used only:

```text
runs/wqb-current/field-ref-requirements-draft.json
runs/wqb-current/sanitized-slot-support-review.json
runs/wqb-current/p0a-repaired-claims-integrated-sanitized.json
```

Did not use raw extraction artifacts.

## F02 search plan

Role: `turnover_entry_hold_exit_diagnostic`

Support: `SCR-008`

Search objective: identify field families that can support information-arrival, state-change, entry, hold, and exit/no-trade interpretation. This is not a search for fields that mechanically lower turnover.

Queries planned/attempted:

```text
volume
liquidity
```

Preferred themes:

```text
trading activity
liquidity
volume
turnover-like semantic fields
event-state proxies
condition-state proxies
```

Forbidden themes:

```text
mechanical turnover reduction without signal logic
sparse coverage
fields likely to create concentration
fields copied from community examples
fields with no information-arrival interpretation
```

## F04 search plan

Role: `operator_grammar_economic_meaning_scout`

Support: `SCR-005`, `SCR-006`, `SCR-020`, `SCR-044`

Search objective: identify field families with economic meaning, operator-role explainability, data-quality reviewability, and either single-family discipline or coherent multi-family interaction.

Queries planned/attempted:

```text
fundamental
analyst
news sentiment
```

Preferred themes:

```text
fundamental fields with clear economic interpretation
analyst or estimate fields with temporal interpretation
sentiment/news fields only if interaction rationale exists
price-volume fields only if they serve explicit mechanism
fields with profileable missingness and distribution
```

Forbidden themes:

```text
community operator example mimicry
symbolic recipe reuse
noise injection
unexplained multi-family mixing
invalid-value heavy fields without cleaning plan
fields likely to create concentration artifacts
```