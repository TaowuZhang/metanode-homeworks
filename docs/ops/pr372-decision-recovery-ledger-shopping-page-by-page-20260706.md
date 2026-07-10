# PR #372 判断链追回：购物逐页 95 节点（2026-07-06）

## 0. 边界

本文件追回 `TopHub > 购物` 目录 95 个节点、12 个内容页逐项复审后的判断链。

纳入文件：

- `docs/ops/tophub-shopping-directory-ledger-20260705.md`
- `docs/ops/tophub-shopping-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-taobao-closure-20260705.md`
- `docs/ops/tophub-shopping-jd-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-jd-closure-20260705.md`
- `docs/ops/tophub-shopping-pinduoduo-empty-closure-20260705.md`
- `docs/ops/tophub-shopping-smzdm-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-deals-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-crowdfunding-reassessment-page-01-20260705.md`
- `docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md`
- `docs/ops/tophub-shopping-credit-cards-reassessment-page-01-20260705.md`
- `docs/ops/tophub-shopping-final-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 购物目录整体：不做“发现可以买什么”的日常流

```yaml
- source_id: "TopHub > 购物 95 节点整体"
  current_route: closed_no_tophub_no_folo_actual_change
  decision_status: recovered_from_pr
  selected_reason: "购物目录已完整审完：淘宝 19/19、京东 21/21、拼多多 0/0、什么值得买 13/13、羊毛线报 13/13、众筹 3/3、图书 18/18、信用卡 8/8，合计 95/95 个节点、12/12 个内容页。"
  rejected_alternatives:
    - "把购物目录变成持续消费发现源"
    - "把价格榜、销量榜、好价榜、羊毛线报加入 TopHub"
    - "把购物流加入 Folo"
    - "建立价格 / 优惠 / 信用卡 / 众筹 / 限免宽泛追踪器"
  route_reason: "最终 TopHub 新增 0、取消 0、替换 0、调序 0；Folo 实际新增 0；追踪器新增 0。购物目录的角色不是发现可以买什么，而是在真实需求出现后按任务调用价格、渠道、用户经验、版本、规格、权益和交付信息；任务结束后退出注意力。"
  reuse_rule: "以后遇到购物来源，先问是否已有真实需求。没有需求，不让促销流生成购买任务；有需求，先写场景、规格、预算和不可妥协条件，再核价、核规格、核渠道、核售后。"
  evidence_locator: "docs/ops/tophub-shopping-directory-ledger-20260705.md#目录状态; docs/ops/tophub-shopping-final-closure-20260705.md#最终架构判断"
  next_action: "无实际配置动作；2026-07-17 复查 Craig Mod。"
```

---

## 2. 淘宝与京东：采购工具，不是质量证据

```yaml
- source_id: "淘宝 19 个节点"
  current_route: on_demand_marketplace_search_and_crowdfunding_discovery
  decision_status: recovered_from_pr
  selected_reason: "淘宝与天猫在明确购买对象后可用于搜索候选、查看价格、店铺、发货、版本和渠道；淘宝众筹有 AI 小硬件、模型、玩具和杂项新品发现价值。"
  rejected_alternatives:
    - "淘宝热销总榜常驻"
    - "每日爆款 / TOP100 人气 / 实时跑量榜常驻"
    - "数码、美食、美妆、母婴、居家等品类切片常驻"
    - "淘宝众筹每日上新常驻"
  route_reason: "19 个节点全部不进 TopHub、Folo 或复查名单。热销、爆款、人气和跑量榜高度重复；品类榜只是同一促销池切片；券后价、月销和短时销量不能证明质量、适用性或真实优惠；普通商品与药品、保健、减肥、私密护理、母婴和医疗器械混排。"
  reuse_rule: "淘宝只在明确需求后使用。价格榜是线索，不是质量证据；健康、药品、保健、母婴和医疗器械不得从促销榜直接选择；众筹先核验交付能力、知识产权和售后。"
  evidence_locator: "docs/ops/tophub-shopping-taobao-closure-20260705.md#淘宝子目录结论; #按需边界"
  next_action: "按需。"

