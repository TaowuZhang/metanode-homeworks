# GitHub Repository Inventory · 2026-07-07

## Goal

`dongxi-heji/worang` is the single long-term versioned source of truth. Other repositories owned by TaowuZhang are inventory targets and should not become parallel canonical stores.

## Authenticated account

- login: `TaowuZhang`
- id: `80830286`

## App installation scope

The installed GitHub App account list visible here contains only:

- `dongxi-heji`

## Owner-scope result

Authenticated GitHub API query `affiliation=owner` returned three repositories on page offset 0 and no repositories on page offset 100. Within this access scope, the complete owned-repository inventory is:

| repo | owner | permission | visibility | archived | default branch | size | readme status | preliminary route |
|---|---|---:|---|---|---|---:|---|---|
| `TaowuZhang/CN-Mirror-CLI` | `TaowuZhang` | admin | public | false | main | 50 | README found | CN mirror / Xget-adjacent material; durable operating notes sedimented into `worang`; local Codex to verify remaining assets before archive |
| `TaowuZhang/Xget` | `TaowuZhang` | admin | public | false | main | 995 | README found | developer-resource acceleration engine; repository-position notes sedimented into `worang`; local Codex to check fork/upstream/release/pages before archive |
| `TaowuZhang/neetcode-submissions` | `TaowuZhang` | admin | public | false | main | 75 | README not found | coding-practice residue; inspect root tree locally or in GitHub UI before final route |

## Keep target

| repo | owner | permission | visibility | archived | default branch | size |
|---|---|---:|---|---|---|---:|
| `dongxi-heji/worang` | `dongxi-heji` | admin | private | false | main | 18156 |

## Organization-member repositories

Authenticated GitHub API query `affiliation=organization_member` returned `dongxi-heji/worang` plus many `MetaNodeAcademy/*` repositories. All observed `MetaNodeAcademy/*` entries had pull-only permission in this session. They are not controlled repository-maintenance targets here.

## Codex local supplement

Local Codex reported WQB / Alpha assets under local paths including `Alpha-Template-Automation`, `Easy WQB`, `Alpha`, `AIACv2_v1`, and `wqb-cnhkmcp-cloudflare-worker`. Repository search in this session did not find these names as `TaowuZhang` GitHub repositories. They are treated as local repositories or directories until local Git remotes prove otherwise.

Their WQB-specific migration contract is now recorded in:

- `docs/ops/wqb-alpha-consolidation-architecture-20260707.md`

## Local Codex execution task

The user confirmed that the listed local repositories/assets are theirs and asked that non-WQB cleanup be assigned to local Codex.

Task package:

- `docs/ops/local-codex-github-cleanup-task-20260707.md`

Scope:

- protect WQB / quant Alpha related repositories from generic cleanup;
- process all other repositories through inventory, migration, archive, or delete-candidate routes;
- write execution evidence into `runs/github-cleanup/2026-07-07/`;
- update this inventory after local execution.

## Visible migration artifacts

The following safe, non-secret sediments were created in `worang` from currently visible non-WQB repositories:

| source repo | sediment file | status |
|---|---|---|
| `TaowuZhang/CN-Mirror-CLI` | `docs/ops/cn-mirror-cli-sediment-20260707.md` | durable operating notes migrated; full implementation not migrated |
| `TaowuZhang/Xget` | `docs/ops/xget-repository-sediment-20260707.md` | repository-position notes migrated; full product implementation not migrated |

## Immediate routing

### `TaowuZhang/CN-Mirror-CLI`

README identifies it as a Bash-based GitHub acceleration / mirror switching tool with scheduled mirror refresh, local refresh script, safe publish script, and Xget prefix integration. Its durable operating notes are now preserved in:

- `docs/ops/cn-mirror-cli-sediment-20260707.md`

Remaining route: local Codex verifies root tree, workflow/release/pages/package status, unique assets, and local dirty state before archive.

### `TaowuZhang/Xget`

README identifies it as a developer-resource acceleration engine. `package.json` shows a deployable/testable TypeScript/Cloudflare Worker-style project surface. Its repository-position notes are now preserved in:

- `docs/ops/xget-repository-sediment-20260707.md`

Remaining route: local Codex checks fork/upstream/release/pages/external link status before deciding archive vs leave.

### `TaowuZhang/neetcode-submissions`

No README was found through the connector. Treat it as an unknown coding-practice repository until root files are inspected.

## External repository freeze attempt

A direct attempt was made to write a public status marker into `TaowuZhang/CN-Mirror-CLI`; the write was blocked by the safety layer. No external repository content was changed by that attempt.

