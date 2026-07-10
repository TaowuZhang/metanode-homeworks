# PR #372 判断链追回：科技任务触发与高频流来源（2026-07-06）

## 0. 边界

本文件是 `科技` 大类的补充 ledger，追回 PR #372 中已闭环的 AI、App、快讯、限免、教育、电商、数码、酷安、Google、TechWeb 等来源判断。

纳入文件：

- `docs/ops/tophub-technology-ai-closure-20260704.md`
- `docs/ops/tophub-technology-app-closure-20260704.md`
- `docs/ops/tophub-technology-flash-news-closure-20260704.md`
- `docs/ops/tophub-technology-freebies-closure-20260704.md`
- `docs/ops/tophub-technology-education-closure-20260704.md`
- `docs/ops/tophub-technology-ecommerce-closure-20260704.md`
- `docs/ops/tophub-technology-digital-closure-20260704.md`
- `docs/ops/tophub-technology-coolapk-closure-20260704.md`
- `docs/ops/tophub-technology-google-closure-20260704.md`
- `docs/ops/tophub-technology-techweb-closure-20260704.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 本组总规则

```yaml
- source_id: "科技任务触发与高频流总规则"
  current_route: mostly_on_demand_or_folo_review_not_tophub_expansion
  decision_status: recovered_from_pr
  selected_reason: "这批目录包含大量真实有用的来源，但它们多数只在具体任务、具体产品、具体项目、具体学习目标或具体问题出现时有价值。"
  rejected_alternatives:
    - "因为 AI / App / 教育 / 电商 / Google / 数码重要就新建常驻入口"
    - "把快讯、日报、限免、榜单和论坛热帖放进 Folo"
    - "把具体任务源改造成每日阅读责任"
    - "建立宽泛追踪词"
  route_reason: "TopHub 负责低噪声发现和互补雷达；Folo 保存稳定作者、官方源和真实持续阅读关系；任务型来源按需使用。高频流、交易窗口、平台榜单和论坛经验不能替代原始来源。"
  reuse_rule: "以后遇到任务触发来源，先问是否有真实责任关系：正在使用、正在买、正在开发、正在经营、正在学习、正在参赛、正在排障。没有责任关系就不常驻。"
  evidence_locator: "technology ai / app / flash-news / freebies / education / ecommerce / digital / coolapk / google / techweb closure files"
  next_action: "作为科技同族复用规则。"
```

---

## 2. 人工智能：不新增 AI 专属入口

```yaml
- source_id: "科技 > 人工智能 小标签整体"
  current_route: no_new_tophub_no_ai_group_no_broad_tracker
  decision_status: recovered_from_pr
  selected_reason: "AI 重要，但它不是一个应与科技、开发、科学、汽车、数据结构并列的独立内容容器。AI 变化已经通过 IT之家、极客公园、TechCrunch、The Verge、The Register、科学、汽车、开发、产品、公司报道、精确追踪器和 Folo 中的 Google / Last Week in AI 关系进入系统。"
  rejected_alternatives:
    - "新增 AI 日报 / AI 快讯 / AI 热榜"
    - "新增 AI 独立第八组"
    - "订阅 Hugging Face Trending / Daily Papers"
    - "订阅全部官方框架博客"
    - "创建人工智能 / 大模型 / Agent 等宽泛追踪器"
  route_reason: "不新增 TopHub，不替换，不删除。AI 新闻继续通过现有科技、开发、科学、数据结构节点和精确追踪进入；深度阅读留给少量直接作者，论文和官方博客按任务调用。"
  reuse_rule: "AI 来源先按真实发生位置分流：模型产品进科技雷达，编程和 Agent 进开发 / 项目，算力芯片进数据结构，AI4S 进科学，治理劳动进思想或结构。不要把所有 AI 重新装进一个总入口。"
  evidence_locator: "docs/ops/tophub-technology-ai-closure-20260704.md#核心结论; #闭环状态"
  next_action: "无；AI 小标签已闭环。"

