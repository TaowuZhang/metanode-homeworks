# PR #372 判断链追回：专栏 1—94 页快照（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `TopHub > 专栏`目录 1—94 页快照的判断链。

纳入文件：

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

边界说明：

- `专栏 3360 个`可以视为页级复审闭环，因为用户提供了第 1—94 页页面快照。
- 这不是 3360 个节点逐条在线核验，也不是平台执行记录。
- 本轮没有在 TopHub 或 Folo 平台真实执行订阅、取消、替换或调序。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 专栏总规则：候选池，不是第二个未读箱

```yaml
- source_id: "TopHub > 专栏 3360 个页级快照整体"
  current_route: page_level_closed_no_tophub_expansion_no_folo_actual_add
  decision_status: recovered_from_pr_snapshot
  selected_reason: "1—94 页覆盖技术作者、安全、AI、科学、医学、设计、播客、慢读、官方博客、个人博客、知识管理、旅行、生活、资源、破解、灰产等混合来源。专栏目录的真实价值在候选发现。"
  rejected_alternatives:
    - "把专栏目录逐项加入 TopHub"
    - "把 3360 个节点写成逐条实时核验完成"
    - "把候选池一次性变成 Folo 订阅池"
    - "把出现 Agent / Codex / FOLO / Claude / OpenClaw 等关键词的来源直接加入"
  route_reason: "最终 TopHub 新增 0、取消 0、替换 0、调序 0；Folo 真实新增 0。当前已有 `专栏｜订阅聚合`作为系统聚合入口，具体专栏源应进入 Folo 复查池或按任务触发，不应逐个变成 TopHub 日常未读流。"
  reuse_rule: "以后遇到专栏来源，先判定它是作者关系、官方产品流、专业研究、节目源、教程源、论坛/聚合、灰产/资源流还是任务工具。TopHub 保留系统聚合入口；Folo 只收少量长期可跟随来源；其余按任务。"
  evidence_locator: "docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏1—94页总收口"
  next_action: "2026-07-17 Folo 复查时统一去重、压缩和排序。"
```

---

## 2. 为什么不扩 TopHub

```yaml
- source_id: "专栏｜订阅聚合"
  current_route: keep_existing_tophub_system_aggregate
  decision_status: recovered_from_pr_snapshot
  selected_reason: "它比任一单个专栏更适合承担 TopHub 中的系统级专栏聚合入口。"
  rejected_alternatives:
    - "从 94 页里挑多个个人源加入 TopHub"
    - "把技术 / 安全 / AI / 科学 / 播客 / 设计分别加成 TopHub 专栏常驻"
  route_reason: "专栏目录的单个节点大多是 RSS、个人博客、官方博客、节目或专业源，适合 Folo 或任务触发。TopHub 的职责是雷达和分组入口，不是把候选池变成未读仓库。"
  reuse_rule: "如果某目录已有上位聚合入口，单个来源默认不进 TopHub；只有当某来源补出系统级缺口且已实际平台添加，才更新真实顺序。"
  evidence_locator: "docs/ops/tophub-columns-pages-01-10-20260706.md#是否加入TopHub今日热榜; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#TopHub总判断"
  next_action: "保留，不调序。"
```

---

## 3. Folo 总规则：候选池必须去重压缩

```yaml
- source_id: "专栏 Folo 复查总池"
  current_route: folo_review_pool_not_subscription_pool
  decision_status: recovered_from_pr_snapshot
  selected_reason: "专栏是 Folo 候选的主要来源，包含长期作者、稳定专业源、节目型源、工程实践源和公共叙事源。"
  rejected_alternatives:
    - "批量订阅所有优先候选"
    - "按页面出现顺序加入"
    - "同一主题池多个来源并收"
  route_reason: "最终应从 94 页里只挑少量长期会跟随的作者、节目、工程源和公共叙事源进入 2026-07-17 Folo 复查，而不是把候选池一次性变成订阅池。AI/Agent、医学期刊、技术博客、个人随笔尤其需要去重。"
  reuse_rule: "Folo 收的是长期关系，不是目录条目。复查必须看更新频率、全文可读性、重复度、用户实际阅读/收听意愿、是否有不可替代角色。"
  evidence_locator: "docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#Folo总判断; #专栏总候选池的收束方向"
  next_action: "下一步只做 Folo 候选去重、压缩和排序。"
```

