# PR #372 判断链追回：浏览器订阅候选 48 条组级证据（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-browser-link-candidate-cohorts-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-recent-chat-source-routing-residuals-20260706.md`

本轮读取：

- `资源/链接/浏览器订阅候选-2026-07-17.md`

该文件列出了 48 条浏览器订阅候选，并明确状态是“链接候选队列，不等于已订阅、已认可或已验证”。它提供了分组与复查边界，但没有提供每个站点的完整逐条旧对话原文。

因此，本文件只补 **组级判断链**：为什么这些组留在 `资源/链接/` 候选队列，为什么不直接进入 Folo / TopHub / 正式账本，2026-07-17 怎样复查。单站逐项理由继续标为 `needs_item_level_reason`，不由当前模型补编。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：48 条是候选队列，不是订阅事实

```yaml
- source_id: "浏览器订阅候选 48 条整体"
  current_route: resource_links_review_queue_not_subscription_ledger
  decision_status: recovered_from_file_group_level
  selected_reason: "文件明确 48 条媒体、平台和持续更新入口当前仍是 `media` / `platform` 类型链接，因此居留在 `资源/链接/`；只有实际进入 Folo 或 TopHub 后，才更新正式订阅账本。"
  rejected_alternatives:
    - "直接写入 Folo 正式订阅"
    - "直接写入 TopHub 正式订阅"
    - "把链接候选当已认可来源"
    - "把 48 条逐项标 completed"
  route_reason: "候选队列只是保留复查对象；不代表已订阅、已认可或已验证。"
  reuse_rule: "任何浏览器导出的持续更新入口，先留 `资源/链接/` 候选；只有通过复查并实际进入平台后才更新正式账本。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#状态; #复查边界"
  next_action: "2026-07-17 复查。"
```

---

## 2. 促销与行动触发：只在动作发生时有用

```yaml
- source_id: "IT之家「喜加一」 / 少数派商城"
  current_route: action_trigger_candidates_not_folo_feeds
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为促销与行动触发，而不是媒体、思想、财经或技术来源。"
  rejected_alternatives:
    - "加入 Folo 常驻"
    - "加入 TopHub 常驻"
    - "用促销流制造日常打开义务"
    - "把商城或领取入口当信息源"
  route_reason: "促销、领取、商城入口的价值依赖具体动作：领取、购买、比价、查权益。没有动作时，它们会制造冲动、噪声和待办感。"
  reuse_rule: "促销与行动触发入口只在明确动作前打开；任务结束退出，不进入持续未读流。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#促销与行动触发"
  next_action: "按任务触发。"
```

---

## 3. 中文综合与公共新闻：TopHub 已承担公共温度，首页不等于合格订阅源

```yaml
- source_id: "央视网 / 凤凰网 / 新浪网 / 腾讯网 / 搜狐 / 网易 / 澎湃新闻 / 环球网"
  current_route: public_news_candidates_check_tophub_overlap_first
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为中文综合与公共新闻，保留在候选队列，说明它们可能与公共温度或中文新闻入口有关。"
  rejected_alternatives:
    - "全部加入 TopHub"
    - "全部加入 Folo"
    - "用门户首页替代公共温度分组"
    - "把首页存在当可订阅质量证明"
  route_reason: "复查边界明确第一步是判断 TopHub 是否已经承担公共温度；综合门户和公共新闻首页容易与已有公共温度、报刊、机构媒体、综合平台热榜重复。"
  reuse_rule: "综合新闻候选先查现有公共温度是否覆盖；只有低重复、稳定、独特编辑价值成立，才考虑 TopHub 或 Folo。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#中文综合与公共新闻; #复查边界"
  next_action: "复查 TopHub 重复度。"
```

---

## 4. 中文财经与商业：已存在数据与结构，不用浏览器首页补流

