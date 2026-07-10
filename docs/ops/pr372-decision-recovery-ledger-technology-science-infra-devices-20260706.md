# PR #372 判断链追回：科技科学、基础设施与设备结构来源（2026-07-06）

## 0. 边界

本文件是 `科技` 大类的补充 ledger，追回 PR #372 中已闭环的科学、科普、企业基础设施、Apple、汽车与报告来源判断。

纳入文件：

- `docs/ops/tophub-technology-popular-science-closure-20260704.md`
- `docs/ops/tophub-technology-science-closure-20260704.md`
- `docs/ops/tophub-technology-general-tech-closure-20260704.md`
- `docs/ops/tophub-technology-apple-closure-20260704.md`
- `docs/ops/tophub-technology-automotive-closure-20260704.md`
- `docs/ops/tophub-technology-reports-closure-20260704.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 科技雷达互补功能总规则

```yaml
- source_id: "科技雷达互补功能总规则"
  current_route: keep_distinct_functions_not_more_tech_feeds
  decision_status: recovered_from_pr
  selected_reason: "科技雷达需要互补功能：中文公共科普、中文解释型科学、国际科学新闻、企业基础设施与技术运营、官方行动入口、产业结构和终端市场研究。"
  rejected_alternatives:
    - "多订几个科学新闻源"
    - "多订几个苹果传闻站"
    - "多订汽车门户和车型榜"
    - "多订报告库和白皮书目录"
    - "把论文 TOC 放进日常流"
  route_reason: "科技目录太宽，不能按数量扩张。每新增或替换一个来源，都必须补出当前系统缺少的角色，并说明为什么旁边来源不能替代。"
  reuse_rule: "以后遇到科技源，先问它承担哪一种功能；如果只是站内切片、传闻、快讯、报告目录、论文洪流、促销或高频门户，就按需或排除。"
  evidence_locator: "popular-science / science / general-tech / apple / automotive / reports closure files"
  next_action: "作为科技同族复用规则。"
```

---

## 2. 科普中国、果壳与 Science Magazine

```yaml
- source_id: "科普中国网｜头条"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "比 `热点排行` 更适合承担中文公共科普入口：更新更完整，同时连接科学知识、现实风险、日常生活、心理健康、环境和技术，标题克制。"
  rejected_alternatives:
    - "科普中国网｜热点排行"
    - "科普中国网｜最新滚动"
    - "科普中国｜今日辟谣文章作为替代"
  route_reason: "替换 `热点排行` 进入科技雷达。它负责发现当前值得理解的问题；`今日辟谣文章` 放生活与社区，负责具体流言和错误认知核验。"
  reuse_rule: "同一机构内要区分发现入口、辟谣入口、滚动入口；旧稿排行和活动滚动不替代公共科普头条。"
  evidence_locator: "docs/ops/tophub-technology-popular-science-closure-20260704.md#一科普中国内部比较"
  next_action: "已在科技 final closeout 中替换。"

- source_id: "果壳｜科学人"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "用中文互联网表达方式，把科学、健康、技术争议、消费现象和公共讨论转成可进入的解释文章。它与科普中国头条不同，更擅长网络语境、文化现象、争议和解释性标题。"
  rejected_alternatives:
    - "果壳｜每日精选作为第二个果壳常驻"
    - "果壳｜吃货研究所"
    - "果壳｜美丽也是技术活"
  route_reason: "继续保留科技雷达。它是果壳唯一综合解释入口；不再同时保留同站更宽的每日精选。"
  reuse_rule: "同一站点只保留功能最清楚的一条；医疗和健康文章必须回研究或指南核验。"
  evidence_locator: "docs/ops/tophub-technology-popular-science-closure-20260704.md#二果壳内部比较"
  next_action: "无；保留。"

- source_id: "Science Magazine｜Latest News"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "提供编辑后的国际科学新闻窗口：科学发现、科研制度、环境、医学、天文、科学政策和现实影响；标题数量可控，不是原始论文洪流，也不依赖中文转载链条。"
  rejected_alternatives:
    - "果壳｜每日精选"
    - "Science News｜最新科学"
    - "Ars Technica｜科学与探索"
    - "NASA｜News"
    - "Scientific American / New Scientist / ScienceDaily / ScienceAlert / Phys.org 等常驻"
    - "Science / Nature 等论文 TOC 常驻"
  route_reason: "替换 `果壳｜每日精选`，仍放科技雷达。它补国际科学新闻，不是增加第四条科学源。Science News 可作为未来替代候选，不与其并存；论文目录按需。"
  reuse_rule: "科学信息流分四层：大众科学新闻、解释型深读、机构 / 论文原始目录、专题按需。TopHub 只收可标题扫描的编辑新闻窗口。"
  evidence_locator: "docs/ops/tophub-technology-science-closure-20260704.md#二TopHub推荐节点; docs/ops/tophub-technology-popular-science-closure-20260704.md#核心结论"
  next_action: "已在科技 final closeout 中替换。"
