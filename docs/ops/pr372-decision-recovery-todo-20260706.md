# PR #372 判断链追回待办清单（2026-07-06）

## 0. 本清单的任务边界

本清单不是重新审核 TopHub / Folo / 浏览器 / 沃壤，也不是给未来另起一套模板。

本清单只做一件事：

> 从 PR #372 已有文件和此前 ChatGPT 对话中，把已经发生的逐源判断追回来，整理成可以合并前核验的判断链。

判断链指：

- 为什么这个来源进入 TopHub / Folo / 沃壤 / 浏览器 / 按需；
- 为什么不是另一个层；
- 为什么只选这个，不选同页或同类其他来源；
- 为什么放在这个分组或这个顺序；
- 为什么只是候选、复查、按需或退出；
- 该判断现在在 PR 哪个文件里，是否已经足够清楚。

禁止把以下内容冒充判断链：

- 只列最终数量；
- 只列文件层级；
- 只写“停更、重复、噪声、未读债务”这类大词；
- 重新编写当前模型觉得合理的理由；
- 大量使用 `not_found` 逃避追索。

---

## 1. 工作对象

PR：`dongxi-heji/worang#372`

当前已核验状态：

- PR 状态：Draft；
- changed files：208；
- commits：490；
- 当前 head：`feat/folo-subscription-system-v0-20260703`；
- 当前真实 TopHub 结果文件、Folo 账本、目录复审文件均在本 PR 中。

---

## 2. 文件分组清单

### A. 真实状态账本

这些文件记录当前真实状态，不负责完整解释所有过程，但必须作为 source_id 抽取入口。

- `资源/雷达/TopHub 全部订阅顺序.md`
- `资源/雷达/TopHub 使用配置.md`
- `资源/雷达/订阅源账本.yml`
- `资源/雷达/Folo 分类重组.md`
- `资源/雷达/订阅系统.md`
- `资源/雷达/信息流最终分工.md`
- `资源/雷达/文化内容入口与影音工作流.md`

待办：

- [ ] T-A1：从 `TopHub 全部订阅顺序.md` 抽取 80 个当前 TopHub 节点。
- [ ] T-A2：从 `订阅源账本.yml` 抽取 27 个当前 Folo 来源。
- [ ] T-A3：从 `Folo 分类重组.md` 抽取五类注意力分类。
- [ ] T-A4：从 `TopHub 使用配置.md` 抽取精确追踪器、关闭项、已执行替换与取消项。
- [ ] T-A5：生成 `source_id -> 当前层 -> 当前状态 -> 当前所在文件` 的初版索引。

### B. 最终收口文件

这些文件优先提取逐源理由，因为它们最接近“当时为什么这么选”。

- `docs/ops/tophub-ai-final-closure-20260705.md`
- `docs/ops/tophub-design-final-closure-20260705.md`
- `docs/ops/tophub-shopping-final-closure-20260705.md`
- `docs/ops/tophub-finance-final-closure-20260705.md`
- `docs/ops/tophub-entertainment-final-reconciliation-20260705.md`
- `docs/ops/tophub-newspapers-final-closure-20260705.md`
- `docs/ops/tophub-government-affairs-snapshot-closure-20260706.md`
- `docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md`
- `docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md`

待办：

- [ ] T-B1：逐个文件抽取“已进入真实账本”的来源理由。
- [ ] T-B2：逐个文件抽取“复查候选”的来源理由。
- [ ] T-B3：逐个文件抽取“按需 / 沃壤 / 专业入口”的来源理由。
- [ ] T-B4：逐个文件抽取“明确不选 / 退出 / 不订阅”的来源理由。
- [ ] T-B5：把“拒绝的替代项”与“保留项”成对记录，不把被拒绝项丢掉。

### C. 逐页账本与 reassessment 文件

这些文件用于补最终收口文件没有写全的逐页过程。

#### AI

- `docs/ops/tophub-ai-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-01-03-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-04-06-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-07-09-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-10-12-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-13-15-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-16-18-20260705.md`

