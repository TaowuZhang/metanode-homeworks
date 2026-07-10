# wo action inventory：current machine-surface facts

Status: current observed fact inventory  
Fact baseline: `main` at `0b684885e9faeb4e66c2e91d3aa42b26b68ac59c`  
Normative runtime owner: `docs/ops/wo-runtime-v0.md` after its PR is updated and merged  
This document does not define the repository-wide runtime contract.

## 1. Ownership and scope

This document owns the inventory of actions implemented by the current Rust `wo` binary: observable inputs and outputs, local side effects, subprocess boundaries, current machine/event surfaces, and fact-level gaps.

It does not own repository-wide stream rules, stable exit numbers, global non-interactive behavior, process lifecycle, a shared machine envelope, shared statuses, receipts, external-change records, or a cross-action compliance matrix. Those are normative runtime questions for the later revision of PR #342 and `docs/ops/wo-runtime-v0.md`. Until that PR is updated and merged, its proposed runtime rules are candidate design, not facts present on `main`.

Action-specific schemas remain owned by their action documents and implementations. In particular, current `excrete --json` output is an action-specific observed shape; PR #340 remains the candidate owner of its future contract.

## 2. Current action inventory

The current top-level surface is 15 Clap subcommands plus one implicit free-text write path. `stdin` below means application-level reads or prompts; current action code does not read stdin or ask for confirmation.

| Action | Implemented | Read-only / local writes | External capability | Auth | Prompt / stdin | Explicit input file | Current stdout | Current stderr | Current exit behavior | Current machine/event surface | AI non-interactive suitability | Evidence |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| `make` | yes | writes digest event | none | none | none | none | human project view | shadow JSON + `WO_LAST_EVENT_ID` export | observed 0; top-level errors usually 1 | shadow + digest JSONL | no: mixed human/diagnostic surface | `src/main.rs`, `src/make.rs`, `src/digest.rs` |
| `ask` | yes | writes digest event | none | none | none | none | human questions | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no | `src/main.rs`, `src/ask.rs` |
| `smell` | yes | reads traces; writes digest event | none | none | none | none | human residue view | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no | `src/main.rs`, `src/smell.rs` |
| `spark` | yes | writes spark state and digest event | none | none | repeated argv stages; no stdin | none | human dialogue | shadow JSON + export | observed 0/1 | shadow + digest JSONL | partial: stateful argv protocol | `src/main.rs`, `src/spark.rs` |
| `touch` | yes | may write residue/event state | optional NLM subprocess | environment/setup dependent | none | no generic file flag | human pack/NLM result | shadow JSON + export; failures vary | observed 0/1 | shadow + digest JSONL | no: external failure surface is not classified | `src/main.rs`, `src/touch.rs`, `scripts/nlm-grain-probe.sh` |
| `palm` | yes | reads/writes/clears `成为/手心/current.md`; writes event | none | none | none | none | human palm view | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no: read/write modes share one action | `src/main.rs`, `src/palm.rs` |
| `taste` | yes | may update palm; writes digest event | none | none | none | none | human taste view | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no | `src/main.rs`, `src/sense.rs` |
| `dream` | yes | may update palm; no real digest event | none | none | none | none | human dream view | shadow JSON + export | observed 0 unless process failure | shadow only | no: machine result absent | `src/main.rs`, `src/sense.rs` |
| `reveal` | yes | may update palm; no real digest event | none | none | none | none | human reveal view | shadow JSON + export | observed 0 unless process failure | shadow only | no: machine result absent | `src/main.rs`, `src/sense.rs` |
| `sense` (`inhale`, `recognize`, `notice`, `exhale`) | yes | writes digest event; `notice` performs local stale check | local Python subprocess for `notice` | none | none | none | human perception view | shadow JSON + export | observed 0/1; child status is flattened | shadow + digest JSONL | no: subprocess/status surface is not classified | `src/main.rs`, `src/sense.rs` |
| `tend` | yes | writes digest event | none | none | none | none | human questions | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no | `src/main.rs`, `src/tend.rs` |
| `bite` | yes | writes digest event | none in current keyword fallback | none | none | none | human sentence view | shadow JSON + export | observed 0/1 | shadow + digest JSONL | no | `src/main.rs`, `src/bite.rs` |
| `weigh` | yes | Rust delegates effects to probe | Node/Maimemo probe | probe-dependent | none in Rust | none | inherited child stdout | inherited child stderr | observed 0/1; child code is flattened | no Rust event/result schema | no: dependency/auth behavior is delegated | `src/main.rs`, `scripts/maimemo-probe.mjs` |
| `excrete` | yes | inspect is read-only; selection writes event; prepare may write dry-run artifacts | Node worker/adapters for prepare | adapter-dependent | none | none | human output or action JSON | worker diagnostic on failure | observed 0/1; invalid Clap value is 2 | inspect JSON and excrete JSONL event | inspect-only `--json` fits an empty isolated root; write modes do not | `src/main.rs`, `src/cli.rs`, `src/excrete.rs`, `scripts/wo-excrete.mjs` |
| `surface` | yes | reads digest snapshot; no verified write | none | none | none | none | human pressure view | empty on success | observed 0/1 | none | partial: read-only but human-only | `src/main.rs`, `src/digest.rs` |
| free text (`wo <words>`) | yes | writes trace CSV and daily note | none | none | none | none | human acknowledgement with paths | empty on success | observed 0/1 | no digest or response schema | no: implicit write has no preview | `src/main.rs`, `src/trace.rs` |

