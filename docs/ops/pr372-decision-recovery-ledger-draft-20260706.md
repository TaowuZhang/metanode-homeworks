# PR #372 判断链追回草案（第一批，2026-07-06）

## 0. 使用边界

本文件只收录已经能从 PR #372 文件中直接追回的判断链。

本文件不做：

- 不重新审核 TopHub 页面；
- 不把当前模型临时推理写成历史判断；
- 不把未追回条目写成 `not_found` 占位；
- 不把大类口号冒充逐源理由。

每条判断链至少回答：

1. 当前去向；
2. 为什么选它或保留它；
3. 为什么不是同类其他来源；
4. 为什么放到 TopHub / Folo / 沃壤 / 按需 / 退出；
5. 证据来自 PR 哪个文件。

未进入本草案的来源，不代表没有价值，只代表第一批尚未完成判断链追回。

---

## 1. AI / Agent / 机器学习来源

来源文件：

- `docs/ops/tophub-ai-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-ai-final-closure-20260705.md`

### 1.1 TopHub AI 独立分组

```yaml
- source_id: TopHub AI 独立分组
  current_route: no_new_tophub_group
  decision_status: recovered_from_pr
  selected_reason: "不建立 AI 独立分组。AI 不是一个单独容器；模型发布属于科技，Coding Agent 属于开发，安全漏洞属于安全，芯片与数据中心属于数据与结构，教育、劳动和权力属于慢读与思想。"
  rejected_alternatives:
    - "AI 日报"
    - "AI 快讯"
    - "模型热榜"
    - "AI 媒体子频道"
    - "完整论文流"
    - "人工智能独立分组"
  route_reason: "将这些重新装进 AI 热榜，只会重复同一批发布、融资、跑分和情绪标题；当前科技雷达中的综合科技、开发、安全和科学入口已经持续显影 AI。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.2 Simon Willison's Weblog

```yaml
- source_id: "Simon Willison's Weblog"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "实际工具行为与 Coding Agent 实验来源；亲自测试模型、API、Agent 和开源工具，重视代码、复现实验和实际失败，同时覆盖 AI 编程、数据、安全与 Web。"
  rejected_alternatives:
    - "Chip Huyen：作为低频替代，不与 Simon 和 Eugene 同时叠加"
    - "一般 AI 工程教程：缺少真实工具行为和失败记录"
  route_reason: "与 Codex、GitHub PR 和本地 Agent 工作流直接相关；但更新频繁，应放在触发 / 实验观察，不是篇篇必读。"
  next_review_or_use_condition: "2026-07-17 Folo 复查；若未读压力过大，用 Chip Huyen 替换。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.3 Eugene Yan

```yaml
- source_id: "Eugene Yan"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "生产系统、Evals、搜索推荐与团队机制来源；覆盖 Product Evals、LLM-as-Judge、长上下文、安全评测、推荐系统、搜索、MCP、新闻 Agent、真实机器学习系统设计与团队协作。"
  rejected_alternatives:
    - "一般 AI 工程教程：更像教程，不像长期构建、验证和维护系统的判断来源"
  route_reason: "更新低频、结构清楚，并连接 Obsidian、阅读俱乐部、学习和写作；比一般 AI 工程教程更接近如何长期构建、验证和维护系统。"
  next_review_or_use_condition: "2026-07-17 Folo 复查；与 Simon 共同构成工程与系统两条互补线。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.4 AI as Normal Technology

```yaml
- source_id: "AI as Normal Technology"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "可靠性、制度条件与反神话判断来源；讨论 Agent 可靠性与开放世界评测，核验 Agent 建成操作系统、AI 替代工程师等宏大说法，把 AI 放回组织、劳动、法律、科学和制度。"
  rejected_alternatives:
    - "厂商叙事"
    - "AGI 里程碑叙事"
    - "x-risk 单一路线"
  route_reason: "它用于抵消厂商叙事、AGI 里程碑叙事和 x-risk 单一路线；但正常技术同样只是分析框架，不提升为新的默认裁决。"
  next_review_or_use_condition: "2026-07-17 Folo 复查；必须与工程证据和其他社会视角并读。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.5 Rest of World

