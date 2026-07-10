# PR #372 判断链追回：最近对话 source_routing 残余证据（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-recent-chat-topic-index-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-recent-chat-topic-triage-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-browser-link-candidate-cohorts-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-community-podcast-shopping-dev-finance-bridge-20260706.md`

本轮只处理最近 7 天对话中被分到 `source_routing` 的残余证据：

1. 浏览器订阅候选；
2. TopHub 早期首轮取消约 28 个高重复节点；
3. App 目录六页；
4. 社区 page 1 与地方门户。

本文件不把摘要冒充完整旧对话原文；不把早期动作冒充当前状态；不把候选写成已执行。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：source_routing 残余证据只补缺口，不重开已承接文件

```yaml
- source_id: "最近对话 source_routing 残余证据"
  current_route: residual_source_evidence_not_full_reaudit
  decision_status: recent_chat_source_evidence_partial
  selected_reason: "最近 7 天对话中仍有若干 source_routing 材料能定位到主题和部分判断，但没有完整逐源原文。它们包括浏览器订阅候选、早期取消 28 项、App 六页、社区 page 1 与地方门户。"
  rejected_alternatives:
    - "重复重审所有已由专项 ledger 承接的文件"
    - "把最近对话摘要当逐源原文"
    - "把 partial 残余证据写成 completed"
    - "把候选写成已执行"
  route_reason: "这些材料的价值是补时间线、范围边界、候选 / 执行 / 按需分层，不是替代完整来源证据。"
  reuse_rule: "能从最近对话确认的写入 `recovered_from_recent_chat_partial`；已由后续 ledger 承接的写 `superseded_by_later_ledger`；缺逐项原文的写 `needs_source_text_for_item_level`。"
  evidence_locator: "recent ChatGPT context; recent-chat-topic-triage ledger"
  next_action: "作为残余 source evidence。"
```

---

## 2. 浏览器订阅候选：持续更新源进入复查队列，不留浏览器

```yaml
- source_id: "浏览器订阅候选 cohort"
  current_route: folo_tophub_review_queue_2026_07_17
  decision_status: recovered_from_recent_chat_superseded_by_browser_candidate_ledger
  selected_reason: "最近对话确认：浏览器链接治理第一轮和第二轮中，许多媒体、平台和持续更新入口被判为 Folo / TopHub 候选；理由是持续更新的来源应由订阅系统承担‘到来’，不在浏览器长期堆媒体首页。"
  rejected_alternatives:
    - "继续留在浏览器收藏夹作为日常打开对象"
    - "直接写入 Folo / TopHub 正式账本"
    - "全部迁入资源正式条目"
    - "在 2026-07-17 前扩容 Folo"
  route_reason: "浏览器保存抵达，Folo / TopHub 保存到来，沃壤保存判断。订阅候选进入 2026-07-17 复查队列，等待角色、重复度、打开率、未读压力和一进一出判断。"
  reuse_rule: "凡是持续更新的媒体 / 平台入口，先判断 TopHub 是否已覆盖公共温度；若是长期原文关系才进 Folo 复查；浏览器不再承担持续到来。"
  evidence_locator: "recent ChatGPT context; 资源/链接/浏览器订阅候选-2026-07-17.md; browser-link-candidate-cohorts ledger"
  next_action: "2026-07-17 复查。"
```

---

## 3. 浏览器订阅候选内部：门户、促销、财经日历不能一键升级