## 3. Current observed CLI baseline

These are observed implementation behaviors, not a stable public contract:

- root `--help` and all 15 subcommand `--help` calls are offline, exit 0, write `Usage:` to stdout, and keep stderr empty;
- an unknown option is rejected by Clap with empty stdout, an `error:` diagnostic on stderr, and currently exits 2;
- `wo excrete --mode unsafe` has the same observed parse-error stream shape and currently exits 2;
- an unknown bare word is not a command error: it enters the implicit free-text write path and may exit 0 after writing the trace CSV/daily note, or 1 on a top-level error;
- top-level Rust errors commonly surface on stderr and currently produce exit 1;
- on an empty isolated `WO_ROOT`, `wo excrete --json` exits 0, keeps stderr empty, creates no file, and writes exactly:

```json
{"schema_version":1,"command":"excrete","mode":"inspect","candidates":[]}
```

The numbers 1 and 2 above describe the present implementation only. This inventory does not reserve or stabilize them.

## 4. Current coverage and gaps

Current coverage:

- help and Clap parse-error stream behavior is exercised offline;
- `excrete --json` inspect has an action-specific JSON result and a no-side-effect golden fixture;
- digest and excrete records are JSONL; shadow payloads have Rust serialization tests;
- no current Rust action reads application stdin or opens a confirmation prompt.

Fact-level gaps:

- only `excrete` exposes direct result JSON; most actions expose human stdout plus shadow diagnostics;
- event coverage differs: most actions emit real + shadow events, `dream`/`reveal` emit shadow only, and `weigh`/`surface`/free text do not expose the same event surface;
- write actions have no shared preview or approval mechanism;
- subprocess actions flatten or render child failures differently;
- current failures commonly collapse to top-level 1, while Clap parse errors currently use 2;
- most actions are not suitable for direct unattended AI invocation because their result, mutation, or external boundary is not machine-explicit.

## 5. Handoff to runtime specification

The later #342 runtime specification owns unresolved normative decisions for:

- stdout/stderr policy;
- exit taxonomy;
- global non-interactive behavior;
- child TERM → KILL → reap → cleanup lifecycle;
- shared machine output and statuses;
- receipts and external-change records;
- cross-action runtime compliance matrix.

This document supplies observed facts to that work; it does not duplicate those rules.

## 6. External candidate boundaries

- PR #341 (GWS read-only adapter) is Open Draft and remains candidate behavior, not current `main` fact.
- PR #333 (DuckDB read-only lens) is Open Draft and remains candidate behavior, not current `main` fact.
- PR #340 (`excrete` JSON contract) is Open Draft and remains candidate behavior, not current `main` fact.
- PR #342 (runtime specification) is Open Draft and remains candidate normative design until revised and merged.
- None of these candidates makes an external tool a new Rust `wo` subcommand.

## 7. Offline baseline checker

The checker is intentionally narrow and read-only. The caller builds the binary outside the repository and passes it explicitly:

```bash
CARGO_TARGET_DIR=/private/tmp/worang-339-cargo-target cargo build
WO_BIN=/private/tmp/worang-339-cargo-target/debug/wo \
  node scripts/check-wo-action-inventory.mjs
```

It checks the root and 15 subcommand help surfaces, two parse failures, the empty-root `excrete --json` golden result, document inventory completeness, forbidden runtime-ownership claims, timeout bounds, and fixture cleanliness. It isolates `WO_ROOT`, `HOME`, XDG directories, and upper/lower-case proxy variables, then removes the entire temporary sandbox in `finally`.