```yaml
- source_id: "Rest of World"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "美国之外的技术社会与地区经验来源；覆盖印度、中国、全球南方、跨国平台和供应链、数据劳动、移民技术劳工、教育和地区产业变化。"
  rejected_alternatives:
    - "36氪出海：以中国企业全球经营为主要视角，不能替代地区经验"
    - "美国中心英语科技媒体：覆盖重心不同"
  route_reason: "补足现有英语科技媒体明显的美国中心倾向，观察技术怎样进入具体地区和普通人的生活。"
  next_review_or_use_condition: "2026-07-17 Folo 复查。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.6 Chip Huyen

```yaml
- source_id: "Chip Huyen"
  current_route: folo_review_replacement_candidate
  decision_status: recovered_from_pr
  selected_reason: "生成式 AI 平台、Agent、模型路由与生产架构来源，本身值得长期阅读。"
  rejected_alternatives:
    - "与 Simon Willison 和 Eugene Yan 同时进入：会造成工程源叠加"
  route_reason: "当前作为唯一替补；若 Simon Willison 更新过密，则取消 Simon 的候选位置，改为 Chip Huyen；Folo 最终仍保持增加 4 个而不是增加 5 个。"
  next_review_or_use_condition: "2026-07-17 Folo 复查时与 Simon / Eugene 比较。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.7 Stanford CRFM

```yaml
- source_id: "Stanford CRFM"
  current_route: task_research_entry
  decision_status: recovered_from_pr
  selected_reason: "AI 评测与 Agent 可靠性主题的首选研究入口；负责 HELM、模型能力、安全和长上下文评测、Agent 攻防、缺陷协调披露与评测效率。"
  rejected_alternatives:
    - "订阅整个机构动态"
  route_reason: "作为研究入口按主题调用，不进入普通未读流。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.8 SemiAnalysis

```yaml
- source_id: "SemiAnalysis"
  current_route: task_research_entry
  decision_status: recovered_from_pr
  selected_reason: "负责 GPU、HBM、晶圆、数据中心功率和可靠性、训练 / 推理 / TCO、AI 资本开支和供应链。"
  rejected_alternatives:
    - "普通阅读流"
  route_reason: "不可替代，但付费墙、篇幅和专业负担较高；研究算力经济时访问。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.9 Interconnects

```yaml
- source_id: "Interconnects"
  current_route: worang_reference_not_folo
  decision_status: recovered_from_pr
  selected_reason: "开放模型、后训练、模型生态、中美 AI 研究与产业结构来源，质量高。"
  rejected_alternatives:
    - "进入本轮 Folo 四个名额"
  route_reason: "与 Last Week in AI、官方模型来源和当前工程候选存在覆盖；保留在沃壤，不进入这一轮 Folo 四个名额。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.10 EleutherAI Blog

```yaml
- source_id: "EleutherAI Blog"
  current_route: task_research_entry
  decision_status: recovered_from_pr
  selected_reason: "负责 Reward Hacking、机制可解释性、开放模型安全、训练数据、RLHF 和模型透明度。"
  rejected_alternatives:
    - "持续阅读流"
  route_reason: "作为研究入口使用，不要求持续阅读。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.11 Hugging Face Blog

```yaml
- source_id: "Hugging Face Blog"
  current_route: task_engineering_entry
  decision_status: recovered_from_pr
  selected_reason: "官方工程和开放生态价值高。"
  rejected_alternatives:
    - "全量 Folo 未读队列"
  route_reason: "更新量大、产品范围广；需要模型、数据集、训练、推理或具体工具时访问，不建立全量未读队列。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

### 1.12 The Gradient / AINOW / fast.ai / Nicholas Carlini

```yaml
- source_id: "The Gradient"
  current_route: slow_read_task_entry
  decision_status: recovered_from_pr
  selected_reason: "研究、数学、哲学和技术社会的长文平台。"
  rejected_alternatives:
    - "与现有慢读来源一起持续堆积"
  route_reason: "值得按篇读，但不进入本轮 Folo 四个名额。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"