Because repository-level archive/delete and external public status mutation are not available safely from this session, the authoritative cleanup state remains here in `worang`.

## Current executable cleanup state

| repo | state | next safe action |
|---|---|---|
| `TaowuZhang/CN-Mirror-CLI` | inventoried; visible durable notes migrated; external marker write blocked; delegated to local Codex | local Codex verifies remaining assets, then archive if gates pass |
| `TaowuZhang/Xget` | inventoried; visible repository-position notes migrated; delegated to local Codex | local Codex checks fork/upstream/release/pages/external surface, then archive or leave |
| `TaowuZhang/neetcode-submissions` | inventoried; no README found | inspect root tree locally; if no durable content, archive/delete candidate |

## Next inventory fields for each non-worang owned repo

```yaml
repo_full_name:
readme_status:
root_tree_status:
unique_assets:
external_links:
workflow_or_pages_surface:
release_or_package_surface:
fork_or_upstream_status:
content_to_sediment_into_worang:
recommended_route:
manual_gate:
```

## Boundary

The original inventory above was the base for one-by-one review. The execution section below is the authoritative local result; detailed action evidence remains in `runs/github-cleanup/2026-07-07/final-actions.md`.

## Local Codex execution result

Execution evidence: `runs/github-cleanup/2026-07-07/`.

### Archived repos

- `TaowuZhang/CN-Mirror-CLI`: archived through GitHub API at 2026-07-07T22:18:37+08:00; immediate API read-back confirmed `archived=true`.
- Already archived before this run and left unchanged: `TaowuZhang/Sync-Notion-Obsidian`, `TaowuZhang/Gemini-Balance-Lite`.

### Deleted repos

- None.

### Left untouched repos

- `TaowuZhang/Xget`: fork/upstream, external homepage, and deployment workflow dependencies remain.
- `TaowuZhang/neetcode-submissions`: repository was pushed on 2026-07-06, so an active NeetCode integration must be ruled out.
- `TaowuZhang/Sync-Notion-Obsidian`: remote already archived; local checkout has an uncommitted workflow change.
- `TaowuZhang/OpenBB`: remote returns 404; local unique application source remains.
- `TaowuZhang/Gemini-Balance-Lite`: remote already archived; external Vercel homepage remains.
- `TaowuZhang/Network-Flow-Forge`: unique product and architecture corpus has not been migrated.
- `TaowuZhang/simplified-chinese-editing-proofreading-tool`: unique Tauri/Svelte/Rust implementation has not been migrated.
- `/Users/jiajia/Documents/GitHub/worang-bibigpt-migration`: clean duplicate checkout of canonical `worang`; no local deletion was performed.

### WQB protected repos confirmed

- `/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/`: Git repository; 679 dirty entries recorded; no mutation performed.
- `/Users/jiajia/Documents/GitHub/Easy WQB/`: protected local non-Git directory; no mutation performed.
- `/Users/jiajia/Documents/GitHub/Alpha/`: protected local non-Git directory; no mutation performed.
- `/Users/jiajia/Documents/GitHub/AIACv2_v1/`: protected local non-Git directory; no mutation performed.
- `/Users/jiajia/wqb-cnhkmcp-cloudflare-worker/`: protected local non-Git directory; no mutation performed.
- `/Users/jiajia/worang/`: canonical repository; updated only with this evidence and inventory result.
- Additional owner repositories `TaowuZhang/Alpha-Miner-Automation`, `TaowuZhang/Alpha-Template-Automation`, and `TaowuZhang/World-Quant-Brain-MCP` were treated as WQB inventory-only objects.

### Migrated files

- `docs/ops/cn-mirror-cli-sediment-20260707.md` (pre-existing migration verified in this run).
- `docs/ops/xget-repository-sediment-20260707.md` (pre-existing migration verified in this run).
- `docs/ops/openbb-local-sediment-20260707.md` (Round 3 safe architecture and exclusion summary; source not copied).
- `docs/ops/network-flow-forge-sediment-20260707.md` (Round 3 BFF contract, submodule, and asset-boundary summary).
- `docs/ops/simplified-chinese-editing-proofreading-tool-sediment-20260707.md` (Round 3 incomplete product topology linked to canonical `技能/编校.md`).

### Unresolved gates

- `Xget`: its homepage returned HTTP 200; 135 workflow runs include successful Cloudflare Pages/Workers deployments; its fork is 18 commits ahead and 464 behind upstream with six changed files; local dirty state remains unavailable.
- `neetcode-submissions`: ten recent commits through 2026-07-06 confirm continued use; local dirty state remains unavailable.
- `Sync-Notion-Obsidian`: preserve or explicitly discard the local workflow modification.
- `OpenBB`: resolve the missing remote and migrate or discard unique app source.
- `Gemini-Balance-Lite`: confirm the Vercel deployment is retired before considering deletion.
- `Network-Flow-Forge`: migrate durable architecture and resolve its submodule/source ownership.
- `simplified-chinese-editing-proofreading-tool`: decide what durable code or knowledge should migrate.
- Duplicate local `worang-bibigpt-migration`: local checkout removal remains unperformed and was not authorized by this no-delete task.

