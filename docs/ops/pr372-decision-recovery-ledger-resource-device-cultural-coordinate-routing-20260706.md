# PR #372 判断链追回：资源总入口、设备生态与文化坐标路由（2026-07-06）

## 0. 边界

本文件追回 PR #372 中三个非 TopHub 目录型文件形成的判断链：

1. `资源/雷达/设备生态关注说明.md`：设备生态与精确追踪词；
2. `成为/灯塔/文化公共谈话者候选.md`：文化公共谈话者、灯塔、坐标、节目与对话入口；
3. `资源.md`、`资源/README.md`、`资源/语境/_index.md`：资源总入口、资源沉积规则和语境 / 链接 / 领域 / 项目边界。

纳入文件：

- `资源/雷达/设备生态关注说明.md`
- `成为/灯塔/文化公共谈话者候选.md`
- `资源.md`
- `资源/README.md`
- `资源/语境/_index.md`
- `docs/ops/subscription-system-closeout-20260704.md`（用于校正 Apple Foundation Models 后续已创建的时间边界）

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：非目录型文件也要写路由，不写成收藏堆

```yaml
- source_id: "非 TopHub 目录型资源路由整体"
  current_route: resource_routing_rules_not_source_expansion
  decision_status: recovered_from_pr
  selected_reason: "设备生态、文化公共谈话者、资源入口和语境索引都不是 TopHub 目录页，但它们定义了用户为什么长期关注某些设备、人物、节目、工具、材料，以及这些材料应在雷达、灯塔、资源、语境、领域、项目或 docs/ops 中怎样居留。"
  rejected_alternatives:
    - "把这些文件当普通链接收藏"
    - "把设备关注误写成品牌站队"
    - "把人名候选直接写入灯塔"
    - "把外部活空间完整镜像进 Git"
    - "把只有 URL 的候选写成已形成使用语境"
  route_reason: "这些文件共同补的是路由规则：设备生态决定精确追踪词；文化谈话者文件决定人物、节目和对话入口如何进入灯塔 / 坐标 / 按需；资源 README 决定外部材料何时沉到 GitHub；语境索引决定链接候选何时上升为使用语境。"
  reuse_rule: "以后遇到非目录型材料，先问它是在定义长期关注、具体来源、人物坐标、使用方法、外部链接、项目证据还是操作审计；不要用一个‘资源’标签吞掉。"
  evidence_locator: "资源/雷达/设备生态关注说明.md; 成为/灯塔/文化公共谈话者候选.md; 资源/README.md; 资源/语境/_index.md"
  next_action: "作为非目录型资源路由规则。"
```

---

## 2. 设备生态：不是品牌站队，而是真实使用关系

```yaml
- source_id: "设备生态关注整体"
  current_route: radar_device_ecosystem_attention
  decision_status: recovered_from_pr
  selected_reason: "用户当前真实设备生态不是单一品牌：手机、平板、手表、耳机存在 vivo 设备，电脑是 MacBook。vivo 办公组件的价值来自把手机、平板、穿戴设备与电脑工作流连接起来。"
  rejected_alternatives:
    - "把 vivo 与 Apple 记录成二选一"
    - "把关注 vivo 理解成一次性新闻兴趣"
    - "把关注 Apple 理解成放弃 vivo"
    - "把设备生态关注写成品牌粉丝关系"
  route_reason: "文件明确：用户认可 vivo 会继续关注新产品、新能力与公司动态，也认可 MacBook 的使用价值，会继续关注 Apple 的开发者大会、硬件产品与平台变化。两者不是二选一，也不应被记录成品牌站队。"
  reuse_rule: "设备相关来源要从真实设备、跨端工作流、维修/兼容/系统更新、开发能力和购买/升级任务出发，不从品牌忠诚出发。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#已确认的现实使用"
  next_action: "继续作为设备生态雷达。"
```

---

## 3. 取消 vivo 宽词：取消的是失真入口，不是取消关注