- source_id: "AINOW"
  current_route: institutional_critique_reference
  decision_status: recovered_from_pr
  selected_reason: "劳动、平台、军事、游说和公共利益视角明确。"
  rejected_alternatives:
    - "进入本轮 Folo 四个名额"
  route_reason: "作为制度批判入口保留，但 Rest of World 更适合作为当前持续关系。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"

- source_id: "fast.ai"
  current_route: task_learning_entry
  decision_status: recovered_from_pr
  selected_reason: "实践教育、模型训练和对 AI 工程文化的批评。"
  rejected_alternatives:
    - "普通持续订阅"
  route_reason: "按课程、文章和具体问题调用。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"

- source_id: "Nicholas Carlini"
  current_route: task_security_entry
  decision_status: recovered_from_pr
  selected_reason: "模型攻击、安全、隐私和研究实践来源。"
  rejected_alternatives:
    - "普通持续订阅"
  route_reason: "出现具体安全或评测问题时优先查阅。"
  evidence_ref: "docs/ops/tophub-ai-final-closure-20260705.md"
```

---

## 2. 设计来源

来源文件：

- `docs/ops/tophub-design-final-closure-20260705.md`

### 2.1 designboom

```yaml
- source_id: "designboom"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "持续覆盖建筑、产品、家具、材料、工艺、公共空间、展览与设计文化。"
  rejected_alternatives:
    - "站酷 / Dribbble / Behance / 500px 的作品瀑布流"
    - "优设、站酷、UI 中国的大量 AI 教程、工具清单和作品集内容"
    - "腾讯、百度旧设计团队留下的案例与方法档案作为常驻订阅"
  route_reason: "它能补当前视觉体系中缺少的‘设计如何进入物、空间与生活’的视角；若通过复查，承担开放天线，不进入 TopHub 热榜排序。"
  next_review_or_use_condition: "2026-07-17 检查实际更新频率、未读负担、建筑 / 产品 / 工艺 / 材料比例，以及是否提供 CNU、胶片的味道、iDaily 没有的视角。"
  evidence_ref: "docs/ops/tophub-design-final-closure-20260705.md"
```

### 2.2 站酷 / Dribbble / Behance / 500px / 设计教程流

```yaml
- source_id: "Dribbble / 站酷 / 500px / Designspiration / UI 中国 / Designmodo"
  current_route: rejected_or_on_demand_visual_browsing
  decision_status: recovered_from_pr
  selected_reason: "这些来源可以作为视觉灵感或任务型参考，但不等于持续设计判断。"
  rejected_alternatives:
    - "Dribbble 实时、本周、本月、年度和历史总榜"
    - "站酷作品总榜、全部推荐、文章总榜及重复推荐流"
    - "500px 热门作品"
    - "Designspiration Popular"
    - "UI 中国推荐文章"
    - "Designmodo 当前邮件营销内容"
  route_reason: "作品瀑布流、教程工具清单、作品集内容不构成持续设计判断；需要视觉灵感时按需浏览。"
  evidence_ref: "docs/ops/tophub-design-final-closure-20260705.md"
```

---

## 3. 购物来源

来源文件：

- `docs/ops/tophub-shopping-final-closure-20260705.md`

### 3.1 购物目录整体

```yaml
- source_id: "TopHub 购物目录"
  current_route: task_invoked_not_daily_subscription
  decision_status: recovered_from_pr
  selected_reason: "购物目录有工具价值，但不应成为发现可以买什么的日常信息源。"
  rejected_alternatives:
    - "把淘宝 / 京东 / 什么值得买 / 羊毛线报做成持续订阅流"
    - "没有明确需求时由促销信息主动生成购买任务"
  route_reason: "当真实需求已经出现后，按任务调用价格、渠道、用户经验、版本、规格、权益和交付信息；任务结束后退出注意力。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"