- source_id: "AI 深读与作者候选"
  current_route: folo_review_not_immediate_add
  decision_status: recovered_from_pr
  selected_reason: "Simon Willison、Chip Huyen、AI as Normal Technology、Rest of World 等直接补出工程实测、生产系统、制度框架和全球技术社会视角。"
  rejected_alternatives:
    - "立即加入多个 AI newsletter"
    - "继续叠加多个综合 AI 周报 / 日报"
    - "AI 播客批量加入"
  route_reason: "Folo 仍为 27 项，2026-07-17 前保持基线。高优先级与条件候选进入比较池，不自动加入。AI 播客因单集时间长、重复度高和当前无固定收听流程，全部按需。"
  reuse_rule: "AI 深读源必须经复查：直接 RSS、更新频率、实际阅读全文、与已有 Last Week in AI / Google 源重复度、阅读负担。"
  evidence_locator: "docs/ops/tophub-technology-ai-closure-20260704.md#五Folo复查候选; #七播客"
  next_action: "Folo 复查。"

- source_id: "AI 官方 / 论文 / 模型 / 数据集来源"
  current_route: on_demand_ai_research_and_engineering_sources
  decision_status: recovered_from_pr
  selected_reason: "Hugging Face、arXiv、JMLR、Nature ML、Google AI、NVIDIA、LangChain、LlamaIndex、CRFM、EleutherAI 等都是可靠或有用来源。"
  rejected_alternatives:
    - "TopHub 常驻论文 / 模型榜"
    - "普通 Folo 未读队列"
  route_reason: "每日数量巨大，榜单热度不等于长期价值，模型 / 数据集 / Spaces 重复多，原始论文标题缺任务上下文。官方博客价值通常与具体技术栈绑定，不能把可查阅文档变成必须持续阅读的消息流。"
  reuse_rule: "从问题出发检索模型、论文、框架和官方博客；能力判断回模型卡、代码、论文、许可证和可复现实验。"
  evidence_locator: "docs/ops/tophub-technology-ai-closure-20260704.md#二HuggingFace与论文节点; #三官方公司框架和工程博客"
  next_action: "按任务。"
```

---

## 3. App 与软件发现：不把发现变成安装责任

```yaml
- source_id: "App 目录整体"
  current_route: scoped_review_no_broad_app_store_expansion
  decision_status: recovered_from_pr
  selected_reason: "App 目录前六页已足以确认：大量地区、设备、榜型和分类排行属于同构 App Store 排行能力，不构成不同信息能力。"
  rejected_alternatives:
    - "逐国逐区复制 App Store 榜"
    - "把 App Store 排行加入 TopHub"
    - "把软件目录、限免、论坛、下载站加入 Folo"
  route_reason: "本文件以六页为完整审核范围；未提供的地区商店排行统一归档为有意排除的同构地区排行。App Store 榜只说明分发或收入位置，不能证明质量、隐私、长期价值和用户适配度。"
  reuse_rule: "App 榜单只能作为市场样本；安装、购买、订阅前必须核对价格、隐私、权限、开发者、更新历史、平台兼容性和现有工具栈。"
  evidence_locator: "docs/ops/tophub-technology-app-closure-20260704.md#页面范围; #同构地区商店排行的统一规则"
  next_action: "无；App 范围裁剪有效。"

- source_id: "小众软件｜每日最新"
  current_route: folo_existing_no_tophub_duplicate
  decision_status: recovered_from_pr
  selected_reason: "能发现开源工具、独立软件、桌面应用、系统组件与自托管项目，并已在 Folo 深度阅读中建立直接来源关系。"
  rejected_alternatives:
    - "小众软件主站重复加入 TopHub"
    - "小众软件官方论坛两个节点常驻"
  route_reason: "主站继续由 Folo 承担，不重复加入 TopHub；论坛混入求助、自荐、活动、游戏限免与重复转载，噪声高，按需。"
  reuse_rule: "已在 Folo 建立直接来源关系的主站，不在 TopHub 重复；论坛只在寻找新工具或开发者自荐时搜索。"
  evidence_locator: "docs/ops/tophub-technology-app-closure-20260704.md#2小众软件"
  next_action: "无；保留 Folo。"