```yaml
- source_id: "TopHub `vivo` 宽关键词追踪"
  current_route: retired_broad_keyword_not_attention_retired
  decision_status: recovered_from_pr
  selected_reason: "原本试图用 `vivo` 追踪 vivo 公司、设备、办公组件、跨端协作和 Agent 能力。"
  rejected_alternatives:
    - "继续使用 `vivo` 宽词"
    - "把取消 `vivo` 追踪解释成不再关注 vivo"
    - "用品牌词承担具体产品与跨端 Agent 追踪"
  route_reason: "`vivo` 宽词误命中 `in vivo` 生物医学论文、Vivobook、Survivorship Bias 等字母片段；品牌相关结果又主要由促销、价格、爆料和硬件评测占据，真正关心的办公组件、跨端协作与 Agent 能力被淹没。"
  reuse_rule: "品牌词如果噪声压倒真实关注对象，就取消品牌词，改追正式产品名、组件名、系统名、版本名或事件名。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#TopHub `vivo` 追踪的取消"
  next_action: "不恢复 vivo 宽词。"
```

---

## 4. 关键词设计：完整专有名优先，不追宽词

```yaml
- source_id: "设备关键词设计原则"
  current_route: precise_named_tracking_terms
  decision_status: recovered_from_pr
  selected_reason: "TopHub 表现出宽松文本匹配特征，因此优先使用完整、可识别的专有名称。"
  rejected_alternatives:
    - "品牌名"
    - "通用技术词"
    - "单个普通名词"
    - "AI"
    - "Agent"
    - "Apple"
    - "MacBook"
    - "vivo"
  route_reason: "合适颗粒度是平台、组件、事件和临时型号：操作系统、AI 平台、芯片平台、有正式名称的办公 / 跨端 / 助手产品、开发者大会、发布会，以及准备购买、升级或研究某条产品线时的临时型号词。"
  reuse_rule: "追踪词要能闭环、能区分、能在结果里看出真实对象。宽词不追；型号和发布会只在窗口期临时追。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#关键词设计原则"
  next_action: "继续用精确词。"
```

---

## 5. 当前长期设备追踪词：各自分工

