# PR #372 判断链追回：开发目录 290 节点（2026-07-06）

## 0. 边界

本文件追回 `TopHub > 开发 / 程序编程` 目录 290 个节点最终收口后的判断链。

纳入文件：

- `docs/ops/tophub-development-directory-audit-20260704.md`
- `docs/ops/tophub-development-page-25-final-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 开发目录整体：不把开发变成 RSS 收件箱

```yaml
- source_id: "TopHub > 开发 / 程序编程 290 节点整体"
  current_route: no_new_tophub_no_folo_now_project_bound_sources
  decision_status: recovered_from_pr
  selected_reason: "开发目录包含大量真实有用来源：开源项目发现、技术资讯、版本发布、中文技术社区、产品经理、AI 商业文章、周刊、独立开发、安全、主机、营销和课程内容。"
  rejected_alternatives:
    - "把开发目录整体加入科技雷达"
    - "把 290 个节点压成开发资讯收件箱"
    - "把 Go、前端、Node、React、数据库、GPU、嵌入式、形式化证明等全部常驻"
    - "在冻结期提前扩 Folo"
  route_reason: "最终 TopHub 新增 0、取消 0、替换 0，Folo 新增 0，追踪器新增 0。原因不是没有好来源，而是 TopHub 负责公共注意力天气，不承担开发资讯收件箱；程序编程来源必须绑定真实学习、仓库或项目，而不是为所有可能技术栈建立未读债务。"
  reuse_rule: "以后遇到开发来源，先确认项目关系：正在学、正在写、正在部署、正在排障、正在选型、正在维护。没有项目关系就按需，不订阅。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#概况; docs/ops/tophub-development-page-25-final-closure-20260705.md#一完成状态; #三290节点最终架构结论; #五最终真实动作"
  next_action: "无配置动作；等待 2026-07-17 Folo 复查两个候选。"
```

---

## 2. HelloGitHub：保持低频开源项目策展入口

```yaml
- source_id: "HelloGitHub｜月刊"
  current_route: keep_existing_tophub_technology_radar_open_source_project_curation
  decision_status: recovered_from_pr
  selected_reason: "它是低频、月度、编辑筛选、面向开源项目的发现入口。每期有明确编号，中文介绍降低初次筛选成本，更新频率不会制造未读压力。"
  rejected_alternatives:
    - "GitHub Trending Today"
    - "GitHub Trending Weekly 取代 HelloGitHub"
    - "CSDN 今日头条热点"
    - "掘金全站 / AI / 工具热榜"
    - "开源中国热门资讯"
    - "GitHub 语言 Trending 全部订阅"
  route_reason: "继续保持，不被 Trending 取代。Trending 提供当前注意力，HelloGitHub 提供编辑筛选和解释。290 个节点里没有出现比它更适合当前系统的通用低频开源发现源。"
  reuse_rule: "开源项目发现优先选择低频、编辑解释、能回仓库核验的策展源；热榜只作临时漫游。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#HelloGitHub月刊; docs/ops/tophub-development-page-25-final-closure-20260705.md#2HelloGitHub月刊继续保持不被Trending取代; #四最终压缩结果"
  next_action: "无；保持。"

- source_id: "GitHub Trending"
  current_route: on_demand_open_source_attention_sample
  decision_status: recovered_from_pr
  selected_reason: "直接对应仓库，能快速看到当前开发者注意力流向，对 coding agent、MCP、浏览器自动化、代码库记忆、开源基础设施等方向有发现价值。"
  rejected_alternatives:
    - "Trending Today 常驻"
    - "Trending Weekly 取代 HelloGitHub"
    - "按语言 Trending 订阅"
  route_reason: "永久按需。日榜波动大，易被短期传播、发布活动和 AI 热点支配；星标数字常是累计星标；语言分类也会被主语言、模板或映射方式污染。上榜不代表成熟、可靠或适合当前工作流。"
  reuse_rule: "GitHub Trending 只用于发现。任何仓库进入判断前必须回 README、issues、release、维护活跃度、许可证、依赖和实际测试。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#GitHubTrendingTodayWeekly; docs/ops/tophub-development-page-25-final-closure-20260705.md#2HelloGitHub月刊继续保持不被Trending取代"
  next_action: "按需。"
