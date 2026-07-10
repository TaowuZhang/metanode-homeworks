# PR #372 判断链追回：AI reassessment 压缩与让位关系（2026-07-06）

## 0. 边界

本文件追回 `TopHub > 人工智能` 215 个节点复审后的判断链，重点不是重新评价 AI 节点，而是沉积：

- 为什么不新增 AI 专属 TopHub 节点；
- 为什么不建立 AI 独立分组；
- 为什么多数来源让位给现有科技 / 开发 / 科学 / 安全 / 数据结构 / Folo 关系；
- 为什么最终 Folo 只保留四个互补候选；
- 为什么高质量来源仍可能只进入沃壤链接、专题入口或任务型检索。

纳入文件：

- `docs/ops/tophub-ai-reassessment-pages-01-03-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-04-06-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-07-09-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-10-12-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-13-15-20260705.md`
- `docs/ops/tophub-ai-reassessment-pages-16-18-20260705.md`
- `docs/ops/tophub-ai-final-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. AI 目录总判断：不把 AI 重新装进一个总入口

```yaml
- source_id: "TopHub > 人工智能 215 节点整体"
  current_route: no_new_tophub_no_ai_group_no_broad_tracker
  decision_status: recovered_from_pr
  selected_reason: "AI 重要，但它不是一个应与科技、开发、科学、数据结构、安全和慢读并列的独立容器。模型发布属于科技，Coding Agent 属于开发，安全漏洞属于安全，芯片与数据中心属于数据与结构，教育、劳动和权力属于慢读与思想。"
  rejected_alternatives:
    - "新增 AI 日报"
    - "新增 AI 快讯"
    - "新增模型热榜"
    - "新增 AI 媒体子频道"
    - "新增完整论文流"
    - "新增人工智能独立分组"
    - "建立 AI / 大模型 / Agent / AI 编程等宽泛追踪器"
  route_reason: "18 页、215 个节点复审完成后，TopHub 动作仍为 0；现有科技雷达、开发、安全、科学、数据结构和 Folo 直接源已经持续显影 AI。把它们重新装进 AI 热榜，只会重复发布、融资、跑分和情绪标题。"
  reuse_rule: "以后遇到 AI 来源，先判断它真实属于哪一层：模型产品、工程实践、论文研究、工具链、数据中心、社会制度、作者关系或任务文档。不要默认进 AI 总入口。"
  evidence_locator: "docs/ops/tophub-ai-final-closure-20260705.md#一句话结论; #一TopHub最终动作0; #九追踪器结论; #十最终动作表"
  next_action: "无 TopHub 动作；等待 2026-07-17 Folo 复查。"
```

---

## 2. 最终四个 Folo 候选

```yaml
- source_id: "Simon Willison's Weblog"
  current_route: folo_review_candidate_actual_tools_and_coding_agent_experiments
  decision_status: recovered_from_pr
  selected_reason: "角色是实际工具行为与 Coding Agent 实验。亲自测试模型、API、Agent 和开源工具，重视代码、复现实验和实际失败，同时覆盖 AI 编程、数据、安全与 Web；与 Codex、GitHub PR 和本地 Agent 工作流直接相关。"
  rejected_alternatives:
    - "philschmid.de 作为唯一工具教程源"
    - "O’Reilly Radar 作为唯一工程判断源"
    - "所有 Agent 工程博客一起加入"
  route_reason: "进入 2026-07-17 Folo 复查四候选之一，不进 TopHub。它是已有沃壤链接关系的升级候选，而非新发现；但更新最频繁，应放在触发 / 实验观察，而不是篇篇必读。"
  reuse_rule: "作者源进入 Folo 前要看是否亲自实验、是否给代码和失败、是否能追原始来源、是否与当前项目直接相关。更新过密时用更低频替补，而不是叠加多个工程作者。"
  evidence_locator: "tophub-ai-reassessment-pages-10-12-20260705.md#SimonWillisonsWeblog; tophub-ai-final-closure-20260705.md#SimonWillisonsWeblog"
  next_action: "2026-07-17 复查；若未读压力过大，用 Chip Huyen 替换。"