```

### 3.2 Craig Mod

```yaml
- source_id: "Craig Mod"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "不是价格榜或书目榜，而是稳定围绕书、步行、日本地方与城市、摄影、软件、语言、LLM 和独立写作展开的个人长文来源。"
  rejected_alternatives:
    - "TopHub 购物榜单"
    - "立即加入 Folo"
  route_reason: "加入 2026-07-17 Folo 复查名单；复查不是承诺加入。"
  next_review_or_use_condition: "检查真实更新频率、全文可读性、未读负担和与现有来源的重复度。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"
```

### 3.3 淘宝 / 京东 / 什么值得买

```yaml
- source_id: "淘宝 / 京东 / 什么值得买多数节点"
  current_route: on_demand_purchase_tools
  decision_status: recovered_from_pr
  selected_reason: "它们可以在真实购买任务出现后用于价格、渠道、历史价、跨平台价格和具体商品用户经验核验。"
  rejected_alternatives:
    - "作为日常信息源"
    - "从低价、券后价、短时销量和热度直接生成购买任务"
  route_reason: "淘宝、京东和什么值得买多数节点是同一商品池的不同切片；低价、券后价、短时销量和热度不能替代质量、适用性或真实优惠判断。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"
```

### 3.4 羊毛线报

```yaml
- source_id: "羊毛线报"
  current_route: rejected_as_daily_subscription
  decision_status: recovered_from_pr
  selected_reason: "羊毛线报可能包含优惠线索，但不适合作为持续信息流。"
  rejected_alternatives:
    - "所有综合羊毛线报常驻订阅"
    - "依赖多账号、虚假资料、漏洞、取消订单或规避风控的方案"
  route_reason: "它要求持续抢时效，混入账号差异、多账号、绑卡、开户、授信、漏洞与规避风控；不进入常驻订阅。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"
```

### 3.5 众筹 / 图书 / 信用卡支付

```yaml
- source_id: "淘宝众筹 / 小米有品众筹 / 摩点众筹"
  current_route: on_demand_product_discovery
  decision_status: recovered_from_pr
  selected_reason: "众筹具有产品发现价值，尤其是 AI 小硬件、米家生态、消费电子、家居产品、桌游、TRPG、出版、模型与文化周边。"
  rejected_alternatives:
    - "常驻订阅"
  route_reason: "众筹不提供充分的量产、交付、退款、售后与知识产权证据；只按需浏览并核验。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"

- source_id: "当当 / 京东 / 豆瓣书店 / 缺书网"
  current_route: on_demand_book_purchase_tools
  decision_status: recovered_from_pr
  selected_reason: "用于确定书目、版本、ISBN 后的购买、比价或促销查询。"
  rejected_alternatives:
    - "拆成多个未读流"
  route_reason: "常驻书籍发现已由书格每日好书和豆瓣新书速递承担；豆瓣分类新书页按阅读主题主动打开。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"

- source_id: "美国信用卡指南 / 美卡论坛 / V2EX 信用卡"
  current_route: on_demand_financial_payment_tools
  decision_status: recovered_from_pr
  selected_reason: "用于美国卡与银行账户任务、特定卡种、里程票、实践经验、大陆 / 香港 / 境外消费与订阅支付线索。"
  rejected_alternatives:
    - "常驻信用卡信息流"
    - "未经核验的虚拟卡平台与跨区支付方案"
  route_reason: "信用卡、银行卡、虚拟卡与跨境支付高度依赖国家、账户、征信、产品条款和时间；最终依据始终是银行官网、价目表、条款、监管披露和客服确认。"
  evidence_ref: "docs/ops/tophub-shopping-final-closure-20260705.md"
```

---

## 4. 财经来源

来源文件：

- `docs/ops/tophub-finance-final-closure-20260705.md`

### 4.1 财新点击排行榜 / 评论排行榜

```yaml
- source_id: "财新网｜点击排行榜"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原先作为财新入口存在。"
  rejected_alternatives:
    - "继续承担持续入口角色"
  route_reason: "直接事实依据是已停止更新；排行榜逻辑是否理想只是次要问题，失效节点本身已经不能继续承担持续入口角色。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "财新网｜评论排行榜"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原先作为财新入口存在。"
  rejected_alternatives:
    - "继续承担持续入口角色"
  route_reason: "直接事实依据是已停止更新；排行榜逻辑是否理想只是次要问题，失效节点本身已经不能继续承担持续入口角色。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"
