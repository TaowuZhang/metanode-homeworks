# PR #372 判断链追回：科技开发与文化边界来源（2026-07-06）

## 0. 边界

本文件是 `科技` 大类的补充 ledger，追回 PR #372 中科技雷达、开发目录、轻量科学漫游和社区链接流之间的边界判断。

纳入文件：

- `docs/ops/tophub-technology-directory-review-20260704.md`
- `docs/ops/tophub-technology-jandan-closure-20260704.md`
- `docs/ops/tophub-technology-chouti-closure-20260704.md`
- `docs/ops/tophub-development-directory-audit-20260704.md`
- `docs/ops/tophub-development-page-25-final-closure-20260705.md`
- `资源/链接/HelloGitHub.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 科技 / 开发 / 社区漫游总规则

```yaml
- source_id: "科技开发与文化边界总规则"
  current_route: keep_distinct_roles_no_extra_feed
  decision_status: recovered_from_pr
  selected_reason: "科技雷达、开发目录、社区链接流和科学漫游容易互相污染：都可能出现 AI、开源项目、工具、科学、文化和互联网社区话题。"
  rejected_alternatives:
    - "把开发目录 290 个节点转成开发资讯收件箱"
    - "把煎蛋和抽屉作为科技或公共温度常驻"
    - "用 GitHub Trending 取代 HelloGitHub"
    - "把论坛热榜、社区转帖和轻量漫游源放进 Folo"
  route_reason: "保留角色差异：HelloGitHub 负责低频开源项目策展；GitHub Trending 只按需看当前注意力；煎蛋是轻量科学与互联网文化漫游；抽屉是社区链接与情绪样本；开发目录按真实学习、仓库和项目启用。"
  reuse_rule: "以后遇到跨界科技源，先判断它是编辑策展、注意力热榜、作者关系、论坛转帖、项目官方源、还是轻量漫游。不同角色不能互相替代。"
  evidence_locator: "technology directory review / jandan closure / chouti closure / development directory audit / development final closure / HelloGitHub link card"
  next_action: "作为科技同族复用规则。"
```

---

## 2. HelloGitHub 与开发目录的关系

```yaml
- source_id: "HelloGitHub｜月刊"
  current_route: tophub_technology_radar_existing_keep
  decision_status: recovered_from_pr
  selected_reason: "它是低频、月度、编辑筛选、面向开源项目的发现入口；每期有明确编号，中文介绍降低初次筛选成本，更新频率不会制造高频未读压力。"
  rejected_alternatives:
    - "GitHub Trending Today"
    - "GitHub Trending Weekly 取代 HelloGitHub"
    - "CSDN 今日头条热点"
    - "掘金全站 / AI / 工具热榜"
    - "开源中国热门资讯"
    - "Hacker News 作为开发未读流"
  route_reason: "继续保留在科技雷达 / 现有常驻源中。HelloGitHub 与 Trending 的区别是：Trending 提供当前注意力，HelloGitHub 提供编辑筛选和解释。开发目录 290 节点最终也确认，没有出现比它更适合当前系统的通用低频开源发现源。"
  reuse_rule: "开源项目发现优先选择低频、编辑解释、可回到仓库核验的来源；热榜只作为临时漫游，不替代策展入口。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#HelloGitHub月刊; docs/ops/tophub-development-page-25-final-closure-20260705.md#HelloGitHub月刊继续保持不被Trending取代; 资源/链接/HelloGitHub.md"
  next_action: "无；保持。"

- source_id: "GitHub Trending"
  current_route: on_demand_open_source_attention
  decision_status: recovered_from_pr
  selected_reason: "能直接看到当前开发者注意力流向，对 coding agent、MCP、浏览器自动化、代码库记忆、开源基础设施等方向有发现价值。"
  rejected_alternatives:
    - "Trending Today 常驻"
    - "Trending Weekly 取代 HelloGitHub"
    - "按语言 Trending 全部订阅"
  route_reason: "日榜波动很大，容易被短期传播、发布活动和 AI 热点支配；星标数字常是累计星标，不能直接解释为当天增长；上榜不代表成熟、可靠或适合当前工作流。"
  reuse_rule: "GitHub Trending 只用于发现。任何仓库进入判断前必须回 README、issues、release、维护活跃度、许可证、依赖和实际安装测试。"
  evidence_locator: "docs/ops/tophub-development-directory-audit-20260704.md#GitHubTrendingTodayWeekly; docs/ops/tophub-development-page-25-final-closure-20260705.md#HelloGitHub月刊继续保持不被Trending取代"
  next_action: "按需。"

