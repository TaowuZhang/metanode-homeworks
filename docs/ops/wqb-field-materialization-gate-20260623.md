# WQB Field Materialization Gate — 2026-06-23

Updated: 2026-06-23T14:52:42+08:00

## Boundary

Field refs are materialized only as precheck inputs. They are not simulation-ready fields and not a Batch F ready packet.

## F02

Accepted for precheck:

```text
pv104_intv_mean
```

Accepted with limits:

```text
pv104_ivam_mean
bid_ask_price_gap
transaction_cost_estimate
```

## F04

Accepted for precheck:

```text
fnd17_qastturn
fnd17_fcfmtt
```

Accepted with limits:

```text
fnd85_sector_awcm_taa
fnd17_rhsfcmtt
snt22_2pos_mean_164
```

Deferred:

```text
analyst27
analyst4
analyst10
news_sentiment_transfer
news97
snt22_3pos_mean_116
```

## Status

```text
batch_f_ready = false
simulation_ready = false
simulation_dispatched = false
submit_attempted = false
```