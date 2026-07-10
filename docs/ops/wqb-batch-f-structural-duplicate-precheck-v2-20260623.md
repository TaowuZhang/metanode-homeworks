# WQB Batch F Structural Duplicate Precheck v2 — 2026-06-24

Updated: 2026-06-24T13:51:47+08:00

## Boundary

This v2 precheck compares family-level metadata only. It does not compare or copy full expression text.

## Result

```text
input_skeleton_count = 9
accepted_for_expression_boundary = 0
accepted_with_limits_for_expression_boundary = 5
deferred = 4
blocked = 0
duplicate_precheck_v2_status = complete_family_level
```

## Accepted with limits

- `F02-SKEL-002`
- `F02-SKEL-003`
- `F04-SKEL-001`
- `F04-SKEL-002`
- `F04-SKEL-005`

## Deferred

- `F02-SKEL-001`
- `F02-SKEL-004`
- `F04-SKEL-003`
- `F04-SKEL-004`

## Main dedupe decisions

- Keep `F02-SKEL-002` over `F02-SKEL-001`.
- Keep `F02-SKEL-003` over `F02-SKEL-004`.
- Keep `F04-SKEL-001` over `F04-SKEL-004`.
- Keep `F04-SKEL-002` over `F04-SKEL-003`.
- Keep `F04-SKEL-005` only with event-state interaction limits.