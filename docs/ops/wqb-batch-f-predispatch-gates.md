# WQB Batch F pre-dispatch gates

## Purpose

Batch F is not ready to simulate just because an authorization packet exists. Before any `create_multi_simulation` call, the packet must pass three layers:

```text
authorization packet validation
→ duplicate / expression-hash precheck
→ forum claim review / scheduler slot review
```

This runbook adds the missing pre-dispatch gate between source ingestion and simulation.

## Gate 1 — authorization packet validator

Run:

```bash
node scripts/wqb-validate-authorization-packet.mjs \
  runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json
```

The packet must satisfy:

- `submitEnabled=false`
- `autonomousBulkMining=false`
- `private_alpha_full_text_stored=false`
- `confirmSimulation_required=true`
- candidate count ≤ safe batch size
- each candidate has role, constellation, operator family, fields, expected information gain, expected failure addressed, and expression reference
- no `expression`, `regular`, or full alpha expression field is stored in Git

## Gate 2 — structural duplicate precheck

Run:

```bash
node scripts/wqb-duplicate-precheck.mjs \
  runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json
```

Output:

```text
runs/wqb-current/duplicate-precheck-BATCH-20260622-F-starmap-scout-and-climb.json
```

This compares structural signatures against historical batch records. It does not use full expression text.

Caveat: because long-term Git intentionally stores only `expression_ref`, this precheck cannot prove full FASTEXPR uniqueness. Immediately before dispatch, materialized expressions should also be hashed in a private/local cache that is not committed.

## Gate 3 — forum claim review

Before dispatch, the community baseline backfill should either be completed or explicitly waived with a reason.

Expected file after local run:

```text
runs/wqb-community/forum-claims-baseline.json
```

Then integrate:

```bash
node scripts/wqb-integrate-forum-claims.mjs
```

If this file is absent, Batch F remains shallow-prior and should not be treated as fully community-informed.

## Gate 4 — exact user authorization

Only after the above gates and user text:

```text
授权 Batch F，confirmSimulation=true
```

may the system materialize and dispatch Batch F.

No blanket authorization or “继续” is enough for WQB simulation.