```

---

## 3. 2026-07-17 复查的两个最高优先候选

```yaml
- source_id: "Golang Weekly｜英文周刊"
  current_route: folo_review_candidate_go_learning_low_frequency_weekly
  decision_status: recovered_from_pr
  selected_reason: "它最符合当前真实 Go 学习阶段：周频、编辑筛选、Go 语言和生态聚焦，文章、release、安全和工程实践并存。"
  rejected_alternatives:
    - "Go 语言中文网"
    - "中文课程搬运"
    - "陈旧招聘页"
    - "GitHub Go Trending"
    - "宽泛 Go 社区流"
  route_reason: "A1，建议路由为 Folo 低频正式源，执行时间为 2026-07-17 冻结复查时确认，不提前扩容，不进入 TopHub。"
  reuse_rule: "语言学习源优先选择低频编辑周刊；热榜、招聘、课程搬运和社区流按具体问题使用。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#3GolangWeekly英文周刊成为Go学习正式候选; #四最终压缩结果"
  next_action: "2026-07-17 Folo 复查。"

- source_id: "Hugging Face｜Blog 英文版"
  current_route: folo_review_candidate_first_party_ai_ml_engineering
  decision_status: recovered_from_pr
  selected_reason: "第一方完整发布，覆盖 Transformers、TRL、vLLM、PyTorch、推理、训练、OCR、ASR、embedding、Agent、MCP、机器人和本地 AI；具体工程内容可追溯代码、模型和论文。"
  rejected_alternatives:
    - "Hugging Face 中文博客同时订阅"
    - "Models / Spaces / Datasets 榜单作为阅读源"
    - "deeplearning.ai Weekly 作为开发目录候选"
  route_reason: "A1，2026-07-17 Folo 复查候选。英文 Blog 是规范源；中文博客只作中文降噪和低负担按需替代，不双订阅。Models、Spaces、Datasets 榜单是检索工具，不是阅读源。"
  reuse_rule: "第一方工程博客优先原始完整源；中文精选和榜单只作辅助入口。官方源也要等真实复查，不提前扩容。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#4HuggingFace英文Blog胜过中文博客; #四最终压缩结果"
  next_action: "2026-07-17 Folo 复查。"
```

---

## 4. 观察而不扩容：有价值，但不变成常驻

```yaml
- source_id: "InfoQ｜Today"
  current_route: on_demand_architecture_platform_engineering_source
  decision_status: recovered_from_pr
  selected_reason: "补足平台工程、可靠性、云原生、JVM / Parquet、OpenTelemetry、Netflix 工程、AWS、Cloudflare 和企业架构。"
  rejected_alternatives:
    - "加入 TopHub"
    - "冻结期加入 Folo"
    - "替代 The Register"
  route_reason: "B+ 生产级架构入口，但当前系统已有 The Register Latest、Google Developers Blog、具体项目官方博客与文档、The New Stack、Software Engineering Daily 等按需入口。做平台工程、云、可靠性或架构研究时主动查。"
  reuse_rule: "架构媒体只有进入实际架构 / 云 / 可靠性任务时启用；不要为可能的架构兴趣建立未读。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#6InfoQToday有价值但不成为常驻订阅"
  next_action: "按需。"

- source_id: "Frontend Focus / web.dev / CSS-Tricks / JavaScript Weekly / Node Weekly / React Status"
  current_route: frontend_project_bound_sources
  decision_status: recovered_from_pr
  selected_reason: "web.dev 是官方 Web Platform、Baseline、Core Web Vitals、性能、兼容性和权限 UX 参考；Frontend Focus 少而精；CSS-Tricks 适合现代 CSS 和交互深读；JS / Node / React 周刊与项目绑定。"
  rejected_alternatives:
    - "前端来源常驻"
    - "所有 Web / JS / React 周刊进入 Folo"
  route_reason: "最终按项目启用，不常驻。没有活跃前端项目时，这些来源会制造与当前仓库无关的浏览器、框架和 API 未读。"
  reuse_rule: "前端源先确认项目栈：Web Platform、CSS、JS、Node、React 各自按项目调用；项目结束后退回按需。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#7前端来源最终按项目启用不常驻; #四最终压缩结果"
  next_action: "按项目。"