- source_id: "开发目录整体"
  current_route: project_bound_not_subscription_inbox
  decision_status: recovered_from_pr
  selected_reason: "开发目录 25 页、290 节点中存在大量好来源，但它们分别绑定 Go、前端、数据库、GPU、嵌入式、形式化证明、安全、AI coding、Web3、Git 等具体技术栈或项目。"
  rejected_alternatives:
    - "新增开发常驻节点"
    - "把 Go、React、Node、数据库、GPU、Nix、Lean 等生态都建未读流"
    - "把开发目录作为科技雷达扩容依据"
  route_reason: "最终 TopHub 新增 0、Folo 新增 0、追踪器新增 0。程序编程来源必须绑定真实学习、仓库或项目，而不是为所有可能技术栈建立未读债务。"
  reuse_rule: "技术栈未采用就不订阅；项目采用后直接跟官方 blog、release、changelog、仓库或规范。TopHub / Trending 只用于发现，不承担持续跟踪。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#完成状态; #三290节点最终架构结论; #数据库GPU嵌入式和系统工程采用项目直接源"
  next_action: "开发目录需要单独 development ledger 展开 290 节点，本文件只收科技边界。"
```

---

## 3. Golang Weekly、Hugging Face Blog 与开发候选边界

```yaml
- source_id: "Golang Weekly｜英文周刊"
  current_route: folo_review_high_priority_go_learning
  decision_status: recovered_from_pr
  selected_reason: "它最符合当前真实学习阶段：周频、编辑筛选、Go 语言和生态聚焦，文章、release、安全和工程实践并存。明显优于 Go 语言中文网、中文课程搬运、招聘页、GitHub Go Trending 和宽泛社区流。"
  rejected_alternatives:
    - "Go 相关热榜常驻 TopHub"
    - "GitHub Go Trending"
    - "中文 Go 门户 / 课程 / 招聘流"
  route_reason: "建议作为 Folo 低频正式源，在 2026-07-17 冻结复查时确认，不提前扩容；不进入 TopHub。"
  reuse_rule: "语言学习源优先选择低频编辑周刊；语言热榜、招聘流、课程搬运和社区热点按具体问题使用。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#GolangWeekly英文周刊成为Go学习正式候选"
  next_action: "Folo 复查。"

- source_id: "Hugging Face｜Blog 英文版"
  current_route: folo_review_high_priority_ai_ml_engineering
  decision_status: recovered_from_pr
  selected_reason: "第一方完整发布，覆盖 Transformers、TRL、vLLM、PyTorch、推理、训练、OCR、ASR、embedding、Agent、MCP、机器人和本地 AI；具体工程内容可追溯代码、模型和论文。"
  rejected_alternatives:
    - "Hugging Face 中文博客同时订阅"
    - "Models / Spaces / Datasets 榜单作为阅读源"
  route_reason: "英文 Blog 胜过中文博客。中文博客作为中文降噪入口按需，榜单作为检索工具，不是阅读源。"
  reuse_rule: "第一方工程博客如果进入 Folo，必须优先选择原始完整源；中文精选和榜单只作辅助入口。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#HuggingFace英文Blog胜过中文博客"
  next_action: "Folo 复查。"

- source_id: "InfoQ / Frontend / TesterHome / Software Engineering Daily 等开发候选"
  current_route: observe_or_task_triggered_not_current_add
  decision_status: recovered_from_pr
  selected_reason: "它们各有价值：InfoQ 补平台工程与架构；Frontend Focus、web.dev、CSS-Tricks 补 Web 项目；TesterHome 补质量与验收；Software Engineering Daily 补长访谈资料库。"
  rejected_alternatives:
    - "全部进入 TopHub"
    - "冻结期加入 Folo"
    - "先订阅再等用途出现"
  route_reason: "当前没有足够项目责任。InfoQ 不替代 The Register 和官方源；前端来源按项目启用；TesterHome 只有沃壤进入系统测试 / Agent 验收 / 回归流程时启用；Software Engineering Daily 收听成本高，作资料库。"
  reuse_rule: "开发候选先绑定项目阶段：学习、实现、测试、部署、架构、访谈资料。没有项目阶段就按需。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#InfoQToday有价值但不成为常驻订阅; #前端来源最终按项目启用不常驻; docs/ops/tophub-development-page-25-final-closure-20260705.md#TesterHome补足质量视角但目前仍按任务调用"
  next_action: "按项目或 Folo 复查。"
