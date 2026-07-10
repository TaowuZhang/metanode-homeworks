# Workflow evidence latest

This is a read-only exception surface. It does not score code quality and does not treat a quiet repository as proof of effectiveness.

## GetNote mirror

- state: **healthy-current**
- severity: quiet
- interpretation: The newest real attempt is also the last accepted healthy mirror.
- workflow: `sync-getnote.yml`
- evidence kind: `real_external_read`

### Last attempt

- run: [28622300627](https://github.com/dongxi-heji/worang/actions/runs/28622300627)
- status: completed
- conclusion: success
- event: schedule
- head SHA: 00ed45f9bbe85a48a8f3e8f1b63255be65f7a4ba
- started: 2026-07-02T21:20:14Z

### Last healthy

- run id: 28622300627
- accepted at: 2026-07-02T21:20:29.435Z
- acceptance status: healthy
- source SHA: 00ed45f9bbe85a48a8f3e8f1b63255be65f7a4ba

## Reading rule

- **last attempt** answers what most recently happened in the workflow;
- **last healthy** answers which real external run was last accepted into main;
- a failed or unexecuted attempt leaves last healthy untouched;
- a repair is not closed until the same real workflow path succeeds again.