- source_id: "京东 21 个节点"
  current_route: on_demand_official_channel_price_and_spec_check
  decision_status: recovered_from_pr
  selected_reason: "京东适合在采购任务发生后核对自营、官方旗舰店、型号、规格、认证、保修、历史价、配送和售后。"
  rejected_alternatives:
    - "京东总榜 / 9.9 特卖榜常驻"
    - "19 个品类热销榜常驻"
    - "京东图书热销榜作为阅读雷达"
    - "京东健康 / 母婴 / 生鲜 / 家装等榜单常驻"
  route_reason: "21 个节点全部不进 TopHub、Folo 或复查名单。总榜、特卖和品类榜仍是同一商品池切片；多个热销榜出现月销 0，榜单语义失真；食品、居家和百货混入药品、保健、艾灸贴、远红外贴和健康功效商品；图书榜以教辅、套装和营销标题为主。"
  reuse_rule: "京东在型号明确后使用；自营和官方旗舰店是渠道候选，不是质量结论；图书先确认 ISBN、出版社、译者和版本；药品、保健、母婴、健康监测和医疗器械不从热销榜直接决定。"
  evidence_locator: "docs/ops/tophub-shopping-jd-closure-20260705.md#最终结论; #按需边界"
  next_action: "按需。"

- source_id: "拼多多"
  current_route: empty_directory_closed
  decision_status: recovered_from_pr
  selected_reason: "本次页面未显示任何节点。"
  rejected_alternatives:
    - "把拼多多列作未审缺口"
    - "对拼多多平台本身补编判断"
  route_reason: "按 0 节点空目录闭环；未来若重新出现节点，作为新变化重新审核。"
  reuse_rule: "空目录只闭环页面事实，不评价平台本体，也不补编。"
  evidence_locator: "docs/ops/tophub-shopping-directory-ledger-20260705.md#拼多多"
  next_action: "无。"
```

---

## 3. 什么值得买：好价线索，不是消费研究源

```yaml
- source_id: "什么值得买 13 个节点"
  current_route: on_demand_price_history_cross_platform_and_user_experience_tool
  decision_status: recovered_from_pr
  selected_reason: "相比淘宝、京东原始热销榜，什么值得买多了编辑标题、社区筛选和跨平台比价，可在明确商品型号后查历史价格、优惠条件和用户经验。"
  rejected_alternatives:
    - "白菜好价榜"
    - "天猫好价榜"
    - "好价品类榜"
    - "3C 家电好价榜"
    - "滚动最新"
    - "3h 最热"
    - "实时消费热点"
    - "各品类好价榜常驻"
  route_reason: "13 个节点全部不进入持续订阅。它们仍主要是价格流，不是完整评测、购买决策或消费研究来源；今日必买、百亿补贴、国家补贴制造紧迫感；多数价格需要会员、地区、客户端、凑单、领券或多件购买；滚动最新是典型未读制造器；实时消费热点停留在 2024 年事件，不具备 2026 年实时性。"
  reuse_rule: "SMZDM 用于明确商品后的历史价、跨平台价格、优惠条件和用户经验检索；旅行、酒店、跨境和生鲜只在目的地、日期和规则明确后使用。"
  evidence_locator: "docs/ops/tophub-shopping-smzdm-reassessment-pages-01-02-20260705.md#总体判断; #按需角色"
  next_action: "按需。"
```

---

## 4. 羊毛线报：明确退出持续信息流

```yaml
- source_id: "羊毛线报 13 个节点"
  current_route: rejected_from_continuous_stream_with_limited_task_exceptions
  decision_status: recovered_from_pr
  selected_reason: "缺书网和其乐 Keylol 有有限任务价值：缺书网用于确定书目与版本后的促销查询；Keylol 用于主动寻找免费游戏和限时领取。"
  rejected_alternatives:
    - "今日热卖 / 线报酷 / 新赚吧 / 网猴线报 / 线板酷 / 南风线报 / 0818团 / 线报迷 / 爱Q生活网常驻"
    - "羊毛线报进入 Folo"
    - "建立优惠 / 羊毛追踪器"
  route_reason: "综合羊毛线报全部退出持续信息流。它们以极短有效期、账号差异、地区限制、口令、抽奖、助力、金融开户和平台漏洞为中心；大量使用有水、大毛、多 V、多薅、白号、黑号、自测等黑话；混入信用卡、消费金融、基金、证券开户、贷款、数字人民币、漏洞、取消订单和规避风控方案。"
  reuse_rule: "优惠信息不能触发新账户、授信、绑卡或资金路径变化；不使用多账号、漏洞、虚假资料、取消订单或规避风控方案；金融活动必须回官方规则核验。"
  evidence_locator: "docs/ops/tophub-shopping-deals-reassessment-pages-01-02-20260705.md#总体判断; #按需边界"
  next_action: "排除常驻；缺书网和 Keylol 按任务。"
