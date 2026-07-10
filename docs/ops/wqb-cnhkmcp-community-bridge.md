# WQB CNHKMCP Community Bridge

## Correction

The WQB community should not be treated as a public web source. Public web fetches can time out, hit cookie/recaptcha pages, or miss authenticated/community-only material. The correct route is to extend CNHKMCP locally, persist only sanitized community-derived artifacts to GitHub, then expose a controlled read-only Cloudflare MCP facade that Notion AI can call.

This replaces the previous mistaken approach of trying to crawl support.worldquantbrain.com directly.

## Target architecture

```text
Local CNHKMCP / browser-authenticated WQB session
  -> read community/forum/support pages through allowed authenticated tooling
  -> sanitize and extract claims locally
  -> write safe JSON/Markdown artifacts to GitHub
  -> Cloudflare Worker MCP reads GitHub artifacts only
  -> Notion AI calls Worker MCP read-only tools
  -> star-map scheduler consumes community priors
```

## Why GitHub is in the middle

GitHub is the durable, auditable sediment layer. It should store:

- source registry and fetch status;
- extracted community claims;
- failure-mode mappings;
- scheduler prior deltas;
- timestamps and source URLs;
- no cookies, no tokens, no WQB session material, no private alpha full text.

GitHub should not store:

- WQB credentials;
- cookies/session headers;
- private community raw HTML if access terms prohibit redistribution;
- private alpha expressions;
- submit-capable code paths.

## Local CNHKMCP extension contract

The local extension should provide read-only commands/tools like:

```text
community_list_posts(topic?, limit?, offset?)
community_get_post(postId | url)
community_search(query, topic?, limit?)
community_extract_claims(postId | url)
community_export_safe_pack(outDir)
```

Minimum safe output schema:

```json
{
  "source_id": "wqb-community-reduce-correlation",
  "url": "https://support.worldquantbrain.com/...",
  "access_route": "local_cnhkmcp_authenticated_readonly",
  "fetched_at": "2026-06-22T22:20:00+08:00",
  "title": "How do you reduce correlation of a good alpha",
  "visibility": "authenticated_community",
  "raw_stored": false,
  "claims": [
    {
      "claim": "...",
      "evidence_excerpt": "short sanitized excerpt if allowed",
      "maps_to_failure_modes": ["PROD_CORRELATION"],
      "scheduler_effect": "increase mechanism pivot slots after repeated prod correlation fail",
      "confidence": "low|medium|high"
    }
  ],
  "do_not": ["do not copy private expressions"],
  "derived_prior": {
    "topic": "correlation_reduction",
    "weight_delta": 0.05,
    "applies_to": ["BATCH-20260622-F-starmap-scout-and-climb"]
  }
}
```

## Cloudflare Worker MCP facade

Cloudflare should not hold WQB credentials and should not fetch WQB community directly in v1. It should read only GitHub-safe artifacts.

Tools exposed to Notion AI:

```text
wqb_community_health()
wqb_community_list_sources(topic?)
wqb_community_search(query, failureMode?)
wqb_community_get_prior(sourceId)
wqb_community_scheduler_deltas(batchId?)
```

Not exposed:

```text
submit_alpha
create_simulation
raw_cookie_fetch
credential_export
raw_private_html_dump
bulk_mining
```

## Integration with star-map scheduler

Before a new simulation batch is dispatched, the scheduler must check:

1. dominant blockers in the latest heatmap;
2. matching community priors from GitHub artifacts;
3. whether a local community fetch is stale;
4. whether the next batch has community-prior rationale per slot;
5. whether community advice says to stop a local micro-repair and pivot mechanism.

If community artifacts are missing for the dominant blocker, the batch can still be prepared but must be labeled:

```text
community_prior_status = missing_required_local_fetch
```

For Batch F, the current required community topics are:

- correlation reduction;
- avoiding overfitting;
- evaluating a new dataset;
- weight coverage / concentration;
- turnover / trade_when if event alphas are used.

## Local execution packet

The concrete local packet is stored at:

```text
runs/wqb-current/cnhkmcp-community-local-extension-request.json
```

User/local runner should implement or run the CNHKMCP extension and write safe extracted files into:

```text
runs/wqb-community/extracted/*.json
```

After that, run:

```bash
node scripts/wqb-community-intake.mjs \
  runs/wqb-current/community-source-registry.json \
  > runs/wqb-current/community-prior-summary.json
```

Then the Cloudflare Worker can expose the GitHub artifacts to Notion AI.

## Current state

- CNHKMCP Yellow currently exposed to this session has docs, dataset, field, alpha, leaderboard, and gate tools.
- It does not currently expose dedicated community/forum tools in the observed `listTools` result.
- GitHub search shows `worldquant-skill` integrates WorldQuant BRAIN / cnhkmcp and includes a `knowledge_base_search` skill for optimization methods and examples, but the external repo is not connected for full file loading in this Notion workspace.
- Therefore the correct next engineering step is local CNHKMCP extension + sanitized GitHub artifacts + Cloudflare MCP facade, not public web crawling.