---

## 4. AI / Agent 工程池：强相关也不能机械加入

```yaml
- source_id: "AI / Agent 工程与边界候选池"
  current_route: folo_review_pool_merge_with_ai_reassessment
  decision_status: recovered_from_pr_snapshot
  selected_reason: "专栏 94 页出现大量与当前工作流强相关的 AI / Agent / Coding Agent / RAG / 评估 / 模型工程源，包括 Simon Willison、Latent.Space、Interconnects、Chip Huyen、Eugene Yan、Addy Osmani、高策、像清水一般清澈透明、Bryan's Blog、LiuShen's Blog、银河美术馆、LangChain、LlamaIndex、Qdrant、Jina AI、Weights & Biases、Google Research、Microsoft Research、BAIR、CRFM、The Gradient 等。"
  rejected_alternatives:
    - "因 Agent / Codex / FOLO / OpenSpec / MCP / Claude / OpenClaw 关键词直接订阅"
    - "LangChain、LlamaIndex、Qdrant、Jina、Weights & Biases 等官方生态源全部并收"
    - "AI 播客、AI 周报、AI 研究源全部加入"
  route_reason: "进入 AI / Agent 复查池，但必须与已完成的 AI reassessment 四候选去重。Simon、Eugene、AI as Normal Technology、Rest of World、Chip Huyen 等已经有角色定义；专栏中后续发现的同类作者和官方源只能补充或替换，不得叠加成 AI 未读箱。"
  reuse_rule: "AI/Agent 来源按角色分：实际工具实验、生产系统、反神话制度判断、地区社会经验、官方框架、RAG/数据管线、研究评测、教程、播客。每个角色只保留少数互补源。"
  evidence_locator: "docs/ops/tophub-columns-pages-41-50-20260706.md#优先候选; docs/ops/tophub-columns-pages-51-60-20260706.md#优先候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "2026-07-17 与 AI reassessment ledger 合并去重。"

- source_id: "Claude / Anthropic 生态相关个人源"
  current_route: cautious_observation_only_not_default_recommendation
  decision_status: recovered_from_pr_snapshot
  selected_reason: "部分源会记录 Claude Code、第三方 Provider、OpenClaw、龙虾、Agent 工程和模型使用边界，作为能力边界或反面材料可能有观察价值。"
  rejected_alternatives:
    - "因为实践内容相关就默认推荐 Claude / Anthropic 生态"
    - "把 Claude 生态源作为默认 AI 工程基线"
  route_reason: "专栏总收口已明确：涉及 Claude / Anthropic 生态的源不默认推荐；若保留，只作为能力边界、工程实践或反面材料观察。"
  reuse_rule: "与用户明确反感的模型/生态相关的来源只能按具体材料价值处理，不把生态本身升级为推荐。"
  evidence_locator: "docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#Folo总判断; docs/ops/tophub-columns-pages-81-90-20260706.md#优先候选"
  next_action: "谨慎观察。"
```

---

## 5. 编程 / 工程实践池：项目和技术栈优先

