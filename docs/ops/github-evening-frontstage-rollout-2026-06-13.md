# GitHub 全天行动总账 · rollout map · 2026-06-13

> 历史 rollout 记录。2026-07-02 校准：Notion 早晨前台已经退场，GitHub Project 1 继续作为全天行动总账；本页保留当时字段如何形成，不再把 Notion Dashboard 写成现行器官。

## 0. 当前完成态

- GitHub Project 1 已改造成全天行动总账：<https://github.com/orgs/dongxi-heji/projects/1>
- GitHub Project 2 已创建为晚上裁决样机：<https://github.com/orgs/dongxi-heji/projects/2>
- Project 1 已写入 #190–#197、#200 的第一批字段值。
- Project 1 是长期事实母体；Project 2 只保留为施工样机 / 可迁移参考。
- 原 Notion 早晨前台已于 2026-07-02 清场；早晨阅读直接由 Project 1 的精简视图或其他 GitHub 前台承担。

## 1. Project 1 的角色

Project 1 = `worang 全天行动总账`。

它包含：

```text
早上怎么开工
白天谁推进
晚上怎么裁决
什么暂留
什么代谢
什么完成 / 退相干
```

早晨视图不是第二事实源，只是 Project 1 中“水位 = 今日”的轻量读法。它可以在 GitHub 内呈现，不再依赖 Notion 数据库或 Dashboard。

## 2. Project 1 已建字段

- `早晨列`
- `晚上列`
- `水位`
- `火`
- `接续之问`
- `给克一句话`
- `晚上判断`
- `出口`
- `代谢`
- `Last touched`
- `地界`

## 3. 字段语义

### 早晨列

```text
直接做
交给克
今天占位
不进早上
```

用途：承载早晨三列结构；当前由 GitHub Project 视图直接读取。

### 晚上列

```text
待裁决
谷 WIP
AI 可推进
暂留
代谢台
完成·退相干
```

用途：晚上控制权裁决。

### 水位

```text
今日
候补
水下
```

用途：决定是否进入人的近期注意力前台，而不是决定是否进入 Notion。

### 火

```text
烧
温
冷
```

用途：记录人的意愿温度。

### 出口

```text
未到出口
关
接
沉
退
```

用途：防止行动无限堆积。

### 代谢

```text
无
巡检
排泄
拍摄
审计
喷射
栈桥候选
```

用途：系统排泄 / 拍摄 / 审计 / 喷射，不抢近期注意力前台。

## 4. 第一批已写入 Project 1

| Issue | 早晨列 | 晚上列 | 水位 | 火 | 出口 | 代谢 | 地界 |
|---|---|---|---|---|---|---|---|
| #190 | 不进早上 | 待裁决 | 候补 | 温 | 未到出口 | 无 | Notion / GitHub |
| #191 | 今天占位 | 谷 WIP | 今日 | 温 | 未到出口 | 拍摄 | PR / 沃壤 |
| #192 | 交给克 | AI 可推进 | 今日 | 温 | 未到出口 | 无 | WQB / 平台 |
| #193 | 不进早上 | AI 可推进 | 候补 | 温 | 未到出口 | 无 | 资源 |
| #194 | 交给克 | AI 可推进 | 今日 | 温 | 未到出口 | 无 | AI 共台 |
| #195 | 不进早上 | 暂留 | 水下 | 冷 | 未到出口 | 无 | Agent |
| #196 | 不进早上 | AI 可推进 | 候补 | 温 | 未到出口 | 排泄 | Notion / GitHub |
| #197 | 不进早上 | 暂留 | 候补 | 温 | 未到出口 | 无 | Notion / GitHub |
| #200 | 不进早上 | 完成·退相干 | 水下 | 冷 | 沉 | 排泄 | PR / 沃壤 |

上表是 2026-06-13 的历史首批写入，不代表这些对象今天仍处于同一状态。

## 5. Project 1 建议视图

### ☀️ 今日开工

```text
Layout: Board
Filter: 水位 = 今日
Group by: 早晨列
Fields: 接续之问 / 火 / 给克一句话
```

### 🌙 晚上裁决

```text
Layout: Board
Group by: 晚上列
Fields: 水位 / 早晨列 / 火 / 晚上判断 / 给克一句话
```

### 🤖 克可推进

```text
Layout: Table
Filter: 晚上列 = AI 可推进
Fields: 接续之问 / 给克一句话 / 晚上判断
```

### 🧯 代谢台

```text
Layout: Board
Filter: 代谢 != 无
Group by: 代谢
Fields: 晚上列 / 水位 / 晚上判断
```

### 🪦 完成退相干

```text
Layout: Board
Filter: 晚上列 = 完成·退相干 OR 出口 != 未到出口
Group by: 出口
Fields: 晚上判断 / Last touched
```

### 🧭 全量总账

```text
Layout: Table
Fields: 早晨列 / 晚上列 / 水位 / 火 / 接续之问 / 给克一句话 / 晚上判断 / 出口 / 代谢 / Last touched / 地界
Sort: Last touched ascending
```

人的默认入口不应是全量总账，而是少量“待裁决 + 今日”、验收与异常对象。全量视图留给机器、审计和主动回看。

## 6. Project 2 的归位

Project 2：<https://github.com/orgs/dongxi-heji/projects/2>

角色：晚上裁决样机。

它验证了晚上六列与字段写入流程。长期事实以 Project 1 为准；若没有继续实验价值，可以保留为历史参考，不再要求日常维护。

## 7. 2026-07-02 校准后的验收

```text
Notion：只保留公开窗口，不再承担早晨前台或行动总账。
GitHub Project 1：全天行动总账与人的精简注意力入口。
GitHub Project 2：晚上裁决样机 / 历史参考。
```

一句话：**GitHub Project 1 保存全量行动事实，但人的默认视野只显示今日待裁决、验收与异常；Notion 不再参与内部调度。**