```

---

## 3. 科学深读与论文目录为什么不常驻

```yaml
- source_id: "Quanta Magazine / 科学空间 / Science Magazine Podcast"
  current_route: folo_review_candidates_not_tophub
  decision_status: recovered_from_pr
  selected_reason: "Quanta 提供数学、物理和基础生命科学长线解释；科学空间提供中文原创数学与机器学习推导；Science Magazine Podcast 可作为科学音频漫游。"
  rejected_alternatives:
    - "直接加入 TopHub 科技雷达"
    - "立即加入 Folo"
  route_reason: "这些来源标题不能替代正文，适合完整阅读或收听，而不是标题扫读。需要在 2026-07-17 复查 RSS、更新频率、实际阅读 / 收听和负担后再决定。"
  reuse_rule: "深读型科学源先复查真实阅读关系；不因质量高就直接扩张未读队列。"
  evidence_locator: "docs/ops/tophub-technology-science-closure-20260704.md#三Folo高价值候选"
  next_action: "Folo 复查。"

- source_id: "Science / Nature / 专业论文目录"
  current_route: on_demand_original_research_directories
  decision_status: recovered_from_pr
  selected_reason: "质量高，适合研究具体论文、学科或项目。"
  rejected_alternatives:
    - "TopHub 常驻"
    - "Folo 常驻"
  route_reason: "原始论文、TOC、专业期刊目录密度过大，不适合通用信息流；Science / AAAS 多节点互相重复，Nature 混合研究、Briefing、职业、政策、评论和出版事务。"
  reuse_rule: "论文目录按专题任务使用；日常发现先看编辑科学新闻，真正研究再回原始目录。"
  evidence_locator: "docs/ops/tophub-technology-science-closure-20260704.md#六学术期刊与原始论文目录"
  next_action: "按需。"
```

---

## 4. The Register 与企业基础设施视角

```yaml
- source_id: "The Register｜Latest"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "补出当前科技雷达缺少的企业 IT、云、数据库、基础设施、网络安全、漏洞、数据中心、芯片、软件供应链、系统管理员和开发者实际问题。"
  rejected_alternatives:
    - "爱范儿｜每日最新继续承担该位置"
    - "The Verge｜Today"
    - "TechCrunch｜Today"
    - "IT之家 / 爱范儿 / 消费电子门户"
    - "纯安全媒体作为替代"
  route_reason: "放科技雷达。它比 The Verge 更偏企业系统，比 TechCrunch 更少融资叙事，比中文消费科技更少新品与设备消费，也比纯安全媒体更能把安全放回软件、云和企业系统。"
  reuse_rule: "海外科技媒体不是只看知名度；要看它补哪种技术运行视角。企业基础设施和技术运营视角可以替代消费产品流。"
  evidence_locator: "docs/ops/tophub-technology-general-tech-closure-20260704.md#一最值得进入科技大类收口的节点"
  next_action: "已在科技 final closeout 中替换进入。"

- source_id: "硬件专业评测与工作站来源"
  current_route: on_demand_hardware_task_sources
  decision_status: recovered_from_pr
  selected_reason: "Puget Systems、igor´sLAB、TechPowerUp、UltrabookReview、Notebookcheck、Tom's Hardware 等在真实购买、性能、兼容性、AI 推理、GPU、内存、笔记本和工作站任务中有价值。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "Folo"
  route_reason: "它们是任务型专业源；不需要在没有购买、性能测试或硬件故障任务时占据常驻信息流。"
  reuse_rule: "硬件源按任务调用：先明确预算、设备、工作负载、兼容性问题，再选评测源。"
  evidence_locator: "docs/ops/tophub-technology-general-tech-closure-20260704.md#核心结论; #按需专业来源"
  next_action: "按需。"
