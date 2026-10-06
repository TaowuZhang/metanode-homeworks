# Uniswap V3：集中流动性、tick 与费率档

V3 相比 V2 最重要的变化不是换掉 AMM，而是把“所有 LP 都在 `(0, +∞)` 全价格范围提供同样形状的流动性”改成：**每个 LP 可以选择一个价格区间，只在该区间内提供流动性。**

---

## 1. 集中流动性的优势是什么？

V2 中，流动性沿整条 `x * y = k` 曲线分布。现实里很多交易对长期只在较窄价格范围活动，因此大量资本被放在几乎不会到达的价格区域。

V3 允许 LP 选择 `[P_lower, P_upper]`，把资本集中在更可能成交的价格区间。

主要优势：

1. **资本效率更高。** 同样资金可以在当前价格附近提供比 V2 更深的有效流动性，从而降低交易者的价格影响。
2. **LP 可以表达自己的做市策略。** 稳定币可以集中在 1:1 附近；波动资产可以选择更宽的区间。
3. **同一 token pair 可以有多个费率池。** LP 可以在低费高周转和高费高风险补偿之间选择。
4. **流动性分布更接近订单簿式的价格区间管理。** LP 可以用窄区间近似有限价格区间的挂单策略。

但集中流动性不是“免费提升收益”：

- 当前价格离开 LP 区间后，该仓位不再是 active liquidity，也不再赚该池的 swap fee；
- 仓位会变成单边资产；
- 区间越窄，需要越频繁管理，并可能承受更强的相对 HODL 损失；
- 主动管理会带来 gas、再平衡和择时风险。

---

## 2. 集中流动性在合约层面大致如何实现？

可以把 V3 Core 的实现拆成四层：**Pool 全局状态 → Tick 边界 → Position → Swap 跨 Tick。**

### 2.1 Pool 保存当前价格与当前有效流动性

`slot0()` 中最关键的两个值是：

```text
sqrtPriceX96   当前 sqrt(token1/token0) 的 Q64.96 定点数
tick           当前价格所在 tick
```

池还直接保存：

```text
liquidity      当前价格范围内真正 active 的流动性
```

注意：`liquidity()` **不是所有 LP 仓位流动性的总和**，而只是当前价格下仍在区间内的有效流动性。

### 2.2 Tick 保存“跨过这个边界时流动性如何变化”

`ticks(int24)` 为已初始化 tick 保存：

```text
liquidityGross
liquidityNet
feeGrowthOutside0X128
feeGrowthOutside1X128
...oracle 相关累计值
initialized
```

核心是 `liquidityNet`：

- 某仓位在 `tickLower` 开始生效；
- 在 `tickUpper` 停止生效；
- 当价格穿过边界时，Pool 根据该 tick 的 `liquidityNet` 增减当前 active liquidity。

因此 V3 不需要遍历所有 LP 仓位来决定当前可用流动性，只需要处理被价格跨过的已初始化 tick。

### 2.3 TickBitmap 高效寻找下一个已初始化 Tick

理论 tick 很多，不可能逐个遍历。V3 用 `tickBitmap` 把 256 个 tick 初始化状态压进一个 `uint256` 位图。

swap 时可以快速寻找当前方向上下一个已初始化 tick，然后只在真正有流动性边界的地方更新状态。

### 2.4 Position 由 owner + tickLower + tickUpper 标识

V3 Core 的 position key 本质上由：

```text
owner + tickLower + tickUpper
```

组合得到。Position 记录：

```text
liquidity
feeGrowthInside0LastX128
feeGrowthInside1LastX128
tokensOwed0
tokensOwed1
```

一个重要区别：**V3 Core 的 position 本身不是 ERC-721。**

用户常见的“Uniswap V3 LP NFT”来自 Periphery 的 `NonfungiblePositionManager`：它把 Core 中的仓位包装成便于用户持有、授权和转移的 NFT。

### 2.5 Swap 是一段一段跨 tick 执行

概念流程：

```text
当前 sqrtPrice / tick / liquidity
        ↓
寻找交易方向上下一个 initialized tick
        ↓
在当前流动性 L 下计算这一段能走到哪里
        ↓
若到达 tick 边界，则 cross tick
        ↓
按 liquidityNet 更新 active liquidity
        ↓
继续下一段，直到 amount 用完或触及 price limit
```

这就是集中流动性能在不遍历全部仓位的前提下参与 AMM 定价的关键。

---

## 3. Tick 和价格是什么关系？

V3 把价格离散成以 `1.0001` 为底的指数网格：

```text
P(i) = 1.0001 ^ i
```

其中 `i` 是 tick。

因此：

```text
tick = log(P) / log(1.0001)
```

一个 tick 代表原始价格比例大约变化 1 basis point（0.01%）。

合约内部为了提高计算效率，不直接长期用普通价格，而是使用平方根价格：

