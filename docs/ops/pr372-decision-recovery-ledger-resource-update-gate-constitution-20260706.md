# PR #372 判断链追回：资源更新宪法与候选确认闸（2026-07-06）

## 0. 边界

本文件追回 PR #372 中资源更新宪法、资源更新候选事件、确认闸、资源登记与自动化接线形成的判断链。

纳入文件：

- `资源/资源更新宪法 v0.md`
- `docs/ops/resource-update-candidate-gate-v0.md`
- `资源/README.md`
- `资源/链接/_index.md`
- `资源/语境/_index.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：资源更新不是外部平台同步

```yaml
- source_id: "资源更新宪法整体"
  current_route: resource_update_constitution_not_sync_pipeline
  decision_status: recovered_from_pr
  selected_reason: "资源更新指外部平台、本地材料、Notion 借港、阅读 / 听看玩过程、评价、回声、沉积位置或可用性发生可记录变化时，沃壤用候选事件或受控写入更新资源地层。"
  rejected_alternatives:
    - "把资源更新做成外部平台同步"
    - "自动打卡"
    - "自动消费记录"
    - "噪声归档"
    - "把外部平台状态直接当沃壤判断"
  route_reason: "宪法把资源更新分成事实层、接触层、沉积层。外部状态只能作为证据，候选事件只是入口，沉积必须能回访、复用或承重。"
  reuse_rule: "以后任何平台或工具变化，先问它属于事实、接触还是沉积。事实可被观察，接触可被候选，沉积必须经过规则或确认。"
  evidence_locator: "资源/资源更新宪法 v0.md#0-总判准; 资源/README.md#资源更新宪法"
  next_action: "作为资源更新总规则。"
```

---

## 2. GitHub 主账：外部发生，GitHub 记账，谷确认判断

```yaml
- source_id: "GitHub 主账本规则"
  current_route: github_as_resource_update_ledger
  decision_status: recovered_from_pr
  selected_reason: "GitHub 可版本化、可回滚、可审计、可批处理、可跨工具、可长期保存、可分支试验。"
  rejected_alternatives:
    - "把主控权放外部平台"
    - "把主控权放本地临时目录"
    - "外部拿不到数据就放弃主动权"
    - "用外部收藏、播放、观看、下载状态当长期事实账本"
  route_reason: "外部网站常见问题包括 API 权限有限、页面结构变化、状态不可可靠导出、平台下架限流封禁、AI 摘要或转录不能稳定回取。因此简式原则是：外部能拿到什么，算证据；GitHub保存什么，才算沃壤可复核事实；谷确认什么，才算沃壤判断。"
  reuse_rule: "外部平台只作为来源能力声明后的证据输入；值得长期保存、复核、批处理、回滚、协作或跨工具复用时，优先进入 GitHub 管理的文本、清单、manifest、报告或候选事件。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#0-总判准; #2-为什么-GitHub-优先"
  next_action: "继续 GitHub 主账。"
```

---

## 3. 四账分离：镜像账、候选账、主库账、沉积账

```yaml
- source_id: "资源更新四账分离"
  current_route: mirror_candidate_master_sediment_accounts
  decision_status: recovered_from_pr
  selected_reason: "资源更新不应直接从外部状态跳到主库，中间至少分四种账：镜像账、候选账、主库账、沉积账。"
  rejected_alternatives:
    - "外部镜像直接改主库"
    - "候选事件直接代表谷判断"
    - "主库身份自动生成"
    - "沉积账由工具自动决定"
  route_reason: "镜像账保存外部平台 / 本地导出的原始事实，不代表判断；候选账记录某事实可能意味着一次更新，不代表判断；主库账是沃壤承认的资源身份、状态、位置，不应无确认自动生成；沉积账进入领域、项目、source pack 或萃览，必须有规则或确认。"
  reuse_rule: "每条资源变化必须先落到正确账层。能自动生成的是镜像和候选，不是主库身份和沉积判断。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#3-四账分离"
  next_action: "作为资源更新分账规则。"
