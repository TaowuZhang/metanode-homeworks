# PR #372 判断链追回：根入口、ops 索引与链接治理边界（2026-07-06）

## 0. 边界

本文件追回 PR #372 中仓库根入口、`docs/ops/` 操作索引与 `资源/链接/` 治理索引形成的判断链。

纳入文件：

- `README.md`
- `docs/ops/README.md`
- `资源/链接/_index.md`
- `资源/README.md`
- `资源/语境/_index.md`
- `docs/ops/browser-bookmark-cleanup-2026-07-05.md`
- `docs/ops/browser-link-governance-round-2-2026-07-05.md`
- `docs/ops/browser-link-placement-audit-2026-07-05.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 根入口总规则：世界来的东西先辨身份，再入沃壤

```yaml
- source_id: "README.md 根入口"
  current_route: cold_start_and_guard_entry_not_full_map
  decision_status: recovered_from_pr
  selected_reason: "根 README 定义沃壤不是单纯资料库，也不是提醒系统，而是让世界来的东西入土、腐熟、生根，长成用户能再用、再吃、再种的东西。"
  rejected_alternatives:
    - "把根目录当完整地图"
    - "把仓库当冷柜式资料库"
    - "把 ChatGPT、Codex、Claude、NotebookLM、Notion 等界面当权威事实源"
    - "把所有东西直接塞入资源或 docs/ops"
  route_reason: "根 README 明确：GitHub / worang 是事实底盘、版本化地层、工程接口和长期生长处，也是系统关键内容的唯一母本；外部界面可以承担对话、施工、阅读、展示或协作，但不另存一套权威事实。"
  reuse_rule: "任何新材料先辨身份：成为、领域、资源、技能、wo、docs/ops、src 或外部界面。界面可以更换，事实必须留根。"
  evidence_locator: "README.md#东西禾集; #沃壤与外部界面; #协作与守门入口"
  next_action: "作为仓库冷启动边界。"
```

---

## 2. 根目录不是完整地图，只保留守门入口

```yaml
- source_id: "根目录入口清单"
  current_route: minimal_guard_and_cold_start_entry
  decision_status: recovered_from_pr
  selected_reason: "根目录只保留冷启动和守门所需入口：破土、技能、AGENTS、CLAUDE、worang-index/current、root-manifest 和 root-guard。"
  rejected_alternatives:
    - "在根目录复刻全部目录"
    - "让根目录承担全仓目录导航"
    - "把旧 docs/ideas 重新扩成想法库"
  route_reason: "README 明确根目录不是完整地图；docs/ideas 已退场，只作为历史草案来源被只读辨认，不再是长期入口或新草案池。"
  reuse_rule: "根目录只放冷启动、守门和必要入口；内容应回到其身份层。旧草案不得批量搬回。"
  evidence_locator: "README.md#入口; #协作与守门入口"
  next_action: "保持根入口克制。"
```

---

## 3. 七个主体入口：不是平行垃圾箱

```yaml
- source_id: "成为 / 领域 / 资源 / 技能 / wo / docs/ops / src"
  current_route: identity_based_repository_layers
  decision_status: recovered_from_pr
  selected_reason: "README 将仓库入口按身份分层：成为保存运行态与成为地层；领域保存能力域、作品域与责任区；资源保存世界流入的材料、参考、补给和旧文明残片；技能保存正式工艺；wo 是动作接口；docs/ops 保存发布、展示、外排、workflow、脚本和操作证据；src 是 wo 实现。"
  rejected_alternatives:
    - "把所有材料按文件类型分类"
    - "把设施文件迁入成为、领域或资源主体正文"
    - "让 WQB、GetNote、worang-index 进入主体正文"
  route_reason: "README 明确 docs/ops、runs、workers、scripts、.claude、wo 是设施层；WQB、GetNote、worang-index 属于工程 / 运行 / 索引设施，留在设施层或既有镜像位置，不迁入成为、领域、资源主体正文。"
  reuse_rule: "判断文件归属时看它回答什么：我是谁 / 怎么长 / 会什么 / 取养什么 / 怎样与 AI 协作 / 机器怎么跑 / 操作怎么验。"
  evidence_locator: "README.md#入口"
  next_action: "作为根层路由规则。"
