# WQB forum refresh workflow

## Goal

Maintain a safe, reusable WQB community knowledge layer without requiring Notion AI to log in to WQB or scrape the forum live.

Architecture:

```text
Local authenticated CNHKMCP
→ metadata-only forum index
→ local distillation into clusters/read queue/claim candidates
→ safety scan
→ GitHub artifact
→ Notion AI reads GitHub and updates scheduler/claims
```

## Why local first

Forum access currently requires WQB authentication. Keeping forum refresh local avoids storing WQB credentials or sessions in Cloudflare, GitHub Actions, or Notion. The committed artifacts contain only sanitized metadata and claim-level outputs.

Cloudflare, if used, should be a read-only facade over GitHub artifacts rather than the component that logs into WQB.

## Files

Scripts:

- `scripts/wqb-forum-update-local.py` — local authenticated metadata refresh via CNHKMCP.
- `scripts/wqb-forum-distill.mjs` — metadata-only clustering, read queue, claim candidates, digest.

Artifacts:

- `runs/wqb-community/forum-post-index-latest.json`
- `runs/wqb-community/forum-post-index-YYYYMMDDTHHMMSSZ.json`
- `runs/wqb-community/forum-post-index-local.json` for compatibility with earlier runs
- `runs/wqb-community/forum-title-clusters-latest.json`
- `runs/wqb-community/forum-title-clusters-YYYYMMDD.json`
- `runs/wqb-community/forum-read-queue.json`
- `runs/wqb-community/forum-claim-candidates-latest.json`
- `runs/wqb-community/forum-claim-candidates-YYYYMMDD.json`
- `docs/ops/wqb-community-digest-YYYYMMDD.md`

## Manual run

From the repository root:

```bash
PY=".venv-wqb-cnhkmcp/bin/python"
"$PY" scripts/wqb-forum-update-local.py
node scripts/wqb-forum-distill.mjs runs/wqb-community/forum-post-index-latest.json
rg -n -i 'cookie|token|password|session|authorization|credential|secret|bearer|1171964523|qq.com' \
  runs/wqb-community docs/ops/wqb-community-digest-*.md || true
git add runs/wqb-community docs/ops/wqb-community-digest-*.md
git commit -m "chore(wqb): refresh forum community artifacts"
git push
```

If running non-interactively, provide credentials only through the local environment:

```bash
WQB_EMAIL="..." WQB_PASSWORD="..." \
  .venv-wqb-cnhkmcp/bin/python scripts/wqb-forum-update-local.py --non-interactive
```

Do not put credentials in Git, command history snippets shared in chat, Cloudflare variables, or GitHub Actions unless a separate explicit security decision is made.

## Full-read policy

Default workflow is metadata-only.

When deeper information is needed, use `forum-read-queue.json` to pick a small number of posts and extract claim-level summaries only:

```json
{
  "post_id": "...",
  "title": "...",
  "claim": "...",
  "failure_modes": ["PROD_CORRELATION"],
  "design_effect": "...",
  "confidence": "full_read_claim",
  "do_not_copy_expression": true
}
```

Do not commit full private post bodies or private alpha expressions.

## Scheduling recommendation

Do not run this as high-frequency scraping. The useful cadence is:

- every 2–3 days,
- before designing a new simulation batch,
- or after major WQB community activity.

## Safety checklist

Before committing artifacts:

1. Confirm outputs are metadata-only or claim-only.
2. Run the sensitive-word scan.
3. Confirm no credentials, sessions, cookies, private alpha full text, or full post bodies are present.
4. Commit only sanitized artifacts.
5. Push to GitHub so Notion AI can read them without live WQB access.