- source_id: "TesterHome｜七日最热"
  current_route: on_demand_quality_testing_agent_acceptance_source
  decision_status: recovered_from_pr
  selected_reason: "比普通开发社区更有独特角色：测试设计、验收、质量保障、Agent 测试、Workflow 与生产级 Skill。"
  rejected_alternatives:
    - "加入 TopHub"
    - "先订阅再等用途出现"
  route_reason: "A2 条件候选降为 B+ 任务入口。当沃壤开始建立系统测试、Agent 验收或回归流程时启用；当前不是每天必须阅读的来源。"
  reuse_rule: "测试与质量来源必须绑定验收流程、回归流程或测试任务；没有测试流程时不制造未读。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#8TesterHome补足质量视角但目前仍按任务调用"
  next_action: "沃壤进入测试 / Agent 验收时启用。"

- source_id: "deeplearning.ai｜Weekly"
  current_route: ai_cross_directory_candidate_not_development_source
  decision_status: recovered_from_pr
  selected_reason: "低频编辑源，质量和节奏成立。"
  rejected_alternatives:
    - "从开发目录新增"
    - "优先于 Hugging Face Blog"
  route_reason: "不在开发目录新增，进入 AI 候选池 A2 观察位。只有当现有 AI 源缺少每周编辑摘要角色时再考虑。"
  reuse_rule: "跨目录来源应回到最合适目录比较；不要因出现在开发页就作为开发源新增。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#5deeplearningaiWeekly保留为AI跨目录候选不作为开发源"
  next_action: "AI 复查按需。"
```

---

## 5. 长访谈、安全、数据库 / GPU / 系统工程：任务和项目直接源

```yaml
- source_id: "Software Engineering Daily"
  current_route: on_demand_long_interview_archive
  decision_status: recovered_from_pr
  selected_reason: "内容质量高，覆盖语言与运行时、可观测性、Kubernetes、数据系统、AI 系统、游戏开发、开源维护和软件供应链。"
  rejected_alternatives:
    - "常驻未读"
    - "替代低频开发周刊"
  route_reason: "一次展示大量历史节目，阅读 / 收听成本高；最终作为 B+ 长访谈资料库，不进入常驻未读流。"
  reuse_rule: "访谈库按人物、公司、技术栈或系统问题检索；没有音频 / 长文阅读流程不订阅。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#10SoftwareEngineeringDaily只作长访谈资料库"
  next_action: "按需。"

- source_id: "安全开发来源"
  current_route: keep_xianzhi_plus_task_security_sources
  decision_status: recovered_from_pr
  selected_reason: "开发目录中有 FreeBuf、安全脉搏、安全内参、首席安全官、CodeQL、Hacker 社区与安全论坛等安全来源。"
  rejected_alternatives:
    - "新增第二组高频安全热榜"
    - "CodeQL / FreeBuf / 安全内参 / 安全脉搏 / 首席安全官全部常驻"
    - "黑客说 / Hacker Top Community / xLog 热榜"
  route_reason: "保持现有 `先知社区｜精华推荐`。CodeQL、FreeBuf、安全内参、安全脉搏、首席安全官按任务调用；含交易、恶意工具、凭据或高风险推广的节点整体排除。"
  reuse_rule: "安全来源按具体产品、漏洞、仓库、供应链或攻防任务启用；不建立第二条高频安全热榜。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#12安全来源不扩容"
  next_action: "按安全任务。"

