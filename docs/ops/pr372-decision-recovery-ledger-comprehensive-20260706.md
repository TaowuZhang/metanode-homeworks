# PR #372 判断链追回：综合大类（第一版，2026-07-06）

## 0. 边界

本文件追回 `TopHub 综合` 大类中已经能从 PR #372 文件直接确认的判断链。

已读取并纳入本版：

- `docs/ops/tophub-comprehensive-closeout-20260704.md`
- `docs/ops/tophub-comprehensive-zhihu-closure-20260704.md`

本版尚未覆盖全部综合同族文件。以下文件仍需继续追回：

- `tophub-comprehensive-directory-review-20260704.md`
- `tophub-comprehensive-hot-search-closure-20260704.md`
- `tophub-comprehensive-rolling-news-review-20260704.md`
- `tophub-comprehensive-rolling-news-closure-20260704.md`
- `tophub-comprehensive-banyuetan-review-20260704.md`
- `tophub-comprehensive-banyuetan-closure-20260704.md`
- `tophub-comprehensive-baidu-closure-20260704.md`
- `tophub-comprehensive-cctv-closure-20260704.md`
- `tophub-comprehensive-health-closure-20260704.md`
- `tophub-comprehensive-local-city-review-20260704.md`
- `tophub-comprehensive-military-closure-20260704.md`
- `tophub-comprehensive-netease-closure-20260704.md`
- `tophub-comprehensive-people-closure-20260704.md`
- `tophub-comprehensive-real-estate-closure-20260704.md`
- `tophub-comprehensive-sina-closure-20260704.md`
- `tophub-comprehensive-sohu-closure-20260704.md`
- `tophub-comprehensive-southern-weekly-closure-20260704.md`
- `tophub-comprehensive-tencent-closure-20260704.md`
- `tophub-comprehensive-tencent-daily-closure-20260704.md`
- `tophub-comprehensive-toutiao-closure-20260704.md`
- `tophub-comprehensive-wechat-closure-20260704.md`
- `tophub-comprehensive-weibo-closure-20260704.md`
- `tophub-comprehensive-xinhua-closure-20260704.md`
- `tophub-comprehensive-zaker-closure-20260704.md`

因此本文件状态是：`partial_recovered_from_pr`。

---

## 1. 综合大类最终动作

来源证据：`docs/ops/tophub-comprehensive-closeout-20260704.md`

综合大类已执行：

- 两项替换；
- 一项扩容新增；
- Folo 新增 0；
- 追踪器新增 0；
- 通知机器人保持关闭。

```yaml
- source_id: "wikiHow 中文｜首页推荐"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原本承担日常技能 / 生活建议类入口。"
  rejected_alternatives:
    - "继续放在生活与社区"
  route_reason: "当前首页偏向测试、关系、自助与泛化建议，不能稳定承担日常技能和事实核验。由 `科普中国｜今日辟谣文章` 替换。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#1-wikihow-中文首页推荐--科普中国今日辟谣文章"
  next_action: "无；已替换并从最终订阅列表移除。"

- source_id: "科普中国｜今日辟谣文章"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "承担明确的具体命题核验，适合处理生活、健康与科学误区；条目短，适合 TopHub 按需发现。"
  rejected_alternatives:
    - "wikiHow 中文｜首页推荐"
    - "把综合科普热点与错误认知核验合并成同一个节点"
  route_reason: "进入 `生活与社区`，与既有 `科普中国网｜热点排行` 不重复：一个是综合科普热点，一个是错误认知核验。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#1-wikihow-中文首页推荐--科普中国今日辟谣文章"
  next_action: "无；已订阅并加入生活与社区。"

- source_id: "壹心理"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原本承担心理 / 健康类入口。"
  rejected_alternatives:
    - "继续放在生活与社区"
  route_reason: "当前混合心理科普、咨询导流、人格测试、营销与猎奇标题，角色不够稳定。由 `半月谈｜健康` 替换。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#2-壹心理--半月谈健康"
  next_action: "无；已替换并从最终订阅列表移除。"

- source_id: "半月谈｜健康"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "稳定覆盖疾病风险、慢病、用药、体检、心理健康与公共健康；比平台健康热榜更少猎奇和流量化。"
  rejected_alternatives:
    - "壹心理"
    - "平台健康热榜"
  route_reason: "进入 `生活与社区`，与 `科普中国｜今日辟谣文章` 互补：一个核验真假，一个解释风险与健康问题。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#2-壹心理--半月谈健康"
  next_action: "无；已订阅并加入生活与社区。"

- source_id: "中央电视台｜新闻调查"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "周频、阅读压力低，以现场、人物和地区为基础，同时呈现政策、执行、行业、社区与普通人的关系。"
  rejected_alternatives:
    - "央视快讯类节点"
    - "只增加更多新闻标题流"
  route_reason: "进入 `数据与结构`，位于 `中央电视台｜每周质量报告` 后；补的是社会结构的实地调查，不是重复增加央视快讯。"
  ordering_reason: "国家统计局数据 / 财新 / FT / 经济半小时 / 每周质量报告之后，作为社会结构实地调查补位。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#3-新增-中央电视台新闻调查"
  next_action: "无；已订阅并加入数据与结构。"
```