```

### 4.2 财新首页推荐 / 日经中文每日最新

```yaml
- source_id: "财新网｜首页推荐"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "中国政策、商业、金融与调查入口；替代已停止更新的财新点击榜和评论榜。"
  rejected_alternatives:
    - "继续使用已停止更新的点击榜和评论榜"
  route_reason: "财新首页推荐保留为可观察入口，但不把名称本身视为精选的充分证据；后续若实际内容证明噪声过高，再按真实使用结果复审。"
  ordering_reason: "在数据与结构中位于国家统计局数据和数据解读之后、日经中文之前，承担国内编辑入口。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "日经中文网｜每日最新"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "日本经济与企业、东亚供应链、日元、央行和产业政策入口。"
  rejected_alternatives:
    - "用联合早报或其他亚洲华语财经源替代当前常驻位"
  route_reason: "财经目录真正缺少日本与东亚产业、供应链、货币政策视角；日经中文已正式进入 TopHub。"
  ordering_reason: "在数据与结构中位于财新首页推荐之后、FT 中文之前，补东亚视角。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"
```

### 4.3 有知有行｜全部

```yaml
- source_id: "有知有行｜全部"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr
  selected_reason: "长期投资、基金与资产配置、保险边界、组合和再平衡、投资行为与数据解释来源。"
  rejected_alternatives:
    - "有知有行｜知行黑板报"
    - "有知有行｜知行小酒馆"
  route_reason: "三者存在重叠，全部是最完整且不会把单一栏目误当全部角色的入口；执行时间仍放在 2026-07-17 Folo 统一复查，不擅自修改实际订阅。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"
```

### 4.4 第一财经 / FT 中文 / 联合早报 / 经济观察 / 21 财经

```yaml
- source_id: "第一财经｜汽车新闻"
  current_route: keep_existing_tophub_node
  decision_status: recovered_from_pr
  selected_reason: "汽车新闻入口继续保留。"
  rejected_alternatives:
    - "第一财经头条"
    - "第一财经排行"
    - "第一财经直播"
    - "第一财经视频"
    - "盘前必读"
  route_reason: "其余第一财经内容已被财新、FT中文、日经、央视栏目和公共新闻覆盖。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "FT中文网｜十大热门文章"
  current_route: keep_existing_tophub_node
  decision_status: recovered_from_pr
  selected_reason: "国际经济与商业、全球政策和市场解释入口。"
  rejected_alternatives:
    - "继续扩容 FT 中文其他入口"
  route_reason: "保持热门文章入口，不扩容。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "联合早报"
  current_route: task_source_not_constant_subscription
  decision_status: recovered_from_pr
  selected_reason: "新加坡、东南亚与亚洲华语财经来源。"
  rejected_alternatives:
    - "常驻 TopHub 节点"
  route_reason: "保持任务源，不进入常驻，避免和日经、FT中文重复。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "经济观察网"
  current_route: task_search_source
  decision_status: recovered_from_pr
  selected_reason: "公司、产业和政策长文来源。"
  rejected_alternatives:
    - "常驻节点"
  route_reason: "保持任务检索，不增加常驻节点。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "21财经｜数读"
  current_route: task_data_source
  decision_status: recovered_from_pr
  selected_reason: "数据可视化和结构性资料来源。"
  rejected_alternatives:
    - "常驻节点"
  route_reason: "保持任务源。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"
