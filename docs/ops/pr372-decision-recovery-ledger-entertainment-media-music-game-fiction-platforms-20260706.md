# PR #372 判断链追回：娱乐影视、音乐、游戏、网文与动漫平台专项（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中娱乐目录剩余影视平台、音乐平台、游戏目录、网文平台和动漫平台的判断链。

纳入文件：

- `docs/ops/tophub-entertainment-imdb-closure-20260705.md`
- `docs/ops/tophub-entertainment-maoyan-closure-20260705.md`
- `docs/ops/tophub-entertainment-tencent-video-closure-20260705.md`
- `docs/ops/tophub-entertainment-iqiyi-closure-20260705.md`
- `docs/ops/tophub-entertainment-qqmusic-closure-20260705.md`
- `docs/ops/tophub-entertainment-netease-cloud-music-closure-20260705.md`
- `docs/ops/tophub-entertainment-qidian-closure-20260705.md`
- `docs/ops/tophub-entertainment-zongheng-closure-20260705.md`
- `docs/ops/tophub-entertainment-qimao-closure-20260705.md`
- `docs/ops/tophub-entertainment-17173-closure-20260705.md`
- `docs/ops/tophub-entertainment-haoyoukuaibao-closure-20260705.md`
- `docs/ops/tophub-entertainment-anime-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 影视平台总规则：作品发现入口少量保留，片库榜按需

```yaml
- source_id: "影视平台榜总规则"
  current_route: keep_cross_platform_attention_and_discussion_use_platform_library_rankings_on_demand
  decision_status: recovered_from_pr
  selected_reason: "影视平台榜能提供电影、剧集、纪录片、动漫、综艺和票房的观看线索。"
  rejected_alternatives:
    - "把猫眼、腾讯视频、爱奇艺、IMDb 的所有电影 / 剧集 / 票房 / 片库榜常驻"
    - "把平台片库榜当跨平台作品发现"
    - "把票房、期待值、站内热播等同于质量"
    - "把平台榜加入 Folo"
  route_reason: "系统只保留少量跨平台或跨文化角色：豆瓣一周口碑、IMDb 热门电影、豆瓣全球口碑剧集、豆瓣影评和 Cinephilia。猫眼、腾讯视频、爱奇艺等平台榜按观看平台、票房、纪录片或片库任务使用。"
  reuse_rule: "影视来源先分：跨平台口碑、全球注意力、中文讨论、专业影评、票房市场、单平台片库、纪录片目录、经典库。只有前四类可能常驻；票房和平台片库按需。"
  evidence_locator: "IMDb / Maoyan / Tencent Video / iQiyi closure files"
  next_action: "无新增。"
```

---

## 2. IMDb：保留热门电影，不扩 7 个榜单

```yaml
- source_id: "IMDb｜热门电影榜"
  current_route: keep_tophub_movie_game_music_global_movie_attention_radar
  decision_status: recovered_from_pr
  selected_reason: "它动态呈现全球受众当前关注的电影，同时包含新片、即将上映作品、流媒体作品和重新受到关注的经典片；与豆瓣中文口碑形成地区与文化视角差异。"
  rejected_alternatives:
    - "IMDb Top250电影榜"
    - "IMDb Top250剧集榜"
    - "IMDb 热门剧集榜"
    - "IMDb 最受好评的英语电影"
    - "IMDb 最高票房榜（美国）"
    - "IMDb 烂片榜"
  route_reason: "继续保留在影游音乐。一个节点已足够承担 IMDb 的动态发现功能。Top250 是经典库，变化慢；热门剧集会把娱乐组扩成剧集清单；美国票房是商业表现；烂片榜是趣味反面目录；最受好评英语片与 Top250 重复且范围更窄。"
  reuse_rule: "IMDb 热门电影是注意力雷达，不是质量结论。经典库、票房、烂片和剧集按需查。"
  evidence_locator: "docs/ops/tophub-entertainment-imdb-closure-20260705.md#是否有值得为用户新增; #节点判断"
  next_action: "保持。"