```yaml
- source_id: "编程 / 工程实践候选池"
  current_route: folo_review_or_project_triggered_engineering_sources
  decision_status: recovered_from_pr_snapshot
  selected_reason: "候选包括 Tony Bai、Martin Kleppmann's blog、Filippo.io、MaskRay、Wang Fenjin's Blog、岁寒、蛮荆、Kenvix's Blog、Yuexun's Blog、people.kernel.org Reader、LWN.net、web.dev、Addy Osmani、AWS / Google / Cloudflare / Vercel / Meta / Netflix / Slack / Dropbox / Sentry 等工程源。"
  rejected_alternatives:
    - "把全部官方工程博客加入 Folo"
    - "把 Google、AWS、Azure、Meta、Netflix、Slack、Cloudflare 等厂商博客全部变成日常流"
    - "把语言、框架、数据库、云原生版本流全部常驻"
  route_reason: "这些源质量高，但多数绑定具体技术栈、项目阶段或厂商生态。它们适合 Folo 复查池或项目触发，而不是 TopHub。开发目录 ledger 已确认：技术栈未采用不订阅，项目采用后直接跟官方 blog、release、changelog、仓库或规范。"
  reuse_rule: "工程来源按项目阶段路由：学习、实现、测试、部署、架构、可观测性、平台、版本治理。没有当前项目关系就按需或观察。"
  evidence_locator: "docs/ops/tophub-columns-pages-61-70-20260706.md#优先候选; docs/ops/tophub-columns-pages-71-80-20260706.md#优先候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "与 development-directory ledger 合并去重。"

- source_id: "Go / 后端 / 数据系统候选"
  current_route: folo_review_or_learning_project_triggered
  decision_status: recovered_from_pr_snapshot
  selected_reason: "Tony Bai、Go 夜聊、Wang Fenjin、岁寒、蛮荆、Kenvix、Postgres Weekly、Martin Kleppmann 等与 Go 学习、数据库、后端、系统工程和数据系统相关。"
  rejected_alternatives:
    - "和 Golang Weekly 全部并收"
    - "把所有 Go / 数据系统作者加入 Folo"
  route_reason: "Golang Weekly 已在开发目录作为 Go 学习最高优先候选；个人作者和数据系统源需要看是否补 Golang Weekly 不能补的工程经验。Wang Fenjin 偏 DuckDB / DuckLake / Arrow Flight SQL；岁寒偏高并发和 MySQL 大表；蛮荆偏 Go / K8s；这些都可能按项目触发。"
  reuse_rule: "语言周刊、个人作者、项目官方源、数据库专题源四类分开。先保留一个学习节奏源，再按项目补作者或官方源。"
  evidence_locator: "docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#尾段较强候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "Folo 复查时与 Golang Weekly 比较。"
```

---

## 6. 安全 / AppSec / 逆向池：安全源不能并收

```yaml
- source_id: "安全 / AppSec / 逆向候选池"
  current_route: folo_review_pool_or_task_security_sources
  decision_status: recovered_from_pr_snapshot
  selected_reason: "强候选包括 Project Zero、Trail of Bits Blog、PortSwigger Blog、watchTowr Labs、GitHub Security Lab、Google Online Security Blog、Objective-See、Huli、Bryan's Blog、A Few Thoughts on Cryptographic Engineering、Mozilla Security、Securelist、Malwarebytes Labs、Exploit-DB、Corelan、HAHWUL、Rasta Mouse 等。"
  rejected_alternatives:
    - "漏洞数据库 / 企业安全 / 作者研究 / CTF / 攻防工具全并收"
    - "安全源全部进入 TopHub"
    - "把破解下载、灰产工具和安全研究混为一类"
  route_reason: "安全池价值高但极易过量。需要按角色选择少量：平台漏洞研究、Web/AppSec、开源供应链、macOS/Apple 安全、恶意软件/威胁情报、密码学、逆向。破解、盗版、激活、资源搬运和灰产工具明确排除。"
  reuse_rule: "安全来源先分一手研究、漏洞数据库、厂商安全、威胁情报、逆向教程、灰产/破解。只有前几类可能复查；灰产破解排除；具体 CVE / 产品 / 仓库 / 供应链事件按任务打开。"
  evidence_locator: "docs/ops/tophub-columns-pages-31-40-20260706.md#优先候选; docs/ops/tophub-columns-pages-41-50-20260706.md#优先候选; docs/ops/tophub-columns-pages-61-70-20260706.md#优先候选"
  next_action: "Folo 复查时压缩安全核心池。"
```

---

## 7. 科学 / 医学 / 数字健康池：不整包期刊化