---

## 2. 综合目录为什么只产生这三个动作

来源证据：`docs/ops/tophub-comprehensive-closeout-20260704.md#三为什么综合目录只产生这三个动作`

```yaml
- source_id: "综合目录其余节点"
  current_route: rejected_or_on_demand
  decision_status: recovered_from_pr_partial
  selected_reason: "部分节点仍可能在具体任务中有查询价值。"
  rejected_alternatives:
    - "同一门户内容池的热榜、频道和早晚报切片"
    - "同一事件被拆成大量近义标题"
    - "标签与内容不符的栏目"
    - "页面陈旧、接口停更或缓存失真的节点"
    - "评论、财经、科技、健康、军事等被其他内容污染的栏目"
    - "转载聚合层"
  route_reason: "综合目录没有因为规模大而继续扩张；新增门槛是必须补出当前七分组尚未具备的稳定功能。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#三为什么综合目录只产生这三个动作"
  next_action: "继续回综合同族 closure 文件补逐项理由。"
```

---

## 3. Folo 分工与综合候选

来源证据：`docs/ops/tophub-comprehensive-closeout-20260704.md#四folo本轮新增-0`

```yaml
- source_id: "综合目录 Folo 新增"
  current_route: no_new_folo_source
  decision_status: recovered_from_pr
  selected_reason: "本轮三项 TopHub 来源都适合标题发现与按需进入。"
  rejected_alternatives:
    - "把科普中国辟谣、半月谈健康、央视新闻调查同步加入 Folo"
  route_reason: "这些来源不适合制造持续未读；Folo 维持 27 项观察基线到 2026-07-17。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#四folo本轮新增-0"
  next_action: "无；保持 2026-07-17 复查。"

- source_id: "晚点 LatePost"
  current_route: existing_folo_source_no_duplicate
  decision_status: recovered_from_pr
  selected_reason: "知乎想法热榜暴露了晚点 LatePost 的内容线索。"
  rejected_alternatives:
    - "知乎想法热榜"
    - "晚点重复入口"
  route_reason: "Folo 已有 `晚点 - 长报道`，因此不再添加知乎想法热榜或晚点重复入口。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#不重复添加晚点"
  next_action: "无；避免重复入口。"

- source_id: "腾讯研究院"
  current_route: folo_low_frequency_deep_read_candidate
  decision_status: recovered_from_pr_partial
  selected_reason: "保留为 Folo 低频深读候选，不立即订阅。"
  rejected_alternatives:
    - "立即加入 Folo"
  route_reason: "进入前必须确认是否存在稳定直接 RSS 或可靠官方订阅、是否需要不稳定 RSSHub、是否持续更新、是否真的被打开阅读、是否与现有 AI / 社会 / 文化来源重复。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#保留但不执行的候选-腾讯研究院"
  next_action: "2026-07-17 或后续 Folo 复查时再判断；需要进一步追回腾讯相关 closure。"

- source_id: "网易人间｜记事"
  current_route: conditional_folo_candidate
  decision_status: recovered_from_pr_partial
  selected_reason: "作为普通人第一人称长篇的条件候选，不立即订阅。"
  rejected_alternatives:
    - "好读 / 特写 / 人间 / 大国小民同时加入"
  route_reason: "只有在 2026-07-17 复查后确认 Folo 缺少普通人第一人称长篇时，才考虑单独试用；试用时要观察实际阅读率、与现有人物 / 调查来源重复，以及疾病、死亡、犯罪、家庭冲突等高负荷题材造成的情绪成本。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#保留但不执行的候选-网易人间记事"
  next_action: "待旧对话或网易同族 closure 继续追回。"
```

