# WQB Controlled Datafield Lookup Results — 2026-06-23

Updated: 2026-06-23T14:43:18+08:00

## Boundary

Read-only lookup only. No simulation, no submit, no correlation/submission checks, no alpha-language generation, and no actual `field_refs`.

## Tools used

```text
dataset_brief
```

No forbidden tools were used.

## Rate-limit note

The `volume` dataset query returned status `429`. Because `stop_on_429=true`, no datafield-level lookup was attempted. Other parallel dataset queries that had already been issued returned successfully and were recorded, but the stage stops at dataset-level shortlist.

## F02 results

Queries:

```text
volume -> 429
liquidity -> 200
```

Datasets considered:

- `order_book_imbalance` — liquidity, order flow, trade imbalance, spread and market microstructure dynamics.
- `pv104` — one-minute bar microstructure, spread, liquidity, volatility, trade size, volume and price statistics.
- `expected_move` — options-derived expected moves and volatility/liquidity context.
- `pv106` — spread levels and recent changes as liquidity/trading-cost state proxies.

No datafields were loaded.

## F04 results

Queries:

```text
fundamental -> 200
analyst -> 200
news sentiment -> 200
```

Datasets considered:

- `fundamental85`
- `fundamental17`
- `analyst27`
- `analyst4`
- `analyst10`
- `sentiment22`
- `news_sentiment_transfer`
- `news97`

No datafields were loaded after the 429 was observed.