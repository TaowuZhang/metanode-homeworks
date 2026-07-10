# PR #372 判断链追回：浏览器旧链接候选组（2026-07-06）

## 0. 边界

本文件追回 PR #372 中六个浏览器旧链接候选文件形成的判断链。此前 `browser-links-routing` 与 `root-ops-link-index-boundary` 已写入总规则，本文件补每个候选组为什么仍停留在 `资源/链接/`，为什么不直接迁入 Folo、TopHub、领域、项目或语境。

纳入文件：

- `资源/链接/浏览器订阅候选-2026-07-17.md`
- `资源/链接/浏览器主题包候选.md`
- `资源/链接/浏览器项目与学习入口候选.md`
- `资源/链接/浏览器账号与服务入口.md`
- `资源/链接/浏览器待复核与一次性入口.md`
- `资源/链接/浏览器垂直检索候选.md`
- `资源/链接/_index.md`
- `资源/语境/_index.md`

边界说明：

- 这些文件是浏览器导出 cohort 的分组索引，不替代 `资源/链接.csv` 或已存在的 canonical 单条记录。
- 它们只保存候选、门牌、触发条件和后续迁居门槛，不代表已经订阅、认可、验证或沉积。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：浏览器旧链接候选不是迁移完成

```yaml
- source_id: "浏览器旧链接候选组整体"
  current_route: resource_links_candidate_cohorts_not_final_residence
  decision_status: recovered_from_pr
  selected_reason: "旧浏览器导出被拆成订阅候选、主题包候选、项目与学习入口、账号与服务入口、待复核与一次性入口、垂直检索候选六类。"
  rejected_alternatives:
    - "把浏览器书签全部搬成正式资源"
    - "把候选文件替代 canonical 链接记录"
    - "把链接存在视为已经形成使用语境"
    - "把候选组直接写入 Folo / TopHub / 领域 / 项目"
  route_reason: "这些文件共同说明：当前只是在 GitHub 留下真实 URL、清理后的站点地址、候选角色和触发条件。是否进入正式账本、语境、领域或项目，要等使用证据、订阅复查或项目动作。"
  reuse_rule: "旧书签先按功能分候选组；只有实际读过、会用、长期调用、影响项目或通过复查，才向上迁居。"
  evidence_locator: "资源/链接/_index.md; six browser candidate files"
  next_action: "作为浏览器旧链接候选总规则。"
```

---

## 2. 订阅候选：48 条只进 2026-07-17 复查，不进账本

```yaml
- source_id: "浏览器订阅候选 48 条"
  current_route: folo_tophub_review_queue_not_subscribed
  decision_status: recovered_from_pr
  selected_reason: "这些媒体、平台和持续更新入口当前仍是 `media` / `platform` 类型链接，包含促销与行动触发、中文综合与公共新闻、中文财经与商业、中文科技产品与行业、国际新闻财经科技、视觉自然与杂志、财经日历等。"
  rejected_alternatives:
    - "直接写入 Folo 订阅源账本"
    - "直接加入 TopHub"
    - "把首页、付费墙、聚合站、财经日历和促销页自动等同于合格订阅源"
    - "在 2026-07-17 前改 Folo 27 项"
  route_reason: "文件明确状态为链接候选队列，不等于已订阅、已认可或已验证。复查边界是：先判断 TopHub 是否已经承担公共温度；Folo 只考虑独特、稳定、真实会读的来源；2026-07-17 先复查现有 27 个 Folo 来源，再按一进一出判断。"
  reuse_rule: "订阅候选必须经过角色、重复度、打开率、付费墙、未读压力和一进一出复查；链接候选不写入正式订阅账本。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#状态; #复查边界"
  next_action: "2026-07-17 复查。"
```

---

## 3. 促销 / 财经日历 / 综合门户：能触发任务，不等于订阅源

