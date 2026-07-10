# PR #372 判断链追回：最终收口残余桥接（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中若干 `final-closure / final-closeout` 文件的残余判断链。它们多数已被专项 ledger 覆盖，但仍有一个共同风险：旧候选、旧数量、旧页面容易被重新打开为“未完成”。

纳入文件：

- `docs/ops/tophub-ai-final-closure-20260705.md`
- `docs/ops/tophub-development-page-25-final-closure-20260705.md`
- `docs/ops/tophub-shopping-final-closure-20260705.md`
- `docs/ops/tophub-technology-final-closeout-20260704.md`
- 已有关联专项 ledger：AI reassessment、development directory、shopping page-by-page、technology 系列、Folo early operation chain。

边界说明：

- 这些文件中出现的 73、84 等数量都是阶段性真实状态，不覆盖当前 80 项 TopHub 基线。
- 本文件不重新审核目录，不更新控制文件，不改 Ready，不合并 PR。
- 本文件追回的是：为什么这些目录已经闭环、哪些候选仍只是复查候选、哪些旧候选不得重开、哪些 TopHub / Folo 分工不得混淆。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：final closure 的作用是关门，不是开启下一轮扩张

```yaml
- source_id: "final closure / final closeout 文件组"
  current_route: closure_and_temporal_boundary_records
  decision_status: recovered_from_pr
  selected_reason: "这些文件记录某个目录或大类已经审完、执行状态如何、哪些候选进入复查队列、哪些动作已经完成或未执行。"
  rejected_alternatives:
    - "把 final closure 当下一轮新增清单"
    - "用旧阶段数量覆盖当前基线"
    - "把候选队列当已执行订阅"
    - "把已闭环目录重新打开成未完成"
  route_reason: "final closure 的主要功能是建立时间边界、执行边界和复查边界。它不是新方案文件，也不是自动触发下一步动作。"
  reuse_rule: "遇到 final closure，先问：它关闭了什么、确认了什么、保留了哪些复查候选、禁止重开的是什么。"
  evidence_locator: "included final closure files"
  next_action: "作为最终收口规则。"
```

---

## 2. AI 目录：215 节点闭环，不建 AI 独立容器

```yaml
- source_id: "TopHub 人工智能 215 节点最终闭环"
  current_route: closed_no_ai_group_no_tophub_expansion
  decision_status: recovered_from_pr
  selected_reason: "AI 目录 215 个节点、18/18 页、215/215 逐项复审完成。本文件取代早期 `tophub-technology-ai-closure-20260704.md` 的目录概括初步结论。"
  rejected_alternatives:
    - "新增 AI 日报"
    - "新增 AI 快讯"
    - "新增模型热榜"
    - "新增 AI 媒体子频道"
    - "新增完整论文流"
    - "新增人工智能独立分组"
  route_reason: "AI 不是单独容器。模型发布属于科技，Coding Agent 属于开发，安全漏洞属于安全，芯片与数据中心属于数据与结构，教育、劳动和权力属于慢读与思想。重新装进 AI 热榜只会重复发布、融资、跑分和情绪标题。"
  reuse_rule: "AI 源按功能和任务路由，不按 AI 标签建新收件箱。"
  evidence_locator: "docs/ops/tophub-ai-final-closure-20260705.md#结论状态; #一句话结论; #TopHub 最终动作"
  next_action: "保持闭环。"
```

---

## 3. AI Folo 四候选：只登记，不提前修改真实 Folo

