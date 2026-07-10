# Branch governance

Purpose: keep remote branches useful as active work pointers, not as a long-term archive.

This guide was extracted from the 2026-06-02 remote branch cleanup pass.

## Core principle

Branches are working pointers, not archives.

A branch can be deleted after its useful evidence has landed in at least one durable place:

- merged PR history;
- docs / ops notes;
- an issue ledger;
- a tag for a meaningful snapshot;
- an explicit archive index.

Do not delete branches that are still active, ambiguous, or carrying unpreserved evidence.

## Branch classes

| Class | Meaning | Default action |
|---|---|---|
| `main` | active mainline | keep |
| active work | currently being developed | keep |
| hold-active | likely active, owner/context not settled | keep until explicitly reopened |
| merged PR head | branch whose work landed through a PR | delete by exact name after evidence check |
| spike / experiment | exploratory branch | index first, then delete by exact name |
| archive candidate | contains useful branch-only evidence | preserve to docs / issue / tag first |
| stale duplicate | no independent useful evidence | delete by exact name after review |

## Deletion rules

- Never pattern-delete remote branches.
- Never run broad commands such as:

```bash
git push origin --delete $(git branch -r | grep 'wasteland/spike')
```

- Prefer exact-name deletion:

```bash
git push origin --delete path/to/exact-branch-name
```

- Before deletion, confirm at least one of:
  - the branch is an ancestor of `origin/main`;
  - its PR is closed and merged;
  - its useful contents are preserved in docs / issue / tag / archive index;
  - the branch is explicitly classified as stale duplicate.

## Spike branch rule

Spike branches are especially easy to misuse as archives.

Before deleting a spike branch group:

1. create or update an archive index;
2. group related branches by topic;
3. map each branch to PR evidence or diff evidence;
4. mark each branch as one of:
   - `delete-ready`;
   - `tag-first`;
   - `keep-review`;
5. delete only the `delete-ready` branches by exact name.

## Tagging rule

Use tags for meaningful snapshots, not for every deleted branch.

A good tag should mark one of:

- a release boundary;
- a preserved experimental milestone;
- a migration phase boundary;
- an externally referenced state;
- a point that future readers may need to check out directly.

Do not tag every merged PR head. PR history is enough for ordinary merged work.

Suggested tag shapes:

```text
archive/<topic>/<yyyy-mm-dd>
exp/<topic>/<phase>
release/<version>
ops/<topic>/<yyyy-mm-dd>
```

## Issue ledger rule

Large cleanup passes should have a ledger issue.

The ledger should record:

- initial branch inventory;
- branches held out of scope;
- evidence before deletion;
- exact deletion commands;
- deletion results;
- final remote branch state.

## AI-assisted cleanup boundary

AI may help review PRs, branch evidence, docs, and issue ledgers, and may merge candidate PRs into `main` on its own when the merge carries no irreducible value conflict. The single handback criterion is an irreducible value conflict (价值之书 §五), not actor identity; there is no separate “Ke must not merge his own PR” gate.

Stop for explicit human authorization only before actions that destroy remote pointers or create ambiguous snapshots, including:

- deleting remote branches;
- force-overwriting files on `main`;
- tagging ambiguous historical points;
- merges that would trigger a release / tag / permission escalation / publish.

Ordinary candidate-PR merges no longer require human authorization. When in doubt, preserve evidence first and delete later.

## Current held branch examples

As of the 2026-06-02 cleanup pass, these were intentionally left out of deletion:

```text
docs/getnote-revoke-rotate-path
getnote-smoke-work
refactor/prioritize-knowledge-dirs
```

They should be reviewed in separate workstreams, not as part of stale branch cleanup.
