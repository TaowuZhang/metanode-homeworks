# PR #372 判断链追回：最近对话第一批可转证据（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-old-chat-partial-recovery-boundary-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-recent-chat-evidence-intake-20260706.md`

本轮依据用户要求主动翻找最近 7 天 ChatGPT 对话，并通过 File Library 检索到 2026-07-06 上传的 `粘贴的文本 (1).txt`。该文本包含第一批已经能从最近对话 / 旧工作流中确认的判断链对象与部分具体理由。

本文件不是重新审查 TopHub / Folo，也不是把当前模型理由倒灌成历史判断。它只做：

1. 将 `粘贴的文本 (1).txt` 中已经明确出现的第一批可转判断链对象落账；
2. 标注这些内容后来已被哪些专项 ledger 承接；
3. 标注哪些仍需要原始旧对话或源文件行才能进一步逐源展开。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. `粘贴的文本 (1).txt` 的证据角色

```yaml
- source_id: "粘贴的文本 (1).txt / 第一批可转判断链"
  current_route: recent_chat_uploaded_evidence_first_batch
  decision_status: old_chat_evidence_recovered_via_file_library
  selected_reason: "该文件明确说第一批可直接转为判断链的对象包括 AI、设计、购物、财经四组，并列出具体来源、角色、去向和部分拒绝项。它还明确说明这不是最终 ledger，只是标明这些已经能直接转。"
  rejected_alternatives:
    - "继续只说旧对话原文缺失"
    - "把这份文件当全部旧对话完整迁移"
    - "把它列出的对象全部重新审查一遍"
    - "把它的流程建议当用户授权的新 CI / MVP 方案"
  route_reason: "该文件提供的不是所有旧对话全文，而是一份可用的旧工作流遗产：它能证明哪些对象已被早期识别为可转判断链，并给出部分 why_in / why_out / 去向。"
  reuse_rule: "把该文件作为 first-batch old-chat evidence intake；只抽已出现的判断，不扩写未出现的细节。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt, created 2026-07-06T13:03:23Z"
  next_action: "作为最近对话第一批证据。"
```

---

## 2. 判断链字段格式已在旧工作流中出现

```yaml
- source_id: "decision recovery entry format"
  current_route: format_constraint_recovered_from_recent_chat
  decision_status: old_chat_evidence_recovered_via_file_library
  selected_reason: "上传文本列出判断链字段：source_id、current_route、source_file、decision_status、selected_reason、rejected_alternatives、route_reason、ordering_reason、next_review_or_use_condition、evidence_ref。"
  rejected_alternatives:
    - "只用自然段总结"
    - "只列来源名和去向"
    - "只列最终数量"
    - "把未追回项塞 not_found 占位"
  route_reason: "这个格式不是为了机械填表，而是为了保证每条判断链至少回答：当前去向、为什么选或保留、为什么不是同类其他来源、为什么放到 TopHub / Folo / 沃壤 / 按需 / 退出、证据来自哪里。"
  reuse_rule: "后续旧对话证据一旦进入，按同一字段抽取；字段缺失则标 `needs_source_text`。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 判断链条目格式"
  next_action: "作为 ledger 字段约束。"
```

---

## 3. TopHub AI 独立分组：旧证据确认不建立

```yaml
- source_id: "TopHub AI 独立分组"
  current_route: rejected_no_ai_group
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写：TopHub 不建立 AI 独立分组。理由是 AI 不是单独容器，模型发布属于科技，Coding Agent 属于开发，安全漏洞属于安全，芯片与数据中心属于数据与结构，教育、劳动和权力属于慢读与思想。"
  rejected_alternatives:
    - "AI 日报"
    - "AI 快讯"
    - "模型热榜"
    - "AI 媒体子频道"
    - "完整论文流"
    - "人工智能独立分组"
  route_reason: "把这些重新装进 AI 热榜，只会重复发布、融资、跑分和情绪标题；当前科技雷达已经持续显影 AI。该判断后来已由 AI reassessment ledger 完整承接。"
  reuse_rule: "遇到 AI 来源先按功能拆分，不按 AI 标签建新收件箱。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / TopHub AI 独立分组; docs/ops/pr372-decision-recovery-ledger-ai-reassessment-compression-20260706.md"
  next_action: "已由 AI 专项承接。"
```