```yaml
- source_id: "AI Folo 四个候选"
  current_route: folo_review_candidates_not_executed
  decision_status: recovered_from_pr
  selected_reason: "AI 最终闭环压缩出 Simon Willison、Eugene Yan、AI as Normal Technology、Rest of World 四个互补候选。"
  rejected_alternatives:
    - "立即加入四个 Folo 来源"
    - "把 Chip Huyen 同时叠加为第五个工程候选"
    - "订阅 Stanford CRFM、SemiAnalysis、Interconnects、EleutherAI、BAIR、Hugging Face Blog、arXiv 等完整研究流"
  route_reason: "四个候选分别覆盖实际工具行为与 Coding Agent 实验、生产系统与 Evals、可靠性 / 制度条件 / 反神话判断、美国之外的技术社会与地区经验。Chip Huyen 是唯一替补，若 Simon 未读压力过大才替换，不叠加。专业研究入口知道它们，但不形成普通未读流。"
  reuse_rule: "AI Folo 候选必须互补、有限、有替代关系；研究入口按具体问题、关键词、作者、引用链或项目需要调用。"
  evidence_locator: "docs/ops/tophub-ai-final-closure-20260705.md#Folo 最终压缩; #唯一替补; #专业研究入口"
  next_action: "2026-07-17 复查。"
```

---

## 4. 开发目录：290 节点闭环，不把技术栈都变成未读债务

```yaml
- source_id: "TopHub 开发 290 节点最终收口"
  current_route: closed_no_tophub_or_folo_change_before_review
  decision_status: recovered_from_pr
  selected_reason: "开发目录 290 个节点、25/25 页、290/290 逐项判断完成，未审核节点 0。真实配置仍保持 TopHub 84、科技雷达 12、Folo 27，TopHub 新增 / 取消 / 替换 0，Folo 新增 0，追踪器新增 0。"
  rejected_alternatives:
    - "新增开发常驻节点"
    - "为 Go、前端、Node、React、数据库、GPU、嵌入式、形式化证明等建立全部未读流"
    - "把开发目录变成技术 RSS 阅读器"
  route_reason: "TopHub 负责公共注意力天气，不承担开发资讯收件箱；Folo / TopHub 候选冻结至 2026-07-17 复查；程序编程来源必须绑定真实学习、仓库或项目。"
  reuse_rule: "开发源必须绑定学习、仓库、项目、排障、选型或维护；没有当前项目就不常驻。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#一、完成状态; #三、290 节点最终架构结论"
  next_action: "保持闭环。"
```

---

## 5. DuckDB 与 xLog：好官方源也按项目启用，污染热榜整体退出

```yaml
- source_id: "DuckDB｜News"
  current_route: b_plus_official_project_source_activate_when_project_uses_duckdb
  decision_status: recovered_from_pr
  selected_reason: "DuckDB News 是官方第一方，版本、协议、格式和生态变化清晰，信噪比高，比数据库社区聚合更可复查。"
  rejected_alternatives:
    - "因为质量高就常驻"
    - "通过 TopHub Trending 或数据库聚合间接跟踪 DuckDB"
  route_reason: "用户当前没有确认 DuckDB 是长期技术栈。项目真实采用 DuckDB 时，升级为项目直接源；否则 B+ 按需。"
  reuse_rule: "官方项目源质量高不等于常驻。项目启用才升级为项目直接源。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#第 25 页"
  next_action: "按需。"

- source_id: "xLog 热榜整体"
  current_route: rejected_structurally_polluted_platform_charts
  decision_status: recovered_from_pr
  selected_reason: "xLog 热榜中有少量正常开源、Obsidian、RSS 和技术文章。"
  rejected_alternatives:
    - "保留今日最热 / 本周最热 / 本月最热 / 史上最热任一热榜"
    - "用正常文章抵消热榜污染"
  route_reason: "四个层级均出现推广、诈骗式代币文案、交易所镜像、成人或侵害性内容、破解、镜像、盗版与绕过内容。正常文章不能抵消热榜本身结构性污染。"
  reuse_rule: "平台热榜如果结构性污染，个别好文章只能按作者或具体文章处理，不能保留平台热榜。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#第 25 页"
  next_action: "永久退出候选。"
```

---

## 6. 开发候选：Golang Weekly、HF Blog、InfoQ、前端源各有边界