```

---

## 4. 候选事件三分法：事实、关系、沉积候选

```yaml
- source_id: "候选事件三分法"
  current_route: fact_relation_sediment_candidate_split
  decision_status: recovered_from_pr
  selected_reason: "所有外部来源先归入三类候选之一：fact_candidate 记录发生了什么；relation_candidate 记录它和谁有关；sediment_candidate 记录它可能该长到哪里。"
  rejected_alternatives:
    - "新增事实直接沉积"
    - "多入口关系直接合并主身份"
    - "候选去处直接变迁移命令"
  route_reason: "fact_candidate 不可直接推出喜欢、重要、沉积或项目归属；relation_candidate 不可直接完成主身份合并、删除重复项、改主标题或路径；sediment_candidate 必须回答为什么不是噪声、未来怎么找、服务哪个领域 / 项目 / 接法、是否需要谷确认。"
  reuse_rule: "先分候选类型，再决定下一步。事实候选观察，关系候选复核，沉积候选过确认闸。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#4-候选事件三分法"
  next_action: "作为候选事件分类规则。"
```

---

## 5. 候选事件格式：它是最小单位，不是主库记录

```yaml
- source_id: "资源更新候选事件格式"
  current_route: candidate_event_minimum_unit
  decision_status: recovered_from_pr
  selected_reason: "候选事件需要记录 event_id、时间、候选类型、来源、来源能力、来源限制、外部状态、观察到的变化、证据、候选沃壤状态、建议去处、是否需谷确认、风险标记和下一步。"
  rejected_alternatives:
    - "只记一句‘有变化’"
    - "只保存外部 URL"
    - "只保存工具输出"
    - "没有来源能力和限制就写主库"
  route_reason: "候选事件是资源更新的最小单位。它不是主库记录，也不是最终判断。格式中必须包含 source_capability、source_limit、requires_gu_confirmation 和 risk_flags，用来防止工具假装全知或自动判断。"
  reuse_rule: "候选事件必须说明来源可靠提供什么、不能提供什么、证据是什么、为什么能或不能直接执行。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#5-候选事件最小格式; 资源/资源更新宪法 v0.md#5-候选事件格式"
  next_action: "作为候选格式。"
```

---

## 6. 来源能力声明：先声明能力，不假装全知

```yaml
- source_id: "来源能力声明"
  current_route: source_capability_declaration_before_integration
  decision_status: recovered_from_pr
  selected_reason: "每个外部来源接入前，都必须先声明能力，包括 can_observe、can_infer_only_weakly、cannot_observe、user_control_level、preferred_worang_landing、writeback_policy、confirmation_policy 和 failure_mode。"
  rejected_alternatives:
    - "工具拿到什么就当真相"
    - "API 能取到字段就写主库"
    - "播放、收藏、source 加入、摘要生成直接推断价值"
    - "不同来源使用同一权重"
  route_reason: "默认判断表将 GitHub 设为高掌控主账本；本地是临时处理场；Notion 借港是工作台 / 离港点；豆瓣、Apple Music、Steam、Eagle、BibiGPT、GetNote、NotebookLM、链接库分别只是外部事实源、弱事实源、资产索引、取水 / 转录源、摘录事实源、source pack / grain probe 或入口 / 活性源。"
  reuse_rule: "新接入外部工具之前，必须写清它可靠观测什么、只能弱推断什么、完全不能观测什么，以及失败时怎么办。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#7-来源能力声明"
  next_action: "作为外部来源接入模板。"
```

---

## 7. 确认闸：候选可自动生成，但不能自动沉积

```yaml
- source_id: "资源更新确认闸"
  current_route: confirmation_gate_before_master_identity_judgment_merge_delete_sediment
  decision_status: recovered_from_pr
  selected_reason: "新外部入口、新播放 / 观看 / 游玩 / 收藏记录、新摘要 / transcript / source pack、链接可访问性变化、疑似重复、manifest 新增行、外部标签 / 文件夹变化都可自动生成候选。"
  rejected_alternatives:
    - "自动写入主库身份"
    - "自动改评价"
    - "自动合并主身份"
    - "自动删除、归档、退相干"
    - "自动从资源升到领域或项目"
    - "把外部 AI 摘要视为沃壤判断"
    - "把外部消费状态视为谷已经消化"
  route_reason: "必须谷确认的包括写入或改动主库身份、改评价、合并主身份、删除归档退相干、沉积升级、外部 AI 摘要成为沃壤判断、外部消费状态成为谷已消化、高噪声批量导入。"
  reuse_rule: "自动化只可生成候选、manifest、报告、索引或观察记录；触及主库、评价、合并、删除、沉积升级、项目 / 领域归属时必须确认。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#8-确认闸; 资源/资源更新宪法 v0.md#6-需要谷确认的场景"
  next_action: "继续执行确认闸。"
