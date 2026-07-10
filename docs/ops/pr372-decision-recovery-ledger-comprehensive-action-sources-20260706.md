# PR #372 判断链追回：综合行动型来源（健康 / 房产 / 本地宝 / 腾讯日报）（2026-07-06）

## 0. 边界

本文件是 `综合` 大类的补充 ledger，追回 PR #372 中已经闭环的四组行动型或摘要型来源：

- `docs/ops/tophub-comprehensive-health-closure-20260704.md`
- `docs/ops/tophub-comprehensive-real-estate-closure-20260704.md`
- `docs/ops/tophub-comprehensive-local-city-review-20260704.md`
- `docs/ops/tophub-comprehensive-tencent-daily-closure-20260704.md`

判断链标准：

> 每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 行动型来源总规则

```yaml
- source_id: "综合行动型来源总规则"
  current_route: task_invoked_not_daily_subscription
  decision_status: recovered_from_pr
  selected_reason: "健康、住房、城市服务、日报摘要都可能有真实行动价值，但多数价值只在具体问题出现时成立。"
  rejected_alternatives:
    - "把行动入口变成每日阅读流"
    - "把平台健康榜、房产频道、城市服务索引、日报摘要放进 Folo"
    - "用宽泛标签建立追踪器"
  route_reason: "行动型来源优先作为 TopHub 按需入口或任务源；只有低噪声、低频、证据可追溯、能补出稳定功能的来源才进入常驻。"
  reuse_rule: "以后遇到行动型来源，先问：是否有具体任务、具体城市、具体政策、具体身体问题、具体办事对象；没有就不常驻。"
  evidence_locator: "health / real-estate / local-city / tencent-daily closure files"
  next_action: "作为复用规则。"
```

---

## 2. 健康小标签

来源证据：`docs/ops/tophub-comprehensive-health-closure-20260704.md`

`健康` 小标签共 13 个节点，完整覆盖。最终动作是两组替换：`wikiHow 中文｜首页推荐` → `科普中国｜今日辟谣文章`，`壹心理` → `半月谈｜健康`。

```yaml
- source_id: "科普中国｜今日辟谣文章"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "功能清楚：纠正常见生活、健康和科学误区；条目短，可按需进入，不要求连续阅读。"
  rejected_alternatives:
    - "wikiHow 中文｜首页推荐"
    - "微信健康榜"
    - "头条健康榜"
    - "ZAKER健康有术"
  route_reason: "进入生活与社区，承担具体命题核验；与 `科普中国网｜热点排行` 不重复，后者是综合科普热点，前者是错误认知核验。"
  reuse_rule: "健康和生活误区优先选择可核验、低噪声、短条目入口，不用平台热榜替代事实核验。"
  evidence_locator: "docs/ops/tophub-comprehensive-health-closure-20260704.md#健康标签最终新增建议"
  next_action: "已执行。"

- source_id: "半月谈｜健康"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "主题稳定在疾病风险、慢病、用药、体检、心理健康和公共健康；比平台健康热榜少猎奇和流量化内容。"
  rejected_alternatives:
    - "壹心理"
    - "人民网健康频道"
    - "网易健康频道"
    - "中新网健康·生活"
  route_reason: "进入生活与社区，替换壹心理。它不是纯心理来源，但能以更稳定的公共健康框架覆盖心理与身体问题。"
  reuse_rule: "健康常驻源要少营销、少猎奇、低噪声，并和辟谣来源互补；平台栏目和泛健康频道多按需。"
  evidence_locator: "docs/ops/tophub-comprehensive-health-closure-20260704.md#半月谈健康"
  next_action: "已执行。"

- source_id: "ScienceDaily｜Health & Medicine News"
  current_route: on_demand_biomedical_research_news
  decision_status: recovered_from_pr
  selected_reason: "可作为具体生物医学研究问题的英文新闻发现源。"
  rejected_alternatives:
    - "生活与社区常驻"
    - "科技雷达常驻"
    - "Folo 常驻或观察期"
  route_reason: "它是高频英文医学研究新闻，不是个人健康指导源；标题新闻化，很多条目可能只是早期、单项或机制性研究，不能直接转化为诊疗、饮食或补充剂建议。"
  reuse_rule: "医学研究新闻只能引出原论文、研究机构或权威指南核验，不直接转化为个人行动。"
  evidence_locator: "docs/ops/tophub-comprehensive-health-closure-20260704.md#ScienceDaily的实际角色"
  next_action: "按具体研究问题使用。"

- source_id: "健康界 / 21财经新健康 / 人民网健康 / 网易健康等"
  current_route: on_demand_health_industry_or_expert_sources
  decision_status: recovered_from_pr
  selected_reason: "健康界适合医院管理、医保、器械和医疗产业；21财经新健康适合创新药、医药公司、融资和资本市场；人民网健康与网易健康可按任务查看。"
  rejected_alternatives:
    - "Folo"
    - "生活与社区常驻"
  route_reason: "这些候选要么是行业栏目，要么边界过宽，要么转载比例较高；不适合制造健康未读流。"
  reuse_rule: "健康产业、医院管理、药企资本市场和专家栏目按任务使用；持续订阅必须是低频、可追溯、稳定且真正阅读。"
  evidence_locator: "docs/ops/tophub-comprehensive-health-closure-20260704.md#其余节点路由"
  next_action: "按需。"
```