- source_id: "Eugene Yan"
  current_route: folo_review_candidate_production_systems_evals_search_recommendation_team_mechanisms
  decision_status: recovered_from_pr
  selected_reason: "角色是生产系统、Evals、搜索推荐与团队机制。覆盖 Product Evals、LLM-as-Judge、长上下文、安全评测、推荐系统、搜索、MCP、新闻 Agent、团队机制、学习和写作。"
  rejected_alternatives:
    - "InfoQ / InfoWorld / MLOps.community 作为生产工程持续源"
    - "TesterHome 或工程论坛热榜"
    - "只保留一般 AI 工程教程"
  route_reason: "进入 Folo 复查四候选之一。它比一般 AI 工程教程更接近长期构建、验证和维护系统；更新低频、结构清楚，和 Codex、GitHub PR、沃壤、阅读系统与 Agent 评测高度一致。"
  reuse_rule: "生产系统来源优先看是否能沉积评测、失败模式、团队机制和可复用结构；不是看教程数量。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#EugeneYan; tophub-ai-final-closure-20260705.md#EugeneYan"
  next_action: "2026-07-17 Folo 复查。"

- source_id: "AI as Normal Technology"
  current_route: folo_or_slow_read_review_candidate_reliability_institutional_conditions_anti_hype
  decision_status: recovered_from_pr
  selected_reason: "角色是可靠性、制度条件与反神话判断。讨论 Agent 可靠性、开放世界评测、AI 替代工程师、AI 建成操作系统等宏大说法的核验，把 AI 放回组织、劳动、法律、科学和制度条件中。"
  rejected_alternatives:
    - "用厂商叙事解释 AI"
    - "用 AGI 里程碑叙事解释 AI"
    - "用 x-risk 单一路线解释 AI"
    - "让 One Useful Thing 或 The Algorithmic Bridge 单独承担判断框架"
  route_reason: "进入四候选之一，但边界明确：正常技术也是一个分析框架，不能被反过来当作否定所有非线性变化的默认答案；必须与工程证据、产业材料和其他社会视角并读。"
  reuse_rule: "判断型来源可以长期保留，但不能成为新的裁判。它的价值是给出一套可比较框架，而不是替用户决定 AI 应该如何被理解。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#AIasNormalTechnology; tophub-ai-final-closure-20260705.md#AIasNormalTechnology"
  next_action: "2026-07-17 Folo 或慢读复查。"

- source_id: "Rest of World"
  current_route: folo_or_slow_read_review_candidate_non_us_technology_society_regions_labor_supply_chain
  decision_status: recovered_from_pr
  selected_reason: "角色是美国之外的技术社会与地区经验。覆盖印度、中国、全球南方、跨国平台、供应链、数据劳动、移民技术劳工、教育和地区产业变化，把技术从美国公司发布会拉回地区、劳工和具体人的生活。"
  rejected_alternatives:
    - "36氪出海作为唯一全球化来源"
    - "TechCrunch / The Verge / 美国科技媒体继续覆盖全球技术社会"
    - "AINOW 或 The Gradient 作为当前持续社会视角"
  route_reason: "进入四候选之一。它与 `36氪出海`不同：36氪出海偏中国企业全球经营，Rest of World 偏技术社会与人的经验，并能补足现有英语科技媒体的美国中心倾向。"
  reuse_rule: "技术全球化来源要区分企业经营、地区社会、劳工经验、供应链和平台扩张。不能用美国科技媒体或中国企业出海报道覆盖所有地区经验。"
  evidence_locator: "tophub-ai-reassessment-pages-16-18-20260705.md#RestofWorld; tophub-ai-final-closure-20260705.md#RestofWorld"
  next_action: "2026-07-17 Folo 或慢读复查。"
```

---

## 3. 唯一替补：Chip Huyen

```yaml
- source_id: "Chip Huyen"
  current_route: folo_backup_candidate_for_lower_frequency_gen_ai_platform_agent_model_routing_production_architecture
  decision_status: recovered_from_pr
  selected_reason: "角色是生成式 AI 平台、Agent、模型路由与生产架构。她长期连接研究、基础设施、产品和生产系统，文章不是发布快讯，而是结构化解释，更新频率低。"
  rejected_alternatives:
    - "Simon Willison + Eugene Yan + Chip Huyen 三者全部叠加"
    - "用 philschmid.de / Latent.Space / MLOps.community 替代低频稳定基线"
  route_reason: "Chip 本身完全值得长期阅读，但为了避免工程源叠加，当前只作为 Simon 的替补：若 Simon 更新过密，则取消 Simon 候选位置，改用 Chip；Folo 最终仍只增加 4 个，不增加 5 个。"
  reuse_rule: "同一工程角色只能保留少数互补源。低频稳定作者可以作为高频实验作者的替补，而不是叠加。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#ChipHuyen; tophub-ai-final-closure-20260705.md#唯一替补ChipHuyen"
  next_action: "作为 Simon 的替补规则保留。"
