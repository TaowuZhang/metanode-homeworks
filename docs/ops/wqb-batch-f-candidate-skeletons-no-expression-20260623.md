# WQB Batch F Candidate Skeletons — No Expression — 2026-06-23

Updated: 2026-06-23T15:03:09+08:00

## Boundary

These are structural skeletons for duplicate/hash/validator precheck only.

They are not alpha-language bodies, not FASTEXPR, not simulation candidates, and not a Batch F ready packet.

## Skeleton count

```text
total = 9
F02 = 4
F04 = 5
```

## F02 skeletons

- `F02-SKEL-001`: `pv104_intv_mean` as trading-activity state proxy
- `F02-SKEL-002`: `pv104_ivam_mean` as intraday demand-pressure state proxy
- `F02-SKEL-003`: `bid_ask_price_gap` as spread/liquidity no-trade condition proxy
- `F02-SKEL-004`: `transaction_cost_estimate` as trading-cost suppression proxy

## F04 skeletons

- `F04-SKEL-001`: `fnd17_qastturn` as asset-turnover economic mechanism
- `F04-SKEL-002`: `fnd17_fcfmtt` as cash-flow quality mechanism
- `F04-SKEL-003`: `fnd17_rhsfcmtt` as cash-flow per-share mechanism
- `F04-SKEL-004`: `fnd85_sector_awcm_taa` as sector asset-efficiency context
- `F04-SKEL-005`: `snt22_2pos_mean_164` as event/sentiment interaction discipline