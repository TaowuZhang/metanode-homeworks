# PR #372 判断链追回：旧 ChatGPT 对话 partial 追索边界（2026-07-06）

## 0. 边界

本文件记录一次旧 ChatGPT 对话 partial 追索的边界核验。

本轮已在仓库内搜索：

- `ChatGPT 旧对话 粘贴 文本 对话 partial 判断链 追索`
- `旧 ChatGPT 对话 订阅源 TopHub Folo 判断链`
- `粘贴的文本 Folo TopHub 订阅系统 判断链`

结果：当前 PR #372 仓库分支中没有找到可直接引用的旧 ChatGPT 对话原文、完整粘贴文本或可作为逐源判断证据的旧对话沉积文件。

因此，本文件不是旧对话内容迁移成果，而是旧对话追索的边界账：说明什么可以从现有文件继续使用，什么不能由当前模型补编，什么必须等待用户提供旧对话材料或本地提交后再做。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：没有旧对话原文，就不能补编旧判断

```yaml
- source_id: "旧 ChatGPT 对话 partial 追索"
  current_route: blocked_until_source_material_available
  decision_status: boundary_recovered_no_source_text_found
  selected_reason: "用户要求追回每个来源背后的判断链，且明确不能由当前模型新编理由。旧 ChatGPT 对话若未沉积到仓库，当前只能从已有 PR 文件和已创建 ledger 中追回，不得把记忆、印象或模型推测写成旧对话事实。"
  rejected_alternatives:
    - "凭当前模型重新解释旧对话可能说过什么"
    - "把后续 ledger 的理由倒灌成旧 ChatGPT 对话原始判断"
    - "把 source index 中的状态当旧对话证据"
    - "把没有原文的 partial 批量标记 completed"
  route_reason: "仓库搜索未找到旧 ChatGPT 对话原文或粘贴文本。没有来源材料时，只能写追索边界，不能写旧对话内容迁移成果。"
  reuse_rule: "以后做旧对话追索，必须先有原文、摘要、粘贴文件、截图转录或仓库沉积证据；没有证据只登记缺口。"
  evidence_locator: "GitHub search results for old ChatGPT / pasted text queries; existing PR #372 coverage map"
  next_action: "等待用户提供旧对话材料，或在仓库出现相应沉积文件后继续。"
```

---

## 2. 可以继续使用的证据：现有 PR 文件与 ledger

```yaml
- source_id: "现有 PR 文件与新增 ledger"
  current_route: usable_recovered_evidence_layer
  decision_status: recovered_from_pr
  selected_reason: "PR #372 已经沉积了大量 TopHub、Folo、资源、浏览器、文化、设备、链接、资源更新、平台核验和 final closure 文件；后续又新增多份 pr372-decision-recovery-ledger-* 文件，已把许多判断链从 PR 文件中追回。"
  rejected_alternatives:
    - "因为没有旧对话原文就否定已有 ledger"
    - "重新从头审查所有 TopHub / Folo 来源"
    - "重复挖已经由专项 ledger 承接的原始文件"
  route_reason: "已有 PR 文件和 ledger 仍是有效证据。反向覆盖映射已说明哪些原始文件由哪些 ledger 承接。旧对话追索缺口不影响已经从仓库文件追回的判断链。"
  reuse_rule: "继续工作时优先查 `pr372-decision-recovery-original-file-coverage-map-20260706.md`，避免重复挖已覆盖文件。"
  evidence_locator: "docs/ops/pr372-decision-recovery-original-file-coverage-map-20260706.md"
  next_action: "作为当前可用证据层。"
```

---

## 3. 不能使用的替代证据

```yaml
- source_id: "旧对话替代证据禁用项"
  current_route: invalid_substitutes
  decision_status: recovered_boundary
  selected_reason: "旧对话 partial 要追回的是此前实际发生过的判断链，不是当前模型对来源的合理化解释。"
  rejected_alternatives:
    - "当前模型按来源类型新编一套理由"
    - "用覆盖映射替代逐源判断"
    - "用 source index 的 `needs_recovery` / `partial` / `omitted_by_choice` 状态替代旧对话证据"
    - "把用户当前口头标准当作旧对话内容"
    - "把候选清单当已判断来源"
  route_reason: "这些材料最多说明当前路由原则、候选状态或追回需求，不能证明旧对话中曾经为何选 / 不选 / 放层 / 复用。"
  reuse_rule: "替代证据只能标注缺口，不能填充判断链正文。"
  evidence_locator: "source index and coverage map boundaries"
  next_action: "保持禁用。"
```

