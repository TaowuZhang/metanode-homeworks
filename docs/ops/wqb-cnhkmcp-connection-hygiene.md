# WQB CNHKMCP connection hygiene

Updated: 2026-06-22T22:55:00+08:00

## Why this exists

The default Notion AI / personal-agent MCP surface currently contains several stale or duplicate CNHKMCP / WQB Cloudflare connections. Notion Support confirmed there is no documented UI path to remove custom MCP server connections one by one from the default Notion AI / personal agent.

This file is the run-level guardrail that prevents future WQB work from accidentally using stale CNHKMCP connections just because they remain visible in the tool surface.

## Support-email finding

Mailbox search found the thread:

- Subject: `Unable to remove duplicate MCP connections from the default Notion AI agent`
- First sent: 2026-06-04
- Notion Support response: for Custom Agents, remove via agent Settings -> Tools & Access -> Remove -> Save.
- Notion Support response for default Notion AI / personal agent: they were not able to find a documented UI path to remove custom MCP server connections one by one, nor documentation for a workaround; feedback would be shared with Product.
- User later drafted an escalation requesting backend cleanup of stale duplicate CNHKMCP entries.

## Current observed duplicate surface

Observed via Notion user connections on 2026-06-22:

- `mcpServer_cnhkmcp_wqb_cloudflare` — ready
- `mcpServer_cnhkmcp_wqb_cloudflare2` — needs_connection
- `mcpServer_cnhkmcp_wqb_cloudflare_v0_4` — needs_connection
- `mcpServer_cnhkmcp_wqb_cloudflare_v0_4_refresh` — ready
- `mcpServer_cnhkmcp_wqb_cloudflare_v0_4_refresh2` — needs_connection
- `mcpServer_cnhkmcp_wqb_cloudflare_yellow3` — ready
- `mcpServer_cnhkmcp_wqb_cloudflare_yellow2` — ready
- `mcpServer_cnhkmcp_wqb_cloudflare_yellow4` — ready
- `mcpServer_cnhkmcp_wqb_cloudflare_yellow` — needs_connection

## Rule

For WQB work, do not choose a CNHKMCP connection opportunistically from the visible tool list.

Use exactly one canonical CNHKMCP connection, recorded in `runs/wqb-current/cnhkmcp-connection-policy.json`.

If the user connects a new CNHKMCP / WQB Cloudflare MCP specifically for this workflow, update `canonical_connection_key` and `canonical_integration_url` there, then treat every older CNHKMCP connection as stale unless explicitly overridden.

## Temporary current state

Until a new canonical CNHKMCP connection is connected and recorded, the old duplicated surface is considered unsafe for new community/forum engineering. Existing historical references are preserved only for audit/history.

The next community/forum work should therefore proceed through local `cnhkmcp` introspection and sanitized artifacts rather than selecting a random visible Yellow connection.

## Operational checklist before any WQB MCP call

1. Read `runs/wqb-current/cnhkmcp-connection-policy.json`.
2. Confirm the requested tool call uses the recorded `canonical_connection_key`.
3. If there is no canonical key, do not use stale CNHKMCP connections for new forum/community work.
4. If a stale connection must be used for legacy read-only comparison, record the exception in run-state and keep it read-only.
5. Never use a `needs_connection` CNHKMCP entry.
6. Never infer that similarly named Yellow/refresh/v0.4 connections are equivalent.
