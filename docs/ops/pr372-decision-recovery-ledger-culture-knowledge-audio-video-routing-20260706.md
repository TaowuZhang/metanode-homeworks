# PR #372 判断链追回：文化、知识内容与影音工作流路由（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `得到与知识内容接触谱系`、`文化内容入口与影音工作流`、`信息流最终分工`、`订阅系统 v0` 与 `豆瓣阅读偏好初步观察` 形成的判断链。

纳入文件：

- `资源/雷达/得到与知识内容接触谱系.md`
- `资源/雷达/文化内容入口与影音工作流.md`
- `资源/雷达/信息流最终分工.md`
- `资源/雷达/订阅系统.md`
- `资源/雷达/豆瓣阅读偏好初步观察.md`

边界说明：

- 本文件不是 TopHub 某个目录的逐页 closure。
- 本文件不是人物排行榜、教师排行榜、课程购买史、书影音完整消费史或平台迁移方案。
- 本文件追回的是：文化内容、知识内容、播客、视频、BibiGPT、得到、小宇宙、豆瓣、Folo、TopHub 与沃壤之间怎样分工。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：文化不是娱乐，知识也不是持续摄入

```yaml
- source_id: "文化内容与知识内容整体"
  current_route: multi_platform_routing_not_single_subscription_system
  decision_status: recovered_from_pr
  selected_reason: "用户当前文化内容生态同时包括公共文化回声与榜单发现、系统课程 / 听书 / 电子书、播客与长期节目关系、豆瓣书影音与评论社区、长期视频和音频节目，以及将选中的影音转成文稿后与 AI 讨论、校勘和沉积。"
  rejected_alternatives:
    - "把文化内容压缩成 TopHub 娱乐目录"
    - "把文化等同于持续获取知识"
    - "把得到、小宇宙、豆瓣、Folo、BibiGPT 全部统一镜像到 GitHub"
    - "把所有高质量人物、课程、节目或书都变成订阅源"
  route_reason: "这些入口承担不同职责，不需要合并成一个平台。TopHub 看公共回声，得到承接结构化课程 / 听书 / 电子书，小宇宙承接播客节目关系，豆瓣承接书影音条目与评论社区，Folo 承接少量长期关系，BibiGPT 承接取水与转录，沃壤只留下讨论后形成的判断、坐标、主题包和项目动作。"
  reuse_rule: "以后遇到文化或知识来源，先判断它是作品发现、评论解释、系统学习、节目关系、对话材料、水温观察、取水对象还是长期坐标。不能用一个平台吞掉所有形态。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#核心判断; 资源/雷达/信息流最终分工.md#一句话"
  next_action: "作为文化与知识内容总路由。"
```

---

## 2. TopHub：公共文化回声，不承担系统学习

```yaml
- source_id: "TopHub 文化入口"
  current_route: public_culture_echo_and_discovery
  decision_status: recovered_from_pr
  selected_reason: "TopHub 适合承担书、电影、音乐、游戏和视频的公共榜单，近期作品、话题与平台热度，跨领域漫游和意外发现，以及看见某个文化对象为什么正在被讨论。"
  rejected_alternatives:
    - "用 TopHub 承担系统学习"
    - "用 TopHub 替代完整课程、听书、电子书和播客逐期收听"
    - "用榜单决定单部作品或议题的最终判断"
    - "把 TopHub 作为未读清零系统"
  route_reason: "TopHub 负责公共注意力、平台温度、跨领域发现和候选发现。它不能替代得到、小宇宙、豆瓣、BibiGPT 或原平台的深消费。"
  reuse_rule: "TopHub 中的文化源只问：它是否补作品发现、公共讨论、审美漫游、平台温度或跨领域偶遇。若需要系统学习、完整观看/收听/阅读或长期关系，转回原平台或 Folo。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#TopHub; 资源/雷达/信息流最终分工.md#TopHub"
  next_action: "不据此新增或删除 TopHub 节点。"
```

---

## 3. 得到：结构化学习，不镜像课程购买史

