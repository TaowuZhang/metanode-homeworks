# PR #372 判断链追回：综合平台注意力源（微信 / 微博 / 今日头条）（2026-07-06）

## 0. 边界

本文件是 `综合` 大类的补充 ledger，追回 PR #372 中已经闭环的三个平台注意力小标签：

- `docs/ops/tophub-comprehensive-wechat-closure-20260704.md`
- `docs/ops/tophub-comprehensive-weibo-closure-20260704.md`
- `docs/ops/tophub-comprehensive-toutiao-closure-20260704.md`

本文件不处理用户真实关注的微信公众号，也不处理所有综合门户。它只处理 TopHub 综合目录中这三组平台榜单节点。

判断链标准：

> 每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 平台注意力源的总规则

```yaml
- source_id: "综合平台注意力源总规则"
  current_route: platform_attention_not_original_truth_source
  decision_status: recovered_from_pr
  selected_reason: "微博、微信、今日头条等平台榜单可以显影平台内部的注意力结构、标题传播、情绪表达、平台推荐和内容工业模式。"
  rejected_alternatives:
    - "把所有平台切片都当作独立来源"
    - "把平台榜单放入 Folo"
    - "把平台榜单当事实来源或质量来源"
    - "用平台分类榜替用户真实关注的作者 / 公众号 / 媒体做判断"
  route_reason: "平台榜单主要留在 TopHub 的公共温度或按需层，用于横向观察；Folo 保存连续作者、编辑刊物或原始来源，不保存平台热度切片。"
  reuse_rule: "以后遇到平台榜单，先问它是否提供一种现有系统没有的注意力结构；如果只是同一平台内容池的分类切片、营销榜、搜索建议或标题流，就按需或拒绝，不常驻、不进 Folo、不建宽泛追踪器。"
  evidence_locator: "tophub-comprehensive-wechat-closure-20260704.md; tophub-comprehensive-weibo-closure-20260704.md; tophub-comprehensive-toutiao-closure-20260704.md"
  next_action: "无；作为复用规则。"
```

---

## 2. 微信小标签

来源证据：`docs/ops/tophub-comprehensive-wechat-closure-20260704.md`

`微信` 小标签共 34 个节点，三页全部查看。

结论：

- 保留 `微信｜24h热文榜`；
- 保留 `微信读书｜新书榜`；
- 不新增微信读书其余 8 个榜单；
- 不新增 24 个微信垂直分类榜；
- 不进入 Folo；
- 不建立微信相关宽泛追踪器；
- 不处理用户真实关注的微信公众号。

