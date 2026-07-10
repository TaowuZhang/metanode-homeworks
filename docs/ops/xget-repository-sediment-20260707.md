# Xget 仓库沉积说明 · 2026-07-07

## 来源

- 来源仓库：`TaowuZhang/Xget`
- 目的：记录该仓库在 GitHub 清理中的定位，避免它继续作为个人长期事实母本。

## 当前观察

`Xget` README 将其定位为开发者资源加速引擎。`package.json` 显示它是一个可部署 / 可测试的 JavaScript/TypeScript 项目，包含 Cloudflare Worker 相关工具链、Vitest、ESLint、Prettier、TypeScript 与 Wrangler deploy/dev 脚本。

## 沉积判断

这不是普通笔记仓库，也不是 WQB / 量化 Alpha 资产。它更像外部工具 / 产品 / fork / reference surface。

因此：

- 不把完整实现复制进 `worang`；
- 不把它当作 `worang` 的设施层继续维护；
- 不在本轮删除；
- 交给本地 Codex 检查 fork/upstream/release/pages/external link 状态；
- 若没有必须保留的公开产品面或外部依赖，优先 archive 而非 delete。

## 可保留经验

- 如果某工具已经发展成独立可部署产品，清理时先判断它是否有 release、外部链接、fork/upstream 或用户可见表面。
- 产品型仓库不应被悄悄并入 `worang`，除非先拆出长期知识、接口合同和运行边界。
- `worang` 保留判断与索引，不直接承接完整产品实现。

## 对 GitHub 清理的影响

`TaowuZhang/Xget` 当前路线：`archive` 或 `migrate-then-archive`，但必须由本地 Codex 补证后决定。

仍需补证：

- 是否 fork；
- 是否有 upstream；
- 是否有 release；
- 是否有 Pages / Worker 部署面；
- 是否有 open issue / PR / external traffic；
- 是否有本地未推送内容；
- 是否有唯一资产需要先迁移。

## Retirement checklist · Round 4

Current evidence still shows `archived=false`, a configured homepage at `https://xuc.xi-xu.me`, HTTP 200 from that homepage, Cloudflare Pages/Workers workflows, and fork-only divergence.

Before archive:

1. inventory every live entry: homepage/domain, Cloudflare Pages, Cloudflare Workers, the `pages` branch, and any upstream-sync workflow;
2. export only non-secret configuration and a list of required environment-variable names; never copy values, tokens, headers, or cookies;
3. decide whether the 18 fork-only commits/six changed files are promoted, summarized, or explicitly historical-only;
4. migrate or stop the deployment and detach/switch the external domain;
5. verify the public endpoint is no longer served by this repository and confirm no required workflow remains;
6. establish local dirty/unpushed state or explicitly record that no local checkout exists;
7. only then archive once and read back `isArchived=true` through the GitHub API.

Until these checks pass, the repository remains `leave`; Round 4 performs no deployment or repository mutation.