#### 开发

- `docs/ops/tophub-development-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-01-03-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-04-06-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-07-09-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-10-12-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-13-15-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-16-18-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-19-21-20260705.md`
- `docs/ops/tophub-development-reassessment-pages-22-24-20260705.md`
- `docs/ops/tophub-development-page-25-final-closure-20260705.md`

#### 财经

- `docs/ops/tophub-finance-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md`

#### 购物

- `docs/ops/tophub-shopping-directory-ledger-20260705.md`
- `docs/ops/tophub-shopping-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-jd-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-smzdm-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-deals-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-credit-cards-reassessment-page-01-20260705.md`

#### 报刊 / 娱乐 / 社区 / 专栏

- `docs/ops/tophub-newspapers-directory-ledger-20260705.md`
- `docs/ops/tophub-entertainment-*-closure-20260705.md`
- `docs/ops/tophub-community-*-closure-20260705.md`
- `docs/ops/tophub-columns-pages-01-10-20260706.md`
- `docs/ops/tophub-columns-pages-11-20-20260706.md`
- `docs/ops/tophub-columns-pages-21-30-20260706.md`
- `docs/ops/tophub-columns-pages-31-40-20260706.md`
- `docs/ops/tophub-columns-pages-41-50-20260706.md`
- `docs/ops/tophub-columns-pages-51-60-20260706.md`
- `docs/ops/tophub-columns-pages-61-70-20260706.md`
- `docs/ops/tophub-columns-pages-71-80-20260706.md`
- `docs/ops/tophub-columns-pages-81-90-20260706.md`
- `docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md`

待办：

- [ ] T-C1：先抽每个 page-by-page ledger 的“主要候选 / 最终压缩结果 / 任务型入口 / 不选项”。
- [ ] T-C2：再回到 reassessment 文件补同页内“为什么只选这个，不选那个”。
- [ ] T-C3：对每个最终候选记录它来自哪一页、哪一批、被哪些替代项挤掉。
- [ ] T-C4：对每个明确排除组记录排除理由，不再只留下“未新增”。

### D. 平台核验与执行文件

这些文件只用于证明动作是否真的发生，不作为价值判断的主来源。

- `docs/ops/tophub-finance-platform-execution-verified-20260705.md`
- `docs/ops/tophub-newspapers-platform-execution-verified-20260705.md`
- `docs/ops/tophub-stale-node-removal-platform-verified-20260705.md`
- `docs/ops/tophub-all-page-order-audit-20260705.md`

待办：

- [ ] T-D1：标记哪些动作是平台真实执行。
- [ ] T-D2：标记哪些只是建议、候选或复查。
- [ ] T-D3：不得把“建议添加”写成“已经添加”。

### E. 浏览器 / 链接 / 文化内容分工

这些文件不是订阅源逐源理由的主战场，但要解释为什么某些东西不进 TopHub / Folo。

- `资源/语境/浏览器与链接路由.md`
- `资源/语境/浏览器工具架.md`
- `资源/链接/浏览器订阅候选-2026-07-17.md`
- `资源/链接/浏览器主题包候选.md`
- `资源/链接/浏览器项目与学习入口候选.md`
- `资源/雷达/文化内容入口与影音工作流.md`
- `资源/雷达/得到与知识内容接触谱系.md`

待办：

- [ ] T-E1：抽取“浏览器保存抵达、Folo/TopHub 保存到来、沃壤保存判断”的路由理由。
- [ ] T-E2：抽取得到 / 小宇宙 / 豆瓣 / BibiGPT 为什么保留原平台价值。
- [ ] T-E3：把不进 GitHub 的内容说明清楚：普通摘要、完整个人库、平台收藏、未形成判断的记录不迁入。

---

## 3. 第一批已经确认可直接转为判断链的对象

### AI 目录

来源文件：

- `docs/ops/tophub-ai-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-ai-final-closure-20260705.md`

已确认可转为判断链的来源：