```yaml
- source_id: "开发最终候选组"
  current_route: folo_or_on_demand_candidates_with_project_binding
  decision_status: recovered_from_pr
  selected_reason: "Golang Weekly、Hugging Face Blog、deeplearning.ai Weekly、InfoQ Today、web.dev、Frontend Focus、CSS-Tricks、JavaScript Weekly、Node Weekly、React Status 等都有价值。"
  rejected_alternatives:
    - "立即加入 Folo"
    - "英文和中文 Hugging Face Blog 同时订阅"
    - "把 InfoQ Today 常驻"
    - "没有当前前端项目时订阅前端周刊流"
  route_reason: "Golang Weekly 最符合 Go 学习阶段，但要到 2026-07-17 复查确认；Hugging Face 英文 Blog 胜过中文博客，中文只按需；deeplearning.ai Weekly 进入 AI 候选 A2；InfoQ 有生产级架构价值但按任务查；前端源按项目启用，不常驻。"
  reuse_rule: "开发候选必须经过项目关系、语言学习阶段、更新频率和重复度复查。官方规范源通常按需，周刊源只有与当前学习 / 项目绑定才进入 Folo。"
  evidence_locator: "docs/ops/tophub-development-page-25-final-closure-20260705.md#Golang Weekly; #Hugging Face 英文 Blog; #deeplearning.ai; #InfoQ; #前端来源最终按项目启用"
  next_action: "2026-07-17 复查候选。"
```

---

## 7. 购物目录：95 节点闭环，不由促销生成购买任务

```yaml
- source_id: "TopHub 购物目录最终收口"
  current_route: closed_no_attention_flow_no_subscription_change
  decision_status: recovered_from_pr
  selected_reason: "购物目录 95/95 个节点、12/12 个内容页全部完成审核并闭环。TopHub 新增 0、取消 0、替换 0、调序 0；Folo 实际新增 0；追踪器新增 0。"
  rejected_alternatives:
    - "把购物目录做成发现可以买什么的日常信息源"
    - "订阅淘宝、京东、什么值得买、羊毛线报或信用卡热榜"
    - "从销量榜直接选择健康、母婴、医疗相关商品"
    - "没有明确需求时由促销信息生成购买任务"
  route_reason: "购物目录角色是当真实需求已经出现后，按任务调用价格、渠道、用户经验、版本、规格、权益和交付信息；任务结束后退出注意力。低价、券后价、短时销量和热度不能替代质量、适用性或真实优惠判断。"
  reuse_rule: "购物信息由需求触发，任务结束退出。促销流、线报流、信用卡流和健康商品销量榜不进入常驻注意力。"
  evidence_locator: "docs/ops/tophub-shopping-final-closure-20260705.md#状态; #最终动作; #最终架构判断; #明确排除"
  next_action: "保持闭环。"
```

---

## 8. Craig Mod：购物目录里出现的非购物 Folo 候选

```yaml
- source_id: "Craig Mod"
  current_route: folo_review_candidate_not_shopping_source
  decision_status: recovered_from_pr
  selected_reason: "Craig Mod 稳定围绕书、步行、日本地方与城市、摄影、软件、语言、LLM 和独立写作展开。"
  rejected_alternatives:
    - "加入 TopHub"
    - "立即加入 Folo"
    - "把它当价格榜或书目榜"
    - "因为来自购物目录审计就放购物路由"
  route_reason: "它不是购物源，而是作者型长文来源。进入 2026-07-17 Folo 复查名单，届时检查真实更新频率、全文可读性、未读负担和与现有来源重复度；复查不是承诺加入。"
  reuse_rule: "目录来源不决定个人路由。购物目录中出现的作者型长文，也应按作者关系和阅读价值判断。"
  evidence_locator: "docs/ops/tophub-shopping-final-closure-20260705.md#唯一新增复查候选"
  next_action: "2026-07-17 复查。"
```