```yaml
- source_id: "界面新闻 / 财新网 / 晚点 LatePost / 第一财经 / 华尔街日报中文网 / FT 中文网 / 36氪"
  current_route: finance_business_candidates_check_existing_data_structure_and_folo_fit
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为中文财经与商业，保留为候选，说明它们可能与数据与结构、财经深度或商业观察有关。"
  rejected_alternatives:
    - "把这些首页全部加入 Folo"
    - "用浏览器候选覆盖已经执行的财经替换"
    - "把付费墙或首页当合格订阅源"
    - "把商业媒体候选全部塞进数据与结构"
  route_reason: "PR #372 后续已通过财经 ledger 确认数据与结构中的财新首页、日经、FT 中文、第一财经等功能位。浏览器候选不能反向制造重复订阅；晚点、界面、WSJ 中文、36氪等要分别看重复度、付费墙、独特稳定阅读关系和 TopHub 是否已有承担。"
  reuse_rule: "财经商业候选先查现有数据与结构功能位；若是工具或付费墙首页，通常留按需；若是独特长期关系，再进入 Folo 复查。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#中文财经与商业; #复查边界; finance ledgers"
  next_action: "复查时逐项看重复度与阅读关系。"
```

---

## 5. 中文科技、产品与行业：科技雷达已很拥挤，候选不能自动扩容

```yaml
- source_id: "虎嗅网 / 少数派 / 数字尾巴 / 爱范儿 / 极客公园 / 品玩 / 人人都是产品经理 / IT之家 / DataLearner AI / XiaoHu.AI 日报 / MAIGOO 行业资讯"
  current_route: chinese_tech_product_candidates_check_radar_overlap_and_noise
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为中文科技、产品与行业，保留为候选，说明它们与科技雷达、产品观察、AI 日报或行业资讯有关。"
  rejected_alternatives:
    - "全部加入科技雷达"
    - "全部加入 Folo"
    - "把 AI 日报 / 行业资讯流当稳定知识源"
    - "用中文科技媒体首页替代具体科技雷达功能位"
  route_reason: "后续科技 ledger 已多次确立：科技雷达看公共发现和结构缺口，不做日报 / 快讯 / 产品资讯收件箱。少数派等曾有旧节点退出或替换，IT之家等适合具体任务或事件，AI 日报类容易重复模型发布与热点。"
  reuse_rule: "中文科技候选先问它补哪个科技功能位：项目发现、官方行动、企业 IT、安全基础设施、个人工具、消费电子、产品方法或 AI 研究。不能补功能位就不扩容。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#中文科技、产品与行业; technology ledgers"
  next_action: "复查时按功能位逐项判断。"
```

---

## 6. 国际新闻、财经与科技：英语媒体不是自动 Folo 关系

```yaml
- source_id: "TED Talks / CNBC / Bloomberg / WSJ / FT / The Guardian / BBC News / The Verge / TechCrunch / The Information / iMore / Product Hunt / Nine.com.au / Reader’s Digest / VentureBeat"
  current_route: international_media_candidates_check_role_paywall_and_overlap
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为国际新闻、财经与科技，保留为候选，说明它们可能补国际公共新闻、英语财经、科技创业、产品发现或文化材料。"
  rejected_alternatives:
    - "全部加入 Folo"
    - "用国际媒体数量弥补中文信息源不足"
    - "忽略付费墙、打开率和重复度"
    - "把 Product Hunt 等产品热榜当长期关系"
  route_reason: "复查边界要求 Folo 只考虑独特、稳定、真实会读的来源。国际媒体中有强来源，也有付费墙、泛新闻、创业热点、产品热榜和低相关文化材料；不能因为英文或国际就直接升级。"
  reuse_rule: "国际来源先分：公共新闻、财经、科技创业、产品发现、文化材料。Folo 只收长期阅读关系，TopHub 收公共雷达，产品热榜按需或任务触发。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#国际新闻、财经与科技; #复查边界"
  next_action: "复查时逐项看真实打开率与替代关系。"
```

---

## 7. 视觉、自然与杂志：视觉漫游不等于 Folo 长期关系