```

---

## 4. docs/ops：操作证据层，不保存长期判断正文

```yaml
- source_id: "docs/ops/README.md"
  current_route: machine_operation_audit_distribution_connectivity_layer
  decision_status: recovered_from_pr
  selected_reason: "docs/ops 承接机器运行层、审计证据层、发布 / 展示 / 分发操作层、连接性与对账层。"
  rejected_alternatives:
    - "把 docs/ops 当内容库"
    - "把 docs/ops 当长期判断库"
    - "把 Notion 公开窗口当协议、判断或技能母本"
    - "把逐 PR 质量打分和个人判断正文放入 docs/ops"
  route_reason: "docs/ops 只回答机器怎么跑、入口怎么验、发布怎么走、失败怎么记、workflow 如何交接，以及这些操作如何不反向污染成为、领域、资源、技能、wo 与公开展示面。长期判断、触发语、关系规则和 AI 进场方式应按内容回破土、成为、领域、资源、技能或 wo。"
  reuse_rule: "凡主要回答长期判断、关系、能力、身体、项目内容或资源沉积的文件，不应放在 docs/ops；凡回答运行、审计、发布、连接、对账和事故的文件，才进 docs/ops。"
  evidence_locator: "docs/ops/README.md#docs/ops · 机器运行、审计与外显操作入口; #子层边界; #不做什么"
  next_action: "作为 ops 层边界。"
```

---

## 5. docs/ops 第一跳：索引是路标，不是全仓复刻

```yaml
- source_id: "docs/ops 第一跳表"
  current_route: operation_index_not_full_catalog
  decision_status: recovered_from_pr
  selected_reason: "docs/ops README 提供冷启动索引、根目录机器账、运行底线、PR 注意力、workflow 证据、浏览器治理、分发、NLM、Notion、Web、Telegram、GetNote 等操作入口。"
  rejected_alternatives:
    - "把第一跳表扩成全仓目录复刻"
    - "把索引维护变成内容目录正文"
    - "把浏览器治理、发布管线、GetNote 接线等操作文件迁出 docs/ops"
  route_reason: "ops 的索引维护负责 root manifest、worang-index 与冷启动指针，不负责全仓目录复刻。浏览器治理、分发、工作流和连接性文件都是操作证据，因此留在 docs/ops。"
  reuse_rule: "ops 索引回答‘处理这个操作问题先读哪一页’，不回答内容层判断。"
  evidence_locator: "docs/ops/README.md#先读哪一页; #子层边界"
  next_action: "保持第一跳表。"
```

---

## 6. Notion 公开窗口：窗不是母本

```yaml
- source_id: "Notion 公开窗口关系"
  current_route: public_window_not_source_of_truth
  decision_status: recovered_from_pr
  selected_reason: "Notion Education Plus 只保留薄公开窗口：个人介绍、精选作品、公开链接和轻量交互陈列。"
  rejected_alternatives:
    - "把 Notion 做成第二套知识库"
    - "把 Notion 内部施工图复刻进 docs/ops"
    - "让 Notion 成为协议、判断或技能母本"
  route_reason: "docs/ops README 明确：文字、结构和 HTML 源码先在 GitHub 留母本，再手工陈列到 Notion。一句话是 Notion 是窗，GitHub 是母本，docs/ops 负责修路、验路、过路和记事故。"
  reuse_rule: "公开展示可以在 Notion，母本必须在 GitHub；外部展示不能反向覆盖仓库判断。"
  evidence_locator: "docs/ops/README.md#与 Notion 公开窗口的关系"
  next_action: "保持 GitHub 母本。"