```

---

## 3. 猫眼：票房和档期按需，不进常驻

```yaml
- source_id: "猫眼 5 个节点"
  current_route: on_demand_china_theatrical_market_and_schedule_tool
  decision_status: recovered_from_pr
  selected_reason: "猫眼优势是中国大陆院线票房、排片、期待值和上映期观察。"
  rejected_alternatives:
    - "猫眼国内票房榜常驻"
    - "猫眼热映口碑榜常驻"
    - "猫眼最受期待榜常驻"
    - "猫眼北美票房榜常驻"
    - "猫眼TOP100榜常驻"
    - "猫眼进入 Folo"
  route_reason: "5 个节点都不新增。国内票房是商业表现，不代表质量或日本可观看性；热映口碑与豆瓣一周口碑重叠；最受期待受 IP、明星和宣发影响；北美票房页面停留 2018 年，失去当前性；TOP100 是长期经典目录。"
  reuse_rule: "票房源用于研究市场、档期和某片商业表现；不当作品质量推荐。时效异常节点直接排除。"
  evidence_locator: "docs/ops/tophub-entertainment-maoyan-closure-20260705.md#是否有值得为用户新增; #节点判断"
  next_action: "按需。"
```

---

## 4. 腾讯视频与爱奇艺：纪录片方向有价值，但单平台榜不常驻

```yaml
- source_id: "腾讯视频｜纪录片榜"
  current_route: on_demand_documentary_platform_library_high_priority
  decision_status: recovered_from_pr
  selected_reason: "它包含《十三邀》《中国通史》《河西走廊》《航拍中国》《地球脉动》《人生第一次》《如果国宝会说话》《何以中国》等，确实能提供中文历史、人文、自然或社会纪录片线索。"
  rejected_alternatives:
    - "直接加入 TopHub"
    - "腾讯视频电影 / 电视剧 / 综艺 / 动漫 / 热搜榜常驻"
    - "腾讯视频科技原创 / 游戏原创 / 少儿榜常驻"
    - "腾讯视频平台榜进入 Folo"
  route_reason: "纪录片榜是高优先级按需入口，不常驻。新旧作品长期混排，受腾讯片库和版权约束，混入明星纪录片、世界杯专题、平台节目和经典自然历史片。它是片库排序，不是完整纪录片策展。"
  reuse_rule: "纪录片平台榜用于找可观看片名；若要常驻，应寻找跨平台、编辑型、更新稳定的纪录片策展源。"
  evidence_locator: "docs/ops/tophub-entertainment-tencent-video-closure-20260705.md#是否有值得为用户新增; #高优先级按需入口"
  next_action: "按需。"

- source_id: "爱奇艺 59 个节点"
  current_route: on_demand_iqiyi_library_filters_with_documentary_priority
  decision_status: recovered_from_pr
  selected_reason: "爱奇艺纪录片三榜、文化综艺和部分高分榜能暴露中文历史、人文、自然、社会、科学和文化节目线索。"
  rejected_alternatives:
    - "爱奇艺电影 / 电视剧 / 综艺 / 动漫 / 儿童 / 知识所有切片常驻"
    - "爱奇艺纪录片高分 / 热播 / 飙升三榜同时订阅"
    - "爱奇艺知识榜进入科技或慢读"
    - "爱奇艺平台榜进入 Folo"
  route_reason: "59 个节点都是同一平台片库的热播、飙升、必看、高分、题材、年龄和节目类型切片。纪录片确实是方向，但三榜重复且受片库、版权和运营影响；知识榜分类质量不足；儿童榜无当前需求。"
  reuse_rule: "平台片库只在明确使用该平台时查。纪录片、文化综艺作为按需目录，不构成常驻源。"
  evidence_locator: "docs/ops/tophub-entertainment-iqiyi-closure-20260705.md#是否有值得为用户新增; #最值得保留为按需入口"
  next_action: "按需。"
```

---

## 5. 音乐平台总规则：三条互补入口，不把音乐组变成榜单目录

```yaml
- source_id: "音乐平台榜总规则"
  current_route: keep_three_complementary_music_entry_points
  decision_status: recovered_from_pr
  selected_reason: "音乐平台榜能反映主流热度、地区市场、短视频传播、游戏/动漫曲、原创音乐人与特定平台生态。"
  rejected_alternatives:
    - "同时订阅 Apple Music、QQ音乐、网易云多个总榜 / 热歌 / 新歌 / 地区 / 类型榜"
    - "把地区榜、短视频榜、K歌榜、车友榜、VIP榜、ACG细分榜全部常驻"
    - "把平台榜加入 Folo"
  route_reason: "最终音乐入口压缩为三条互补线：Apple Music 专辑排行负责全球主流和专辑层面；豆瓣华语新碟负责华语专辑、独立发行和唱片评价；网易云原创榜负责中文原创单曲、平台原生创作者与新名字。"
  reuse_rule: "音乐源先判断是专辑、单曲、原创、地区市场、短视频传播、ACG/游戏、K歌场景还是平台会员/车友运营。常驻只保留少数互补发现层；具体音乐人、厂牌或节目另进 Folo 复查。"
  evidence_locator: "QQMusic / Netease Cloud Music closure files"
  next_action: "网易云原创榜替换飙升榜已由娱乐 final 处理。"