---

## 4. AI Folo 候选：Simon、Eugene、AI as Normal Technology、Rest of World、Chip Huyen

```yaml
- source_id: "Simon Willison's Weblog"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写 Simon 是 Folo 复查候选，角色是实际工具行为与 Coding Agent 实验；理由是亲自测试模型、API、Agent 和开源工具，重视代码、复现实验和实际失败，同时覆盖 AI 编程、数据、安全与 Web。"
  rejected_alternatives:
    - "Chip Huyen 与 Simon / Eugene 同时叠加"
    - "一般 AI 工程教程"
  route_reason: "它与 Codex、GitHub PR、本地 Agent 工作流直接相关；但更新频繁，应放触发 / 实验观察，不是篇篇必读。"
  reuse_rule: "工程型 AI 作者要看是否有真实实验、失败记录和工具行为，而不是泛教程。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / Simon Willison; AI reassessment ledger"
  next_action: "2026-07-17 复查。"

- source_id: "Eugene Yan"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写 Eugene 是 Folo 复查候选，角色是生产系统、Evals、搜索推荐与团队机制来源；覆盖 Product Evals、LLM-as-Judge、长上下文、安全评测、推荐系统、搜索、MCP、新闻 Agent、真实机器学习系统设计与团队协作。"
  rejected_alternatives:
    - "一般 AI 工程教程"
  route_reason: "它补生产系统和组织机制，不是工具试用作者。与 Simon 的实际工具行为形成互补。"
  reuse_rule: "生产系统型来源要补真实部署、评测、组织协作和系统设计，不与单纯教程混淆。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / Eugene Yan; AI reassessment ledger"
  next_action: "2026-07-17 复查。"

- source_id: "AI as Normal Technology"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写 AI as Normal Technology 是 Folo 复查候选，角色是可靠性、制度条件与反神话判断来源；讨论 Agent 可靠性与开放世界评测，核验 Agent 建成操作系统、AI 替代工程师等宏大说法，把 AI 放回组织、劳动、法律、科学和制度。"
  rejected_alternatives:
    - "厂商叙事"
    - "AGI 里程碑叙事"
    - "x-risk 单一路线"
  route_reason: "它抵消厂商叙事、AGI 里程碑叙事和 x-risk 单一路线，但不能成为新的默认裁决。"
  reuse_rule: "反神话来源用于校正叙事，不授权其成为唯一裁判。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / AI as Normal Technology; AI reassessment ledger"
  next_action: "2026-07-17 复查。"

- source_id: "Rest of World"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写 Rest of World 是美国之外的技术社会与地区经验来源；覆盖印度、中国、全球南方、跨国平台和供应链、数据劳动、移民技术劳工、教育和地区产业变化。"
  rejected_alternatives:
    - "36氪出海替代地区经验"
    - "美国中心英语科技媒体覆盖一切"
  route_reason: "它补足现有英语科技媒体的美国中心倾向，观察技术怎样进入具体地区和普通人的生活。"
  reuse_rule: "地区经验来源不能被出海商业报道或美国中心媒体完全替代。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / Rest of World; AI reassessment ledger"
  next_action: "2026-07-17 复查。"

- source_id: "Chip Huyen"
  current_route: folo_review_replacement_candidate
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本明确写 Chip Huyen 是生成式 AI 平台、Agent、模型路由与生产架构来源，本身值得长期阅读。"
  rejected_alternatives:
    - "与 Simon Willison 和 Eugene Yan 同时进入"
  route_reason: "同时加入会造成工程源叠加；若 Simon 更新过密，则用 Chip 替换 Simon。Folo 最终仍保持增加 4 个，而不是 5 个。"
  reuse_rule: "候选之间要有替代关系；不能因为都好就叠加。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / Chip Huyen; AI reassessment ledger"
  next_action: "替补候选。"
```

---

## 5. AI 专业研究入口：知道它们，但不做普通未读流

