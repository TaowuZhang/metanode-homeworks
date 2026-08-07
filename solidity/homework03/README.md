# MetaNode Homework03 — Upgradeable NFT Auction Market

MetaNode Academy Solidity 基础二「任务3」作业：使用 **Hardhat 3** 实现一个可升级 NFT 拍卖市场，支持 ETH / ERC20 跨币种竞价、Chainlink Data Feed 美元计价、UUPS 升级、完整测试与 Sepolia 部署脚本。

## 作业要求对应关系

| 作业要求 | 实现 |
|---|---|
| ERC721 NFT | `MetaNodeNFT.sol`，支持 owner mint 与标准 ERC721 转移 |
| 创建拍卖 | 卖家将 NFT 托管进 `NFTAuction`，设置 USD 起拍价与持续时间 |
| ETH 出价 | `bidWithEth` |
| ERC20 出价 | `bidWithToken` |
| 跨币种比较 | 所有出价通过 Chainlink `AggregatorV3Interface` 转成 18 位 USD 值后比较 |
| 结束拍卖 | NFT 给最高出价者，资金给卖家；无出价则 NFT 退回卖家 |
| 被超价退款 | pull-payment 模式，ETH/ERC20 均通过 `withdrawRefund` 提取 |
| 合约升级 | OpenZeppelin UUPS；`NFTAuctionV2` 增加最小加价比例 |
| 单元/集成测试 | 18 个 Hardhat + Mocha 测试 |
| 覆盖率 | 96.00% line / 96.30% statement |
| 测试网部署 | `scripts/deploy-sepolia.ts` |
| 升级脚本 | `scripts/upgrade-sepolia.ts` |

## 项目结构

```text
contracts/
  MetaNodeNFT.sol
  NFTAuction.sol
  NFTAuctionV2.sol
  mocks/
    MockERC20.sol
    MockV3Aggregator.sol
scripts/
  deploy-sepolia.ts
  upgrade-sepolia.ts
test/
  NFTAuction.ts
TEST_REPORT.md
hardhat.config.ts
```

## 关键设计

### 1. 同一拍卖支持 ETH 与 ERC20 交叉竞价

拍卖不固定支付币种。每次出价都携带自己的支付资产：

- ETH 使用 `address(0)`；
- ERC20 使用其 token 地址；
- 合约读取对应 Chainlink Data Feed；
- 金额统一换算为 18 位精度 USD；
- 最高 USD 值决定领先者。

这样预言机直接参与核心竞价逻辑，而不是一个独立的演示函数。

### 2. Chainlink 价格归一化

`quoteUsd(token, amount)` 同时处理：

- token decimals；
- feed decimals；
- `latestRoundData()` 的 answer / round / updatedAt 有效性；
- 18 位 USD 统一精度。

Sepolia 默认配置：

- ETH/USD feed: `0x694AA1769357215DE4FAC081bf1f309aDC325306`
- LINK token: `0x779877A7B0D9E8603169DdbD7836e478b4624789`
- LINK/USD feed: `0xc59E3633BAAC79493d908e63626716e204A45EdF`

部署脚本允许通过环境变量覆盖这些地址。

### 3. Pull-payment 退款

被超价后不在新出价交易中主动向旧 bidder 转账，而是记入：

```solidity
pendingReturns[bidder][paymentToken]
```

旧 bidder 再调用 `withdrawRefund` 提取。这样避免某个拒收 ETH 的 bidder 阻塞后续正常出价。

### 4. UUPS 升级

V1 使用：

- `Initializable`
- `OwnableUpgradeable`
- `UUPSUpgradeable`
- `_authorizeUpgrade` + `onlyOwner`
- implementation constructor 中 `_disableInitializers()`

V2 只在 V1 存储之后追加：

```solidity
uint256 public minimumBidIncrementBps;
```

因此保持原有拍卖、feed、退款等状态布局。测试会先建立 V1 拍卖状态，再升级到 V2，并验证原状态仍存在。

## 安装

要求 Node.js 22+。

```bash
npm install
```

## 编译

```bash
npm run compile
```

## 测试

```bash
npm test
```

当前实测：

```text
18 passing / 0 failing
```

详细结果见 [TEST_REPORT.md](./TEST_REPORT.md)。

## Coverage

```bash
npm run coverage
```

当前实测：

```text
Total line coverage:      96.00%
Total statement coverage: 96.30%
```

HTML 报告生成到 `coverage/html/`，该目录属于可再生成产物，不提交 Git。

## Sepolia 部署

先准备测试网专用 RPC 与测试网账户。不要把私钥提交到仓库。

```bash
export SEPOLIA_RPC_URL='https://...'
export SEPOLIA_PRIVATE_KEY='0x...'
npm run deploy:sepolia
```

脚本会：

1. 部署 `MetaNodeNFT`；
2. 以 UUPS proxy 部署 `NFTAuction`；
3. 配置 Sepolia ETH/USD feed；
4. 配置 Sepolia LINK/USD feed；
5. 输出 NFT、proxy、implementation 与 feed 地址。

已于 **2026-08-07** 完成 Sepolia 实际部署、Chainlink feed 配置与 UUPS V1 → V2 升级。实际地址、交易哈希和升级后链上回读结果记录在 `DEPLOYMENTS.md`。

## 升级到 V2

```bash
export SEPOLIA_RPC_URL='https://...'
export SEPOLIA_PRIVATE_KEY='0x...'
export AUCTION_PROXY_ADDRESS='0x...'
export MIN_BID_INCREMENT_BPS='500'
npm run upgrade:sepolia
```

`500` 表示最低加价 5%。升级由 proxy owner 授权。

## 安全与边界

- 所有资金路径使用 `nonReentrant`；当前 Sepolia / Hardhat Cancun 环境采用 OpenZeppelin `ReentrancyGuardTransient`。
- NFT 创建拍卖时进入合约托管。
- 结算前先更新 `settled` 状态，再执行外部转移。
- ERC20 使用 `SafeERC20`。
- 被超价资金使用 pull-payment，不让退款接收方控制新出价是否能成功。
- Chainlink feed 必须存在、answer 为正、round 有效。
- 只有 owner 可配置 feed 和授权 UUPS 升级。
- 私钥、RPC token 等秘密由环境变量提供，`.env` 被 `.gitignore` 排除。

## 参考

- MetaNode homework03: <https://github.com/MetaNodeAcademy/LearningRoadmap/blob/main/contract/homework03.md>
- MetaNode reference project: <https://github.com/lc3091/hardhatV3Nft>
- OpenZeppelin Contracts: <https://docs.openzeppelin.com/contracts/5.x/>
- OpenZeppelin Hardhat Upgrades: <https://docs.openzeppelin.com/upgrades-plugins/hardhat-upgrades>
- Chainlink Data Feeds: <https://docs.chain.link/data-feeds>
- Hardhat 3: <https://hardhat.org/>