```yaml
- source_id: "得到"
  current_route: native_platform_structured_learning_and_content_consumption
  decision_status: recovered_from_pr
  selected_reason: "用户主要在得到获取课程、听书、电子书和人文社科等结构化知识内容。得到承担相对完整、经过组织的知识消费。"
  rejected_alternatives:
    - "把得到上的全部课程、听书、电子书和收藏镜像进沃壤"
    - "把得到上的年更新闻或参考消息式专栏当主要新闻入口"
    - "把得到内容重复搬进 TopHub 或 Folo"
    - "把得到接触谱系写成人物排行榜"
  route_reason: "得到原平台明显优于聚合入口。真正改变长期理解的内容，才进入沃壤的领域、萃览、坐标或项目层；普通课程消费和听书历史不迁移。"
  reuse_rule: "得到材料按课程 / 听书 / 电子书 / 具体人物 / 具体问题返回。只有形成明确判断变化、领域理解、人物坐标或项目动作，才沉积到沃壤。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#得到; 资源/雷达/订阅系统.md#得到"
  next_action: "继续由得到原平台承担。"
```

---

## 4. 得到人物谱系：不是教师排行榜，而是关系强度与返回方式

```yaml
- source_id: "熊逸 / 王烁 / 吴军 / 吴伯凡 / 何帆 / 郑也夫 / 刘瑜 / 包刚升等"
  current_route: personal_knowledge_contact_spectrum_not_subscription_list
  decision_status: recovered_from_pr
  selected_reason: "这些人物在用户过去的阅读、课程或节目中形成过不同程度的影响：熊逸让古典文本和历史人物重新活起来；王烁提供认知训练和媒体判断方法；吴军曾是早期科技、工程、教育、方法和写作入口；吴伯凡是长期思想、商业与系统观察坐标；何帆承担年度经济观察；郑也夫、刘瑜、包刚升分别进入社会学、公共政治写作和政治学课程关系。"
  rejected_alternatives:
    - "把所有重要人物直接加入灯塔或 Folo"
    - "按得到名师知名度排序"
    - "把早期重要自动等同于当前每次输出都有增量"
    - "把听过课等同于长期追随"
  route_reason: "文件明确区分：较强长期影响、曾经重要但关系变化、按领域/课程/问题返回、只是听过、明确没有购买或不认同表达位置。真正留下的通常是框架、书、课程、具体单集和判断变化，而不是人物品牌。"
  reuse_rule: "人物源先判断关系类型：长期坐标、早期坐标、年度观察、按需课程入口、节目关系、水温入口、明确反例。只有反复返回并改变理解或选择的人，才进入坐标或灯塔候选。"
  evidence_locator: "资源/雷达/得到与知识内容接触谱系.md#第一层; #当前结构判断"
  next_action: "自然补具体材料，不强行盘点。"

- source_id: "梁宁 / 刘润"
  current_route: early_important_later_content_specific_review
  decision_status: recovered_from_pr
  selected_reason: "梁宁早期产品思维重要，刘润早期商学内容有价值。"
  rejected_alternatives:
    - "继续把早期重要视为长期无条件追随"
    - "全盘否定早期价值"
  route_reason: "两者代表‘曾经重要，但后来关系发生变化’。梁宁后期现实感和互联网现场质感减弱；刘润后期高频输出稀释独特性，结构化和整合能力强，但不等于篇篇有洞察。"
  reuse_rule: "早期重要不是永久免检；后续按具体内容复核，不按人物名气保留长期追更。"
  evidence_locator: "资源/雷达/得到与知识内容接触谱系.md#第二层"
  next_action: "按内容返回。"

- source_id: "薄世宁"
  current_route: important_medical_and_evidence_based_reasoning_coordinate_candidate
  decision_status: recovered_from_pr
  selected_reason: "他的急诊和医学视角改变了用户对医学与循证的理解：医学没有想象中神秘，治疗不是只靠权威解释，医学判断依赖证据、概率、经验、条件与持续修正。"
  rejected_alternatives:
    - "只记作普通医学按需入口"
    - "把循证理解简化为绝对答案"
  route_reason: "文件明确将薄世宁从普通医学按需入口提升为医学与循证思维的重要入口 / 个人健康认知坐标候选。"
  reuse_rule: "医学内容若改变用户对证据、概率和判断边界的理解，可以进入坐标候选；具体医疗行动仍回专业医生、指南、监管和现实条件。"
  evidence_locator: "资源/雷达/得到与知识内容接触谱系.md#薄世宁"
  next_action: "保留为坐标候选。"

- source_id: "万维钢"
  current_route: explicit_non_relationship_not_missing_item
  decision_status: recovered_from_pr
  selected_reason: "不是遗漏，而是明确没有形成购买关系，且用户不认同其常见表达位置。"
  rejected_alternatives:
    - "因知名度默认补入得到谱系"
    - "把偶尔点开等同于长期关系"
    - "把精英表达位置作为默认知识入口"
  route_reason: "用户自始至终没有购买其课程，只可能因具体问题偶尔点开一两节；关键不是内容完全没有价值，而是其表达常带有用户不认同的精英视角。"
  reuse_rule: "明确非关系也要记录，避免模型按知名度补编。没有购买、没有共鸣、表达位置不认同，不进入长期谱系。"
  evidence_locator: "资源/雷达/得到与知识内容接触谱系.md#明确的反例与非关系"
  next_action: "不补入长期谱系。"
```