```yaml
- source_id: "macOS"
  current_route: keep_precise_environment_radar
  decision_status: recovered_from_pr
  selected_reason: "用户每天使用 MacBook；macOS 结果虽多，但绝大多数与 Mac 平台直接相关，包括系统更新、Beta、兼容性、系统故障、桌面工具、安全、恶意软件、终端、Safari 扩展、Mac 专用软件和系统功能变化。"
  rejected_alternatives:
    - "因为 999+ 条结果就取消"
    - "用 Apple 或 MacBook 宽词替代"
    - "用 Apple Silicon 替代 macOS"
  route_reason: "macOS 是系统与应用生态词，不是品牌宽词。它定位为低频环境雷达，不启用通知，不形成未读债务。需要具体版本问题时临时搜索版本词，而不是再建永久追踪。"
  reuse_rule: "当前使用环境相关系统词可长期保留；但只作为低频雷达，具体问题用具体版本词。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#现有 macOS 追踪审计"
  next_action: "保留。"

- source_id: "Apple Silicon"
  current_route: keep_precise_chip_platform_radar
  decision_status: recovered_from_pr
  selected_reason: "关注芯片架构、本地 AI、MLX / Metal、容器、虚拟化、性能、硬件路线、Apple Silicon 上的本地 LLM 和开发能力。"
  rejected_alternatives:
    - "用 macOS 取代 Apple Silicon"
    - "用 Apple 宽词取代"
    - "用 Apple Intelligence 泛词取代"
  route_reason: "Apple Silicon 与 macOS 职责不同：macOS 看系统、应用、兼容、安全和日常故障；Apple Silicon 看芯片平台、本地 AI、容器、虚拟化、性能和硬件路线。两者同时保留，不替换。"
  reuse_rule: "同属 Apple 生态也可分层：系统体验、芯片平台、开发者大会、基础模型、具体设备型号分别追踪或按需。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#当前已确认创建并保留; #第二批关键词实测"
  next_action: "保留。"

- source_id: "vivo办公套件"
  current_route: keep_precise_cross_device_workflow_tracking
  decision_status: recovered_from_pr
  selected_reason: "结果数量可控，直接涉及 vivo 手机、平板与 Windows / Mac 的互联、文件传输、投屏、备份恢复、远程控制、iPad 与 Apple 生态连接、用户体验、故障、版本更新和小 V Claw。"
  rejected_alternatives:
    - "用 vivo 宽词替代"
    - "用 OriginOS 长期替代"
    - "只按新闻兴趣处理"
  route_reason: "它最契合真实设备关系和跨端工作流，已创建长期追踪。"
  reuse_rule: "跨设备能力要追具体组件名，而不是追品牌名。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#vivo办公套件：通过，已创建"
  next_action: "保留。"

- source_id: "小 V Claw"
  current_route: keep_precise_cross_device_agent_tracking
  decision_status: recovered_from_pr
  selected_reason: "数量很小且高度集中，全部围绕小 V Claw PC 版、vivo X Fold6 与办公套件联动、手机向电脑下达任务、原子笔记、日程与跨设备智能体、Windows 与 Mac 适配。"
  rejected_alternatives:
    - "用 Agent 宽词替代"
    - "用 vivo 办公套件完全替代"
  route_reason: "它比 vivo办公套件更窄，几乎没有噪声，直接对应用户曾提到的 vivo 跨设备 Agent。"
  reuse_rule: "Agent 相关追踪优先用正式产品名，不用 Agent / AI / 大模型等宽泛词。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#小 V Claw：通过，已创建"
  next_action: "保留。"

- source_id: "蓝河操作系统"
  current_route: keep_precise_vivo_wearable_os_tracking
  decision_status: recovered_from_pr
  selected_reason: "中文正式名消除 Blue Robotics 的 BlueOS 同名污染，结果直接涉及 vivo WATCH、iQOO WATCH、蓝河操作系统 3、AIoT、蓝心智能、vivo 开发者大会、Rust 内核、开源、开发者竞赛和手表与 iPhone / vivo 手机连接。"
  rejected_alternatives:
    - "BlueOS"
    - "BlueOS 3"
    - "蓝河操作系统 3"
    - "vivo 宽词"
  route_reason: "用户实际使用 vivo 手表，因此与现实设备关系直接相关。中文总词可继续捕捉未来版本；具体版本词只在发布窗口临时使用。"
  reuse_rule: "有同名污染时优先使用中文正式名或更完整专有名；版本词按窗口期临时使用。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#蓝河操作系统：通过，已创建"
  next_action: "保留。"

- source_id: "Codex CLI"
  current_route: keep_local_agent_development_workflow_tracking
  decision_status: recovered_from_pr
  selected_reason: "它指向本地代理与开发工作流。"
  rejected_alternatives:
    - "Codex 宽词"
    - "AI Agent 宽词"
  route_reason: "设备生态文件把它列入当前已确认创建并保留；系统总分工中也将其归入精确追踪器。其价值是本地 CLI、权限、沙箱、MCP 和工作流变化，而不是宽泛 AI 新闻。"
  reuse_rule: "开发工作流工具追踪用具体工具 / CLI / 项目名，不用宽泛类别。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#当前已确认创建并保留; docs/ops/subscription-system-closeout-20260704.md#追踪器"
  next_action: "保留。"
```

---

## 6. AFM、WWDC、OriginOS、Apple Intelligence、BlueOS：候选、季节性、改词或不建