---

## 4. 追踪器与通知机器人

来源证据：`docs/ops/tophub-comprehensive-closeout-20260704.md#五追踪器本轮新增-0` 与 `#六通知机器人保持关闭`

```yaml
- source_id: "综合目录追踪器"
  current_route: no_new_tracker
  decision_status: recovered_from_pr
  selected_reason: "综合目录没有发现必须主动通知的新对象。"
  rejected_alternatives:
    - "人民网 / 新华社 / 央视 / 腾讯 / 网易 / 搜狐 / 百度 / 新浪 / 今日头条 / ZAKER"
    - "热搜 / 财经 / 健康 / 军事 / 台海 / 国际 / AI / 汽车 / 教育等宽泛词"
    - "栏目品牌或门户名称"
  route_reason: "只有具体、有限、可闭环的政策、产品、公司、项目、人物或事件，才值得进入追踪器。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#五追踪器本轮新增-0"
  next_action: "无。"

- source_id: "综合目录通知机器人"
  current_route: notification_bot_off
  decision_status: recovered_from_pr
  selected_reason: "当前通知机器人订阅列表为空、通知渠道为空。"
  rejected_alternatives:
    - "企业微信"
    - "钉钉"
    - "飞书"
    - "Telegram"
    - "Discord"
    - "自动刷新"
  route_reason: "真正需要提醒的事项应进入日历或专门的条件监测，而不是把信息流推送到聊天工具。"
  evidence_locator: "docs/ops/tophub-comprehensive-closeout-20260704.md#六通知机器人保持关闭"
  next_action: "无。"
```

---

## 5. 知乎小标签完整判断链

来源证据：`docs/ops/tophub-comprehensive-zhihu-closure-20260704.md`

知乎小标签共 8 个节点已一次性全部查看。本版对知乎可视为 `recovered_from_pr`。

