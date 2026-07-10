# CN-Mirror-CLI 沉积说明 · 2026-07-07

## 来源

- 来源仓库：`TaowuZhang/CN-Mirror-CLI`
- 目的：把长期有用的仓库维护经验沉到 `worang`，让外部仓库不再承担事实母本角色。

## 可保留经验

1. 写入 GitHub 前，应确认临时镜像或 rewrite 没有污染 push 路径。
2. 网络代理只属于运行态，不应写入长期仓库配置。
3. 镜像测速结果属于 cache，不属于长期事实；长期保留的是测速方法、边界和风险提示。
4. 示例配置可以留存；真实本地配置、token、header、私有 prefix 不迁移。
5. `mirrors.conf`、`xget.conf`、本地个人说明应保持 ignored 或只留在本地。

## 处置判断

| 内容 | 处理 |
|---|---|
| GitHub 镜像与下载加速经验 | 已在本页摘要保留 |
| 安全发布边界 | 已在本页摘要保留 |
| 私密配置 | 不迁移 |
| 一次性测速结果 | 不迁移 |
| 完整工具实现 | 不迁移；等待本地 Codex 判断 archive / reference |

## 清理影响

`TaowuZhang/CN-Mirror-CLI` 的长期经验已经在 `worang` 中留根。若本地 Codex 后续确认没有其他唯一资产、外部依赖、release/pages/workflow 约束，则该仓库可以进入 `migrate-then-archive` 路线。

## 仍需本地 Codex 补证

- root tree；
- workflow 状态；
- release / package / pages 状态；
- 是否有未迁移的唯一脚本；
- 是否存在本地 dirty change；
- archive 或 delete 的最终人工闸门。