```

---

## 5. 众筹：创意发现，不提供交付证据

```yaml
- source_id: "众筹 3 个节点"
  current_route: on_demand_product_and_culture_discovery
  decision_status: recovered_from_pr
  selected_reason: "淘宝众筹、小米有品众筹和摩点比普通热销榜更接近产品与创意发现。淘宝众筹可看 AI 小硬件、模型、玩具与杂项新品；小米有品适合米家生态、消费电子与家居产品；摩点适合桌游、TRPG、出版、模型与文化周边。"
  rejected_alternatives:
    - "淘宝众筹每日上新常驻"
    - "小米有品众筹每日上新常驻"
    - "摩点众筹最新上线常驻"
    - "进入 Folo"
  route_reason: "3 个节点全部按需。TopHub 页面只显示项目标题，缺少交付记录、团队背景、生产进度、退款规则、知识产权、售后和真实样机证据；众筹是销售前置页，不是独立评测。"
  reuse_rule: "支持任何众筹项目前，必须核验团队、样机、量产、交付历史、退款规则、售后和知识产权；创意发现不等于购买。"
  evidence_locator: "docs/ops/tophub-shopping-crowdfunding-reassessment-page-01-20260705.md#总体判断; #按需角色"
  next_action: "按需。"
```

---

## 6. 图书：保留现有阅读雷达，Craig Mod 作为作者关系复查

```yaml
- source_id: "书格｜每日好书"
  current_route: keep_existing_tophub_books_and_versions_source
  decision_status: recovered_from_pr
  selected_reason: "已是当前 TopHub 第 52 项，提供古籍、图册、版本与数字化资源，角色独立。"
  rejected_alternatives:
    - "重复添加书格"
    - "用购物图书热销榜替代"
  route_reason: "保持，不重复添加。它不是促销榜，而是古籍与数字版本入口。"
  reuse_rule: "版本、古籍和数字化资源来源与销售榜分开；不因购物目录出现图书就替换阅读雷达。"
  evidence_locator: "docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md#书格每日好书"
  next_action: "无；保留。"

- source_id: "豆瓣｜新书速递"
  current_route: keep_existing_tophub_general_new_books_radar
  decision_status: recovered_from_pr
  selected_reason: "已是当前 TopHub 第 48 项，作为综合新书雷达继续保留。"
  rejected_alternatives:
    - "豆瓣社会纪实 / 小说 / 历史文化 / 商业经管 / 文学 / 科学新知 / 艺术设计 / 绘本漫画分类全部常驻"
    - "豆瓣书店作为常驻"
    - "当当 / 京东热销榜作为阅读雷达"
  route_reason: "分类新书页质量可能很高，但都是综合新书速递的子集；拆成多个未读流会扩大阅读债。豆瓣书店是销售选品页；当当和京东热销榜受促销、库存、教辅、套装和营销影响，不承担推荐来源。"
  reuse_rule: "新书发现保留一个综合入口；分类页按当前阅读主题主动打开，不拆未读流。购书先核对豆瓣条目、出版社、译者、ISBN、版本和书评。"
  evidence_locator: "docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md#总体判断; #豆瓣新书速递; #阅读与购书分工"
  next_action: "无；保留。"

- source_id: "Craig Mod"
  current_route: folo_review_candidate_personal_longform_books_japan_walking_software_language
  decision_status: recovered_from_pr
  selected_reason: "它不是价格榜或书目榜，而是稳定围绕书、步行、日本城市与地方、摄影、软件、语言、LLM 和独立会员写作展开的个人长文来源；与用户对日本生活、技术、阅读、地方经验和非默认视角的关注相合。"
  rejected_alternatives:
    - "加入 TopHub"
    - "立即加入 Folo"
    - "把它当图书榜处理"
  route_reason: "加入 2026-07-17 Folo 复查候选，届时核验真实更新量、全文可读性、未读负担和与现有来源的重复度。复查不是承诺加入。"
  reuse_rule: "作者型来源要按作者关系复查，而不是按目录名处理；看长期视角、全文可读、更新频率、重复度和真实阅读率。"
  evidence_locator: "docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md#CraigMod; docs/ops/tophub-shopping-final-closure-20260705.md#唯一新增复查候选CraigMod"
  next_action: "2026-07-17 Folo 复查。"