- source_id: "效率火箭"
  current_route: folo_high_priority_review_candidate
  decision_status: recovered_from_pr
  selected_reason: "与 Obsidian、Notion、Agent、Markdown、本地工具和跨平台工作流直接相关，不只是列 App，而是在讨论工作流、自动化、本地性和工具结构。"
  rejected_alternatives:
    - "立即加入 TopHub"
    - "立即加入 Folo"
    - "用 AppSo / 软件目录替代"
  route_reason: "文章需要阅读全文，不适合只靠标题扫过；进入 2026-07-17 Folo 高优先级复查。"
  reuse_rule: "工具工作流作者源先验证是否持续原创、是否真实阅读、是否和已有源重复，再进 Folo。"
  evidence_locator: "docs/ops/tophub-technology-app-closure-20260704.md#8效率火箭"
  next_action: "Folo 复查。"

- source_id: "反斗软件｜最新发布"
  current_route: tophub_candidate_from_app_review_but_later_superseded_to_on_demand
  decision_status: recovered_from_pr
  selected_reason: "App 审核阶段它是小型软件与扩展发布候选：标题能看出平台、用途和是否值得进一步打开，更新量低，不以折扣返利和下载量为核心。"
  rejected_alternatives:
    - "AppSo 鲜面连线"
    - "软件目录 / 下载站"
    - "限免站替代"
  route_reason: "App 文件中形成 TopHub 新增候选，但科技 final closeout 后未采用，退为按需；说明候选可被后续全局收口修订。"
  reuse_rule: "小工具新发布源有价值，但必须和少数派 Matrix、小众软件、HelloGitHub、Apple / macOS 追踪器比较后再决定是否常驻。"
  evidence_locator: "docs/ops/tophub-technology-app-closure-20260704.md#9反斗软件最新发布; docs/ops/tophub-technology-final-closeout-20260704.md#本轮最终未采用候选"
  next_action: "按需。"
```

---

## 4. 快讯与限免：发现不等于责任

```yaml
- source_id: "科技 > 快讯 小标签整体"
  current_route: no_new_tophub_no_folo_no_tracker
  decision_status: recovered_from_pr
  selected_reason: "快讯能发现公共事件、公司公告、政策、市场价格、行业变化和产品消息。"
  rejected_alternatives:
    - "通用财经快讯常驻"
    - "AI 快讯常驻"
    - "加密 7×24 流常驻"
    - "电商、汽车、融资、安全等行业快讯常驻"
  route_reason: "37 个节点没有形成新的 TopHub 候选。快讯目录最大问题是同一件事被太多流重复，且快讯不提供最终状态、原始文件、上下文和传闻 / 正式公告边界。"
  reuse_rule: "快讯负责发现，最终状态回原始文件。按问题选择：监管 / 交易所 / 公司公告优先，行业专业源其次，通用快讯只作线索。"
  evidence_locator: "docs/ops/tophub-technology-flash-news-closure-20260704.md#总体判断; #TopHub候选; #来源优先顺序"
  next_action: "无；快讯小标签已闭环。"

- source_id: "加密与 Web3 快讯"
  current_route: on_demand_crypto_news_only
  decision_status: recovered_from_pr
  selected_reason: "对短线交易者或具体协议研究可能有用。"
  rejected_alternatives:
    - "PANews 多节点并存"
    - "律动 / Odaily / BitKan / 深潮 / Followin 常驻"
    - "币价 / 巨鲸 / ETF / 空投 / 代币宽追踪"
  route_reason: "这些来源高度重复推送币价、巨鲸仓位、ETF、杠杆爆仓、监管、上币、Meme 和链上动作。用户当前没有分钟级加密市场工作流。"
  reuse_rule: "研究以太坊、稳定币、RWA、某协议或监管时临时选一个行业源，并回链上数据、协议公告、监管文件和项目原文。"
  evidence_locator: "docs/ops/tophub-technology-flash-news-closure-20260704.md#四加密货币与Web3快讯"
  next_action: "按需。"

