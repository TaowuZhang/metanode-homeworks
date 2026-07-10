# Workflow status and next actions — 2026-07-03

This snapshot records what each active or recently retired workflow has actually proved, what remains unproved, and the smallest next action. It is evidence inventory, not a permanent quality score.

## Summary

- Most workflows have a real successful run on current or recent `main`.
- The main systemic gap is generated PR assembly: content generators must refresh and include the generated Worang index in the same PR.
- Dry-run evidence is not treated as proof of real external publication.
- Stable workflows are not changed merely to create activity.

## Workflow-by-workflow status

| Workflow | Current evidence | Status | Next action |
|---|---|---|---|
| `ci-linux.yml` | Current Linux repository checks pass. A runtime deprecation warning remains advisory. | healthy | Identify the exact Action emitting the warning before changing versions; do not guess-upgrade. |
| `sync-steam-quarterly.yml` | Real quarterly candidate run completed; PR #362 merged; follow-up index repair PR #364 merged. | healthy and closed-loop | Observe the next scheduled quarter. Its index-refresh pattern is the reference for other generated PR workflows. |
| `sync-getnote.yml` | Real scheduled sync completed with 20 records and healthy evidence. | healthy | Make the rolling `recent/` semantics and Git-history recovery path explicit; do not turn it into an unbounded mirror by default. |
| `worang-index.yml` | Current-main maintenance run succeeded and produced no unnecessary change. | healthy | Keep it as the canonical generator/checker; generated PR workflows should call the same build/check pair before opening a PR. |
| `weekly-governance-report.yml` | Report generation succeeded, but PR #365 failed because the index was not refreshed in the same branch. | assembly gap | PR #367 adds index generation, verification, and index files to the generated governance PR. |
| `sync-douban.yml` | Real sync and cookie authentication succeeded; PR #366 contains the expected mirror boundary but fails the index check. | assembly gap | Apply the same index-refresh pattern as Steam. No data-source redesign is required. |
| `cockpit-ignition-wire.yml` | Completed and stable in the current workflow set. | healthy | No immediate code change. Revisit only when cockpit wiring or evidence contracts change. |
| `publish-dry-run.yml` | Manual run 28648125223 succeeded on current main and uploaded an artifact. | healthy dry-run | Validate the generated summary and declared artifact files, not only the job exit code. |
| `publish-dispatch.yml` | Manual `dry_run` dispatch 28648462725 succeeded; the existing gate prevented treating this as live evidence. | dry-run path proved | Keep real publication gated. Exercise supported modes through offline adapter/plan contracts and validate generated summaries. |
| `link-index-dry-run.yml` | Candidate generation completed historically; automatic PR creation failed, then manual PR #185 was merged. | experiment complete | Retire the Actions entry point and its write permissions; PR #368 proposes the retirement. |

## Generated PR closure contract

Any workflow that writes repository content and opens a PR should follow this order:

```text
generate owned content
→ audit owned-path boundaries
→ detect whether owned content changed
→ node scripts/build-worang-index.mjs
→ node scripts/check-worang-index.mjs
→ include content and docs/ops/worang-index/current.{json,md} in one PR
```

The content-change decision must be made before index generation, because the generated index itself will otherwise make every run appear changed.

The PR must include only:

1. the workflow-owned output boundary;
2. `docs/ops/worang-index/current.json`;
3. `docs/ops/worang-index/current.md`.

A red index check after correct content generation is an assembly failure, not evidence that the source sync itself failed.

## Current PR map

- PR #367: close the weekly governance report/index assembly gap.
- PR #368: retire the completed link-index workflow and remove its active write surface.
- Publish summary-contract changes are prepared on branch `test/publish-run-contracts-20260703`; opening a PR was blocked by the connected GitHub safety layer because the branch touches publication-control code. The branch is intentionally not merged or force-moved.
- `sync-douban.yml` remains unchanged: the connected GitHub safety layer refused to rewrite a workflow containing an existing repository-secret reference. This is a tooling boundary, not a judgment that the fix is unnecessary.

## Non-actions

- Do not rerun an obsolete workflow merely to replace a historical red run with green.
- Do not enable real publication to test a workflow audit.
- Do not modify stable cockpit or Steam workflows without a new counterexample.
- Do not interpret GitHub Action startup or permission failures as source-generator failures.