```

---

## 7. 资源/链接：门牌与候选池，不是迁移命令

```yaml
- source_id: "资源/链接/_index.md"
  current_route: external_link_signposts_and_candidate_pool
  decision_status: recovered_from_pr
  selected_reason: "资源/链接 是外部链接的门牌与候选池，串联链接治理文件与当前候选入口。"
  rejected_alternatives:
    - "把链接候选当已形成语境"
    - "把候选文件替代资源/链接.csv 或正式单条记录"
    - "把浏览器导出 cohort 当 canonical 记录"
    - "把居留建议当迁移命令"
  route_reason: "索引明确：链接进入语境、领域或项目仍需相应使用证据。浏览器旧链接候选组只是导出 cohort 的分组索引，不替代资源/链接.csv 或既有正式条目；同名或同 URL 已有正式条目时，以既有 canonical 记录为主。"
  reuse_rule: "链接层只保门牌、候选、风险、重复和居留建议；迁移、沉入语境、领域或项目必须另有使用证据。"
  evidence_locator: "资源/链接/_index.md#资源/链接 治理索引 v0; #浏览器旧链接候选组"
  next_action: "作为链接层边界。"
```

---

## 8. 链接治理文件组：规则先于动作

```yaml
- source_id: "链接治理文件组"
  current_route: link_governance_rules_before_operations
  decision_status: recovered_from_pr
  selected_reason: "_rules、_field_mapping、_interfaces、_alive_check_evidence、_pr_checklist 和事件样本共同定义链接活性检查、字段映射、接口草案、证据边界、PR 自检和候选事件。"
  rejected_alternatives:
    - "直接改链接条目"
    - "不读规则就跑脚本"
    - "把旧 Markdown 类型字段直接当最终分类"
    - "把失效链接字段直接判 dead"
  route_reason: "索引给出推荐读取顺序：先读规则，再读字段映射、接口、活性证据、候选样本、重复样本和 PR 检查清单。扩样结论也明确旧类型字段漂移明显，只能作为候选信号；失效链接字段缺少检查证据，不能直接判 dead。"
  reuse_rule: "链接治理任务先读规则，生成候选事件和证据，再开 PR 人工审阅；不能直接自动改主库。"
  evidence_locator: "资源/链接/_index.md#治理文件组; #推荐读取顺序; #当前扩样结论"
  next_action: "作为链接治理流程。"
```

---

## 9. 浏览器旧链接候选组：六类候选，不替代正式资源

```yaml
- source_id: "浏览器旧链接候选组"
  current_route: browser_export_cohort_candidates
  decision_status: recovered_from_pr
  selected_reason: "浏览器导出整理形成六类候选：主题包候选 37、订阅候选 48、项目与学习入口 96、账号与服务入口 43、待复核与一次性入口 71、垂直检索候选 2。"
  rejected_alternatives:
    - "把这些候选直接写入资源/雷达"
    - "把订阅候选直接加入 Folo"
    - "把项目入口直接沉入成为/项目"
    - "把账号入口当知识资源"
  route_reason: "索引明确这些文件仍属于链接入口与候选，不放在资源/雷达；已经形成稳定使用语境的浏览器分工和工具架放在资源/语境；迁移、清理和对账证据放在 docs/ops。"
  reuse_rule: "浏览器导出 cohort 要先分候选类型，再按真实使用证据进入语境、雷达、项目、账号服务或历史。"
  evidence_locator: "资源/链接/_index.md#浏览器旧链接候选组"
  next_action: "作为候选索引。"
```

---

## 10. 链接施工边界：允许候选，不允许自动处置

```yaml
- source_id: "链接施工边界"
  current_route: candidate_generation_and_manual_review_only
  decision_status: recovered_from_pr
  selected_reason: "链接治理允许读取旧链接条目、提取最小字段、生成候选事件、标记分类和居留建议、记录风险/重复/低价值/失效候选，并开 PR 供人工审阅。"
  rejected_alternatives:
    - "自动删除链接"
    - "自动合并重复链接"
    - "全量抓网页正文"
    - "把 source metadata 当实时证据"
    - "把可访问当有价值"
    - "把居留权建议当迁移命令"
    - "把近重复当可合并"
    - "把同领域邻居放进去重流程"
  route_reason: "索引明确禁止自动处置。本次浏览器文件迁居是用户明确要求后的人工修正，不把该动作改写为自动迁移权限。"
  reuse_rule: "链接治理只能生成候选和 PR；删除、合并、迁居和价值判断必须人工确认。"
  evidence_locator: "资源/链接/_index.md#施工边界"
  next_action: "继续人工审阅。"