```

### 4.5 财经日历 / Web3 / 官方披露

```yaml
- source_id: "华尔街见闻｜财经日历"
  current_route: worang_tool_link
  decision_status: recovered_from_pr
  selected_reason: "用于安排一周、核对央行、数据、财报或政策事件。"
  rejected_alternatives:
    - "TopHub 每日订阅"
  route_reason: "财经日历是查询工具，不是内容流。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "Foresight News｜文章"
  current_route: on_demand_web3_structure_source
  decision_status: recovered_from_pr
  selected_reason: "稳定币、支付、监管、协议、安全和行业结构长文的中文结构源。"
  rejected_alternatives:
    - "当前 TopHub 常驻"
    - "当前 Folo 常驻"
  route_reason: "保持按需，不进入当前 TopHub 或 Folo。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "慢雾科技｜技术研究 / 漏洞披露"
  current_route: worang_professional_source
  decision_status: recovered_from_pr
  selected_reason: "Web3 安全技术研究与漏洞披露来源。"
  rejected_alternatives:
    - "慢雾公司新闻"
  route_reason: "写入沃壤专业来源，不选择公司新闻。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "CoinDesk｜Today"
  current_route: on_demand_english_original_reporting
  decision_status: recovered_from_pr
  selected_reason: "英文原始报道入口。"
  rejected_alternatives:
    - "常驻英文位置"
  route_reason: "价格稿占比仍高，不占常驻英文位置。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"

- source_id: "BBX / PANews 快讯 / Odaily 7×24 / BitKan / CryptoPanic / TechFlow 7×24 / Followin / U.Today / Bitcoin News"
  current_route: rejected_web3_price_or_flash_streams
  decision_status: recovered_from_pr
  selected_reason: "这些来源可能提供加密市场信息，但不承担本系统需要的 Web3 技术、安全、监管与结构判断。"
  rejected_alternatives:
    - "7×24 快讯流"
    - "价格预测流"
    - "巨鲸、杠杆、爆仓、支撑位信息流"
  route_reason: "不需要币价交易流，不把 Web3 技术 / 安全 / 监管与币价流混在一起。"
  evidence_ref: "docs/ops/tophub-finance-final-closure-20260705.md"
```

---

## 5. 报刊来源

来源文件：

- `docs/ops/tophub-newspapers-final-closure-20260705.md`

### 5.1 人民日报电子版 / 新华每日电讯电子报

```yaml
- source_id: "人民日报｜电子版"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原先作为全国官方日报入口。"
  rejected_alternatives:
    - "继续承担每日公共温度角色"
  route_reason: "TopHub 节点样本约停留在 2024 年，已经不能承担每日公共温度角色。"
  evidence_ref: "docs/ops/tophub-newspapers-final-closure-20260705.md"

- source_id: "新华每日电讯｜电子报"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "保持当前更新，题材覆盖全国政策、社会、区域、产业和人物。"
  rejected_alternatives:
    - "单一财经报作为公共温度入口"
    - "单一法制报作为公共温度入口"
    - "继续使用失效的人民日报电子版"
  route_reason: "比单一财经报或法制报更适合作为一个低频、官方、广谱的日报入口；替换一个已经失效的全国广谱日报节点。"
  ordering_reason: "加入公共温度分组。"
  evidence_ref: "docs/ops/tophub-newspapers-final-closure-20260705.md"
```

### 5.2 财经报 / 法制报 / 中央级报刊

```yaml
- source_id: "法治日报 / 检察日报 / 上海法治报 / 地方法制报"
  current_route: on_demand_legal_task_sources
  decision_status: recovered_from_pr
  selected_reason: "法治日报全国性强，检察日报和上海法治报适合作为法律任务来源。"
  rejected_alternatives:
    - "作为全国公共温度常驻入口"
  route_reason: "法治日报角色过窄，头版仍与综合时政重复；地方法制报不适合承担全国公共温度入口。"
  evidence_ref: "docs/ops/tophub-newspapers-final-closure-20260705.md"

- source_id: "每日经济新闻 / 21世纪经济报道 / 经济参考报 / 中国经营报 / 证券日报"
  current_route: on_demand_financial_newspaper_sources
  decision_status: recovered_from_pr
  selected_reason: "财经报有任务检索价值。"
  rejected_alternatives:
    - "新增持续订阅"
    - "替代新华每日电讯作为公共温度入口"
  route_reason: "财经报样本没有产生更好的公共温度候选；财经内容更适合在具体任务中检索。"
  evidence_ref: "docs/ops/tophub-newspapers-final-closure-20260705.md"