```yaml
- source_id: "科学 / 医学 / 数字健康候选池"
  current_route: on_demand_or_folo_review_limited_science_health_sources
  decision_status: recovered_from_pr_snapshot
  selected_reason: "候选包括 Quanta Magazine、Scientific American、Science / Nature / Science Advances、JMLR、The Lancet Digital Health、The Lancet 系列、JAMA、PNAS、JCI、Science Translational Medicine、CRFM、JMLR 等。"
  rejected_alternatives:
    - "Nature / Science / Lancet / JAMA / PNAS 等整包进入 Folo"
    - "把医学期刊流作为日常阅读"
    - "用专业期刊标题替代医学/健康行动建议"
  route_reason: "科学与医学质量高但专业密度过大。final closure 已要求对医学期刊、技术博客、个人随笔只选少量入口，不整包加入。若未来建立健康研究池，也只能选 1—2 个上位入口，例如 The Lancet Digital Health 作为医学 AI / 数字健康专题候选。"
  reuse_rule: "科学/医学源按主题调用；标题不能直接转行动。健康、医学、药物、心理和诊疗问题必须回专业指南、医生、监管和原始论文。"
  evidence_locator: "docs/ops/tophub-columns-pages-21-30-20260706.md#优先候选; docs/ops/tophub-columns-pages-51-60-20260706.md#优先候选; docs/ops/tophub-columns-pages-81-90-20260706.md#优先候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#Folo总判断"
  next_action: "按任务或极少量 Folo 复查。"
```

---

## 8. 设计 / UX / 视觉文化池：灵感不能变未读债

```yaml
- source_id: "设计 / UX / 视觉文化候选池"
  current_route: folo_review_or_project_inspiration_sources
  decision_status: recovered_from_pr_snapshot
  selected_reason: "候选包括 NN/g latest articles and announcements、A List Apart、Little Big Details、Core77、Fonts In Use、ArchDaily、谷德设计网、Design Milk、Minimalissimo、Concept Art World、Figma Design、UX Booth、Inspect Element 等。"
  rejected_alternatives:
    - "ArchDaily、谷德、Design Milk、Minimalissimo、Concept Art World 等灵感源全部并收"
    - "把设计图文流加入 TopHub"
    - "把广告/软文式设计商业稿纳入系统"
  route_reason: "设计源有灵感价值，但图文量大，容易变成灵感未读池。NN/g 属于 UX / 研究方法 / AI 解释 / 业务结果判断，更像产品判断来源；建筑、视觉和字体源按项目或审美补充使用。"
  reuse_rule: "设计源分 UX 研究、前端细节、字体、建筑空间、视觉灵感、商业软文。长期 Folo 只保留少量方法型或不可替代作者/机构；灵感流按项目。"
  evidence_locator: "docs/ops/tophub-columns-pages-01-10-20260706.md#优先候选; docs/ops/tophub-columns-pages-31-40-20260706.md#优先候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "Folo 复查时压缩为极少量设计/UX源。"
```

---

## 9. 播客 / 慢读 / 公共叙事池：节目要有收听流程

```yaml
- source_id: "播客 / 慢读 / 公共叙事候选池"
  current_route: folo_review_or_slow_read_pool
  decision_status: recovered_from_pr_snapshot
  selected_reason: "候选包括 Longreads、法庭線 The Witness、Craig Mod、The MIT Press Reader、Rest of World、工劳小报、EmacsTalk、Go 夜聊、内核恐慌、Latent.Space、DataTalks.Club、The TWIML AI Podcast、随机波动、声东击西、忽左忽右、故事 FM、一天世界、文化有限、Wait But Why 等。"
  rejected_alternatives:
    - "高质量播客全部加入"
    - "英文长文源全部加入"
    - "社会议题源因重要就机械加入"
  route_reason: "慢读和播客价值高，但阅读/收听成本高。专栏总收口要求节目型源、公共叙事源进入复查池，不直接订阅。AI 播客尤其拥挤，未来若建立音频流程也只能少量选择。"
  reuse_rule: "播客必须先有真实收听流程；慢读源必须确认会读全文；公共叙事源要看主题负担、地区关系和不可替代性。"
  evidence_locator: "docs/ops/tophub-columns-pages-81-90-20260706.md#优先候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "Folo 复查时按节目/慢读池去重。"
```

---

## 10. 知识管理 / RSS / 数字花园池