- source_id: "科技 > 限免 小标签整体"
  current_route: on_demand_freebies_no_subscription
  decision_status: recovered_from_pr
  selected_reason: "官方商店和部分媒体说明能帮助发现限时免费游戏、软件、App、小说和促销。"
  rejected_alternatives:
    - "限免流常驻 TopHub"
    - "限免流进入 Folo"
    - "限免 / 免费 / Epic / Steam 宽追踪"
  route_reason: "限免是交易窗口，不是长期关系。免费不等于值得安装、值得使用或值得长期保留；过期后价值迅速归零。Epic 官方是条件升级路径，但当前没有稳定领取习惯。"
  reuse_rule: "临时看限免时优先官方商店，其次带上下文媒体说明，再是聚合发现；最终领取状态、价格、地区限制和版本信息必须回官方页面核验。"
  evidence_locator: "docs/ops/tophub-technology-freebies-closure-20260704.md#总体判断; #来源优先级; #三个固定问题"
  next_action: "按需。"
```

---

## 5. 教育与电商：领域重要不等于常驻

```yaml
- source_id: "科技 > 教育 小标签整体"
  current_route: no_new_education_group_task_triggered
  decision_status: recovered_from_pr
  selected_reason: "教育目录混合教育产业、政策媒体、科研机构新闻、考研、竞赛、招生、英语学习、论坛、个人博客和泛热榜。"
  rejected_alternatives:
    - "新增教育组"
    - "芥末堆 / 多知网 / 新京报教育常驻"
    - "MIT Research News 立即加入 Folo"
    - "考研 / 高考 / NOI 宽追踪"
  route_reason: "21 个节点没有形成新的 TopHub 候选。教育产业媒体、学校招生、竞赛报名和英语学习都依赖具体任务；当前没有证据表明它们需要持续占据注意力。"
  reuse_rule: "教育来源必须先问任务：某所学校、某年招生、某个教育产品、某项政策、某省 NOI/CSP。没有具体责任，不建立常驻。"
  evidence_locator: "docs/ops/tophub-technology-education-closure-20260704.md#核心判断; #TopHub候选; #精确关键词追踪"
  next_action: "按任务。"

- source_id: "MIT Research News"
  current_route: high_value_on_demand_research_news
  decision_status: recovered_from_pr
  selected_reason: "包含医学影像、语言网络、Agentic AI、石墨烯、数据中心能源、黑洞、材料建模、医疗传感器等有研究价值主题。"
  rejected_alternatives:
    - "TopHub 常驻"
    - "Folo 候选立即增加"
  route_reason: "同一页面也混有任命、获奖、校园活动、招生、媒体综述和行政信息；现有候选已有 Quanta、Science Magazine、MIT Press Reader 等更集中来源。"
  reuse_rule: "大学机构新闻有价值但不能替代科学深读或论文目录；具体研究主题出现时查原文和论文链接。"
  evidence_locator: "docs/ops/tophub-technology-education-closure-20260704.md#三科研与大学机构新闻; #Folo候选"
  next_action: "按需。"

- source_id: "科技 > 电商 小标签整体"
  current_route: task_triggered_cross_border_ecommerce_sources
  decision_status: recovered_from_pr
  selected_reason: "雨果网、AMZ123、知无不言、电商报、卖家之家等在跨境平台规则、侵权、税务、账号、物流、活动和卖家经验上有任务价值。"
  rejected_alternatives:
    - "新增电商节点"
    - "新增独立电商组"
    - "跨境电商 / 亚马逊 / Shopee / TikTok Shop 宽追踪"
    - "RSSHub 承接平台快讯、活动、报告下载和论坛榜单"
  route_reason: "用户当前没有持续经营跨境店铺项目。中国企业全球化由 `36氪出海｜热门推荐`承担结构观察；跨境店铺运营来源全部按任务调用。"
  reuse_rule: "跨境电商实务只有真实店铺、具体站点、具体政策和明确日期形成责任时才临时追踪；最终状态回平台官方、监管、税务、法律与公司文件。"
  evidence_locator: "docs/ops/tophub-technology-ecommerce-closure-20260704.md#核心判断; #三个固定问题"
  next_action: "按任务。"
