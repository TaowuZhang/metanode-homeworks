# Branch archive index｜2026-06-01

Purpose: reduce remote branch congestion without losing evidence.

This index is a pre-delete ledger. It does not delete branches by itself. It records where evidence already lives, which branches still need review, and which branches should only be treated as legacy or spike material.

## Rules

- Branches are not archives.
- Do not batch-delete remote branches.
- Delete only after evidence is either in `main`, in a PR, in this index, in `docs/legacy/`, or in a tag.
- `wasteland/spike*` branches are not merged directly into `main`; valuable evidence should be indexed or tagged first.
- `terminal-bridge` branches are legacy only; do not restore terminal bridge as current architecture.
- CSV pair / `*_all.csv` layered branches are historical after canonical CSV migration; do not reintroduce paired runtime assumptions.

## Current congestion sources

| Source | Shape | Handling |
|---|---|---|
| Merged PR head branches | Many `chore/*`, `content/*`, `audit/*`, `tooling/*`, `feature/*` | Delete after `git merge-base --is-ancestor origin/<branch> origin/main` or PR evidence check |
| `audit/*-2026-05-31` | Mostly merged audit reports | Keep reports in `docs/ops/`; delete merged heads; review `audit/fuhai-deep-2026-05-31` separately |
| `wasteland/spike*` | Large Rust CLI / digest experiment forest | Build spike index or archive tags, then delete branches |
| terminal bridge branches | Old Notion/local terminal bridge path | Only legacy archive; never current path |
| CSV layered branches | `docs/csv-layer-policy-98`, `fix/sense-emotion-layered-csv-80` | Historical / replaced by canonical CSV migration |
| Open PR heads | #100, #103, #104 | Review/merge/split before branch deletion |

## Open PR branches

| Branch | PR | Current recommendation | Evidence / note |
|---|---:|---|---|
| `docs/nlm-runtime-boundary` | #104 | Merge first after light review | Adds `docs/ops/nlm-runtime-boundary.md`; matches current no-terminal-bridge, GitHub MCP first, VPS-last posture |
| `release/wo-v0.1-space-snapshot-100` | #100 | Update then merge | Snapshot is valuable, but PR state references should be updated after #102 / closed #97/#99 |
| `distribution-safety-gate-2026-06-01` | #103 | Split / rebase before merge | Contains distribution safety gate plus broader skill-governance material; base is `refactor/prioritize-knowledge-dirs` |

## Merged PR branches safe to delete after local verification

Run before deleting each branch:

```bash
git merge-base --is-ancestor origin/<branch> origin/main \
  && echo "MERGED: <branch>" \
  || echo "CHECK DIFF: <branch>"
```