---

## 5. 小宇宙：节目关系与收听，不等于 Folo 全量镜像

```yaml
- source_id: "小宇宙"
  current_route: native_platform_podcast_discovery_relationship_and_listening
  decision_status: recovered_from_pr
  selected_reason: "小宇宙用于发现节目与单集，建立对主持人、节目和系列的长期关系，直接收听，并从节目页、简介和社区反应理解节目上下文。"
  rejected_alternatives:
    - "把全部小宇宙节目搬入 Folo"
    - "把播客热榜等同于节目关系"
    - "把每个高价值单集都变成长期订阅"
  route_reason: "不是每个节目都需要进入 Folo。只有跨平台统一查看、避免推荐流、稳定 RSS 或不易漏更确实带来价值时，才考虑在 Folo 重复订阅。"
  reuse_rule: "播客先在小宇宙建立真实收听关系；Folo 只收少量长期节目、跨平台阅读室需求或 RSS 明显改善的节目。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#小宇宙; 资源/雷达/信息流最终分工.md#小宇宙"
  next_action: "继续由小宇宙原生承担。"
```

---

## 6. 豆瓣：条目、评论和社区关系，不被 TopHub 豆瓣榜替代

```yaml
- source_id: "豆瓣"
  current_route: native_platform_books_films_music_reviews_marks_and_communities
  decision_status: recovered_from_pr
  selected_reason: "豆瓣承担书、电影、音乐与创作者关系，书评、影评和条目资料，小组、话题与文化社区，以及公共口碑与作品发现。"
  rejected_alternatives:
    - "用 TopHub 豆瓣榜单替代豆瓣原生条目、评论和标记"
    - "把近期豆瓣镜像统计成完整阅读史"
    - "把猜书名游戏中 AI 出题书统一写成用户已读"
  route_reason: "TopHub 中的豆瓣榜单只提供公共热度和发现，不替代豆瓣原生的条目、评论、标记与社区关系。豆瓣近期镜像只适合辨认稳定倾向，不适合统计最喜欢类型或形成完整人格画像。"
  reuse_rule: "阅读证据要分强弱：用户主动出题通常是更强真实阅读关系；AI 出题说明候选方向或共同语境，不能自动写成已读。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#豆瓣; 资源/雷达/豆瓣阅读偏好初步观察.md#数据边界与证据强度"
  next_action: "继续由豆瓣原生承担。"
```

---

## 7. BibiGPT：取水与转录，不自动沉 Git