```

---

## 6. QQ 音乐：腾讯音乐人原创榜值得注意，但仍按需

```yaml
- source_id: "QQ音乐 35 个节点"
  current_route: on_demand_music_market_slices_no_persistent_subscription
  decision_status: recovered_from_pr
  selected_reason: "QQ 音乐覆盖热歌、飙升、新歌、流行指数、地区、语种、短视频、K歌、听歌识曲、动漫、游戏音乐、腾讯音乐人原创和国际榜。"
  rejected_alternatives:
    - "QQ音乐热歌 / 飙升 / 新歌 / 流行指数常驻"
    - "QQ音乐腾讯音乐人原创榜常驻"
    - "QQ音乐游戏音乐榜 / 动漫音乐榜常驻"
    - "日本公信、台湾 KKBOX、香港 TVB、Melon、UK、Billboard 等地区榜常驻"
    - "QQ音乐榜单进入 Folo"
  route_reason: "不新增。腾讯音乐人原创榜最接近新增价值，但仍是单曲榜，且豆瓣华语新碟和网易云原创榜已经承担更清晰的专辑/原创分工。地区榜和类型榜会把音乐组变成全球榜单目录；部分节点明显过时。"
  reuse_rule: "QQ 音乐适合按主题歌单、地区市场、游戏/动漫音乐或中文原创单曲任务调用；不常驻。"
  evidence_locator: "docs/ops/tophub-entertainment-qqmusic-closure-20260705.md#是否有值得为用户新增; #节点判断"
  next_action: "按需。"
```

---

## 7. 网易云：原创榜替换飙升榜，不增加数量

```yaml
- source_id: "网易云音乐｜原创榜"
  current_route: tophub_movie_game_music_chinese_original_music_replacement
  decision_status: recovered_from_pr
  selected_reason: "原创榜比飙升榜更接近当前系统缺少的中文原创音乐人与平台原生创作者发现。样本包含房东的猫、陈粒、河图、黄诗扶、银临、h3R3、ICE杨长青、万妮达、塞壬唱片-MSR、Life Awaits、木马等。"
  rejected_alternatives:
    - "继续保留网易云飙升榜"
    - "原创榜与飙升榜同时订阅"
    - "网易云 62 个可区分节点中的 ACG、地区、车友、VIP、音乐合伙人、AI、听歌识曲等榜单常驻"
  route_reason: "替换而不是新增。飙升榜混合综艺 Live、旧歌翻红、儿歌、短视频 BGM、DJ 与速度改编、平台突然传播单曲；原创榜能把网易云入口从传播热度转向创作者发现。双订会让音乐组被网易云单曲流占据。"
  reuse_rule: "平台内音乐源只保留一个角色清楚的入口。中文原创发现优先原创榜；ACG、游戏、地区、车友和 VIP 榜按歌单或研究任务调用。"
  evidence_locator: "docs/ops/tophub-entertainment-netease-cloud-music-closure-20260705.md#是否有值得新增; #为什么用原创榜替换"
  next_action: "已纳入娱乐最终动作。"
```

---

## 8. 网文平台：有市场价值，但没有持续使用场景

```yaml
- source_id: "起点中文网 28 个节点"
  current_route: on_demand_web_fiction_market_and_reading_tool
  decision_status: recovered_from_pr
  selected_reason: "起点能观察中文网文市场结构，覆盖男频和女频，阅读指数、签约作者新书、畅销、月票、推荐、收藏、书友等信号各有用途。"
  rejected_alternatives:
    - "起点中文网 / 起点女生网阅读指数榜常驻"
    - "畅销、月票、推荐、收藏、书友、新书、限时免费等全部常驻"
    - "把起点榜单放慢读与思想或影游音乐"
    - "起点榜单进入 Folo"
  route_reason: "不新增。存在网络文学空白不等于必须用平台榜填满。起点男频/女频若只订一个会偏市场，两个都订会增加重复；榜单高度重叠；新书和新人榜噪声高；限免是促销；更新字数奖励产量。"
  reuse_rule: "若找书，优先阅读指数和签约作者新书；若研究大众幻想、性别叙事或类型市场，男女频成对比较；若追更，直接跟具体作品或作者。"
  evidence_locator: "docs/ops/tophub-entertainment-qidian-closure-20260705.md#是否有值得为用户新增; #网络文学缺口的最终判断"
  next_action: "按需。"

