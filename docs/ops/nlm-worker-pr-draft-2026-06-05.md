# PR draft: NLM Worker MCP facade and runner smoke chain

## Title

```plain text
docs/worker: add controlled NLM Worker MCP facade
```

## Summary

This PR adds a controlled Cloudflare Worker MCP facade for future NotebookLM access, plus runner smoke scripts and deployment/rollback runbooks.

The intent is to keep NotebookLM as a library assistant / grain retriever, not a reasoning authority or writeback path.

## What changed

### Worker facade

- Add `workers/nlm-worker-mcp/worker.js`.
- Add `workers/nlm-worker-mcp/README.md`.
- Add `workers/nlm-worker-mcp/wrangler.toml.example`.

The Worker exposes only a small MCP surface:

- `nlm_manifest_list_packs`
- `nlm_manifest_get_pack`
- `nlm_check_pack_freshness`
- `nlm_list_notebooks`
- `nlm_list_sources`
- `nlm_query_grains`

Runner-backed tools fail closed until `NLM_RUNNER_URL` is explicitly configured.

### Runner / smoke helpers

- Add runner smoke helpers.
- Add HTTP runner scaffold.
- Add manifest dry-run helper.
- Add Worker local smoke helper.
- Add Wrangler health smoke helper.
- Add MCP HTTP smoke helper.
- Add bearer auth smoke helper.
- Add deploy preflight helper.

### Ops docs

- Add architecture note.
- Add runner smoke docs.
- Add manifest binding compatibility note.
- Add Worker smoke docs.
- Add deploy preflight doc.
- Add first deploy runbook.
- Add rollback runbook.
- Add branch review packet.

## Safety boundaries

This PR does **not**:

- deploy the Worker;
- configure Cloudflare secrets;
- configure `NLM_RUNNER_URL`;
- install or authenticate NotebookLM;
- store Google cookies;
- call NotebookLM;
- write back to GitHub from the Worker;
- expose source add/sync/delete/share/deep research/audio/video/slides.

## Expected first-stage behavior

Before runner is configured, `nlm_query_grains` should return:

```json
{
  "ok": false,
  "error": "runner_not_configured"
}
```

This is intentional and should be treated as a successful safety condition.

`/mcp` fails closed if `MCP_BEARER_TOKEN` is not configured, and returns 401 for missing or wrong bearer tokens once configured.

## Review checklist

- [ ] Worker only exposes the intended MCP tools.
- [ ] `/mcp` requires configured `MCP_BEARER_TOKEN` and fails closed when missing.
- [ ] Public pack output does not reveal raw notebook/source ids.
- [ ] Existing `upload.notebook_id_env + upload.source_id` manifest bindings are supported.
- [ ] Smoke helpers default to safe / no-op behavior.
- [ ] No runner or NotebookLM auth is configured in this PR.
- [ ] First deploy runbook keeps `NLM_RUNNER_URL` unset.
- [ ] Rollback runbook covers token rotation and route disablement.

## Suggested validation

```bash
bash scripts/nlm-worker-deploy-preflight.sh
node scripts/nlm-worker-manifest-dry-run.mjs list
node scripts/nlm-worker-local-smoke.mjs
```

In an environment with Wrangler:

```bash
MCP_BEARER_TOKEN="<local-token>" NLM_WRANGLER_SMOKE_MODE=dev bash scripts/nlm-worker-wrangler-smoke.sh
MCP_BEARER_TOKEN="<local-token>" bash scripts/nlm-worker-wrangler-mcp-smoke.sh
MCP_BEARER_TOKEN="<local-token>" bash scripts/nlm-worker-wrangler-auth-smoke.sh
```

## Next step after merge

Deploy facade only, without runner:

1. Set `MCP_BEARER_TOKEN`.
2. Optionally set read-only `GITHUB_TOKEN`.
3. Do **not** set `NLM_RUNNER_URL`.
4. Deploy Worker.
5. Validate `/healthz`, `tools/list`, manifest tools, and `runner_not_configured`.
