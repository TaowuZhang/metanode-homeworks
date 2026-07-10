# Local Codex Task · GitHub Repository Cleanup · 2026-07-07

## Intent

Final target: only `dongxi-heji/worang` remains as the long-term maintained canonical repository.

The user confirms: the repositories/paths they listed are their repositories/assets. WQB / quant Alpha related repositories are excluded from immediate non-WQB cleanup and should follow the WQB consolidation contract. All other repositories should be cleaned, migrated, archived, or prepared for removal.

## Canonical keep

- `dongxi-heji/worang`

## WQB / quant Alpha protected set

Do not archive/delete these as generic repository cleanup. They belong to the WQB consolidation path:

```text
/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/
/Users/jiajia/Documents/GitHub/Easy WQB/
/Users/jiajia/Documents/GitHub/Alpha/
/Users/jiajia/Documents/GitHub/AIACv2_v1/
/Users/jiajia/wqb-cnhkmcp-cloudflare-worker/
/Users/jiajia/worang/
```

Route these through:

- `docs/ops/wqb-alpha-consolidation-architecture-20260707.md`

## Non-WQB cleanup targets observed from GitHub API

```text
TaowuZhang/CN-Mirror-CLI
TaowuZhang/Xget
TaowuZhang/neetcode-submissions
```

Also scan `/Users/jiajia/Documents/GitHub/` for any other local git repositories not in the WQB protected set.

## Hard boundaries

Do not perform irreversible actions before writing a local evidence packet:

- no `rm -rf` on repositories;
- no `git push --force`;
- no secret printing;
- no repository deletion before evidence packet exists;
- no deletion of unique assets before migration or explicit discard record;
- no mutation of WQB protected repositories except inventory/reporting.

Archive/delete can be completed only if all acceptance gates below pass. If using GitHub UI or `gh repo archive/delete`, record the exact command or UI step in the final report.

## Phase 1 — Freeze inventory

From local shell, produce:

```bash
mkdir -p /Users/jiajia/worang/runs/github-cleanup/2026-07-07
cd /Users/jiajia/Documents/GitHub

find . -maxdepth 2 -name .git -type d \
  | sed 's#^./##; s#/.git$##' \
  | sort > /Users/jiajia/worang/runs/github-cleanup/2026-07-07/local-git-repos.txt

for repo in */.git; do
  d=${repo%/.git}
  (
    cd "$d" || exit 0
    echo "## $d"
    echo "path: $(pwd)"
    echo "branch: $(git branch --show-current 2>/dev/null || true)"
    echo "head: $(git rev-parse --short HEAD 2>/dev/null || true)"
    echo "status:"
    git status --porcelain=v1 || true
    echo "remotes:"
    git remote -v || true
    echo
  )
done > /Users/jiajia/worang/runs/github-cleanup/2026-07-07/local-repo-status.md
```

## Phase 2 — Classify

Create:

```text
/Users/jiajia/worang/runs/github-cleanup/2026-07-07/repository-disposition.csv
```

Columns:

```csv
path,remote,owner_repo,is_wqb_related,has_uncommitted_changes,has_unique_assets,has_secrets_risk,has_pages_or_release,has_workflow_or_package,recommended_action,reason,next_manual_gate
```

Classification values:

- `keep`: only `worang`;
- `wqb-protected`: WQB / quant Alpha path;
- `migrate-then-archive`: durable material exists but belongs in `worang`;
- `archive`: history useful, no active maintenance;
- `delete-candidate`: no unique assets, no external dependency, no open work;
- `unknown`: missing evidence.

## Phase 3 — Non-WQB targets

For each non-WQB repo:

1. Read README / root tree / workflows / package files.
2. Check whether durable notes or scripts should move to `worang`.
3. If migration needed, copy only safe summaries or source files into the correct `worang` stratum.
4. Write one disposition card per repo:

```text
/Users/jiajia/worang/runs/github-cleanup/2026-07-07/repos/<safe-repo-name>.md
```

Card template:

```markdown
# Repository disposition: <repo>

- path:
- remote:
- current branch:
- head:
- dirty status:
- README:
- workflows/packages/releases/pages:
- unique assets:
- secrets risk:
- migration performed:
- recommended action:
- final gate:
```

## Phase 4 — Preferred routes for known non-WQB repos

### `CN-Mirror-CLI`

Likely route: `migrate-then-archive`.

Look for durable material about CN mirror setup, Xget prefix usage, safe publish script, and local mirror config. If useful, summarize or migrate into `worang` under either:

- `资源/链接/`
- `领域/代码/`
- `docs/ops/` if purely operational

### `Xget`

Likely route: `archive` or `migrate-then-archive`, but first check whether it is fork/upstream/product/reference. Do not delete if it has releases, stars, external links, or fork/upstream meaning.

### `neetcode-submissions`

Likely route: `delete-candidate` or `archive`, but only after checking root tree and dirty status. If it is pure coding-practice residue with no unique long-term notes, record deletion candidate.

## Phase 5 — GitHub final actions

If all gates pass, local Codex may perform final GitHub actions one-by-one. Prefer archive over delete unless the repo is truly disposable.

Before each final action, append to:

```text
/Users/jiajia/worang/runs/github-cleanup/2026-07-07/final-actions.md
```

Format:

```markdown
## <repo>

- action: archive | delete | leave
- reason:
- evidence card:
- migrated content:
- command_or_ui_step:
- timestamp:
```

## Acceptance gates

A repo may be archived/deleted only if:

- it is not `worang`;
- it is not WQB protected;
- local dirty status is recorded;
- remote URL is recorded;
- README/root tree were inspected;
- workflow/package/release/pages status was checked;
- unique assets were migrated or explicitly marked discardable;
- no secrets were printed or committed;
- final action was recorded in `final-actions.md`.

## Final report back to `worang`

After execution, update:

```text
docs/ops/github-repository-inventory-20260707.md
```

Add a section `Local Codex execution result` with:

- repos archived;
- repos deleted;
- repos left untouched;
- WQB protected repos confirmed;
- migrated files;
- unresolved gates.

Do not claim completion unless the GitHub UI/API state confirms it.