```text
sqrtPriceX96 = sqrt(P) * 2^96
```

`TickMath.getSqrtRatioAtTick()` 的定义正是：

```text
sqrt(1.0001 ^ tick) * 2^96
```

### decimals 必须考虑

Tick 对应的是 token 的**原始整数单位**比例：

```text
P_raw = raw token1 / raw token0 = 1.0001 ^ tick
```

若 token0 decimals 为 `d0`、token1 decimals 为 `d1`，人类单位的 token1/token0 价格为：

```text
P_human = 1.0001 ^ tick * 10^(d0 - d1)
```

所以不能只拿 tick 直接当“UI 上看到的价格”，否则 USDC(6) / ETH(18) 等交易对会差很多数量级。

### 一个边界细节

`slot0.tick` 表示最后执行过的 tick transition。若价格恰好停在 tick 边界，它可能与直接用 `TickMath.getTickAtSqrtRatio(sqrtPriceX96)` 算出的值存在边界差异。因此写合约时应理解具体接口语义，而不是认为二者在所有瞬间严格等价。

---

## 4. 为什么 V3 需要 Tick？

如果 LP 可以任意选连续实数价格作为区间边界，链上需要维护的边界状态会难以索引、聚合和高效遍历。

Tick 同时解决四件事：

1. **把连续价格空间离散化。** 所有仓位边界落在可枚举的价格网格上。
2. **成为流动性变化的事件点。** 跨 tick 时只需应用 `liquidityNet`。
3. **让位图索引成为可能。** `TickBitmap` 可以用 bit operation 快速找到下一个已初始化边界。
4. **让价格、流动性和预言机累计值使用统一离散坐标。** V3 oracle 会累计 tick 随时间的变化。

因此 tick 不只是“显示价格的编号”，而是 **V3 流动性状态机的价格坐标。**

---

## 5. 集中流动性的数学直觉

对区间 `[P_a, P_b]`、流动性 `L`，V3 可以看成对 V2 曲线做平移后的“虚拟储备”模型：

```text
(x + L / sqrt(P_b)) * (y + L * sqrt(P_a)) = L^2
```

当当前价格 `P` 位于区间内：

```text
amount0 = L * (sqrt(P_b) - sqrt(P))
          / (sqrt(P) * sqrt(P_b))

amount1 = L * (sqrt(P) - sqrt(P_a))
```

直觉：

- 价格靠近下界时，仓位更多由 token0 构成；
- 价格向上移动时，token0 被逐渐换成 token1；
- 价格达到上界后，仓位成为单边 token1；
- 反方向则相反。

这也是“价格离开区间后停止赚手续费”的数学基础：仓位此时已经不再提供当前价格处的 active liquidity。

---

## 6. Tick spacing 和手续费费率是什么关系？

### 6.1 它们是 Factory 中的一对治理参数，不是公式推导关系

`UniswapV3Factory` 保存：

```solidity
mapping(uint24 => int24) public feeAmountTickSpacing;
```

创建池时：

```text
fee -> 查询 tickSpacing -> 部署 Pool
```

原始 V3 Factory 构造函数默认启用：

| fee | 人类费率 | tick spacing |
|---:|---:|---:|
| `500` | 0.05% | 10 |
| `3000` | 0.30% | 60 |
| `10000` | 1.00% | 200 |

后来治理又为以太坊主网加入了 1 bp 档（具体部署还可以由治理启用更多费率与 spacing 组合）：

| fee | 人类费率 | tick spacing |
|---:|---:|---:|
| `100` | 0.01% | 1 |

`fee` 的单位是百万分之一，因此：

```text
500    = 0.05%
3000   = 0.30%
10000  = 1.00%
```

Factory owner 可以用：

```solidity
enableFeeAmount(uint24 fee, int24 tickSpacing)
```

增加新的组合；一个已经启用的 fee tier 不能随意改成另一套 spacing。

### 6.2 tick spacing 限制哪些 tick 能作为仓位边界

若：

```text
tickSpacing = 60
```

则 LP 只能在 spacing 对齐的 tick 上初始化边界，例如：

```text
..., -120, -60, 0, 60, 120, ...
```

spacing 越小：

- 区间边界越精细；
- LP 可以更精确地集中资本；
- 可初始化的 tick 更多；
- 状态更容易分散到更多 tick，大额 swap 可能跨越更多 initialized tick，增加 gas。

spacing 越大：

- 区间更粗；
- LP 精细调价能力下降；
- 可初始化 tick 更少；
- 状态和跨 tick 成本更容易受控。

---

## 7. 为什么不同费率池对应不同 tick spacing？

核心是 **市场波动特征、LP 风险补偿、价格精度与 gas 成本之间的工程折中。**

### 低费率池：适合低波动、高相关资产

如稳定币/稳定币：