- source_id: "纵横中文网 9 个节点"
  current_route: on_demand_male_frequency_web_fiction_sample
  decision_status: recovered_from_pr
  selected_reason: "纵横提供起点之外的男频网文样本，尤其玄幻、仙侠、军事、历史、都市和长篇连载；完结榜对偶尔找完整作品有独立价值。"
  rejected_alternatives:
    - "纵横 24 小时畅销 / 月票 / 推荐 / 点击 / 捧场榜常驻"
    - "纵横替代起点按需优先级"
    - "纵横进入 Folo"
  route_reason: "不新增。9 个节点全部偏男频；商业榜高度重复；新书样本少且类型化；更新榜奖励产量；与起点相比覆盖更窄，不足以成为网文总入口。"
  reuse_rule: "纵横适合研究男频、完整作品和老头部时按需；不承担全网文市场总览。"
  evidence_locator: "docs/ops/tophub-entertainment-zongheng-closure-20260705.md#是否有值得为用户新增; #与起点中文网的比较结论"
  next_action: "按需。"

- source_id: "七猫中文网 14 个节点"
  current_route: on_demand_free_web_fiction_platform_sample
  decision_status: recovered_from_pr
  selected_reason: "七猫能观察免费阅读平台上的题材、标题、商业化趋势和男女频差异；完结榜对低负担试读有价值。"
  rejected_alternatives:
    - "七猫大热 / 收藏 / 原创风云 / 原创飞跃 / 更新榜常驻"
    - "七猫替代起点阅读指数榜"
    - "男女频多个热度榜同时订阅"
    - "七猫进入 Folo"
  route_reason: "不新增。榜单高度重复，标题和题材高度平台化，更新榜只反映活跃度，原创风云/飞跃口径不透明，七猫与起点、纵横、微信读书有大量作品和类型重叠。"
  reuse_rule: "七猫用于观察免费阅读商业模式、标题包装和男女频题材；完结榜和新书榜按需，不常驻。"
  evidence_locator: "docs/ops/tophub-entertainment-qimao-closure-20260705.md#是否有值得为用户新增; #与起点纵横微信读书的比较结论"
  next_action: "按需。"
```

---

## 9. 游戏目录：老网游档案和手游新品都按需

```yaml
- source_id: "17173 12 个节点"
  current_route: on_demand_chinese_old_online_game_and_legacy_market_sample
  decision_status: recovered_from_pr
  selected_reason: "17173 能作为研究中国网络游戏长期记忆、老网游生态、怀旧文化和传统网游题材分类的样本。"
  rejected_alternatives:
    - "17173 热门游戏榜常驻"
    - "17173 新游期待榜 / 手游期待榜常驻"
    - "17173 手游热门、二次元、策略、回合、武侠、玄幻、奇幻等分类榜常驻"
    - "17173 进入 Folo"
  route_reason: "不新增。热门榜更像多年累计投票旧游戏目录，期待榜维护状态失真，手游分类榜旧作和多代产品混排，不能回答当前哪些游戏活跃、哪些新作值得投入时间或哪些独立作品值得发现。"
  reuse_rule: "17173 用作老网游和怀旧文化档案样本，不作当前游戏发现。"
  evidence_locator: "docs/ops/tophub-entertainment-17173-closure-20260705.md#结论; #按需入口"
  next_action: "按需。"

- source_id: "好游快爆 5 个节点"
  current_route: on_demand_chinese_mobile_game_new_product_sample
  decision_status: recovered_from_pr
  selected_reason: "好游快爆能反映中国移动游戏市场的预约、下载和版本热度；预约榜能看到尚未上线或准备移植到移动端的作品。"
  rejected_alternatives:
    - "好游快爆预约榜常驻"
    - "好游快爆飙升 / 热门 / 下载 / 二次元榜常驻"
    - "好游快爆进入 Folo"
  route_reason: "不新增。预约榜过宽，混有低透明度 IP 手游、端游移植、云游戏版本、玩家自制、恐怖版、抽卡版和样本量不透明项目；飙升、热门和下载榜高度重叠；二次元榜是商业手游热度，不是动画或文化入口。"
  reuse_rule: "好游快爆预约榜是国内手游新品研究入口，不是游戏文化常驻。下载不等于质量或长期留存。"
  evidence_locator: "docs/ops/tophub-entertainment-haoyoukuaibao-closure-20260705.md#是否值得新增; #按需入口"
  next_action: "按需。"