```yaml
- source_id: "微信｜24h热文榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "不是高质量内容榜，也不是公众号推荐榜，而是微信公众号生态中高传播文章的混合样本。它最有价值的是暴露公众号生态正在传播的标题、情绪、地区信息和内容工业模式。"
  rejected_alternatives:
    - "把它当高质量内容榜"
    - "把 10 万+ 当作内容质量"
    - "从榜单直接关注公众号"
    - "用它替代真实公众号整理"
  route_reason: "放在 `公共温度`，只作为公众号传播生态样本；公共温度的任务是比较平台差异，而不是提供低噪声事实。"
  reuse_rule: "以后遇到高传播公众号榜，只把它当传播样本，不当事实源、质量榜或公众号整理依据；健康、投资、法律、教育、天气、政策一律回原始来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#一微信24h热文榜"
  next_action: "无；保留但降格使用。"

- source_id: "微信读书｜新书榜"
  current_route: tophub_slow_read_and_thought
  decision_status: recovered_from_pr
  selected_reason: "在九个微信读书节点里，它最能提供微信读书最近新增和推动什么的变化信号；可作为轻量新书发现入口。"
  rejected_alternatives:
    - "微信读书｜总榜"
    - "微信读书｜飙升榜"
    - "微信读书｜热搜榜"
    - "微信读书｜男生小说榜 / 女生小说榜 / 小说榜"
    - "微信读书｜神作榜 / 神作潜力榜"
  route_reason: "放在 `慢读与思想`，只用于发现书名，不根据推荐值直接判断质量；后续仍用豆瓣书评、新书速递、作者与出版社信息二次判断。"
  reuse_rule: "平台书榜只保留能提供近期变化信号的一项；长期总榜、营销飙升、热搜、类型小说榜和平台命名榜不叠加。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#三微信读书新书榜"
  next_action: "无；保留。"

- source_id: "微信垂直分类榜 24 项"
  current_route: rejected_misclassified_platform_categories
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "微信·军事"
    - "微信·科技"
    - "微信·文化"
    - "微信·生活"
    - "微信·职场"
    - "微信·财经"
    - "微信·教育"
    - "微信·历史"
    - "微信·健康"
    - "微信·搞笑"
    - "微信·娱乐"
    - "微信·社会"
    - "微信·汽车"
    - "微信·美食"
    - "微信·房产"
    - "微信·旅行"
    - "微信·时尚"
    - "微信·宠物"
    - "微信·情感"
    - "微信·游戏"
    - "微信·育儿"
    - "微信·体育"
    - "微信·星座"
    - "微信·动漫"
  route_reason: "分类标签发生系统性错位，绝大多数垂直节点无法按标签解释；这些榜单只能证明微信文章分类聚合当前失真，不能提供可靠垂直信息。"
  reuse_rule: "分类名不能被信任时，不得按标签把来源放入科技、财经、健康、教育、娱乐或生活分组。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#二24个微信垂直分类榜"
  next_action: "无；全部不新增。"

- source_id: "微信读书｜总榜"
  current_route: on_demand_long_term_popularity
  decision_status: recovered_from_pr
  selected_reason: "可按需查看微信读书长期大众偏好。"
  rejected_alternatives:
    - "慢读与思想常驻"
  route_reason: "总榜主要是长期畅销和经典作品，反映长期大众偏好，不是近期变化；新增会让经典畅销书长期占据首页，变化率过低。"
  reuse_rule: "长期总榜变化率低，适合按需，不适合常驻发现。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#四微信读书总榜"
  next_action: "按需。"

- source_id: "微信读书｜飙升榜 / 热搜榜"
  current_route: on_demand_platform_marketing_signal
  decision_status: recovered_from_pr
  selected_reason: "可偶尔观察平台营销、影视联动、活动和搜索变化。"
  rejected_alternatives:
    - "慢读与思想常驻"
  route_reason: "飙升和热搜与平台推荐、营销、影视联动、活动、搜索高度相关，不能形成独立阅读判断。"
  reuse_rule: "平台推动信号不等于质量；只在追问某本书为何走红时按需打开。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#五微信读书飙升榜; #六微信读书热搜榜"
  next_action: "按需。"

- source_id: "微信读书｜男生小说榜 / 女生小说榜 / 小说榜"
  current_route: rejected_type_fiction_homepage_expansion
  decision_status: recovered_from_pr
  selected_reason: "不否定其中作品价值。"
  rejected_alternatives:
    - "慢读与思想常驻"
  route_reason: "这些是网文与类型小说分类榜，男/女分类粗糙，影视原著会强烈推动排名，三个榜单之间大量重复，会扩大网络小说在慢读与思想中的权重。"
  reuse_rule: "类型小说榜按具体类型和作品寻找，不把平台性别化分类热榜加入首页。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#七微信读书男生小说榜女生小说榜小说榜"
  next_action: "无；不新增。"

- source_id: "微信读书｜神作榜 / 神作潜力榜"
  current_route: rejected_platform_marketing_label
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "作为文学评价来源"
    - "作为新书判断来源"
  route_reason: "榜名带强烈平台营销语言，混合经典文学、网文、影视原著、股票书、AI通俗内容、自我成长、赚钱书和命理教材，推荐值异常高，不能承担文学评价或新书判断。"
  reuse_rule: "平台命名为‘神作’或类似营销标签，不构成评价标准。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#八微信读书神作榜神作潜力榜"
  next_action: "无；明确不新增。"

- source_id: "真实微信公众号整理"
  current_route: out_of_scope_for_this_ledger
  decision_status: recovered_from_pr
  selected_reason: "本次只审 TopHub 的微信榜单节点。"
  rejected_alternatives:
    - "根据失真的微信分类榜批量判断用户真实关注的公众号"
  route_reason: "用户当前关注的公众号哪些保留、哪些迁入 Folo、哪些留在微信、是否建立公众号分组、如何处理历史关注与未读压力，都要等七分组与 Folo 基线复查后单独处理。"
  reuse_rule: "平台榜单判断不能替代真实订阅关系判断；真实作者 / 公众号另开。"
  evidence_locator: "docs/ops/tophub-comprehensive-wechat-closure-20260704.md#十一与真实公众号整理的边界"
  next_action: "后续单独处理。"
```