```

---

## 6. 数码与酷安：用户经验样本不等于常驻热榜

```yaml
- source_id: "科技 > 数码 小标签整体"
  current_route: no_new_tophub_with_folo_review_candidates
  decision_status: recovered_from_pr
  selected_reason: "数码目录中确实有真实使用经验、硬件维修、设备方法论和专业购买入口。"
  rejected_alternatives:
    - "新增 Engadget / 数字尾巴 / 中关村在线 / Tom's Guide 等门户"
    - "用硬件新闻流替代 The Register"
    - "腕表、投影、充电、无人机等垂直站常驻"
  route_reason: "本标签不产生 TopHub 增删。消费电子门户已过剩；真正值得保留的是按需专业入口和少数作者型 Folo 复查候选。"
  reuse_rule: "数码源分层：门户热榜、作者经验、维修拆解、专业评测、垂直设备、促销。常驻只在补出系统缺口时成立。"
  evidence_locator: "docs/ops/tophub-technology-digital-closure-20260704.md#核心结论"
  next_action: "无 TopHub 动作；Folo 复查作者。"

- source_id: "夜航船夫 / Slot 4 / jax"
  current_route: folo_review_or_on_demand_author_sources
  decision_status: recovered_from_pr
  selected_reason: "夜航船夫与 Obsidian、OpenCode、iCloud、Google Calendar、Rime、NAS、本地 AI 和工具生活经验高度贴近；Slot 4 是真实动手型硬件与家庭网络来源；jax 有 Agent、WWDC、OpenClaw 和浏览器工具项目。"
  rejected_alternatives:
    - "立即加入多个作者"
    - "把作者源压成 TopHub"
  route_reason: "夜航船夫进入 Folo 高优先级观察候选；Slot 4 条件候选；jax 因 Claude Code 比例、重叠和单一技术浪潮风险，按需访问。"
  reuse_rule: "作者型设备 / 工具源看长期关系和真实阅读，不看单篇命中；风险是阅读债和主题偏移。"
  evidence_locator: "docs/ops/tophub-technology-digital-closure-20260704.md#三作者型节点"
  next_action: "Folo 复查或按需。"

- source_id: "酷安小标签整体"
  current_route: on_demand_user_experience_samples
  decision_status: recovered_from_pr
  selected_reason: "酷安能提供媒体榜单难以获得的真实用户体验、系统适配、设备反馈、长期使用感受和社区争议。"
  rejected_alternatives:
    - "酷安今日热门常驻"
    - "酷安评论榜常驻"
    - "酷安优惠话题常驻"
    - "酷图榜进视觉与自然"
    - "酷安进 Folo"
  route_reason: "五个节点全不常驻。全站榜单混入品牌阵营、跑分争论、个人生活、疾病、情感、工作和优惠信息。酷安的价值只在具体设备、系统版本、App、影像、性能或购买问题中显现。"
  reuse_rule: "酷安作为用户经验样本，不作为事实裁决。至少核对设备型号、系统版本、使用时间、测试条件、多个独立反馈和官方日志 / 专业评测。"
  evidence_locator: "docs/ops/tophub-technology-coolapk-closure-20260704.md#核心结论; #六酷安的正确使用方式; #七与现有节点的关系"
  next_action: "按需。"
```

---

## 7. Google 与 TechWeb：官方关系归 Folo，门户切片按需

```yaml
- source_id: "Google 小标签整体"
  current_route: no_new_tophub_google_folo_baseline_kept
  decision_status: recovered_from_pr
  selected_reason: "Google 有大量高质量官方源，但当前 Folo 已经建立 Google Developers Blog、Google DeepMind News、Google AI Developers 三条一手关系。"
  rejected_alternatives:
    - "Google AI Blog 常驻"
    - "Google Developers Blog 在 TopHub 重复"
    - "Official Google Blog / Cloud Blog / 9to5Google 常驻"
    - "Google / Gemini / Android / Google Cloud 宽追踪"
  route_reason: "不新增任何 Google TopHub 节点。Google 官方持续关系由 Folo 承担；TopHub 不重复官方源。陈旧、重复、宽泛、营销或条件任务源按需。"
  reuse_rule: "官方生态来源先检查是否已有 Folo 一手关系；已有时不在 TopHub 重复。只有具体框架、产品、漏洞或 API 进入项目后才精确追踪。"
  evidence_locator: "docs/ops/tophub-technology-google-closure-20260704.md#核心结论; #十四三个固定问题"
  next_action: "无 TopHub 动作；Folo 复查候选。"

