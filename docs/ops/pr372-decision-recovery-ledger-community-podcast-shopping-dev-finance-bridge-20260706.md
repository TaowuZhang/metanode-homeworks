# PR #372 判断链追回：社区播客购物、开发与财经闭环桥接（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中三份目录审计 / 最终收口文件中仍有独立价值的判断链。

纳入文件：

- `docs/ops/tophub-community-and-podcast-directory-audit-20260704.md`
- `docs/ops/tophub-development-directory-audit-20260704.md`
- `docs/ops/tophub-finance-final-closure-20260705.md`
- 已有关联专项 ledger：社区、购物、开发目录、财经 page-by-page、Folo 早期操作链等。

边界说明：

- 本文件不重复已完成的逐页 ledger，只补桥接判断：早期混合目录如何拆到后续专项；哪些候选为何暂不执行；哪些平台总类为什么不能直接进入个人路由。
- 文件中出现的 73、84 等数量是阶段快照；当前真实 TopHub 基线以 `资源/雷达/TopHub 全部订阅顺序.md` 的 80 项为准。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 社区目录：论坛与地方社区仓库，不是生活入口

```yaml
- source_id: "TopHub 社区目录整体"
  current_route: forum_and_local_community_repository_not_single_life_group
  decision_status: recovered_from_pr
  selected_reason: "社区目录约 235 个节点、20 页，覆盖虎扑、贴吧、水木、北大未名、吾爱破解、安全社区、开发者社区、V2EX、LINUXDO、汽车论坛、地方门户、高校论坛、旅游、数码、汽车等主题。"
  rejected_alternatives:
    - "把社区目录整体加入生活与社区"
    - "因为论坛有群体真实经验就常驻"
    - "把论坛热榜当公共生活事实源"
  route_reason: "社区目录不是同质社区，而是论坛与地方社区仓库。不同节点提供公共情绪、工具经验、地方生活、专业讨论或持续关系，必须分开路由。"
  reuse_rule: "论坛名气和目录归类不等于个人价值；先判断它提供公共情绪、工具经验、地方生活还是长期关系。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#一、社区目录不是一个同质社区"
  next_action: "已由社区专项 ledger 承接。"
```

---

## 2. 高噪声公共讨论：按需看群体注意力，不进日常首页

```yaml
- source_id: "贴吧 / 虎扑 / 知乎话题榜 / 虫部落热门讨论"
  current_route: on_demand_public_discussion_samples
  decision_status: recovered_from_pr
  selected_reason: "这些来源能显示某些群体在谈什么。"
  rejected_alternatives:
    - "加入生活与社区"
    - "加入公共温度"
    - "进入 Folo"
    - "作为事实源"
  route_reason: "标题党、重复话题、性别与关系内容、投资晒单、猎奇视频和情绪争论占比高；知乎话题榜还出现同一问题重复成多个条目，节点聚合质量不稳定。"
  reuse_rule: "高噪声论坛热榜只在需要观察某个群体注意力时按需打开；不形成日常首页和事实判断。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#社区目录当前样本"
  next_action: "按需。"
```

---

## 3. 专业与工具社区：问题触发，不整榜订阅

```yaml
- source_id: "吾爱破解 / V2EX / LINUXDO / 开发者与安全社区"
  current_route: task_triggered_professional_community_search
  decision_status: recovered_from_pr
  selected_reason: "V2EX、LINUXDO 和其他开发者 / 安全社区可能比综合社区更贴近实际技术使用。"
  rejected_alternatives:
    - "因为属于开发者社区就自动订阅"
    - "把吾爱破解放进个人常用首页"
    - "把专业社区热榜放入 Folo"
  route_reason: "吾爱破解当前热帖和精品软件区混有绿色版软件、破解、反调试、下载器和绕过授权内容；V2EX、LINUXDO 应等明确技术需求出现后再审查具体节点。专业社区也要避免整榜成为未读流。"
  reuse_rule: "解决技术问题时直接搜索具体社区和关键词；只有稳定低噪声、周频或明确长期关系成立，才考虑常驻。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#专业与工具社区"
  next_action: "按需。"
```

---

## 4. 经验型社区：豆瓣小组当前更稳定，马蜂窝是旅行候选