---

## 3. 微博小标签

来源证据：`docs/ops/tophub-comprehensive-weibo-closure-20260704.md`

`微博` 小标签共 5 个节点，全部查看。

结论：

- 保留 `微博｜热搜榜`；
- 保留 `微博｜话题榜`；
- 不新增 `微博｜社会榜`、`微博｜文娱榜`、`微博｜生活榜`；
- 不进入 Folo；
- 不建立微博相关宽泛追踪器。

```yaml
- source_id: "微博｜热搜榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "让微博平台的即时注意力结构显影：赛事、娱乐、社会、科技、商业、健康议题混合，同一事件被拆成大量人物、争议、梗和情绪条目。"
  rejected_alternatives:
    - "微博｜社会榜作为替代或叠加"
    - "微博｜文娱榜作为替代或叠加"
    - "微博｜生活榜作为替代或叠加"
    - "把微博热搜放入 Folo"
  route_reason: "继续放在 `公共温度`；它不是事实重要性排序，而是平台传播与互动结果。"
  reuse_rule: "看到同一事件多条时，把它理解为传播裂变，不重复阅读；健康、投资、政策、国际冲突和产品消息回原始来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#一微博热搜榜"
  next_action: "无；保留。"

- source_id: "微博｜话题榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "观察微博正在持续组织哪些关系、心理、工作和生活问题；它是长期讨论题与情绪 / 关系问答入口，不是实时热搜。"
  rejected_alternatives:
    - "微博｜生活榜"
    - "把话题榜当心理建议或关系判断来源"
  route_reason: "继续放在 `公共温度`，与热搜榜互补：一个显影即时注意力，一个显影持续讨论结构。"
  reuse_rule: "话题榜只能带走问题，不能把榜单答案当结论；真正触发思考时再进入正式来源或后续讨论。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#二微博话题榜"
  next_action: "无；保留。"

- source_id: "微博｜社会榜"
  current_route: on_demand_social_slice
  decision_status: recovered_from_pr
  selected_reason: "当微博总热搜被赛事或娱乐完全占据，而只想观察微博上的社会议题时，可临时打开。"
  rejected_alternatives:
    - "公共温度常驻"
    - "数据与结构常驻"
  route_reason: "它不是稳定社会议题编辑页，而是微博热搜的社会化切片；与微博热搜、澎湃、新京报、中国新闻周刊等高度重叠，且仍受微博热度机制支配，不提供更可靠事实来源。"
  reuse_rule: "同一平台的垂直切片只提高某类内容密度，不等于新来源关系；按需，不常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#三微博社会榜"
  next_action: "按需。"

- source_id: "微博｜文娱榜"
  current_route: on_demand_entertainment_propagation_slice
  decision_status: recovered_from_pr
  selected_reason: "可在具体明星、剧集、演唱会或粉圈事件成为研究对象时临时观察传播。"
  rejected_alternatives:
    - "影游音乐常驻"
  route_reason: "它反映明星与粉圈的即时传播，不是作品发现、影评、音乐评论或游戏文化。加入文娱榜只会把明星与粉圈热度带进作品分组，破坏分组用途。"
  reuse_rule: "明星 / 粉圈传播不等于作品入口；影游音乐只收作品、媒介和文化观察。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#四微博文娱榜"
  next_action: "按需。"

- source_id: "微博｜生活榜"
  current_route: on_demand_life_expression_slice
  decision_status: recovered_from_pr
  selected_reason: "可在需要观察微博上的打工人话语、生活梗、情绪表达和轻趋势时临时打开。"
  rejected_alternatives:
    - "生活与社区常驻"
  route_reason: "它更接近微博生活方式、情绪表达与网络梗切片，不是可执行、可验证的生活信息；健康内容也不能直接采用。"
  reuse_rule: "生活表达和生活行动分开；能观察话语不等于能提供行动建议。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#五微博生活榜"
  next_action: "按需。"

- source_id: "微博平台节点整体"
  current_route: keep_two_nodes_only
  decision_status: recovered_from_pr
  selected_reason: "两个现有节点已经分别覆盖即时平台注意力与持续性讨论题。"
  rejected_alternatives:
    - "社会榜"
    - "文娱榜"
    - "生活榜"
    - "微博相关 Folo 来源"
    - "微博宽泛追踪器"
  route_reason: "后三者能提高某类内容密度，但没有形成新的来源关系、编辑能力或证据能力；再加只会提高微博权重并重复同一内容池。"
  reuse_rule: "同一平台保留总入口和互补入口即可，垂直切片只有按需价值。"
  evidence_locator: "docs/ops/tophub-comprehensive-weibo-closure-20260704.md#核心结论; #六与七分组的关系; #七tophubfolo与追踪器; #最终结论"
  next_action: "无；微博已闭环。"
```

