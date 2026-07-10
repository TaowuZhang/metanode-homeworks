# Web 第零期封板说明 v0.1

时间：2026-05-09
对象：web/index.html + web/content/issue-zero.json / objects.json / scenes.json
状态：第零期第一个可保存版本

## 一句话

第零期是一个可运行的个人网站现场原型：用场刊入场，用工作室重组材料，用证据场显影来路，用幕后施工保留未完成。

它不是正式首页，不是作品集，不是简历，也不是 Notion 公共页的替代品。它是网站生成前的第一个现场版本。

## 已完成

### 研究与方法

- 建立「场刊 / 工作室 / 证据场」三意象。
- 将三意象从视觉风格校正为组织方法。
- 写入 Web 物理语法方法文档。
- 写入第零期 brief、回顾、物件语法。

### 数据化

- 生成 `web/content/issue-zero.json`。
- 生成 `web/content/objects.json`。
- 生成 `web/content/scenes.json`。
- 生成数据预览文档，确认 issue / objects / scenes 可读、可索引、可重排。

### 页面骨架

- `web/index.html` 已接入第零期数据。
- 展示 opening。
- 展示 Programme。
- 支持 4 个场位切换。
- 每个场位按 `objectOrder` 重排物件。
- 每个物件有远看 / 中看 / 靠近解释。
- 靠近解释保留来源、痕迹、证明、反默认。

### NNGroup 复审

NNGroup Topics 已导入并参与复审。复审重点：

- Information Scent：增强场位说明与物件远看。
- Progressive Disclosure：保留远看 / 中看 / 靠近解释三层。
- Cognitive Load：降低术语负荷。
- Information Architecture：让 scene 不只是排序，而是观看规则。
- Usability Heuristics：补当前场位说明、识别而非记忆。
- Trustworthiness：为 near 层预留 `links`。

已落地：

- `scenes.json` 增加 `description`、`question`、`emphasis`。
- `objects.json` 的 `near` 增加 `links: []`。
- 页面显示当前场位说明和「这一场在问」。

### Copy Pass

完成从研究语气到场刊语气的第一轮文案修整。

已写回：

- `issue.opening`
- 4 个 `scenes.description`
- 8 个物件 `objects.far`

目标：让外部访客不用先理解整套方法，也能知道自己正在进入什么、能点什么、靠近后会得到什么。

### Apple Design 复审

Apple Design 两个拆包已导入，导入日志页已归档，内容页保留。

Apple 复审重点：

- Continuity：场位切换应像同一现场里的观看位置变化。
- Spatial layout：重要内容在中心，避免窗口过多。
- Designing for visionOS：靠近解释应像物件被拿近，而不是弹窗。
- Liquid Glass：控制层可以轻轻浮起，内容层要清楚，避免整页材质化。
- Typography：文字仍是第一界面，不为深度牺牲可读性。

已落地：

- 场位按钮区增加 shelf / 工作台边缘感。
- `scene-note` 增加便签感。
- 卡片 hover / focus 轻微 lift。
- `details[open]` 时卡片像被拿近。
- 增加 `prefers-reduced-motion`。

## 当前页面能力

当前第零期页面可以：

- 读取本地 JSON 内容数据。
- 展示第零期 opening。
- 展示节目单。
- 在 4 个场位之间切换：本期场刊、工作室、证据场、幕后施工。
- 根据场位重排同一批物件。
- 让访客从远看进入中看，再展开近看证据。
- 在不猜访客意图的前提下，响应访客主动选择。

## 明确不做

当前阶段不做：

- 不做完整正式首页。
- 不做作品集。
- 不做简历页。
- 不做 Apple 风拟态。
- 不做整页 Liquid Glass。
- 不新增场位。
- 不新增物件。
- 不引入前端框架。
- 不把 `details` 改成 modal。
- 不牺牲文字可读性。

## 当前风险与待补

- `objects.near.links` 已预留，但还没有回填证据链接。
- `Apple / NNGroup 缺席位` 这个物件已经过期，应改成「两块校正石入场」。
- `web/index.html` 仍是单文件，后续扩展时需要考虑拆分。
- 尚未更新 `web/README.md`。
- 尚未做人工视觉截图 / 视觉走查记录。
- 尚未决定部署策略。

## 下一轮入口

建议下一轮按以下顺序推进：

1. 把 `Apple / NNGroup 缺席位` 改成「两块校正石入场」。
2. 回填 `objects.near.links`。
3. 做一次数据 + 浏览器 smoke test。
4. 更新 `web/README.md`。
5. 记录人工视觉走查。
6. 决定是否拆分 `web/index.html`。
7. 决定部署方式。

## 封板判断

第零期已经从概念、研究、资料导入，推进到一个可运行、可解释、可校正、可保存的版本。

它还不是正式网站，但已经是一个站得住的现场。
