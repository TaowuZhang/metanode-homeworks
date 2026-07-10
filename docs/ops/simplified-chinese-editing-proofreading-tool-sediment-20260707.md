# Simplified Chinese editing/proofreading tool sediment · 2026-07-07

## Source and status

- source: `TaowuZhang/simplified-chinese-editing-proofreading-tool`
- GitHub state: private, active, `archived=false`, last pushed 2025-03-15
- local checkout: none found under `/Users/jiajia` to depth 6

## What exists

The repository is a Tauri 2 + SvelteKit 5 + Rust desktop scaffold. The remote tree has 50 files, including seven Rust files. It contains proposed `process_text` and `load_rules` commands plus local/AI processor types, but the live Tauri builder and Svelte page still wire only the template `greet` command. The proofreading surface is therefore incomplete and not a production tool.

## Canonical relationship

`技能/编校.md` is already the canonical judgment and quality-control process for editing, independent review, proofreading, title review, and release. This software repository does not replace that skill. Its durable residue is the possible product shape: a desktop wrapper that loads rules and delegates text processing while keeping the human review protocol separate.

## Disposition

The product topology and incomplete implementation state are preserved here. Full source and application icons were not copied. Archive remains blocked by unavailable local dirty state and the need to explicitly decide whether the unfinished implementation is discardable. If retired, preserve `技能/编校.md` as the process source and treat this repository only as historical product exploration.

## Implementation-state decision

- `src-tauri/src/commands.rs` proposes `process_text` and `load_rules` commands.
- `src-tauri/src/text_processor/mod.rs` proposes local and AI processors.
- the active Tauri builder registers only the template `greet` command;
- the active Svelte page calls only `greet`;
- therefore the proofreading commands/processors are not wired into a usable product path.

This is an implementation prototype, not a skill definition. It remains in `docs/ops`; `技能/编校.md` should not absorb incomplete software details. The remaining choice is whether to finish the wiring in an explicitly maintained product or mark the scaffold/source/icons historical-only before archive.

## Round 5 finish-or-historicize gate

| option | meaning | required action | consequence |
|---|---|---|---|
| finish | resume as an explicitly maintained desktop product | choose a maintained code target, wire commands/processors into Tauri/Svelte, define rule/data contracts, add tests, then reassess repository topology | creates a new product maintenance commitment; `技能/编校.md` remains the process source |
| historicize | preserve only as an unfinished product prototype | keep this sediment and repository history, explicitly mark source/icons historical-only, establish local dirty state, then archive | no parallel tool is maintained; canonical editing judgment remains in `技能/编校.md` |

Recommended default: `historicize`, because the active UI/backend still run the template greeting path and the durable process already exists in `技能/编校.md`. Decision `D-chinese-proofreading-tool` belongs to the user; Codex does not create a new skill or product commitment automatically.