```yaml
- source_id: "浏览器订阅候选中的门户 / 促销 / 财经日历"
  current_route: candidate_subtypes_need_separate_review
  decision_status: recovered_from_recent_chat_superseded_by_browser_candidate_ledger
  selected_reason: "最近对话与候选文件显示，订阅候选内混有中文综合门户、国际媒体、科技媒体、财经商业入口、促销 / 行动触发入口、财经日历和视觉自然来源。"
  rejected_alternatives:
    - "把首页和门户都当合格订阅源"
    - "把促销 / 喜加一 / 商城入口加入 Folo"
    - "把财经日历变成每日未读流"
    - "因为候选文件存在就写入正式订阅源账本"
  route_reason: "门户多与 TopHub 公共温度重复；促销入口只有明确领取或购买动作时有价值；财经日历是查询工具，不是阅读源；国际 / 科技 / 视觉来源则需看是否与既有 TopHub / Folo 重复。"
  reuse_rule: "订阅候选先分 subtype：公共门户、作者 / 媒体关系、工具日历、促销动作、视觉漫游。不同 subtype 不用同一迁移规则。"
  evidence_locator: "recent ChatGPT context; browser subscription candidate file; browser-link-candidate-cohorts ledger"
  next_action: "候选 subtype 复查。"
```

---

## 4. TopHub 早期首轮取消约 28 项：时间线证据，不是当前清单

```yaml
- source_id: "TopHub 早期首轮取消约 28 项"
  current_route: early_unsubscribe_operation_fact_needs_item_text_for_source_level
  decision_status: recovered_from_recent_chat_partial_needs_source_text_for_item_level
  selected_reason: "最近对话确认：早期完成 TopHub account settings 后，曾给出第一波取消列表，约 28 个高重复节点，并明确提醒不要继续删除超过这一波。"
  rejected_alternatives:
    - "把这 28 项作为当前全部退订清单"
    - "把首轮取消扩展为无限删除规则"
    - "用早期取消覆盖后续平台核验"
    - "在没有原始 28 项文本时逐项编理由"
  route_reason: "这是一条重要时间线证据：早期系统先去掉高重复节点，但随后进入审慎阶段。当前真实状态必须回后续全部订阅顺序、pending / platform verification、失效节点移除核验和专项 ledger。"
  reuse_rule: "早期批量取消只用于说明系统从扩张转向压缩；若要逐项追回 28 个来源为什么退，必须先取得原始取消列表文本。"
  evidence_locator: "recent ChatGPT conversation 2026-07-03; pending/platform verification ledgers"
  next_action: "needs_source_text_for_item_level。"
```

---

## 5. TopHub 早期取消的复用边界：不能因为旧重复就永久拒绝同类

```yaml
- source_id: "TopHub 早期取消规则"
  current_route: early_dedup_principle_not_permanent_ban
  decision_status: recovered_from_recent_chat_partial
  selected_reason: "早期取消列表的共同方向是高重复、低独特增量、旧状态或平台切片过多。"
  rejected_alternatives:
    - "把被取消来源所属品牌永久拉黑"
    - "未来同类来源一律拒绝"
    - "用删除数量衡量系统质量"
  route_reason: "删除的是当时那些节点的角色，不是永久拒绝其品牌或主题。后续仍可能按功能位、更新状态、低噪声和结构缺口重新引入同类更合适入口。"
  reuse_rule: "评估同类来源时，不问‘这个品牌以前删过吗’，而问‘这个具体节点现在是否补独特功能位，是否低重复，是否有证据更新，是否比旁边入口更好’。"
  evidence_locator: "recent ChatGPT context; stale node and comprehensive / technology / entertainment replacement ledgers"
  next_action: "作为早期取消复用规则。"
```

---

## 6. App 目录六页：范围由用户给定，不主动补页

```yaml
- source_id: "TopHub App 目录六页"
  current_route: bounded_page_scope_closed_for_round
  decision_status: recovered_from_recent_chat_partial_superseded_by_technology_closeout
  selected_reason: "最近对话确认：用户明确只给前六页，不主动浏览 TopHub 补页，不把未提供的地区 App Store 榜单写入附录；六页就是本轮 App 审核完整范围。"
  rejected_alternatives:
    - "主动补浏览更多 App 页"
    - "把未提供地区榜单写入附录"
    - "把六页闭环说成整个 App 宇宙完成"
  route_reason: "该判断首先是范围边界：尊重用户提供页面，不扩张任务。后续科技 closeout 已记录 App 目录六页闭环。"
  reuse_rule: "用户限定页数 / 页面时，按给定范围闭环；不把未提供页面补成新审计义务。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; tophub technology closeout"
  next_action: "若要逐源展开，读 App closure 原始文件。"
```

