# PR #372 判断链追回：最近 7 天 ChatGPT 对话证据接入（2026-07-06）

## 0. 边界

本文件接续 `pr372-decision-recovery-ledger-old-chat-partial-recovery-boundary-20260706.md`。

前一份边界账只检查了仓库内是否已有旧 ChatGPT 对话原文，结论是：PR #372 仓库分支中未找到可直接引用的完整旧对话沉积文件。

本轮根据用户补充：此前保留的判断在 ChatGPT 最近 7 天对话里，应主动翻找最近 7 天对话。因此，本文件把“旧对话追索”从“仓库无原文”推进为：

- 最近 7 天 ChatGPT 对话索引中有可确认的任务边界与若干事实；
- 2026-07-06 上传的 `粘贴的文本 (1).txt` 中有可引用的对话遗留问题描述；
- 但这些仍不是完整逐源原文，不能替代逐源判断链本身；
- 能写入的是证据接入与处理协议，不是把所有旧对话 partial 标记完成。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 最近 7 天对话索引能确认什么

```yaml
- source_id: "最近 7 天 ChatGPT 对话索引"
  current_route: recent_chat_index_evidence_not_full_transcript
  decision_status: evidence_intake_recovered
  selected_reason: "用户明确补充：此前保留的判断在 ChatGPT 最近 7 天对话里，应主动翻找最近 7 天对话。对话索引能确认 PR #372 与 TopHub、Folo、浏览器书签、source judgment chains 有关，并能确认用户在 2026-07-06 反复要求追回每个来源背后的判断链。"
  rejected_alternatives:
    - "继续只说仓库里没有旧对话原文"
    - "把最近对话索引当完整原文"
    - "用索引摘要直接生成逐源 why_in / why_out"
    - "把所有 partial 标记 completed"
  route_reason: "最近 7 天对话索引可以证明任务边界和部分事实，但不能提供每个来源完整旧对话上下文、原句、相邻替代来源或当时排序因果。"
  reuse_rule: "最近对话索引用于建立追索方向和状态边界；逐源落账仍需要原文、粘贴文本、已沉积文件或可回查的文件证据。"
  evidence_locator: "recent ChatGPT conversation search 2026-07-06; file_library `粘贴的文本 (1).txt`"
  next_action: "用作旧对话追索入口，不作完成声明。"
```

---

## 2. 用户的核心要求在最近对话中已确认

```yaml
- source_id: "用户核心要求：每个来源背后的判断链"
  current_route: binding_task_requirement
  decision_status: recovered_from_recent_chat
  selected_reason: "用户在 2026-07-06 明确说：所有判断都在 ChatGPT 对话历史里，要追回到 PR 中；具体要求包括为什么选进 TopHub / 今日热榜 / Folo、为什么只选这个来源、为什么这样排序、以后遇到类似来源怎样复用。用户又反复给出固定标准：每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。"
  rejected_alternatives:
    - "只评价 PR 是否可合并"
    - "只总结 PR 改了哪些文件"
    - "生成新的方案、MVP 或流程"
    - "用大类原则替代逐源判断"
    - "把 partial 当完成"
  route_reason: "这个要求本身是当前所有 ledger 的约束，不是可选格式。后续任何补充如果不能回答四个 why，就只能标为索引、边界或候选，不能标为判断链完成。"
  reuse_rule: "后续新增任何 ledger 条目，必须至少能回答 selected_reason、rejected_alternatives、route_reason、reuse_rule；否则标 `needs_source_text` 或 `candidate_only`。"
  evidence_locator: "recent ChatGPT conversation facts 2026-07-06; `粘贴的文本 (1).txt` task statement"
  next_action: "作为所有后续旧对话追索硬约束。"
```

---

## 3. 上传粘贴文本提供了可用任务边界

```yaml
- source_id: "粘贴的文本 (1).txt"
  current_route: uploaded_chat_legacy_problem_statement
  decision_status: recovered_from_file_library_summary
  selected_reason: "该文件明确说明任务不是评价 PR 能不能合，也不是总结 PR 改了哪些文件，而是复述 PR #372 里已经发生过的重要判断：为什么留下、为什么不要、为什么进 TopHub、为什么只进 Folo 复查、为什么按需、为什么退出、为什么是真实平台动作、为什么只是候选不能写进账本。"
  rejected_alternatives:
    - "把任务改成 PR 评价"
    - "发明漂亮说法解释 PR"
    - "只写系统原则"
    - "只列文件类别"
    - "只列结果数量"
  route_reason: "该粘贴文本把协作问题定位得很清楚：缺的是把 ChatGPT 里发生过的判断落档成可核验取舍，而不是更换表达格式或重新设计方案。"
  reuse_rule: "后续旧对话追索的最小产出不是总结，而是取舍还原：source / chosen / rejected / layer / evidence / reuse。"
  evidence_locator: "file_library `粘贴的文本 (1).txt`, created 2026-07-06T13:03:23Z"
  next_action: "作为旧对话追索的任务声明证据。"
```

---

## 4. 最近对话确认的系统事实

