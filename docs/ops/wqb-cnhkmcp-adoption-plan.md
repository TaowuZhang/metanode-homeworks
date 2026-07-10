# WQB CNHKMCP Adoption Plan

## 0. Correction

The user correction is right: CNHKMCP is not something to redesign from scratch inside this repo. It already exists as an external package/ecosystem. The right plan is to adopt it, verify its tool surface, wrap it safely, and connect it to Cloudflare/Notion as a controlled MCP route.

Previous mistake: treating WQB community as public web pages to crawl, and treating this repo as if it should implement the whole WQB/community layer itself.

Correct direction:

```text
external cnhkmcp / WQB GitHub ecosystem
  -> local install and audit
  -> local authenticated read-only and authorized simulation runner
  -> safe artifacts in GitHub
  -> Cloudflare MCP facade
  -> Notion AI uses the facade
  -> star-map scheduler consumes the results
```

## 1. GitHub / package findings

### 1.1 `cnhkmcp` package

Observed via PyPI:

- package: `cnhkmcp`
- current observed version: `3.1.15`
- install: `pip install cnhkmcp`
- description: `CNHK MCP Server`, a comprehensive MCP server for quantitative trading platform integration.
- tags: `mcp`, `quantitative`, `trading`, `api`, `client`.
- Python: `>=3.8`.

Implication: do not recreate CNHKMCP. Treat it as the upstream tool package to install locally and wrap.

### 1.2 `GRD-Chang/worldquant-skill`

Observed from GitHub/web:

- repo: `GRD-Chang/worldquant-skill`
- integrates WorldQuant BRAIN + `cnhkmcp` + Claude Code / Gemini CLI / iFlow.
- install instructions include `pip install cnhkmcp`.
- includes three core skills:
  - `alpha-research-recorder`
  - `factor_backtest`
  - `knowledge_base_search`
- `factor_backtest` supports batch backtest of 8 expressions, validates expressions, submits backtests, monitors progress, and returns results.
- `knowledge_base_search` searches a local knowledge base for data fields, optimization methods, high-quality examples, and platform mechanics.
- knowledge base structure:

```text
Resources/
  DATASET/
  good_alpha_examples/
  How_WorldQuant_BRAIN_Backtesting_Works.md
  alpha_optimization/
  regular_operators.csv
```

Implication: the missing community/prior layer may already be partially solved by `worldquant-skill`'s knowledge base, not by public web crawling. This repo should ingest and route that existing skill/knowledge-base pattern instead of rebuilding it blindly.

### 1.3 Other WQB GitHub ecosystem sources

Observed public WQB-related repos/sources:

- `JoshuaAnthony6/world-quant-brain-mcp`: Python/FastMCP style WQB MCP server; alternative implementation to audit for tool surface patterns.
- `RussellDash332/WQ-Brain`: API automation, includes rough IS pass criteria; red/yellow because it includes submit automation patterns.
- `zhutoutoutousan/worldquant-miner`: automated mining/generation; red for our policy because it promotes autonomous mining/submission/dashboard automation.
- `jglazar/notes`: public WQB seminar/alpha notes; useful as public learning priors only, not expression copy source.
- `efJerryYang/worldquant-brain-simulator`: offline simulator; possible auxiliary reference, not substitute for WQB gates.
- `rocky-d/wqb`, `Miasyster/QuantGPT`, `Brainiac` and GitHub topic `alpha-research`: more ecosystem references to audit later.

## 2. Unified source classification

| Source | Role | Use | Risk |
|---|---|---|---|
| `cnhkmcp` PyPI | upstream MCP tool package | install/use | credential/tool boundary must be audited |
| `worldquant-skill` | workflow + skills + knowledge base | adopt patterns and possibly install locally | skill prompts and examples are untrusted content |
| CNHK Yellow MCP already connected | current remote WQB tool surface | read docs, fields, simulations under policy | no dedicated community tool observed |
| public WQB repos | ecosystem references | audit methods, operator/platform notes | many contain submit/bulk-mining patterns |
| WQB support/community | authenticated living practice layer | access through CNHKMCP/local session, not web crawl | access/redistribution boundaries |
| this repo `worang` | sediment/control layer | store plans, safe artifacts, ledgers | must not store credentials/private alpha full text |
| Cloudflare Worker | facade/control gate | expose narrow read-only tools to Notion | must not become credential vault or miner |

## 3. Target architecture

