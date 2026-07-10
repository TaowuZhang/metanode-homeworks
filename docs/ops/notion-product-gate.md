# Notion 产品门与公开窗口

日期：2026-07-02

状态：

- **Notion AI 付款：HOLD**
- **自动搬运与 Cloudflare / GitHub App 扩建：HOLD**
- **Education Plus 公开窗口：ACTIVE**
- **GitHub / worang：唯一版本化母本**

## 账户事实

当前工作区是用户过去学生资格获得的单成员 **Education Plus**。它提供个人 Plus 能力，不等于 Business，也不自带可依赖的无限 Notion AI。

以后判断 Notion 产品能力时，必须以这个真实账户为基线，不把普通 Free、Business 或单独购买 AI 的说明混进来。

## 当前分工

```text
GitHub / worang
  = 事实、规则、技能、项目状态、作品母稿、公开文案源码、Issue、PR、版本历史

Notion Education Plus
  = 少量公开页面、手工排版、轻量交互陈列
```

Notion 不再承担：

- 内部协议、三本书或 AI 协作规则的唯一母本；
- GitHub Issue / PR / Project 镜像；
- 技能调度或项目总账；
- 每日投递、自动沉淀与第二套事实层；
- 为了维持旧页面树而继续付费或建设基础设施。

## 公开窗口的母本

Notion 公开页面的文案和结构保存在：

- `web/content/notion-site.md`
- `web/content/profile.json`

Notion 可以改变视觉排版、分栏和交互呈现，但重要文字与 HTML 源码不能只存在于 Notion。

当前保留的最小页面集合：

1. `Notion · 轻页面`：私有薄后台，只说明 GitHub 是母本并指向公开页；
2. `棱窗 · 公开页面`：公开页面的薄父页 / 管理入口；
3. `你好，欢迎来到这里`：公开首页；
4. `第 0 期外显入口 v1`：当前保留的一件公开展品。

除此之外，旧星盘、公约、腐海、暗房、驿港、迁移记录、压测工作台、旧数据库和复制的外部资料空间不再作为 Notion 现行器官。

## Interactive HTML 一手观察

2026-07-02，用户在实际 Education Plus 工作区的区块菜单中直接看到 `HTML · 嵌入`。因此当前可以确认：

- 手动 HTML 插入在该账户中可用；
- HTML 可以由 ChatGPT、Codex 或本地编辑器生成，再放入 Notion；
- 这条基础展示路径不依赖 Notion Agent，也不依赖购买 Notion AI。

仍需一次直接 smoke test 才能声称公开运行成立：

- 发布一个带简单交互的 HTML 区块；
- 用未登录窗口和手机打开；
- 验证交互是否仍然可用。

在验证前，不声明 HTML 区块内部动态内容一定能被搜索引擎索引。

## 为什么暂停 Notion AI 付款

Notion AI 的模型级额度、重置周期、降级 / 回退方式和可用性承诺不足以支撑付费依赖。功能吸引力不能替代容量和行为契约。

重新考虑付款前，至少需要能回答：

- 实际账户可用哪些模型；
- 每类限制怎样计数、何时重置；
- 达限后是拒绝、冷却、切模型还是降级；
- 重试是否重复计数；
- 长文写作与修改能否在真实节奏下持续完成；
- 是否可以先按月测试，而不是先承担年度承诺。

## 保留的逃生工具

PR #345 / #346 已留下手动单页 Notion 导出 CLI、测试与真实金丝雀。它继续作为审计和一次性撤离工具，而不是恢复自动化工程的理由。

当前不部署：

- Notion webhook；
- Cloudflare Worker / Queue / Workflow / D1；
- GitHub App；
- 私有 export fork；
- 自动 PR 或自动 landing；
- 新的 Notion 状态数据库。

## 维护纪律

- 新的内部知识直接进 GitHub，不先造 Notion 容器。
- 新公开页先在 GitHub 留母稿，再手工陈列到 Notion。
- Notion 页面失去公开用途后，先核对 GitHub 替代物，再放入 Trash。
- 不为保留历史导航而保存空壳；需要历史时看 Git 与 Notion Trash。
- 不把大批外部课程或别人的空间全量复制进 Notion；只保留轻索引，消化后再进入沃壤。