```yaml
- source_id: "知识管理 / RSS / 数字花园候选池"
  current_route: folo_review_or_project_triggered_knowledge_infra_sources
  decision_status: recovered_from_pr_snapshot
  selected_reason: "候选包括 Taxodium、Save The Web Project、质数人生、软通达、Jason Lee、我叫尤加利、顾宇的研习笔记、文武科技社等，涉及 Elfeed、RSS、Web 存档、第二大脑、FOLO、Obsidian、个人知识基础设施和数字花园。"
  rejected_alternatives:
    - "因为 FOLO / RSS / Obsidian 相关就全部加入"
    - "把一次性工具教程加入 Folo"
  route_reason: "这些与当前信息系统整理任务高度相关，但必须看作者是否长期输出、是否有实践深度、是否与已有小众软件/少数派/夜航船夫/系统文档重复。FOLO 相关不等于必订。"
  reuse_rule: "知识管理源要看长期实践关系和可复用方法，不看关键词。一次性教程进入项目笔记，长期作者才进 Folo 复查。"
  evidence_locator: "docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#尾段较强候选; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#专栏总候选池的收束方向"
  next_action: "Folo 复查。"
```

---

## 11. 明确排除类型

```yaml
- source_id: "专栏排除信号"
  current_route: reusable_rejection_rules
  decision_status: recovered_from_pr_snapshot
  selected_reason: "用于复用专栏判断。"
  rejected_alternatives:
    - "破解、盗版、激活、灰产、主机优惠、虚拟卡、挂机赚钱、低质 SEO、课程引流、测试站、空站、泛知乎入口、泛即刻圈子、酷安羊毛流进入系统"
  route_reason: "1—94 页多次出现破解下载、盗版、repack、激活、虚拟卡、跨境支付教程、博彩、网赚、挂机赚钱、羊毛、VPS/主机优惠、低质资源站、泛知乎、泛即刻、测试文章、Welcome/Coming soon、无长期主题个人流水、一次性工具发布、镜像下载、素材包和音效包。全部不进入 TopHub 或 Folo。"
  reuse_rule: "专栏源只要核心是灰产、破解、盗版、资源搬运、虚拟卡、薅羊毛、低质 SEO、课程引流、测试站或空站，直接排除；单篇可用也不抵消结构性污染。"
  evidence_locator: "docs/ops/tophub-columns-pages-01-10-20260706.md#不建议进入Folo的类型; docs/ops/tophub-columns-pages-41-50-20260706.md#本批明确不建议进入Folo的类型; docs/ops/tophub-columns-pages-91-94-and-closure-20260706.md#本批明确不建议进入Folo的类型"
  next_action: "作为复用规则。"
```

---

## 12. 本补充 ledger 的复用规则

1. `专栏 3360 个`是页级复审闭环，不是 3360 个节点逐条实时核验。
2. TopHub 保留 `专栏｜订阅聚合`，不从专栏页逐个扩容。
3. Folo 收长期关系，不收目录条目；候选池必须统一去重、压缩、排序。
4. AI / Agent 来源按角色去重：工具实验、生产系统、制度判断、地区经验、官方框架、RAG、研究评测、教程、播客。
5. 工程源按项目阶段和技术栈触发；官方博客质量高也不能全订。
6. 安全源按平台漏洞、AppSec、供应链、macOS、威胁情报、密码学、逆向分层，只留少数。
7. 科学/医学源不整包期刊化；健康行动回专业指南、医生、监管和原始论文。
8. 设计/视觉灵感源不能变成图文未读池；优先少量方法型或不可替代来源。
9. 播客必须先有收听流程；慢读必须确认会读全文。
10. FOLO / RSS / Obsidian / Agent / Codex 等关键词不是订阅理由。
11. 涉及 Claude / Anthropic 生态的源不默认推荐；若保留，只作能力边界、工程实践或反面材料观察。
12. 灰产、破解、盗版、虚拟卡、挂机赚钱、资源搬运和低质 SEO 排除。

---

## 13. 仍需继续追回

专栏 1—94 页已补成页级 snapshot ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 浏览器与链接；
- 报刊目录如果需要可进一步补专项 ledger；
- 娱乐目录大量 closure 仍可继续追回；
- 旧 ChatGPT 对话追索所有 partial 条目。
