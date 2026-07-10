# WQB community baseline backfill — 2026-06-22

## Decision

Run a one-time baseline backfill before dispatching Batch F.

The forum workflow now solves ingestion:

```text
local CNHKMCP → metadata-only index → distill → GitHub artifact → Notion AI read
```

But it has not yet solved absorption:

```text
selected useful posts → claim-level summaries → source-claim-registry → star-map weights → Batch F revision
```

Therefore the next useful action is not another simulation dispatch. It is to convert the most relevant current and previously discovered community sources into scheduler-consumable claims.

## What is included

New baseline files:

- `runs/wqb-community/source-inventory-baseline.json`
- `runs/wqb-community/baseline-read-queue.json`
- `runs/wqb-current/local-forum-claim-extraction-request.json`

These combine:

1. the refreshed 50-post forum metadata index,
2. the generated clusters / read queue / digest,
3. older official-linked community URLs already discovered from WQB docs,
4. WQB official docs / Yellow-derived source packs,
5. worldquant-skill / CNHKMCP route artifacts.

## First-pass read priorities

The first pass should read only a small set and extract claims only:

1. anti-invalid-backtest / small-batch falsification,
2. simulation optimization,
3. prod_corr / virgin dataset pivot,
4. self-robustness → prod_corr staging,
5. prod_corr persistence / recovery workflow,
6. turnover reduction,
7. `trade_when` semantics,
8. concentrated weight diagnosis,
9. frequency-aware seed / CW pre-diagnosis / multi-GEM decorrelation,
10. sub-universe Sharpe repair,
11. overfitting avoidance,
12. dataset evaluation.

The Super Alpha combo post is intentionally postponed or method-only because it has high expression-copying risk.

## Safety rules

- Do not commit full forum bodies.
- Do not commit credentials, cookies, sessions, tokens, or auth headers.
- Do not commit private alpha expressions or Super Alpha combo expressions.
- Do not use forum write tools.
- Store only short paraphrased claims, failure modes, design effects, confidence, and source id/title.

## Expected output of the local/full-read phase

```text
runs/wqb-community/forum-claims-baseline.json
```

Each claim should look like:

```json
{
  "claim_id": "...",
  "claim": "short paraphrased rule, no copied expression",
  "failure_modes": ["PROD_CORRELATION"],
  "design_effect": "how scheduler/candidate grammar changes",
  "confidence": "full_read_claim",
  "do_not_copy_expression": true,
  "source_locator": "post id/title only"
}
```

## Scheduler impact before Batch F

Batch F remains proposal-only. Before dispatch, use the baseline claims to answer:

- Should behavioral decorrelation stay at 2 slots or be reduced/increased?
- Should fnd65 repair be delayed until CW / overfitting claims are extracted?
- Should scout slots shift toward dataset-evaluation-backed sources?
- Should duplicate-hash and anti-invalid-backtest checks be mandatory hard gates?
- Which current Batch F candidate is still shallow or unsupported?

No simulation is authorized by this backfill.
