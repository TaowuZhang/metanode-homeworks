# 豆瓣 mirror v0.1 / v0.2 种子切片

日期：2026-06-03
位置：资源库自动化候选 / 豆瓣同步区

## 本刀目标

先把 v0 的全局 `--limit` 改成按类型限额，再把脚本压成一个小而灵活的“导入归一器”：既能吃旧 Markdown，也能吃浏览器 / 油猴导出的豆瓣 CSV。

## v0.1 改动

- `scripts/sync-douban.mjs` 新增 `--limit-per-type`；
- `--limit` 保留为兼容别名，但语义改为每类限额；
- 环境变量优先读 `DOUBAN_SYNC_LIMIT_PER_TYPE`，兼容 `DOUBAN_SYNC_LIMIT`；
- 收集逻辑改为每个类型内部按 `rated_at` 倒序排序后 slice，再合并 manifest；
- 输出日志改为 `limitPerType`；
- README 使用示例改为 `--limit-per-type 20`。

## v0.2 改动

- 新增 `--source csv`；
- 新增 `--input <file-or-directory>` / `DOUBAN_SYNC_INPUT`；
- 新增 `--type movie|book|music`，用于单个 CSV 文件无法从文件名判断类型时；
- 目录导入时按文件名识别 `movie.csv` / `book.csv` / `music.csv`；
- 新增内置轻量 CSV parser，支持引号、逗号、换行短评；
- 新增 `csv_import` source_mode；
- `manifest.schema.json` 允许 `csv_import`；
- 输出仍只写 `资源/_douban/manifest.jsonl` 与 `*/recent/*.md`。

## 字段映射

CSV 列名保持朴素兼容：

```plain text
标题 / title / 名称 / 片名 / 书名 / 专辑 -> title
个人评分 / 评分 / rating -> personal_rating
打分日期 / 标记日期 / 日期 / rated_at / date -> rated_at
我的短评 / 短评 / 评语 / comment / review -> short_comment
条目链接 / 豆瓣链接 / 链接 / url / link -> douban_url / douban_id
上映日期 / 发行日期 / release_date -> release_date
出版日期 / publication_date -> publication_date
作者 / author -> author
音乐家 / 艺术家 / artist -> artist
制片国家 / 制片国家/地区 / 国家 / 地区 -> country_or_region
```

## 本地烟测：existing-markdown

在 `/data/douban-v01` 用每类 3 条 fixture 运行：

```plain text
node scripts/sync-douban.mjs --source existing-markdown --limit-per-type 2 --dry-run
node scripts/sync-douban.mjs --source existing-markdown --limit-per-type 2
```

结果：

```plain text
Douban mirror v0.1: source=existing-markdown rows=6 limitPerType=2 dryRun=true
{
  "movie": 2,
  "book": 2,
  "music": 2
}

Douban mirror v0.1: source=existing-markdown rows=6 limitPerType=2 dryRun=false
{
  "movie": 2,
  "book": 2,
  "music": 2
}
```

`manifest.jsonl` 结果为 6 行：movie 2、book 2、music 2。

## 本地烟测：csv

在 `/data/douban-v02` 用 `资源/_douban/imports/movie.csv`、`book.csv`、`music.csv` fixture 运行：

```plain text
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 2 --dry-run
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 2
```

结果：

```plain text
Douban mirror v0.2: source=csv rows=6 limitPerType=2 dryRun=true
{
  "movie": 2,
  "book": 2,
  "music": 2
}

Douban mirror v0.2: source=csv rows=6 limitPerType=2 dryRun=false
{
  "movie": 2,
  "book": 2,
  "music": 2
}
```

`manifest.jsonl` 结果为 6 行：movie 2、book 2、music 2。额外检查：CSV 中带换行的 `我的短评` 可正确进入生成 Markdown。

## 外部做法速查

本轮按“研究”入口查了几个外部豆瓣导出 / 备份项目，结论不是要接一个大框架，而是确认最小可运转形态。

### 看到的稳定模式

- `bambooom/douban-backup`：油猴脚本导出 CSV，样例字段与当前沃壤旧 Markdown 高度一致：`标题`、`个人评分`、`打分日期`、`我的短评`、`上映日期`、`制片国家`、`条目链接`。它也提到 RSS 很轻，但每次只保留最近 10 条，而且会混入想看 / 想听 / 想读。
- `UlyC/DouBanExport`：也是浏览器脚本导出；核心经验是不要爬条目详情页，优先读个人标记列表页，降低反爬与封号风险。
- `hqweay/douban-getter`：目标明确是备份用户标记，而不是爬豆瓣条目详情；抽象成三步：取列表页、解析、保存本地。它提到列表页一页约 15 条，并建议请求间隔不要太快。
- `einverne/douban-export`：提供 movie / book / music 三类导出命令，说明三类分开处理是常见做法。

### 对沃壤脚本的判断

最简单、最稳的下一步不是马上做 GitHub Action 自动登录豆瓣，也不是把 RSS、列表页爬虫、浏览器插件、Notion 同步揉成一个系统。

更合适的最小形态是：

```plain text
手动 / 浏览器导出 CSV 或既有 Markdown
→ sync-douban.mjs 归一化字段
→ 写入 资源/_douban/manifest.jsonl + recent/*.md
```

也就是让 `scripts/sync-douban.mjs` 只做一个小而灵活的“导入归一器”：输入可以换，输出稳定；不碰账号、不写回豆瓣、不做大同步平台。

## 边界

本刀仍然不做：

- 不写回豆瓣；
- 不接账号密钥；
- 不写 `.github/workflows/*`；
- 不改旧 `资源/电影/`、`资源/书籍/`、`资源/音乐/`；
- 不删除旧文件；
- 不把 Apple Music / 微信读书 / 得到大脑纳入本轮。

## 下一步

如果要生成真实 `_douban` 种子，优先在可 checkout 的环境运行：

```plain text
node scripts/sync-douban.mjs --source existing-markdown --limit-per-type 20
```

如果要从豆瓣导出的 CSV 生成镜像，先把文件放入：

```plain text
资源/_douban/imports/movie.csv
资源/_douban/imports/book.csv
资源/_douban/imports/music.csv
```

再运行：

```plain text
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 20 --dry-run
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 20
```

生成输出仍然只提交：

```plain text
资源/_douban/manifest.jsonl
资源/_douban/movie/recent/*.md
资源/_douban/book/recent/*.md
资源/_douban/music/recent/*.md
```