```yaml
- source_id: "BibiGPT"
  current_route: audio_video_water_extraction_transcript_and_discussion_entry
  decision_status: recovered_from_pr
  selected_reason: "长期更新的视频或音频节目中，真正值得进一步处理的单集，可以使用 BibiGPT 保存或处理视频音频，取得摘要、章节、字幕或完整文稿，根据当前问题自定义角度重切，再与 ChatGPT 共同讨论、校勘、比较和提炼。"
  rejected_alternatives:
    - "把 BibiGPT saved library 当 Git 仓库待迁移目录"
    - "把普通转录和摘要全部镜像进 Git"
    - "把 BibiGPT 转录当成待办或完成判断"
    - "只抽嘉宾观点，忽略提问、追问、沉默、迟疑和关系变化"
  route_reason: "BibiGPT saved library 是第一入口。普通转录和摘要留在 BibiGPT；只有形成证据、主题包、长期判断、规则变化或项目动作时才进入沃壤。对话类内容要视需要保留谁提出问题、哪个追问改变回答、哪里出现迟疑或转向、双方关系怎样影响内容。"
  reuse_rule: "BibiGPT 完成的是取水，不是沉积。沉积条件是讨论后形成判断、主题包、领域理解、项目动作或路由规则变化。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#BibiGPT; 资源/雷达/信息流最终分工.md#BibiGPT"
  next_action: "作为影音处理工作流。"
```

---

## 8. Folo：少量长期关系与多媒体阅读室，不镜像文化平台

```yaml
- source_id: "Folo 文化与知识来源边界"
  current_route: limited_long_term_relationships_and_multimedia_reading_room
  decision_status: recovered_from_pr
  selected_reason: "Folo 负责持续跟随少量选定来源，并尽量保留内容原本的媒介形态：文章、图片、音频、视频和项目更新。"
  rejected_alternatives:
    - "把 Folo 当 TopHub 的另一份新闻镜像"
    - "把得到课程、听书、电子书镜像到 Folo"
    - "把全部小宇宙节目、豆瓣关系和视频频道加入 Folo"
    - "让 Folo 依赖 AI 摘要、翻译或自动任务"
  route_reason: "Folo 免费计划足够，当前不需要付费扩容。它优先留下少量一手官方源、真正会完整阅读的作者、长报道、专业深读、少量长期播客与稳定节目、Folo 明显改善呈现的多媒体来源、原生平台无法干净提供的频道和真正触发行为的提醒。"
  reuse_rule: "Folo 收长期关系，不收平台全集。是否加入由持续关系、实际打开、RSS 稳定性、媒介形态和不可替代角色决定。"
  evidence_locator: "资源/雷达/信息流最终分工.md#Folo; 资源/雷达/订阅系统.md#Folo承担"
  next_action: "2026-07-17 复查；不立即扩容。"
```

---

## 9. 沃壤：留下讨论后形成的东西

```yaml
- source_id: "沃壤文化内容沉积规则"
  current_route: judgments_coordinates_topic_packs_projects_not_platform_mirror
  decision_status: recovered_from_pr
  selected_reason: "沃壤保存为什么关注某个来源、节目、作品或人物，它当前由哪个平台承担，什么时候复查，什么条件下重新启用，以及哪些内容真正改变了领域理解、人物坐标或项目动作。"
  rejected_alternatives:
    - "保存每条平台更新"
    - "镜像得到、小宇宙、豆瓣、Folo 或 BibiGPT"
    - "把订阅关系本身当项目"
    - "把只看过、听过、觉得有趣的内容全部沉积"
  route_reason: "只有出现以下变化才沉积：形成了自己的判断，多个材料形成互证主题包，文稿校勘后能脱离原节目阅读，对话暴露长期关系机制或现实处境，内容改变长期领域理解，服务具体项目或行动，暴露工作流、重复或来源路由问题。"
  reuse_rule: "单篇内容先略过、读过即止、平台收藏或 Folo Collection；形成摘要或判断进萃览；形成趋势信号进雷达；改变长期理解进领域；改变动作进项目。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#沃壤; 资源/雷达/订阅系统.md#下沉规则"
  next_action: "继续按沉积门槛执行。"
```

---

## 10. 内容形态：演讲、采访、对话、辩论、水温不是同一层