---

## 4. 旧对话材料进入后的处理协议

```yaml
- source_id: "旧对话原文进入后的处理协议"
  current_route: future_old_chat_recovery_protocol
  decision_status: protocol_recovered
  selected_reason: "如果后续用户提供旧 ChatGPT 对话原文、粘贴文件、截图转录或已经沉到仓库的对话材料，需要有一个不污染现有 ledger 的处理方式。"
  rejected_alternatives:
    - "整段复制旧对话进 ledger"
    - "只摘结论不摘判断链"
    - "把旧对话里的错误、误判或过时数量当当前事实"
    - "没有证据 locator 就写 completed"
  route_reason: "旧对话进入后应按来源逐条抽取：来源名、当时页面或候选、为什么选、为什么不选相邻来源、当时放在哪一层、当时的复用规则、后来是否被平台执行或后续 ledger 覆盖。旧对话中的数量、状态和建议必须带时间边界。"
  reuse_rule: "旧对话 ledger 采用 `old_chat_evidence` 状态；若旧判断已被后续平台核验覆盖，写 `superseded_by_platform_verification`；若旧判断只是想法，写 `candidate_only`；若证据不足，写 `needs_source_text`。"
  evidence_locator: "future user-provided old chat material"
  next_action: "等待材料。"
```

---

## 5. 旧对话条目状态词

```yaml
- source_id: "旧对话追索状态词"
  current_route: status_vocabulary_for_old_chat_recovery
  decision_status: protocol_recovered
  selected_reason: "旧对话追索需要比 `covered / partial` 更细的状态词，避免把材料性质混在一起。"
  rejected_alternatives:
    - "统一写 completed"
    - "统一写 not_found"
    - "统一写 recovered_from_pr"
  route_reason: "旧对话可能有原始判断、候选想法、被后续覆盖的建议、已执行动作、错误判断、过时数量、只剩口头印象或完全无证据。必须分开。"
  reuse_rule: "建议使用：`old_chat_evidence_recovered`、`old_chat_candidate_only`、`superseded_by_later_ledger`、`superseded_by_platform_verification`、`needs_source_text`、`do_not_backfill_from_model`。"
  evidence_locator: "this boundary ledger"
  next_action: "作为后续旧对话 ledger 状态词。"
```

---

## 6. 当前可确认的未完成边界

```yaml
- source_id: "旧 ChatGPT 对话 partial 未完成边界"
  current_route: explicit_unfinished_boundary
  decision_status: not_completed
  selected_reason: "当前没有可引用旧对话原文；用户本地仍有 master checklist、coverage audit、military boundary ledger 等待提交；社区热门第 2—20 页仍无实际页面。"
  rejected_alternatives:
    - "宣布旧 ChatGPT 对话已迁移"
    - "宣布 PR #372 判断链全完成"
    - "宣布 TopHub 80 + Folo 27 逐源判断全补齐"
  route_reason: "没有证据的完成不能说完成。旧对话追索目前只完成边界核验，未完成内容迁移。"
  reuse_rule: "对外只能说：旧对话追索已建立边界，尚未迁移旧对话内容。"
  evidence_locator: "this boundary ledger; original coverage map"
  next_action: "等待旧对话材料或转入用户指定的具体缺口。"
```

---

## 7. 本补充 ledger 的复用规则

1. 没有旧对话原文、粘贴文件、截图转录或仓库沉积证据，不能写旧对话判断链。
2. 现有 PR 文件与 ledger 仍是有效证据；旧对话缺口不否定已追回内容。
3. source index、coverage map、候选清单和当前模型解释都不能替代旧对话证据。
4. 旧对话进入后必须逐源抽取判断链，并带时间边界。
5. 旧判断若已被后续平台核验覆盖，应写被覆盖，而不是删除旧判断。
6. 当前只能宣布“旧对话追索边界已建立”，不能宣布旧对话迁移完成。

---

## 8. 仍需继续追回

旧 ChatGPT 对话 partial 追索目前只完成边界账。

继续待办：

- 用户若提供旧对话原文、粘贴文本或本地文件路径，再建立 `old-chat-evidence-*` 专项 ledger；
- 等用户本地提交 master checklist、coverage audit、military boundary ledger 后，再统一登记；
- 若不提供旧对话材料，则下一步只能处理用户指定的具体来源缺口，或继续等本地提交后做控制文件登记。