```

---

## 4. 为什么 O’Reilly、Ahead、philschmid、Hugging Face、Interconnects 等让位

```yaml
- source_id: "O’Reilly Radar"
  current_route: on_demand_or_worang_link_engineering_judgment_source
  decision_status: recovered_from_pr
  selected_reason: "Agent 工程、系统设计、失败模式与组织实践很强，样本覆盖 prompt injection、Agent memory、agentic code review、loop engineering、principal drift、Agent experience、平台自建等问题。"
  rejected_alternatives:
    - "进入最终四个 Folo 名额"
    - "替代 Eugene Yan 或 AI as Normal Technology"
  route_reason: "最终未进入四个名额，不是质量差，而是角色被拆分后已有更合适持续源：Eugene 承担生产系统与评测，AI as Normal Technology 承担可靠性与制度判断，Simon 承担实际工具实验。O’Reilly 退为按需或沃壤链接。"
  reuse_rule: "工程判断媒体如果角色被作者型来源覆盖，就保留为任务源，不进 Folo 长期队列。"
  evidence_locator: "tophub-ai-reassessment-pages-04-06-20260705.md#OReillyRadar; tophub-ai-final-closure-20260705.md#工程来源的最终去重"
  next_action: "按需。"

- source_id: "Ahead of AI"
  current_route: on_demand_or_worang_link_llm_architecture_learning_source
  decision_status: recovered_from_pr
  selected_reason: "LLM 架构、Coding Agent、论文与从零实现很强，比论文榜更重视理解路径，也比教程聚合更稳定。"
  rejected_alternatives:
    - "进入最终四个 Folo 名额"
    - "替代 Simon / Eugene / Chip"
  route_reason: "最终让位于四个更互补的持续关系。它仍是学习 LLM 架构和 coding agent 底层的重要入口，但不是当前新增 Folo 名额里的不可替代角色。"
  reuse_rule: "学习型 LLM 来源优先按学习阶段调用；只有当学习路线变成持续主线时才升级。"
  evidence_locator: "tophub-ai-reassessment-pages-07-09-20260705.md#AheadofAI; tophub-ai-final-closure-20260705.md#工程来源的最终去重"
  next_action: "按需或沃壤链接。"

- source_id: "philschmid.de"
  current_route: on_demand_agent_mcp_skill_harness_deployment_tutorial_source
  decision_status: recovered_from_pr
  selected_reason: "Agent Skill、AGENTS.md、Harness、Context Engineering、MCP、Codex CLI、Gemini API、Computer Use、Deep Research、Hugging Face、开放模型、微调、评测与部署都非常贴近实践。"
  rejected_alternatives:
    - "进入最终四个 Folo 名额"
    - "作为唯一 Agent 工程源"
  route_reason: "让位原因是强绑定 Google / Gemini、云平台和具体产品教程，部分内容随 API 快速过期。当前由 Simon 承担真实工具实验，Eugene 承担系统评测，Chip 作为低频替补。"
  reuse_rule: "强教程型作者源按具体技术栈 / API / 工程任务调用；产品教程可能过期，不能直接变成长期默认。"
  evidence_locator: "tophub-ai-reassessment-pages-10-12-20260705.md#philschmidde; tophub-ai-final-closure-20260705.md#工程来源的最终去重"
  next_action: "按需。"

- source_id: "Hugging Face｜Blog"
  current_route: on_demand_official_open_model_engineering_source
  decision_status: recovered_from_pr
  selected_reason: "官方开放模型、工具与工程生态价值高，覆盖模型、数据集、训练、推理、工具、Agent、MCP、安全和本地模型。"
  rejected_alternatives:
    - "Hugging Face 中文博客同时订阅"
    - "Hugging Face 模型 / Spaces / 数据集榜单成为阅读源"
    - "进入最终四个 Folo 名额"
  route_reason: "最终未进持续流，因为更新量大、产品范围广、官方文章与合作方 / 模型发布内容混合。需要模型、数据集、训练、推理或具体工具时访问。中文博客只是低负担替代，不与英文 Blog 双订阅。"
  reuse_rule: "官方工程博客是任务入口；除非它成为当前核心技术栈，否则不要把全量官方更新变成未读。"
  evidence_locator: "tophub-ai-reassessment-pages-04-06-20260705.md#HuggingFaceBlog; tophub-ai-final-closure-20260705.md#HuggingFaceBlog; #AI周报和中文来源"
  next_action: "按任务。"