---

## 7. 反斗软件｜最新发布：App 六页中进入科技雷达的候选

```yaml
- source_id: "反斗软件｜最新发布"
  current_route: executed_or_candidate_tophub_technology_radar_according_to_later_closeout
  decision_status: recovered_from_recent_chat_partial_superseded_by_technology_ledger
  selected_reason: "最近对话确认：App 目录六页闭环后，新增候选 `反斗软件｜最新发布` 进入科技雷达，科技候选净变化修订为 +2。"
  rejected_alternatives:
    - "把所有 App Store 地区榜单加入科技雷达"
    - "把 App 目录整体加入 TopHub 常驻"
    - "把 App 榜单作为应用质量判断"
  route_reason: "反斗软件是软件 / 应用发现入口，区别于 App Store 销量、免费、付费或地区榜单。它补的是具体软件发布发现，而不是市场排名。"
  reuse_rule: "App / 软件来源先区分软件发布、市场排名、促销、下载站、评测、官方 changelog；只有低噪声发布发现才可能进科技雷达。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; technology closeout / App closure"
  next_action: "已由科技 closeout 承接；若要补更细，读取 App closure。"
```

---

## 8. App Store 地区榜单：市场样本，不主动写入附录

```yaml
- source_id: "未提供的地区 App Store 榜单"
  current_route: not_reviewed_not_backfilled
  decision_status: recovered_from_recent_chat_boundary
  selected_reason: "最近对话确认：不把未提供的地区 App Store 榜单写入附录。"
  rejected_alternatives:
    - "主动补全地区榜单"
    - "把未审地区榜单当已判断"
    - "把 App Store 排名当应用质量判断"
  route_reason: "App Store 榜单本质是市场天气图，且本轮只审用户提供的六页。未提供页面不构成审计对象。"
  reuse_rule: "未提供页面不反向补；App Store 榜单按具体市场研究任务调用。"
  evidence_locator: "recent ChatGPT conversation 2026-07-04; early App Store audit bridge ledger"
  next_action: "not_reviewed。"
```

---

## 9. 社区 page 1：无 immediate adds / cancels，只留下两个候选

```yaml
- source_id: "社区 page 1 audit"
  current_route: page_level_partial_with_two_candidates
  decision_status: recovered_from_recent_chat_partial_superseded_by_community_ledgers
  selected_reason: "最近对话确认：社区 page 1 audit 已上 GitHub；当时无 immediate adds / cancels，只留下 `马蜂窝｜热门游记` 和 `水木社区｜十大热门话题` 两个候选。"
  rejected_alternatives:
    - "把社区 page 1 完成说成社区 235 节点全量闭环"
    - "把两个候选都写成已执行"
    - "用 page 1 结果覆盖社区热门第 2—20 页缺口"
  route_reason: "page 1 是社区目录的部分审查；后续马蜂窝已进入生活与社区，水木仍为候选 / 任务触发。社区整体仍 partial。"
  reuse_rule: "社区页级审查必须保留页码边界；page-level candidate 不等于 directory-level completion。"
  evidence_locator: "recent ChatGPT conversation 2026-07-05; community ledgers"
  next_action: "社区整体 partial。"
```

---

## 10. 马蜂窝｜热门游记：候选后来执行，补旅行经验

