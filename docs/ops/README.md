# docs/ops · 机器运行、审计与外显操作入口

`docs/ops/` 承接沃壤里的机器运行层、审计证据层、发布 / 展示 / 分发操作层、连接性与对账层。

它不是内容库，也不是长期判断库。它只回答：机器怎么跑、入口怎么验、发布怎么走、失败怎么记、workflow 如何交接，以及这些操作如何不反向污染 `成为/`、`领域/`、`资源/`、`技能/`、`wo/` 与公开展示面。

长期判断、触发语、关系规则和 AI 的进场方式，按内容回 `破土.md`、`成为/`、`领域/`、`资源/`、`技能/` 或 `wo/`。Notion 只保留公开窗口，不再是协议、判断或技能的母本。

## 先读哪一页

| 要处理什么 | 第一跳 |
|---|---|
| 冷启动索引 / AI 第一跳 | `worang-index/current.md` |
| 根目录机器账 | `root-manifest.json` 与 `../../scripts/root-guard.mjs` |
| 机器输出、非交互调用、外部子进程的共同底线 | `wo-runtime-v0.md` |
| `wo` 当前 15 个 action 的事实盘点与离线基线检查 | `wo-action-inventory.md` |
| 新建文件夹打报告 | `new-folder-report-template.md` |
| PR 如何只递交必要的人类注意力 | `../../.github/pull_request_template.md` |
| 通用 CLI 当前状态 / 证据门 | `cli-observatory.md` |
| workflow 的 last attempt / last healthy 与修复复验 | `workflow-evidence-loop.md` |
| 浏览器书签清理与可见界面核验 | `browser-bookmark-cleanup-2026-07-05.md` |
| 浏览器旧链接第二轮治理过程 | `browser-link-governance-round-2-2026-07-05.md` |
| 浏览器链接居留与路径对账 | `browser-link-placement-audit-2026-07-05.md` |
| 分发总图 / 出口边界 | `distribution-map.md` |
| 外显发布通道铺路 | `distribution-road-v0.md` |
| 外部 CLI / 开源仓候选 | `distribution-cli-registry.v0.md` |
| 平台 CLI 选型证据与接入状态 | `platform-cli-selection-2026-06-28.md` |
| 发布候选放行 | `distribution-publish-candidate-gate.md` |
| 发布包字段与检查 | `publishing-package.md` |
| 发布失败 / 风控 / 限流 | `publish-failure-ledger.md` |
| Workflow 待安装记录 | `publish-workflow-install-note.md` |
| NLM / NotebookLM 运行环境边界 | `nlm-runtime-boundary.md` |
| NLM source 架构、灵感 corpus 与本地验证 | `nlm-cli-integration-2026-06-27.md` |
| NLM 灵感显影首轮动态验证记录 | `nlm-ideas-visibility-local-validation-2026-06-27.md` |
| GWS 只读 adapter 的 no-input、timeout、exit 与 receipt 契约 | `gws-readonly-adapter-contract.md` |
| 灵感 corpus 合同 | `../../wo/nlm/corpus/becoming-ideas/README.md` |
| 灵感跨卡价值评测 | `../../wo/nlm/ideas-eval.md` |
| Notion 产品门与公开窗口 | `notion-product-gate.md` |
| Notion 清场记录 | `notion-cleanup-2026-07-02.md` |
| 个人展示 / Web 外壳 | `personal-display-and-web.md` |
| Web 静态数据 | `web-static-data.md` |
| Web 第零期试运行 | `web-issue-zero-brief.md` |
| Telegram / 外排管线 | `excrete-telegram.md` 与 `distribution-map.md` |
| GetNote mirror / MCP 读入口 | `getnote-mirror-mcp-worker.md` |
| GetNote 好问层 / Resource Registry 接口 | `getnote-mirror-v0.2-ask-layer.md` |

## 子层边界

| 子层 | 放什么 | 不放什么 |
|---|---|---|
| 机器运行 | workflow、guard、dry-run、检查脚本说明、生成器说明 | 个人判断正文 |
| 审计证据 | 一次性审计、迁移记录、对账报告、失败 ledger、真实运行证据口径 | 可长期复用的领域知识、逐 PR 质量打分 |
| 发布 / 外显 | Web、分发、公共入口、发布包、平台适配 | 未筛隐私的原始色散 |
| 连接 / 对账 | MCP、worker、GetNote mirror、资源候选闸、Project 对账 | 活的第二事实源或应用内工作台镜像 |
| 索引维护 | root manifest、worang-index、冷启动指针 | 全仓目录复刻 |

## 与 Notion 公开窗口的关系

Notion Education Plus 只保留一个薄公开窗口：个人介绍、精选作品、公开链接和轻量交互陈列。其文字、结构和 HTML 源码先在 GitHub 留母本，再手工陈列到 Notion。

本目录保存具体操作地层：产品门、清场记录、发布管线、Web 结构、静态数据、发布包、平台适配、试运行记录、机器账与审计证据。

一句话：

> Notion 是窗；GitHub 是母本；`docs/ops/` 负责修路、验路、过路和记事故。

## 与三立基及其他承重层的关系

- `成为/`：谷正在成为的运行态、身体、情绪、项目、日用结构。
- `领域/`：谷正在训练和持有的能力域、作品域、责任域。
- `资源/`：外部材料、课程、影音、链接、跟练、GetNote / Douban / BibiGPT 等已沉积资源。
- `技能/`：可复用的路径、协议、QA 与 QC。
- `wo/`：AI 协作设计、本地动作接口、观察台与 source pack。
- `docs/ops/`：服务这些层的机器、审计、发布、连接与对账。

如果一个文件主要是在回答“我是谁 / 我怎么长 / 我会什么 / 我取养什么 / 我怎样与 AI 协作”，不要放在 `docs/ops/`。

## 不做什么

- 不把网页做成第二套知识库。
- 不把社交反馈自动反向写回沃壤结构。
- 不保存未筛隐私的原始色散。
- 不替谷执行发布、署名或不可逆公共动作。
- 不复刻 Notion 内部施工图。
- 不保存长期判断正文。
- 不把 `docs/ideas/` 扩成想法库。