```yaml
- source_id: "Apple Foundation Models"
  current_route: precise_technical_tracking_created_after_device_note
  decision_status: recovered_with_temporal_boundary
  selected_reason: "设备生态文件中它实测通过但尚未创建：结果数量可控，主题集中于 Apple Foundation Models 模型架构、设备端 / 云端模型、Private Cloud Compute、Swift / Apple 平台开发框架、第三方模型、统一 API 和本地 AI 应用。"
  rejected_alternatives:
    - "Apple Intelligence"
    - "Apple AI"
    - "AI"
    - "Agent"
    - "大模型"
  route_reason: "设备生态文件当时只列为技术观察候选，未创建；随后订阅系统收口确认科技大类最终新增精确追踪 `Apple Foundation Models`。因此 ledger 记录时间边界：它从候选变成已创建，但理由仍是精确、低频、技术雷达，不是普通用户功能追踪。"
  reuse_rule: "同一来源状态会随平台执行变化而改变；旧文件的候选状态不能覆盖后续执行证据。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#Apple Foundation Models：通过，但尚未创建; docs/ops/subscription-system-closeout-20260704.md#追踪器"
  next_action: "按已创建追踪器处理。"

- source_id: "WWDC"
  current_route: seasonal_tracking_candidate_not_persistent
  decision_status: recovered_from_pr
  selected_reason: "关键词有效，能覆盖 Apple 开发者大会、系统平台更新、Xcode、Swift、Foundation Models、Apple Intelligence、Siri、Mac、Apple Silicon、本地 AI、会后复盘和教程。"
  rejected_alternatives:
    - "全年常驻"
    - "用 Apple 宽词代替"
  route_reason: "会后重复解读、开发教程堆积和市场叙事噪声较多。推荐窗口是每年大会前 4—6 周开始追踪，会后保留 3—4 周，然后取消。"
  reuse_rule: "发布会、开发者大会和展会词是季节性追踪，不是长期追踪。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#WWDC：通过，但尚未创建"
  next_action: "季节性使用。"

- source_id: "OriginOS"
  current_route: on_demand_or_version_window_tracking
  decision_status: recovered_from_pr
  selected_reason: "结果中确有系统升级、AI 轻办公、原子工作台、跨端能力和数字车钥匙等内容。"
  rejected_alternatives:
    - "长期追踪 OriginOS"
    - "用 OriginOS 替代 vivo办公套件"
  route_reason: "结果混入酷安吐槽、主题讨论、自制系统设计、第三方 ROM 移植、具体手机爆料与促销，仍不足以成为长期常驻。新版本发布窗口可临时追踪 `OriginOS 7` 等具体版本名。"
  reuse_rule: "操作系统总词如果社区噪声和硬件促销多，改为具体版本窗口词。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#OriginOS：不建立长期追踪"
  next_action: "按需。"

- source_id: "Apple Intelligence"
  current_route: not_created_overlap_and_repost_noise
  decision_status: recovered_from_pr
  selected_reason: "内容基本与 Apple AI 相关。"
  rejected_alternatives:
    - "长期追踪 Apple Intelligence"
    - "用它替代 AFM / WWDC / macOS / Apple Silicon"
  route_reason: "重复度高，主要被分析师或供应链报告多站转载、iPhone 换机与股价叙事、新机内存与硬件传闻、Siri 与 Apple Intelligence 重复报道占据；重要内容已可由 WWDC、macOS、Apple Silicon 和 Apple Foundation Models 更精确覆盖。"
  reuse_rule: "概念词如果被供应链、股价、换机和重复报道淹没，应拆成具体技术、平台或事件。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#Apple Intelligence：暂不创建"
  next_action: "不创建。"

- source_id: "BlueOS"
  current_route: rejected_due_to_homonym_use_chinese_official_name
  decision_status: recovered_from_pr
  selected_reason: "多数结果与 vivo 蓝河操作系统、WATCH、AIoT、Rust 内核和开发者大会有关。"
  rejected_alternatives:
    - "直接使用 BlueOS"
    - "同时创建 BlueOS 与蓝河操作系统"
  route_reason: "顶部存在 Blue Robotics 开源机器人平台同名污染，因此不用英文词，改用中文正式名 `蓝河操作系统`。"
  reuse_rule: "英文名有同名污染时，用官方中文名、全称或更窄对象名。"
  evidence_locator: "资源/雷达/设备生态关注说明.md#BlueOS：不直接使用"
  next_action: "不用 BlueOS。"
```