```yaml
- source_id: "经验型社区"
  current_route: life_experience_candidates_not_forum_expansion
  decision_status: recovered_from_pr
  selected_reason: "汽车之家、马蜂窝、豆瓣小组、高校论坛与地方论坛都有生活经验材料。马蜂窝热门游记能提供具体路线、城市、徒步与长途旅行记录，豆瓣小组当前仍比综合论坛更稳定。"
  rejected_alternatives:
    - "新增汽车之家论坛精选"
    - "新增高校论坛或地方论坛"
    - "把马蜂窝立即加入"
  route_reason: "汽车之家混有营销、提车作业模板和旧式论坛内容；高校论坛高度依赖特定社区语境；马蜂窝作为旅行经验候选更清楚，但仍需等生活与社区出现旅行经验缺口或具体旅行任务。"
  reuse_rule: "生活经验源要看是否提供可复用经验、低噪声结构和明确任务关系；论坛身份本身不够。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#经验型社区; #当前决定"
  next_action: "马蜂窝后续已进入生活与社区。"
```

---

## 5. 播客目录：发现榜和持续收听关系分开

```yaml
- source_id: "TopHub 播客目录整体"
  current_route: podcast_repository_split_discovery_vs_relationship
  decision_status: recovered_from_pr
  selected_reason: "播客目录约 95 个节点、8 页，混合喜马拉雅热门免费榜、头条榜、音乐榜、飙升榜、有声小说、相声、儿童故事、助眠音乐、广播剧、荔枝单集热榜、iTunes Store 中国区播客 Top100，以及 UX Coffee、博物志、疯投圈、梁文道·八分等独立节目。"
  rejected_alternatives:
    - "把播客目录整体加入影游音乐"
    - "把喜马拉雅热门免费榜理解为播客策展"
    - "从播客榜单直接批量加入 Folo"
  route_reason: "播客目录本身仍是平台内容仓库。TopHub 适合放播客发现榜；Folo 适合放真正持续听的节目。"
  reuse_rule: "播客先分发现榜、平台综合热榜、单集热度、独立节目和历史档案。发现榜不等于节目关系。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#二、播客目录实测; #TopHub 与 Folo 的播客分工"
  next_action: "已由播客专项继续处理。"
```

---

## 6. iTunes Top100：播客市场注意力候选，不直接订阅

```yaml
- source_id: "iTunes Store 中国区播客 Top100"
  current_route: tophub_podcast_discovery_candidate_not_subscribed
  decision_status: recovered_from_pr
  selected_reason: "它是当时最接近真正播客发现榜的节点，覆盖商业、文化、英语学习、喜剧、新闻、投资、科技和个人成长等节目。"
  rejected_alternatives:
    - "立即加入影游音乐"
    - "把排名当质量"
    - "由榜单直接建立 Folo 订阅"
  route_reason: "它适合回答当前中文播客市场哪些节目正在获得注意力，但排名不等于质量，也不等于用户会长期收听；当时保留为候选，不立即加入。"
  reuse_rule: "播客榜单可用于市场发现；具体节目进入 Folo 前必须试听、核验更新状态和节目关系。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#播客目录的三种节点"
  next_action: "候选。"
```

---

## 7. 博物志：技术条件通过，仍需实际收听

```yaml
- source_id: "博物志"
  current_route: folo_candidate_verified_source_structure_waiting_listening_experience
  decision_status: recovered_from_pr
  selected_reason: "独立站可访问，第 246 期有明确发布时间、时长、完整节目简介、分章节时间戳、图片与延伸说明、前后集导航、独立 RSS 入口和 Typlog 托管归档；内容围绕博物馆展览、古籍、城市、历史建筑与地方饮食展开。"
  rejected_alternatives:
    - "确认 RSS 后立即加入 Folo"
    - "只因主题契合就写成长期节目关系"
    - "把 TopHub 内部详情页当独立节目站"
  route_reason: "博物志具备成为 Folo 稳定订阅源的技术条件与内容结构，且与城市、博物馆、旅行、文化、视觉和历史兴趣高度重合。但当前只完成来源验证，是否长期订阅仍以实际收听体验为准。"
  reuse_rule: "播客候选需要两步：先验证独立站、元数据和 RSS，再通过实际收听判断长期关系。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#博物志独立站验证; #当前播客候选排序"
  next_action: "试听后复查。"
```

---

## 8. UX Coffee / 喜马拉雅 / 荔枝：节目候选与平台热榜分开

