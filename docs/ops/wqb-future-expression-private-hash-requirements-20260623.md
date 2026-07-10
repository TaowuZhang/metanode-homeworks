# WQB Future Expression Private Hash Requirements — 2026-06-24

Updated: 2026-06-24T13:51:47+08:00

If a future draft-only expression packet is created, it must run private/local hash and no-copy checks before any authorization packet can exist.

Required before any future simulation authorization packet:

- expression hash ready
- no-copy risk pass
- duplicate precheck pass or explained limits
- gate-aware score present
- kill rule present
- recovery controller present
- exact future user authorization supplied

Current state:

```text
no_expression = true
simulation_ready_now = false
submit_enabled = false
```