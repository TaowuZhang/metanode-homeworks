# Spike branch archive plan｜2026-06-02

Purpose: prepare `wasteland/spike*` remote branch cleanup without deleting experimental evidence.

This is an index / archive plan only. It does not delete branches. It exists because spike branches are not archives, but they may contain useful experimental evidence that should be represented before deletion.

## Rules

- Do not batch-delete by pattern.
- Do not run `git push origin --delete $(git branch -r | grep 'wasteland/spike')`.
- Delete only by exact branch name after the branch is marked `delete-ready` below.
- Prefer PR evidence when it exists.
- If no PR evidence exists but the experiment still has value, preserve a one-line index entry or archive tag first.
- Do not merge spike branches directly into `main`.
- `main`, GetNote active branches, and `refactor/prioritize-knowledge-dirs` are out of scope for this plan.

## Group summary

| Group | Branch pattern | Topic | Current recommendation |
|---|---|---|---|
| Early residue | `spike-0`, `spike-1`, `spike-2` | deps / raw text / raw residue / decay / palm excavation | index-first |
| Entry series | `spike-3-5*` | command entry variants | index-first |
| Shadow-event series | `spike-3-6*` | shadow event variants and fixes | index-first |
| Real-emit series | `spike-3-7*` | real emit variants | delete some by PR evidence; index remainder |
| Digest / mapping / trace | `spike-3a` to `spike-3f` | digest types, node mapping, snapshots, trace entry | index-first |

## Known merged PR evidence

These were already recorded in `docs/ops/branch-archive-2026-06-01.md` and can be treated as first delete candidates after local/remote confirmation:

| Branch | Evidence | Recommendation |
|---|---|---|
| `wasteland/spike-3-7j-sense-notice-real-emit` | PR #57 merged | delete-ready |
| `wasteland/spike-3-7k-sense-inhale-real-emit` | PR #58 merged | delete-ready |
| `wasteland/spike-3-7l-sense-recognize-real-emit` | PR #59 merged | delete-ready |
| `wasteland/spike-3-7m-sense-exhale-real-emit` | PR #60 merged | delete-ready |
| `wasteland/spike-3-7n-taste-real-emit` | PR #61 merged | delete-ready |

Exact-name deletion commands for the known merged set, after `git fetch origin --prune`:

```bash
git push origin --delete wasteland/spike-3-7j-sense-notice-real-emit
git push origin --delete wasteland/spike-3-7k-sense-inhale-real-emit
git push origin --delete wasteland/spike-3-7l-sense-recognize-real-emit
git push origin --delete wasteland/spike-3-7m-sense-exhale-real-emit
git push origin --delete wasteland/spike-3-7n-taste-real-emit
```

## Branch inventory

### Early residue

| Branch | Head | Archive posture | Delete condition |
|---|---|---|---|
| `wasteland/spike-0-deps` | `fd8ae80` | index-first | Review diff or link PR evidence |
| `wasteland/spike-1-5-raw-text` | `04843be` | index-first | Review diff or link PR evidence |
| `wasteland/spike-1-raw-residue` | `f2f76c0` | index-first | Review diff or link PR evidence |
| `wasteland/spike-2a-decay-view` | `a825915` | index-first | Review diff or link PR evidence |
| `wasteland/spike-2b-palm-excavate` | `ff3a803` | index-first | Review diff or link PR evidence |

### Entry series: `spike-3-5*`

| Branch | Head | Archive posture |
|---|---|---|
| `wasteland/spike-3-5a-make-entry` | `4128402` | index-first |
| `wasteland/spike-3-5b-ask-entry` | `6569e73` | index-first |
| `wasteland/spike-3-5c-bite-entry` | `2b6c198` | index-first |
| `wasteland/spike-3-5d-smell-entry` | `52a1a6b` | index-first |
| `wasteland/spike-3-5e-palm-entry` | `1c3500c` | index-first |
| `wasteland/spike-3-5f-spark-entry` | `41a7b83` | index-first |
| `wasteland/spike-3-5g-touch-entry` | `92754c0` | index-first |
| `wasteland/spike-3-5h-tend-entry` | `e676eca` | index-first |
| `wasteland/spike-3-5i-sense-default-entry` | `6ee6d6c` | index-first |
| `wasteland/spike-3-5j-sense-notice-entry` | `5a0e4f1` | index-first |
| `wasteland/spike-3-5k-sense-inhale-entry` | `d132df9` | index-first |
| `wasteland/spike-3-5l-sense-recognize-entry` | `fffa5b1` | index-first |
| `wasteland/spike-3-5m-sense-exhale-entry` | `cd2814f` | index-first |
| `wasteland/spike-3-5n-taste-entry` | `117ec9e` | index-first |
| `wasteland/spike-3-5o-dream-entry` | `d32d6b6` | index-first |
| `wasteland/spike-3-5p-reveal-entry` | `a288008` | index-first |
| `wasteland/spike-3-5q-dispatch-explicit` | `8168c0d` | index-first |

