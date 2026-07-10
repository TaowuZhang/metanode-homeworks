# WQB Alpha Consolidation Architecture · 2026-07-07

## Purpose

This file records the local Codex inventory supplement and turns it into the current `worang` consolidation contract for WQB / Alpha research assets.

## Freeze baseline

Local Codex reported:

- baseline to freeze: `main@393bc06`;
- export `git status --porcelain` before further construction;
- confirm whether 213 deletions are planned document migration before continuing on dirty `main`.

This repository-side file does not validate the local dirty tree. It records the gate that local Codex must satisfy before mutation.

## Canonical source routing

| asset class | canonical source | route |
|---|---|---|
| current code | `Alpha-Template-Automation/src` | migrate into the future `wqb-alpha/app` / packages only after baseline freeze |
| lightweight research tool | `Easy WQB` | preserve as lightweight loop; map into `wqb-alpha/app` or `research-store` only after code contract review |
| knowledge judgment | `worang/领域/量化` | remains the judgment source |
| old POC / workers / snapshots | local repos and old dirs | archive/runtime/vendor unless explicitly promoted |

## Target directory shape

```text
wqb-alpha/
├── app/                    # main pipeline and API
├── packages/
│   ├── expression-core/    # types, AST, operator/field contracts
│   ├── wqb-client/         # auth, retry, simulation, read-only queries
│   └── research-store/     # SQLite schema and repository layer
├── research/
│   ├── hypotheses/         # human-readable hypothesis cards
│   ├── experiments/        # one experiment per run_id
│   ├── notebooks/          # analysis; not production logic
│   └── factor-library/     # normalized factor knowledge
├── data/
│   ├── manifests/
│   ├── raw/                # read-only, partitioned by fetch date
│   └── derived/            # probes, heatmaps, features
├── runtime/                # db/log/checkpoints; gitignored
├── tests/
└── archive/                # POC, old Worker, historical snapshot index
```

## Asset manifest contract

Field caches, Operators, Universe, and Neutralization files must have an `assets-manifest.json` entry:

```json
{
  "source": null,
  "fetched_at": null,
  "region": null,
  "delay": null,
  "sha256": null,
  "row_count": null,
  "schema_version": null
}
```

## Unified Alpha schema

Each Alpha record must preserve at least:

```yaml
alpha_key: logic_hash + environment_hash
name: human-readable
expression: FASTEXPR
hypothesis: why it should work
fields: []
operators: []
settings:
  region: IND
  universe: TOP500
  delay: 1
  decay: 4
  neutralization: SUBINDUSTRY
  truncation: 0.08
lineage:
  parent_alpha_key: null
  mutation: null
evidence:
  simulation_id: null
  run_id: null
  in_sample: {sharpe: null, fitness: null, turnover: null, returns: null, drawdown: null, margin: null}
  checks: []
  self_corr: null
  prod_corr: null
  neutralization_counterfactuals: []
status: drafted | syntax_valid | simulated | cycle1_passed | robust | submitted | rejected
provenance: {created_at: null, code_commit: null, data_manifest: null}
```

## Experiment record template

Every experiment must record:

- `run_id`;
- code commit;
- field manifest;
- platform environment;
- candidate count;
- unique `logic_hash` count;
- preregistered primary metrics;
- preregistered stop conditions;
- all attempts, not only the best attempt;
- neutralization / window / decay as explicit dimensions;
- raw and counterfactual results;
- DSR / multiple-testing correction;
- time split;
- region / universe robustness;
- correlation matrix;
- `runtime_conditions` for platform 429, billing, and startup failure, separate from expression failure rate.

## Priority repairs

1. **Data contract repair**: audit the source of `cycle2_passed` 3,347 rows. Records missing metrics / region / category must not be treated as passed.
2. **Light loop closure**: make `wqb_loop.py` UPSERT into `wqb_db.py` with `run_id` and `logic_hash`.
3. **Judgment as tests**: implement LLM Kill Test, CROWDING counterfactual, and DSR screening as repeatable scripts or notebooks.

## Local paths reported by Codex

| role | path | route |
|---|---|---|
| main code | `/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/src/` | `wqb-alpha/app` and packages after freeze |
| main config | `/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/config/` | manifest-backed config import |
| main runtime DB | `/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/state/brain.db` | read-only audit; do not commit raw runtime DB |
| active nursery | `/Users/jiajia/Documents/GitHub/Alpha-Template-Automation/alpha_nursery/` | archive or normalize into experiments |
| light research | `/Users/jiajia/Documents/GitHub/Easy WQB/` | lightweight loop candidate |
| expression cards | `/Users/jiajia/Documents/GitHub/Alpha/` | migrate durable cards into `research/hypotheses` or `factor-library` |
| LLM notebook | `/Users/jiajia/Documents/GitHub/AIACv2_v1/ai_alpha_comp.ipynb` | notebook archive or repeatable test extraction |
| judgment | `/Users/jiajia/worang/领域/量化/亭子_v1.md` | judgment source |
| platform worker | `/Users/jiajia/wqb-cnhkmcp-cloudflare-worker/wqb-cnhkmcp-mcp/src/index.js` | archive/runtime/vendor unless promoted |
| maintenance logs | `/Users/jiajia/Documents/资料/系统维护记录/wqb_*.txt` | extract safe summaries only |

## Manual review queue

- `config/worldquant-brain-platform.mcp.json` contains only non-secret connection config.
- `cycle2_passed` 3,347 missing-metric rows are migration, fixture, or business pass.
- 213 deletes in dirty main are planned document migration.
- Five lost cockpit knowledge pages still exist in Notion or migrated repositories.
- IND `mdl192` CROWDING neutralization retry result.
- Easy WQB extreme metrics: test period, book size, data integrity, raw platform response.
- `alpha_nursery` duplicate submission / duplicate count risk after platform 429 recovery.

## Repository-cleanup implication

The local Codex paths are not currently visible as `TaowuZhang` GitHub repositories through this connector. They are treated as local repositories or directories until local Git remotes prove otherwise. Final GitHub cleanup still targets all non-`worang` maintained repositories, but local-only assets are first migrated or archived through this WQB consolidation contract.