```

---

## 8. 低风险径行：只增加候选，不代表评价

```yaml
- source_id: "低风险径行条件"
  current_route: reversible_candidate_only_changes
  decision_status: recovered_from_pr
  selected_reason: "某些操作可低风险径行，但必须同时满足：不改主库判断、不删除不合并不覆盖、不代表谷评价、可回滚、只增加候选 / manifest / 报告 / 索引 / 观察记录、发生在 branch / PR 或明确允许的候选层。"
  rejected_alternatives:
    - "在主干静默改主库"
    - "用低风险名义完成沉积升级"
    - "把候选层写入当成已确认"
  route_reason: "低风险径行的意义是允许观察和试压，不是绕过确认闸。"
  reuse_rule: "可自动或半自动执行的只限候选层增量；进入主库和承重层仍需确认。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#8-确认闸"
  next_action: "作为自动化边界。"
```

---

## 9. PR / branch 工作法：试压，不污染 main

```yaml
- source_id: "资源更新 PR / branch 工作法"
  current_route: branch_pr_candidate_review_flow
  decision_status: recovered_from_pr
  selected_reason: "资源更新机制优先通过 GitHub branch / PR 试压。推荐流程是观察外部变化 → 生成候选事件 / manifest / 报告 → 新建 branch → 写入候选层 → PR 审阅 → 谷确认是否沉积 → 决定 merge / 修改 / 退相干。"
  rejected_alternatives:
    - "外部变化直接推 main"
    - "工具直接合并资源更新"
    - "PR body 不说明是否只做候选"
  route_reason: "PR body 必须说明来源、是否只做候选、是否触及主库、是否需要谷确认、不做什么、后续去处。这样候选机制可以先在 branch / PR 中试压，不污染 main。"
  reuse_rule: "资源更新 PR 不能只写‘同步完成’，必须写明候选属性、确认需求和不做事项。"
  evidence_locator: "docs/ops/resource-update-candidate-gate-v0.md#9-PR--branch-工作法"
  next_action: "继续 branch / PR 试压。"
```

---

## 10. 外部状态映射：平台行为只说明发生，不说明价值

```yaml
- source_id: "外部状态到沃壤状态映射"
  current_route: external_state_as_evidence_not_worang_state
  decision_status: recovered_from_pr
  selected_reason: "外部平台状态只能回答平台记录了什么、资源入口在哪里、是否发生过某种接触迹象、是否有可提取材料、是否有时间、标签、评论、摘要、播放、收藏、文件等证据。"
  rejected_alternatives:
    - "外部显示看过就写已消化"
    - "播放次数或高频播放直接写重要"
    - "平台收藏直接写入库"
    - "AI 摘要直接写评价"
    - "外部下架自动删除沃壤记录"
  route_reason: "外部平台状态不能直接回答谷是否真正消化、是否值得沉积、是否进入资源 / 领域 / 项目、是否代表谷评价、是否应长期保存。冲突时，谷当前确认优先，沃壤沉积状态优先于外部平台状态，新证据优先触发复核，不自动覆盖旧判断。"
  reuse_rule: "外部事实可以触发复核，不能覆盖沃壤判断。平台行为只说明发生，不说明价值成立。"
  evidence_locator: "资源/资源更新宪法 v0.md#4-外部状态--沃壤状态映射原则; docs/ops/resource-update-candidate-gate-v0.md#11-反依赖外部平台原则"
  next_action: "作为外部状态映射规则。"