- source_id: "Google Online Security Blog / Google Testing Blog / Digital Inspiration"
  current_route: folo_review_or_conditional_candidates
  decision_status: recovered_from_pr
  selected_reason: "Google Online Security Blog 补 AI Agent、浏览器、Android、开源供应链和基础设施安全的一手技术；Google Testing Blog 补测试、review、小 PR、软件工程方法；Digital Inspiration 补 Gmail、Sheets、Apps Script 与 Workspace 自动化教程。"
  rejected_alternatives:
    - "立即加入 TopHub"
    - "立即加入 Folo"
  route_reason: "它们是连续技术关系或常青知识库，不是公共热榜。7 月 17 日复查 RSS、更新频率、阅读率和与现有 darkreading、Krebs、Google 源的重复度。"
  reuse_rule: "安全和工程方法源看是否真正改变项目实践；Workspace 自动化源只有近期多次命中真实任务时再进观察期。"
  evidence_locator: "docs/ops/tophub-technology-google-closure-20260704.md#三GoogleOnlineSecurityBlog; #四GoogleTestingBlog; #七DigitalInspiration"
  next_action: "Folo 复查。"

- source_id: "TechWeb 小标签整体"
  current_route: on_demand_traditional_tech_portal
  decision_status: recovered_from_pr
  selected_reason: "TechWeb 可用于搜索某家公司或产品的中文门户报道，比较传统科技门户如何组织产业事件。"
  rejected_alternatives:
    - "TechWeb 今日焦点 / 精华 / 推荐 / 滚动 / 业界常驻"
    - "TechWeb 外电编译作为外媒替代"
    - "TechWeb 原创报道进入 Folo"
    - "TechWeb 宽追踪"
  route_reason: "10 个节点全不常驻。TechWeb 主要是传统科技门户多入口分发，同一批 AI、手机、汽车、公司、融资和政策稿件在多栏目重复；原创报道是最优先按需页，但仍不足以常驻。"
  reuse_rule: "传统科技门户先看是否只是焦点、精华、推荐、滚动、业界的同池切片；原创栏目也要检查采访对象、原始文件、数据方法和是否只是公开消息重包装。"
  evidence_locator: "docs/ops/tophub-technology-techweb-closure-20260704.md#核心结论; #十一TechWeb在系统中的正确位置; #十二三个固定问题"
  next_action: "按需。"
```

---

## 8. 本补充 ledger 的复用规则

1. AI 不靠总入口解决，回到实际发生层：科技、开发、科学、数据结构、精确追踪和少量作者。
2. App 发现不等于安装责任；App Store 排行是市场样本，不是质量判断。
3. 快讯只负责发现，最终状态回原始文件、公司公告、监管、官方博客或项目原文。
4. 限免是交易窗口，不是长期关系；官方商店优先，聚合只作发现。
5. 教育、电商、工程、NOI、Google、Android、SEO、Workspace 自动化都需要具体任务触发。
6. 数码和酷安最有价值的是具体设备 / 系统 / 故障 / 长期体验样本，不是全站热榜。
7. 官方生态源如果已经在 Folo 建立一手关系，TopHub 不重复。
8. 传统科技门户多栏目切片不构成多个来源；原创也要验采访、文件、数据和方法。
9. 宽泛追踪词一律不加；只追踪具体对象、具体版本、具体项目、具体政策、具体设备或具体事件。

---

## 9. 仍需继续追回

科技大类已补入：

- 商业叙事与全球化来源；
- 科学、科普、企业基础设施、Apple、汽车、报告与设备市场结构；
- 当前科技雷达来源；
- AI、App、快讯、限免、教育、电商、数码、酷安、Google、TechWeb 等任务触发与高频流来源。

继续待办：

- HelloGitHub 与开发目录之间的关系；
- Jandan / Chouti / Coolapk 后续相邻娱乐 / 社区边界（若 PR 中还有对应科技 closure）；
- AI reassessment 六批的让位关系；
- 科技补充 ledger 后续等用户本地提交 master checklist 与 coverage audit 后，再统一登记。
