# WQB Controlled Datafield-Level Lookup Results — 2026-06-23

Updated: 2026-06-23T14:52:42+08:00

## Boundary

Read-only datafield-level lookup only. No simulation, no submit, no correlation/submission checks, no forum writes, and no alpha-language generation.

## Tools used

```text
field_brief
```

## Rate limit

`field_brief` on `analyst27` returned `429`. The `sentiment22` request was already issued in parallel and is recorded. No further datafield-level calls were made after the 429 result.

## F02

Queried:

```text
pv104
pv106
```

Field briefs loaded for microstructure volume, liquidity, spread, and transaction-cost families.

## F04

Queried:

```text
fundamental85
fundamental17
analyst27
sentiment22
```

`analyst27` was rate-limited. Fundamental and sentiment field briefs were loaded.