---

## 4. 今日头条小标签

来源证据：`docs/ops/tophub-comprehensive-toutiao-closure-20260704.md`

`今日头条` 小标签共 14 个节点，两页全部查看。

结论：

- 不新增任何 TopHub 常驻节点；
- 14 个节点全部视为同一推荐与搜索内容池的主题切片；
- 不进入 Folo；
- 不建立今日头条相关宽泛追踪器。

```yaml
- source_id: "今日头条平台节点整体"
  current_route: rejected_or_on_demand_platform_slices
  decision_status: recovered_from_pr
  selected_reason: "如需比较今日头条平台当天注意力，可按需打开 `头条热榜｜总榜` 一个节点。"
  rejected_alternatives:
    - "今日头条｜头条热榜"
    - "头条搜索｜猜你想搜"
    - "头条热榜｜总榜"
    - "头条热榜｜财经榜"
    - "头条热榜｜健康榜"
    - "头条热榜｜国际榜"
    - "头条热榜｜军事榜"
    - "头条热榜｜台海榜"
    - "头条热榜｜科技榜"
    - "头条热榜｜娱乐榜"
    - "头条热榜｜文旅榜"
    - "头条热榜｜教育榜"
    - "头条热榜｜汽车榜"
    - "头条热榜｜体育榜"
  route_reason: "这 14 个节点不是 14 个独立来源，而是今日头条同一推荐与搜索内容池按主题重新切片；主题榜只改变标签，不改变来源质量。"
  reuse_rule: "同一推荐池的主题切片不当独立来源；最多按需打开一个总入口比较平台注意力。"
  evidence_locator: "docs/ops/tophub-comprehensive-toutiao-closure-20260704.md#核心判断; #tophubfolo与追踪器; #最终结论"
  next_action: "无；今日头条已闭环。"

- source_id: "今日头条｜头条热榜 / 头条热榜｜总榜"
  current_route: on_demand_platform_attention_check
  decision_status: recovered_from_pr
  selected_reason: "可按需观察今日头条平台推荐和消费强度。"
  rejected_alternatives:
    - "公共温度常驻"
  route_reason: "当前样本被单一世界杯事件高度占据，同一事件被拆成比分、门将、纪录、球迷、情绪和评论等大量条目；它反映平台推荐和消费强度，不是公共议题完整性，也不是事实可靠性排序。"
  reuse_rule: "总榜和主榜高度重复时，不并存；同一事件重复拆分说明它在放大平台消费，而非拓宽视野。"
  evidence_locator: "docs/ops/tophub-comprehensive-toutiao-closure-20260704.md#今日头条头条热榜; #头条热榜总榜"
  next_action: "按需。"

- source_id: "头条搜索｜猜你想搜"
  current_route: rejected_search_suggestion_not_source
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "公共温度常驻"
  route_reason: "它混合热点事件、车型、小说、字体设置、熊猫、人物和社会案件，是平台搜索联想，不是内容来源；只能说明部分用户可能继续搜索什么。"
  reuse_rule: "搜索联想不等于内容来源，也不等于公共榜单。"
  evidence_locator: "docs/ops/tophub-comprehensive-toutiao-closure-20260704.md#头条搜索猜你想搜"
  next_action: "无；不新增。"

- source_id: "头条热榜｜财经榜 / 健康榜 / 国际榜 / 军事榜 / 台海榜 / 科技榜 / 娱乐榜 / 文旅榜 / 教育榜 / 汽车榜 / 体育榜"
  current_route: rejected_vertical_platform_slices
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "数据与结构"
    - "生活与社区"
    - "影游音乐"
    - "科技雷达"
    - "公共温度"
  route_reason: "这些垂直榜单混合事实新闻、媒体评论、自媒体观点、情绪表达和搜索联想；财经、健康、国际、军事、台海、科技、娱乐、文旅、教育、汽车、体育等标签不能替代对应领域的原始来源、官方文件、专业媒体或任务入口。"
  reuse_rule: "平台垂直榜只改变标签，不改变证据等级；领域问题回领域源，不让平台热榜进入领域分组。"
  evidence_locator: "docs/ops/tophub-comprehensive-toutiao-closure-20260704.md#各节点判断; #与七分组的关系"
  next_action: "无；不新增。"

- source_id: "今日头条 Folo / 追踪器"
  current_route: no_folo_no_tracker
  decision_status: recovered_from_pr
  selected_reason: "今日头条节点是高频推荐流和热度切片。"
  rejected_alternatives:
    - "今日头条 Folo 来源"
    - "今日头条"
    - "头条热榜"
    - "国际榜 / 军事榜 / 台海榜 / 财经榜 / 科技榜 / 健康榜等宽泛追踪"
  route_reason: "这些来源会制造重复、情绪化标题和大量未读，不适合建立持续来源关系；若事件值得持续关注，应追踪具体事件、政策、公司、人物或项目，而不是平台榜单。"
  reuse_rule: "高频推荐流不进 Folo；宽泛平台标签不进追踪器。"
  evidence_locator: "docs/ops/tophub-comprehensive-toutiao-closure-20260704.md#tophubfolo与追踪器"
  next_action: "无。"
```