```

---

## 11. v0 总判断：链接、萃览、语境、领域、项目、ops 分层

```yaml
- source_id: "资源/链接 v0 总判断"
  current_route: layered_resource_promotion_path
  decision_status: recovered_from_pr
  selected_reason: "索引给出资源层晋升路径：资源/链接是门牌和候选池；资源/萃览是读过后压出的可复用内容；资源/语境是知道什么时候怎么用的资源；领域是稳定知识结构；成为/项目是已经影响具体行动的链接证据；docs/ops 是迁移、清理、对账与运行证据。"
  rejected_alternatives:
    - "把所有外部链接都留在链接层"
    - "把只有链接的候选直接沉领域"
    - "把操作对账证据放进资源正文"
    - "把读过内容和链接门牌混在一起"
  route_reason: "链接层是入口，不是消化层。只有经过阅读、使用、领域化或项目动作后，才进入上层。"
  reuse_rule: "链接晋升需要证据：读过变萃览，会用变语境，长期调用变领域，影响动作变项目，操作证据进 docs/ops。"
  evidence_locator: "资源/链接/_index.md#v0 总判断; 资源/语境/_index.md#居留边界"
  next_action: "作为链接晋升规则。"
```

---

## 12. docs/ops、资源/链接、资源/语境三者关系

```yaml
- source_id: "浏览器治理三层关系"
  current_route: ops_evidence_links_candidates_context_usage
  decision_status: recovered_from_pr
  selected_reason: "浏览器旧链接清理产生三类材料：迁移、清理和对账证据；链接入口与候选；已经形成稳定使用方法的浏览器分工和工具架。"
  rejected_alternatives:
    - "把浏览器清理过程写入资源/语境"
    - "把浏览器工具架留在 docs/ops"
    - "把候选链接写入语境"
  route_reason: "索引明确：迁移、清理和对账证据位于 docs/ops；浏览器旧链接候选位于资源/链接；已经形成稳定使用语境的浏览器分工和工具架位于资源/语境。"
  reuse_rule: "同一个浏览器工作流的不同产物按证据、候选、语境三层拆开存放。"
  evidence_locator: "资源/链接/_index.md#浏览器旧链接候选组; 资源/语境/_index.md#浏览器与链接"
  next_action: "保持三层。"
```

---

## 13. 本补充 ledger 的复用规则

1. 根 README 是冷启动和守门入口，不是全仓目录复刻。
2. GitHub 是母本；外部界面是工位或窗口，不是另一套权威事实。
3. 根目录入口按身份分层：成为、领域、资源、技能、wo、docs/ops、src。
4. docs/ops 是机器运行、审计、发布、连接和对账层，不保存长期判断正文。
5. Notion 是窗，GitHub 是母本，docs/ops 负责修路、验路、过路和记事故。
6. 资源/链接是门牌与候选池，不是迁移命令。
7. 链接治理先读规则，后做候选事件；不能直接自动改主库。
8. 浏览器导出 cohort 不替代资源/链接.csv 或正式条目。
9. 链接治理允许生成候选和 PR，禁止自动删除、自动合并、全量抓正文和把可访问当有价值。
10. 链接晋升路径是：链接 → 萃览 / 语境 / 领域 / 项目 / docs/ops，必须有使用证据。
11. 同一浏览器治理工作流的产物要拆成 ops 证据、链接候选和语境材料。

---

## 14. 仍需继续追回

根入口、ops 索引与链接治理边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如果继续补仓库内残余，可检查是否还剩 `docs/ops/resource-update-*`、`资源/_registry/*`、`资源更新宪法` 这一组资源更新闸文件需要单独追回。