- source_id: "数据库 / GPU / 嵌入式 / 系统工程项目源"
  current_route: project_adopted_direct_sources
  decision_status: recovered_from_pr
  selected_reason: "DuckDB News、阿里云数据库内核月报、CUDA Trending、Nix Today、HCL / Terraform Today、BitBake / Yocto Today、Erlang / Elixir、Lean、LLVM、CodeQL 等质量高，且接近官方、项目或生态。"
  rejected_alternatives:
    - "宽泛榜单常驻"
    - "在未采用技术栈前订阅"
    - "通过 TopHub Trending 间接跟踪项目"
  route_reason: "越接近底层或专业领域，越不适合订阅一个宽泛榜单。项目未采用则不订阅；项目采用后直接跟官方 Blog、release、changelog、仓库或规范。"
  reuse_rule: "底层技术源只在技术栈真实进入项目时升级；采用之后跟一手源，不跟聚合热榜。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#13数据库GPU嵌入式和系统工程采用项目直接源"
  next_action: "按项目采用情况。"

- source_id: "DuckDB｜News"
  current_route: on_demand_upgrade_if_project_adopts_duckdb
  decision_status: recovered_from_pr
  selected_reason: "官方第一方，版本、协议、格式和生态变化清晰，信噪比高，比数据库社区聚合更可复查。"
  rejected_alternatives:
    - "当前常驻"
    - "数据库社区聚合"
  route_reason: "用户当前没有确认 DuckDB 是长期技术栈。若沃壤、数据分析、BiBiGPT 数据处理或其他项目真实采用 DuckDB，再升级为项目直接源。"
  reuse_rule: "第一方项目源也不能因为质量高就自动常驻；必须等项目采用。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#第25页DuckDBNews"
  next_action: "项目采用时升级。"
```

---

## 6. AI Coding 与 Web3 / Rust / Git：回项目，不建宽泛流

```yaml
- source_id: "AI Coding 内容"
  current_route: route_to_ai_precise_trackers_or_project_docs_not_development_subscription
  decision_status: recovered_from_pr
  selected_reason: "开发目录大量节点反复包装 Codex、Claude Code、Agent、MCP、Skill、Loop、Harness、Vibe Coding。"
  rejected_alternatives:
    - "在开发目录重复订阅 AI Coding 媒体"
    - "用普通媒体二次包装替代项目文档"
    - "宽泛 Agent / AI Coding 追踪器"
  route_reason: "模型和 AI 能力变化进入 AI 目录；Codex CLI 等明确产品变化由精确追踪器处理；沃壤仓库实际工作流写入项目文档、PR 和复盘；普通媒体二次包装不订阅。只有出现可复现代码、基准、架构或真实失败分析时，才按任务阅读。"
  reuse_rule: "AI Coding 不用再建一条开发新闻流；回到产品、项目、代码、基准和复盘。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#15AICoding内容统一转项目与AI目录不在开发目录重复订阅"
  next_action: "按项目。"

- source_id: "Rust / Web3 / Git 开发来源"
  current_route: on_demand_no_new_general_source
  decision_status: recovered_from_pr
  selected_reason: "Rust、Web3 和 Git 都可能与用户学习和项目有关。"
  rejected_alternatives:
    - "Rust Trending 常驻"
    - "Web3 / 区块链热榜常驻"
    - "Git 文章聚合常驻"
  route_reason: "Rust Trending 是项目发现，不是生态解释；Web3 热榜大量交易、平台、量化与推广污染；Git 内容散落在掘金、博客园、DEV 和课程中，缺少稳定编辑源。Rust 回具体项目和官方生态，Web3 沿课程、官方文档和已核验书籍路线，Git 围绕真实仓库问题查官方文档、Missing Semester 和具体技术文章。"
  reuse_rule: "语言、区块链和 Git 概念源必须回学习路线或仓库问题，不订阅宽泛热榜。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#14RustWeb3与Git没有出现新的通用常驻源"
  next_action: "按需。"
