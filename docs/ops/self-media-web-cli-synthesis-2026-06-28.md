# 自媒体 × 个人网站 × 外排 CLI 汇总与推进规格 · 2026-06-28

## 目标

把禾集、自我网站、Notion 栈桥、GitHub 发布路和 `wo excrete` 收束为一条可复用管线：内容主源保持不动，机器生成统一发布包，各平台只消费适配稿；默认不发生公共发布。

## 当前事实

- 禾集已有品牌、四桥内容体系、创作管线和候选平台；0→1 仍缺一次完整首发链路。
- 个人网站已有第零期静态壳、外部信息架构和 `dispatch.json` 空位；网站是长期门面，不是社交排泄口。
- Notion 栈桥/棱窗负责外显方向、陈列与放行，不做发布器；具体脚本、schema、失败记录留在 `docs/ops/` 与 Git。
- PR #291 已落一套安全的发布路：统一 schema、目标配置、dry-run 适配器和结果汇总。
- 旧 `wo excrete → scripts/wo-excrete.mjs` 另有一套不兼容发布包，并会直接调用 Telegram / 小红书外手。

## 本轮实现规格

### 命令

```bash
# 只读查看候选，供人或 AI 稳定解析
cargo run -- excrete --json

# 留下排泄标记；默认只 dry-run
cargo run -- excrete last --platforms web,rss,xiaohongshu --mode dry_run --json

# 同步生成统一发布包和平台草稿，不外发
cargo run -- excrete last --platforms web,rss,xiaohongshu --mode dry_run --prepare --json
```

### 契约

- 唯一发布包契约：`schemas/publish-package.schema.json`。
- `wo excrete` 默认 `dry_run`；`last_click` 只生成最后一步材料；`live` 仍受目标白名单限制。
- `--json` 的 stdout 只能有一个可解析 JSON 对象。
- `--prepare` 调用本地 Node 外手，但外手只生成、校验和运行受闸门控制的适配器。
- 不保存 token、cookie、账号密码；不把平台结果写回内容主源或 `events.jsonl`。

### 验收

- Rust 测试覆盖模式序列化与 JSON 候选形状。
- Node 测试证明旧入口生成的新包通过统一 validator。
- 临时 `WO_ROOT` 端到端试跑能生成 package、平台草稿和 summary，且不访问外部平台。
- `cargo test`、Node tests、发布路 smoke 全部通过。

## 平台优先级（谷已确认）

- 第一批：小红书、X。
- 第二批：LinkedIn、脉脉。
- 国内全表：微信公众号、视频号、小红书、知乎、B站、抖音、脉脉、豆瓣。
- 国外全表：X、YouTube、TikTok、LinkedIn。
- 知识网络/知识付费：小报童（按用户口述“小爆同”理解）、小鹅通；国外重点评估 Ghost，另保留 Substack。

具体 CLI 选型与安装记录见 `platform-cli-selection-2026-06-28.md`。

## 可以反复使用的资产

| 资产 | 复用方式 | 自动化位置 |
|---|---|---|
| 禾集“四桥 × 表达姿势”与内容创作公式 | 每个选题都可生成母稿 metadata，不按平台重造选题 | 后续放进 package 的内容标签与模板层 |
| `结撰 → 编校 → 赋形` 工序 | 所有图文、长文、视频脚本共享前半程 | 先产母稿与审查结果，平台 adapter 只做裁剪/排版 |
| canonical publish package | 人检查、脚本消费、CI 校验、失败定位共用一个托盘 | `wo excrete --prepare` + schema + validator |
| `dry_run / last_click / live` 三闸 | 所有平台共用安全语义 | CLI、target config、GitHub Actions 三层同时限制 |
| Web 的 Pulse / Workbench / Cabinet / Dispatch | 个人网站可持续更新，不复制整座知识库 | 静态 JSON 构建；Dispatch 消费已放行发布包 |
| risk notes + human check | 隐私、版权、身份、公共反作用不随平台丢失 | package 必填字段；live 也不能移除人工确认 |
| summary / artifact / callback sink | 每次执行可回看、可重跑、可定位 | `publish-runs/<event-id>/summary.json`，不写回主源 |

## 自动化边界

### 可以立即自动化

- 从 digest 选择材料，输出稳定 JSON 候选。
- 生成母标题、正文、摘要、平台变体骨架与风险检查位。
- 校验 schema、目标平台、模式与必填人工检查。
- 为 Web、RSS、Telegram、小红书等生成草稿 artifact。
- 在 PR / 手动 workflow 中重跑 dry-run 并上传结果。
- 收集 adapter 成败、日志路径和下一步动作。

### 必须保留人工判断

- 选题是否值得公开、是否代表谷当下判断。
- 引文版权、真人信息、署名与身份暴露。
- 平台最终标题、封面、视觉与发送时点。
- 高风控国内平台的最后点击。
- 任何 `live` 开闸，以及失败后的重试/换平台决定。
- 外部反馈如何回收到栈桥/棱窗；它不能自动改写沃壤结构。

## Issues / PR 收束

- [Issue #208](https://github.com/dongxi-heji/worang/issues/208) 的真正缺口不是再做品牌规划，而是跑通一次“选题 → 结撰 → 编校 → 赋形 → 发布 → 复盘”。本轮补的是发布托盘与机器路，尚未代替首篇内容和账号注册。
- [PR #225](https://github.com/dongxi-heji/worang/pull/225) 确认 Web 是“在建慢炖”的长期门面，`dispatch.json` 是分发位；因此 Web adapter 应消费发布包，不应让网站变成第二内容库。
- [PR #291](https://github.com/dongxi-heji/worang/pull/291) 已给出正确的统一发布契约与安全闸；本轮把旧 `wo excrete` 接到这条路，并消除第二套包格式。
- [Issue #293](https://github.com/dongxi-heji/worang/issues/293) 的“验证 main + 安装专用 workflows”已在本轮完成本地实现；真正的平台 live wiring 仍等待平台优先级与账号/官方 API 条件。
- [Issue #315](https://github.com/dongxi-heji/worang/issues/315) 要求稳定 JSON、非交互和 receipt。本轮先在 `excrete` 完成首个垂直切片；全局 `--json / --plain / --schema`、统一退出码与 plan/apply 仍是后续通用化工作。

## 下一刀

1. 由谷完成一次 `xurl` OAuth 与 Postiz device login；凭证不进入 Agent 上下文。
2. 选一篇现有作品或新材料，实际走完禾集 0→1；不要再用演示文案替代真实内容。
3. 给小红书发布包补真实图片资产，让 `sau-plan` 形成完整 argv。
4. 先实跑 X 的 plan/apply；LinkedIn 只创建 remote draft；小红书由谷确认 `sau` plan 后单独放行。
5. Ghost 作为个人知识网络/知识付费主站候选另开部署实验。

## 证据

- GitHub Issue #208：禾集 0→1。
- GitHub Issue #293 / PR #291：publish road v0 与后续闸门。
- GitHub Issue #315：AI 稳定 CLI 机器契约。
- GitHub PR #225：个人网站第零期与 `dispatch.json` 空位。
- Notion 栈桥审计记录：`docs/ops/zhanqiao-docsops-audit-2026-05-31.md`（原页：<https://www.notion.so/997c8a52e08249508ed7ff444516570e>）。
- `领域/作品/彼岸此地·集/推·自媒体.md`、`领域/作品/建·个人页面.md`。