```yaml
- source_id: "National Geographic / MAGAZINELIB / 考拉新媒体导航"
  current_route: visual_nature_magazine_candidates_not_auto_subscription
  decision_status: recovered_from_file_group_level_needs_item_level_reason
  selected_reason: "这组被归为视觉、自然与杂志，说明它们可能补视觉、自然、杂志入口或媒体发现。"
  rejected_alternatives:
    - "全部加入 Folo"
    - "把杂志库当正式阅读关系"
    - "把导航站当内容源"
    - "用视觉漫游替代已有视觉与自然分组"
  route_reason: "视觉 / 自然来源要区分原始内容、图像策展、杂志库、导航站和版权 / 可访问性问题。导航站不是内容源，杂志库不等于阅读关系。已有视觉与自然分组已承担低频视觉漫游。"
  reuse_rule: "视觉来源先确认它是原始内容、策展、资料库还是导航；只有稳定原文或策展关系成立才进 Folo / TopHub。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#视觉、自然与杂志; visual-nature ledgers"
  next_action: "复查时分清来源类型。"
```

---

## 8. 财经日历：工具链接，不是阅读流

```yaml
- source_id: "华尔街见闻财经日历 / 东方财富财经日历"
  current_route: worang_tool_or_on_demand_calendar_not_feed
  decision_status: recovered_from_file_group_level
  selected_reason: "这组被单独归为财经日历。文件复查边界也明确财经日历不自动等同于合格订阅源。"
  rejected_alternatives:
    - "加入 Folo"
    - "加入 TopHub 常驻阅读流"
    - "作为每日未读内容"
    - "用财经日历生成投资动作"
  route_reason: "财经日历是查询工具，用于具体宏观、财报、数据公布、市场事件核对，不是持续阅读源。此前财经 ledger 已将华尔街见闻财经日历路由为沃壤工具链接。"
  reuse_rule: "财经日历写工具链接或按需调用，不制造未读，不生成投资动作。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#财经日历; finance ledger"
  next_action: "工具 / 按需。"
```

---

## 9. 2026-07-17 复查协议

```yaml
- source_id: "浏览器订阅候选复查协议"
  current_route: review_protocol_for_48_candidates
  decision_status: recovered_from_file_group_level
  selected_reason: "文件明确复查边界：先判断 TopHub 是否已经承担公共温度；Folo 只考虑独特、稳定、真实会读的来源；首页、付费墙、聚合站、财经日历和促销页不自动等同于合格订阅源；2026-07-17 先复查现有 27 个 Folo 来源，再按一进一出原则判断。"
  rejected_alternatives:
    - "复查前提前扩容 Folo"
    - "先加候选再说"
    - "忽略现有 27 项打开率"
    - "按候选数量补齐订阅"
  route_reason: "复查不是扩容许可，而是筛选机制。先处理现有负担，再考虑候选一进一出。"
  reuse_rule: "到复查日先看现有 Folo 27 项是否退出、合并或降级；新候选必须替代旧来源或证明独特关系。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#复查边界"
  next_action: "2026-07-17 复查。"
```

---

## 10. 本补充 ledger 的复用规则

1. 浏览器订阅候选 48 条是 `资源/链接/` 候选，不是订阅事实。
2. 只要文件没有逐项旧理由，单站保持 `needs_item_level_reason`。
3. 促销 / 行动触发按任务打开，不进 Folo。
4. 中文综合新闻先查 TopHub 公共温度重叠。
5. 中文财经商业先查数据与结构功能位、付费墙和真实阅读关系。
6. 中文科技产品行业先查科技雷达功能位，不做日报 / 快讯 / 产品资讯收件箱。
7. 国际媒体先分公共新闻、财经、科技创业、产品发现、文化材料；英文不等于自动 Folo。
8. 视觉自然杂志要区分原始内容、策展、资料库和导航站。
9. 财经日历是工具，不是阅读流。
10. 2026-07-17 复查先处理现有 Folo 27 项，再一进一出。

---

## 11. 仍需继续追回

浏览器订阅候选 48 条已完成组级证据追回。

继续待办：

- 单站逐项 why_in / why_out 仍需要更多旧对话原文或复查当天实际判断；
- 若继续 recent-chat source evidence，可尝试寻找 TopHub 早期 28 个取消项原始列表；
- 若找不到原始列表，则只能转入 workflow / content boundary ledger，不把 source evidence 硬编完成；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
