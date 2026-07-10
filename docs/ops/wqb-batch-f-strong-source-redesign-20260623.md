# WQB Batch F Strong-Source Redesign Draft — 2026-06-23

Status: draft only, not ready. No simulation. No submit. Not an authorization packet.

## Why not ready

Batch F now has safer slot roles and reviewed source references, but it still lacks filled `field_refs`, duplicate precheck, private/local hash precheck after materialization, authorization-packet validator, exact future authorization, and stronger extraction for several P0A sources.

Validation of this file only means design consistency, not readiness.

## Slots

- F01: prod_corr / originality mechanism pivot. Uses `SCR-011`. Not ready: field refs and duplicate/hash gates missing.
- F02: turnover entry-hold-exit diagnostic. Uses `SCR-043` and `SCR-008` as re-extraction target. Not ready: SCR-008 body evidence zero and field refs missing.
- F03: dataset / coverage / field-profile scout. Uses `SCR-041`. Not ready: field profiles missing.
- F04: operator grammar / economic meaning scout. Uses `SCR-005` and `SCR-006` as re-extraction targets. Not ready: generic claims.
- F05: Single Data / mixed-signal robustness diagnostic. Uses `SCR-020` as re-extraction target. Not ready: generic claim.
- F06: holdout / control / anti-waste. Uses workflow guardrails `SCR-037` and `SCR-049`. Allowed to have no strong alpha-design source because it is a holdout slot.