### API-verified online state · round 2

- `TaowuZhang/CN-Mirror-CLI`: `archived=true`; completed record reconciled, archive action not repeated.
- `TaowuZhang/Xget`: `archived=false`; live homepage returned HTTP 200; leave unchanged.
- `TaowuZhang/neetcode-submissions`: `archived=false`; latest commit 2026-07-06T06:33:13Z; leave unchanged.
- `TaowuZhang/Sync-Notion-Obsidian`: `archived=true`.
- `TaowuZhang/Gemini-Balance-Lite`: `archived=true`; configured Vercel homepage returned HTTP 200.
- `TaowuZhang/Network-Flow-Forge`: `archived=false`.
- `TaowuZhang/simplified-chinese-editing-proofreading-tool`: `archived=false`.
- `TaowuZhang/OpenBB`: API lookup remains HTTP 404.

### Next manual gates · round 2

- Xget: retire the live homepage/deployments, decide the 18 fork-only commits, and establish local dirty/unpushed state.
- neetcode: retire the active submission integration and establish local dirty/unpushed state; then prefer archive over delete.
- Remaining local/product gates are detailed in `runs/github-cleanup/2026-07-07/unresolved-gates-round-2.md`.

### Round 3 execution result

- No GitHub repository was archived or deleted in Round 3.
- Xget and neetcode remain untouched because deployment/continued-write evidence and unavailable local dirty state fail archive gates.
- Sync's dirty workflow was classified as a cron replacement and left intact for human cadence confirmation.
- OpenBB, Network-Flow-Forge, and the Chinese proofreading tool received minimal safe sediment pages; full source/vendor/data was not copied.
- The duplicate `worang-bibigpt-migration` checkout was reconfirmed clean and ancestral; no local deletion occurred.
- WQB protected assets remained inventory-only and unchanged.
- Round 3 evidence and remaining gates: `runs/github-cleanup/2026-07-07/unresolved-gates-round-3.md`.

### Round 4 execution result

- No repository, deployment, workflow, or local checkout was mutated.
- OpenBB's eight personal source files are now represented by an exact behavior/history manifest; exact source remains pending a promotion-versus-historical-only choice.
- Network-Flow-Forge now has a canonical core-spec and overlay-purpose index; the full corpus remains archive evidence.
- The Chinese proofreading prototype's unwired state is documented in `docs/ops`; `技能/编校.md` remains the canonical process and was not changed.
- Xget and Gemini now have explicit deployment-retirement checklists; both external endpoints returned HTTP 200.
- neetcode has zero GitHub Actions workflows; continued commits come from an authenticated external/manual path that must be disabled before archive.
- Sync's local choice is now explicit: hourly cron versus every 12 hours; the patch remains untouched.
- Duplicate `worang` checkout remains local-only and untouched; WQB protected assets remain inventory-only.
- Round 4 ledger: `runs/github-cleanup/2026-07-07/unresolved-gates-round-4.md`.

## Round 5 cleanup gate status

Round 5 converted remaining blockers into either Codex-actionable preservation policies or explicit user decisions. No external or local destructive action was taken.

### Codex-actionable

- OpenBB source-file retention policy: complete at per-file summary/index level.
- Network-Flow-Forge asset preservation policy: complete at spec/diagram/overlay/submodule index level.
- Chinese editing/proofreading tool finish-or-historicize gate: technical state, target layer, and recommended default documented.
- Duplicate `worang` checkout: recorded as user-confirmed local deletion only; never a remote repository action.

### User decisions required

- `D-openbb-source-retention`: keep, promote, or explicitly discard exact OpenBB source.
- `D-network-preservation`: archive-only preservation or promote selected Network specs/overlays.
- `D-chinese-proofreading-tool`: finish or historicize the unfinished tool.
- `D-xget-retire`: Xget deployment retirement.
- `D-neetcode-stop-writer`: neetcode writer shutdown.
- `D-sync-cron`: Sync cron cadence.
- `D-gemini-vercel`: Gemini Vercel retirement.
- `D-duplicate-worang-checkout`: duplicate checkout deletion authorization.

Decision details and safe defaults: `runs/github-cleanup/2026-07-07/unresolved-gates-round-5.md`.
