# Branch archive continuation index｜2026-06-02

Purpose: continue remote branch cleanup after `branch-archive-2026-06-01.md` without losing evidence.

This index records the remaining non-spike branches after the first cleanup batch. It is a pre-delete / pre-merge ledger, not a deletion command.

## Rules carried forward

- Branches are not archives.
- Do not batch-delete remote branches.
- Delete only after evidence is in `main`, in a PR, in this index, in `docs/legacy/`, or in a tag.
- Do not delete by namespace or pattern.
- `wasteland/spike*` remains out of scope until a spike index exists.
- `docs/getnote-revoke-rotate-path` is hold-active because it may be active GetNote / 得到大脑 Get 笔记 work with another AI.
- Terminal bridge branches are legacy only; do not restore terminal bridge as current architecture.
- CSV pair / `*_all.csv` layered branches are historical after canonical CSV migration; do not reintroduce paired runtime assumptions.

## Current non-spike review bucket

| Branch | Head | Current status | Next action | Delete condition |
|---|---|---|---|---|
| `audit/fuhai-deep-2026-05-31` | `82506117` | archive-to-doc | Preserve the #87 audit report path / summary before deletion | Delete only after `docs/ops/fuhai-deep-audit-2026-05-31.md` or equivalent summary is merged/indexed |
| `deprecation/terminal-bridge-source-map` | `4e00051d` | archive-to-doc / legacy boxing | Preserve the source-map deprecation marker as terminal-bridge legacy evidence | Delete only after marker is merged, copied to legacy docs, or recorded here as intentionally discarded |
| `distribution-safety-gate-2026-06-01` | `3465aab3` | archive-to-doc / maybe-merge-docs | Reclassified: no longer just old #103 residue; contains new docs/ideas evidence | Delete only after new docs/ideas evidence is merged, extracted, or explicitly indexed |
| `docs/csv-layer-policy-98` | `f634d5dc` | close-after-index | Preserve historical CSV policy only | Delete only after useful wording is indexed; do not restore `_all.csv` pair runtime route |
| `fix/sense-emotion-layered-csv-80` | `af3aca03` | close-after-index | Preserve test / formatting ideas only | Delete only after useful test idea is indexed; do not merge obsolete layered runtime code |
| `docs/getnote-revoke-rotate-path` | `37d6e157` | hold-active / do-not-clean-now | Skip in this cleanup pass | Do not delete, archive, or fold unless user explicitly reopens GetNote cleanup |
| `refactor/prioritize-knowledge-dirs` | `9b6f2076` | keep-active | Active local dirty / side-mainline branch | Do not delete until local state is clean and user confirms |
| `wo/grade-scale` | `0ea8764f` | likely delete-if-evidence-preserved | Earlier index maps it to #84 merged / `wo/grade.md` | Delete after local guard or explicit PR evidence check |
| `wo/identity` | `68234813` | unknown / later | Needs PR/diff review | Do not delete in this pass |

## Branch notes

### `audit/fuhai-deep-2026-05-31`

Observed local / remote evidence:

```text
82506117 docs(audit): add fuhai deep audit report
1 file changed, 178 insertions(+)
docs/ops/fuhai-deep-audit-2026-05-31.md
```

Linked PR: #87, closed draft, unmerged.

Value to preserve:

- One-time Fuhai deep structure audit.
- Covers Notion-side topology, Fuhai seven gates, Gongyueshu/L1/L2 boundary, and Git/Worang landing context.
- Lists defects only; no fixes.

Posture:

- Do not delete until the audit report is merged or summarized in an audit index.
- If preserving by index only, keep at minimum the branch name, head SHA, file path, PR #87, and the summary above.

### `deprecation/terminal-bridge-source-map`

Observed evidence:

```text
4e00051d docs(wo): add terminal bridge deprecation marker to source-map
1 file changed, 4 insertions(+)
wo/source-map.md
```

Value to preserve:

- Source-map marker that terminal bridge is deprecated / legacy.

Posture:

