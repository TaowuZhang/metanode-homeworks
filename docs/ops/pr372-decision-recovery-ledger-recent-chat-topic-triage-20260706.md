# PR #372 判断链追回：最近 7 天对话主题分流（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-recent-chat-topic-index-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-recent-chat-first-batch-evidence-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-browser-link-candidate-cohorts-20260706.md`

本轮根据用户要求继续主动翻找最近 7 天 ChatGPT 对话，将相邻主题分流为：

1. `source_routing`：可纳入 PR #372 的来源 / 订阅 / 链接路由判断链；
2. `workflow_automation`：属于 CI、自动化、工作流、日历、邮件扫描、复查节奏；
3. `content_generation_boundary`：属于晚间见闻、歌单、封面、日报、内容生成能力边界；
4. `personal_philosophy_or_method`：属于清单、判断权、协作方式、方法论，不强行塞进订阅源账本。

本文件不是旧对话全文迁移，不宣布逐源完成。它的作用是：防止把所有最近对话都塞进 PR #372，也防止把真正属于 source routing 的材料漏掉。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：先分流，再落账

```yaml
- source_id: "最近 7 天相邻主题分流"
  current_route: topic_triage_before_source_ledger
  decision_status: recent_chat_topic_triage_recovered
  selected_reason: "最近 7 天对话同时包含 PR #372 来源路由、浏览器链接治理、TopHub/Folo 设置、CI 失败、自动化任务、晚间见闻、清单系统等主题。"
  rejected_alternatives:
    - "全部纳入 PR #372 source ledger"
    - "全部排除为无关"
    - "把 workflow / automation / content generation 写成来源价值判断"
    - "把清单哲学写成订阅源路由"
  route_reason: "PR #372 只应承接来源、订阅、链接、TopHub/Folo/沃壤路由的判断链；其他主题可另入 ops / workflow / content / method ledger。"
  reuse_rule: "每个最近对话主题先判断它回答的是来源去向、机器怎么跑、内容怎么生成，还是协作方法。只有来源去向才进入 PR #372 source ledger。"
  evidence_locator: "recent ChatGPT conversation context 2026-07-03 to 2026-07-06; recent-chat-topic-index ledger"
  next_action: "作为 topic triage 规则。"
```

---

## 2. 明确属于 source_routing 的主题

```yaml
- source_id: "PR #372 / TopHub / Folo / 沃壤三层分工"
  current_route: source_routing
  decision_status: recovered_from_recent_chat_and_files
  selected_reason: "订阅系统 v0、TopHub/Folo/沃壤三层分工、来源管道、来源状态 active / trial / standby / on_demand / retired_candidate 都直接回答来源如何进入、观察、常驻、待机、按需或退出。"
  rejected_alternatives:
    - "归入 workflow_automation"
    - "只当 PR 工程问题"
    - "只当个人信息管理方法"
  route_reason: "这些内容直接决定来源放 TopHub、Folo、沃壤、按需或退出，是 PR #372 的主线。"
  reuse_rule: "涉及来源状态和层级分工的材料，优先纳入 PR #372 判断链 ledger。"
  evidence_locator: "recent ChatGPT context; `订阅系统-v0.md`; system routing ledgers"
  next_action: "继续作为 PR #372 主线。"

- source_id: "浏览器链接治理中的 Folo / TopHub 候选"
  current_route: source_routing
  decision_status: recovered_from_recent_chat_and_files
  selected_reason: "浏览器链接治理第一轮和第二轮完整清单中，许多条目被判为 Folo / TopHub 候选，理由是持续更新的来源应由订阅系统承担到来，不在浏览器长期堆媒体首页。第二轮还明确冻结至 2026-07-17 订阅复查。"
  rejected_alternatives:
    - "继续留在浏览器收藏夹"
    - "直接写入 Folo / TopHub 正式账本"
    - "把所有链接迁入资源正式条目"
  route_reason: "这类材料属于 PR #372，因为它决定浏览器、订阅系统、资源链接和沃壤之间的来源路由。"
  reuse_rule: "持续更新源先入候选队列，等复查时按一进一出和打开率判断；浏览器只保存抵达，不保存到来。"
  evidence_locator: "浏览器链接治理_第一轮候选.csv; 浏览器链接治理_第二轮完整清单.csv; browser link candidate ledgers"
  next_action: "已由 browser-link-candidate-cohorts 承接。"

- source_id: "TopHub 功能与设置"
  current_route: source_routing_with_system_boundary
  decision_status: recovered_from_recent_chat_and_ledgers
  selected_reason: "TopHub 首页、日报、动态、追踪、榜中榜、热文库、话题、日历、通知机器人、过滤、投稿节点等都影响来源如何被看见、追踪、通知和过滤。"
  rejected_alternatives:
    - "把平台功能全部启用"
    - "把通知机器人和外部通知作为默认能力"
    - "把过滤器当判断替代品"
  route_reason: "这些设置属于来源路由系统边界，不是具体来源价值本身。后续已由系统 ledgers 承接。"
  reuse_rule: "TopHub 新功能只在改善路由时启用；平台支持不等于当前系统需要。"
  evidence_locator: "recent ChatGPT context; account/widgets/feature/system ledgers"
  next_action: "已承接。"
```