```

---

## 4. 煎蛋：轻量科学与互联网文化漫游

```yaml
- source_id: "煎蛋｜每日更新"
  current_route: on_demand_science_and_internet_culture_wandering
  decision_status: recovered_from_pr
  selected_reason: "能提供科学与技术之外的意外连接、互联网历史、网络文化、奇特物理 / 动物 / 医学 / 社会题目，以及可继续追原论文或原始报道的线索。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "慢读与思想"
    - "生活与社区"
    - "Folo"
    - "替代果壳 / 科普中国"
  route_reason: "不新增。每日更新混合科学、医学、政治、社会、科技、AI、金融、互联网史和争议研究标题，边界不稳；医学和健康标题常有强结论表达，不能直接转为行动。"
  reuse_rule: "轻量科学奇闻源可以按需漫游；涉及医学、健康或重大科学结论时必须回论文、专业机构或正式科普核验。"
  evidence_locator: "docs/ops/tophub-technology-jandan-closure-20260704.md#核心结论; #一煎蛋每日更新; #三与现有来源的关系"
  next_action: "按需。"

- source_id: "煎蛋｜无聊图"
  current_route: rejected_or_on_demand_image_wandering
  decision_status: recovered_from_pr
  selected_reason: "若恢复更新，可能有社区图片与娱乐漫游价值。"
  rejected_alternatives:
    - "视觉与自然常驻"
    - "Folo"
  route_reason: "当前页面暂无内容；即使以后恢复，也缺少稳定摄影、自然、艺术或视觉策展标准，不能与 NASA、国家地理、北京天文馆、胶片的味道等视觉节点相比。"
  reuse_rule: "社区图片流不因有图就进入视觉层；必须有稳定策展标准、主题边界和可持续内容。"
  evidence_locator: "docs/ops/tophub-technology-jandan-closure-20260704.md#二煎蛋无聊图"
  next_action: "按需或排除。"
```

---

## 5. 抽屉：社区链接与情绪样本

```yaml
- source_id: "抽屉新热榜｜最新 / 一周 / 24小时 / 3天"
  current_route: on_demand_community_link_and_mood_sample
  decision_status: recovered_from_pr
  selected_reason: "能显示某类中文社区正在转发什么，适合按需观察近期传播、社区情绪、段子和链接流。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "公共温度常驻"
    - "生活与社区常驻"
    - "多个时间窗口并存"
    - "Folo"
  route_reason: "不新增。科技、社会、政治、娱乐、科学、体育混在同一流；标题依赖情绪、猎奇和未核实叙述；原始来源不稳定，与微博热搜、知乎热榜、微信热文和公共温度高度重复。时间窗口变化不是新来源能力。"
  reuse_rule: "社区热榜只能作为情绪和传播样本，不作事实核验。多个时间窗口视为同一来源能力。"
  evidence_locator: "docs/ops/tophub-technology-chouti-closure-20260704.md#核心结论; #一抽屉新热榜最新; #二一周最热榜24小时最热榜3天最热榜"
  next_action: "按需。"

- source_id: "抽屉｜挨踢1024"
  current_route: on_demand_tech_community_spread_sample
  decision_status: recovered_from_pr
  selected_reason: "九个节点中最接近科技社区，能看到 AI 使用体验、技术事件、工具、开源项目、Codex 配额、GitHub 工具和科技周刊在民间社区如何扩散。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "Folo"
    - "作为事实来源"
    - "替代少数派 Matrix"
  route_reason: "不新增。大量内容来自二次转帖和未经验证传闻；AI 与特定公司争议占比高；与 Matrix 热榜、IT之家、极客公园、HelloGitHub 和科技追踪器重复。若只保留一个民间技术社区入口，Matrix 热榜更稳定，因为它以原创产品、开发过程和真实工作流为主。"
  reuse_rule: "技术社区转帖源用于看扩散，不用于定案；找到关键词后回原始博客、GitHub、官方公告或论文。"
  evidence_locator: "docs/ops/tophub-technology-chouti-closure-20260704.md#七挨踢1024; #八与少数派Matrix热榜的关系"
  next_action: "按需。"