```

---

## 7. 明确排除：污染、陈旧、时间口径不可靠、社区榜单过宽

```yaml
- source_id: "xLog 热榜"
  current_route: rejected_permanently_as_platform_hotlist
  decision_status: recovered_from_pr
  selected_reason: "无保留价值。"
  rejected_alternatives:
    - "xLog 今日最热"
    - "xLog 本周最热"
    - "xLog 本月最热"
    - "xLog 史上最热"
  route_reason: "四个层级均出现推广、诈骗式代币文案、交易所镜像、成人或侵害性内容。少量正常开源、Obsidian、RSS 和技术文章不能抵消热榜结构性污染。个别作者文章未来按作者或具体文章处理。"
  reuse_rule: "平台热榜不能作为信任背书；结构性污染出现时，正常文章不能抵消热榜风险。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#第25页xLog史上最热"
  next_action: "排除。"

- source_id: "CSDN / 掘金宽泛榜 / 博客园排行 / 开源中国社区讨论 / 51CTO / 奇绩创坛齐思 / Go中文课程搬运等"
  current_route: rejected_or_on_demand_broad_chinese_developer_platforms
  decision_status: recovered_from_pr
  selected_reason: "这些平台中有单篇可用技术材料。"
  rejected_alternatives:
    - "CSDN 今日头条热点"
    - "掘金全站 / AI / 工具热榜"
    - "开源中国热门资讯"
    - "博客园排行"
    - "Go语言中文网课程搬运与陈旧招聘"
  route_reason: "宽泛平台榜单存在时间口径不可靠、标题膨胀、课程 / 会议 / 活动推广、旧新闻、重复技术关键词、招聘或社区争论混杂。单篇可由搜索命中，但热榜不承担质量筛选。"
  reuse_rule: "中文开发平台按关键词检索单篇，不订阅全站、AI、工具或排行热榜。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#CSDN今日头条热点; #掘金; #开源中国; docs/ops/tophub-development-page-25-final-closure-20260705.md#四最终压缩结果"
  next_action: "按需或排除。"

- source_id: "Hacker News 各版本"
  current_route: on_demand_programmer_public_discussion
  decision_status: recovered_from_pr
  selected_reason: "Front Page 可观察程序员公共讨论。"
  rejected_alternatives:
    - "HN Front Page 常驻"
    - "HN 最新 / 热门 / 中文精选 / Top HN / 日报多版本常驻"
    - "Folo"
  route_reason: "Newest 噪声过大；Front Page 相对最好但仍混入科学、政治、历史与社会；中文精选有二次摘要和重复；Top HN、热门、日报角色重复。TopHub 已承担公共注意力天气，不再建立 HN 未读流。"
  reuse_rule: "HN 用于按需观察程序员讨论，不做事实源；多个镜像和时间窗口不并存。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#9HackerNews全部退回按需"
  next_action: "按需。"
```

---

## 8. 本补充 ledger 的复用规则

1. 开发目录不变成 RSS 收件箱；必须绑定学习、仓库、项目、排障、选型或维护。
2. HelloGitHub 保持低频开源项目策展，不被 Trending 替代。
3. GitHub Trending 只作注意力样本；仓库判断回 README、issues、release、许可证和维护状态。
4. Golang Weekly 与 Hugging Face 英文 Blog 是 2026-07-17 复查最高优先候选，不提前扩容。
5. InfoQ、Frontend、TesterHome、Software Engineering Daily、web.dev、CSS-Tricks 等按项目阶段启用。
6. 安全来源不扩容；保持先知社区，其他安全源按具体产品、漏洞、仓库或供应链任务调用。
7. 数据库、GPU、Nix、Lean、LLVM、CodeQL、DuckDB 等项目源采用项目直接源规则：未采用不订阅，采用后跟官方。
8. AI Coding 内容回 AI 目录、精确追踪器、项目文档和复盘，不在开发目录重复订阅。
9. Web3、Rust、Git 走学习路线 / 官方文档 / 项目任务，不订阅宽泛热榜。
10. 结构性污染、陈旧、时间口径不可靠和高风险推广，是开发热榜排除信号。

---

## 9. 仍需继续追回

开发目录 290 节点已补成单独 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前科技 / AI 补充 ledger；
- 社区同族 closure；
- 财经逐页；
- 购物逐页；
- 政务、校务、专栏、浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