```yaml
- source_id: "浏览器订阅候选中的促销、门户和财经日历"
  current_route: task_trigger_or_public_reference_candidates
  decision_status: recovered_from_pr
  selected_reason: "候选中包含 IT之家喜加一、少数派商城、央视网、凤凰网、新浪、腾讯网、搜狐、网易、澎湃、环球网、华尔街见闻财经日历和东方财富财经日历。"
  rejected_alternatives:
    - "把促销入口加入 Folo"
    - "把综合门户加入公共温度常驻"
    - "把财经日历做成每日未读流"
  route_reason: "促销与行动触发入口只有在明确领取或购买动作时有价值；综合门户多与当前 TopHub 公共温度和报刊 / 媒体节点重复；财经日历是查询工具，不是阅读源。"
  reuse_rule: "促销、门户和日历优先按任务触发；只有形成稳定、独特、低噪声信息增量时才复查订阅。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#促销与行动触发; #中文综合与公共新闻; #财经日历"
  next_action: "候选。"
```

---

## 4. 主题包候选：参考门牌，不是领域知识

```yaml
- source_id: "浏览器主题包候选 37 条"
  current_route: resource_links_reference_gateways_not_domain_knowledge
  decision_status: recovered_from_pr
  selected_reason: "主题包候选按百科与文化、学术研究、安全与隐私、技术标准与数据分组，包含 Internet Archive、SEP、IEP、MedlinePlus、Gutenberg、World History Encyclopedia、ACM、Semantic Scholar、arXiv、CNVD、NVD、OWASP、MDN、国家数据等。"
  rejected_alternatives:
    - "直接进入领域"
    - "直接进入资源/语境"
    - "替代已有 canonical 链接记录"
    - "按主题包分组当成已经消化的知识结构"
  route_reason: "文件明确这些条目仍是外部链接候选，按检索角色分组，但尚未成为领域知识，也没有正式进入资源/链接.csv。同名或同 URL 已有正式条目时，继续复用既有 canonical 记录。"
  reuse_rule: "参考站、百科、标准、数据库和学术工具先留门牌；实际读过形成可复用内容才进萃览，会用才进语境，长期领域调用才进领域。"
  evidence_locator: "资源/链接/浏览器主题包候选.md#浏览器主题包候选; #后续迁居门槛"
  next_action: "留链接层。"
```

---

## 5. 安全 / 学术 / 数据主题包：高价值也不自动沉积

```yaml
- source_id: "安全、学术与数据主题包"
  current_route: high_value_reference_links_waiting_use_context
  decision_status: recovered_from_pr
  selected_reason: "ACM、Scopus、Semantic Scholar、arXiv、CNVD、NVD、OWASP、FIRST、ENISA、OpenSSL、Wireshark、pfSense、MDN、Google 趋势、国家数据等明显有工具和参考价值。"
  rejected_alternatives:
    - "因为权威就进入领域"
    - "因为安全相关就进入科技雷达"
    - "因为可查数据就写入数据与结构"
  route_reason: "权威性不是居留层级。它们当前只保存门牌和候选状态；需要在具体安全、论文、统计、标准、工程或排障任务中被调用，才形成语境或项目引用。"
  reuse_rule: "高价值参考源要等具体使用场景把它拉起来；不要让权威链接本身制造长期维护义务。"
  evidence_locator: "资源/链接/浏览器主题包候选.md#学术研究; #安全与隐私; #技术、标准与数据; #后续迁居门槛"
  next_action: "按任务调用。"
```

---

## 6. 项目与学习入口：从浏览器移出，不重建通用书签目录

```yaml
- source_id: "浏览器项目与学习入口候选 96 条"
  current_route: project_or_course_activation_candidates
  decision_status: recovered_from_pr
  selected_reason: "这些入口包含在线课程与技能提升、开发与编程、学术研究与文献、常用工具集、深度阅读与资讯分析等。"
  rejected_alternatives:
    - "重新建立通用书签目录"
    - "全部沉入领域"
    - "全部沉入项目"
    - "作为浏览器常驻书签继续保留"
  route_reason: "文件明确这些入口不再常驻浏览器，但已在 GitHub 留下真实 URL。进入具体项目、课程或比较任务时，从这里挑选并写入对应项目，而不是重建通用书签目录。"
  reuse_rule: "学习和项目入口由当前项目 / 课程激活；没有项目关系时只保留链接候选。"
  evidence_locator: "资源/链接/浏览器项目与学习入口候选.md#浏览器项目与学习入口候选; #居留说明"
  next_action: "按项目激活。"
```

