# NLM 运行环境边界

> 状态：草案 / 内部试运行  
> 日期：2026-06-01  
> 目的：回答 NLM / NotebookLM 远手到底应该跑在哪里，以及什么时候才需要 VPS。

## 一句话

NLM 的重点不是“有一台服务器”，而是克能不能在需要时用上一只远手。

运行环境按债务从轻到重排序：

1. 当前克可用的 GitHub MCP / GitHub 文档与 PR 流
2. 本地手动 CLI
3. GitHub Actions 短跑
4. Cloudflare Worker / Tunnel 作为轻量入口
5. Notion Worker（观察项）
6. VPS 长期后端

VPS 是最后的身体，不是第一选择。

---

## 已拆除的旧路

终端桥已经不存在。

不要再把未来方案建立在这些假设上：

- Notion 页面投递 shell 命令
- 本地 watcher 轮询 Notion
- 本地常开电脑回填结果
- 自动粘贴“继续”唤醒克
- 终端槽位 / 车道 / 本地桥作为执行层

相关 legacy 文档只作历史参考，不是现行路径。

---

## 当前真正可用的手

### 1. 克能直接用的：GitHub MCP

当前克能用 GitHub MCP 对 `dongxi-heji/worang` 做这些事：

- 读文件、目录、issue、PR
- 新建分支
- 创建 / 更新文件
- 提 PR
- 评论 issue / PR
- 做轻量 repository 操作

因此，凡是“沉到沃壤”的判断、文档、路线、实验记录，优先走 GitHub MCP。

这条路的优点：

- 克已经能用
- 不需要终端桥
- 不需要 VPS
- 变更可审计
- 可通过 PR 保留边界

这条路的限制：

- 它能改仓库，不等于能运行任意 CLI
- 它适合沉积和编排，不适合实时执行 NotebookLM 查询

### 2. 谷能用的：本地 CLI

如果 `nlm` 只需要偶尔用，最轻路径仍是谷本地执行，再把结果带回。

适合：

- 验证输出质量
- 检查 grain 是否干净
- 检查引用是否可用
- 低频查询

不适合：

- 克在 Notion 对话里实时调用
- 自动定时任务
- 多入口共享登录态

---

## GitHub Actions：短跑，不是常驻后端

GitHub Actions 适合把 NLM 相关动作变成可复现短跑。

适合：

- `scripts/nlm-grain-probe.sh` sanity check
- source pack stale check
- 定时构建 `wo/nlm/*.md`
- workflow_dispatch 手动触发一次查询
- 把结果写成 artifact、issue comment 或 commit

不适合：

- 低延迟实时 MCP
- 长期保管复杂 Google 登录态
- 需要交互登录的 CLI
- 需要随时被克同步调用的工具

判断：

> GitHub Actions 是任务工坊，不是实时远手。

---

## Cloudflare：优先试的轻量入口

如果要花钱或上外部基础设施，Cloudflare 应优先于 VPS 被试。

Cloudflare 适合：

- 提供稳定 HTTPS 入口
- 做简单鉴权、限流、日志
- 转接 GitHub workflow_dispatch
- 暴露轻量 HTTP / MCP facade
- 保护源站或隐藏复杂后端

Cloudflare 不一定适合：

- 直接跑普通 CLI
- 持久保存复杂登录态
- 长任务队列
- 浏览器自动化
- 需要本地文件系统语义的工具

判断：

> Cloudflare 是门房和轻入口，不默认是发动机。

若某个 NLM 动作最终可以拆成“HTTP 请求 → 外部 API → 短响应”，Cloudflare 可能就够。若必须跑 `nlm` CLI 并保留登录态，再看后续层。

---

## Notion Worker：观察项，不设为默认

Notion Worker 可能是更新的 Notion 原生工具入口，但当前不设为默认路线。

原因：

- 谷几乎不使用 Notion Custom Agent
- 可用性、权限、费用、限制仍需观察
- 即便可用，也未必适合跑复杂 CLI 或长期登录态

如果未来它能让克直接调用轻量 tool，并且成本/权限可接受，再评估。

判断：

> Notion Worker 是可能的新器官，但不是当前承重柱。

---

## VPS：什么时候才成立

只有当以下条件同时命中多条，才进入 VPS：

- 克必须在 Notion 对话中实时调用 NLM
- 现有 GitHub MCP / GitHub Actions / Cloudflare 都无法承接
- 工具只能作为普通 CLI 跑
- 需要长期保存登录态、缓存或任务状态
- 需要队列、重试、日志、超时控制
- 需要多个入口共享同一个远手
- 单次任务较长，超过 serverless / Actions 的舒适区

否则不买 VPS。

VPS 的合理身份是：

> 一只长期在线、能跑普通 CLI、能保管状态、能被克远程叫到的手。

不是：

> 为了有后端而有后端。

---

## 当前路线

### 现在

1. 把运行环境边界沉到 GitHub。
2. 继续维护 `wo/nlm/` 的远手语义和 source pack。
3. 不恢复终端桥。
4. 不把 VPS 作为默认答案。

### 下一步候选实验

按轻重排序：

1. GitHub Actions：跑一个不含机密的 NLM 脚本 sanity check。
2. GitHub Actions：验证 source pack stale check。
3. Cloudflare Worker：做一个最小 `/ping` 或 `/grain-proxy` facade。
4. Cloudflare Worker：转接 GitHub workflow_dispatch，观察能否让克通过 GitHub 结果间接用上。
5. 若以上都不足，再讨论 VPS。

---

## 判断表

| 需求 | 首选运行环境 | 备注 |
|---|---|---|
| 克沉积文档 / 改路线 | GitHub MCP | 当前已可用 |
| 谷低频查询 NLM | 本地 CLI | 零后端 |
| 定时构建 source pack | GitHub Actions | 短跑 |
| 检查 source pack 是否 stale | GitHub Actions | 可写 artifact / issue |
| 提供轻量 HTTPS 入口 | Cloudflare Worker | 先试 |
| Notion Agent 原生工具 | Notion Worker / Custom MCP | 观察项 |
| 实时跑 CLI + 保持登录态 | VPS | 最后才上 |

---

## 不做

- 不恢复终端桥。
- 不让本地常开电脑成为默认路径。
- 不为了一次实验购买长期 VPS。
- 不把 NotebookLM 当主脑。
- 不让 NLM 直接写 `taste / dream / reveal / trace`。
- 不把 Cloudflare 误认为能解决所有 CLI 执行问题。

---

## 归位

- NLM 原则与 source pack：`wo/nlm/`
- 运行环境边界：`docs/ops/nlm-runtime-boundary.md`
- 具体 workflow：未来放 `.github/workflows/`
- 已废弃终端桥：仅保留在 `docs/legacy/`

一句话：

> 远手先要能被克用上；身体越晚出现越好。