### Shadow-event series: `spike-3-6*`

| Branch | Head | Archive posture |
|---|---|---|
| `wasteland/spike-3-6a-make-shadow-event` | `dddb743` | index-first |
| `wasteland/spike-3-6b-ask-shadow-event` | `71d9fc9` | index-first |
| `wasteland/spike-3-6c-bite-shadow-event` | `d3ada5b` | index-first |
| `wasteland/spike-3-6d-smell-shadow-event` | `1a5eb35` | index-first |
| `wasteland/spike-3-6e-palm-shadow-event` | `9199cec` | index-first |
| `wasteland/spike-3-6f-spark-shadow-event` | `f765d89` | index-first |
| `wasteland/spike-3-6g-touch-shadow-event` | `3a4146b` | index-first |
| `wasteland/spike-3-6g-touch-shadow-event-fix` | `a3d5b5c` | index-first |
| `wasteland/spike-3-6h-tend-shadow-event` | `830bee5` | index-first |
| `wasteland/spike-3-6h-tend-shadow-event-fix` | `cc633ce` | index-first |
| `wasteland/spike-3-6i-sense-default-shadow-event` | `f9ebef2` | index-first |
| `wasteland/spike-3-6j-sense-notice-shadow-event` | `059d0c2` | index-first |
| `wasteland/spike-3-6k-sense-inhale-shadow-event` | `b3a187c` | index-first |
| `wasteland/spike-3-6l-sense-recognize-shadow-event` | `1047733` | index-first |
| `wasteland/spike-3-6m-sense-exhale-shadow-event` | `a884ab9` | index-first |
| `wasteland/spike-3-6n-taste-shadow-event` | `dcf0434` | index-first |
| `wasteland/spike-3-6o-dream-shadow-event` | `5242aa2` | index-first |
| `wasteland/spike-3-6p-reveal-shadow-event` | `3dda051` | index-first |

### Real-emit series: `spike-3-7*`

| Branch | Head | Archive posture |
|---|---|---|
| `wasteland/spike-3-7a-make-real-emit` | `bd131a3` | index-first |
| `wasteland/spike-3-7b-ask-real-emit` | `9817fbe` | index-first |
| `wasteland/spike-3-7c-bite-real-emit` | `09d4222` | index-first |
| `wasteland/spike-3-7d-smell-real-emit` | `885d297` | index-first |
| `wasteland/spike-3-7e-palm-real-emit` | `8263c41` | index-first |
| `wasteland/spike-3-7f-spark-real-emit` | `3e6960d` | index-first |
| `wasteland/spike-3-7g-touch-real-emit` | `41055d6` | index-first |
| `wasteland/spike-3-7h-tend-real-emit` | `2e20ccf` | index-first |
| `wasteland/spike-3-7i-sense-real-emit` | `df7f010` | index-first |
| `wasteland/spike-3-7j-sense-notice-real-emit` | `95a7572` | delete-ready via PR #57 |
| `wasteland/spike-3-7k-sense-inhale-real-emit` | `1d63bfa` | delete-ready via PR #58 |
| `wasteland/spike-3-7l-sense-recognize-real-emit` | `ff30452` | delete-ready via PR #59 |
| `wasteland/spike-3-7m-sense-exhale-real-emit` | `c2180a9` | delete-ready via PR #60 |
| `wasteland/spike-3-7n-taste-real-emit` | `fe2cea3` | delete-ready via PR #61 |

### Digest / mapping / trace series

| Branch | Head | Archive posture |
|---|---|---|
| `wasteland/spike-3a-digest-types` | `0b34bea` | index-first |
| `wasteland/spike-3b-node-mapping` | `37271f2` | index-first |
| `wasteland/spike-3b-plus-bias-demo` | `00cfcf0` | index-first |
| `wasteland/spike-3c-event-snapshot` | `5418986` | index-first |
| `wasteland/spike-3d-abort-snapshot` | `532e600` | index-first |
| `wasteland/spike-3e-sense-vomit` | `52d4fad` | index-first |
| `wasteland/spike-3f-trace-entry` | `abf3218` | index-first |

## Suggested cleanup phases

1. Merge this index to `main`.
2. Delete the five known merged real-emit branches by exact name.
3. For each remaining group, do one read-only review batch:
   - `git diff --stat origin/main...origin/<branch>`
   - search linked PR evidence by head branch
   - mark `delete-ready`, `tag-first`, or `keep-review`
4. Delete only exact-name `delete-ready` branches.

## Out of scope

```text
ci/getnote-sync-workflow-v0
docs/getnote-revoke-rotate-path
getnote-smoke-work
refactor/prioritize-knowledge-dirs
main
```