- `Simon Willison's Weblog`：Folo 复查候选；角色是实际工具行为与 Coding Agent 实验；理由包括亲自测试模型、API、Agent、开源工具，重视代码、复现实验和实际失败，并与 Codex、GitHub PR、本地 Agent 工作流相关。
- `Eugene Yan`：Folo 复查候选；角色是生产系统、Evals、搜索推荐与团队机制；理由包括 Product Evals、LLM-as-Judge、长上下文、安全评测、推荐系统、搜索、MCP、新闻 Agent、团队协作和低频结构化更新。
- `AI as Normal Technology`：Folo 复查候选；角色是可靠性、制度条件与反神话判断；理由包括 Agent 可靠性、开放世界评测、反厂商叙事、反 AGI 里程碑单线叙事。
- `Rest of World`：Folo 复查候选；角色是美国之外的技术社会与地区经验；理由包括印度、中国、全球南方、跨国平台、供应链、数据劳动和普通人生活经验。
- `Chip Huyen`：唯一替补；理由是生成式 AI 平台、Agent、模型路由与生产架构；不与 Simon Willison 和 Eugene Yan 同时进入，避免工程源叠加。
- `Stanford CRFM`：任务型研究入口；不订阅整个机构动态。
- `SemiAnalysis`：任务型研究入口；不可替代但付费墙、篇幅和专业负担高。
- `Interconnects`：沃壤保留；质量高，但与已有来源和当前工程候选覆盖。
- `EleutherAI Blog`：研究入口；不要求持续阅读。
- `Hugging Face Blog`：官方工程和开放生态价值高，但更新量大，按具体模型、数据集、训练、推理或工具任务访问。
- `The Gradient` / `AINOW` / `fast.ai` / `Nicholas Carlini`：慢读或任务型入口，不进入本轮四个 Folo 名额。

### 设计目录

来源文件：

- `docs/ops/tophub-design-final-closure-20260705.md`

已确认可转为判断链的来源：

- `designboom`：唯一 Folo 复查候选；理由是持续覆盖建筑、产品、家具、材料、工艺、公共空间、展览与设计文化，补当前视觉体系中缺少的“设计如何进入物、空间与生活”的视角。
- `站酷 / Dribbble / Behance / 500px`：作品瀑布流，不等于持续设计判断；按需视觉漫游，不常驻。
- `优设 / UI 中国 / 站酷文章`：大量 AI 教程、工具清单、作品集内容，不进入持续关系。
- `腾讯 CDC / 百度用户体验中心`：案例与方法档案，作为任务或低频来源。

### 购物目录

来源文件：

- `docs/ops/tophub-shopping-final-closure-20260705.md`

已确认可转为判断链的来源：

- `Craig Mod`：2026-07-17 Folo 复查候选；理由是它不是价格榜或书目榜，而是围绕书、步行、日本地方与城市、摄影、软件、语言、LLM 和独立写作的个人长文来源。
- 淘宝 / 京东 / 什么值得买多数节点：同一商品池不同切片；低价、券后价、短时销量和热度不能替代质量、适用性或真实优惠判断。
- 羊毛线报：需要持续抢时效，混入账号差异、多账号、绑卡、开户、授信、漏洞与规避风控；不做常驻订阅。
- 众筹：有产品发现价值，但不提供充分的量产、交付、退款、售后与知识产权证据；按需。
- 信用卡 / 银行卡 / 虚拟卡 / 跨境支付：高度依赖国家、账户、征信、产品条款和时间；按具体任务核验。

### 财经目录

来源文件：

- `docs/ops/tophub-finance-final-closure-20260705.md`

已确认可转为判断链的来源：

