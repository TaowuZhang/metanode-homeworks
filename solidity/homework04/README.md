# Solidity Homework 04 — Uniswap 与常用 ERC / EIP

本目录整理本阶段的三个学习任务：

1. [Uniswap V2：AMM、交易风险、LP 与核心接口](./uniswap-v2.md)
2. [Uniswap V3：集中流动性、tick 与费率档](./uniswap-v3.md)
3. [常用 ERC / EIP 标准速查](./erc-eip-standards.md)

## 学习目标

完成后应能独立解释：

- V2 常数乘积 AMM 的定价、手续费、价格影响与 `amountOut` 计算；
- 滑点与滑点容忍度的区别，以及为什么过宽的滑点会放大 MEV / 三明治风险；
- LP 份额、手续费收入与无常损失之间的关系；
- V2 Core 与 Periphery 的职责边界及主要接口；
- V3 为什么把流动性限制在价格区间，以及 tick、`sqrtPriceX96`、tick spacing、fee tier 如何组合；
- ERC-20、2612、721、1155、165、2981、1271、4626、4337 与 EIP-712、EIP-7702 的职责；
- ERC-8004、x402、ERC-8183 在 AI + Web3 中分别解决“身份/信任、支付、任务托管结算”的哪一层问题。

## 标准状态说明

标准状态按 **2026-10-06** 的 Ethereum EIPs 官方页面整理。特别注意：

- ERC-8004（Trustless Agents）目前仍为 **Draft**；
- ERC-8183（Agentic Commerce）目前仍为 **Draft**；
- x402 **不是 Ethereum ERC/EIP**，而是独立支付协议；HTTP `402 Payment Required` 是其核心入口，当前规范也覆盖 MCP / A2A 等传输。

因此实现时不能把 Draft 当作已经冻结、不会变化的生产标准。

## 主要参考资料

- Ethereum EIPs：https://eips.ethereum.org/all
- Ethereum ERCs：https://github.com/ethereum/ERCs
- Ethereum EIPs GitHub：https://github.com/ethereum/EIPs
- Uniswap V2 Core：https://github.com/Uniswap/v2-core
- Uniswap V2 Periphery：https://github.com/Uniswap/v2-periphery
- Uniswap V3 Core：https://github.com/Uniswap/v3-core
- Uniswap V3 Periphery：https://github.com/Uniswap/v3-periphery
- Dapp-Learning V3 白皮书导读：https://github.com/Dapp-Learning-DAO/Dapp-Learning/blob/main/defi/Uniswap-V3/whitepaperGuide/understandV3Witepaper.md
- x402：https://x402.org/