---

## 7. 开发 / 学术 / AI 工具入口：工具可用不等于当前项目采用

```yaml
- source_id: "项目学习候选中的开发、学术与 AI 工具"
  current_route: on_demand_project_tools_not_current_stack
  decision_status: recovered_from_pr
  selected_reason: "候选中包含 AWS、Azure、Gitee、HelloGitHub、LeetCode、Ollama、Stack Overflow、shadcn/ui、Apple Developer、Context7、Tauri、Vite、AMiner、Elicit、Scite、TXYZ、Open LLM Leaderboard、OpenCompass、Statista 等。"
  rejected_alternatives:
    - "把工具可用性写成当前技术栈"
    - "把论文工具写成默认研究流程"
    - "把 AI 排行榜写入科技雷达"
    - "把课程平台写成当前学习计划"
  route_reason: "这些入口只有在课程、研究、工程采用或比较任务激活时才有上下文。否则只是旧书签留下的可访问门牌。"
  reuse_rule: "工具入口要等任务定义后再升格：采用进项目，稳定方法进语境，长期能力进领域。"
  evidence_locator: "资源/链接/浏览器项目与学习入口候选.md#开发与编程; #学术研究与文献; #深度阅读与资讯分析"
  next_action: "按需。"
```

---

## 8. 账号与服务入口：保存门牌，不保存秘密，不因敏感就擅自排除

```yaml
- source_id: "浏览器账号与服务入口 43 条"
  current_route: account_and_service_entry_signposts
  decision_status: recovered_from_pr
  selected_reason: "这些条目属于 account_entry、service_entry 或仍需区分服务与知识属性的外部入口，包括邮箱、学校、办公、服务后台、网络与账号服务、私人用途入口、Gmail、Slack、印象笔记和石墨文档等。"
  rejected_alternatives:
    - "保存密码、验证码、令牌、Cookie、会话 ID 或临时回跳参数"
    - "因为是账号入口就全部排除出仓库"
    - "因为能登录就进入项目"
    - "把账号入口当知识资源"
  route_reason: "文件明确只保存清理后的站点地址，不保存秘密、会话或一次性授权参数。账号入口和服务入口的默认居留地就是资源/链接；只有某个入口实际绑定具体项目动作时，才在相应成为/项目中引用。"
  reuse_rule: "账号与服务只存门牌，不存秘密；入口是否进入项目取决于项目动作，不取决于账号属性。"
  evidence_locator: "资源/链接/浏览器账号与服务入口.md#浏览器账号与服务入口; #居留说明"
  next_action: "留链接层。"
```

---

## 9. 服务 / 知识属性混合入口：先区分角色，再迁居

```yaml
- source_id: "服务与知识属性待区分入口"
  current_route: needs_role_split_before_residence
  decision_status: recovered_from_pr
  selected_reason: "CTF 在线工具、Cybersecurity Mind Maps、OWASP ZAP、PaywallBuster、号码查询、dCode、百变小樱等同时可能是工具、知识、服务、账号或临时入口。"
  rejected_alternatives:
    - "按当前文件夹直接定性"
    - "全部作为知识资源"
    - "全部作为账号服务"
    - "因混合属性直接删除"
  route_reason: "文件明确这些入口仍需区分服务与知识属性。不同角色决定不同居留：工具可能进语境或项目，知识资源可能进萃览或领域，账号服务只保门牌。"
  reuse_rule: "混合入口先补使用场景和角色，再决定居留。"
  evidence_locator: "资源/链接/浏览器账号与服务入口.md#需区分服务入口与知识资源"
  next_action: "待角色复核。"
```

---

## 10. 待复核与一次性入口：不能把候选状态冒充迁移结论

