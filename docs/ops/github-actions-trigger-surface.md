# GitHub Actions Trigger Surface

Date: 2026-07-03

Purpose: document why `notion/export/**` canary branches should create zero GitHub Actions runs.

## Policy

`scripts/check-github-actions-policy.mjs` enforces:

- every `push` trigger must explicitly constrain `branches`, `branches-ignore`, `tags`, or `tags-ignore`;
- no `push` pattern may match `notion/export/**`;
- `branches-ignore` policies must explicitly exclude `notion/export/**`;
- `create`, `workflow_run`, and `repository_dispatch` triggers are rejected;
- workflow `on` declarations must use block style; inline/flow forms such as `on: push`, `on: [push, pull_request]`, `"on": push`, and `'on': [push]` are rejected;
- pure scheduled workflows do not receive meaningless `push` configuration.

## Current Workflows

| Workflow | Triggers | Private runner cost surface | `notion/export/**` result |
|---|---|---|---|
| `ci-linux.yml` | `pull_request` to `main`; `push` to `main`; manual dispatch | PRs, pushes to `main`, manual runs | Not matched by `push.branches: [main]` |
| `cockpit-ignition-wire.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |
| `publish-dispatch.yml` | manual dispatch | Manual only | No push trigger |
| `publish-dry-run.yml` | manual dispatch; path-limited pull request | Manual/PR only | No push trigger |
| `sync-douban.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |
| `sync-getnote.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |
| `sync-steam-quarterly.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |
| `weekly-governance-report.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |
| `worang-index.yml` | schedule; manual dispatch | Scheduled/manual only | No push trigger |

## Retired Workflows

### `link-index-dry-run.yml`

Retired on 2026-07-03.

The experiment successfully generated its candidate output, and the manually opened follow-up PR #185 was merged. The workflow no longer has an active operational role, so retaining `contents: write` and `pull-requests: write` would only preserve an unnecessary write surface. Historical evidence remains in Git history and PR #185.

## Canary Verification

The real canary must use global before/after run snapshots, not branch-filtered run lists:

```bash
gh run list --limit 100 --json databaseId,workflowName,event,headBranch,status,createdAt > before.json

# push notion/export/<export-id>

gh run list --limit 100 --json databaseId,workflowName,event,headBranch,status,createdAt > after.json
```

Acceptance is `after - before = 0` by `databaseId`.

Verified: 2026-07-02

- Export ID: `59e6f271-ff4a-5734-ae03-aa4334b736b3`
- Export branch: `notion/export/59e6f271-ff4a-5734-ae03-aa4334b736b3`
- Export commit: `b6b78edbc4d1d67b65150ef748d3f9ffc4f71b3b`
- Global Actions runs created: `0`
- Repeated run: `already_exported`
- Repeated resume: `already_exported`
