# 豆瓣 mirror v0 烟测记录

日期：2026-06-03  
分支：`feat/douban-mirror-v0`

## 烟测目标

验证 `scripts/sync-douban.mjs` 的 v0 种子模式是否能在不碰旧资源目录的前提下：

- 从 `资源/电影/`、`资源/书籍/`、`资源/音乐/` 解析豆瓣字段；
- 识别电影 / 书籍 / 音乐三类豆瓣 URL；
- 生成 `资源/_douban/*/recent/*.md`；
- 生成 `资源/_douban/manifest.jsonl`；
- 保留多行 `我的短评`；
- 将 `@2021年1月1日` 这类出版日期去掉前导 `@`；
- 只写入 `资源/_douban/`。

## 本地夹具

本地 sandbox 建了三个最小样本：

```plain text
资源/电影/测试电影.md
资源/书籍/测试书.md
资源/音乐/测试专辑.md
```

字段覆盖：

- 电影：标题、上映日期、个人评分、制片国家、打分日期、条目链接；
- 书籍：标题、个人评分、作者、出版日期、两行短评、打分日期、条目链接；
- 音乐：标题、个人评分、发行日期、打分日期、条目链接、音乐家。

## 命令

```plain text
node scripts/sync-douban.mjs --source existing-markdown --limit 20 --dry-run
node scripts/sync-douban.mjs --source existing-markdown --limit 20
```

## 结果

Dry-run 输出：

```json
{
  "movie": 1,
  "book": 1,
  "music": 1
}
```

实际写入生成：

```plain text
资源/_douban/book/recent/2345678-测试书.md
资源/_douban/manifest.jsonl
资源/_douban/movie/recent/1234567-测试电影 - Test Movie.md
资源/_douban/music/recent/3456789-测试专辑.md
```

`manifest.jsonl` 三行均包含：

- `source: douban`
- `source_mode: existing_markdown_seed`
- `type`
- `douban_id`
- `douban_url`
- `title`
- `personal_rating`
- `rated_at`
- `original_path`
- `path`
- `raw_hash`
- `synced_at`

书籍样本验证通过：多行短评被保留为：

```plain text
第一行短评
第二行短评
```

## 未完成项

`.github/workflows/sync-douban.yml` 未能写入。GitHub 返回：

```plain text
403 Resource not accessible by integration
```

判断：这是 workflow 路径的特殊写入权限缺口，不影响普通文件写入和脚本主体。后续若要加 GitHub Actions，需要补 workflow scope / 由本地或有 workflow 权限的通道提交。

## 当前结论

v0 种子脚本烟测通过。下一刀可以选择：

1. 在真实仓库上运行 `--dry-run`；
2. 手动补 workflow；
3. 开始设计 `douban-account` 数据源探针。