- source_id: "抽屉段子 / 你问我答 / 图片 / 42区"
  current_route: rejected_or_on_demand_community_misc
  decision_status: recovered_from_pr
  selected_reason: "可在研究网络语言、社区情绪、图片传播或社区精选时按需看。"
  rejected_alternatives:
    - "生活与社区"
    - "视觉与自然"
    - "公共温度"
    - "Folo"
  route_reason: "段子不是科技来源；你问我答混合个人求助和新闻残留且时间错位；图片没有稳定审美、自然、摄影或视觉策展标准；42区边界混乱，有商业、新闻、娱乐、健康和社会混杂。"
  reuse_rule: "社区分区名不能替代内容边界。问答、图片、段子、精选区要分别看是否有稳定策展和责任关系。"
  evidence_locator: "docs/ops/tophub-technology-chouti-closure-20260704.md#三段子; #四你问我答; #五图片; #六42区"
  next_action: "按需或排除。"
```

---

## 6. 科技目录首页：为什么不直接从首页扩容

```yaml
- source_id: "TopHub 科技大类首页"
  current_route: directory_index_not_closure
  decision_status: recovered_from_pr
  selected_reason: "科技首页只是跨媒体、跨主题总目录与热门节点展示页，不是可直接闭环的小标签。"
  rejected_alternatives:
    - "连续翻 35 页并把首页候选直接加入"
    - "少数派最新与少数派热门双订阅"
    - "36氪最新与 24 小时热榜双订阅"
    - "威锋网今日新信息立即加入"
  route_reason: "首页不产生立即新增、替换或删除动作。它只暴露复查压力：36氪、虎嗅、果壳、IT之家、爱范儿、Readhub 等需要回各自小标签判断。"
  reuse_rule: "目录首页用于规划后续审查，不用于直接决策；同站最新 / 热门 / 分类页要回小标签做同族比较。"
  evidence_locator: "docs/ops/tophub-technology-directory-review-20260704.md#页面性质; #本页对三个固定问题的阶段性回答"
  next_action: "无；作为科技目录入口边界。"
```

---

## 7. 本补充 ledger 的复用规则

1. HelloGitHub 是低频开源项目策展，不被 GitHub Trending 取代。
2. GitHub Trending 是注意力样本，不是项目质量判断；仓库必须回 README、issues、release、许可证和维护状态。
3. 开发目录不变成开发 RSS 收件箱；语言、框架、数据库、前端、安全、GPU、嵌入式等都按真实项目启用。
4. Golang Weekly、Hugging Face Blog 等高价值来源进入 Folo 复查，不提前扩容。
5. 煎蛋是低责任科学与互联网文化漫游，不替代果壳、科普中国或正式科学源。
6. 抽屉是社区传播与情绪样本，不是事实源；挨踢1024 也不能替代 Matrix 热榜。
7. 科技目录首页只做复查路线图，不直接新增节点。
8. 社区转帖、轻量漫游、开发热榜和项目官方源必须分层，不得混为“科技来源”。

---

## 8. 仍需继续追回

科技大类已补入：

- 商业叙事与全球化来源；
- 科学、科普、企业基础设施、Apple、汽车、报告与设备市场结构；
- 当前科技雷达来源；
- AI、App、快讯、限免、教育、电商、数码、酷安、Google、TechWeb 等任务触发与高频流来源；
- HelloGitHub / 开发目录 / 煎蛋 / 抽屉 / 科技首页边界。

继续待办：

- 科技补充 ledger 后续等用户本地提交 master checklist 与 coverage audit 后，再统一登记；
- `AI reassessment` 六批的让位关系；
- `开发目录` 290 节点需要单独 development ledger 展开；
- 社区、财经、购物、政务、校务、专栏、浏览器与链接仍未追回完。
