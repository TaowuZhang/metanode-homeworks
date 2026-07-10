# Folo 路由审查与操作清单 v2

日期：2026-07-03

> 状态：两轮清理和直接 RSS 替换均已完成；当前进入分类重组。

## 当前实际状态

- 总订阅：27
- RSSHub：15
- 直接 RSS：10
- 第三方桥接：2

## 已完成

1. 第一波公共雷达清理。
2. Hacker News 与 NASA APOD 改为直接 RSS。
3. 第二波查询型、阶段型与重复快讯清理。
4. 最新 OPML 验证没有误删或额外变化。

## 当前操作：只移动，不删除

在 Folo 中创建：

1. `01 必看触发`
2. `02 深度阅读`
3. `03 漫游`
4. `04 行动提醒`
5. `05 观察期`

完整 27 项移动清单保存在本地文件：

- `folo-分类重组清单-2026-07-03.md`
- `folo-分类重组清单-2026-07-03.csv`

这一阶段不再导入 OPML，以免把已有订阅的分类移动与新增订阅混在一起。

## 版本化名称清单

### 01 必看触发

- Google Developers Blog
- Google DeepMind News
- Obsidian Plugins

### 02 深度阅读

- 小众软件
- 晚点 - 长报道
- Last Week in AI
- darkreading
- Krebs on Security
- Steph Ango
- 有知有行

### 03 漫游

- Nat Geo Photo of the Day
- NASA APOD
- Magnum Photos
- 中国爬楼联盟
- 街拍中国

### 04 行动提醒

本组 4 个来源的具体名称当前只保留在本地清单，版本化记录保留数量与用途。

### 05 观察期

- Stack Overflow Blog
- Google AI Developers
- Logan Kilpatrick
- News Minimalist
- Hacker News

另有 2 个观察来源的名称当前只保留在本地清单。

## 观察期规则

两周后检查：

1. 中国爬楼联盟与街拍中国最多保留一个。
2. Google AI Developers 与 Google Developers Blog 重复明显时保留内容更完整的一边。
3. Stack Overflow Blog、News Minimalist、Logan Kilpatrick、Hacker News 没有真实打开或独特价值就退出常驻。
4. 行动提醒只有在确实触发领取、观看或使用时才继续保留。

## 边界

- 当前不再继续删除。
- 不要求清空未读。
- 7 个来源名称当前省略；是否写入版本化账本不作为去留条件。
- 分类是注意力用途，不是价值等级。