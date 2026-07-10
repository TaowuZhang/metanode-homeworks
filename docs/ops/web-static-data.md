# Web 静态数据规格

Web 第一版只做外壳，不做第二套知识库。

## 数据文件

建议先用：

- `web/content/profile.json`
- `web/content/pulse.json`
- `web/content/workbench.json`
- `web/content/cabinet.json`
- `web/content/dispatch.json`

## 字段规范与 Mock 样例

### `profile.json`
负责顶部 Hero 区及联系方式。
```json
{
  "name": "稻谷",
  "tagline": "拆解系统、构建系统、用数据验证系统",
  "intro": "这不是简历。这是一个打开的好奇心陈列柜和工作台。北邮电信工程本科 → 悉尼大学网络安全硕士 → 正在探索网络安全、编程与量化金融。",
  "links": {
    "github": "https://github.com/dongxi-heji",
    "email": "mailto:1171964523@qq.com"
  }
}
```

### `pulse.json`
负责生命体征与当前重心。
```json
{
  "metrics": [
    { "label": "晨跑签到", "value": "2768 次", "note": "7 年+" },
    { "label": "渗透靶机", "value": "12 个", "note": "100% Root" }
  ],
  "currentDirections": [
    "网络安全：渗透测试全链路",
    "编程系统构建：算法与后端",
    "量化信号验证：WorldQuant 因子设计"
  ],
  "narrative": "同一种思维方式的三个出口。"
}
```

### `workbench.json`
进行中的工程/项目卡片。
```json
[
  {
    "title": "RenoPilot.JS.Shapes2",
    "kind": "构建 · 编程",
    "status": "In Progress",
    "summary": "独立负责 94% 核心代码（10.5 万行），获 'Excellent Design' 评价。",
    "href": "#"
  },
  {
    "title": "渗透测试靶机攻防",
    "kind": "拆解 · 网安",
    "status": "In Progress",
    "summary": "信息收集 → 漏洞利用 → 权限维持 的全链路笔记。",
    "href": "#"
  }
]
```

### `cabinet.json`
决定性坐标故事（Teaser → Depth）。
```json
[
  {
    "title": "问心无愧天地宽",
    "hook": "在代写泛滥的环境中，过一种不需要观众的生活。",
    "bodyRef": "我没有对抗什么。我只是没有滑过去。每天做饭、上课、写作业...",
    "tags": ["独立", "选择"]
  },
  {
    "title": "半江瑟瑟半江红",
    "hook": "从文字开始认识坂本龙一，理解自然和环境不一定要配音乐。",
    "bodyRef": "他为末代皇帝作曲，在商业和艺术之间找到了位置...",
    "tags": ["感知", "音乐"]
  }
]
```

### `dispatch.json`
排泄/发布的历史切片。
```json
[
  {
    "title": "内容界面与认知界面的载体分裂",
    "date": "2026-05-07",
    "channel": "Notebook",
    "href": "#",
    "excerpt": "列表式展示回答\"有什么\"，力导向图回答\"它们之间是什么\"..."
  }
]
```

## 原则

本地 / Notion 摘要生成 JSON；网页只消费 JSON。  
网页不回写 Notion，不回写 events，不接受社媒反馈改结构。  
第一版可用 Astro / Vite / Next 任一静态路线；未拍板前不建框架。

## 待补缺口

- [x] 提供以上各个 JSON 的结构定义及 mock 数据。
- [ ] 根据后续 Web 框架（如 Astro）的具体情况，确认这些 JSON 是打包进项目还是通过运行时拉取。