---

## 5. 本补充 ledger 的复用规则

1. 平台榜单可以保留少量总入口，用来显影平台注意力结构。
2. 同一平台的垂直切片通常只提高某类内容密度，不产生新来源关系。
3. 领域标签不能替代证据等级；财经、健康、教育、科技、军事、汽车、文旅问题要回官方、专业或任务源。
4. 平台热榜不进 Folo；Folo 保存连续作者、编辑刊物或原始来源。
5. 宽泛平台词不进追踪器；追踪器只收具体事件、政策、产品、公司、人物或项目。
6. 平台榜单判断不能替代真实订阅关系判断；真实公众号、作者、媒体应另开处理。

---

## 6. 仍需继续追回

综合大类已经补入：

- closeout 三项动作；
- 知乎小标签；
- 热搜小标签；
- 滚动新闻小标签；
- 微信小标签；
- 微博小标签；
- 今日头条小标签。

仍需继续追回：

- 百度、腾讯、网易、搜狐、新浪、ZAKER 等门户同族 closure；
- 半月谈、央视、人民网、新华社、南方周末等同族 closure；
- 健康、军事、房产、地方城市等 closure；
- `腾讯研究院` 与 `网易人间｜记事` 的更完整理由。

综合大类仍是 `partial_recovered_from_pr`，不能写成全量完成。