```yaml
- source_id: "浏览器待复核与一次性入口 71 条"
  current_route: link_only_low_value_needs_review_or_one_off_candidates
  decision_status: recovered_from_pr
  selected_reason: "这些网站仍是链接候选、低价值候选、一次性事项或需要补证据的条目。"
  rejected_alternatives:
    - "迁入资源/雷达"
    - "自动迁入领域或项目"
    - "因为可访问就认为有价值"
    - "因为旧书签存在就保留为正式资源"
  route_reason: "文件明确它们不是资源/雷达对象，也没有被自动迁入领域或项目。复核后可以继续留在链接层、进入语境、沉入项目，或只保留历史；当前不能把候选状态冒充成迁移结论。"
  reuse_rule: "待复核入口要么补证据，要么保留历史，要么退相干；不能因存在 URL 就升级。"
  evidence_locator: "资源/链接/浏览器待复核与一次性入口.md#浏览器待复核与一次性入口; #居留说明"
  next_action: "待复核。"
```

---

## 11. 低独特性、一次性、角色未稳：分别处理

```yaml
- source_id: "低独特性 / 一次性 / 角色未稳入口"
  current_route: differentiated_review_outcomes
  decision_status: recovered_from_pr
  selected_reason: "文件把低独特性、已有替代或不进入正式资源表的入口，一次性阅读或处理事项，以及角色尚未稳定的入口分开。"
  rejected_alternatives:
    - "所有待复核都走同一处理"
    - "把一次性事项写入长期资源"
    - "把 AI 检测器、词典、百科、文库和日历工具全部纳入领域"
  route_reason: "低独特性入口往往已有更优替代；一次性事项处理完即可散；角色未稳入口需要先确认它是工具、内容、服务、资料库、检索入口还是项目材料。"
  reuse_rule: "复核时先分低价值、一次性、角色未稳，再决定保留、迁居、退相干或项目引用。"
  evidence_locator: "资源/链接/浏览器待复核与一次性入口.md#低独特性; #一次性阅读或处理事项; #角色尚未稳定"
  next_action: "分类型复核。"
```

---

## 12. 垂直检索候选：播客检索工具，未形成语境

```yaml
- source_id: "浏览器垂直检索候选 2 条"
  current_route: podcast_search_tool_candidates
  decision_status: recovered_from_pr
  selected_reason: "问问小宇宙是中文播客内容检索候选，Dexa 是英文播客与专家内容检索候选。"
  rejected_alternatives:
    - "直接沉入资源/语境"
    - "直接进入播客工作流"
    - "把鸠摩搜索重复计入垂直检索"
  route_reason: "文件明确这些入口仍是检索工具候选，默认居留在资源/链接；只有形成稳定使用语境后，才考虑沉入资源/语境；影响具体研究或项目动作后，再在相应项目中引用。鸠摩搜索仍属于后续复核，不在这里重复计数。"
  reuse_rule: "检索工具先证明稳定使用场景；一次查询不等于语境成立。"
  evidence_locator: "资源/链接/浏览器垂直检索候选.md#浏览器垂直检索候选"
  next_action: "待使用语境。"
```

---

## 13. 本补充 ledger 的复用规则

1. 浏览器旧链接候选是 cohort 分组，不是迁移完成。
2. 订阅候选 48 条不等于 Folo / TopHub 已订阅；2026-07-17 才按一进一出复查。
3. 首页、付费墙、聚合站、财经日历和促销页不自动等同合格订阅源。
4. 主题包候选是参考门牌，不是领域知识。
5. 权威、高价值、安全、学术或数据来源也要等使用场景，不因权威直接沉积。
6. 项目与学习入口由具体项目、课程、研究或比较任务激活；不重建通用书签目录。
7. 账号与服务入口只保存清理后的站点地址，不保存秘密、会话或一次性授权参数。
8. 服务 / 知识属性混合入口先区分角色，再决定居留。
9. 待复核与一次性入口不能把候选状态冒充迁移结论。
10. 低独特性、一次性、角色未稳要分别处理。
11. 垂直检索工具要形成稳定使用语境，才迁入资源/语境或项目。
12. 链接晋升路径仍是：链接 → 萃览 / 语境 / 领域 / 项目 / docs/ops，必须有使用证据。

---

## 14. 仍需继续追回

浏览器旧链接候选组已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内可继续做一轮“已存在 ledger 文件覆盖 PR 原始文件”的反向映射，找出完全未覆盖原始文件；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