- source_id: "中国教育报 / 科技日报官方网站 / 光明日报官方网站"
  current_route: on_demand_central_newspaper_sources
  decision_status: recovered_from_pr
  selected_reason: "中央级报刊有任务调用价值。"
  rejected_alternatives:
    - "新增持续订阅"
  route_reason: "当前缺口不是更多报纸，而是替换一个已经失效的全国广谱日报节点；报纸电子版通常高频、大体量、版面型，不适合进入 Folo 制造未读负担。"
  evidence_ref: "docs/ops/tophub-newspapers-final-closure-20260705.md"
```

---

## 6. 娱乐 / 文化作品来源

来源文件：

- `docs/ops/tophub-entertainment-final-reconciliation-20260705.md`

### 6.1 娱乐大类整体

```yaml
- source_id: "TopHub 娱乐大类"
  current_route: tophub_culture_work_discovery
  decision_status: recovered_from_pr
  selected_reason: "最终保留作品发现、评论与解释、创作与产业、少量传播温度、视觉漫游与摄影叙事五类功能。"
  rejected_alternatives:
    - "明星私生活"
    - "粉圈热搜"
    - "平台重复细分榜"
    - "历史累计榜"
    - "高奢商品连续流"
    - "来源边界不清节点"
  route_reason: "娱乐目录不是粉圈流，而是文化作品的公共回声、发现、评论、产业与视觉入口。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"
```

### 6.2 新增 10 个 TopHub 节点

```yaml
- source_id: "华丽志"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "作为消费产业 / 时尚商业与结构观察入口，加入数据与结构。"
  rejected_alternatives: []
  route_reason: "补娱乐与消费产业之间的结构视角。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "豆瓣电影｜最受欢迎的影评"
  current_route: tophub_slow_read_and_thought
  decision_status: recovered_from_pr
  selected_reason: "承担电影作品的评论与解释。"
  rejected_alternatives: []
  route_reason: "加入慢读与思想，不作为影游音乐榜单。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "书格｜每日好书"
  current_route: tophub_slow_read_and_thought
  decision_status: recovered_from_pr
  selected_reason: "承担书籍与文献发现。"
  rejected_alternatives:
    - "在购物图书目录中再拆多个图书未读流"
  route_reason: "作为现有常驻书籍发现入口，后续购物图书按需，不再拆成多个未读流。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "豆瓣｜华语新碟榜"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "华语新音乐作品发现。"
  rejected_alternatives: []
  route_reason: "进入影游音乐分组，承担作品发现。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "网易云音乐｜原创榜"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "原创音乐作品发现。"
  rejected_alternatives:
    - "网易云音乐｜飙升榜"
  route_reason: "原创榜比飙升榜更接近作品发现，不是短时传播热度。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "豆瓣｜全球口碑剧集榜"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "剧集作品发现。"
  rejected_alternatives: []
  route_reason: "进入影游音乐分组，补剧集口碑发现。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "新京报｜娱乐"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "娱乐新闻与文化传播温度来源。"
  rejected_alternatives: []
  route_reason: "承担少量传播温度，不扩张为粉圈热搜。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "INDIENOVA｜文章"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "独立游戏文章与创作文化来源。"
  rejected_alternatives: []
  route_reason: "进入影游音乐，承担创作与产业中的游戏侧。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "喷嚏网｜乐影"
  current_route: tophub_film_game_music
  decision_status: recovered_from_pr
  selected_reason: "音乐 / 影像文化内容入口。"
  rejected_alternatives: []
  route_reason: "进入影游音乐，承担乐影作品漫游。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "CNU视觉联盟｜每日精选"
  current_route: tophub_visual_and_nature
  decision_status: recovered_from_pr
  selected_reason: "低频人工策展和作品发现。"
  rejected_alternatives:
    - "用它替代胶片的味道"
  route_reason: "与胶片的味道功能不同；CNU 负责低频人工策展和作品发现。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"
