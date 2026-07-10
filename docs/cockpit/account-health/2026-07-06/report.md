# Cockpit account health inspector · 2026-07-06

> Read-only report. 只报告，不关 issue、不改 label、不写 Notion、不替谷裁决退相干。
> 器官分三档：坏户口（缺出生即存在的器官）· 待补全（成熟/晚裁时补，不算坏）· 健康。公约母法：真实>形式齐整，不逼出未长出的器官填满。

- Open cockpit issues scanned: 17
- 坏户口·needs repair: 11
- 待补全·incomplete: 0
- 健康·healthy: 6

## 坏户口 · Needs repair（缺出生器官）

- #235 [cockpit] 同步：Notion 工作台 × GitHub Project 1 项目户口
  - 硬伤 missing core: 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:候补, lane:暂留
  - https://github.com/dongxi-heji/worang/issues/235
- #204 [cockpit] 落地服务台：AI 工作流落地服务台硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, lane:待裁决, waterline:候补
  - https://github.com/dongxi-heji/worang/issues/204
- #203 [cockpit] 题库工程：自考会计题库资料源硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:候补, lane:AI可推进
  - https://github.com/dongxi-heji/worang/issues/203
- #197 [cockpit] 外运 Notion：现场沉淀到 GitHub 硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:候补, lane:暂留
  - https://github.com/dongxi-heji/worang/issues/197
- #196 [cockpit] 清理链接：Notion / GitHub 地层健康硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:候补, lane:AI可推进
  - https://github.com/dongxi-heji/worang/issues/196
- #195 [cockpit] 运行 Agent：后台自动化边界硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, lane:暂留, waterline:水下
  - https://github.com/dongxi-heji/worang/issues/195
- #194 [cockpit] 递交任务包：AI 共台投喂硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:今日, lane:AI可推进
  - https://github.com/dongxi-heji/worang/issues/194
- #193 [cockpit] 守门资源：多源入库判断硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:候补, lane:AI可推进
  - https://github.com/dongxi-heji/worang/issues/193
- #192 [cockpit] 推进 WQB：平台状态分析硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, waterline:今日, lane:AI可推进
  - https://github.com/dongxi-heji/worang/issues/192
- #191 [cockpit] 梳理 PR：#200 快照候选硬锚点
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）
  - labels: cockpit, frontstage:github-evening, lane:谷WIP, waterline:今日
  - https://github.com/dongxi-heji/worang/issues/191
- #190 [cockpit] 两前台底盘：Notion 早晨 × GitHub 晚上
  - 硬伤 missing core: 一句话, 接续入口
  - 顺手待补 pending: 证据锚点（成熟时补）, 出口（成熟时补）, morning prompt（已上浮点火线·必补）
  - labels: cockpit, frontstage:github-evening, lane:待裁决, waterline:候补
  - https://github.com/dongxi-heji/worang/issues/190

## 待补全 · Incomplete（不算坏，成熟/晚裁时补）

- None.

## 健康 · Healthy

- #212 [cockpit] 学习线集群（七条学·线一户总管）
- #211 [cockpit] 建·个人页面（坐标库替代简历）
- #210 [cockpit] 建·技术栈路线（roadmap.sh 补短板）
- #209 [cockpit] 研·CC源码碰法（挖出改造克谷碰法的东西）
- #208 [cockpit] 外显：推·自媒体（禾集 0→1）
- #207 [cockpit] 视觉索引：Eagle 字段规范

## 七器官 · Seven organs（出生必填 vs 补强）

出生即存在·缺了就漂（坏户口）：
1. 一句话 / one-liner
2. 当前状态 / current state
3. 接续入口 / handoff entry

成熟/晚裁时补强·缺了不算坏：
4. Area · 领域（出生选未归域也行；人看的软分类，不进自动化）
5. Lane / Waterline（晚裁产物，后补挂 label）
6. 证据锚点 / evidence anchor（没真锚点就留空，别造假）
7. 出口 / exit；Morning prompt（只对 waterline:今日/候补 的户口必需）

Morning prompt 被明早点火线读取：waterline:今日 → 主区，候补 → 旁路，水下 → 排除。