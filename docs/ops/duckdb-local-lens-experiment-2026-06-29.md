# DuckDB 本地只读透镜实验 · 2026-06-29

对应：Issue #312 `explore(cli): validate DuckDB as a local analysis lens`

## 结论

**ADOPT**：值得进入后续最小实现 PR，但本次只提交实验包，不实现正式接入。

结论门全部满足：7 个真实问题成功回答；时间聚合、跨 CSV schema 漂移和嵌套 JSON 连续性检查明显比临时 Node/Python 脚本更直接；中文路径与内容正常；JSON 可稳定交给 `jq`；混合日期没有被静默误判为 `DATE`，而是安全推断为 `VARCHAR` 后用 `try_cast` 显式处理；全程使用 `:memory:`，禁用扩展自动安装/加载，无 `.duckdb` 文件、网络扩展或 canonical 写回；7 个问题的整批运行耗时为 0.08–0.18 秒。

ADOPT 只表示可以另开后续最小实现 PR。今晚没有增加 doctor、正式 wrapper、mise、CI、`wo` 子命令或写回能力。

## 实验边界与环境

| 项目 | 实测值 |
|---|---|
| Git 基线 | `2da5a6a67ab7177acb2c0820a1312f5b89763e26` (`origin/main`) |
| 分支 | `experiment/duckdb-local-lens-20260629`，独立 worktree |
| macOS | 26.5.1 (Build 25F80), arm64 |
| DuckDB | v1.5.4 (Variegata) `08e34c447b` |
| DuckDB CLI 资产 | 官方 `duckdb_cli-osx-arm64.zip`；SHA-256 `d6c35195683fd1378e5624b01ca390069d399f8341c38986b7e3dfa0b3470d10` |
| jq | 1.8.1 |
| Node.js | v20.19.0 |
| Python | 3.14.4 |

本机最初没有 `duckdb`。`brew install duckdb` 两次都在下载 1.5.4 bottle 时因 `pkg-containers.githubusercontent.com` TLS `SSL_ERROR_SYSCALL` 失败；为避免把传输故障当成产品结论，改用 DuckDB 官方 GitHub Release 的同版本 arm64 CLI，下载到 `/private/tmp` 并与 release API 的 SHA-256 校验一致。没有安装或加载任何 DuckDB 扩展。

