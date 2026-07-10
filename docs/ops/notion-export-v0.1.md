# Notion Export v0.1

## Boundary

This is the canary exporter for the Notion AI to Worang/GitHub outbound path. It is intentionally narrow:

- One page per run.
- Manual CLI only.
- No webhook, launchd, attachments, batch concurrency, pull request creation, merge, deletion, or bidirectional sync.
- The only exported content path is `runs/notion-export/<export-id>/`.
- GitHub remains the versioned source of truth; Notion is the writing and dialogue surface.
- The Notion `状态` property must be a Select property. It is not a Notion Status property.

## Fixed Contract

- Contract: `notion-export/v0.1`
- Notion API version: `2026-03-11`
- Namespace UUID: `7b4027da-f653-5359-85fa-dc238274f6fd`

The namespace UUID is permanent for this protocol version. It was generated from:

```text
https://github.com/dongxi-heji/worang/notion-export/v0.1
```

Export identity:

```text
payload_hash = SHA-256(canonical_payload)
export_id = UUIDv5(7b4027da-f653-5359-85fa-dc238274f6fd, page_id + ":" + payload_hash)
```

## Canonical Payload

The payload is UTF-8 canonical JSON with fixed key order:

```json
{
  "contract": "notion-export/v0.1",
  "notion_version": "2026-03-11",
  "namespace_uuid": "7b4027da-f653-5359-85fa-dc238274f6fd",
  "page_id": "<page-id>",
  "title": "<title>",
  "markdown": "<UTF-8 LF normalized markdown>"
}
```

Markdown is normalized to LF and exactly one trailing newline. The payload contains no timestamp, media URL, credential, or runtime host data.

## Validation

The Notion Markdown response is fail-closed. The response body must be an object with explicit fields:

- `markdown`: string
- `truncated`: boolean
- `unknown_block_ids`: array

Missing fields, wrong field types, non-object responses, or array responses fail at `fetch:invalid_markdown_response`. `truncated: true` still fails at `fetch:markdown_truncated`, and non-empty `unknown_block_ids` still fails at `fetch:unknown_block_ids`. Empty or whitespace-only Markdown fails at `validate:empty_markdown`.

The Markdown validator permits ordinary text, blank lines, headings 1-4, paragraphs, ordered and unordered lists, checkboxes, blockquotes, fenced code blocks, horizontal rules, normal Markdown links, emphasis, strikethrough, inline code, emoji, CJK text, and backslash escapes.

It rejects unrecognized HTML-like tags, attribute suffixes, tables, unsupported enhanced structures that appear as concrete syntax, and Markdown image/media syntax such as `![alt](url)`, `![alt][ref]`, and `![alt]`. Fenced code block content is not scanned for tags, attributes, or image/media examples. Escaped image syntax such as `\![alt](url)` is treated as ordinary displayed text.

Reports never copy the full Markdown body. They only keep status, stage, error code, and position.

## Git Sink

The sink uses argument-array subprocess calls only. It never uses `sh -c`, `eval`, string-built shell commands, `git add -A`, or force push.

Flow:

```text
detached worktree from base
write runs/notion-export/<export-id>/
git diff --check
status boundary check
single detached commit
push HEAD directly to refs/heads/notion/export/<export-id>
ls-remote SHA verification
remove worktree
```

The sink does not create a persistent local `notion/export/<export-id>` branch. The export branch exists only on the remote after a successful push.

If the remote branch already exists, the exporter reads the remote manifest by querying the branch with `git ls-remote --exit-code`, fetching that head into a temporary local ref under `refs/worang/notion-export-check/`, and reading `runs/notion-export/<export-id>/manifest.json` with `git show <temp-ref>:<path>`. The temporary ref is deleted after success or failure.

Matching page ID, export ID, and payload hash produce an ack-only or already-exported recovery path. The Notion receipt branch is compared to the deterministic branch name, and the receipt commit is compared to the remote branch head SHA; branch/commit are not required in the manifest. Missing branches, fetch failures, missing manifest paths, invalid manifest JSON, manifest conflicts, and receipt commit conflicts are distinct failures.

## CLI

```bash
node scripts/notion-export.mjs doctor
node scripts/notion-export.mjs plan
node scripts/notion-export.mjs dry-run --page <page-id>
node scripts/notion-export.mjs run --page <page-id>
node scripts/notion-export.mjs resume --page <page-id>
```

Live Notion export requires `NOTION_TOKEN` and `NOTION_DATA_SOURCE_ID`. Missing credentials are a real external block, not a simulated success.

`dry-run` writes the export package to a temporary directory and prints that path. It must not leave new `runs/notion-export/<export-id>/` files in the implementation branch.