- source_id: "偶书"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "把偶书作为电子书资源入口"
  route_reason: "页面以数百 GB 电子书合集、网络小说合集和下载工具为主；来源授权、版本质量、文件安全和版权状态无法确认。"
  reuse_rule: "电子书资源不能以容量和下载便利为依据；必须核验授权、版本、格式、文件安全和版权状态。"
  evidence_locator: "docs/ops/tophub-shopping-books-reassessment-pages-01-02-20260705.md#偶书"
  next_action: "排除。"
```

---

## 7. 信用卡与支付：具体任务才查，最终回官方条款

```yaml
- source_id: "信用卡 8 个节点"
  current_route: on_demand_payment_and_card_tasks_official_terms_final
  decision_status: recovered_from_pr
  selected_reason: "美国信用卡指南、美卡论坛和 V2EX 信用卡在美国卡、美国银行账户、点数、里程票、大陆 / 香港 / 境外支付、Apple Pay、OpenAI 支付和订阅支付任务中有线索价值。"
  rejected_alternatives:
    - "我爱卡 24H 热帖榜"
    - "美卡论坛日榜 / 周榜 / 月榜常驻"
    - "美国信用卡指南 Today 常驻"
    - "虚拟信用卡巴士"
    - "V2EX 信用卡常驻"
    - "我爱卡论坛最新发表"
  route_reason: "8 个节点没有一个适合持续推送。论坛节点高噪声，日周月榜重复，部分内容涉及套利、规避风控、灰色虚拟卡、贷款与债务；美国信用卡指南质量最高但主要面向美国信用卡与银行账户，当前长期相关性不足；虚拟信用卡巴士有商业导流、合规与资金安全风险。"
  reuse_rule: "信用卡、银行卡、虚拟卡和跨境支付按具体国家、账户、产品和时间重新核验。最终依据是银行官网、价目表、条款、监管披露和客服确认；不因短期奖励新增授信，不采用绕地区、身份、风控或平台规则的虚拟卡方案。"
  evidence_locator: "docs/ops/tophub-shopping-credit-cards-reassessment-page-01-20260705.md#总体判断; #按需来源; #使用边界"
  next_action: "按需。"
```

---

## 8. 明确排除与证据门槛

```yaml
- source_id: "购物排除信号"
  current_route: reusable_rejection_rules
  decision_status: recovered_from_pr
  selected_reason: "用于复用购物判断。"
  rejected_alternatives:
    - "让促销主动生成购买任务"
    - "让低价、热销、券后价、榜单位置替代质量判断"
  route_reason: "明确排除：所有综合羊毛线报常驻订阅；依赖多账号、虚假资料、漏洞、取消订单或规避风控的方案；未经核验的虚拟卡平台与跨区支付方案；以数百 GB 电子书合集、下载工具和来源不明文件为主的电子书站；从销量榜直接选择药品、保健、减肥、母婴、健康监测或医疗器械；没有明确需求时由促销信息主动生成购买任务。"
  reuse_rule: "购物源的证据门槛随类别提高：普通商品核规格和售后；高价值商品核历史价、保修和第三方评测；健康、金融、母婴、药品和医疗器械必须回官方、监管、医生 / 专业标准和产品条款。"
  evidence_locator: "docs/ops/tophub-shopping-final-closure-20260705.md#明确排除; docs/ops/tophub-shopping-directory-ledger-20260705.md#购物目录的最终角色"
  next_action: "作为复用规则。"
```

---

## 9. 本补充 ledger 的复用规则

1. 购物目录不承担持续发现“可以买什么”；只有真实需求出现后才调用。
2. 淘宝、京东和 SMZDM 是任务工具：核价、核型号、核渠道、核用户经验，不是质量证据。
3. 低价、券后价、月销、短时销量、热度和百亿补贴不能替代适用性判断。
4. 羊毛线报、金融优惠和虚拟卡风险源不能进入持续流；不为几元奖励新增账户、授信、绑卡或资金路径。
5. 众筹只做创意发现；支持前核验团队、样机、量产、交付、退款、售后和知识产权。
6. 图书保留 `书格｜每日好书`和`豆瓣｜新书速递`；分类页按主题主动打开，不拆成多个未读流。
7. Craig Mod 是作者关系，不是购物源；只进入 Folo 复查。
8. 信用卡与跨境支付最终回银行官网、条款、监管披露和客服确认。
9. 健康、金融、母婴、药品、医疗器械、身份、账户和支付相关内容提高证据门槛。
10. 购买完成后结束任务，不把促销流长期留在注意力中。

---

## 10. 仍需继续追回

购物逐页 95 节点已补成单独 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 政务、校务、专栏、浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
