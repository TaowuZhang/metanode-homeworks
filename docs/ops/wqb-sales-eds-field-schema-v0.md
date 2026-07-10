# Sales EDS 字段清单 v0

来源：Notion《Sales EDS 字段清单 v0》。沉降时间：2026-06-05。

## r / N / D

- r：字段清单来自 read-only `/datafields` 通路，已经从现场发现变成可复现数据族地形。
- N：字段簇、第一层判断与 Prototype 接法可以沉 Git；simulation 现场留 Notion / WQB。
- D：本文件保存 Sales EDS 的稳定字段地形，不保存账号凭证或提交行为。

## 本次确认

- Dataset：`analyst4`
- Dataset name：Analyst Estimate Data for Equity
- Region：USA
- Delay：1
- Universe：TOP3000
- Theme：USA/D1 Fast Datasets Power Pool June`26
- Query：`search=Sales`
- Count：1324
- Worker version：`528d0ccd-0a57-41c8-9d01-d6b20f770c87`

## Sales EDS 主线字段簇

### Center｜中心预期

- `sales_estimate_average`｜Sales - mean of estimations with a delay of 1 quarter｜MATRIX｜coverage 1｜userCount 90｜alphaCount 139
- `sales_estimate_average_quarterly`｜Sales - mean of estimations｜MATRIX｜coverage 1｜userCount 68｜alphaCount 112
- `sales_estimate_average_annual`｜Sales - mean of estimations｜MATRIX｜coverage 1｜userCount 69｜alphaCount 96
- `sales_estimate_median_value`｜Sales - Median value among forecasts｜MATRIX｜coverage 1｜userCount 59｜alphaCount 89
- `median_sales_estimate`｜Sales - median of estimations｜MATRIX｜coverage 1｜userCount 63｜alphaCount 81
- `sales_estimate_median_quarterly`｜Sales - median of estimations｜MATRIX｜coverage 1｜userCount 70｜alphaCount 114

### High / Low｜分布两端

- `sales_estimate_maximum`｜Sales - The highest estimation｜MATRIX｜coverage 1｜userCount 102｜alphaCount 134
- `sales_estimate_minimum`｜Sales - The lowest estimation｜MATRIX｜coverage 1｜userCount 61｜alphaCount 70
- `sales_estimate_maximum_quarterly`｜Sales - The highest estimation｜MATRIX｜coverage 1｜userCount 49｜alphaCount 78
- `sales_estimate_minimum_quarterly`｜Sales - The lowest estimation｜MATRIX｜coverage 1｜userCount 58｜alphaCount 91
- `highest_sales_estimate`｜Sales - The highest estimation for the annual period｜MATRIX｜coverage 1｜userCount 54｜alphaCount 66
- `lowest_sales_estimate`｜Sales - The lowest estimation for the annual period｜MATRIX｜coverage 1｜userCount 48｜alphaCount 65

### Dispersion｜分歧 / 宽度

- `sales_estimate_standard_deviation`｜Sales - standard deviation of estimations｜MATRIX｜coverage 0.7259｜userCount 373｜alphaCount 574
- `sales_estimate_stddev_quarterly`｜Standard deviation of Sales estimations｜MATRIX｜coverage 0.7259｜userCount 390｜alphaCount 679
- `sales_estimate_dispersion`｜Standard deviation of Sales estimations for the annual period｜MATRIX｜coverage 0.6792｜userCount 49｜alphaCount 74
- `anl4_fs_detail_estimates_basic_qf_v4_nd_sales_std`｜Standard deviation of Sales estimations｜MATRIX｜coverage 0.7069｜userCount 19｜alphaCount 26
- `anl4_fs_detail_estimates_basic_qf_delay1_v4_nd_sales_std`｜Standard deviation of Sales estimates｜MATRIX｜coverage 0.7069｜userCount 17｜alphaCount 22

### Breadth｜覆盖 / 广度

- `sales_estimate_count`｜Sales - number of estimations｜MATRIX｜coverage 1｜userCount 1529｜alphaCount 2007
- `sales_estimate_count_quarterly`｜Sales - number of estimations｜MATRIX｜coverage 1｜userCount 1433｜alphaCount 1858
- `sales_estimate_count_2`｜Number of Sales estimates｜MATRIX｜coverage 1｜userCount 106｜alphaCount 133
- `anl4_fs_detail_estimates_basic_qf_delay1_v4_nd_sales_number`｜Number of Sales estimates｜MATRIX｜coverage 0.9852｜userCount 313｜alphaCount 389
- `anl4_fs_detail_estimates_basic_qf_v4_nd_sales_number`｜Sales - number of estimations｜MATRIX｜coverage 0.9852｜userCount 268｜alphaCount 320

### Current / Previous｜当前估值 / 前值

- `sales_estimate_value`｜Sales - Estimated value｜VECTOR｜coverage 0.9931｜userCount 18｜alphaCount 42
- `sales_previous_estimate_value`｜The previous estimation of Sales｜VECTOR｜coverage 0.9867｜userCount 15｜alphaCount 29
- `anl4_ads1detailafv110_estvalue`｜Estimation value｜VECTOR｜coverage 0.6878｜userCount 86｜alphaCount 160
- `anl4_ads1detailafv110_prevval`｜The Previous Estimation of Financial Item｜VECTOR｜coverage 0.6856｜userCount 69｜alphaCount 126
- `anl4_ads1detailqfv110_estvalue`｜Estimation value｜VECTOR｜coverage 0.6292｜userCount 63｜alphaCount 132
- `anl4_ads1detailqfv110_prevval`｜The previous estimation of financial item｜VECTOR｜coverage 0.6268｜userCount 161｜alphaCount 361

### Guidance｜公司指引

- `sales_guidance_value`｜Sales - Guidance value for the annual period｜MATRIX｜coverage 0.4262｜userCount 244｜alphaCount 278
- `sales_guidance_value_quarterly`｜Sales - guidance value｜MATRIX｜coverage 0.3262｜userCount 60｜alphaCount 69
- `sales_max_guidance_value`｜Maximum guidance value for annual sales｜MATRIX｜coverage 1｜userCount 53｜alphaCount 58
- `sales_min_guidance_value`｜Minimum sales guidance for the annual period.｜MATRIX｜coverage 1｜userCount 27｜alphaCount 39
- `sales_max_guidance_quarterly`｜The maximum guidance value for sales.｜MATRIX｜coverage 1｜userCount 239｜alphaCount 280
- `sales_min_guidance_quarterly`｜Minimum guidance value for Sales｜MATRIX｜coverage 1｜userCount 226｜alphaCount 314
- `anl4_fs_guidances_basic_qf_nd_sales_maxguidance`｜Upper bound of the company’s Sales guidance range｜MATRIX｜coverage 1｜userCount 8｜alphaCount 16
- `anl4_fs_guidances_basic_qf_nd_sales_minguidance`｜Lower bound of the company’s Sales guidance range｜MATRIX｜coverage 1｜userCount 3｜alphaCount 11

### Actual｜事后真实值

- `actual_sales_value_annual`｜Sales - Actual Value｜MATRIX｜coverage 0.9943｜userCount 240｜alphaCount 316
- `actual_sales_value_quarterly`｜Sales - Value in financial services income statement (in millions)｜MATRIX｜coverage 1｜userCount 276｜alphaCount 421
- `anl4_fs_actuals_basic_af_nd_sales_value`｜Sales/Revenue actual value for the annual period｜MATRIX｜coverage 0.9683｜userCount 0｜alphaCount 0

## 第一层量化判断

- 不先追最热字段：`sales_estimate_count`、`sales_estimate_count_quarterly`、`sales_estimate_standard_deviation`、`sales_estimate_stddev_quarterly` 热度高，更适合 filter / regime / normalization。
- 首选组合：中心变化 + 分歧状态 + 覆盖约束。
- 低拥挤候选：`sales_estimate_value` / `sales_previous_estimate_value`，适合做 revision magnitude 原子字段；若 VECTOR 处理困难，退回 MATRIX consensus 字段。

## 接到 Prototype 亭

| Prototype | 核心问题 | 候选字段 |
|---|---|---|
| P1 中心迁移 | Sales consensus 是否持续上修 | `sales_estimate_average`, `sales_estimate_median_value` |
| P2 修正幅度 | 当前估值相对 previous 是否有方向性变化 | `sales_estimate_value`, `sales_previous_estimate_value` |
| P3 分歧压缩 / 扩张 | 预期分布宽度扩大还是收敛 | `sales_estimate_standard_deviation`, `sales_estimate_dispersion` |
| P4 广度确认 | 是否有足够 analyst count 支撑 | `sales_estimate_count`, `sales_estimate_count_quarterly` |
| P5 公司指引偏移 | company guidance 是否领先或背离 analyst consensus | `sales_guidance_value`, `sales_max_guidance_quarterly`, `sales_min_guidance_quarterly` |

## Ξ 声明 / 折叠审计

- 已沉 Git：字段簇、第一层判断、Prototype 接法。
- 留 Notion：后续 simulation、结果回收、活动现场判断。
- 不外运：账号凭证、平台会话、自动 submit 或未授权批量行为。
