# PR #372 判断链追回：pending actions 与平台核验边界（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `pending-actions` 与 `platform-execution-verified` 文件形成的判断链。它们不是来源逐项 ledger，而是用于区分：建议、待办、用户已执行、平台已核验、后续又被新核验覆盖。

纳入文件：

- `docs/ops/tophub-community-pending-actions-20260705.md`
- `docs/ops/tophub-entertainment-pending-actions-20260705.md`
- `docs/ops/tophub-newspapers-platform-execution-verified-20260705.md`
- `docs/ops/tophub-finance-platform-execution-verified-20260705.md`
- `docs/ops/tophub-stale-node-removal-platform-verified-20260705.md`
- 相关专项 ledger：community、entertainment、newspapers、finance、stale-design-visual-nature。

边界说明：

- `pending-actions` 文件会随平台执行与后续核验改变状态；不能把旧待办永久化。
- `platform-execution-verified` 是当时平台证据；如果后续又发生取消、替换或失效移除，应以后续核验为当前基线。
- 当前真实 TopHub 基线仍是 80 项。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：pending、verified、current 三层分开

```yaml
- source_id: "pending-actions / platform-verified 文件组"
  current_route: state_transition_evidence_layers
  decision_status: recovered_from_pr
  selected_reason: "pending-actions 记录待执行或阶段动作；platform-execution-verified 记录用户已在平台执行并提供列表后的核验；当前基线由最新核验和当前顺序文件决定。"
  rejected_alternatives:
    - "把旧 pending 永久当待办"
    - "把 verified 当永远不变的当前状态"
    - "用计划文件覆盖平台证据"
    - "用单个分组核验覆盖全部页"
  route_reason: "订阅系统有连续状态迁移：建议 → 用户执行 → 平台列表核验 → 后续变更。每一层都只能回答自己的问题。"
  reuse_rule: "以后读 PR #372 文件时，先标注该文件是计划、待办、执行核验还是当前基线；不要跨层推断。"
  evidence_locator: "included pending and platform verification files"
  next_action: "作为状态层规则。"
```

---

## 2. 社区 pending：部分动作已完成，但社区目录仍为 partial

```yaml
- source_id: "社区大类 pending actions"
  current_route: executed_selected_actions_but_directory_partial
  decision_status: recovered_from_pr
  selected_reason: "社区文件确认本阶段已完成三个动作：先知社区精华推荐进入科技雷达，V2EX 周报进入生活与社区，马蜂窝热门游记进入生活与社区；分组排序和全部页位置均已完成。"
  rejected_alternatives:
    - "把社区目录 235 个节点宣称全量闭环"
    - "因为三个动作完成就关闭所有社区目录缺口"
    - "把虎扑、LINUXDO 继续列为待审缺口"
  route_reason: "文件明确社区目录 235 个节点全量审核仍未整体闭环，缺口是社区热门总目录第 2—20 页及尚未提供实际页面内容；但虎扑已由娱乐目录完整审核，LINUXDO 为空页面闭环，不重复记为缺口。"
  reuse_rule: "一个大类可以部分动作已执行、部分标签闭环、整体仍 partial。不得把动作完成冒充全目录完成，也不得把跨目录重复入口重复审。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#状态; #已完成动作; #跨目录复用与空页闭环"
  next_action: "社区热门第 2—20 页仍是 partial。"
```

---

## 3. 社区已完成三动作：为什么选这三个、不选旁边项

