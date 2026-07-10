# WQB premium source full discovery runbook

## Purpose

The user corrected the source policy: checking only the current 50 forum metadata rows is not enough. We need a full discovery pass for high-vote premium community posts.

Rule:

```text
Any WQB community post with vote_count >= 50 must be read or explicitly blocked/deferred before claiming the source layer is adequate.
```

## What happened in the first attempt

The first local full-discovery attempt returned:

```text
topics_discovered = 0
unique_posts_discovered = 0
topic_results[0].error = support authentication exchange failed with status 403 / Cloudflare "Just a moment"
```

Interpretation: this is not evidence that no high-vote posts exist. The list route was blocked by a support-auth / Cloudflare challenge. The script now records this and can run search fallback.

## Discovery command

```bash
cd "/Users/jiajia/Library/Mobile Documents/com~apple~CloudDocs/Obsidian/kb/worang-main"

git pull --rebase origin feat/wqb-alpha-production-scaffold-20260621

.venv-wqb-cnhkmcp/bin/python -m py_compile scripts/wqb-forum-discover-premium-local.py

.venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-discover-premium-local.py \
  --topic-max-results 200 \
  --posts-per-topic 300 \
  --vote-threshold 50 \
  --high-value-threshold 20 \
  --search-fallback always \
  --search-max-results 100
```

Outputs:

```text
runs/wqb-community/premium-source-full-discovery-latest.json
runs/wqb-community/premium-source-full-discovery-YYYYMMDDTHHMMSSZ.json
runs/wqb-community/premium-source-full-discovery-queue.json
```

Look at:

```text
list_route_blocked
search_fallback_ran
unique_posts_discovered
must_read_count
high_value_count
topic_results[].error
search_result_summary[].post_count
```

## Claim extraction for discovered high-vote posts

Only run this if `must_read_count > 0` or `high_value_count > 0`.

```bash
.venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-claim-extract-local.py \
  --queue runs/wqb-community/premium-source-full-discovery-queue.json \
  --out runs/wqb-community/premium-source-full-discovery-claims.json \
  --limit 100
```

## Integrate claims

```bash
node scripts/wqb-integrate-forum-claims.mjs runs/wqb-community/premium-source-full-discovery-claims.json
```

## Safety scan

Use the narrower real-secret scan, not a broad schema-word scan:

```bash
rg -n -i '1171964523|qq\.com|bearer [A-Za-z0-9._-]+|set-cookie|sessionid|WQB_PASSWORD|password=' \
  runs/wqb-community/premium-source-full-discovery-latest.json \
  runs/wqb-community/premium-source-full-discovery-queue.json \
  runs/wqb-community/premium-source-full-discovery-claims.json \
  runs/wqb-current/source-claim-registry.json \
  runs/wqb-current/starmap-weight-model.json \
  runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json \
  runs/wqb-current/forum-claims-integration-summary.json || true
```

If anything real appears, do not commit.

## Commit

```bash
git add \
  runs/wqb-community/premium-source-full-discovery-latest.json \
  runs/wqb-community/premium-source-full-discovery-*.json \
  runs/wqb-community/premium-source-full-discovery-queue.json \
  runs/wqb-community/premium-source-full-discovery-claims.json \
  runs/wqb-current/source-claim-registry.json \
  runs/wqb-current/starmap-weight-model.json \
  runs/wqb-current/authorization-ready-BATCH-20260622-F-starmap-scout-and-climb.json \
  runs/wqb-current/forum-claims-integration-summary.json

git commit -m "chore(wqb): add full premium forum discovery claims"
git push
```

If discovery still returns zero because list and search routes are blocked, commit only the discovery artifact if it accurately records the platform boundary; do not claim full source coverage.

## No simulation

This runbook does not authorize simulation. It is source discovery and claim extraction only.