---

## 3. 明确属于 workflow_automation 的主题

```yaml
- source_id: "CI Linux failures for PR #372"
  current_route: workflow_automation
  decision_status: recovered_from_recent_chat_and_gmail_context
  selected_reason: "2026-07-03 至 2026-07-04 多次出现 PR #372 / branch `feat/folo-subscription-system-v0-20260703` 的 CI Linux 失败通知。"
  rejected_alternatives:
    - "把 CI 失败写成来源为什么选 / 不选"
    - "用 CI 状态解释 TopHub / Folo 取舍"
    - "因为 CI 失败就停止判断链追回"
  route_reason: "CI 失败是工程操作状态，只说明 workflow 有失败，不说明来源价值。"
  reuse_rule: "CI / build / PR status 进入 docs/ops 或今日报状态核验，不进入 source judgment 的 why_in / why_out。"
  evidence_locator: "recent Gmail context; recent-chat-topic-index ledger"
  next_action: "ops 状态。"

- source_id: "自动化只生成差异与候选，不自动退订 / 提交"
  current_route: workflow_automation
  decision_status: recovered_from_recent_chat_and_files
  selected_reason: "最近对话和订阅系统文件确认：自动化只能生成差异、候选、manifest、报告、索引或观察记录；不能自动退订、自动提交、自动改主库。"
  rejected_alternatives:
    - "自动退订"
    - "自动提交信源"
    - "自动启用通知机器人"
    - "自动把候选写入 Folo"
  route_reason: "这是操作边界，不是具体来源判断。它约束工具不能接管路由权。"
  reuse_rule: "自动化产物只能进入候选层；触及主库、评价、合并、删除、沉积升级、项目 / 领域归属必须确认。"
  evidence_locator: "recent ChatGPT context; resource-update-gate constitution; Folo early operation chain"
  next_action: "ops / automation 边界。"

- source_id: "2026-07-17 Folo 观察期复查"
  current_route: workflow_automation_for_source_review
  decision_status: recovered_from_recent_chat_and_files
  selected_reason: "最近对话确认已记下 2026-07-17 09:00 日本时间的 Folo 观察期复查。它约束候选来源何时复查，而不是说明候选已经被订阅。"
  rejected_alternatives:
    - "提前把候选加入 Folo"
    - "把复查事件当执行事实"
    - "没有打开率和一进一出判断就扩容"
  route_reason: "这是来源复查的时间工作流。它和 source routing 相连，但本身是 workflow / schedule。"
  reuse_rule: "所有 Folo 候选在复查日前保持候选状态；复查时先看现有 27 项，再一进一出。"
  evidence_locator: "recent ChatGPT context 2026-07-03; Folo early operation chain; browser subscription candidates"
  next_action: "2026-07-17 复查。"
```

---

## 4. 明确属于 content_generation_boundary 的主题