实现依据核对了 DuckDB 官方的 [CLI](https://duckdb.org/docs/current/clients/cli/overview.html)、[输出格式](https://duckdb.org/docs/current/clients/cli/output_formats.html)、[CSV 导入](https://duckdb.org/docs/current/data/csv/overview.html) 和 [JSON](https://duckdb.org/docs/current/data/json/overview.html) 文档。

## 数据源与敏感性

三个输入均由 `git ls-files --error-unmatch` 确认为 Git 跟踪文件；实验前先检查了表头、样本结构和敏感性。

| 路径 | 结构 | 选择理由 | 敏感性处理 |
|---|---|---|---|
| `成为/色散.csv` | 10 列 CSV，含混合日期与稀疏字段 | Issue #312 指定候选；能回答真实时间模式、空值与重复问题 | 含个人叙事，属于高敏感但非凭据的已跟踪 canonical 数据；SQL 和报告不返回正文，只返回聚合、日期范围和布尔中文检查 |
| `成为/探筹.csv` | 8 列 CSV，枚举较完整 | 能验证类别分布，并与色散做多 CSV `union_by_name` 和 schema 漂移检查 | 动作目录，无凭据、账户导出、邮件或聊天；报告只列枚举频率 |
| `wo/nlm/corpus/becoming-ideas/catalog.json` | 单个 JSON 对象，含标量和嵌套 `parts[]` 对象数组 | 是真实 NLM corpus 合同；可验证 JSON 展开、分片连续性和总数一致性 | 派生元数据；不含卡片正文或远端认证信息 |

未使用未提交文件、凭据、账户导出、邮件正文、私密聊天、HTTP/S3/云存储或外部服务。时间模式只对真实含时间字段的色散执行；探筹和 catalog 的时间分析为“不适用”。

runner 将三个显式输入复制到 `mktemp` 目录下的 `色散 样本.csv`、`探筹 样本.csv`、`NLM 目录 样本.json`，以同时验证中文和空格路径。退出 trap 清理目录；查询只建立临时 view，并显式设置：

```sql
SET autoinstall_known_extensions = false;
SET autoload_known_extensions = false;
```

## 七个真实问题

完整可执行 SQL 位于 `docs/ops/duckdb-local-lens-experiment.sql`。下面保留每个问题的核心 SQL 和去敏结果。

### Q1. 三个数据源各有多少逻辑行，中文内容是否正常读取？

```sql
SELECT
  (SELECT count(*) FROM trace_rows) AS trace_rows,
  (SELECT count(*) FROM motion_rows) AS motion_rows,
  (SELECT count(*) FROM catalog_rows) AS catalog_rows,
  (SELECT count(*) FROM trace_rows
   WHERE regexp_matches(coalesce("色散内容", ''), '[一-龥]')) AS trace_rows_with_chinese,
  (SELECT count(*) FROM motion_rows
   WHERE regexp_matches(coalesce("动作", ''), '[一-龥]')) AS motion_rows_with_chinese;
```

结果：色散 291 行、探筹 39 行、catalog 1 个顶层对象；分别有 73 行色散正文和 38 行动作名命中中文字符。CSV 物理行数会被带换行的 quoted field 扰动，DuckDB 的 291 是 CSV 逻辑记录数，因此比 `wc -l` 可靠。

### Q2. 哪些字段为空，是否存在重复记录或可能重复的键？

```sql
SELECT
  count(*) FILTER (WHERE coalesce(trim(cast("色散内容" AS VARCHAR)), '') = '')
    AS blank_content,
  count(*) FILTER (WHERE coalesce(trim(cast("关联项目" AS VARCHAR)), '') = '')
    AS blank_project
FROM trace_rows;

SELECT count(*) AS duplicate_groups, sum(copies - 1) AS duplicate_excess_rows
FROM trace_duplicate_counts
WHERE copies > 1;

SELECT "色散内容"
FROM trace_rows
WHERE coalesce(trim("色散内容"), '') <> ''
GROUP BY "色散内容"
HAVING count(*) > 1;
```

结果：色散空字段数依次为日期 0、色散内容 216、关联项目 278、判断线索 280、完成质量 284、情绪 287、执行情境强度 281、状态 56、色散类型 270、认知挑战 280。探筹没有任何含空字段的记录。

色散有 37 个完整重复组、160 条超额重复；37 个组全部是正文为空的记录。非空色散正文候选键重复数为 0，探筹动作候选键重复数为 0。这里揭示的是 canonical CSV 的真实质量问题，但本实验只报告，不修改文件。

首轮曾误用只选择 `count(*)` 的 `GROUP BY ALL`，得到错误的 1 组/290 条；DuckDB 返回结构使错误可立即与独立 Python 基线对照定位，改成显式十列分组后得到上述结果。这是唯一一轮 SQL 修复。

### Q3. 色散在时间上有什么峰值、空档和异常值？

```sql
SELECT
  strftime(try_cast("日期" AS DATE), '%Y-%m') AS month,
  count(*) AS records
FROM trace_rows
WHERE try_cast("日期" AS DATE) IS NOT NULL
GROUP BY month
ORDER BY month;
```

结果：264 行可转成日期，27 行为非日期值；范围为 2025-11-14 至 2026-05-14。月分布为 2025-11: 85、2025-12: 13、2026-01: 50、2026-03: 52、2026-04: 22、2026-05: 42。明显峰值是 2025-11，2026-02 是空档。DuckDB 正确把混合字段推断成 `VARCHAR`；日期语义必须显式 `try_cast`，没有静默丢行。

### Q4. 探筹枚举是否紧凑，是否有孤立值或拼写漂移？

```sql
SELECT "场景", count(*) AS records FROM motion_rows GROUP BY "场景";
SELECT "强度", count(*) AS records FROM motion_rows GROUP BY "强度";
SELECT "来源", count(*) AS records FROM motion_rows GROUP BY "来源";
SELECT "象限", count(*) AS records FROM motion_rows GROUP BY "象限";
```

结果：场景为地面 19、站立 18、坐姿 2；强度为轻 18、中 15、重 6；来源为克·首批 31、克·首批·地名系列 8；象限为流动 13、释放 11、唤醒 10、维持 5。`坐姿` 是低频真实值；未发现同义拼写漂移或单条孤立枚举。

### Q5. 两个 CSV 的字段和推断类型如何漂移，多文件联合查询是否稳定？

```sql
SELECT column_name, data_type
FROM duckdb_columns()
WHERE table_name IN ('trace_rows', 'motion_rows')
ORDER BY table_name, column_index;

SELECT source_kind, count(*)
FROM combined_csv_rows
GROUP BY source_kind;
```

结果：色散 10 列、探筹 8 列，共享列为 0；两边的 CSV 列均推断为 `VARCHAR`。`union_by_name` 一次扫描仍保留色散 291 行和探筹 39 行，并以 NULL 填补异构列。它适合临时观察，但也会掩盖“两个文件其实没有共同 schema”的事实，所以必须同时输出 schema drift，不能只看 union 成功。

### Q6. 真实 JSON 的嵌套数组、对象与类型能否稳定展开？

```sql
SELECT
  part.path,
  part.start,
  part."end",
  part.count
FROM read_json_auto(getvariable('catalog_path')),
UNNEST(parts) AS unpacked(part);
```

结果：`parts` 稳定推断为 `STRUCT(path VARCHAR, start BIGINT, end BIGINT, count BIGINT)[]`；展开 6 个对象，总数为 293；顶层必需字段缺失行 0、part 必需字段缺失行 0、路径重复 0、范围计数错误 0、序列空档 0。标量 `version/count/part_size` 推断为 `BIGINT`，字符串字段为 `VARCHAR`。不猜造不存在的缺失情况。

### Q7. NLM corpus catalog 是否均衡且内部完整？

```sql
SELECT
  c.count AS declared_total,
  p.computed_total,
  p.first_id = 1 AND p.last_id = c.count AS covers_declared_range
FROM catalog_rows AS c
CROSS JOIN (
  SELECT
    sum(item_count) AS computed_total,
    min(start_id) AS first_id,
    max(end_id) AS last_id
  FROM catalog_parts
) AS p;
```

结果：声明总数与分片求和均为 293；5 个满 50 条分片、1 个末尾 43 条分片、0 个超大分片；覆盖从 1 到 293，合同为真。这不是演示题：此前需要 Python 构建器或临时脚本才能审计 NLM corpus 分片健康；SQL 不打开任何卡片正文即可回答。

## 输出、耗时与错误行为

三次 JSON 与三次 CSV 独立运行结果逐字节一致。所有查询作为一个只读批次执行；为避免反复读取私人输入，没有为了微基准拆成 7 个进程。因而各问题的保守上界都是整批耗时：

| 输出 | 三次 wall time | 结论 |
|---|---|---|
| JSON | 0.14s / 0.08s / 0.18s | 7 条结果，`jq` 验证通过 |
| CSV | 0.15s / 0.11s / 0.08s | header + 7 条记录，Python `csv.DictReader` 稳定读取 |

核心验证：

```bash
DUCKDB_BIN=/private/tmp/duckdb-cli-v1.5.4/duckdb \
  docs/ops/run-duckdb-local-lens-experiment.sh \
  '成为/色散.csv' \
  '成为/探筹.csv' \
  'wo/nlm/corpus/becoming-ideas/catalog.json' \
  json | jq -e 'length == 7'
```

结果为 `true`。无匹配查询返回 `[]` 且退出码为 0，语义清楚。引用不存在列时退出码为 1，并给出 `Binder Error`、候选列、SQL 行列位置和插入符；足够让 Codex 自主修复查询。

## schema inference 与现有方式比较

推断观察：

- 两个 CSV 的所有列均为 `VARCHAR`；这对稀疏中文 canonical 文件是保守而安全的。
- 色散日期混入非日期状态值，因此没有推断成 `DATE`。时间问题使用 `try_cast`，27 行异常被显式计数。
- JSON 的整数、字符串、嵌套 struct 数组推断正确，本实验无需手工 schema。
- 若未来出现前导零 ID、大整数越界、日期格式混杂或跨批 schema 漂移，应显式指定列类型，不依赖采样推断。

实际收益：

- **明显更清楚 1：跨 CSV 漂移。** `read_csv_auto([..], union_by_name=true, filename=true)` 加 `duckdb_columns()` 同时完成解析、类型观察、字段集合和联合计数。Node 标准库没有 CSV parser；Python 需要自行解析、合并列集并另写类型推断。
- **明显更清楚 2：时间聚合。** `try_cast + strftime + GROUP BY` 同时保留异常数和月桶，查询意图直接可读，不需要 Python 的循环、异常捕获和 Counter 管线。
- **明显更清楚 3：嵌套 JSON 连续性。** `UNNEST(parts)` 配合 `lag(end_id)` 能直接表达分片空档；jq 可以处理这个单文件，但与 CSV 联合、schema 类型和多问题统一输出时明显变复杂。
- **没有优势的部分：** 对 39 行探筹做单个 Counter，或遍历所有 CSV 字段算空值，短 Python 可能更短。DuckDB 的价值不在替代每一个 `jq`/Python 小操作，而在同一进程里跨格式扫描、聚合、连接和生成稳定结构化输出。

这些文件规模很小，性能不是采用理由；0.08–0.18 秒只证明启动和查询成本可接受。采用理由是临时分析表达力与跨格式一致性。

## 限制与风险

1. Homebrew bottle 本轮受 TLS 传输故障阻塞；后续实现 PR 必须单独验证可重复安装与版本固定，不能把本次 `/private/tmp` 下载方式变成正式供应链。
2. `union_by_name` 会让完全异构的 CSV 看似“联合成功”；正式模板必须并列展示 schema drift。
3. CLI JSON 模式下若连续执行多个有结果的 `SELECT`，会输出多个相邻 JSON 文档而不是一个可供 `jq` 解析的文档。本实验用“临时 view + 单个最终 UNION 结果集”规避。
4. 色散包含个人叙事。正式只读 wrapper 仍应默认聚合输出，不能提供会意外回显正文的宽松模板。
5. 小数据上的 DuckDB 不比 Python 更快到有实际意义；不要因此迁移 canonical 文件或维护持久数据库。
6. schema inference 是探索辅助，不是数据合同。需要运算的日期/数字字段必须显式转换并统计转换失败。

## 只读与产物检查

- runner 明确使用 `:memory:`；没有 `.duckdb` 或 `.duckdb.wal`。
- SQL 禁用扩展自动安装和自动加载，且不含 `INSTALL`、`LOAD`、HTTP 或 S3。
- 不含 `INSERT`、`UPDATE`、`DELETE` 或 `COPY TO`。
- 三个输入与 `HEAD` 无 diff；没有向 `成为/`、`领域/`、`资源/` 写回。
- 临时结果只存在于 `mktemp` 目录，runner 退出时清理。
- 没有后台进程、CI、mise、Rust CLI、Omega 文件或 Omega 测试改动。

## 后续建议（不在本 PR 实现）

另开最小实现 PR，再评估 Issue #312 的 doctor、正式只读 wrapper、小型公开 smoke fixture、mise 固定版本和少量 SQL 模板。该 PR 必须继续保持 canonical 文件优先、默认内存数据库、无网络扩展、无自动写回，并单独处理 Homebrew 安装可靠性。