```

---

## 5. Apple 官方行动入口、精确追踪与作者型关系

```yaml
- source_id: "Apple 支持｜更换和维修扩展计划"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "39 个苹果节点里最明确、最少噪声、最具行动价值的官方入口。它补的是设备所有者可能错过、但一旦命中就影响维修、费用、安全和处理动作的服务计划。"
  rejected_alternatives:
    - "Apple Newsroom 常驻"
    - "Apple Developer 官方新闻常驻"
    - "MacRumors / 9to5Mac / 苹果传闻站常驻"
    - "Apple 相关 Folo 立即扩容"
  route_reason: "放科技雷达。现有 Apple Silicon、macOS 追踪器和科技媒体能覆盖新闻与技术变化，但不能稳定替代官方维修与召回入口。"
  reuse_rule: "官方行动入口和官方新闻不同；一旦命中会产生动作的低噪声官方页，可以少量常驻。"
  evidence_locator: "docs/ops/tophub-technology-apple-closure-20260704.md#一建议新增Apple支持更换和维修扩展计划"
  next_action: "已在科技 final closeout 中新增。"

- source_id: "Apple Foundation Models"
  current_route: precise_tracker_candidate
  decision_status: recovered_from_pr
  selected_reason: "与用户当前 Mac、本地模型、Apple Silicon、自动化、Agent 和开发工作流直接相关；精确指向苹果端侧基础模型、开发框架、API、AFM Core、模型能力与应用接入。"
  rejected_alternatives:
    - "Apple Intelligence"
    - "苹果 AI"
    - "Apple Newsroom 常驻"
  route_reason: "它不是宽泛苹果 AI 新闻，而是具体技术框架与模型家族。建议新增精确追踪词，不启用通知、不自动刷新，低频按需检查。"
  reuse_rule: "追踪器只收具体技术框架、产品家族或项目；不收宽泛品牌词和营销词。"
  evidence_locator: "docs/ops/tophub-technology-apple-closure-20260704.md#三建议建立AppleFoundationModels精确追踪"
  next_action: "已在科技 final closeout 中列为追踪器。"

- source_id: "MacStories / MacTalk / Apple Developer / Swift.org"
  current_route: folo_review_or_on_demand_apple_sources
  decision_status: recovered_from_pr
  selected_reason: "MacStories 贴近 Mac 工作流、App、自动化、Shortcuts、MCP、AI 工具和独立软件生态；MacTalk 是中文独立作者、产品与 AI 工作流来源；Apple Developer 和 Swift.org 是官方开发者任务源。"
  rejected_alternatives:
    - "立即进入 TopHub"
    - "多个苹果作者站一起进入 Folo"
    - "Apple Developer 中英文同时订阅"
  route_reason: "作者型和官方开发者关系不适合 TopHub 标题热榜。MacStories 和 MacTalk 需 2026-07-17 复查 RSS、阅读率和重复度；Apple Developer / Swift.org 只有在真实 Apple 平台开发任务出现时启用。"
  reuse_rule: "苹果生态来源分为官方行动、官方新闻、官方开发者、作者型工作流、传闻站、教程站。不同层不能互相替代。"
  evidence_locator: "docs/ops/tophub-technology-apple-closure-20260704.md#四Folo高优先级候选MacStories; #五Folo条件候选MacTalk; #六官方开发者来源"
  next_action: "Folo 复查或按需。"
```

---

## 6. 汽车产业结构与购车任务源

```yaml
- source_id: "第一财经｜汽车新闻"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "把汽车作为社会与产业系统观察：税制、国家标准、销量周期、供应链、利润、安全、法规、自动驾驶、消费政策和全球重组。"
  rejected_alternatives:
    - "IT之家｜智能汽车"
    - "Readhub｜汽车"
    - "财联社｜汽车同时常驻"
    - "中国汽车报 / 盖世汽车常驻"
    - "车型热榜、视频榜、论坛榜和品牌发布流"
  route_reason: "放数据与结构，不放科技雷达。目的不是追每一款车，而是观察中国制造、能源转型、消费政策、交通安全与全球产业竞争。财联社是高质量替代，不与第一财经并存。"
  reuse_rule: "汽车来源先拆四层：产业结构、技术路线、购车评测、车主经验。常驻只收结构层；购车和车型任务按需。"
  evidence_locator: "docs/ops/tophub-technology-automotive-closure-20260704.md#核心结论; #一为什么选择第一财经汽车新闻"
  next_action: "已在科技 final closeout 中新增。"

