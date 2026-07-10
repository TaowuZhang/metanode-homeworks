# 豆瓣 CSV 导入闭环

日期：2026-06-04
位置：`资源/_douban/` 手动同步入口

## 本刀目标

在第一批真实 `_douban` seed 已进入 `main` 之后，把豆瓣的最小继续同步路径写清楚：

```plain text
豆瓣导出 CSV
→ 放入 资源/_douban/imports/
→ dry-run 审计
→ 真实生成 _douban manifest + recent
→ 只提交 _douban 生成物
```

这不是账号自动同步，也不是音乐 / 得到 / 微信读书的统一桥；只是把「谷在豆瓣记录过的东西，至少能继续同步到沃壤」这件事先做稳。

## 当前状态

已完成：

- `scripts/sync-douban.mjs` 已在 `main`。
- `资源/_douban/README.md` 与 `manifest.schema.json` 已在 `main`。
- 第一批真实 seed 已合入 `main`：
  - `资源/_douban/manifest.jsonl`：60 行；
  - `资源/_douban/movie/recent/`：20 条；
  - `资源/_douban/book/recent/`：20 条；
  - `资源/_douban/music/recent/`：20 条。

## 输入位置

CSV 手动导入文件放在：

```plain text
资源/_douban/imports/movie.csv
资源/_douban/imports/book.csv
资源/_douban/imports/music.csv
```

文件名用于类型识别：

- `movie.csv` → 豆瓣电影；
- `book.csv` → 豆瓣书籍；
- `music.csv` → 豆瓣音乐。

若只有单个 CSV，文件名又无法识别类型，则显式指定：

```plain text
node scripts/sync-douban.mjs --source csv --input path/to/file.csv --type movie --limit-per-type 20
node scripts/sync-douban.mjs --source csv --input path/to/file.csv --type book --limit-per-type 20
node scripts/sync-douban.mjs --source csv --input path/to/file.csv --type music --limit-per-type 20
```

## 运行方式

先 dry-run：

```plain text
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 20 --dry-run
```

确认输出数量、字段和样本后，再真实生成：

```plain text
node scripts/sync-douban.mjs --source csv --input 资源/_douban/imports --limit-per-type 20
```

默认每类 20 条。需要扩大时，先只扩大 `--limit-per-type`，不要同时改脚本、workflow 或账号接入。

## 最小审计

生成后先看数量：

```plain text
wc -l 资源/_douban/manifest.jsonl
find 资源/_douban/movie/recent -type f | wc -l
find 资源/_douban/book/recent -type f | wc -l
find 资源/_douban/music/recent -type f | wc -l
```

再看字段缺口：

```plain text
node - <<'NODE'
const fs = require("fs");
const rows = fs.readFileSync("资源/_douban/manifest.jsonl", "utf8").trim().split("\n").map(JSON.parse);
const byType = {};
for (const r of rows) {
  byType[r.type] ??= [];
  byType[r.type].push(r);
}
for (const [type, items] of Object.entries(byType)) {
  const missing = key => items.filter(r => r[key] == null || r[key] === "").length;
  console.log(type, {
    count: items.length,
    missing_douban_id: missing("douban_id"),
    missing_title: missing("title"),
    missing_rating: missing("personal_rating"),
    missing_rated_at: missing("rated_at"),
    missing_url: missing("douban_url"),
    missing_release_date: missing("release_date"),
    missing_publication_date: missing("publication_date"),
    missing_author: missing("author"),
    missing_artist: missing("artist"),
    missing_country_or_region: missing("country_or_region"),
  });
}
NODE
```

核心门槛：

- 三类的 `douban_id` 必须 0 缺失；
- 三类的 `title` 必须 0 缺失；
- 三类的 `personal_rating` / `rated_at` / `douban_url` 应尽量 0 缺失；
- 电影优先看 `release_date` / `country_or_region`；
- 书籍优先看 `publication_date` / `author`；
- 音乐优先看 `release_date` / `artist`。

## 提交范围

CSV 导入生成后，原则上只提交：

```plain text
资源/_douban/manifest.jsonl
资源/_douban/movie/recent/*.md
资源/_douban/book/recent/*.md
资源/_douban/music/recent/*.md
```

`imports/*.csv` 是否提交需另判：默认不提交，除非它本身被视为可公开 / 可版本化的原始导入材料。

## 边界

本闭环不做：

- 不写回豆瓣；
- 不接账号、cookie 或 token；
- 不要求把密钥放进聊天；
- 不碰 `.github/workflows/*`；
- 不改旧 `资源/电影/`、`资源/书籍/`、`资源/音乐/`；
- 不把 Apple Music、微信读书、得到大脑混进豆瓣 mirror；
- 不把豆瓣音乐记录等同于真实听歌史；
- 不把得到读书笔记等同于豆瓣书籍评价。

## 后续桥位

后续可以另开：

1. **豆瓣账号同步**：研究是否需要本地环境变量 / GitHub Secrets / 手动触发，不在聊天中暴露密钥。
2. **音乐桥位**：区分豆瓣音乐评价与 Apple Music / 其他听歌行为。
3. **读书桥位**：区分豆瓣书籍条目、微信读书行为、得到课程 / 笔记。
4. **字段清洗**：例如电影 `country_or_region` 中的电影节 / 首播集脏值。

这些都不属于当前 CSV 手动导入闭环。
