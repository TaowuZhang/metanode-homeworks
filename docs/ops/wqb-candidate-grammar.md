# WQB Candidate Grammar v1

状态：候选表达式抽象语法，不作为公开 Alpha 原式库  
生成时间：2026-06-21T23:36:28+08:00

## 1. Cell key

region × universe × delay × dataset × field_family × field × field_type × operator_family × horizon × decay × neutralization × truncation × sign

## 2. Field type adapter

| field_type | first adapter | allowed next layer |
|---|---|---|
| MATRIX | direct | rank / ts_zscore / ts_delta / ts_rank |
| VECTOR | vec_avg / vec_sum / vec_count | 再接 MATRIX time-series operator |
| EVENT-like | event count / recency / vec_count | 先转可比较截面 |
| GROUP / categorical | group transform | 避免直接 ts operator |

## 3. Operator families

- rank_ts_zscore：rank(ts_zscore(field, z))
- rank_ts_delta_ts_zscore：rank(ts_delta(ts_zscore(field, z), d))
- rank_ts_rank：rank(ts_rank(field, h))
- inverse_rank_ts_zscore：-rank(ts_zscore(field, z))
- vector_avg_rank_ts_zscore：rank(ts_zscore(vec_avg(field), z))

表达式全文不进入长期 Git ledger；batch candidate 只保存 expression_ref、operator_family 与参数。真实表达式只在受控平台执行上下文或本地私有短期上下文中出现。

## 4. Horizon buckets

- short：3–10
- medium：15–30
- long：45–120

## 5. Mutation slots

标准 family batch 6 个候选：baseline / horizon / decay / truncation / neutralization / sign-or-operator-sibling。