```

---

## 10. 动漫：没有合格的跨平台低噪音发现榜

```yaml
- source_id: "动漫 18 个节点"
  current_route: on_demand_anime_platform_and_forum_sources_no_persistent_subscription
  decision_status: recovered_from_pr
  selected_reason: "动漫标签显示了平台播放榜、论坛/资讯流和动漫音乐榜。Stage1st 是最有按需价值的核心观众反应入口。"
  rejected_alternatives:
    - "美漫百科常驻"
    - "微信动漫 24h 热文榜常驻"
    - "百度视频 / 腾讯视频 / 360影视 / 爱奇艺多个动漫榜常驻"
    - "羁绊网二次元资讯常驻"
    - "QQ音乐动漫音乐榜作为动漫入口"
    - "Stage1st 动漫论坛进入 TopHub 或 Folo"
  route_reason: "不新增。平台播放榜受库存、版权、儿童内容影响，不能回答跨平台、当季、真正值得看的动画是什么；微信动漫标签污染严重；羁绊网陈旧；美漫百科以流言为主；Stage1st 有价值但仍是论坛热帖，不是编辑筛选。"
  reuse_rule: "动漫来源按当季作品、制作研究、核心观众反应、平台库存、音乐歌单、美漫传闻分层。真正常驻应等待稳定动画评论、制作研究或季番策展来源。"
  evidence_locator: "docs/ops/tophub-entertainment-anime-closure-20260705.md#是否有值得新增; #按需优先级"
  next_action: "按需。"
```

---

## 11. 明确排除信号

```yaml
- source_id: "娱乐影视音乐游戏网文动漫平台排除信号"
  current_route: reusable_rejection_rules
  decision_status: recovered_from_pr
  selected_reason: "用于后续复用。"
  rejected_alternatives:
    - "把单平台片库、站内热度、票房、期待值、下载量、更新时间、累计投票、更新字数、播放量和分区标签当成常驻理由"
  route_reason: "明确排除：时效异常的票房/音乐/资讯榜；平台片库多切片；高分/热播/飙升/类型榜重复；票房和期待值替代质量判断；音乐地区榜堆叠；网文平台运营指标当文学价值；老网游累计投票当当前热度；手游下载/预约当质量；动漫播放榜当跨平台发现。"
  reuse_rule: "平台指标只能回答平台内部行为，不能直接回答作品是否值得看、听、玩或读。"
  evidence_locator: "all included media/music/game/fiction/anime platform files"
  next_action: "作为复用规则。"
```

---

## 12. 本补充 ledger 的复用规则

1. 影视常驻入口要跨平台、跨文化或能提供评论解释；平台片库榜按需。
2. IMDb 热门电影是全球注意力雷达，不是质量榜；猫眼是票房与档期工具。
3. 腾讯视频和爱奇艺的纪录片榜有价值，但只是片库目录，不是完整策展。
4. 音乐入口压缩为 Apple Music 专辑、豆瓣华语新碟、网易云原创榜三条互补线。
5. QQ 音乐地区/类型/短视频/游戏/动漫榜按歌单或市场任务调用，不常驻。
6. 网易云原创榜替换飙升榜，是从传播热度转向创作者发现；不增加网易云数量。
7. 起点、纵横、七猫用于网文市场、题材和按需找书，不制造日常待读库存。
8. 17173 是老网游和怀旧文化档案样本；好游快爆预约榜是国内手游新品样本。
9. 动漫暂缺合格的跨平台、当季、低噪音发现榜；Stage1st 按需观察核心观众反应。
10. 平台榜不进 Folo；具体导演、影评人、音乐人、厂牌、作者、工作室、栏目或节目另行复查。

---

## 13. 仍需继续追回

娱乐影视、音乐、游戏、网文与动漫平台专项已补成第三份娱乐 ledger。

继续待办：

- 若还要继续娱乐，可补 `QQ音乐 / 网易云音乐` 更细的 ACG / 游戏音乐 / 地区榜使用规则，或补 `得到与知识内容接触谱系`、`文化内容入口与影音工作流` 这些雷达文件的判断链；
- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目。
