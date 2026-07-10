# Web 物理语法综合方法 v0.1

时间：2026-05-09
用途：把「场刊 / 工作室 / 证据场」与 UIUX 导入资料综合成可施工方法。

## 判断
这个网站不是作品集、AI 产品、Notion 公共页，也不是猜访客意图的 UI。
它是：由场刊入场、由工作室重组、由证据标签显影的个人现场；访客不被猜测，只被允许显式选择靠近方式。

AI 是时代背景，不是主语。吸收 AI-UX 的控制权、上下文、透明度、过程可见、多入口、信任线索；拒绝「猜你想看」、自动个性化、AI 导览员、隐性行为推荐。

## 三意象的反惯性
- 场刊不是杂志模板：保留入场感、当前期、现场边界、节目顺序；避免精美目录和永久 About Me。
- 工作室不是卡片墙：保留未完成性、同物多义、临时摆放、可重组；避免 gallery / dashboard / 普通筛选器。
- 证据场不是数据面板：保留来源活动、遗留物、证明关系、反默认判断、可追溯链接；避免 KPI、自夸、红线侦探板。

## 意志保护 / 反猜意图护栏
不用「意图驱动」作核心词，改为：
- 意志保护
- 反猜意图护栏
- 显式场位驱动
- 观看姿势驱动
- 靠近方式驱动

不是系统猜用户意图，而是保护第一使用者的意志生成：稻谷的表达、访客的显式选择、愿意 / 不愿意的保留。所有场位必须明确可选，所有自动变化必须可解释、可回退、可关闭。

理由链：
- 「意图」容易被第三人称化：系统、工程师、模型都可以把它当成可预测对象。
- 「意志」是第一人称生成：它不是被识别的对象，而是在靠近、拒绝、选择和承担中练出来。
- 大厂工程语境常把「猜用户意图」当作能力；这里反而要保护不被猜、不被代替、不被自动补全。
- 所以目标不是 intent fidelity，而是 will protection / anti-intent-guessing / explicit scene selection。

第一版场位：
1. 本期场刊
2. 工作室
3. 证据场
4. 外排出口
5. 幕后施工

## 从资料吸收
### Shape of AI
- Wayfinders -> 场刊入场线索
- Inputs + Context -> 点击、lens、靠近物件是显式上下文
- Governors -> lens、展开、证据、返回都必须可控
- Trust indicators -> 来源、证据、更新时间、可追溯贴在物件上
- Synthesis / Remix / Footprints -> 同一批材料进入不同编排，并显示来路

### AI-UX interactions
吸收 Human in the loop、Multi-modal、Agentive UX 的「协助但不替人判断」「多入口」「主动但不打扰」。
不吸收隐形 agent、行为推断、自动改写叙事。

### Refactoring UI
先做功能，不先做 layout。第一版只验证：
1. 入场
2. 重组
3. 靠近解释

层级优先；用弱化背景突出重点；用重叠、阴影、层次制造可进入空间。


### Sajid Code UI
这一批资料补的是落地手感，而不是新概念。
- Only Noobs Build Beautiful Websites -> 抵抗“漂亮网站”冲动；网页首先要让进入、重组、解释成立。
- Original Ideas are Overrated -> 可用性胜过原创性；三意象可以陌生，但入口、展开、返回要熟悉。
- The 80% of UI Design - Typography -> 首版视觉重点放在文字层级、密度、行距、标签和边注。
- Good Design Is As Little Design As Possible -> 少设计，不堆效果；用最少手段让物理语法成立。

迁移到本站：先打磨场刊节目单、物件标题、气味短句、证据标签、来源说明，而不是先追求视觉奇观。

### Laws of UX
- Cognitive Load：不要一次性压出三意象、lens、证据。
- Jakob’s Law：用熟悉交互承载陌生内容。
- Hick / Choice Overload：首屏 3-5 个主要选择足够。
- Gestalt：用留白、边界、邻近、相似组织物件。
- Peak-End：设计强入场峰值和好的离开时刻。
- Tesler：复杂性不会消失，应沉到 schema、材料层、生成层。

## 两层架构
可见层：入场层、工作室层、证据层、幕后层、出口层。
不可见层：
1. Threshold Engine 阈限引擎
2. Object Grammar 物件语法
3. Recomposition System 重组系统
4. Perceptual Choreography 感知编舞
5. Provenance Layer 来路层

## 第一版数据语法
先用 JSON / Markdown，不急上框架。
- issue：当前期、主题、opening、programme
- object：id、title、kind、state、teaser、roles、provenance、display
- scene：id、label、description、rules

scene / lens 是显式选择，不是推断用户意图。

## 技术初判
采纳：HTML / CSS / 原生 JS、CSS Grid、CSS custom properties、JSON/Markdown。
试验：View Transition、IntersectionObserver、Popover/Dialog、Anchor Positioning。
评估：Astro / Eleventy、Web Components、GSAP。
暂缓：React / Next、Three.js/WebGL、AI chat/agent 导览。

## 检查清单
- 是否猜访客意图？
- 是否自动改写叙事？
- 是否把行为信号当成猜测许可？
- 是否保护显式场位与靠近方式？
- 是否保留拒绝、命名、判断、承担的空间？
- 场刊是否真的入场？
- 工作室是否真的重组？
- 证据是否贴在物件上？
- 首屏是否只给 3-5 个主要线索？
- 是否渐进披露？
- 是否有可回退、可关闭的控制？

## 后续补审
Apple Design：补空间连续性、motion as orientation、物件被拿近而不是弹窗。
NNGroup Topics：补 information scent、progressive disclosure、navigation、scannability。
Airbnb / Duolingo / Sajid Code UI / Material / Fonts / Icons 后续按需补。

## 一句话
不设计页面，设计进入条件。
不设计栏目，设计物件的可变身份。
不设计证据模块，设计遗留物如何被解释。
不设计动画，设计注意力如何移动。
不猜访客意图，只保护显式靠近方式。