```yaml
- source_id: "PR #372 最近对话事实"
  current_route: recent_chat_context_fact_layer
  decision_status: recovered_from_recent_chat
  selected_reason: "最近 7 天对话确认：PR #372 位于 `dongxi-heji/worang`，标题为 `订阅系统 v0：按 TopHub 能力重构 Folo 与沃壤分工`，分支为 `feat/folo-subscription-system-v0-20260703`，保持 Draft / open / unmerged。GitHub 存系统设计和决策记录，不存完整 OPML 或私人订阅。"
  rejected_alternatives:
    - "把 PR 当已合并"
    - "继续尝试 merge Draft PR"
    - "把私人完整订阅关系写进公开仓库"
  route_reason: "这些事实影响后续动作边界：不能 merge、不能改 Ready、不能把候选或私有来源写成公开真实账本。"
  reuse_rule: "任何后续写入只新增证据文件或 ledger；不合并 PR，不改 Draft，不公开私人订阅。"
  evidence_locator: "recent ChatGPT conversation facts 2026-07-03 to 2026-07-06"
  next_action: "保持 PR 不合并。"
```

---

## 5. 最近对话确认的三层分工表达

```yaml
- source_id: "TopHub / Folo / 沃壤 三层分工"
  current_route: recovered_recent_chat_phraseology_already_supported_by_pr_files
  decision_status: recovered_from_recent_chat_and_pr_ledgers
  selected_reason: "最近 7 天对话中曾把 TopHub 描述为 public weather map / 今日热榜，Folo 描述为 long-term contact list，沃壤描述为 judgment layer。这个表述与已经写入的系统 ledger 一致：TopHub 看公共注意力和阶段性雷达，Folo 接长期关系，沃壤保存判断。"
  rejected_alternatives:
    - "把 TopHub 当第二个 Folo"
    - "把 Folo 当公共热榜收件箱"
    - "把沃壤当外部平台镜像"
  route_reason: "该表述可作为最近对话中已确认的系统分工短句，但不能替代具体来源判断。它是上层路由，不是逐源证据。"
  reuse_rule: "遇到来源时先按三层分工粗分，再回具体文件或旧对话原文补 why_in / why_out。"
  evidence_locator: "recent ChatGPT conversation fact 2026-07-03; existing `ledger-tophub-folo-system-routing`"
  next_action: "只作为总路由短句复用。"
```

---

## 6. 最近对话确认的 guided review 三问题

```yaml
- source_id: "TopHub / Folo guided review three questions"
  current_route: recovered_recent_chat_protocol_already_in_pr_file
  decision_status: recovered_from_recent_chat_and_pr_file
  selected_reason: "最近 7 天对话确认：2026-07-04 用户要求记住后续 PR / GitHub 更新的三个问题，并创建了 `docs/ops/tophub-folo-guided-review-20260704.md`。三个问题是：TopHub 是否新增 / 替换 / 删除及放置；是否新增精确追踪器并区分通知机器人；内容该放 TopHub 还是免费 Folo。"
  rejected_alternatives:
    - "只做页面点评"
    - "只问是否值得订阅"
    - "把追踪器和通知机器人混为一谈"
    - "把 TopHub 候选自动写进 Folo"
  route_reason: "这个对话事实已经被仓库文件承接。它是目录复查流程，不是单个来源最终判断。"
  reuse_rule: "后续每个页面 / 小标签先回答这三个问题，再决定是否进入 ledger。"
  evidence_locator: "recent ChatGPT conversation fact 2026-07-04; docs/ops/tophub-folo-guided-review-20260704.md"
  next_action: "已由 `ledger-folo-early-operation-chain` 承接。"
```

---

## 7. 最近对话确认的社区 early facts

```yaml
- source_id: "社区第 1 页与地方门户最近对话事实"
  current_route: recent_chat_facts_already_partly_covered_by_ledgers
  decision_status: recovered_from_recent_chat_partial
  selected_reason: "最近 7 天对话确认：社区 page 1 audit 文件曾记录到 GitHub；页面有 12 个节点；当时建议未来候选包括 `马蜂窝｜热门游记` 和 `水木社区｜十大热门话题`；另有地方门户 closure 文件记录到 GitHub，`地方门户` 49/49 marked fully closed，无新增 TopHub / Folo。"
  rejected_alternatives:
    - "把社区第 1 页完成说成社区 235 节点全量完成"
    - "把马蜂窝 / 水木都写成已执行"
    - "把地方门户继续列为未闭环"
  route_reason: "这些最近对话事实与后续社区 ledger 一致：马蜂窝最终已进入生活与社区；水木仍是未来任务触发或短期试验候选；地方门户已闭环，不应反复打开。"
  reuse_rule: "最近对话事实可以校正待办边界，但仍需以仓库 closure / pending / platform verification 文件判断真实执行状态。"
  evidence_locator: "recent ChatGPT conversation facts 2026-07-05; community and pending-action ledgers"
  next_action: "无需重复审地方门户；社区整体仍 partial。"
```

---

## 8. 最近对话确认的 36氪 early fact

