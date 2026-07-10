# PR #372 判断链追回：最近 7 天 ChatGPT 对话主题索引（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-recent-chat-evidence-intake-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-recent-chat-first-batch-evidence-20260706.md`

本轮依据用户要求继续主动翻找最近 7 天 ChatGPT 对话，建立主题级索引。它不是旧对话全文迁移，也不是逐源判断链完成证明。

本文件只记录：

1. 最近 7 天对话中哪些主题与 PR #372 判断链追回有关；
2. 哪些主题能直接转为判断链；
3. 哪些只是状态、操作或 CI 证据，不能当来源价值判断；
4. 哪些已经被后续 GitHub ledger 承接；
5. 哪些仍需要原始对话或具体页面证据。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：最近对话主题索引不是逐源完成证明

```yaml
- source_id: "最近 7 天 ChatGPT 对话主题索引"
  current_route: topic_index_not_source_level_completion
  decision_status: recent_chat_topic_index_recovered
  selected_reason: "最近 7 天对话里有 PR #372、浏览器清理、清单系统、晚间见闻、任务设置、TopHub/Folo 系统设置、社区与地方门户、App 目录、综合动作、36氪科技页、CI 失败等相关主题。"
  rejected_alternatives:
    - "把主题索引当完整旧对话原文"
    - "把主题事实直接转成每个来源的 why_in / why_out"
    - "把状态和 CI 邮件当来源价值判断"
    - "因为找到最近对话索引就宣布旧对话迁移完成"
  route_reason: "主题索引只能告诉后续应追哪条线，不能替代逐源证据。"
  reuse_rule: "后续继续追回时，先按主题定位，再找原文、上传文本、已沉积 GitHub 文件或平台核验证据；没有证据就标 `needs_source_text`。"
  evidence_locator: "recent ChatGPT conversation context 2026-07-03 to 2026-07-06"
  next_action: "作为主动翻找路线图。"
```

---

## 2. PR #372 主线：判断链追回，不是合并 PR

```yaml
- source_id: "PR #372 最近对话主线"
  current_route: recovery_context_not_merge_instruction
  decision_status: recovered_from_recent_chat
  selected_reason: "最近对话确认 PR #372 是 `订阅系统 v0：按 TopHub 能力重构 Folo 与沃壤分工`，分支为 `feat/folo-subscription-system-v0-20260703`，核心工作是追回 TopHub / Folo / 沃壤分工背后的判断链。"
  rejected_alternatives:
    - "继续尝试 merge Draft PR"
    - "把 PR 是否值得合并作为当前主任务"
    - "只总结 PR 改了哪些文件"
    - "只看 TopHub 80 / Folo 27 数字"
  route_reason: "最近对话中曾尝试合并但因为 PR 仍是 Draft 而失败；后续用户转向要求追回判断链。因此当前任务不是合并，而是补证据。"
  reuse_rule: "涉及 PR #372 的后续动作只新增证据文件或 ledger；不合并、不改 Ready、不用数字正确代替判断链。"
  evidence_locator: "recent ChatGPT conversations 2026-07-06; GitHub PR #372 context"
  next_action: "保持 Draft / open / unmerged。"
```

---

## 3. TopHub 设置与系统功能：平台能力不等于启用

```yaml
- source_id: "TopHub 设置与功能对话"
  current_route: system_configuration_evidence_already_covered
  decision_status: recovered_from_recent_chat_superseded_by_system_ledgers
  selected_reason: "最近 7 天对话确认用户曾粘贴 TopHub 功能：首页、日报、动态、追踪、榜中榜、热文库、话题、日历，分类包括综合、科技、娱乐、社区、购物、财经、开发、简报、AI，以及账号设置、过滤、通知、投稿节点。"
  rejected_alternatives:
    - "因为 TopHub 支持功能就全部启用"
    - "启用通知机器人和外部通知"
    - "把过滤器当判断替代品"
    - "把日报、榜中榜、热文库、热点日历当默认首页"
  route_reason: "当时形成的分工是：Folo=关系层，TopHub=公共世界 / 阶段性雷达，沃壤=判断沉积；平台能力需要按使用目的启用，而不是按存在感启用。后续已由系统层 ledger 承接。"
  reuse_rule: "TopHub 新功能先问是否改善路由；不能因为功能存在就启用。"
  evidence_locator: "recent ChatGPT conversation 2026-07-03; tophub account/widgets/feature/system ledgers"
  next_action: "已由系统 ledgers 承接。"
```

---

## 4. TopHub 首轮取消 28 项：只能作为早期动作，不能扩展为无限删除

