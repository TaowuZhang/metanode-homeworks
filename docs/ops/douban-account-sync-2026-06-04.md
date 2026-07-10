# 豆瓣账号同步闭环

日期：2026-06-04
位置：`scripts/sync-douban.mjs --source douban-account`

## 本刀目标

在第一批 `_douban` seed 和 CSV 手动导入闭环之后，补上豆瓣账号的**只读同步入口**：

```plain text
本地 cookie / cookie 文件
→ 读取豆瓣账号「我看 / 我读 / 我听」
→ dry-run 审计
→ 真实生成 _douban manifest + recent
→ 只提交 _douban 生成物
```

这一步确认「账号同步是得做的」，但仍不做自动 workflow、不写回豆瓣、不把密钥放进聊天。

## 凭据边界

脚本只从本地环境读取凭据：

```plain text
DOUBAN_COOKIE_FILE=cookies/douban.txt
```

或：

```plain text
DOUBAN_COOKIE='...'
```

推荐 `DOUBAN_COOKIE_FILE`，因为它避免把 cookie 放进 shell history。

仓库已有 `.gitignore`：

```plain text
cookies/
.env
.env.*
!.env.example
```

因此本地 cookie 文件和 `.env` 不应进入提交。

## Dry-run

```plain text
DOUBAN_COOKIE_FILE=cookies/douban.txt \
node scripts/sync-douban.mjs --source douban-account --input <douban-user-id> --limit-per-type 20 --dry-run
```

只跑单类：

```plain text
DOUBAN_COOKIE_FILE=cookies/douban.txt \
node scripts/sync-douban.mjs --source douban-account --input <douban-user-id> --type movie --limit-per-type 20 --dry-run

DOUBAN_COOKIE_FILE=cookies/douban.txt \
node scripts/sync-douban.mjs --source douban-account --input <douban-user-id> --type book --limit-per-type 20 --dry-run

DOUBAN_COOKIE_FILE=cookies/douban.txt \
node scripts/sync-douban.mjs --source douban-account --input <douban-user-id> --type music --limit-per-type 20 --dry-run
```

## 真实生成

```plain text
DOUBAN_COOKIE_FILE=cookies/douban.txt \
node scripts/sync-douban.mjs --source douban-account --input <douban-user-id> --limit-per-type 20
```

生成范围仍然只在：

```plain text
资源/_douban/manifest.jsonl
资源/_douban/movie/recent/*.md
资源/_douban/book/recent/*.md
资源/_douban/music/recent/*.md
```

## 审计

数量审计：

```plain text
wc -l 资源/_douban/manifest.jsonl
find 资源/_douban/movie/recent -type f | wc -l
find 资源/_douban/book/recent -type f | wc -l
find 资源/_douban/music/recent -type f | wc -l
```

字段缺口审计仍沿用 CSV 闭环文档里的脚本。

账号同步的第一优先级：

- `douban_id` 稳定；
- `title` 正常；
- `douban_url` 完整；
- `personal_rating` 能从豆瓣星级样式读出；
- `rated_at` 能从收藏页日期读出；
- `source_mode` 为 `douban_account`。

第二优先级才是类型字段：

- 电影：`release_date` / `country_or_region`；
- 书籍：`publication_date` / `author`；
- 音乐：`release_date` / `artist`。

豆瓣页面结构若变动，先以 dry-run 样本判断是否需要修 parser，不要直接大批量真实生成。

## 已知限制

- 这不是官方 API；它读取的是豆瓣账号收藏页 HTML。
- Cookie 可能过期；过期时脚本会失败，不会写空 mirror。
- 页面结构变化会影响字段完整度。
- 音乐条目的字段仍可能比电影 / 书籍稀薄。
- 账号同步只读，不会改豆瓣状态、标签、评分、短评。
- 不接 GitHub Actions；若未来接，也必须是手动触发 + secret，不定时自动跑。

## 仍然不做

- 不写回豆瓣；
- 不把 cookie / token / 密钥放进聊天；
- 不提交 `cookies/` 或 `.env`；
- 不碰 `.github/workflows/*`；
- 不改旧 `资源/电影/`、`资源/书籍/`、`资源/音乐/`；
- 不把 Apple Music、微信读书、得到大脑混进豆瓣 mirror；
- 不把豆瓣音乐记录当真实听歌史；
- 不把得到读书笔记当豆瓣书籍评价。