# WQB Batch F Structural Duplicate Precheck — 2026-06-23

Updated: 2026-06-23T15:03:09+08:00

## Boundary

Duplicate precheck used structural signatures only. It did not compare or store any full alpha-language body.

## Result

```text
skeletons_total = 9
pass = 0
pass_with_limits = 8
defer = 1
block = 0
status = partial
```

The status is partial because current Batch F and source-registry/starmap family metadata were reviewed, but prior Batch B/C/D/E family summaries were not loaded.

## Deferred

- `F04-SKEL-004`: overlaps `F04-SKEL-001` in the asset-turnover family and should not move forward until family dedupe chooses whether aggregate or direct asset-turnover context is useful.

## Main duplicate risks

- `F02-SKEL-001` and `F02-SKEL-002` both use `pv104` microstructure activity.
- `F02-SKEL-003` and `F02-SKEL-004` both use `pv106` spread/cost state.
- `F04-SKEL-001` and `F04-SKEL-004` both use asset-turnover family.
- `F04-SKEL-002` and `F04-SKEL-003` both use cash-flow family.