| Branch | PR / evidence | Evidence location | Deletion note |
|---|---:|---|---|
| `content/lighthouse-expansion-2026-05-07` | #63 merged | main history / `成为/灯塔*`, `成为/色散*` | delete-if-merged |
| `content/becoming-updates-2026-05-07` | #64 merged | main history / `成为/*` | delete-if-merged |
| `feature/excrete-content-ref-resolution` | #65 merged | main history / excrete digest code | delete-if-merged |
| `feature/touch-nlm-manifest-routing` | #66 merged | main history / `wo/nlm/manifest.json`, touch route | delete-if-merged; continue NLM work from `main` |
| `tooling/notion-terminal-bridge` | #67 merged | git history; later legacy docs | delete-if-merged; do not revive |
| `tooling/notion-markdown-bridge` | #68 merged | main history / Notion markdown bridge script | delete-if-merged |
| `chore/sediment-quadrants` | #70 merged | main history / `成为/沉积/` | delete-if-merged |
| `cleanup/notion-terminal-bridge-legacy` | #71 merged | `docs/legacy/terminal-bridge-ops.md`, `docs/legacy/notion-local-github-architecture-v1.md` | delete-if-merged |
| `chore/ai-dispute-ledger-routes` | #72 merged | main history / AI dispute routing docs | delete-if-merged |
| `chore/direction-trace-wo-recall` | #73 merged | main history / `wo/field-notes/` | delete-if-merged |
| `chore/personal-charter-origin` | #74 merged | main history / `成为/项目/建·个人宪章.md` | delete-if-merged |
| `chore/weight-light-site` | #75 merged | main history / `成为/项目/研·重与轻的系统代偿.md` | delete-if-merged |
| `chore/book-guessing-intent-field` | #76 merged | main history / guess-book project docs | delete-if-merged |
| `chore/greenhouse-time-handle` | #77 merged | main history / greenhouse time handle docs | delete-if-merged |
| `chore/seedlings-routing` | #78 merged | main history / seedling routing docs | delete-if-merged |
| `chore/fuhai-axiom-footprints` | #83 merged | main history / `docs/ideas/*` | delete-if-merged |
| `wo/grade-scale` | #84 merged | main history / `wo/grade.md` | delete-if-merged |
| `content/ai-six-symbiotic-actions` | #86 merged | main history / `wo/ai.md` | delete-if-merged; do not return to old four-position model |
| `audit/zhanqiao-docsops-2026-05-31` | #88 merged | `docs/ops/zhanqiao-docsops-audit-2026-05-31.md` | delete-if-merged |
| `audit/fuhai-seven-2026-05-31` | #89 merged | `docs/ops/fuhai-seven-audit-2026-05-31.md` | delete-if-merged |
| `audit/caoman-2026-05-31` | #90 merged | `docs/ops/caoman-audit-2026-05-31.md` | delete-if-merged |
| `audit/jiegang-resources-2026-05-31` | #91 merged | `docs/ops/jiegang-resources-audit-2026-05-31.md` | delete-if-merged |
| `audit/yong-chaoci-2026-05-31` | #92 merged | `docs/ops/yong-chaoci-audit-2026-05-31.md` | delete-if-merged |
| `audit/local-notion-position-2026-05-31` | #93 merged | `docs/ops/local-notion-position-audit-2026-05-31.md` | delete-if-merged |
| `audit/ai-position-map-2026-05-31` | #94 merged | `docs/ops/ai-position-map-audit-2026-05-31.md` | delete-if-merged |
| `audit/becoming-projects-fuhai-manifestation-2026-05-31` | #95 merged | `docs/ops/becoming-projects-fuhai-manifestation-audit-2026-05-31.md` | delete-if-merged |
| `audit/gongyueshu-l1-l2-runtime-2026-05-31` | #96 merged | `docs/ops/gongyueshu-l1-l2-runtime-audit-2026-05-31.md` | delete-if-merged |
| `pr/sense-emotion-library` | #62 merged | main history / digest event wiring | delete-if-merged |
| `chore/repo-governance-pr100` | branch tip equals old main audit merge | main history / PR #96 merge commit | delete-if-merged or delete if no independent diff |

## Branches needing review before deletion

| Branch | Why not auto-delete | Next check |
|---|---|---|
| `audit/fuhai-deep-2026-05-31` | PR #87 was closed unmerged draft | Decide whether to extract summary into audit index, then delete branch |
| `feature/getnote-mirror-v0` | No PR evidence in current audit | Run `git diff --name-status origin/main...origin/feature/getnote-mirror-v0` |
| `deprecation/terminal-bridge-source-map` | Legacy terminal evidence may be unique | Extract to `docs/legacy/terminal-bridge-source-map.md` or record here, then delete |
| `refactor/prioritize-knowledge-dirs` | #103 currently uses it as base; contains canonical CSV equivalent work | Keep until #103 is split/rebased/closed; then delete if no unique diff |
| `docs/csv-layer-policy-98` | PR #99 closed unmerged; policy may contain unique wording | Extract useful policy into CSV migration note if needed; do not restore paired `_all.csv` route |
| `fix/sense-emotion-layered-csv-80` | PR #97 closed unmerged; layered CSV implementation is obsolete after #102 | Extract test idea if needed; do not merge code |

## Spike branches

Default handling for all `wasteland/spike*`:

1. Do not merge directly into `main`.
2. Check whether branch tip is already included in main or represented by a merged PR.
3. If not included but historically useful, add a one-line entry below or create an archive tag.
4. Delete the remote branch only after evidence is indexed or tagged.

| Branch group | Topic | Keep evidence where | Delete after |
|---|---|---|---|
| `wasteland/spike-0-*` to `wasteland/spike-2*` | early dependency / raw residue / decay / palm excavation | this index or future `docs/ops/spike-index-2026-06-01.md` | local diff reviewed |
| `wasteland/spike-3a*` to `wasteland/spike-3f*` | digest types, node mapping, event/abort snapshot, trace entry | spike index | local diff reviewed |
| `wasteland/spike-3-5*` | command entry series | spike index | confirmed absorbed or indexed |
| `wasteland/spike-3-6*` | shadow event series | spike index | confirmed absorbed or indexed |
| `wasteland/spike-3-7*` | real emit series | main PR history where merged; spike index for remainder | `--merged` or PR evidence checked |

Known merged evidence from previous PR history:

| Branch | PR / evidence | Note |
|---|---:|---|
| `wasteland/spike-3-7j-sense-notice-real-emit` | #57 merged | delete-if-merged |
| `wasteland/spike-3-7k-sense-inhale-real-emit` | #58 merged | delete-if-merged |
| `wasteland/spike-3-7l-sense-recognize-real-emit` | #59 merged | delete-if-merged |
| `wasteland/spike-3-7m-sense-exhale-real-emit` | #60 merged | delete-if-merged |
| `wasteland/spike-3-7n-taste-real-emit` | #61 merged | delete-if-merged |

## Terminal bridge legacy

| Branch | Legacy evidence | Current posture |
|---|---|---|
| `tooling/notion-terminal-bridge` | original watcher in git history / PR #67 | historical only |
| `cleanup/notion-terminal-bridge-legacy` | `docs/legacy/terminal-bridge-ops.md`, `docs/legacy/notion-local-github-architecture-v1.md` / PR #71 | delete branch after merge verification |
| `deprecation/terminal-bridge-source-map` | unknown, needs diff | extract to legacy docs if unique; never current path |

## CSV / NLM / touch decisions

| Branch | Current posture | Evidence |
|---|---|---|
| `feature/touch-nlm-manifest-routing` | already relevant but merged; delete branch after verification | PR #66 / main |
| `docs/nlm-runtime-boundary` | active PR, merge first | PR #104 |
| `docs/csv-layer-policy-98` | obsolete / historical after canonical CSV migration | PR #99 closed, PR #102 merged canonical CSV |
| `fix/sense-emotion-layered-csv-80` | obsolete implementation after canonical CSV migration | PR #97 closed, PR #102 merged canonical CSV |

## Safe deletion commands

Never batch delete. For each branch:

```bash
BR=content/lighthouse-expansion-2026-05-07

git fetch origin --prune

git merge-base --is-ancestor "origin/$BR" origin/main \
  && git push origin --delete "$BR" \
  || echo "Do not delete yet: $BR has non-ancestor tip"
```

For squash/cherry-pick/equivalent branches, record the evidence above first, then delete manually:

```bash
git diff --name-status origin/main...origin/$BR
git push origin --delete "$BR"
```

## Commands not to run

```bash
# Do not batch-delete merged remote branches.
git branch -r --merged origin/main | xargs -n 1 git push origin --delete

# Do not delete all spike branches by pattern.
git push origin --delete $(git branch -r | grep 'wasteland/spike')

# Do not delete terminal branches by pattern.
git push origin --delete $(git branch -r | grep terminal)

# Do not delete unknown feature branches without diff.
git push origin --delete feature/getnote-mirror-v0
```

## Next suggested order

1. Merge #104 after light review.
2. Update #100 snapshot facts, then merge if still desired.
3. Split or rebase #103.
4. Delete verified merged PR heads one by one.
5. Create spike index / archive tags for `wasteland/spike*`, then delete indexed branches.
6. Review and box terminal / CSV legacy branches.