```yaml
- source_id: "Stanford CRFM / SemiAnalysis / Interconnects / EleutherAI Blog / Hugging Face Blog / The Gradient / AINOW / fast.ai / Nicholas Carlini"
  current_route: professional_or_on_demand_research_sources
  decision_status: old_chat_evidence_recovered_superseded_by_ai_ledger
  selected_reason: "上传文本把这些列入第一批可转判断链对象，作为 AI 目录中专业研究、产业分析、研究博客、工程或安全相关来源。"
  rejected_alternatives:
    - "完整研究流进入 Folo"
    - "完整论文洪流进入 TopHub"
    - "把所有 AI 专业源都当常驻阅读"
  route_reason: "这些来源有专业价值，但多数适合具体问题、关键词、作者、引用链或项目需要时调用；普通订阅流会制造未读负担。"
  reuse_rule: "研究入口先按任务调用，只有持续低噪声、独特、真实阅读关系成立才进入 Folo 复查。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / AI 目录对象列表; AI reassessment ledger"
  next_action: "按需或专业入口。"
```

---

## 6. 设计目录：designboom 是唯一 Folo 复查候选

```yaml
- source_id: "designboom"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_design_ledger
  selected_reason: "上传文本明确写 designboom 是设计目录唯一 Folo 复查候选；其独特角色是建筑、产品、家具、材料、工艺、公共空间、展览与设计文化，补‘设计如何进入物、空间与生活’的视角。"
  rejected_alternatives:
    - "站酷 / Dribbble / Behance / 500px 的作品瀑布流"
    - "优设 / UI 中国 / 站酷文章的教程与工具清单流"
    - "腾讯 CDC / 百度用户体验中心直接常驻"
  route_reason: "designboom 更接近设计文化、材料和空间，而不是作品瀑布流或工具教程；但仍只是 Folo 复查候选，不是已订阅事实。"
  reuse_rule: "设计源要看是否提供设计如何进入物、空间、材料、公共生活的判断，而不是只看作品流量。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / designboom; stale-design-visual-nature ledger"
  next_action: "2026-07-17 复查。"
```

---

## 7. 设计拒绝 / 按需组：作品瀑布流、教程流、案例档案

```yaml
- source_id: "站酷 / Dribbble / Behance / 500px"
  current_route: on_demand_visual_wander_not_folo_relationship
  decision_status: old_chat_evidence_recovered_superseded_by_design_ledger
  selected_reason: "上传文本明确写这些是作品瀑布流，不等于持续设计判断，按需视觉漫游。"
  rejected_alternatives:
    - "作为 Folo 长期关系"
    - "作为 TopHub 设计常驻"
  route_reason: "作品流能提供视觉刺激，但不能稳定提供设计判断、材料、工艺、空间与文化脉络。"
  reuse_rule: "视觉作品站按需漫游；只有持续策展关系成立才复查订阅。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 设计目录; stale-design-visual-nature ledger"
  next_action: "按需。"

- source_id: "优设 / UI 中国 / 站酷文章"
  current_route: rejected_or_on_demand_design_tutorial_flow
  decision_status: old_chat_evidence_recovered_superseded_by_design_ledger
  selected_reason: "上传文本明确写这些来源包含大量 AI 教程、工具清单、作品集内容，不进入持续关系。"
  rejected_alternatives:
    - "进入 Folo 常驻"
    - "进入 TopHub 设计长期节点"
  route_reason: "教程和工具清单容易变成低密度工具流，不补设计文化与材料判断。"
  reuse_rule: "教程站按任务搜索，不作为持续设计关系。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 设计目录; stale-design-visual-nature ledger"
  next_action: "按需或排除。"

- source_id: "腾讯 CDC / 百度用户体验中心"
  current_route: task_or_low_frequency_case_archive
  decision_status: old_chat_evidence_recovered_superseded_by_design_ledger
  selected_reason: "上传文本明确写它们是案例与方法档案，作为任务或低频来源。"
  rejected_alternatives:
    - "直接进入 Folo 常驻"
    - "直接替代设计文化来源"
  route_reason: "企业 UX 案例有方法价值，但依赖具体任务，不承担持续设计文化入口。"
  reuse_rule: "企业设计案例按项目、产品、UX 方法任务调用。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 设计目录; stale-design-visual-nature ledger"
  next_action: "按任务或低频。"
```