```

---

## 11. 各来源映射：每个平台只给对应候选

```yaml
- source_id: "豆瓣 / Apple Music / Steam / Eagle / BibiGPT / GetNote / NotebookLM / 链接库 / 手动输入"
  current_route: source_specific_candidate_mapping
  decision_status: recovered_from_pr
  selected_reason: "宪法把不同来源映射为不同候选：豆瓣想看/想读是轻触，评分/短评是评价候选；Apple Music 播放是经过，资料库是收藏，歌单是情境簇，高频播放是强接触但不自动沉积；Steam 愿望单、已拥有、游玩时长、成就各自只是候选；BibiGPT 摘要和 transcript 是可处理材料候选；GetNote 摘录是强接触候选；NotebookLM source 是 source pack 候选；链接 URL 是入口候选；谷手动输入是最高权重。"
  rejected_alternatives:
    - "用同一规则处理所有平台"
    - "把播放、愿望单、收藏、下载、摘要、source 加入都当沉积"
    - "把自动摘要当谷评价"
  route_reason: "来源能力不同，用户掌控度不同，默认权重不同，所以必须按来源映射。"
  reuse_rule: "每个平台只给它可靠支持的候选类型。真正评价、主身份、沉积和项目归属仍需确认。"
  evidence_locator: "资源/资源更新宪法 v0.md#4-3-来源映射细则; docs/ops/resource-update-candidate-gate-v0.md#7-来源能力声明"
  next_action: "作为平台映射规则。"
```

---

## 12. 状态词区分：听过、看过、收藏、下载、入库、回声、沉积

```yaml
- source_id: "资源状态词区分"
  current_route: terminology_guardrail
  decision_status: recovered_from_pr
  selected_reason: "宪法区分听过、看过、玩过、收藏、下载、入库、有回声、沉积。"
  rejected_alternatives:
    - "听过等于喜欢"
    - "看过等于理解"
    - "收藏等于已接触"
    - "下载等于已整理"
    - "入库等于已评价"
    - "有回声等于必然长期保存"
    - "沉积等于外部收藏或平台状态"
  route_reason: "状态词混淆会把外部消费和平台行为误写为沃壤判断。"
  reuse_rule: "每次写资源状态都使用严格词义；没有证据不要升级词。"
  evidence_locator: "资源/资源更新宪法 v0.md#7-状态词区分"
  next_action: "作为术语护栏。"
```

---

## 13. 资源接口：资源/、领域/、成为/项目、借港、NotebookLM、退相干

```yaml
- source_id: "资源更新后续接口"
  current_route: resource_domain_project_borrow_port_nlm_decoherence_routes
  decision_status: recovered_from_pr
  selected_reason: "宪法定义资源进入不同后续接口的条件：资源/ 接收已消化、可批量处理、可复用、有明确来源和稳定主身份、有未来回访价值或与领域 / 项目 / 成为线连接的材料；领域/ 需要说明支撑哪个领域判断；成为/项目需要说明服务哪个项目；Notion 借港适合尚未消化、需要共同治理或可视化操作的材料；NotebookLM source pack 是推理容器；退相干防止沃壤变成噪声仓库。"
  rejected_alternatives:
    - "纯平台收藏进资源"
    - "纯播放记录进资源"
    - "无回声链接进资源"
    - "未判断的大批量素材进资源"
    - "NotebookLM source pack 当主库"
    - "退相干视为失败"
  route_reason: "资源/ 不接收纯平台收藏、纯播放记录、纯下载痕迹、无回声链接、未判断的大批量素材和‘也许以后有用’的噪声。借港材料必须带离港判断点。source pack 是推理容器，不是主库。退相干不是失败，而是防噪声。"
  reuse_rule: "候选要按成熟度和作用落位：未消化留借港，source pack 留推理容器，可回访复用进资源，支撑判断进领域，影响推进进项目，无回声退相干。"
  evidence_locator: "资源/资源更新宪法 v0.md#8-后续支线接口; 资源/README.md#什么时候沉到这里"
  next_action: "作为后续接口规则。"
```

---

## 14. 防噪声原则：反自动打卡、反自动消费记录、反噪声仓库

```yaml
- source_id: "资源更新防噪声原则"
  current_route: anti_checkin_anti_consumption_log_anti_noise_archive
  decision_status: recovered_from_pr
  selected_reason: "宪法禁止把播放过、打开过、下载过、收藏过、摘要过、加入资料库、加入 source pack 自动记为成就。"
  rejected_alternatives:
    - "完整记录谷消费过什么"
    - "把所有未来也许有用的材料入库"
    - "保存没有回声、不能复用、不能归位的资源"
  route_reason: "资源更新不追求完整消费记录。它只保留未来可能回访、复用、判断、项目化、领域化的材料。任何资源若不能回答以后怎么找、为什么要找、服务哪个领域或项目、是否有回声、是否可复用、是否可批量处理、是否保留为证据或是否代表重要接触，就不应进入主库。"
  reuse_rule: "消费可发生即散；资源须能回访；沉积须能承重。"
  evidence_locator: "资源/资源更新宪法 v0.md#9-防噪声原则"
  next_action: "作为防噪声规则。"