---

## 7. 文化公共谈话者：不是名人清单，而是关系类型地图

```yaml
- source_id: "文化公共谈话者候选整体"
  current_route: candidate_map_not_lighthouse_registry
  decision_status: recovered_from_pr
  selected_reason: "文件记录用户与人物、节目和谈话方式的不同关系：灯塔、坐标、对话入口、节目入口、水温入口和按需内容。"
  rejected_alternatives:
    - "把所有提到的人直接写入成为/灯塔.csv"
    - "按名人、学者或主持人知名度排序"
    - "把文化小众爱好和大城市生活方式包装成身份象征"
    - "把节目长期收听直接等同于灯塔"
  route_reason: "用户真正寻找的不是值得关注的名人，而是能把抽象问题带回现实生活、让不同立场进入谈话、把文化学术与普通经验连接起来、不要求听众把自己当权威、能在对话中显出如何承接冲突、误解、脆弱和真实处境的人。"
  reuse_rule: "人物进入灯塔前必须说明关系类型和真实影响；候选地图不是灯塔登记。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#状态; #核心判断"
  next_action: "继续作为候选地图。"
```

---

## 8. 为什么对话比演讲更重要

```yaml
- source_id: "对话入口优先规则"
  current_route: dialogue_form_as_evidence_not_speech_ranking
  decision_status: recovered_from_pr
  selected_reason: "演讲更多呈现整理完成、愿意公开的自我；真实对话可能显出未准备问题、攻击、误解、尴尬、沉默、矛盾、不确定、脆弱、脚本破裂和双方是否真的发生变化。"
  rejected_alternatives:
    - "用 TED、一席、专业演讲和学术讲座替代对话观察"
    - "把演讲者整理好的观点当长期关系证据"
    - "把主持人气质自动视为深度"
  route_reason: "长期评价对象常常不是主持人本人，而是这个人与另一个人在这次对话里究竟发生了什么。"
  reuse_rule: "处理访谈、播客、对谈、长谈时，要记录具体对话如何发生，而不只抽观点。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#为什么对话比演讲更重要"
  next_action: "用于影音与文化坐标判断。"
```

---

## 9. 已较明确人物位置：灯塔、坐标、节目坐标候选分开

```yaml
- source_id: "迈克尔·桑德尔"
  current_route: high_confidence_lighthouse_candidate_from_existing_coordinate
  decision_status: recovered_from_pr
  selected_reason: "高中观看《公正》课程，已经影响用户对公平、责任、优绩主义、运气与结构的理解；重要的是用户真实看过、讲过并带入生活的经历。"
  rejected_alternatives:
    - "只因其学者身份写入灯塔"
    - "把公共知名度当证据"
  route_reason: "已有深层坐标，可优先形成灯塔条目。"
  reuse_rule: "灯塔证据来自真实接触、持续影响和生活判断变化，不来自名气。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#已较明确的位置; #后续规则"
  next_action: "可优先形成灯塔条目。"

- source_id: "梁文道 / 刘擎 / 戴锦华"
  current_route: high_priority_lighthouse_candidates_require_real_materials
  decision_status: recovered_from_pr
  selected_reason: "梁文道可能照亮阅读如何回到经验、历史感与人的处境；刘擎可能照亮多元价值、政治哲学与中文公共生活；戴锦华已对用户产生影响，并在电影、文学和文化阅读中提供新视角，尤其让形式、镜头、叙事、性别、历史和意识形态成为理解现实的方式。"
  rejected_alternatives:
    - "直接写入灯塔"
    - "只按名望或学科标签判断"
  route_reason: "它们是高优先候选，但仍要从真实接触材料写‘为何点亮’，不把候选状态冒充已录入灯塔。"
  reuse_rule: "高优先灯塔候选需要具体材料、具体影响和可回访证据。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#梁文道; #刘擎; #戴锦华; 成为/灯塔/文化公共谈话者候选.md#后续规则"
  next_action: "从真实接触材料补为何点亮。"

- source_id: "Steve说 / 史蒂夫说"
  current_route: program_coordinate_candidate_from_changed_views
  decision_status: recovered_from_pr
  selected_reason: "用户曾在小宇宙较多收听，其中有几期真实改变过自己的看法，已经超过普通节目入口。"
  rejected_alternatives:
    - "因为长期收听直接写主持人为灯塔"
    - "只作为普通播客入口"
  route_reason: "更适合作为节目坐标候选：后续从真正改变判断的具体单集出发，确认它影响了什么。"
  reuse_rule: "节目坐标从改变判断的具体单集建立，不从收听时长或节目名建立。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#Steve说"
  next_action: "从具体单集补坐标证据。"
```

