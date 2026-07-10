# Worang Index workflow handoff

`.github/workflows/worang-index.yml` 本轮由 GitHub 集成写入时被权限拦截：`403 Resource not accessible by integration`。

因此本页保留完整 workflow 交接稿。谷需要时可手动创建 `.github/workflows/worang-index.yml`。

## YAML

```yaml
name: Worang Index

on:
  pull_request:
    paths:
      - "**/*.md"
      - "**/*.json"
      - "scripts/**"
      - "docs/ops/worang-index/**"
      - "docs/ops/root-manifest.json"
  workflow_dispatch:
  schedule:
    - cron: "0 20 * * *" # Asia/Shanghai 04:00

permissions:
  contents: write
  pull-requests: write

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: "20"
      - name: Build index
        run: node scripts/build-worang-index.mjs
      - name: Check index
        run: node scripts/check-worang-index.mjs
      - name: Check root manifest
        run: node scripts/root-guard.mjs
      - name: Create maintenance PR
        if: github.event_name == 'schedule' || github.event_name == 'workflow_dispatch'
        uses: peter-evans/create-pull-request@v6
        with:
          branch: chore/worang-index-refresh
          delete-branch: true
          title: "chore(index): refresh worang index"
          commit-message: "chore(index): refresh worang index"
          body: |
            ## Summary
            - Refresh generated `docs/ops/worang-index/current.json`.
            - Refresh generated `docs/ops/worang-index/current.md`.
            - Keep root section concise and first-jump only.

            ## Verification
            - node scripts/build-worang-index.mjs
            - node scripts/check-worang-index.mjs
            - node scripts/root-guard.mjs
```

## 边界

- workflow 不直推 `main`。
- schedule / workflow_dispatch 只开维护 PR。
- PR 检查只验证生成物、根部精要与 root guard。
- 语义归位、根目录白名单、外部 URL、Notion 链接不自动改。