---

## 9. 科技最终 closeout：真实执行证据，但旧数量已被后续覆盖

```yaml
- source_id: "TopHub 科技大类最终收口"
  current_route: executed_technology_closeout_with_superseded_stage_counts
  decision_status: recovered_from_pr_platform_verified_with_temporal_boundary
  selected_reason: "科技大类完成逐页审核、候选池总收口、用户手工执行节点调整、全部订阅节点 / 科技雷达 / 数据与结构页面核验、Apple Foundation Models 精确追踪创建与结果页检查、GitHub 真实状态文件更新。"
  rejected_alternatives:
    - "把科技 closeout 中的 73 项当当前最终总量"
    - "把评测页面当已审核"
    - "把旧候选重新打开"
  route_reason: "该文件记录当时已经发生并由页面证据确认的真实状态：删除 2、替换 4、新增 3、新增 1 个精确追踪。但后续娱乐、报刊、财经、失效节点等动作已改变总量；因此它是科技阶段执行证据，不是当前全局数量基线。"
  reuse_rule: "执行证据必须带阶段边界。科技阶段已闭环，旧候选和旧数量不得重新打开；当前总量看最新顺序文件。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#状态; #最终真实状态; #已执行动作; #闭环"
  next_action: "保持科技闭环。"
```

---

## 10. 科技 / Folo 分工：TopHub 承担公共发现，Folo 不重复

```yaml
- source_id: "科技 closeout TopHub / Folo 分工"
  current_route: tophub_technology_public_discovery_folo_no_duplicate
  decision_status: recovered_from_pr
  selected_reason: "科技 closeout 明确 TopHub 承担科技门户、企业 IT、科学、全球化、汽车与终端市场研究、设备维修计划和少量精确追踪；Folo 继续保持 27 个直接关系与观察来源。"
  rejected_alternatives:
    - "Folo 重复订阅 The Register、Science Magazine、Counterpoint 或 Apple 维修计划"
    - "用 RSSHub 承接快讯、日报、榜单、报告库、限免和论坛"
    - "启用追踪机器人、频道或自动刷新"
  route_reason: "TopHub 低成本公共发现不制造未读债务；Folo 保持少量直接关系；按需网页处理电商、教育、报告、快讯、评测、硬件购买、竞赛和平台政策。"
  reuse_rule: "已经由 TopHub 承担公共雷达的来源，不在 Folo 重复订阅；Folo 只在需要长期原文关系或多媒体原貌时介入。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#三个固定问题的最终答案; #Folo后续"
  next_action: "2026-07-17 复查 Folo。"
```

---

## 11. 本补充 ledger 的复用规则

1. final closure 用来关门，不是开下一轮扩张。
2. AI 目录已 215/215 闭环，不建 AI 独立容器。
3. AI Folo 候选只有四个互补名额，Chip Huyen 是替补，不叠加。
4. 专业研究入口按具体问题调用，不形成普通未读流。
5. 开发目录已 290/290 闭环，不把所有技术栈建成未读债务。
6. 官方项目源质量高也要项目采用才升级常驻。
7. 结构性污染平台热榜整体退出，个别好文章按作者或具体文章处理。
8. 开发候选要绑定当前学习、项目、语言阶段和未读压力复查。
9. 购物目录已 95/95 闭环，购物信息由真实需求触发，不由促销触发购买任务。
10. 目录来源不决定个人路由，Craig Mod 是作者型长文候选，不是购物源。
11. 科技 final closeout 是阶段执行证据，旧 73 项数量不覆盖当前 80 项基线。
12. TopHub 已承担公共雷达时，Folo 不重复订阅。

---

## 12. 仍需继续追回

最终收口残余桥接已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内剩余可查项主要是完整扫描是否还有 `tophub-*final-closure`、`*-platform-execution-verified`、`*-directory-audit` 未被 ledger 覆盖；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
