# WQB Candidate Field Materialization Validation Report — 2026-06-23

Updated: 2026-06-23T14:52:42+08:00

## Checks

- NO_SIMULATION_OR_SUBMIT: pass
- READ_ONLY_LOOKUP_ONLY: pass
- NO_EXPRESSION_OR_FASTEXPR: pass
- NO_BATCH_READY_OR_AUTHORIZATION_READY: pass
- ONLY_ACCEPTED_SLOT_SUPPORT_USED: pass
- NO_BLOCKED_SOURCE_USED: pass
- MATERIALIZED_ONLY_FOR_PRECHECK: pass
- RATE_LIMIT_RESPECTED: pass

## Rate limit

`analyst27` field lookup returned `429`; no further datafield calls were made after that result.

## Current state

```text
batch_f_ready = false
simulation_ready = false
simulation_dispatched = false
submit_attempted = false
```