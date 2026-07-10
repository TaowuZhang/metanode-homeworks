# WQB Private Expression Workspace Policy — 2026-06-24

Full expressions are local-private only.

Git artifacts may contain only:

- expression_ref
- structural metadata
- private expression hash status
- redacted summary
- validator results
- risk score
- run-state

## Private paths

The following paths are ignored by Git and must be used for local-only draft bodies:

```text
local/private/
runs/wqb-private/
*.private.json
*.private.ndjson
```

Recommended local private file:

```text
local/private/wqb/batch-f-expression-drafts-private.json
```

## Current cloud-agent decision

This run cannot create or verify a local private workspace from the cloud GitHub-only execution surface. Therefore it did not create any private draft body and produced a local expression draft request instead.

```text
full_expression_committed = false
private_expression_hash_ready = false
simulation_dispatched = false
submit_attempted = false
batch_f_ready = false
authorization_ready_artifact_created = false
```