---

## 8. 购物目录：Craig Mod 是作者候选，购物平台按任务或排除

```yaml
- source_id: "Craig Mod"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_shopping_ledger
  selected_reason: "上传文本把 Craig Mod 列为购物目录中第一批可转判断链对象，并与购物平台类来源分开。后续 shopping ledger 已明确它不是购物源，而是作者型长文来源。"
  rejected_alternatives:
    - "把 Craig Mod 当价格榜或书目榜"
    - "因为出现在购物目录就按购物来源处理"
    - "立即加入 Folo"
  route_reason: "目录来源不决定个人路由。Craig Mod 应按作者关系、步行、日本地方、摄影、软件、语言、LLM 和独立写作等长期阅读价值复查。"
  reuse_rule: "来源所在目录不决定它的层级；作者型长文按作者关系判断。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 购物目录; shopping ledger; final-closure residual bridge"
  next_action: "2026-07-17 复查。"

- source_id: "淘宝 / 京东 / 什么值得买多数节点"
  current_route: purchase_task_triggered_on_demand
  decision_status: old_chat_evidence_recovered_superseded_by_shopping_ledger
  selected_reason: "上传文本明确把淘宝 / 京东 / 什么值得买多数节点列为按任务处理对象；购物文件已写明购物目录不是发现可以买什么的日常信息源，而是在真实需求出现后按任务调用价格、渠道、用户经验、版本、规格、权益和交付信息。"
  rejected_alternatives:
    - "加入 TopHub 常驻"
    - "加入 Folo"
    - "用销售榜单生成购买任务"
  route_reason: "购物平台榜单服务销售转化，价格和优惠依赖账号、地区、会员和叠券条件；不能制造持续注意力流。"
  reuse_rule: "购物来源由明确需求触发，任务结束退出。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 购物目录; shopping ledger"
  next_action: "按需。"

- source_id: "羊毛线报 / 众筹 / 信用卡 / 银行卡 / 虚拟卡 / 跨境支付"
  current_route: high_risk_or_task_specific_finance_purchase_entries
  decision_status: old_chat_evidence_recovered_superseded_by_shopping_ledger
  selected_reason: "上传文本把这些列为购物目录第一批可转对象。后续 shopping ledger 已承接：羊毛线报退出持续流；众筹按需核验交付、量产、退款、售后和知识产权；信用卡 / 支付类只在具体跨境支付或卡产品任务中检索。"
  rejected_alternatives:
    - "把线报、众筹、信用卡流加入常驻"
    - "采用规避风控、漏洞、多账号或虚假资料方案"
  route_reason: "这些来源要么高噪声，要么高风险，要么高度依赖具体任务与个人条件。"
  reuse_rule: "购物与支付类来源必须任务触发，并保留风险边界。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 购物目录; shopping ledger"
  next_action: "按任务或排除。"
```

---

## 9. 财经目录：替换、复查、工具、按需和专业来源