```yaml
- source_id: "UX Coffee / 喜马拉雅综合榜 / 荔枝热榜"
  current_route: uxcoffee_trial_candidate_platform_charts_rejected_or_on_demand
  decision_status: recovered_from_pr
  selected_reason: "UX Coffee 与设计、产品、职业经验和创作实践相关；喜马拉雅和荔枝可反映平台音频内容热度。"
  rejected_alternatives:
    - "把喜马拉雅热门免费榜、头条榜、音乐榜、飙升榜当播客入口"
    - "把荔枝热榜当播客策展"
    - "立即加入 UX Coffee"
  route_reason: "喜马拉雅平台榜混有新闻、有声书、相声、儿童内容、睡眠音乐和网络小说，不能承担播客策展；荔枝热榜更接近单集热度，且情绪化内容占比较高；UX Coffee 需要先确认近期节目是否仍持续符合兴趣。"
  reuse_rule: "独立节目可试听；平台综合音频榜不等于播客来源。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#播客目录的三种节点; #当前播客候选排序"
  next_action: "UX Coffee 待试听；平台榜按需或排除。"
```

---

## 9. 购物目录：销售转化入口，不进入注意力流

```yaml
- source_id: "TopHub 购物目录整体"
  current_route: purchase_task_triggered_directory_not_attention_flow
  decision_status: recovered_from_pr
  selected_reason: "购物目录约 85 个节点，覆盖淘宝、京东、拼多多、什么值得买、羊毛线报、众筹、图书与信用卡。"
  rejected_alternatives:
    - "加入任何首页分组"
    - "订阅淘宝、天猫、京东或线报热榜"
    - "把销量、券后价和平台榜单当质量证明"
    - "用当当畅销图书榜补慢读与思想"
  route_reason: "购物目录主体是销售转化而非信息价值：标题服务即时下单；价格离不开账号、地区、会员、叠券和限时活动；健康、儿童、保健等商品风险与噪声混在一起；线报聚合不可读；过时节点缺乏维护。当当畅销图书榜虽相对清楚，但已有微信读书、豆瓣新书和书评等阅读节点。"
  reuse_rule: "购物信息只在具体购买需求后按商品搜索、比较和核验；购物榜单不进入常驻注意力流。"
  evidence_locator: "docs/ops/tophub-community-and-podcast-directory-audit-20260704.md#三、购物目录实测; #购物目录的路由原则"
  next_action: "已由购物专项完整闭环。"
```

---

## 10. 开发目录早期 audit：技术目录混杂，不能整体进科技雷达

```yaml
- source_id: "TopHub 开发目录早期 audit"
  current_route: mixed_development_repository_split_by_use
  decision_status: recovered_from_pr
  selected_reason: "开发目录约 290 个节点、25 页，覆盖 GitHub、GitOther、掘金、CSDN、博客园、Hugging Face、产品经理、周刊、安全、独立开发、人工智能、营销、大模型、思维导图和主机等。"
  rejected_alternatives:
    - "把整个开发目录加入科技雷达"
    - "订阅中文技术社区热榜"
    - "把产品经理和 AI 商业文章当开发源"
  route_reason: "目录混合开源项目发现、技术资讯、版本发布、中文社区热文、产品经理与 AI 商业文章、周刊、月刊、独立开发、安全、主机、营销与课程型内容。必须按节点用途判断。"
  reuse_rule: "开发来源先分开源发现、官方版本发布、社区热文、教程、产品方法、周刊、安全、独立开发和课程营销。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#概况"
  next_action: "后续已由开发目录专项完整闭环。"
```

---

## 11. GitHub Trending：注意力发现，不是采用证据

```yaml
- source_id: "GitHub Trending Today / Weekly"
  current_route: on_demand_or_low_frequency_project_attention_discovery
  decision_status: recovered_from_pr
  selected_reason: "它最接近真实开源项目发现，能快速看到当前开发者注意力流向，对 coding agent、MCP、浏览器自动化、代码库记忆、开源基础设施等方向有发现价值。"
  rejected_alternatives:
    - "常驻 Trending Today"
    - "把星标和上榜解释成成熟可靠"
    - "用 Trending 代替 README / issues / release / license 审查"
  route_reason: "日榜波动很大，容易被短期传播、发布活动和 AI 热点支配；页面星标通常是累计星标，不代表当天增长；上榜不代表适合当前工作流。Weekly 比 Today 更适合低频项目发现候选。"
  reuse_rule: "Trending 只做项目发现。采用判断必须回仓库 README、issues、release、维护活跃度、许可证和代码。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#GitHub Trending Today / Weekly"
  next_action: "按需或低频候选。"
```

---

## 12. HelloGitHub / 湾区日报：主题契合仍需活跃度和通道验证

