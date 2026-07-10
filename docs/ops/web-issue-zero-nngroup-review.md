# Web 第零期 NNGroup 复审 v0.1

时间：2026-05-09
对象：web/index.html + web/content/issue-zero.json / objects.json / scenes.json
用途：用 NNGroup Topics 对第零期最小页面做可用性复审，不继续泛读资料。

## 复审结论

当前骨架方向成立：入场、场位切换、物件靠近解释三步已经能工作。主要问题不在功能，而在信息气味、场位解释、证据追溯和术语负荷。

## 1 Information Scent 信息气味

依据：Information Scent / Information Foraging。

现状：物件的远看文案有气味，但场位按钮只有「本期场刊 / 工作室 / 证据场 / 幕后施工」，对外部访客仍偏概念。

建议：给每个 scene 增加 description 和 question，让用户在点击前知道会看到什么。

## 2 Progressive Disclosure 渐进披露

依据：Accordions、Contextual Help、Progressive Disclosure。

现状：远看、中看、details 靠近解释的三层结构正确，没有一上来压满证据。

建议：保留 details；不要展开全部。近看信息保持四项：来源、痕迹、证明、反默认。后续可让证据场默认展开第一张卡，但现在先不做。

## 3 Cognitive Load 认知负荷

依据：Few Guesses More Success、CASTLE、Cognitive Load。

现状：数量可控（4 场位、8 物件），但术语浓度高：场刊、证据场、反猜意图、物理语法等。

建议：用 scene.description 翻译术语；用 scene.question 说明用户为什么要看这个场位。不要增加新场位。

## 4 Information Architecture 信息架构

依据：IA vs Sitemaps、3 Key Models。

现状：IA 不再是页面树，而是同一批物件在不同场位下重排，这个方向正确。

建议：scenes.json 不应只有 label 和 objectOrder，应增加：description、question、emphasis。这样场位不仅排序，也说明观看规则。

## 5 Usability Heuristics 启发式检查

依据：Jakob Nielsen 10 Heuristics。

通过：
- 用户控制：场位按钮可随时切换。
- 一致性：物件卡结构一致。
- 美学与极简：当前只保留必要信息。

待补：
- 系统状态可见：当前场位除了按钮黑底外，应有一句当前场位说明。
- 识别而非记忆：不要要求用户记住「证据场」含义。

## 6 Trustworthiness 可信度

依据：Trustworthiness in Web Design。

现状：近看已有来源、痕迹、证明、反默认，证据结构成立。

建议：objects.near 预留 links 字段；证据场后续显示关联文档/commit/来源页面。先补 schema，不急着补满链接。

## 最小修改建议

1. scenes.json：每个 scene 增加 description、question、emphasis。
2. objects.json：near 增加 links: []。
3. index.html：在按钮下显示当前场位 description/question。
4. 不改视觉大方向，不加动效，等 Apple Design 进来后再做空间与动效复审。

## 复审后的原则

第零期不是更炫，而是更清楚：用户进入前有气味，切换时知道自己在哪，靠近时能看到证据，离开时知道这些判断从哪里来。
