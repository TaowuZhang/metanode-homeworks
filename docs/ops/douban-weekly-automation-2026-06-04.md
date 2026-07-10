# 豆瓣账号每周同步自动化

日期：2026-06-04
位置：`.github/workflows/sync-douban.yml`

## 意图

豆瓣账号 mirror 进入稳定的每周自动同步节奏：

```plain text
每周一 04:00 Asia/Shanghai
→ 只读读取豆瓣账号「我看 / 我读 / 我听」
→ dry-run
→ real-run
→ 审计
→ 有实质变化才开 PR
```

GitHub Actions 使用 UTC，因此实际 cron 是：

```plain text
0 20 * * 0
```

也就是每周日 20:00 UTC，对应每周一 04:00 东八区。

## 前置条件

自动化建立在 no-op churn 防护之后：重复同步如果没有稳定字段变化，不应因为 `synced_at` 变化重写 manifest 和 60 个 Markdown 文件。

已验证的目标输出：

```plain text
Douban mirror reconcile: unchanged=60 changed=0
Douban mirror writes: files=0 manifest=0
```

## 凭据

Workflow 使用 GitHub Secret，不从仓库读取 cookie 文件：

```plain text
DOUBAN_COOKIE
```

当前默认值：

```plain text
DOUBAN_USER_ID=158583636
DOUBAN_SYNC_LIMIT=20
DOUBAN_USER_AGENT=Mozilla/5.0 DoubanMirrorSync/0.3
```

边界：

- 不把 cookie / token / `.env` 写入仓库；
- 不把密钥贴到聊天；
- 不写回豆瓣；
- 不提交 `cookies/`。

## 审计闸门

Workflow 会执行：

1. `node --check scripts/sync-douban.mjs`；
2. `--dry-run`；
3. 真实同步；
4. manifest 审计：
   - 三类数量符合 limit；
   - 所有行都是 `source_mode: douban_account`；
   - `douban_id` / `title` / `personal_rating` / `rated_at` / `douban_url` / `path` 不为空；
   - 每个 manifest path 都存在。

审计失败时 workflow 失败，不应创建同步 PR。

## 写入模型

Workflow 不直接提交到 `main`，而是使用固定分支创建/更新 PR：

```plain text
chore/douban-account-mirror-sync
```

如果没有文件变化，`create-pull-request` 不会创建新 PR。

## 来源边界

豆瓣 mirror 只表示豆瓣账号里的书影音标记 / 评分 / 条目记录。

特别是书籍：

```plain text
得到 = 读书 / 学习主源
豆瓣书籍 = 外部条目、评分、标记、metadata
微信读书 = 当前不作为读书源
```

因此该 workflow 不处理得到、微信读书、Apple Music，也不把它们混入豆瓣 mirror。