```yaml
- source_id: "演讲 / 采访 / 对话 / 辩论 / 水温内容"
  current_route: form_sensitive_content_routing
  decision_status: recovered_from_pr
  selected_reason: "不同内容形态产生不同价值。演讲适合快速获取整理完成的框架；采访可能只是信息问答，也可能进入真实处境；真正有价值的对话会让双方都不再只是完成准备好的输出；辩论提供立场碰撞和表达技术；水温内容观察公共空间、知识行业或商业人群当前认为哪些议题值得讲。"
  rejected_alternatives:
    - "把所有长内容都当观点摘要"
    - "把采访形式直接认定为真实对话"
    - "把辩论技巧当真实判断"
    - "把水温入口误当核心 AI、科技或学术来源"
  route_reason: "用户越来越重视真实对话中出现的无准备反应、误解、沉默、迟疑、转向、回避、自我修正和双方关系变化。文化内容不能被简化成观点摘要。"
  reuse_rule: "处理影音时先判断内容形态。演讲可抽框架；采访看是否进入真实处境；对话要保留提问、追问、关系和转向；辩论看赛制与真实声音差异；水温只用于观察公共叙事，不作一手知识。"
  evidence_locator: "资源/雷达/文化内容入口与影音工作流.md#内容形态不是同一件事"
  next_action: "用于 BibiGPT 与 ChatGPT 讨论规则。"
```

---

## 11. 阅读偏好：不是学科画像，而是多条兴趣轴

```yaml
- source_id: "豆瓣阅读偏好与猜书名现场"
  current_route: reading_axes_not_complete_profile
  decision_status: recovered_from_pr
  selected_reason: "近期豆瓣镜像与猜书名现场显示，用户喜欢把抽象系统重新放回人的处境，对系统、失败、周期、预测和方法有稳定兴趣，政治、公共生活和普通人的位置是长期主线，同时阅读空间跨越日本日常散文、科幻、戏剧、漫画、视觉叙事、研究方法、商业投资和现实判断。"
  rejected_alternatives:
    - "把近期豆瓣镜像当完整阅读史"
    - "把用户书架简化成哲学、社科、科学、欧美文学的思想型书架"
    - "把严肃阅读、小说、漫画、设计分成高低等级"
    - "让 AI 推荐只输出一张‘像你’的书单"
  route_reason: "阅读偏好文件明确：豆瓣近期镜像只适合辨认若干稳定倾向；猜书名 PR 是互动中的活体样本。形式、人物、叙事、图像、版式、镜头和结构都可能承载思想。"
  reuse_rule: "荐书或文化来源判断要说明补的是哪条兴趣轴，是否打开盲区，中文版本与现实可得性是否核实，为什么可能产生新视角，以及现在读还是留候选池。"
  evidence_locator: "资源/雷达/豆瓣阅读偏好初步观察.md#当前可用的阅读偏好描述; #AI推荐更适合扩展盲区"
  next_action: "用于阅读与文化推荐边界。"
```

---

## 12. 本补充 ledger 的复用规则

1. 文化不等于娱乐目录，知识不等于持续摄入。
2. TopHub 看公共回声，Folo 接住少量原声与媒介，沃壤留下坐标。
3. 得到承担课程、听书、电子书和结构化知识消费，不镜像完整购买、收藏和学习历史。
4. 得到人物关系按影响、变化、按需返回和明确非关系分层，不按名气排序。
5. 小宇宙承担播客发现、节目关系和直接收听，不全量搬入 Folo。
6. 豆瓣承担书影音条目、评论、标记和社区关系；TopHub 豆瓣榜只提供公共发现。
7. BibiGPT 是影音取水和转录入口，普通转录不自动沉 Git。
8. Folo 不镜像得到、小宇宙、豆瓣或视频平台，只收少量长期关系和多媒体阅读室需要。
9. 沃壤只保存关系、判断、重新启用条件、主题包、领域理解和项目动作。
10. 演讲、采访、对话、辩论和水温内容要分形态处理。
11. 形式、叙事、图像、镜头、版式和媒介本身可以承载思想，不能只抽观点。
12. AI 推荐用于扩展盲区，不替用户宣布“应该读什么”。

---

## 13. 仍需继续追回

文化、知识内容与影音工作流已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 若需要，可继续补 `TopHub 能力地图 / 使用配置 / 全部订阅顺序 / 订阅源账本` 的总分工 ledger。