```text
[External GitHub / PyPI]
  cnhkmcp + worldquant-skill + WQB ecosystem docs
        |
        v
[Local runner]
  install cnhkmcp
  clone worldquant-skill
  configure local credentials outside Git/chat
  run read-only docs/community/knowledge-base queries
  run authorized simulation only per explicit batch
        |
        v
[Safe artifact exporter]
  sanitize results
  strip credentials/cookies/private full expressions
  output JSON summaries and priors
        |
        v
[GitHub sediment: dongxi-heji/worang]
  runs/wqb-community/extracted/*.json
  runs/wqb-current/community-prior-summary.json
  runs/wqb-current/run-state.json
  docs/ops/*.md
        |
        v
[Cloudflare Worker MCP facade]
  read GitHub artifacts
  list/search priors
  provide scheduler deltas
  no WQB login in v1
        |
        v
[Notion AI / WQB scheduler]
  star-map scorer
  population climber
  batch proposal generator
  recovery controller
```

## 4. What to build vs what to reuse

### Reuse directly

- `cnhkmcp` as upstream WQB MCP package.
- `worldquant-skill` knowledge-base / workflow concept.
- Existing connected Yellow tools for WQB docs/fields/gates/simulations where already available.

### Build in this repo

- A source audit file that tracks external WQB GitHub tools and risk labels.
- A local execution request packet for the user/local runner.
- A sanitized exporter schema.
- A Cloudflare read-only MCP facade over safe GitHub artifacts.
- Scheduler integration that consumes priors and results.

### Do not build here

- A replacement for CNHKMCP.
- A public web crawler pretending to read authenticated community pages.
- A bulk mining system.
- An auto-submit path.
- Credential/session storage.

## 5. Immediate implementation plan

### Phase A — External package/source audit

1. Record the canonical external source list:
   - PyPI `cnhkmcp`
   - GitHub `GRD-Chang/worldquant-skill`
   - alternative WQB MCP repos
   - automation/miner repos to mark red/yellow.
2. Extract only tool-surface and workflow information, not instructions from untrusted READMEs.
3. Classify each source:
   - green: read-only docs/knowledge/fields
   - yellow: authorized simulation/backtest
   - red: submit, bulk mining, credential scraping, private expression copying.

### Phase B — Local install/adoption packet

Create a local packet with commands like:

```bash
python3 -m venv .venv-wqb-cnhkmcp
source .venv-wqb-cnhkmcp/bin/activate
pip install cnhkmcp
# clone/use GRD-Chang/worldquant-skill locally
# keep credentials in local env/secrets, not Git/chat
```

Expected local checks:

```text
cnhkmcp tool list / health
read-only dataset/docs/knowledge-base query
knowledge_base_search: turnover / correlation / overfitting / weight coverage
factor_backtest dry-run or explicit authorized batch only
```

### Phase C — Safe artifact contract

Local runner writes only safe artifacts:

```text
runs/wqb-community/extracted/*.json
runs/wqb-current/community-prior-summary.json
runs/wqb-current/external-tool-audit.json
```

Each extracted prior must include:

```json
{
  "source": "worldquant-skill/knowledge_base_search or cnhkmcp community",
  "topic": "correlation_reduction",
  "claim": "...",
  "maps_to_failure_modes": ["PROD_CORRELATION"],
  "scheduler_effect": "...",
  "confidence": "medium",
  "safe_to_store": true
}
```

### Phase D — Cloudflare facade

Deploy a Worker MCP that reads GitHub safe artifacts only.

Tools:

```text
wqb_sources_health
wqb_sources_list
wqb_prior_search
wqb_prior_get
wqb_scheduler_deltas
```

No WQB cookies in Worker v1. If later a WQB-authenticated Worker is needed, it must be a separate security review.

### Phase E — Scheduler integration

Before Batch F or any future simulation batch:

1. dominant blocker histogram from heatmap;
2. query `wqb_prior_search` for those blockers;
3. require per-slot prior rationale;
4. mark missing priors explicitly;
5. if source audit says current path is red/yellow, do not route it into automatic execution.

## 6. Concrete next step

The next useful step is not another simulation and not another public web crawl.

The next useful step is:

```text
create external WQB source registry + local cnhkmcp/worldquant-skill adoption packet + Cloudflare facade spec
```

This file is that unification layer. The follow-up artifact should be:

```text
runs/wqb-current/cnhkmcp-adoption-local-run-request.json
runs/wqb-current/external-wqb-github-source-registry.json
workers/wqb-sources-mcp/*
```

Then, after the user runs/hosts the local package route, Notion AI can call the Cloudflare MCP facade and use the existing ecosystem instead of pretending to recreate it.