```yaml
- source_id: "财新点击排行榜 / 财新评论排行榜"
  current_route: retired_stale_finance_ranking_nodes
  decision_status: old_chat_evidence_recovered_superseded_by_finance_ledger
  selected_reason: "上传文本明确写财新点击排行榜退出、财新评论排行榜退出。后续财经 ledger 已承接：直接事实依据是两个旧排行榜停止更新，不能继续承担持续入口。"
  rejected_alternatives:
    - "继续保留失效排行榜"
    - "只因财新品牌权威就保留"
  route_reason: "失效节点不能承担持续财经入口。"
  reuse_rule: "财经源先核更新状态；品牌权威不能覆盖节点失效。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 财经目录; finance ledger"
  next_action: "已退出。"

- source_id: "财新首页推荐 / 日经中文网每日最新"
  current_route: executed_tophub_data_structure_finance_replacements
  decision_status: old_chat_evidence_recovered_superseded_by_finance_ledger
  selected_reason: "上传文本明确写财新首页推荐进入 TopHub 数据与结构，日经中文每日最新进入 TopHub 数据与结构。后续平台核验确认两者已执行。"
  rejected_alternatives:
    - "继续使用两个财新排行榜"
    - "新增更多财经快讯或交易流"
  route_reason: "财新首页补仍在更新的国内编辑入口；日经补日本、东亚产业、供应链、货币政策视角。"
  reuse_rule: "财经新增必须补结构缺口，不补交易噪声。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 财经目录; finance platform execution ledger"
  next_action: "已执行。"

- source_id: "有知有行｜全部"
  current_route: folo_review_candidate_2026_07_17
  decision_status: old_chat_evidence_recovered_superseded_by_finance_ledger
  selected_reason: "上传文本明确写有知有行｜全部进入 Folo 复查。后续财经 ledger 已承接其角色：长期投资、基金与资产配置、保险边界、组合和再平衡、投资行为与数据解释。"
  rejected_alternatives:
    - "立即加入 Folo"
    - "选择单栏目替代全部"
    - "把投资教育内容生成投资动作"
  route_reason: "`全部`覆盖完整角色，但只能提供理解和行为尺度，不生成投资指令；执行仍放到 2026-07-17 复查。"
  reuse_rule: "投资教育源只能作为行为尺度与理解来源，不当交易动作生成器。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 财经目录; finance ledger"
  next_action: "2026-07-17 复查。"

- source_id: "华尔街见闻财经日历"
  current_route: worang_tool_link_not_feed
  decision_status: old_chat_evidence_recovered_superseded_by_finance_ledger
  selected_reason: "上传文本明确写华尔街见闻财经日历进沃壤工具链接。"
  rejected_alternatives:
    - "作为每日订阅"
    - "作为财经内容流"
  route_reason: "财经日历是查询工具，不是阅读源。"
  reuse_rule: "工具链接写沃壤，不制造未读。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 财经目录; finance ledger"
  next_action: "工具链接。"

- source_id: "Foresight News / 慢雾 / CoinDesk"
  current_route: web3_structure_security_and_original_news_on_demand
  decision_status: old_chat_evidence_recovered_superseded_by_finance_ledger
  selected_reason: "上传文本明确写 Foresight News 按需、慢雾写入沃壤专业来源、CoinDesk 按需。"
  rejected_alternatives:
    - "Web3 快讯 / 价格预测 / 巨鲸 / 杠杆 / 爆仓流"
    - "把交易流加入常驻"
  route_reason: "Web3 要把技术、安全、监管、行业结构与币价交易流彻底分开。慢雾作为安全 / 技术研究专业来源，Foresight 与 CoinDesk 作为按需结构或英文原始报道入口。"
  reuse_rule: "Web3 只保留结构、安全、监管与专业来源；交易流排除。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 财经目录; finance ledger"
  next_action: "按需 / 专业来源。"
```

---

## 10. 平台核验与执行文件：证明动作，不证明价值

```yaml
- source_id: "platform execution verification files"
  current_route: execution_evidence_not_value_judgment
  decision_status: old_chat_evidence_recovered_superseded_by_pending_verification_ledger
  selected_reason: "上传文本明确列出 finance platform execution、newspapers platform execution、stale node removal verified、all page order audit，并写明这些文件只证明动作是否真的发生，不作为价值判断主来源。"
  rejected_alternatives:
    - "把建议添加写成已经添加"
    - "用平台核验文件解释来源价值"
    - "把执行证据与候选理由混在一起"
  route_reason: "平台核验负责标记真实执行、建议 / 候选 / 复查，不能替代为什么选择或退出的价值判断。"
  reuse_rule: "平台核验文件回答 did_it_happen；why_it_matters 回到 closure / ledger / source evidence。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 平台核验与执行文件; pending-actions platform verification ledger"
  next_action: "作为状态证据。"
```

---

## 11. 浏览器 / 链接 / 文化内容分工：旧证据确认待抽事项