- source_id: "Interconnects"
  current_route: worang_link_or_on_demand_model_ecosystem_governance_source
  decision_status: recovered_from_pr
  selected_reason: "开放模型、后训练、蒸馏、开放工件、中国 AI 实验室、模型治理、权重开放、性能差距和长期下注有独立价值。"
  rejected_alternatives:
    - "进入最终四个 Folo 名额"
    - "替代 AI as Normal Technology 或 Rest of World"
  route_reason: "质量高，但与 Last Week in AI、官方模型来源和工程候选有覆盖。保留在沃壤，不进入这一轮 Folo 四个名额。"
  reuse_rule: "模型生态判断源可作为重要视角保存，但若与现有周报和官方源覆盖，就不进入新增队列。"
  evidence_locator: "tophub-ai-reassessment-pages-10-12-20260705.md#Interconnects; tophub-ai-final-closure-20260705.md#Interconnects"
  next_action: "沃壤链接或按需。"
```

---

## 5. 专业研究入口：知道它们，但不形成普通未读流

```yaml
- source_id: "Stanford CRFM"
  current_route: on_demand_ai_evaluation_agent_reliability_research_source
  decision_status: recovered_from_pr
  selected_reason: "HELM、模型能力、安全和长上下文评测、Agent 攻防、缺陷协调披露、评测效率，是 AI 评测与 Agent 可靠性主题首选研究入口。"
  rejected_alternatives:
    - "完整 arXiv 分类"
    - "JMLR / Nature / 论文目录常驻"
    - "进入 Folo 普通阅读流"
  route_reason: "专业性强、更新不稳定，学术项目页面与公告混合；需要建立 AI 评测与 Agent 可靠性研究关系时优先调用，不订阅整个机构动态。"
  reuse_rule: "专业研究中心按主题调用；研究关系成立时追具体项目、benchmark、论文和漏洞披露，不追机构全站。"
  evidence_locator: "tophub-ai-reassessment-pages-16-18-20260705.md#StanfordCRFM; tophub-ai-final-closure-20260705.md#StanfordCRFM"
  next_action: "按研究任务。"

- source_id: "SemiAnalysis"
  current_route: on_demand_ai_compute_economics_infrastructure_source
  decision_status: recovered_from_pr
  selected_reason: "负责 GPU、HBM、晶圆、数据中心功率与可靠性、训练、推理、TCO、AI 资本开支和供应链。能把模型竞赛落到真实算力、产能、网络、内存和成本。"
  rejected_alternatives:
    - "进入 Folo 普通未读"
    - "用 Counterpoint 或一般科技媒体替代"
  route_reason: "不可替代，但付费墙、篇幅、专业门槛和核验负担高。研究算力经济时访问，不成为普通阅读流。"
  reuse_rule: "基础设施经济来源只在芯片、算力、数据中心、资本开支和供应链问题出现时启用，并与财报、供应链和其他来源交叉核验。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#SemiAnalysis; tophub-ai-final-closure-20260705.md#SemiAnalysis"
  next_action: "按任务。"

- source_id: "EleutherAI Blog"
  current_route: on_demand_open_model_safety_interpretability_research_source
  decision_status: recovered_from_pr
  selected_reason: "开放模型研究、Reward Hacking、机制可解释性、开放权重安全、训练数据、RLHF 和模型透明度具有独立价值。"
  rejected_alternatives:
    - "进入普通 Folo 未读"
    - "替代 Hugging Face Blog 或 Interconnects"
  route_reason: "研究密度高、更新不规则，并非每篇都适合当前学习阶段。作为研究入口使用，不要求持续阅读。"
  reuse_rule: "开放模型研究源按论文、代码和数据验证使用；不因研究质量高就变成通用未读流。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#EleutherAIBlog; tophub-ai-final-closure-20260705.md#EleutherAIBlog"
  next_action: "按研究任务。"