```yaml
- source_id: "先知社区｜精华推荐"
  current_route: executed_tophub_technology_radar_security_entry
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "补足应用安全、漏洞研究、AI Agent、MCP、提示注入和工具信任边界雷达。"
  rejected_alternatives:
    - "先知社区无副标题节点"
    - "看雪论坛｜最新精华"
  route_reason: "先知社区无副标题节点与精华推荐重复；看雪更偏逆向工程、内核和移动安全项目，保留为项目触发按需来源。"
  reuse_rule: "安全社区先选覆盖面和精选程度最适合当前雷达的入口；更深更窄的逆向源按项目触发。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作"
  next_action: "已执行。"

- source_id: "V2EX 周报"
  current_route: executed_tophub_life_community_weekly_tech_user_observation
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "以周频观察开发者、独立创作者与技术使用者持续讨论的问题，避免每日热门、推广和中转站信息形成未读堆积。"
  rejected_alternatives:
    - "V2EX 最热主题"
    - "今日热议"
    - "首页最新"
    - "技术 / 程序员 / 分享创造等高频节点"
    - "交易、站务和单个账号历史主题聚合"
  route_reason: "周报比高频热榜更低噪声，更适合生活与社区中的技术职业社区观察。"
  reuse_rule: "论坛型社区优先选周频、精选、聚合入口；日榜和分区榜按需。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作"
  next_action: "已执行。"

- source_id: "马蜂窝｜热门游记"
  current_route: executed_tophub_life_community_travel_experience_entry
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "补足完整游记、路线、自驾、徒步和旅行现场经验。"
  rejected_alternatives:
    - "背包客栈自助旅行论坛｜背包严选好文"
    - "旅游目录其他节点"
  route_reason: "马蜂窝补当前生活与社区的旅行经验缺口；背包客栈国际旅行信息较具体，但与马蜂窝部分重叠，本轮不同时加入，若国际旅行成为连续项目再补。"
  reuse_rule: "旅行经验源先补一个主入口；国际旅行、保险、交通、装备等任务变连续时再补第二入口。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作; #高价值按需来源"
  next_action: "已执行。"
```

---

## 4. 社区高价值按需：有价值也不常驻

```yaml
- source_id: "社区高价值按需来源组"
  current_route: on_demand_task_entries_not_subscription_flow
  decision_status: recovered_from_pr
  selected_reason: "V2EX 专题节点、数码任务入口、夜航船夫、Slot 4、背包客栈、看雪、V2EX 汽车、深圳论坛、高楼迷等都有具体任务价值。"
  rejected_alternatives:
    - "全部加入 TopHub 常驻"
    - "全部加入 Folo"
    - "把任务入口当日常信息流"
  route_reason: "这些来源价值高度依赖任务：Apple / Android / 硬件 / DNS / 信用卡 / 汽车、硬件评测、国际旅行、逆向工程、购车、城市问题和商业项目研究。没有任务时常驻只会增加噪声。"
  reuse_rule: "高价值按需来源要保留触发条件；没有触发条件不升级为常驻。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#高价值按需来源"
  next_action: "按需。"
```

---

## 5. 娱乐 pending：待办已经完成，当前真实状态是 80

```yaml
- source_id: "娱乐大类执行清单"
  current_route: completed_platform_execution_record
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "用户已于 2026-07-05 在 TopHub 界面完成全部新增、取消与排序，最终状态已经写入 `资源/雷达/TopHub 全部订阅顺序.md`。"
  rejected_alternatives:
    - "把娱乐 pending 继续当待办"
    - "把 10 个新增节点再作为候选讨论"
    - "用娱乐执行前的 73 覆盖当前 80"
  route_reason: "文件明确状态为已完成。新增 10、取消 3、胶片的味道复核后保留，最终真实 TopHub 为 80，Folo 为 27。"
  reuse_rule: "pending 文件一旦写明已完成并有当前顺序文件承接，就不再是待办，只作为执行核验证据。"
  evidence_locator: "docs/ops/tophub-entertainment-pending-actions-20260705.md#状态; #最终真实状态"
  next_action: "已完成。"
```

---

## 6. 娱乐最终顺序：分组功能位已经重新定型

