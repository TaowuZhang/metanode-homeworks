# WQB Failure Taxonomy v1

状态：failure-mode → next-action 规则  
生成时间：2026-06-21T23:36:28+08:00

| failure | meaning | next_action |
|---|---|---|
| LOW_SHARPE | 信号弱或方向错 | invert_sign / change_operator_family / discard |
| LOW_FITNESS | 综合不够 | repair_turnover / repair_concentration / preserve low-corr parent |
| HIGH_TURNOVER | 换手过高 | increase_decay / lengthen_horizon / slower_operator |
| LOW_TURNOVER | 信号太慢 | decrease_decay / shorten_horizon / delta_probe |
| CONCENTRATED_WEIGHT | 持仓集中 | lower_truncation / group_rank / neutralization_probe |
| LOW_SUB_UNIVERSE_SHARPE | 子宇宙不稳 | neutralization_probe / universe_diagnostic / field_sibling |
| LOW_2Y_SHARPE | 近两年不稳 | regime_diagnostic / lengthen_horizon / reduce family budget |
| SELF_CORRELATION | 与自有 Alpha 太像 | change_field_family / change_operator_family / change_time_scale |
| PROD_CORRELATION | 与生产池拥挤 | change_mechanism / explore_new_dataset |
| DATA_DIVERSITY | 数据多样性不足 | change_dataset / combine distinct data family |
| REGULAR_SUBMISSION | 账户或提交节奏 | no expression repair; record platform state |
| POWER_POOL_CORRELATION | power pool 拥挤 | low-crowding dataset / new mechanism |
| EXPRESSION_TYPE_ERROR | 类型不匹配 | grammar repair; VECTOR first vec_avg/sum/count |
| AUTH | auth false | stop platform action; user restores auth |
| CAPTCHA | human verification | stop; do not bypass |
| 429 | rate limit | respect Retry-After; cooldown |
| CONCURRENT_LIMIT | 并发过高 | lower batch size; wait/poll |
