# 分发管线地图

这份文档给「排泄 / 发布 / 展示」划边界。它不是内容计划，也不是平台运营 SOP。它只回答：一坨从沃壤长出的东西，出土以后走哪根管、谁是真相源、是否允许回写。

## 总判准

- **本地沃壤**：生长与沉积的主源。
- **Notion**：当前工作台、活体展示入口、AI 协作调度面。
- **网页**：外部门面，不成为第二套内容系统。
- **社交媒体**：排泄口 / 外部反馈来源，不反向决定沃壤结构。
- **Git**：可迁徙凭证，记录已经沉地的判断与工具。

一句话：

> 本地负责生长，Notion 负责活体展示，网页负责门面，社交媒体负责排泄。

## 出口表

| 出口 | 作用 | 上游源 | 生成物 | 是否外发 | 是否回写 |
|---|---|---|---|---|---|
| Telegram | payload / 受控通道烟测 | `type=excrete` event | 消息 payload；live adapter 待显式接通 | 默认否 | summary / artifact，不回写土壤 |
| 小红书 | 国内图文分发试验口 | 发布包 + Playwright 外手 | 图文笔记 / 截图 / 失败快照 | 是 | 只写外手日志与失败快照 |
| 公众号 | 长文母版口 | 发布包 / 本地 Markdown | 图文草稿 / 长文 | 半自动 | 不回写土壤 |
| Notion 公共页 | 活体自我介绍 / 展示入口 | Notion 页面 + 本地已沉淀摘要 | 展示页、数据库视图、脉搏入口 | 是 | Notion 自身是展示层，不当本地真相源 |
| Web | 外部门面 / 可控视觉层 | 本地 Markdown / JSON / Notion 摘要 | 静态站 | 是 | 不回写土壤 |
| GitHub | 工程落痕 / 可迁徙凭证 | 本地代码与文档 | commit / PR | 是 | Git 是沉积历史，不是社媒反馈口 |
| CFnew | 连接备用供应链 / 分线验收 | 上游 CFnew + 本地 secrets | Worker、KV、品牌化入口、smoke 结果 | 只给授权对象 | 只回写 runbook、模板、脚本，不回写秘密 |

## 现有最小安全通道

当前已经存在：

- `wo excrete --json`
- `wo excrete last --platforms xiaohongshu,x-twitter,linkedin,maimai --mode dry_run --prepare --json`
- `node scripts/wo-excrete.mjs`（兼容分步入口）
- `xurl` / Postiz / `sau` 固定版本外手（见 `platform-cli-selection-2026-06-28.md`）
- `scripts/lib/pipes/telegram.mjs`
- `scripts/lib/pipes/xiaohongshu.mjs`
- `scripts/publish/adapters/*`
- `.excrete-cursor`

边界：

- `wo excrete` 追加带 `dry_run / last_click / live` 模式的排泄标记；`--prepare` 可同步生成托盘。
- Node 外手扫描未处理标记，生成统一发布包、平台草稿与 `summary.json`。
- 旧 Telegram / 小红书直接投递 pipe 保留为历史实现，但已从 runner 断开，不会被默认或 `--prepare` 调用。
- 不把平台成功 / 失败状态反写 `成为/消化/events.jsonl`。

## 发布包位置

平台适配不应直接吃原始事件。中间要有可检查的发布包：

- `ai-browse/out/packages/<event-id>.md`
- `ai-browse/out/packages/<event-id>.json`

发布包是「出土前的托盘」：人能看，外手能吃，失败能定位。

最小字段见 `docs/ops/publishing-package.md`。

## 分发节奏

### 第一阶段：真通道

目标：确认一根外排管确实能跑。

- Telegram 作为最小验证口。
- 只要求能送达、能记 log、失败能看见。
- 不做增长，不看点赞，不分析反馈。

### 第二阶段：发布包

目标：把原始事件转成平台可吃的成品托盘。

- 生成标题、正文、摘要、标签、平台备注。
- 做隐私 / 版权 / 外发风险检查。
- 先只生成包，不自动发布到复杂平台。

### 第三阶段：平台适配

目标：让不同平台从同一发布包取材。

- 小红书：图文包 / 封面字 / 标签 / Playwright 半自动。
- 公众号：长文母版 / 草稿化。
- Telegram：轻量文本 / smoke test。
- Web / Notion：展示层抽取，而不是排泄层回流。

## 连接供应链边界

CFnew 这类连接供应链不是社交媒体外排，也不是公开发布管线。

它归 `docs/ops/`，因为它管的是：

- 可迁徙部署；
- 依赖记录；
- 分线管理；
- 失败回滚；
- 个人 / 授权对象的可达性验收。

边界见：

- `docs/ops/cfnew-supply-chain.md`
- `docs/ops/cfnew-connectivity-checks.md`

它不允许把真实订阅、UUID、KV、Cloudflare token、账号后台地址写进 Git。

## 社交媒体边界

可以：

- 把已经出土的 residue / spark / fossil 排出去。
- 记录投递结果、失败原因、payload hash。
- 保存失败快照，供下次修管子。

不可以：

- 用点赞、阅读量、评论反向改沃壤结构。
- 把平台状态写回 `events.jsonl`。
- 为迎合平台重写本地真相源。
- 自动外发含隐私、账号、住址、token、未授权图片或真人信息的内容。
- 在没有明确平台边界时，把「生成发布包」和「真正发布」混成一步。

## Notion 与网页展示边界

Notion 公共页和网页不是排泄管。

- Notion 公共页：适合「脉搏 + 工作台 + 陈列柜」的活体入口。
- 网页：适合更可控的视觉外壳。
- 两者都应从本地 / Notion 已沉淀材料抽取，不要复制出第二套维护宇宙。

个人展示的核心分叉：

- 内部管理视图：完整、可编辑、字段多。
- 外部展示视图：Teaser → Depth，先给卡片和好奇缺口，再进入深层。

## 下一步

- [ ] 给 Telegram 重接显式 live adapter、chat 白名单、receipt 与幂等策略；未完成前只生成 payload。
- [x] 发布包生成器已接入 `wo excrete --prepare` / `node scripts/wo-excrete.mjs`，并统一到 canonical schema。
- [ ] 为小红书先做可人工检查的图文包输出格式，之后再做 Playwright。
- [ ] 为公众号生成长文草稿模板，不急自动发。
- [ ] CFnew 供应链只保留 runbook / env 模板 / smoke 脚本，不提交真实秘密。