```yaml
- source_id: "娱乐执行后分组顺序"
  current_route: executed_group_order_map
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "娱乐执行后，数据与结构、慢读与思想、生活与社区、影游音乐、视觉与自然等分组均给出最终顺序。"
  rejected_alternatives:
    - "把新来源留在临时候选"
    - "把胶片的味道误判为与 CNU 重复"
    - "把娱乐新增全部放入影游音乐"
  route_reason: "新增来源按功能位分散：华丽志进入数据与结构；豆瓣影评与书格进入慢读与思想；音乐、剧集、娱乐新闻、独立游戏和乐影进入影游音乐；CNU 进入视觉与自然。CNU 是低频人工视觉精选，胶片的味道是摄影叙事、城市观察、胶片器材与画幅方法，二者不重复。"
  reuse_rule: "新增来源按功能位路由，不按大类来源一股脑放进娱乐分组。"
  evidence_locator: "docs/ops/tophub-entertainment-pending-actions-20260705.md#最终分组顺序; #复核后保留"
  next_action: "已执行。"
```

---

## 7. 报刊平台核验：当时财经仍待执行，后来被财经核验覆盖

```yaml
- source_id: "报刊替换平台执行核验"
  current_route: platform_verified_newspaper_replacement_with_superseded_finance_pending
  decision_status: recovered_from_pr_platform_verified_with_temporal_boundary
  selected_reason: "2026-07-05 根据全部订阅列表和公共温度分组列表核验，`新华每日电讯｜电子报`已订阅并加入公共温度，`人民日报｜电子版`已不在全部订阅中。"
  rejected_alternatives:
    - "只凭替换计划说已完成"
    - "继续保留人民日报电子版失效节点"
    - "把当时的财经待办视为当前待办"
  route_reason: "报刊替换当时真实完成，TopHub 总量仍为 84；但同一文件记录的财经替换待办后来已由财经平台执行核验覆盖。因此本文件只保留报刊执行证据和当时状态，不保留旧财经待办为当前状态。"
  reuse_rule: "一个核验文件中的待办可能被后续核验覆盖；后读文件不能只看旧待办，要查是否已有后续执行证据。"
  evidence_locator: "docs/ops/tophub-newspapers-platform-execution-verified-20260705.md#结果; #仍待执行; docs/ops/tophub-finance-platform-execution-verified-20260705.md"
  next_action: "报刊已执行；财经待办已被后续覆盖。"
```

---

## 8. 失效节点移除核验：当前 80 项基线从这里成立

```yaml
- source_id: "四个失效节点移除核验"
  current_route: platform_verified_current_80_baseline_transition
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "用户提供了移除后的全部订阅节点完整列表；`微信｜热词`、`后续｜跟踪 Live`、`国家留学网｜综合项目专栏`、`Yoho!潮流志｜每日潮闻`均已不在列表中。"
  rejected_alternatives:
    - "四个失效节点等量补位"
    - "把此前 84 项继续当当前基线"
    - "把新华每日电讯或喷嚏乐影误删"
  route_reason: "新增 0，取消 4，真实订阅基线从 84 调整为 80。此前曾被怀疑但确认仍更新的 `新华每日电讯｜电子报` 与 `喷嚏网｜乐影`仍在列表中。"
  reuse_rule: "当前 TopHub 基线以最新平台核验为准。失效节点取消不必等量补位；被怀疑节点需区分作品年份 / 更新时间与真实更新状态。"
  evidence_locator: "docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据; #数量核验; #结论"
  next_action: "作为当前 80 项基线转换证据。"
```

---

## 9. 本补充 ledger 的复用规则

1. pending、verified、current 三层分开。
2. 社区可以部分动作已执行、部分标签闭环、整体仍 partial。
3. 跨目录重复入口不重复审；空页面可闭环，未来有内容再审。
4. 已完成的 pending 文件不再是待办，而是执行核验证据。
5. 新增来源按功能位路由，不按来源大类一股脑归组。
6. 一个平台核验文件中的旧待办，可能被后续平台核验覆盖。
7. 最新平台核验把 TopHub 当前基线从 84 调整为 80。
8. 失效节点取消后不为总量补位。
9. 被怀疑节点要区分真正失效与标题中出现旧作品 / 旧年份。
10. 当前状态必须看最新顺序文件与最新平台核验，不能从旧 pending / verified 文件反推。

---

## 10. 仍需继续追回

pending actions 与平台核验边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内可继续做一轮文件名覆盖审计，确认是否还有新增前的 `tophub-*closure` 未被任一 ledger 覆盖；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