```yaml
- source_id: "晚间见闻 / 今日空间"
  current_route: content_generation_boundary
  decision_status: recovered_from_recent_chat_context
  selected_reason: "最近 7 天对话中，晚间见闻 / 今日空间涉及读取 Gmail、网页、GitHub，筛选真实见闻，调用 Apple Music 生成歌单草案，调用图片工具生成封面，以及模型在工具没实际运行时不应拿普通文字冒充歌单和图片。"
  rejected_alternatives:
    - "把晚间见闻整体塞进 PR #372 source ledger"
    - "把内容生成失败写成来源选择判断"
    - "把未调用工具的文字冒充 Apple Music 歌单或图片"
  route_reason: "晚间见闻是内容生成与工具能力边界，不是订阅源路由本身。只有其中具体来源如何进入见闻筛选，才可能回到 source routing。"
  reuse_rule: "内容生成任务要分离来源筛选、工具调用、输出呈现和失败边界；不要把生成工作流混入来源账本。"
  evidence_locator: "recent ChatGPT conversations 2026-07-04; user workflow corrections"
  next_action: "另立内容工作流边界，不进 PR #372 source ledger。"

- source_id: "生成歌单与封面"
  current_route: content_generation_boundary_with_tool_truthfulness
  decision_status: recovered_from_recent_chat_context
  selected_reason: "用户要求实际连接 Apple Music、显示可点击按钮、生成 1x1 图片封面，而不是只输出普通文字列表或描述。"
  rejected_alternatives:
    - "工具没运行时用文字冒充结果"
    - "把歌单草案当来源判断"
    - "把封面生成当订阅系统动作"
  route_reason: "这类任务检验工具调用和输出真实性，不属于 PR #372 来源取舍。"
  reuse_rule: "工具产物必须真实调用工具；无法调用时明确边界，不用文字冒充。"
  evidence_locator: "recent ChatGPT conversations 2026-07-04; Apple Music / image generation workflow"
  next_action: "内容工具边界。"
```

---

## 5. 明确属于 personal_philosophy_or_method 的主题

```yaml
- source_id: "清单系统 / 9 维清单矩阵讨论"
  current_route: personal_philosophy_or_method
  decision_status: recovered_from_recent_chat_context
  selected_reason: "最近 7 天对话中，用户讨论清单到底是不是清单、是否值得长期维护、是在维护清单还是维护其他东西，并质疑把清单系统化成 Mental OS 的表达。"
  rejected_alternatives:
    - "直接写入 PR #372 来源判断 ledger"
    - "把清单哲学变成订阅源规则"
    - "用 AI 的系统化术语覆盖用户对清单本质的追问"
  route_reason: "这是个人方法论与协作哲学问题，不是具体来源路由。它可能影响怎样维护 ledger，但不直接说明某个来源 why_in / why_out。"
  reuse_rule: "清单 / 方法论讨论可影响 ledger 的形式和维护节奏，但不能替代来源证据。"
  evidence_locator: "recent ChatGPT conversations 2026-07-06; checklist discussions"
  next_action: "不纳入 PR #372 source ledger；必要时另立方法论记录。"

- source_id: "判断权与反接管原则"
  current_route: personal_philosophy_or_method_with_collaboration_constraint
  decision_status: recovered_from_recent_chat_context
  selected_reason: "用户反复强调不要由 AI 擅自压缩选择空间、不要把风险判断变成路由接管、不要用模型生成方案替代用户问题。"
  rejected_alternatives:
    - "把 AI 风险判断作为自动路由权"
    - "把用户问题改写成方案生成"
    - "用漂亮原则替代已经发生的判断"
  route_reason: "这是协作约束，影响所有 ledger 写法：可以显影风险和证据门槛，但不能接管判断。"
  reuse_rule: "ledger 只还原证据和取舍，不替用户决定下一步；建议与执行必须基于用户任务阶段。"
  evidence_locator: "recent ChatGPT conversations and standing collaboration constraints"
  next_action: "作为协作护栏。"
```

---

## 6. 混合主题：浏览器清理

