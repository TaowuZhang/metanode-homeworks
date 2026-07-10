# GWS read-only adapter contract: Phase 2

Status: local fake-GWS process contract only. No real Google account, OAuth grant, Workspace payload, or installed GWS version has been tested.

## Covered action

The only covered action is today's read-only Calendar agenda through the already merged bridge:

```bash
node scripts/gws-bridge.mjs --no-input --timeout 30s calendar +agenda --today
```

This does not create a `wo gws` command or claim that `wo touch --gws` exists. Other GWS reads and every write path remain outside this contract. Authentication, plan/apply, rollback, and automatic approval remain forbidden.

## Invocation boundary

- `--no-input` is required and closes child stdin. It never sends a newline, accepts a prompt, authenticates, or treats `CI=true`/non-TTY as consent.
- `--timeout` is required. Accepted units are positive integer milliseconds (`ms`) or seconds (`s`), up to 24 hours.
- Timeout sends TERM to the complete child process group, waits 250 ms, then sends KILL and reaps the direct child.
- Parent SIGINT/SIGTERM uses the same process-group cleanup and exits 130/143 after writing an interrupted receipt when possible.
- The executable is invoked directly without a shell. Its path may contain spaces or quotes.

## Streams and exits

Success emits exactly one normalized JSON document to stdout. Safe child warnings are normalized and written to stderr. Invalid JSON, timeout, interruption, and failures leave stdout empty. Suspected credential values in child diagnostics are redacted and diagnostics are bounded.

| Exit | Meaning | Receipt status | Retryable |
|---:|---|---|---|
| 0 | success | `success` | no |
| 64 | invalid wrapper arguments | no receipt | after correcting input |
| 69 | GWS executable unavailable | `failed` | after installation |
| 70 | invalid output or wrapper/receipt failure | `failed` | after fixing contract/local state |
| 74 | unclassified external failure | `failed` | unknown |
| 75 | classified temporary external failure | `failed` | yes |
| 77 | authentication required or permission rejected | `needs_auth` or `rejected` | after human action |
| 124 | timeout | `timeout` | yes |
| 130 / 143 | parent interrupted by SIGINT / SIGTERM | `interrupted` | yes |

Arbitrary child exit codes are recorded as `adapter_exit_code` but never passed through as the public exit code.

## Receipt `wo.run.v0`

Every started contract run attempts an atomic, mode-0600 receipt at `target/wo-runs/<run-id>.json` under `WO_ROOT` (or the repository root). The containing directory is mode 0700 and already ignored through `target/`.

Required fields are `schema_version`, `run_id`, `action`, `adapter`, `mode`, `started_at`, `finished_at`, `duration_ms`, `tool_versions`, `input_refs`, `artifacts`, `external_changes`, `status`, `exit_code`, and nullable `adapter_exit_code`. `external_changes` is always empty and artifacts are empty for this action. The checked-in producer schema is [`../../schemas/wo-run-v0.schema.json`](../../schemas/wo-run-v0.schema.json).

Receipts save only the safe action reference `gws:calendar.agenda:today`, observed tool version, process outcome, and timing evidence. They never save argv, stdin, stdout payload, child diagnostics, environment variables, tokens, account identifiers, Google bodies, user-home paths, or receipt paths. A receipt write failure converts apparent business success into exit 70 and suppresses stdout.

## Evidence boundary

`scripts/fixtures/gws/fake-gws.mjs` exercises the process contract completely offline. This evidence cannot validate real authentication, Google permission semantics, actual Calendar data, upstream command compatibility, or GWS `0.22.5`. If PR #338 merges, its GWS evidence must be reviewed against this result; this contract alone does not justify `ADOPT`.