---

## 3. 房产小标签

来源证据：`docs/ops/tophub-comprehensive-real-estate-closure-20260704.md`

`房产` 小标签共 5 个节点，完整展示。结论：不新增常驻房产节点，不进 Folo，不建宽泛追踪。

```yaml
- source_id: "央广网｜房产"
  current_route: on_demand_housing_policy_source
  decision_status: recovered_from_pr
  selected_reason: "本页中最有实际信息增量，覆盖公积金、保障房、城市更新、租房、老旧小区、物业、住房政策、统计局房价和开发投资数据。"
  rejected_alternatives:
    - "数据与结构常驻"
    - "生活与社区常驻"
    - "Folo"
  route_reason: "它仍混合土拍、房企拿地、楼市情绪、商业地产展览和项目宣传；住房民生、土地交易和行业宣传没有清晰分层。"
  reuse_rule: "住房议题按具体问题进入：公积金、保障房、租房、城市更新、住建部门或统计数据；不订阅整个行业频道。"
  evidence_locator: "docs/ops/tophub-comprehensive-real-estate-closure-20260704.md#央广网房产"
  next_action: "住房政策或城市更新任务时按需。"

- source_id: "房产小标签其余节点"
  current_route: rejected_or_on_demand_real_estate_noise
  decision_status: recovered_from_pr
  selected_reason: "少数节点可在具体北京楼市、房企或市场任务中按需查看。"
  rejected_alternatives:
    - "微信房产榜"
    - "人民网房产频道"
    - "中华网地产财经频道"
    - "新浪财经房产频道"
  route_reason: "微信房产标签失真；人民网房产大量陈旧和企业稿；中华网更像北京交易与房企资讯流；新浪财经房产标签污染严重。"
  reuse_rule: "住房重要不等于房产频道可常驻；优先回国家统计局、住建、公积金、住房保障部门和具体城市公告。"
  evidence_locator: "docs/ops/tophub-comprehensive-real-estate-closure-20260704.md#各节点判断; #最终结论"
  next_action: "按需或排除。"
```

---

## 4. 本地宝小标签

来源证据：`docs/ops/tophub-comprehensive-local-city-review-20260704.md`

`本地宝` 小标签共 32 个城市节点，三页全部查看。结论：不新增任何城市节点，不进 Folo。

```yaml
- source_id: "本地宝 32 个城市节点"
  current_route: on_demand_city_service_index
  decision_status: recovered_from_pr
  selected_reason: "本地宝的真实价值是按城市提供行动信息：考试、招生、落户、公积金、补贴、限行、交通、住房、招聘、展览、演出、消费券和本地活动。"
  rejected_alternatives:
    - "多个城市一起加入生活与社区"
    - "因为城市知名度或以后可能旅行而订阅"
    - "Folo"
  route_reason: "这些节点高度依赖具体地域。对某城市没有居住、家人、工作、学习、办事、频繁往返或长期停留关系时，几乎全是噪声。"
  reuse_rule: "城市服务源只在现实城市关系成立时使用；没有明确城市任务，不常驻、不进 Folo。"
  evidence_locator: "docs/ops/tophub-comprehensive-local-city-review-20260704.md#TopHub与Folo路由判断; #最终决定"
  next_action: "出现明确城市任务时，从本地宝进入具体官方页面。"

- source_id: "本地宝城市节点进入 TopHub 的条件"
  current_route: conditional_tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "城市节点只有在用户与城市存在持续现实关系时才可能成为行动型入口。"
  rejected_alternatives:
    - "全国城市焦点资讯通用订阅"
  route_reason: "判断条件包括：当前居住地、家人长期所在地、正在办理落户 / 公积金 / 教育 / 住房 / 考试 / 求职、高频往返、阶段性长期停留。"
  reuse_rule: "城市类来源先确认现实关系，再决定是否临时加入；优先选择人社、公积金、交通、教育考试院、住建或政府公告等具体官方来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-local-city-review-20260704.md#TopHub与Folo路由判断"
  next_action: "按城市任务。"
```