```yaml
- source_id: "浏览器清理"
  current_route: mixed_source_routing_and_resource_workflow
  decision_status: recovered_from_recent_chat_and_files
  selected_reason: "浏览器清理同时包含来源路由和资源治理：Zen / Chrome / Safari 分工、浏览器保存抵达、Folo/TopHub 保存到来、沃壤保存判断；旧链接候选分为订阅候选、主题包、项目学习、账号服务、待复核、垂直检索。"
  rejected_alternatives:
    - "全部归入 PR #372 source ledger"
    - "全部归入浏览器偏好"
    - "全部迁入资源正式条目"
  route_reason: "其中 Folo / TopHub 候选、持续更新来源、来源路由属于 PR #372；浏览器分工、工具架、账号服务、候选链接治理属于资源 / 语境 / 链接 / ops。"
  reuse_rule: "浏览器材料先拆：订阅候选进 PR #372 复查队列；工具架和浏览器分工进资源/语境；账号服务留资源/链接；迁移证据进 docs/ops。"
  evidence_locator: "recent ChatGPT context 2026-07-04; browser link candidate ledgers"
  next_action: "已由 browser/resource ledgers 承接；仍可按具体候选继续补。"
```

---

## 7. 混合主题：任务设置 / 邮件关注扫描 / 今日报

```yaml
- source_id: "任务设置 / 邮件关注扫描 / 今日报"
  current_route: workflow_automation_with_source_inputs
  decision_status: recovered_from_recent_chat_context
  selected_reason: "最近 7 天对话中，用户设置每日简报、邮件关注扫描、今日报，并指出今日报与邮件关注扫描有重叠、今日报内容过窄、晚间见闻生成内容不足。"
  rejected_alternatives:
    - "把自动任务本身纳入 PR #372 source ledger"
    - "让任务自动修改邮件、日历、GitHub 或订阅"
    - "把今日报做成罗列动态或寻找今天可以做什么"
  route_reason: "这些属于 workflow / automation：它们可能读取来源，但主题本身是任务触发、扫描条件、闭环状态和输出边界。只有具体来源为什么进入简报 / 今日报时，才回到 source routing。"
  reuse_rule: "自动任务先写触发条件、输入边界、禁止动作、输出格式和闭环规则；不要与订阅源去留混在一起。"
  evidence_locator: "recent ChatGPT conversations 2026-07-04 to 2026-07-05"
  next_action: "另入 workflow / automation 记录。"
```

---

## 8. 本补充 ledger 的复用规则

1. 最近对话主题先分流，再决定是否进入 PR #372 source ledger。
2. `source_routing` 包括 TopHub/Folo/沃壤三层分工、订阅候选、浏览器链接中的持续更新来源、TopHub 设置中影响来源路由的功能。
3. `workflow_automation` 包括 CI 失败、自动化边界、Folo 复查日程、邮件扫描、今日报、任务设置。
4. `content_generation_boundary` 包括晚间见闻、歌单、封面、日报正文生成、工具真实调用边界。
5. `personal_philosophy_or_method` 包括清单系统、判断权、反接管原则、协作方式。
6. 浏览器清理是混合主题，要拆成 source routing / resource workflow / account-service / ops evidence。
7. 任务设置是 workflow，只有具体来源进入任务的理由才纳入 source ledger。
8. 不把相邻主题硬塞进订阅源账本，也不把真正的来源路由材料排除掉。

---

## 9. 仍需继续追回

最近 7 天对话主题分流已建立。

继续待办：

- 对 `source_routing` 中仍未逐源落账的浏览器订阅候选、TopHub 早期 28 个取消、App 六页、社区 page 1，可继续按具体来源补 `recent-chat-source-evidence-*`；
- 对 `workflow_automation` 另立任务 / 自动化边界 ledger，不混进 PR #372 source ledger；
- 对 `content_generation_boundary` 可另立晚间见闻 / Apple Music / 图片工具边界 ledger；
- 对 `personal_philosophy_or_method` 不写入订阅源账本，除非用户要求整理协作原则；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
