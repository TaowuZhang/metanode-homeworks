# WQB Expression Draft Boundary Spec — 2026-06-24

Updated: 2026-06-24T13:51:47+08:00

## Boundary

This is a boundary spec only. It does not create an expression draft, FASTEXPR draft, authorization-ready artifact, simulation-ready packet, or submit path.

```text
expression_draft_allowed_next = true
expression_draft_created_now = false
fast_expression_created_now = false
```

## Accepted skeletons for next step

- `F02-SKEL-002`
- `F02-SKEL-003`
- `F04-SKEL-001`
- `F04-SKEL-002`
- `F04-SKEL-005`

## Slot constraints

F02 may have at most 2 future drafts and must keep entry/hold/exit interpretation explicit.

F04 may have at most 3 future drafts and must keep economic mechanism and data-quality role explicit.

## Not ready

```text
simulation_ready = false
batch_f_ready = false
authorization_ready_artifact_created = false
```