```yaml
- source_id: "马蜂窝｜热门游记"
  current_route: executed_tophub_life_community_travel_experience_entry
  decision_status: recovered_from_recent_chat_superseded_by_platform_verification_and_community_ledgers
  selected_reason: "最近对话确认它最早是社区 page 1 两个候选之一；后续 pending actions / community ledger 已确认它进入生活与社区，角色是完整游记、路线、自驾、徒步和旅行现场经验。"
  rejected_alternatives:
    - "把马蜂窝候选继续当未执行"
    - "把所有旅游社区都常驻"
    - "把背包客栈同时加入"
  route_reason: "马蜂窝补的是生活与社区里的旅行经验缺口；背包客栈保留为国际旅行连续项目时的补充按需入口。"
  reuse_rule: "旅行经验源先补一个主入口；任务连续化后再补第二入口。"
  evidence_locator: "recent ChatGPT conversation 2026-07-05; community pending actions ledger"
  next_action: "已执行。"
```

---

## 11. 水木社区｜十大热门话题：候选，不等于已执行

```yaml
- source_id: "水木社区｜十大热门话题"
  current_route: future_candidate_or_short_trial_not_executed
  decision_status: recovered_from_recent_chat_partial
  selected_reason: "最近对话确认它是社区 page 1 早期候选之一，与马蜂窝并列出现。"
  rejected_alternatives:
    - "把水木写成已执行"
    - "把高校 / 老论坛社区直接常驻"
    - "把它当社区目录闭环证据"
  route_reason: "水木可观察高校 / 职业 / 城市 / 技术人群旧式论坛气质，但噪声、上下文和社区语境较强，需未来任务触发或短期试验，不应自动加入。"
  reuse_rule: "老论坛 / 高校论坛来源先短期试验或任务触发；不要因群体样本价值直接常驻。"
  evidence_locator: "recent ChatGPT conversation 2026-07-05; community ledgers"
  next_action: "候选 / 任务触发。"
```

---

## 12. 地方门户：49/49 闭环，无新增订阅

```yaml
- source_id: "地方门户 49/49"
  current_route: closed_no_subscription_additions
  decision_status: recovered_from_recent_chat_superseded_by_community_ledgers
  selected_reason: "最近对话确认：地方门户 49/49 closed，无新增订阅；深圳论坛和高楼迷只作按需工具。"
  rejected_alternatives:
    - "继续把地方门户列为未闭环"
    - "把深圳论坛常驻"
    - "把高楼迷常驻"
    - "为地方门户建立 Folo 或 TopHub 组"
  route_reason: "地方门户高度依赖具体城市、住房、规划、交通、学校、政务或生活任务。没有明确城市任务时常驻会制造噪声。"
  reuse_rule: "地方门户按城市 / 事项触发，任务结束退出；不做全国地方论坛日常流。"
  evidence_locator: "recent ChatGPT conversation 2026-07-05; community ledgers"
  next_action: "已闭环，不重开。"
```

---

## 13. 本补充 ledger 的复用规则

1. 浏览器订阅候选属于 PR #372 source_routing，但仍是复查队列，不是正式订阅。
2. 订阅候选内部要按 subtype 分：门户、作者 / 媒体关系、工具日历、促销动作、视觉漫游。
3. TopHub 首轮约 28 个取消只作早期时间线；没有原始 28 项文本，不逐项补编。
4. 早期取消不是永久拉黑品牌；未来同类来源仍可按功能位和证据重新判断。
5. App 六页按用户提供范围闭环，不主动补页。
6. 反斗软件补软件发布发现，不等于 App Store 榜单质量判断。
7. 未提供地区 App Store 榜单不审、不补、不写附录。
8. 社区 page 1 是 page-level partial，不是社区目录全量闭环。
9. 马蜂窝候选后来已执行；水木仍是候选 / 任务触发。
10. 地方门户 49/49 已闭环，无新增订阅，按城市任务触发。

---

## 14. 仍需继续追回

最近对话 source_routing 残余证据已补一层。

继续待办：

- 若要逐项追回 TopHub 早期首轮 28 个取消项，必须先获得原始 28 项列表文本；
- 若要逐源展开 App 六页，需读取或找到 `tophub-technology-app-closure` 相关原始文件；
- 若要继续 recent-chat source evidence，可转向浏览器订阅候选 48 条中的具体条目，但必须只抽文件中已有理由；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