---

## 5. 腾讯日报小标签

来源证据：`docs/ops/tophub-comprehensive-tencent-daily-closure-20260704.md`

`腾讯日报` 小标签共 11 个节点，完整展示。结论：不新增任何节点，不进 Folo，不建追踪器。

```yaml
- source_id: "腾讯日报 11 个节点"
  current_route: on_demand_tencent_digest_only
  decision_status: recovered_from_pr
  selected_reason: "如果想快速看腾讯如何编排当天议程，可按需打开综合早报或综合晚报其中一个。"
  rejected_alternatives:
    - "综合早报 / 综合晚报常驻"
    - "财经早报 / 财经晚报常驻"
    - "国际早报 / 国际晚报常驻"
    - "科技早报 / 汽车早报 / 体育早晚报 / 游戏早报常驻"
    - "Folo"
  route_reason: "11 个节点不是 11 个独立编辑来源，而是腾讯同一内容池按早晚和主题重新切片；跨节点重复明显。现有公共温度、数据与结构、科技雷达、影游音乐已有更清楚来源。"
  reuse_rule: "日报摘要产品按需看一个总入口即可；不要把同一内容池按早晚和主题切片拆成多个常驻源。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-daily-closure-20260704.md#核心结构判断; #最终结论"
  next_action: "无；腾讯日报已闭环。"

- source_id: "腾讯日报追踪器"
  current_route: no_tracker
  decision_status: recovered_from_pr
  selected_reason: "固定摘要产品不是需要条件触发的对象。"
  rejected_alternatives:
    - "腾讯早报"
    - "腾讯晚报"
  route_reason: "具体事件仍使用精确关键词、官方来源或后续跟踪，而不是追踪摘要品牌。"
  reuse_rule: "摘要产品不建追踪器；追踪器只跟具体对象。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-daily-closure-20260704.md#追踪器"
  next_action: "无。"
```

---

## 6. 本补充 ledger 的复用规则

1. 行动信息源只在具体任务成立时有价值，不自动变成持续阅读关系。
2. 健康来源必须区分：误区核验、公共健康解释、医学研究新闻、医疗产业信息、个人诊疗建议。不同层不能互相替代。
3. 房产和住房问题优先回具体城市、统计数据、住建、公积金和保障房官方来源，不订阅泛房产流。
4. 城市服务源只在现实城市关系成立时使用；没有城市关系就是噪声。
5. 日报 / 早报 / 晚报摘要产品如果来自同一内容池，只按需看一个总入口，不多节点订阅。
6. 行动型来源不进 Folo，除非它是低频、稳定、可追溯、真实持续阅读的具体来源。
7. 宽泛健康、房产、城市、早报、晚报词不进追踪器；追踪具体政策、城市、项目、办事对象或事件。

---

## 7. 仍需继续追回

综合大类已经补入：

- closeout 三项动作；
- 知乎；
- 热搜 / 滚动新闻；
- 微信 / 微博 / 今日头条；
- 百度 / 腾讯 / 网易 / 搜狐 / 新浪 / ZAKER；
- 人民网 / 新华社 / 南方周末 / 半月谈 / CCTV；
- 健康 / 房产 / 本地宝 / 腾讯日报。

仍需继续追回：

- 综合目录中尚未落账的 directory review / rolling review 边界；
- `腾讯研究院` 与 `网易人间｜记事` 的 Folo 复查条件可继续细化；
- 如果仍有安全层不适合展开的敏感类别，应单独低展开度记录边界，不在总控中铺开完整节点名。

综合大类仍是 `partial_recovered_from_pr`，不能写成全量完成。