```yaml
- source_id: "TopHub 首轮 high-duplication unsubscribe list"
  current_route: early_operation_fact_not_open_ended_deletion_rule
  decision_status: recovered_from_recent_chat_partial
  selected_reason: "最近 7 天对话确认：早期完成 account settings 后，曾给出第一波 TopHub 取消列表，约 28 个高重复节点，并明确警告不要继续删除超过这一波。"
  rejected_alternatives:
    - "把第一波取消扩展成继续无限删除"
    - "把早期取消列表当当前完整状态"
    - "用旧取消动作覆盖后续平台核验"
  route_reason: "这是早期操作事实，不是当前判断链完整证据。后续已有平台核验、当前顺序和失效节点移除文件覆盖当前状态。"
  reuse_rule: "早期操作清单只能作时间线证据；当前状态回最新平台核验和全部订阅顺序。"
  evidence_locator: "recent ChatGPT conversation 2026-07-03; pending/platform verification ledgers"
  next_action: "无需重开早期取消列表。"
```

---

## 5. 综合动作：新闻调查、辟谣、半月谈健康进入，wikiHow / 壹心理退出

```yaml
- source_id: "综合大类最近对话动作"
  current_route: recovered_recent_chat_action_already_covered_by_comprehensive_ledgers
  decision_status: recovered_from_recent_chat_superseded_by_comprehensive_ledgers
  selected_reason: "最近 7 天对话确认：TopHub 综合动作包括 `中央电视台｜新闻调查` 加入数据与结构，`科普中国｜今日辟谣文章` 和 `半月谈｜健康` 加入生活与社区，`wikiHow 中文｜首页推荐` 与 `壹心理` 被移除。"
  rejected_alternatives:
    - "继续保留 wikiHow 中文首页推荐"
    - "继续保留壹心理"
    - "把辟谣、健康和新闻调查放进 Folo"
    - "把综合动作说成全综合目录完成"
  route_reason: "这些动作补的是事实核验、健康公共信息与调查结构，不是长期作者关系。后续已由综合 action-sources / institutional-media / comprehensive ledger 承接。"
  reuse_rule: "综合来源按功能位判断：辟谣与健康是行动 /公共信息源，新闻调查是数据与结构入口；低密度泛建议源退出。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; comprehensive ledgers"
  next_action: "已由综合 ledgers 承接。"
```

---

## 6. 36氪科技页：早期 no immediate add，后续被科技专项吸收

```yaml
- source_id: "36氪科技页最近对话事实"
  current_route: early_technology_page_fact_superseded
  decision_status: recovered_from_recent_chat_superseded_by_technology_ledgers
  selected_reason: "最近对话确认：36氪页当时基线有 13 个节点；不要新增 `36氪｜最新 / 收藏榜 / 综合榜 / 人气榜 / 热议榜 / AI频道 / 创投频道 / 资讯推荐`；`36氪｜快讯` 是 2020 陈旧缓存，`未来汽车日报｜最新要闻` 是 2024 陈旧节点；当时 strongest replacement candidate 是 `36氪｜深氪` 对比既有 `36氪｜24小时热榜`，但 no immediate add / no tracker / no Folo change。"
  rejected_alternatives:
    - "立即新增 36氪 深氪"
    - "新增 36氪 多个榜单"
    - "用陈旧快讯或未来汽车日报补科技雷达"
    - "立即改 Folo 或追踪器"
  route_reason: "这是早期科技页判断，后来被科技商业全球化、当前雷达与任务触发 ledger 吸收。"
  reuse_rule: "早期候选不等于当前待办；看当前科技状态时回科技专项 ledger。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; technology ledgers"
  next_action: "不再作为当前待办。"
```

---

## 7. App 目录六页：用户提供六页就是本轮完整范围

```yaml
- source_id: "TopHub App 目录六页"
  current_route: bounded_page_scope_closed_for_that_round
  decision_status: recovered_from_recent_chat_partial_to_technology_context
  selected_reason: "最近对话确认：用户明确只给前六页，不主动浏览 TopHub 补页，不把未提供的地区 App Store 榜单写入附录；六页就是本轮 App 审核完整范围。"
  rejected_alternatives:
    - "主动补浏览更多 App 页"
    - "把未提供地区榜单写入附录"
    - "把六页闭环说成整个 App 宇宙完成"
  route_reason: "该边界控制的是审查范围，不是来源价值本身。后续记录显示 App 目录六页闭环，新增候选 `反斗软件｜最新发布` 进科技雷达，科技候选净变化修订为 +2。"
  reuse_rule: "用户限定页数时，严格按用户提供范围闭环；不主动补页扩张。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; tophub technology app closure file"
  next_action: "已由科技 closeout 承接；若要逐源展开需读 App closure。"
```

---

## 8. 社区 page 1 与地方门户：候选和闭环边界分开

