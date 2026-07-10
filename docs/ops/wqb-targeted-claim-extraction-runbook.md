# WQB Targeted Claim Extraction Runbook

This runbook executes the local CNHKMCP claim-only extraction request created at:

```text
runs/wqb-current/local-targeted-extractor-request.json
```

## Boundary

Do not run simulation. Do not submit. Do not generate a Batch F ready packet. Do not store full forum bodies, credentials, cookies, sessions, or expression-like details.

## Local command

```bash
cd "/Users/jiajia/Library/Mobile Documents/com~apple~CloudDocs/Obsidian/kb/worang-main"

git pull --rebase origin feat/wqb-alpha-production-scaffold-20260621

.venv-wqb-cnhkmcp/bin/python scripts/wqb-targeted-claim-extract-local.py \
  --request runs/wqb-current/local-targeted-extractor-request.json \
  --out runs/wqb-current/local-targeted-extractor-output.json \
  --limit 12

node scripts/wqb-integrate-targeted-claims.mjs \
  runs/wqb-current/local-targeted-extractor-output.json
```

## Safety scan

Run this before committing local outputs:

```bash
rg -n -i '1171964523|qq\.com|bearer [A-Za-z0-9._-]+|set-cookie|sessionid|WQB_PASSWORD|password=|full_expression|fast_expression|raw_template|formula' \
  runs/wqb-current/local-targeted-extractor-output.json \
  runs/wqb-current/p0a-repaired-claims-integrated.json \
  docs/ops/wqb-p0a-repaired-claims-integration-20260623.md
```

Expected result: no secret values and no expression-like content. Schema-level mentions of banned field names in validation metadata are allowed only when they do not carry values.

## Commit local outputs

```bash
git status --short

git add \
  runs/wqb-current/local-targeted-extractor-output.json \
  runs/wqb-current/p0a-repaired-claims-integrated.json \
  docs/ops/wqb-p0a-repaired-claims-integration-20260623.md \
  runs/wqb-current/run-state.json

git commit -m "chore(wqb): integrate targeted forum claim extraction"
git push
```

## What counts as success

The final report should include:

```text
attempted_count
read_success_count
improved_count
still_generic_count
metadata_only_count
blocked_count
which SCR upgraded to medium/strong
which SCR remain not_scheduler_eligible
whether expression-like content exists: must be no
Batch F slot support changes
simulation=false
submit=false
Batch F ready=false
```

## Interpretation

If no source improves after local extraction, keep Batch F not-ready. Do not force weak sources into scheduler slots.

If some sources improve, they may only become slot support candidates. They still do not authorize simulation. Field refs, duplicate/hash prechecks, validator, and future explicit user authorization remain required.