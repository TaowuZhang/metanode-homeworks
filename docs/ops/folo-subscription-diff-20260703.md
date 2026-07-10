# Folo 路由审查差异记录（2026-07-03）

## 当前实际结果

- 原始订阅：45
- 当前订阅：27
- 原始 RSSHub：33
- 当前 RSSHub：15
- 原始直接 RSS：10
- 当前直接 RSS：10
- 第三方桥接：2（未变化）

## 已完成的直接源替换

### Hacker News

- 删除：`https://rsshub.app/hackernews`
- 新增并验证：`https://news.ycombinator.com/rss`

### NASA APOD

- 删除：`https://rsshub.app/nasa/apod`
- 新增并验证：`https://apod.nasa.gov/apod.rss`

## 第二波精确差异

从 33 个订阅降到 27 个，只删除了以下 6 项：

1. Lifehacker
2. Cyber Security News
3. 领研 论文 计算机
4. 中国气象局 每日天气提示
5. 东方财富网 策略报告
6. 财经日历 华尔街见闻

新增来源：0。

## 总体结果

从原始快照到当前快照：

- 总订阅：45 → 27
- RSSHub：33 → 15
- 直接 RSS：10 → 10
- 第三方桥接：2 → 2

清理阶段已经完成。下一阶段转为分类重组和观察，不再继续无差别删除。