```yaml
- source_id: "浏览器 / 链接 / 文化内容分工"
  current_route: routing_reasons_recovered_and_later_expanded
  decision_status: old_chat_evidence_recovered_superseded_by_resource_and_browser_ledgers
  selected_reason: "上传文本明确列出浏览器语境、浏览器工具架、浏览器订阅候选、主题包候选、项目入口候选、文化内容入口、得到谱系，并写明要抽：浏览器保存抵达，Folo / TopHub 保存到来，沃壤保存判断；得到 / 小宇宙 / 豆瓣 / BibiGPT 为什么保留原平台价值；普通摘要、完整个人库、平台收藏、未形成判断的记录不迁入 GitHub。"
  rejected_alternatives:
    - "把所有链接迁入 GitHub 正式资源"
    - "把所有文化内容统一搬进 TopHub / Folo"
    - "把普通摘要和平台收藏当沃壤沉积"
  route_reason: "这些是路由判断而非普通订阅源判断。后续 browser-links、resource-device-cultural-coordinate、culture-knowledge-audio-video、browser-link-candidate-cohorts 等 ledger 已展开承接。"
  reuse_rule: "链接保存抵达，订阅系统保存到来，沃壤保存判断；平台原生价值未被沃壤替代。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 浏览器与文化内容分工; browser/resource/culture ledgers"
  next_action: "已由后续 ledger 承接。"
```

---

## 12. 验收条件：不能只因数字正确就合并

```yaml
- source_id: "first batch recovery acceptance conditions"
  current_route: validation_constraints_for_pr372_recovery
  decision_status: old_chat_evidence_recovered_via_file_library
  selected_reason: "上传文本明确写：不能只因为 TopHub 80 / Folo 27 数字正确就合并；真实账本要说明现在有什么；判断链要说明为什么是这些；候选池要说明为什么只是候选；拒绝项要说明为什么不选同页其他来源；平台核验文件要说明什么动作真的发生；找不回理由的条目不能冒充完整沉积。"
  rejected_alternatives:
    - "数字正确就合并"
    - "真实账本替代判断账本"
    - "候选池替代复查理由"
    - "平台核验替代价值判断"
    - "找不回理由也写完成"
  route_reason: "这些验收条件直接回应用户的核心问题：需要判断链，不需要漂亮总论。"
  reuse_rule: "任何后续完成声明都必须逐项通过这些验收条件。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt / 验收条件"
  next_action: "作为验收护栏。"
```

---

## 13. 本补充 ledger 的复用规则

1. `粘贴的文本 (1).txt` 是最近对话 / 旧工作流证据，不是完整旧对话原文。
2. 它能确认第一批可转判断链对象：AI、设计、购物、财经，以及平台核验、浏览器 / 文化分工、验收条件。
3. 已被后续专项 ledger 承接的条目标 `superseded_by_*_ledger`，不重复当新当前状态。
4. 只抽文本中已经出现的理由和拒绝项，不扩写未出现的细节。
5. AI 不建独立组；AI Folo 候选四个互补，Chip 为替补。
6. 设计只把 designboom 作为唯一 Folo 复查候选；作品瀑布流和教程流按需或排除。
7. 购物不做日常发现，Craig Mod 作为作者型候选；平台、线报、众筹、支付按任务或排除。
8. 财经替换和候选已清楚：财新两个榜退，财新首页 / 日经进， 有知有行复查，财经日历进工具，Web3 分结构 / 安全 / 交易流。
9. 平台核验只证明动作，不证明价值。
10. 浏览器、得到、小宇宙、豆瓣、BibiGPT 的原平台价值已在该文本中成为明确待抽事项，后续已被资源 / 文化 ledgers 承接。
11. 不能因为 TopHub 80 / Folo 27 数字正确就宣布判断链完成。

---

## 14. 仍需继续追回

最近对话第一批可转证据已落账。

继续待办：

- 主动按最近 7 天主题继续查：PR #372 的后续对话、浏览器清理、清单系统、晚间见闻、任务设置；
- 若用户提供完整旧对话导出或更精确对话标题，可建立 `old-chat-evidence-source-level-*`；
- 不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