```yaml
- source_id: "社区 page 1 / 地方门户最近对话事实"
  current_route: community_partial_and_local_portal_closed_boundary
  decision_status: recovered_from_recent_chat_superseded_by_community_ledgers
  selected_reason: "最近对话确认：社区 page 1 audit 已上 GitHub；当时无 immediate adds / cancels，只留下 `马蜂窝｜热门游记` 和 `水木社区｜十大热门话题` 两个候选。地方门户 49/49 closed，无新增订阅；深圳论坛和高楼迷只作按需工具。"
  rejected_alternatives:
    - "把社区 page 1 说成社区 235 节点全量闭环"
    - "把马蜂窝和水木都写成已执行"
    - "继续把地方门户列为未闭环"
    - "把深圳论坛 / 高楼迷常驻"
  route_reason: "社区 page 1 是早期部分审查，地方门户则已闭环。后续马蜂窝已进入生活与社区，水木仍为候选或任务触发。社区整体仍 partial。"
  reuse_rule: "同一大类内要区分 page-level partial、candidate、executed 和 closed subdirectory。"
  evidence_locator: "recent ChatGPT conversations 2026-07-05; community ledgers"
  next_action: "社区整体 partial；地方门户不重开。"
```

---

## 9. CI 失败邮件：操作证据，不是来源选择证据

```yaml
- source_id: "PR #372 CI failure notifications"
  current_route: operations_evidence_not_source_judgment
  decision_status: recovered_from_recent_gmail_context
  selected_reason: "最近 7 天 Gmail / 对话上下文多次出现 PR #372 和分支 `feat/folo-subscription-system-v0-20260703` 的 CI Linux 失败通知，说明该 PR 在 2026-07-03 至 2026-07-04 期间有持续提交和失败检查。"
  rejected_alternatives:
    - "把 CI 失败当来源价值判断"
    - "用 CI 失败解释 TopHub / Folo 取舍"
    - "因为 CI 失败就停止判断链追回"
  route_reason: "CI 失败只说明工程操作状态，不说明来源为何选或不选。它可进入 docs/ops 或今日报状态核验，但不能作为 source judgment ledger 的 why_in / why_out。"
  reuse_rule: "CI / PR 通知用于操作层；来源判断回 ledger、closure、平台核验或旧对话证据。"
  evidence_locator: "recent Gmail context 2026-07-03 to 2026-07-04"
  next_action: "作为 ops 状态，不进来源判断。"
```

---

## 10. 浏览器清理、清单系统、晚间见闻、任务设置：相关但不全属于 PR #372 逐源判断

```yaml
- source_id: "最近 7 天其他主题：浏览器清理 / 清单系统 / 晚间见闻 / 任务设置"
  current_route: adjacent_topics_need_separate_filter_before_pr372_recovery
  decision_status: recent_chat_topic_indexed_needs_specific_source_text
  selected_reason: "用户明确说最近 7 天对话里还有浏览器清理、清单系统、晚间见闻、任务设置等相关判断。"
  rejected_alternatives:
    - "全部纳入 PR #372 source ledger"
    - "全部排除为无关"
    - "用主题名直接补来源判断"
  route_reason: "这些主题与信息系统、订阅、浏览器、自动任务和内容呈现有关，但不一定都是 PR #372 的逐源判断。需要先过滤：哪些是来源路由，哪些是工作流，哪些是定时任务，哪些是内容生成失败或能力边界。"
  reuse_rule: "相邻主题先做 topic triage：source routing → 可纳入 PR #372；workflow / automation / content generation → 另立 ops 或 workflow ledger；个人表达 / 清单哲学 → 不强行塞进订阅源 ledger。"
  evidence_locator: "recent ChatGPT conversation context 2026-07-04 to 2026-07-06"
  next_action: "下一轮可逐项做 topic triage。"
```

---

## 11. 本补充 ledger 的复用规则

1. 最近对话主题索引只定路线，不替代逐源原文。
2. PR #372 主线是判断链追回，不是合并 PR。
3. TopHub 功能设置里的核心判断是平台能力不等于启用。
4. 早期取消列表不能扩展为无限删除。
5. 综合动作已由后续 ledger 承接，但不能宣称综合全量完成。
6. 36氪 early fact 是时间边界，已被科技专项吸收。
7. App 六页遵守用户限定范围，不主动补页。
8. 社区 page 1、地方门户、候选与已执行要分开。
9. CI 失败是操作证据，不是来源选择证据。
10. 浏览器清理、清单系统、晚间见闻、任务设置需要下一轮 topic triage，不直接塞进 PR #372 source ledger。

---

## 12. 仍需继续追回

最近 7 天 ChatGPT 对话主题索引已建立。

继续待办：

- 下一轮建议做 `recent-chat-topic-triage`：浏览器清理、清单系统、晚间见闻、任务设置分别属于 source routing / workflow / automation / content generation / personal philosophy 哪一类；
- 对属于 source routing 的条目，再建立 `recent-chat-source-evidence-*`；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