---

## 10. 对话入口、节目观察、水温入口与表达者按需

```yaml
- source_id: "许知远 / 十三邀"
  current_route: high_value_dialogue_entry_not_host_lighthouse_by_default
  decision_status: recovered_from_pr
  selected_reason: "部分访谈曾打动用户，但真正重要的常常是被访者与具体对话，不一定是许知远本人；其笨拙、不顺滑有时制造尴尬、错位和未被完全控制的空间。"
  rejected_alternatives:
    - "因知识分子气质自动视为深度"
    - "把主持人直接写入灯塔"
  route_reason: "应逐期判断谈话是否真的发生。"
  reuse_rule: "访谈节目按具体对话记录，不按主持人稳定人设记录。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#许知远 /《十三邀》"
  next_action: "逐期判断。"

- source_id: "窦文涛 / 鲁豫"
  current_route: dialogue_mechanism_observation
  decision_status: recovered_from_pr
  selected_reason: "窦文涛通过自嘲、把自己放低和不抢夺判断中心，使嘉宾更容易自然说话；鲁豫的不轻易相信可能诱发嘉宾继续解释、修正和主动袒露。"
  rejected_alternatives:
    - "按主持技巧直接写入灯塔"
    - "只抽结论不看互动机制"
  route_reason: "价值在于谈话机制：一个通过自我缩小承接，一个通过持续不信诱发更深袒露。"
  reuse_rule: "对话入口要记录机制和关系变化，而不仅是主持人观点。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#窦文涛; #鲁豫"
  next_action: "按具体对话记录。"

- source_id: "小林说 / 凉子访谈录 / 姜DORA"
  current_route: dialogue_entry_pending_identity_or_case_specific
  decision_status: recovered_from_pr
  selected_reason: "用户通常不是长期追随主理人，而是因具体对话对象进入；凉子访谈录、姜DORA 等名称或账号存在记忆待核。"
  rejected_alternatives:
    - "强行补全节目名或人物身份"
    - "因账号辨识度自动进入灯塔"
  route_reason: "获得更多线索前保持克制；重点判断双方是否真的相遇，还是各自完成准备好的输出。"
  reuse_rule: "记忆不完整时保留待核状态，不用模型补编身份。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#小林说; #凉子访谈录; #姜DORA; #后续规则"
  next_action: "待核。"

- source_id: "罗振宇 / 脱不花"
  current_route: long_form_program_and_life_stage_observation
  decision_status: recovered_from_pr
  selected_reason: "罗振宇 /《文明之旅》呈现把整个人投入长期节目、借节目处理人生中段、远行、文明和自身位置的状态；脱不花长谈把组织者、管理者与中年人的阶段问题同时放进长时间对话。"
  rejected_alternatives:
    - "仅按得到创始人或管理者身份理解"
    - "直接写成灯塔"
  route_reason: "它们当前更适合知识媒介坐标、长期节目观察、组织实践坐标，而不是简单名人追随。"
  reuse_rule: "长期节目可观察媒介形态与人生阶段，不一定转成灯塔。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#长期节目与知识媒介观察"
  next_action: "按节目和阶段观察。"

- source_id: "快刀青衣"
  current_route: chinese_ai_public_water_temperature_entry
  decision_status: recovered_from_pr
  selected_reason: "用于观察国内知识媒体、大众与商业人群当前认为哪些 AI 议题值得讲。"
  rejected_alternatives:
    - "当作 AI 信息源"
    - "当作技术判断来源"
  route_reason: "文件明确不把他当 AI 信息源，而是中文 AI 公共水温入口。"
  reuse_rule: "水温入口只观察公共叙事和话题选择，不当一手知识或技术来源。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#快刀青衣"
  next_action: "按水温使用。"
```