- source_id: "CnEVPost / Electrek / Ars Technica Cars / 购车评测与车主经验来源"
  current_route: conditional_folo_or_on_demand_auto_sources
  decision_status: recovered_from_pr
  selected_reason: "CnEVPost 连续跟踪中国新能源汽车产业；Electrek 适合全球电动交通与能源；Ars Cars 适合海外车型、监管与自动驾驶事故；38号、汽车之家论坛、V2EX 汽车、精真估等适合具体购车和使用问题。"
  rejected_alternatives:
    - "全部常驻"
    - "全部进入 Folo"
    - "建立汽车 / 新能源汽车 / 智能驾驶宽泛追踪"
  route_reason: "没有持续汽车研究或购车任务时，这些都会制造销量、车型、品牌、促销和个人经验未读压力。CnEVPost 只是未来条件候选，不进当前核心复查。"
  reuse_rule: "具体车型、预算、保险、维修、二手残值、自动驾驶事故或监管问题出现时，临时调用对应来源；追踪只收具体车型、法规、事故或项目。"
  evidence_locator: "docs/ops/tophub-technology-automotive-closure-20260704.md#二为什么不选择其他产业型来源常驻; #三购车与车型评测来源"
  next_action: "按需或未来条件复查。"
```

---

## 7. 报告与终端市场结构

```yaml
- source_id: "Counterpoint Research｜最新见解"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "持续把设备、芯片、出货量、终端形态和 AI 平台放入同一市场结构中；与用户长期关注的 Apple、vivo、手机厂商、Apple Silicon、AI 硬件、智能眼镜和消费电子市场直接相关。"
  rejected_alternatives:
    - "CBNData资讯 / 报告"
    - "艾瑞体系多个报告节点"
    - "QuestMobile 常驻"
    - "东方财富策略报告"
    - "国家金融与发展实验室 TopHub 节点"
    - "AMZ123 跨境报告"
    - "同时加入 Folo"
  route_reason: "放数据与结构。它不是泛报告目录、消费白皮书、券商策略或每日快讯，而是明确的科技终端市场研究入口。先由 TopHub 承担标题级发现，不同时加入 Folo。"
  reuse_rule: "报告源不能因名字叫报告就常驻；只有具备稳定独立数据能力、与用户项目长期相关、标题可筛选、且不与现有结构源重复时，才进入数据与结构。"
  evidence_locator: "docs/ops/tophub-technology-reports-closure-20260704.md#核心判断; #五科技终端与硬件市场"
  next_action: "已在科技 final closeout 中新增。"

- source_id: "商业报告与数据研究目录"
  current_route: on_demand_research_sources
  decision_status: recovered_from_pr
  selected_reason: "CBNData、艾瑞、QuestMobile、东方财富策略、国家金融与发展实验室、AMZ123 等在具体消费、行业、移动互联网、宏观金融、跨境项目中有任务价值。"
  rejected_alternatives:
    - "多个报告目录常驻"
    - "报告源进入 Folo"
  route_reason: "报告页同时存在市场研究、品牌白皮书、营销内容、券商策略、学术报告、乱码和陈旧入口。方法、样本、口径、合作方和报告目的必须逐份核验。"
  reuse_rule: "报告不是证据本身；它是研究起点和关键词地图。使用时检查方法、样本、口径、赞助和时间。"
  evidence_locator: "docs/ops/tophub-technology-reports-closure-20260704.md#报告目录的价值不在于每份都读; #15个节点路由总表"
  next_action: "按任务调用。"
```

---

## 8. 本补充 ledger 的复用规则

1. 科技雷达不是科技新闻越多越好，而是互补功能要清楚。
2. 科学源分层：公共科普、中文解释、国际科学新闻、深读、原始论文目录、专题任务。
3. 企业基础设施视角可以替代消费电子流；The Register 的价值在技术运行层，不在品牌知名度。
4. Apple 来源分层：官方行动、官方新闻、官方开发者、作者型工作流、传闻站、教程站。只有官方行动入口常驻，作者型进入复查。
5. 汽车来源分层：产业结构、技术路线、购车评测、车主经验。常驻只收结构层。
6. 报告源分层：独立数据能力、商业白皮书、策略报告、学术报告、任务目录。报告不是证据本身。
7. 追踪器只收具体对象，不收宽泛科学、Apple、汽车、芯片、AI、报告或市场词。

---

## 9. 仍需继续追回

科技大类已补入科学、科普、企业基础设施、Apple、汽车、报告与终端市场结构判断，但仍是 partial。

继续待办：

- 少数派、IT之家、极客公园、TechCrunch、The Verge 等现有科技雷达来源的同族判断；
- AI、App、快讯、限免、教育、电商、数码、酷安、Google、TechWeb 等 closure；
- 后续等用户本地提交 master checklist 与 coverage audit 后，再登记本文件。
