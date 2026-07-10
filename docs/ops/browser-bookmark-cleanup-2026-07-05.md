# 浏览器书签清理执行与核验｜2026-07-05

本页属于一次性执行、审计和可见界面核验记录，因此居留在 `docs/ops/`。

## 完成标准

1. 用户实际界面符合目标状态；
2. 浏览器重启后状态不恢复；
3. 不误伤标签组、密码、Cookies、历史、扩展和其他非目标数据；
4. GitHub 记录与当前界面一致。

磁盘文件、自动化日志和节点计数只能作为辅助证据，不能覆盖用户实际看到的界面。

## Chrome

最终可见状态：

- 旧书签文件夹已经全部消失；
- 书签栏只剩 `工具架`；
- `工具架` 中只有 Photopea 与 Apifox；
- `Myself / Job / Network / WQB / Leetcode / Hack` 保存标签组保持不变；
- 没有把旧 Chrome 书签重新迁入 Zen。

此前声称“274 → 2”但未核验界面的自动化报告不作为完成依据；最终完成以用户提供的可见界面为准。

## Zen

活动 profile 为 `Default (release)`。通过正式界面导入 `工具架 v0`，核验 8 条：

- PDF24 Tools；
- Regex Vis；
- JSON Crack；
- Draw.io；
- Excalidraw；
- iconfont；
- Pexels；
- Freesound。

未操作另一个 Zen profile，也未直接修改 `places.sqlite`。

## Safari

旧导出包含 303 条书签与 22 条阅读列表。历史由 Safari ZIP 与 HTML 导出保存。

2026-07-05 用户已完成界面清空：

- 用户书签：0；
- 阅读列表：0；
- 自定义书签文件夹：0。

Safari 的长期角色为不承载收藏体系的干净隔离浏览器。

## 结果

- Chrome：完成；
- Zen 工具架：完成；
- Safari：完成。

长期浏览器分工见 `资源/语境/浏览器与链接路由.md`；工具使用语境见 `资源/语境/浏览器工具架.md`。