---

## 11. 公共学者观察候选：观察，不补编身份

```yaml
- source_id: "沈奕斐 / 梁永安 / 杜素娟 / 钱进"
  current_route: public_scholar_observation_candidates
  decision_status: recovered_from_pr
  selected_reason: "沈奕斐关注家庭、亲密关系、性别与社会结构怎样回到普通人的生活；梁永安可能通过文学、城市、爱情、青年与现代生活谈普通人的处境；杜素娟有眼缘和关注；钱进有模糊记忆。"
  rejected_alternatives:
    - "直接写入灯塔"
    - "补编机构、履历和身份"
    - "把眼缘或模糊记忆写成长期影响"
  route_reason: "杜素娟姓名已订正，但尚未说清具体课程、研究领域或长期影响；钱进具体身份待核。应保持观察和待核，不擅自补全。"
  reuse_rule: "人物身份不清或影响证据不足时，保留观察候选或待核状态。"
  evidence_locator: "成为/灯塔/文化公共谈话者候选.md#公共学者观察候选; #后续规则"
  next_action: "待更多线索。"
```

---

## 12. 资源总入口：活库与参考分层，不是大杂烩

```yaml
- source_id: "资源.md 总入口"
  current_route: para_resources_quick_entry_with_active_and_reference_layers
  decision_status: recovered_from_pr
  selected_reason: "`资源.md` 是 PARA 中 Resources 的支持与补给入口，按维护模式分为活库与参考两层。活库是结构化数据库，需要定期分拣与维护；参考页是低频修订的非结构化参考内容。"
  rejected_alternatives:
    - "把资源入口写成所有材料的大杂烩"
    - "让资源.md 承担治理规则"
    - "把所有参考页纳入自动代理覆盖"
  route_reason: "资源.md 用于快速找入口；资源/README.md 才负责判断资源是否该进入 Git、留借港、沉到领域或项目。"
  reuse_rule: "找入口看资源.md；判断居留看资源/README.md。活库和参考页分开。"
  evidence_locator: "资源.md#资源"
  next_action: "保持总入口。"
```

---

## 13. 资源 README：沉积要能回访、复用或承重

```yaml
- source_id: "资源/README.md 治理规则"
  current_route: resource_residence_and_update_gate
  decision_status: recovered_from_pr
  selected_reason: "`资源/` 保存已经导出、消化、可检索、可批量处理或可复用的材料。"
  rejected_alternatives:
    - "把 Notion 借港完整镜像进 Git"
    - "把外部仍在生长的空间搬进资源"
    - "只是为了找得到而复制外部空间"
    - "用外部平台状态直接覆盖沃壤状态"
  route_reason: "README 明确：外部空间仍在生长、需要阅读现场时留借港；已经被用户消化成摘录、判断、结构或作品材料时才沉资源或领域；需要脚本、版本、批处理或 source pack 时沉 GitHub。外部状态是证据，不是判决；候选事件是入口，不是主库；沉积必须能回访、复用或承重。"
  reuse_rule: "进入资源前先判材料状态：原始导出、半加工、可引用、操作参考、领域前体、身体候选。不能只看文件名。"
  evidence_locator: "资源/README.md#与借港的分工; #资源更新宪法; #材料状态; #什么时候沉到这里"
  next_action: "作为资源沉积规则。"
```

