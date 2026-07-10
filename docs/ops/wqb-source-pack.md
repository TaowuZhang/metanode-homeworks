# WQB Source Pack v1

状态：社区经验转规则，不沉公开 Alpha 原式  
生成时间：2026-06-21T23:36:28+08:00

## 吸收原则

外部资料只进入三类资产：source summary、candidate grammar、failure taxonomy。禁止复制公开 alpha examples 的完整表达式；只抽象机制、operator family 和研究纪律。

| 来源 | 可吸收 | 不可照搬 | 转成规则 |
|---|---|---|---|
| WQB seminar notes | 平台 gates、字段/设置习惯、operator 思路 | 具体表达式 | rulebook + grammar templates |
| Learn2Quant | 因子研究流程、风险解释 | 脱离 WQB gate 的指标崇拜 | seed → prototype → gate_check |
| Finding Alphas | 经济假设、简单表达、数据源优先 | 书中公式 | hypothesis pattern |
| 101 Formulaic Alphas | rank / delta / ts_rank / correlation 结构 | 公开原式 | operator_family only |
| WQB automation repos | queue、async polling、ledger、dedupe | autonomous bulk mining / auto submit | 逐批授权队列 |
| jglazar notes | 实操经验、失败修复线索 | 平台规避捷径 | failure-mode repair table |
| alpha examples repos | 结构启发 | 可识别复制 | negative rule：不得复制 |
| Bailey / López de Prado | 多重测试与过拟合警告 | 学术检验机械套用 | family budget + stop rules |
| Adaptive Markets | alpha decay / crowding | 把历史强信号当永久真理 | recent stability + cooldown |

## 当前 prior

- 简单结构优先于复杂表达式。
- 字段家族比 operator 花活更决定命运。
- operator_count=2 的 rank + ts_zscore / ts_delta 曾高频存活。
- USA / TOP3000 / D1 是当前安全港。
- fnd65 是强信号但稳定带毒，主毒为 LOW_2Y_SHARPE + CONCENTRATED_WEIGHT。
