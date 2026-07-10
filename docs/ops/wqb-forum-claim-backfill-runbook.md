# WQB forum claim backfill runbook

## Purpose

Convert premium/baseline community read queues into claim-level artifacts that the scheduler can use.

This is the second half of the community workflow:

```text
metadata index
→ premium/baseline read queue
→ selected local forum/help-center reads
→ claim-only summaries
→ source-claim-registry / star-map / Batch F update
```

## Safety rule

The output must not contain:

- WQB email/password
- cookies, sessions, tokens, auth headers
- full forum post bodies
- full comment bodies
- private Alpha full expressions
- copied Super Alpha combo expressions

The extractor reads full post/help-center content in local memory only, then writes short paraphrased claim drafts and keyword evidence counts.

## Premium-source run

Use this first. It supports the user's rule: posts with `vote_count >= 50` must be read.

```bash
.venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-claim-extract-local.py \
  --queue runs/wqb-community/premium-source-must-read-queue.json \
  --out runs/wqb-community/premium-source-claims.json \
  --limit 12
```

For non-interactive local runs:

```bash
WQB_EMAIL="..." WQB_PASSWORD="..." \
  .venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-claim-extract-local.py \
  --queue runs/wqb-community/premium-source-must-read-queue.json \
  --out runs/wqb-community/premium-source-claims.json \
  --non-interactive \
  --limit 12
```

To read only specific P0 sources:

```bash
.venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-claim-extract-local.py \
  --queue runs/wqb-community/premium-source-must-read-queue.json \
  --out runs/wqb-community/premium-source-claims.json \
  --source-ids forum:15238206176279,forum:41254129745815,official-community:wqb-community-reduce-correlation,official-community:wqb-community-avoid-overfitting \
  --limit 10
```

## Baseline run

After premium sources:

```bash
.venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-claim-extract-local.py --limit 12
```

Expected output:

```text
runs/wqb-community/forum-claims-baseline.json
```

## Safety scan

```bash
rg -n -i 'cookie|token|password|session|authorization|credential|secret|bearer|1171964523|qq.com' \
  runs/wqb-community/premium-source-claims.json \
  runs/wqb-community/forum-claims-baseline.json || true
```

If anything meaningful appears, do not commit; inspect and sanitize locally.

## Integrate premium claims

```bash
node scripts/wqb-integrate-forum-claims.mjs runs/wqb-community/premium-source-claims.json
```

Expected updates:

```text
runs/wqb-current/source-claim-registry.json
runs/wqb-current/starmap-weight-model.json
runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json
runs/wqb-current/forum-claims-integration-summary.json
```

## Commit

```bash
git add \
  runs/wqb-community/premium-source-claims.json \
  runs/wqb-community/forum-claims-baseline.json \
  runs/wqb-current/source-claim-registry.json \
  runs/wqb-current/starmap-weight-model.json \
  runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json \
  runs/wqb-current/forum-claims-integration-summary.json

git commit -m "chore(wqb): integrate premium forum claims"
git push
```

## Interpretation

The deterministic extractor produces claim drafts, not final truth. Treat `full_read_claim_draft_needs_review` as useful but still reviewable. Strong scheduler changes should be made only after checking whether the claim is consistent with WQB official docs and current gate failures.

No simulation is authorized by this runbook.