- Legacy boxing only.
- Do not restore terminal bridge as current path.
- If the 4-line marker is not needed in `main`, this index may serve as the preservation record before branch deletion.

### `distribution-safety-gate-2026-06-01`

Observed evidence after recheck:

```text
3465aab3 docs(ops): record tool authorization boundary
a8a63480 docs(ideas): collect meta-skill beachcombing pearls
6 files changed, 101 insertions(+)
```

Important correction:

- This branch is no longer only the old #103 distribution safety gate residue.
- The strict ancestor guard blocked deletion, correctly.
- It now contains new unmerged docs / ideas evidence.

Value to preserve:

- Tool-action authorization boundary.
- Meta-skill beachcombing pearls.

Posture:

- Reclassify from `delete-if-evidence-preserved` to `archive-to-doc / maybe-merge-docs`.
- Do not delete until these docs/ideas are merged, extracted, or explicitly indexed as discarded.

### `docs/csv-layer-policy-98`

Observed evidence:

```text
f634d5dc docs: record csv layer policy
3 files changed, 128 insertions(+), 7 deletions(-)
AGENTS.md
wo/csv-layer-policy.md
wo/export-cleaning.md
```

Linked PR: #99, closed unmerged.

Superseded by:

- #102 `data: canonicalize csv sources`, merged.

Value to preserve:

- Historical CSV layer policy wording.
- Export-cleaning cautions around merge/delete/write safety.

Forbidden restoration:

- Do not restore ordinary `.csv` / `*_all.csv` pair runtime assumptions.
- Do not revive this as the current CSV route.

Posture:

- Extract useful historical policy wording if needed.
- Otherwise this index is enough to support deletion after user confirmation.

### `fix/sense-emotion-layered-csv-80`

Observed evidence:

```text
af3aca03 fix(sense): read layered emotion csv
2 files changed, 146 insertions(+), 23 deletions(-)
src/sense.rs
wo/sense.md
```

Linked PR: #97, closed unmerged.

Superseded by:

- #102 `data: canonicalize csv sources`, merged.

Value to preserve:

- Test idea for layered emotion merging.
- Emotion temperature / formatting idea.

Forbidden restoration:

- Do not merge the obsolete layered runtime implementation into current `sense`.
- Do not restore `_all.csv` pair runtime route.

Posture:

- Extract test/formatting idea if needed.
- Otherwise this index is enough to support deletion after user confirmation.

### `docs/getnote-revoke-rotate-path`

User boundary:

- This may be active work with another AI around GetNote / 得到大脑 Get 笔记.
- Skip it in this cleanup pass.

Posture:

- Do not inspect further for deletion.
- Do not archive, fold, or delete unless the user explicitly reopens GetNote cleanup.

## Candidate next commands

### Safe read-only audit

```bash
for BR in \
  audit/fuhai-deep-2026-05-31 \
  deprecation/terminal-bridge-source-map \
  distribution-safety-gate-2026-06-01 \
  docs/csv-layer-policy-98 \
  fix/sense-emotion-layered-csv-80 \
  wo/grade-scale \
  wo/identity

do
  echo
  echo "===== $BR ====="
  git log --oneline --decorate origin/main.."origin/$BR"
  git diff --stat origin/main..."origin/$BR"
done
```

### Do not run yet

```bash
# Do not delete GetNote active work.
git push origin --delete docs/getnote-revoke-rotate-path

# Do not delete refactor active local branch.
git push origin --delete refactor/prioritize-knowledge-dirs

# Do not delete all spike branches by pattern.
git push origin --delete $(git branch -r | grep 'wasteland/spike')
```

## Next suggested order

1. Merge this continuation index or keep it as a candidate PR.
2. If user wants small cleanup before spike work, check `wo/grade-scale` first because the previous index already maps it to #84 merged evidence.
3. Review `wo/identity` separately; it has no current decision in this continuation index.
4. Decide whether `audit/fuhai-deep-2026-05-31` and the terminal / CSV branches can be deleted with this index as sufficient preservation, or whether they need content extraction first.
5. Only after that, begin `wasteland/spike*` indexing.