- source_id: "BAIR Blog / JMLR / Nature Machine Learning / arXiv cs.CL-cs.CV-cs.LG / TensorFlow Blog"
  current_route: on_demand_research_and_framework_sources
  decision_status: recovered_from_pr
  selected_reason: "这些源质量高，分别承担实验室研究、原始期刊、科学研究、论文分类和官方框架任务。"
  rejected_alternatives:
    - "完整论文 / 期刊 / 框架流常驻"
    - "把论文标题流变成每日责任"
  route_reason: "论文和框架目录密度太高、跨度太广，标题不能替代方法质量、复现性、许可证、安全性和任务适配。"
  reuse_rule: "论文和框架源围绕具体问题、关键词、作者、论文引用链、项目需要调用；不订阅完整分类。"
  evidence_locator: "tophub-ai-reassessment-pages-16-18-20260705.md#BAIRBlog; #JMLR; #NatureMachineLearning; #TensorFlowBlog; tophub-ai-final-closure-20260705.md#NatureMachineLearningJMLRarXiv"
  next_action: "按任务。"
```

---

## 6. 慢读与视角候选：保留，但不进入本轮四个名额

```yaml
- source_id: "The Gradient"
  current_route: slow_read_or_on_demand_research_math_philosophy_technology_society
  decision_status: recovered_from_pr
  selected_reason: "研究、数学、哲学和技术社会长文平台，能讨论 alignment、数学、purpose、wellbeing、bias、模型失败和技术社会问题。"
  rejected_alternatives:
    - "文字版和播客双订阅"
    - "进入本轮四个 Folo 名额"
  route_reason: "值得按篇读，但不需要和现有慢读来源一起持续堆积；若选择文字版，不再订阅播客版。"
  reuse_rule: "慢读来源要防止主题堆积；文字与播客同品牌二选一。"
  evidence_locator: "tophub-ai-reassessment-pages-16-18-20260705.md#TheGradient; tophub-ai-final-closure-20260705.md#TheGradient"
  next_action: "按篇或慢读候选。"

- source_id: "AINOW"
  current_route: slow_read_institutional_critique_candidate
  decision_status: recovered_from_pr
  selected_reason: "劳动、平台、军事、游说和公共利益视角明确，补充工程来源忽略的制度与政治经济问题。"
  rejected_alternatives:
    - "进入本轮四个 Folo 名额"
    - "替代 Rest of World"
  route_reason: "机构立场明确，更新包含招聘、证词和组织动态；当前 Rest of World 更适合作为持续关系，AINOW 作为制度批判入口保留。"
  reuse_rule: "批判机构源是视角，不是唯一解释；和工程、地区、产业来源并读。"
  evidence_locator: "tophub-ai-reassessment-pages-13-15-20260705.md#AINOW; tophub-ai-final-closure-20260705.md#AINOW"
  next_action: "按需或慢读。"

- source_id: "fast.ai / Nicholas Carlini / Max Woolf / David Stutz / One Useful Thing / The Intrinsic Perspective"
  current_route: high_value_authors_on_demand_or_links
  decision_status: recovered_from_pr
  selected_reason: "这些作者都有原创价值：fast.ai 连接实践教育和社会判断；Nicholas Carlini 低频安全与评测；Max Woolf 实验与反例；David Stutz 评测方法；One Useful Thing 工作与管理；The Intrinsic Perspective 认知、教育和 AI。"
  rejected_alternatives:
    - "全部进入 Folo"
    - "用其中一个替代最终四个角色"
  route_reason: "最终四个候选已经覆盖实际工具、生产系统、反神话制度判断和美国之外地区经验。这些作者主题角色被覆盖或不够当前核心，保留为沃壤链接或按篇任务源。"
  reuse_rule: "高价值作者也要有明确不可替代角色；没有当前长期角色就不要加入未读队列。"
  evidence_locator: "tophub-ai-reassessment-pages-07-09-20260705.md#fastai; #OneUsefulThing; tophub-ai-reassessment-pages-10-12-20260705.md#NicholasCarlini; tophub-ai-reassessment-pages-13-15-20260705.md#MaxWoolf; #DavidStutz; tophub-ai-reassessment-pages-16-18-20260705.md#TheIntrinsicPerspective; tophub-ai-final-closure-20260705.md#慢读候选"
  next_action: "沃壤链接或按篇。"