```yaml
- source_id: "HelloGitHub 月刊 / 湾区日报"
  current_route: hellogithub_low_frequency_project_curation_candidate_bay_area_requires_activity_check
  decision_status: recovered_from_pr
  selected_reason: "HelloGitHub 是低频策展节点：月度更新、有明确编号、聚焦开源项目、中文介绍降低初筛成本、不会制造太多未读压力；湾区日报主题与用户兴趣高度相关，包含完成的软件、File over app、写作即思考、开源项目商业模式、独立开发、软件长期维护、内部工具、技术组织与产品历史。"
  rejected_alternatives:
    - "立即订阅 HelloGitHub Folo"
    - "只因主题契合就订阅湾区日报"
    - "用长档案页面当活跃来源"
  route_reason: "HelloGitHub 与 Trending 分工不同：Trending 提供当前注意力，HelloGitHub 提供编辑筛选和解释；但早期仍列为优先验证候选。湾区日报页面呈现长档案、缺少明确日期，需确认独立站最近是否仍持续更新。"
  reuse_rule: "低频策展源要同时验证更新节奏、通道、内容结构和未读压力；主题契合不能替代活跃度核验。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#HelloGitHub 月刊; #湾区日报; #与现有科技雷达的关系"
  next_action: "后续已由开发目录专项承接。"
```

---

## 13. 中文技术社区热榜：有材料，但按问题搜索

```yaml
- source_id: "CSDN / 人人都是产品经理 / 掘金 / 开源中国 / PMCAFF"
  current_route: on_demand_chinese_tech_and_product_search_not_subscription
  decision_status: recovered_from_pr
  selected_reason: "这些来源能提供技术材料、产品概念、版本发布、工具实践、产品方法论和行业讨论。"
  rejected_alternatives:
    - "订阅 CSDN 今日头条热点"
    - "订阅人人都是产品经理日榜"
    - "订阅掘金全站 / AI / 工具热榜"
    - "订阅开源中国热门资讯"
    - "把 PMCAFF 当日报"
  route_reason: "CSDN 时间口径和聚合新鲜度不可靠；人人都是产品经理一页信息密度高，营销语言、课程、个人品牌和真实实践混在一起；掘金虽有具体技术材料，但标题膨胀、职场戏剧化、AI 题材重复和流量文案明显；开源中国版本发布最好回项目 release；PMCAFF 更像旧文章档案而非当前日报。"
  reuse_rule: "中文技术社区适合带着具体问题和关键词搜索，不适合整榜订阅。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#CSDN 今日头条热点; #人人都是产品经理; #掘金; #开源中国; #PMCAFF / 咖啡日报"
  next_action: "按需。"
```

---

## 14. 财经最终闭环：缺口不是快讯，是结构、工具和边界

```yaml
- source_id: "TopHub 财经目录最终收口"
  current_route: finance_page_by_page_closure_execution_bridge
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "财经目录 324 个节点、27 页已逐项判断，剩余 0。最终等量替换已执行：取消财新点击排行榜、评论排行榜，加入财新首页推荐和日经中文每日最新。"
  rejected_alternatives:
    - "继续增加财经快讯、7×24、荐股、龙虎榜、资金流、盘前盘后和市场复盘"
    - "新增宽泛财经追踪器"
    - "把财经信息流当投资动作生成器"
  route_reason: "财经目录不缺新闻、快讯或市场意见；真正缺的是日本与东亚产业 / 供应链 / 货币政策视角、仍在更新的财新编辑入口、长期投资和行为尺度、官方披露与专业研究工具、Web3 技术 / 安全 / 监管与币价交易流的分离。"
  reuse_rule: "财经源先看它补结构缺口还是制造交易噪声；事实与行动回官方披露、监管、统计、公司原始材料和用户确认。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#审核完成; #最终判断; #TopHub 最终执行结果; #明确不做"
  next_action: "已由财经专项完整闭环。"
```

---

## 15. 财经来源分工：数据与结构 12 项各有角色