```yaml
- source_id: "36氪科技页 early fact"
  current_route: recent_chat_fact_superseded_by_later_technology_ledgers
  decision_status: recovered_from_recent_chat_superseded
  selected_reason: "最近 7 天对话确认：2026-07-04 针对 36氪科技页，当时判断为 no immediate add； strongest replacement candidate 是 `36氪｜深氪` vs existing `36氪｜24小时热榜`；no tracker / Folo changes yet。"
  rejected_alternatives:
    - "立即添加 36氪 深氪"
    - "立即改追踪器"
    - "立即改 Folo"
  route_reason: "该事实是早期科技页判断的一部分，后来已经被科技目录、科技商业全球化、当前科技雷达等 ledger 吸收。若后续要追 36氪 逐源原句，还需原始对话或对应审计文件行。"
  reuse_rule: "早期对话中的 no immediate add 只能作为时间边界；当前真实状态以科技专项 ledger 和 TopHub 当前顺序为准。"
  evidence_locator: "recent ChatGPT conversation fact 2026-07-04; technology ledger series"
  next_action: "不再把 36氪 early candidate 当当前待办。"
```

---

## 9. 对上传粘贴文本中的可执行流程取其边界，不照单生成新方案

```yaml
- source_id: "粘贴文本里的 decision-ledger / CI 建议"
  current_route: process_suggestions_not_current_execution_order
  decision_status: recovered_as_context_not_adopted_plan
  selected_reason: "粘贴文本里出现过 decision-ledger.yml、why_in、why_out、review_cycle、CI、grep 旧对话、TopHub 80 + Folo 27 等流程建议。"
  rejected_alternatives:
    - "立刻把建议当用户已授权的 MVP / PR-B 方案"
    - "把 80 + 27 直接生成空白行并宣布完成"
    - "用 CI 建议替代当前的手工 ledger 追回"
  route_reason: "用户此前已经明确反感被擅自生成方案。该文本对本任务有用的部分是边界：逐源理由要可验证，真实账本与判断账本要分开，候选 / 复查 / 已执行要分清。不是要求当前模型立刻实施新的 CI 方案。"
  reuse_rule: "从粘贴文本中只提取已被用户认可的约束与问题定义；流程方案必须等用户明确要求才执行。"
  evidence_locator: "file_library `粘贴的文本 (1).txt` snippets about decision-ledger / CI; user corrections in current conversation"
  next_action: "继续手工追回 ledger，不新增 CI。"
```

---

## 10. 对旧对话 partial 的状态修正

```yaml
- source_id: "旧对话 partial 状态"
  current_route: recent_chat_index_available_but_raw_transcript_missing
  decision_status: updated_boundary
  selected_reason: "前一份边界账写的是仓库里找不到旧对话原文；本轮用户指出旧材料在最近 7 天 ChatGPT 对话中。主动搜索后，确实找到了最近对话索引与上传粘贴文本作为任务边界证据，但仍没有完整逐源原文。"
  rejected_alternatives:
    - "继续说完全没有旧对话证据"
    - "宣布旧对话迁移完成"
    - "把最近对话摘要当逐源原文"
    - "把当前模型补写当旧判断"
  route_reason: "状态应从 `no_source_text_found_in_repo` 修正为 `recent_chat_index_available_raw_transcript_needed`。"
  reuse_rule: "后续若继续旧对话追索，先用最近对话索引锁定具体对话和主题，再要原文或已沉积片段；只有原文可见时才逐源落账。"
  evidence_locator: "recent ChatGPT conversation search; this intake ledger"
  next_action: "等待具体旧对话原文 / 导出 / 粘贴，或由用户指定某个最近对话标题继续追。"
```

---

## 11. 本补充 ledger 的复用规则

1. 最近 7 天 ChatGPT 对话索引是有效追索入口，但不是完整原文。
2. 上传的 `粘贴的文本 (1).txt` 可以作为任务边界证据：要复述已发生判断，不评价 PR，不发明系统。
3. 最近对话确认的 PR 元信息、三层分工、guided review 三问题、社区 early facts、36氪 early fact 可以作为边界或时间线证据。
4. 这些证据不能直接生成 80 + 27 逐源 why_in / why_out。
5. 粘贴文本里的 CI / decision-ledger 方案只能作为上下文，不自动执行。
6. 旧对话 partial 状态应更新为：最近对话索引可用，但逐源原文仍需提供。
7. 后续旧对话 ledger 必须按 `old_chat_evidence_recovered`、`old_chat_candidate_only`、`superseded_by_later_ledger`、`superseded_by_platform_verification`、`needs_source_text`、`do_not_backfill_from_model` 标注。

---

## 12. 仍需继续追回

最近 7 天 ChatGPT 对话证据接入已完成一层。

继续待办：

- 若用户提供具体对话标题、原文、导出、粘贴文本或截图转录，再建立 `old-chat-evidence-*` 专项 ledger；
- 若用户只要求继续主动找，则下一步可按最近 7 天对话主题逐项建立 `recent-chat-topic-index`：PR #372、浏览器清理、清单系统、晚间见闻、任务设置等；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