```

---

## 7. 周报、中文来源与播客为什么让位

```yaml
- source_id: "AIGC Weekly / TheSequence / The Rundown AI / The Decoder / AI 日报快讯热榜"
  current_route: rejected_or_on_demand_ai_digest_sources
  decision_status: recovered_from_pr
  selected_reason: "AIGC Weekly 周频、中文、低负担，曾进入比较池；TheSequence 结构优于一般 AI 日报。"
  rejected_alternatives:
    - "与 Last Week in AI 双周报"
    - "新增 AI 日报 / 快讯 / 热榜"
  route_reason: "已有 Folo `Last Week in AI`，因此不再加入第二套 AI 周报。AIGC Weekly 只有在中文周报需求明确时重新比较；其他 AI 日报、快讯和热榜继续排除。"
  reuse_rule: "周报是节奏关系，不是内容越多越好；已有一个可用周报时，新周报必须证明不可替代。"
  evidence_locator: "tophub-ai-reassessment-pages-01-03-20260705.md#AIGCWeekly; tophub-ai-final-closure-20260705.md#AI周报和中文来源"
  next_action: "按需。"

- source_id: "宝玉的分享 / Hugging Face 中文博客"
  current_route: on_demand_chinese_translation_and_low_friction_alternatives
  decision_status: recovered_from_pr
  selected_reason: "宝玉的分享贴近 Codex、Agent、MCP、Skill、工程团队和 AI 产品管理；Hugging Face 中文博客降低官方工程内容的中文阅读门槛。"
  rejected_alternatives:
    - "进入最终四个 Folo 名额"
    - "Hugging Face 中英文双订阅"
  route_reason: "宝玉保留为中文翻译和产品观察入口，遇到具体英文长文、模型发布或产品讨论时按需查看。Hugging Face 中文博客只作为英文 Blog 的低负担替代入口，不与英文 Blog 双订阅；本轮两者都不进入持续流。"
  reuse_rule: "中文转述源要看原始链接、过度转述和立场偏移；中文替代源不能和英文原始源双订阅。"
  evidence_locator: "tophub-ai-reassessment-pages-01-03-20260705.md#宝玉的分享; #HuggingFace中文博客; tophub-ai-final-closure-20260705.md#AI周报和中文来源"
  next_action: "按需。"

- source_id: "AI 播客候选"
  current_route: no_podcast_now_future_three_way_choice
  decision_status: recovered_from_pr
  selected_reason: "The TWIML AI Podcast、Practical AI、AI Engineering Podcast 都有质量，分别偏研究产业访谈、较容易进入的工程工具讨论、生产系统和基础设施。"
  rejected_alternatives:
    - "本轮新增播客"
    - "三档播客全部加入"
    - "用音频填充已经足够丰富的文字来源"
  route_reason: "当前无固定英语播客流程，单集时间长，文字来源已覆盖大部分主题。未来若建立音频流程，三者最多选一个。"
  reuse_rule: "播客必须先有真实收听流程；没有流程就按集检索，不变成未听债务。"
  evidence_locator: "tophub-ai-reassessment-pages-16-18-20260705.md#未来若建立音频流程再三选一; tophub-ai-final-closure-20260705.md#播客结论"
  next_action: "未来建立音频流程时再三选一。"
```

---

## 8. 最终复用规则

1. AI 来源不按“AI”这个大词路由，而按真实功能路由：工具实验、生产系统、制度判断、地区经验、研究入口、官方工程、论文、周报、播客、任务文档。
2. TopHub 不新增 AI 专属节点；AI 发布、融资、跑分和情绪标题已由现有科技雷达和公共温度显影。
3. Folo 新增只看少数互补关系，不把高质量候选全部订阅。
4. Simon / Eugene / AI as Normal Technology / Rest of World 是四个不同角色，不是四个“好 AI 源”。
5. Chip Huyen 是 Simon 的低频替补；不把 Simon、Eugene、Chip 三者全部叠加。
6. 专业研究入口按具体问题调用；论文流、模型榜、数据集榜、Spaces 榜不是未读流。
7. 中文转述和中文博客是低负担替代或按需入口，不和英文原始源双订阅。
8. 播客必须先有收听流程；没有流程就不新增。
9. 追踪器只收单一产品、项目、漏洞、论文、公司事件或有限对象，不收 AI / 大模型 / Agent 等宽词。

---

## 9. 仍需继续追回

AI reassessment 六批和 final closure 已补成压缩 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前科技补充 ledger；
- `开发目录` 290 节点需要单独 development ledger 展开；
- 社区、财经、购物、政务、校务、专栏、浏览器与链接仍未追回完。