```

### 6.3 取消与保留

```yaml
- source_id: "豆瓣｜热门图书-非虚构类"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原为图书热度入口。"
  rejected_alternatives:
    - "继续保留历史 / 重复图书榜"
  route_reason: "娱乐大类最终不保留平台重复细分榜、历史累计榜；书籍发现由豆瓣新书、书评、微信读书、书格等承担。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "网易云音乐｜飙升榜"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原为音乐短时热度入口。"
  rejected_alternatives:
    - "继续作为音乐发现入口"
  route_reason: "由网易云原创榜和 Apple Music / 豆瓣新碟等承担作品发现，减少短时传播热度。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "开眼视频｜日报"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原为视频日报入口。"
  rejected_alternatives:
    - "继续保留来源边界不清或重复视频日报"
  route_reason: "娱乐大类不保留来源边界不清节点。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "胶片的味道"
  current_route: keep_tophub_visual_and_nature
  decision_status: recovered_from_pr
  selected_reason: "承担摄影叙事、城市观察、胶片器材和画幅方法。"
  rejected_alternatives:
    - "被 CNU视觉联盟｜每日精选替代"
  route_reason: "它与 CNU 每日精选功能不同；CNU 负责低频人工策展和作品发现，胶片的味道负责摄影叙事与方法。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"
```

### 6.4 娱乐相关 Folo 复查候选

```yaml
- source_id: "基本読書"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "娱乐大类最终闭环中列为 Folo 复查候选。"
  rejected_alternatives: []
  route_reason: "复查前 Folo 保持 27 项，不提前添加。具体逐源理由需要从对应娱乐逐页 closure 或旧对话继续追回。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "木遥的窗子"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "娱乐大类最终闭环中列为 Folo 复查候选。"
  rejected_alternatives: []
  route_reason: "复查前 Folo 保持 27 项，不提前添加。具体逐源理由需要从对应娱乐逐页 closure 或旧对话继续追回。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "Yuko's Blog"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "娱乐大类最终闭环中列为 Folo 复查候选。"
  rejected_alternatives: []
  route_reason: "复查前 Folo 保持 27 项，不提前添加。具体逐源理由需要从对应娱乐逐页 closure 或旧对话继续追回。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"

- source_id: "Huiliu"
  current_route: folo_review_candidate_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "娱乐大类最终闭环中列为 Folo 复查候选。"
  rejected_alternatives: []
  route_reason: "复查前 Folo 保持 27 项，不提前添加。具体逐源理由需要从对应娱乐逐页 closure 或旧对话继续追回。"
  evidence_ref: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md"
```

---

## 7. 下一批追回对象

本草案第一批已经覆盖：

- AI 最终四个 Folo 候选、一个替补、多个任务型入口；
- 设计唯一候选与被拒绝的作品瀑布流 / 教程流；
- 购物整体路由、Craig Mod、购物工具、羊毛线报、众筹、图书、信用卡支付；
- 财经替换、Folo 候选、财经任务源、财经日历、Web3、官方披露；
- 报刊替换、财经报 / 法制报 / 中央级报刊按需逻辑；
- 娱乐新增、取消、保留、Folo 候选的第一层链条。

下一批必须继续追回：

1. 开发目录 25 页：哪些进入任务入口、哪些保留为开发雷达、哪些不进 Folo；
2. AI 6 个 reassessment 文件：把每批“主要候选”中未进入最终四个名额的来源逐个写清为什么让位；
3. 娱乐各单项 closure：补 `基本読書`、`木遥的窗子`、`Yuko's Blog`、`Huiliu`、`maxOS`、`Velas 电波站`的逐源理由；
4. 专栏 94 页：从页面级复审中抽出候选池、明确排除组、关键词相关但不加入的理由；
5. 浏览器 / 链接：把“为什么留在浏览器、资源链接、语境、沃壤，而不是 TopHub / Folo”的关系链补齐。