- 价格通常在窄范围活动；
- LP 愿意用非常窄的区间提高资本效率；
- 交易者又非常敏感于手续费；
- 因此低 fee + 小 tick spacing 更合理。

1 bp fee tier 使用 spacing 1，就是为了让稳定币 LP 可以在大约 1 bp 粒度上设置边界。

### 高费率池：更适合高波动资产

波动更大的资产：

- LP 需要更高 fee 补偿库存/无常损失和主动管理风险；
- 通常不需要在每 1 bp 都建立极细价格边界；
- 更大的 spacing 限制 tick 状态密度，也降低极端情况下跨大量 tick 的 gas 压力。

### 关键结论

```text
fee 高  ≠ 数学上必然要求 spacing 大
fee 低  ≠ 数学上必然要求 spacing 小
```

二者是由协议治理选择并在 Factory 中绑定的**市场设计参数**。历史上的组合体现了经验性权衡，而不是 `tickSpacing = f(fee)` 这种公式。

---

## 8. V2 与 V3 的关键差异

| 维度 | V2 | V3 |
|---|---|---|
| 流动性范围 | `(0, +∞)` 全范围 | LP 自选 `[P_lower, P_upper]` |
| 资本效率 | 相对低 | 可显著提高 |
| LP 份额表示 | Pair ERC-20 LP Token | Core position；Periphery 常包装为 ERC-721 NFT |
| 费率 | 固定 0.30% swap fee | 多 fee tier |
| 价格坐标 | reserves 比例 | `sqrtPriceX96` + tick |
| 区间边界 | 无 | tick |
| 当前有效流动性 | 整池统一 | 随价格跨 tick 动态变化 |
| 管理复杂度 | 较低 | 较高，需要区间管理 |

---

## 9. 六个学习问题的压缩答案

### 1）集中流动性的优势是什么？

把资本放在更可能成交的价格范围，提高单位资本的交易深度和手续费效率；代价是离开区间后停止赚费，并增加管理风险。

### 2）合约层面大致如何实现？

Position 指定 `tickLower/tickUpper`；tick 保存 `liquidityNet`；当前价格只使用 in-range liquidity；swap 找到下一个 initialized tick，跨过时更新 active liquidity。

### 3）tick 和价格是什么关系？

```text
P_raw = 1.0001 ^ tick
sqrtPriceX96 = sqrt(P_raw) * 2^96
```

UI 价格还要校正 token decimals。

### 4）为什么需要 tick？

把连续价格离散成可索引、可位图搜索的边界，使链上能高效维护集中流动性和跨价格区间状态变化。

### 5）tick spacing 和手续费费率是什么关系？

Factory 为每个启用的 fee tier 固定一个 tick spacing；spacing 决定 LP 可以在哪些 tick 初始化区间边界。

### 6）为什么不同费率池对应不同 tick spacing？

不同市场的波动、合理 fee、价格精度需求和跨 tick gas 成本不同。低波动池通常适合低 fee + 小 spacing，高波动池通常容忍高 fee + 大 spacing；它是治理参数组合，不是数学必然关系。

---

## 10. 建议自己手工追的源码

按这个顺序最容易建立完整模型：

1. `TickMath.getSqrtRatioAtTick()`：tick → `sqrtPriceX96`；
2. `IUniswapV3PoolState.slot0()`：当前价格、tick；
3. `IUniswapV3PoolState.ticks()`：理解 `liquidityGross/liquidityNet`；
4. `TickBitmap`：理解如何找下一个 initialized tick；
5. `UniswapV3Pool._modifyPosition()` / `mint()`：理解上下界如何改变 tick 状态；
6. `UniswapV3Pool.swap()`：理解一步步推进 sqrt price、跨 tick 和更新 liquidity；
7. `UniswapV3Factory.enableFeeAmount()`：理解 fee tier 与 tick spacing 的绑定；
8. Periphery `NonfungiblePositionManager`：理解为什么用户最终看到的是 LP NFT。

---

## 参考

- Dapp-Learning V3 白皮书导读：https://github.com/Dapp-Learning-DAO/Dapp-Learning/blob/main/defi/Uniswap-V3/whitepaperGuide/understandV3Witepaper.md
- Uniswap V3 Core：https://github.com/Uniswap/v3-core
- `UniswapV3Factory.sol`：https://github.com/Uniswap/v3-core/blob/main/contracts/UniswapV3Factory.sol
- `TickMath.sol`：https://github.com/Uniswap/v3-core/blob/main/contracts/libraries/TickMath.sol
- `IUniswapV3PoolState.sol`：https://github.com/Uniswap/v3-core/blob/main/contracts/interfaces/pool/IUniswapV3PoolState.sol
- Uniswap Governance — Add 1 Basis Point Fee Tier：https://gov.uniswap.org/t/proposal-add-1-basis-point-fee-tier/14745