```yaml
- source_id: "知乎｜热榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "不是简单新闻标题榜，而是把事件改写成‘如何评价’‘为何发生’‘哪些信息值得关注’等问题；显示公众正在要求怎样的解释，而不只是哪些词正在被搜索。"
  rejected_alternatives:
    - "知乎｜热搜"
    - "微博热搜"
    - "小红书热榜"
    - "抖音总榜"
  route_reason: "继续保留在 `公共温度`；它提供问题入口，但不是答案质量保证。医疗、法律、金融、政策与科学问题必须回正式来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#一知乎热榜"
  next_action: "无；只从热榜带走值得继续追问的问题，不在榜单层完成判断。"

- source_id: "知乎日报｜Today"
  current_route: tophub_slow_read_and_thought
  decision_status: recovered_from_pr
  selected_reason: "数量低、题材跨度大、不完全受当天单一热点支配，更像一组可以停下来读的解释题。"
  rejected_alternatives:
    - "把它作为第二个公共热榜"
  route_reason: "继续保留在 `慢读与思想`；它是低频精选入口，不形成每日清空义务，真正有价值的文章才进入沃壤或后续讨论。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#二知乎日报today"
  next_action: "无。"

- source_id: "知乎｜热搜"
  current_route: on_demand_only
  decision_status: recovered_from_pr
  selected_reason: "可在需要比较知乎用户主动搜索什么与平台热榜推荐什么时临时查看。"
  rejected_alternatives:
    - "进入公共温度"
  route_reason: "当前条目几乎是热榜事件的更短关键词版本，只显示搜索意图，不提供问题结构、讨论内容或来源判断；信息量低且与其他热搜重复。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#三知乎热搜"
  next_action: "无；按需。"

- source_id: "知乎｜想法热榜"
  current_route: rejected_with_original_source_lead
  decision_status: recovered_from_pr
  selected_reason: "暴露出 `晚点 LatePost` 这一原始来源线索。"
  rejected_alternatives:
    - "订阅知乎想法热榜"
    - "通过知乎转载层订阅晚点"
  route_reason: "当前前列几乎被晚点 LatePost 占据，来源集中度高，短想法多为文章预告或摘录，夹杂写作素材、营销、财经荐股和日常打卡。若晚点值得长期阅读，应评估原始网站、公众号或可订阅源，而不是订阅知乎想法热榜。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#四知乎想法热榜"
  next_action: "晚点已在 Folo 有长报道入口；不重复添加。"

- source_id: "知乎｜话题榜"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "进入公共温度"
    - "作为话题发现入口"
  route_reason: "当前 14 条中大多重复同一个问题，节点名称与实际内容不一致，重复率极高，无法承担话题发现，也与热榜和进站必看没有互补价值。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#五知乎话题榜"
  next_action: "无；明确排除。"

- source_id: "知乎｜进站必看"
  current_route: on_demand_archive
  decision_status: recovered_from_pr
  selected_reason: "部分问题和回答具有长期阅读价值，适合回看知乎早期高质量问答或针对具体主题寻找经典回答。"
  rejected_alternatives:
    - "首页常驻"
    - "Folo 持续订阅"
  route_reason: "更像平台经典内容陈列和新用户导览，而不是持续更新来源；列表包含多年前热点和早期代表性问答。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#六知乎进站必看"
  next_action: "无；按需档案。"

- source_id: "知乎书店｜知乎周刊"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "进入慢读与思想"
    - "作为当前出版趋势或非虚构阅读来源"
  route_reason: "时间口径明显陈旧，混合平台专题、营销合作、考试材料和品牌内容；不是当前出版趋势、非虚构阅读或新书发现来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#七知乎书店知乎周刊"
  next_action: "无；明确不新增。"

- source_id: "知乎书店｜新书抢鲜"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "豆瓣热门图书"
    - "豆瓣新书速递"
    - "微信读书新书榜"
    - "豆瓣书评"
  route_reason: "页面没有清楚的编辑标准、出版时间、评价、主题分区或读者反馈，更多是知乎书店的上架与销售目录，不能替代现有书籍发现来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#八知乎书店新书抢鲜"
  next_action: "无；不进入慢读与思想。"
```

---

## 6. 知乎与 TopHub / Folo / 追踪器关系

```yaml
- source_id: "知乎平台节点整体"
  current_route: keep_existing_two_nodes_only
  decision_status: recovered_from_pr
  selected_reason: "知乎热榜与知乎日报各自提供清楚且互补的功能。"
  rejected_alternatives:
    - "知乎热搜"
    - "知乎想法热榜"
    - "知乎话题榜"
    - "知乎进站必看"
    - "知乎周刊"
    - "新书抢鲜"
    - "知乎平台节点进入 Folo"
    - "知乎相关宽泛追踪器"
  route_reason: "TopHub 维持保留 `知乎｜热榜` 与 `知乎日报｜Today`；Folo 应保存连续作者、媒体或原始来源，不保存知乎热搜、想法榜、话题榜、经典目录或商店目录；追踪器不建立知乎、AI、Agent、科技、健康、法律、教育等宽泛词。"
  evidence_locator: "docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#九与七分组的关系; #十tophubfolo与追踪器; #最终结论"
  next_action: "知乎已闭环；继续追综合其他同族文件。"
```

---

## 7. 本版仍需继续追回

本文件只完成综合大类中的：

- closeout 三项真实动作；
- Folo / 追踪器 / 通知机器人分工；
- 知乎小标签 8 个节点。

仍需继续追回：

- 搜索 / 热搜小标签；
- 微信 / 微博 / 公众号相关判断；
- 百度、腾讯、网易、搜狐、新浪、今日头条、ZAKER 等门户同族 closure；
- 半月谈、央视、人民网、新华社、南方周末等同族 closure；
- 健康、军事、房产、地方城市、滚动新闻等 closure；
- `腾讯研究院` 与 `网易人间｜记事` 的更完整旧对话或同族文件理由。

完成这些以后，才能把综合大类标为 `recovered_from_pr`，目前只能标为 `partial_recovered_from_pr`。