```yaml
- source_id: "数据与结构 财经后顺序"
  current_route: tophub_data_structure_finance_role_split
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "最终数据与结构顺序把原始数据、国内编辑入口、东亚视角、国际财经、调查节目、出海产业、汽车、科技硬件与消费产业分开。"
  rejected_alternatives:
    - "用第一财经其他栏目扩容"
    - "用联合早报替代日经或 FT 中文"
    - "用经济观察网或 21财经数读常驻"
    - "把财新首页名称本身当精选充分证据"
  route_reason: "国家统计局两个入口承担数据底稿；财新首页承担中国政策、商业、金融与调查，但体量较大，需要真实使用观察噪声；日经承担日本与东亚供应链、日元、央行和产业政策；FT 中文承担国际经济与商业；央视栏目承担调查；36氪出海、第一财经汽车、Counterpoint、华丽志分别补全球化、汽车、终端和消费产业。其他财经源保持任务检索。"
  reuse_rule: "数据与结构按功能位保留，不按财经品牌越多越好。新增源必须说明替代谁或补哪个结构缺口。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#当前“数据与结构”顺序; #来源分工"
  next_action: "观察财新首页真实噪声。"
```

---

## 16. 有知有行：Folo 候选选全部，不选单栏目

```yaml
- source_id: "有知有行｜全部"
  current_route: folo_review_candidate_long_term_investing_and_behavior_scale
  decision_status: recovered_from_pr
  selected_reason: "它承担长期投资、基金与资产配置、保险边界、组合和再平衡、投资行为与数据解释。"
  rejected_alternatives:
    - "立即加入 Folo"
    - "选择有知有行｜知行黑板报"
    - "选择有知有行｜知行小酒馆"
    - "把平台组合调整当用户投资指令"
  route_reason: "三者存在重叠，`全部`最完整，不会把单一栏目误当整体角色；执行时间放在 2026-07-17 Folo 统一复查中，不因财经审核擅自修改实际订阅。"
  reuse_rule: "投资教育类来源若进入 Folo，要选能覆盖完整角色的入口；但内容只能作理解和行为尺度，不生成投资动作。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#Folo 最终建议"
  next_action: "2026-07-17 复查。"
```

---

## 17. 财经工具与 Web3：工具进沃壤，交易流不进常驻

```yaml
- source_id: "财经日历 / 官方披露 / Web3"
  current_route: tools_and_professional_sources_on_demand_not_feeds
  decision_status: recovered_from_pr
  selected_reason: "华尔街见闻财经日历、巨潮资讯、交易所、中国货币网、FRB、国家金融与发展实验室、中国会计视野网、Foresight News 文章、慢雾技术研究、CoinDesk Today 等各有工具或专业来源价值。"
  rejected_alternatives:
    - "把财经日历作为每日订阅"
    - "把官方披露工具制造阅读义务"
    - "订阅 Web3 快讯、价格预测、巨鲸、杠杆、爆仓、支撑位流"
    - "选择慢雾公司新闻而不是技术研究 / 漏洞披露"
  route_reason: "财经日历是查询工具，不是内容流；官方披露与研究工具发生具体公司、规则、货币政策、债券或市场制度问题时直接查询；Web3 要把技术、安全、监管和行业结构长文，与币价交易流彻底分开。"
  reuse_rule: "工具链接写沃壤，不制造未读；Web3 结构 / 安全源按需查，交易和价格流排除。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#财经日历; #Web3; #官方与研究工具组"
  next_action: "按需。"
```

---

## 18. 本补充 ledger 的复用规则

1. 社区目录是论坛与地方社区仓库，不是一个统一生活入口。
2. 高噪声论坛热榜只能按需观察群体注意力，不进日常首页。
3. 专业社区按具体技术问题搜索，不整榜订阅。
4. 经验型社区要看可复用经验与低噪声结构；马蜂窝这种候选要等真实旅行 / 生活缺口。
5. 播客目录要分发现榜、平台综合热榜、单集热度、独立节目和历史档案。
6. 播客进入 Folo 前要先试听，RSS 可用和独立站完整只是候选条件。
7. 购物目录是销售转化入口，不进入注意力流。
8. 开发目录混杂，必须按节点用途判断，不能整体进科技雷达。
9. GitHub Trending 是注意力发现，不是采用证据。
10. HelloGitHub / 湾区日报这类主题契合源仍需验证活跃度、通道、内容结构和未读压力。
11. 中文技术社区热榜按具体问题搜索，不常驻。
12. 财经目录闭环后，新增只补结构缺口，不加快讯和交易噪声。
13. 有知有行是复查候选，不是已执行 Folo 变更。
14. 财经工具写沃壤，不制造未读；Web3 结构 / 安全源和交易流必须分开。

---

## 19. 仍需继续追回

社区播客购物、开发与财经闭环桥接已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内剩余可查项可继续检查是否还有 `tophub-*final-closure`、`*-platform-execution-verified`、`*-directory-audit` 未被 ledger 覆盖；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