---

## 14. 语境索引：已知怎样用，才进语境

```yaml
- source_id: "资源/语境/_index.md"
  current_route: context_layer_for_known_use_cases
  decision_status: recovered_from_pr
  selected_reason: "`资源/语境/` 保存已经形成什么时候、为什么、怎样使用的资源语境。"
  rejected_alternatives:
    - "把只有 URL、标题和候选状态的链接放进语境"
    - "用语境替代资源/链接候选池"
    - "把浏览器导出 cohort 在语境里重复"
  route_reason: "索引明确：只有 URL、标题和候选状态留在资源/链接；已经知道何时使用、为何使用、与什么替代，才进入资源/语境；被长期能力域反复调用候选沉入领域；影响具体项目动作，在成为/项目中引用；迁移、清理、审计和对账证据进入 docs/ops。"
  reuse_rule: "链接 → 语境 → 领域 / 项目 / docs/ops 的上升路径要有使用证据。"
  evidence_locator: "资源/语境/_index.md#居留边界"
  next_action: "作为链接与语境边界。"
```

---

## 15. 身体资源边界：资源条目不是训练命令

```yaml
- source_id: "身体资源与跟练 Lens"
  current_route: resource_candidates_for_body_lens_not_health_strategy
  decision_status: recovered_from_pr
  selected_reason: "资源/README 将外部跟练视频、课程、动作资源、候选样本和 Body Resource Lens 候选、缺口、source candidates、samples、c01-intake 放入资源/跟练。"
  rejected_alternatives:
    - "把健康判断散落在资源条目"
    - "把跟练资源写成训练命令"
    - "把 Lens helper / prototype / tests 放进资源"
  route_reason: "训练选择原则、身体复利和健康策略归领域/健康；用户身体日用体验归成为/身体或成为/项目；Lens helper、prototype、tests、CLI / MCP 候选归 wo/；审计、迁移、workflow 证据归 docs/ops。"
  reuse_rule: "跟练 / 身体资源进入 Lens 前，必须保留候选状态、适用状态与避用条件，不把资源条目写成训练命令。"
  evidence_locator: "资源/README.md#身体资源与跟练 Lens; #维护原则"
  next_action: "作为身体资源边界。"
```

---

## 16. 本补充 ledger 的复用规则

1. 设备生态关注来自真实使用关系，不是品牌站队。
2. 取消宽词追踪不等于取消关注；只是移除失真入口。
3. 追踪器优先使用完整、可识别的专有名称，不用品牌名、AI、Agent、MacBook 等宽词。
4. `macOS`、`Apple Silicon`、`vivo办公套件`、`小 V Claw`、`蓝河操作系统`、`Codex CLI` 各有不同设备 / 平台 / 组件 / 工作流职责。
5. `Apple Foundation Models` 在设备生态文件中是候选，后续科技收口已创建；旧候选状态不能覆盖后续执行证据。
6. `WWDC`、`Apple Event`、具体版本和型号是季节性或窗口期追踪，不是长期常驻。
7. 文化公共谈话者文件是候选地图，不等于灯塔登记。
8. 灯塔、坐标、对话入口、节目入口、水温入口和按需内容要分开。
9. 对话要记录关系如何发生，不只抽观点。
10. 记忆不完整的人名或节目名保持待核，不补编身份。
11. 资源.md 是快速入口，资源/README.md 是治理规则。
12. 外部活空间留借港；沉资源必须能回访、复用或承重。
13. 只有已知何时、为何、怎样使用的材料，才进资源/语境。
14. 跟练 / 身体资源是候选材料，不是训练命令。

---

## 17. 仍需继续追回

资源总入口、设备生态与文化坐标路由已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如继续补仓库内残余，可检查 `README.md`、`docs/ops/README.md`、`资源/链接/_index.md` 是否还需要一份总索引 / 操作索引 ledger。