```

---

## 15. 资源登记：registry 只登记规则，不自动同步

```yaml
- source_id: "资源/_registry"
  current_route: registry_rules_not_auto_sync_engine
  decision_status: recovered_from_pr
  selected_reason: "资源/_registry 是资源层的总登记与治理入口，只登记规则，不自动同步。它包括 sources.yml、residence-rules、update-events、resource-update-gate。"
  rejected_alternatives:
    - "把 registry 当自动同步引擎"
    - "把 registry 条目当已执行迁移"
    - "绕过资源 README 直接按 registry 改库"
  route_reason: "资源 README 说明：进入资源前先看 README；需要判断材料是否可同步、是否先建索引、是否沉到领域或项目时，再看 registry；需要判断为什么更新资源、内容上值不值得更新、克能否直接拾掐时，先看 resource-update-gate。"
  reuse_rule: "registry 是治理入口，不是自动行动器。它提供规则和登记，执行仍走候选事件、确认闸和 PR。"
  evidence_locator: "资源/README.md#资源总登记"
  next_action: "作为 registry 边界。"
```

---

## 16. 自动化接线：workflow 是捷径边，不是判决

```yaml
- source_id: "资源层自动化接线"
  current_route: workflow_registration_with_warning_not_judgment
  decision_status: recovered_from_pr
  selected_reason: "资源 README 登记了三条 workflow：sync-douban 每周写入资源/_douban，dry-run + 审计 → PR，末刀归谷；sync-getnote 每日写入资源/得到/_getnote，但直接 push 主干，未走候选闸；sync-steam-quarterly 待登记核对。"
  rejected_alternatives:
    - "把 workflow 直推当正确范式"
    - "因直推未出事就取消确认闸"
    - "把 workflow 输出当沃壤判断"
  route_reason: "README 明确：⚠️ 是告警不是判决；直推未带入危险内容，但是否补 dry-run / 审计 / PR 闸留给后续管道 PR 判断。workflow 只是捷径边，服务于动作。"
  reuse_rule: "workflow 接线变化必须同步更新 README 自动化表和 registry automation_links；workflow 输出仍需按候选、manifest、确认闸分层。"
  evidence_locator: "资源/README.md#自动化接线现状登记"
  next_action: "后续管道 PR 处理直推风险。"
```

---

## 17. 本补充 ledger 的复用规则

1. 资源更新不是同步、打卡、消费记录或噪声归档。
2. 外部状态是证据，不是判决；候选事件是入口，不是主库；沉积必须能回访、复用或承重。
3. GitHub 是主账本；本地是临时处理场；外部平台是低到中置信证据源。
4. 四账分离：镜像账、候选账、主库账、沉积账。
5. 候选三分：fact_candidate、relation_candidate、sediment_candidate。
6. 每个来源接入前必须声明能力、限制、写回策略、确认策略和失败模式。
7. 候选可自动生成；主库身份、评价、合并、删除、沉积升级、项目 / 领域归属必须确认。
8. 低风险径行只增加可回滚候选、manifest、报告、索引或观察记录。
9. 资源更新优先 branch / PR 试压，不污染 main。
10. 外部平台行为只说明发生，不说明价值成立。
11. 各平台只给各自可靠支持的候选类型。
12. 状态词不能升级：听过不等于喜欢，看过不等于理解，收藏不等于入库。
13. 资源/、领域/、成为/项目、借港、NotebookLM source pack、退相干各有接口。
14. 资源系统反自动打卡、反自动消费记录、反噪声仓库。
15. registry 只登记规则，不自动同步。
16. workflow 是捷径边，不是判决；接线变化要同步登记。

---

## 18. 仍需继续追回

资源更新宪法与候选确认闸已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如继续补仓库内残余，可检查 `docs/ops/resource-update-*` 是否只剩本文件，或转入旧 ChatGPT 对话追索。