- `财新网｜点击排行榜`：退出；直接事实依据是已停止更新。
- `财新网｜评论排行榜`：退出；直接事实依据是已停止更新。
- `财新网｜首页推荐`：进入 TopHub 数据与结构；角色是中国政策、商业、金融与调查；但仍需观察噪声与重复。
- `日经中文网｜每日最新`：进入 TopHub 数据与结构；角色是日本经济与企业、东亚供应链、日元、央行和产业政策。
- `有知有行｜全部`：Folo 复查候选；角色是长期投资、基金与资产配置、保险边界、组合和再平衡、投资行为与数据解释；不选择黑板报或小酒馆，因为 `全部` 不会把单一栏目误当全部角色。
- `第一财经｜汽车新闻`：保留现有入口；不扩头条、排行、直播、视频、盘前必读等。
- `华尔街见闻｜财经日历`：沃壤工具链接，不是 TopHub 每日订阅；财经日历是查询工具，不是内容流。
- `Foresight News｜文章`：Web3 中文结构源，按需，不进入当前 TopHub 或 Folo。
- `慢雾科技｜技术研究 / 漏洞披露`：写入沃壤专业来源；不选择公司新闻。
- `CoinDesk｜Today`：英文原始报道按需；价格稿占比高，不占常驻英文位置。

---

## 4. 判断链条目格式

本 PR 后续补出的每条判断链应包含以下字段。字段名可以调整，但含义不可丢。

```yaml
- source_id:
  current_route:
  source_file:
  decision_status: recovered_from_pr | recovered_from_chat | partial | needs_recovery
  selected_reason:
  rejected_alternatives:
  route_reason:
  ordering_reason:
  next_review_or_use_condition:
  evidence_ref:
```

说明：

- `selected_reason` 写为什么选这个；
- `rejected_alternatives` 写为什么同页、同类、相近来源不选；
- `route_reason` 写为什么进 TopHub / Folo / 沃壤 / 浏览器 / 按需 / 退出；
- `ordering_reason` 只在涉及排序时填写；
- `decision_status` 不能批量填 `needs_recovery` 后交差。

---

## 5. 下一批执行顺序

### 第一批：把已有最终收口文件里的判断抽成链

- [ ] AI：`tophub-ai-final-closure` + `page-by-page-ledger`
- [ ] 设计：`tophub-design-final-closure`
- [ ] 购物：`tophub-shopping-final-closure`
- [ ] 财经：`tophub-finance-final-closure`
- [ ] 报刊：`tophub-newspapers-final-closure`
- [ ] 娱乐：`tophub-entertainment-final-reconciliation`

交付物：

- [ ] `docs/ops/pr372-decision-recovery-ledger-draft-20260706.md`

### 第二批：补逐页文件里的候选与拒绝项

- [ ] AI 18 页；
- [ ] 开发 25 页；
- [ ] 财经 27 页；
- [ ] 购物 12 页；
- [ ] 专栏 94 页。

交付物：

- [ ] 每批形成 “候选 / 拒绝项 / 去向 / 理由 / 来源文件” 表。

### 第三批：追索 ChatGPT 旧对话

仅在 PR 文件已经不足以还原逐源理由时启动。

- [ ] 用来源名 + TopHub / Folo / 今日热榜 / Follow / 排序 / 为什么 / 候选 / 按需 / GitHub 检索旧对话；
- [ ] 找到旧对话理由后补入判断链；
- [ ] 找不到时不得补编，只能标为“PR 文件无足够过程证据，需要从旧对话继续追索”。

### 第四批：合并前清理

- [ ] 删除或降权没有判断价值的总结文件；
- [ ] 保留真实账本；
- [ ] 保留过程证据；
- [ ] 明确标注哪些文件只是操作流水；
- [ ] PR 描述改写为“状态账本 + 判断追回”的真实说明。

---

## 6. 合并前验收条件

PR #372 不能只因为 TopHub 80 / Folo 27 数字正确就合并。

最低验收条件：

- [ ] 当前真实账本能说明“现在有什么”；
- [ ] 判断链能说明“为什么是这些”；
- [ ] 候选池能说明“为什么只是候选”；
- [ ] 拒绝项能说明“为什么不选同页其他来源”；
- [ ] 平台核验文件能说明“什么动作真的发生”；
- [ ] 找不回理由的条目不能冒充完整沉积。

只有满足这些，PR 才能作为订阅系统的判断资产，而不是只作为一次执行残留。
