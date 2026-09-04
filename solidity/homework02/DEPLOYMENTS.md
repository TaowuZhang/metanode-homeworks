# BeggingContract 部署记录

## Sepolia

Status: **已部署，donate / getDonation / withdraw 链上测试通过（2026-09-04）**，`chainId = 11155111`。

使用仓库原始 `BeggingContract.sol`，通过 **Remix + Chrome MetaMask** 部署和测试。编译器为 `0.8.28+commit.7893614a`。Remix 2.5.6 对应环境名称为 `Browser Extension → Sepolia Testnet - MetaMask`（注入钱包环境）。未修改合约，`withdraw` 保持 `payable(owner).transfer(amount)`；未导出私钥或助记词。

### 地址与交易

| 项目 | 地址 / 交易哈希 |
| --- | --- |
| 网络 | Sepolia (`11155111`) |
| 合约 | `BeggingContract` |
| 合约地址 | `0x25c98d2940eDc3f8E34e19e49753bF101C68E762` |
| owner / 部署 / 捐赠账户 | `0x5fCE1A4989A8bE17CcECfFB08Ed01614769F91eE` |
| 部署交易 | [`0x0c8f15371849324282fad5767c22f01fc5d07ba7c6e1c5ee02d1ec58c3a776c8`](https://sepolia.etherscan.io/tx/0x0c8f15371849324282fad5767c22f01fc5d07ba7c6e1c5ee02d1ec58c3a776c8) |
| donate 测试交易 | [`0xd9384b55c427d47bb1304ae246178e910c4d9e5fd5cb7a790e5efcd648ad802f`](https://sepolia.etherscan.io/tx/0xd9384b55c427d47bb1304ae246178e910c4d9e5fd5cb7a790e5efcd648ad802f) |
| withdraw 测试交易 | [`0xdbd87c65e07f2ee9adda50dbce547762c2b00d967cb45da13e7832cdeef0b152`](https://sepolia.etherscan.io/tx/0xdbd87c65e07f2ee9adda50dbce547762c2b00d967cb45da13e7832cdeef0b152) |
| Etherscan 合约 | [查看合约及余额](https://sepolia.etherscan.io/address/0x25c98d2940eDc3f8E34e19e49753bF101C68E762) |

### 链上回读与功能测试

通过 Remix 回读，并使用公共 Sepolia RPC 独立复核，原始请求和回执保存在 [evidence/sepolia-verification.json](./evidence/sepolia-verification.json)。

| 检查 | 结果 |
| --- | --- |
| 部署回执 | `status = 0x1`，区块 `11630262` |
| 部署后的代码 | `eth_getCode` 非空，运行时代码 `4266 bytes` |
| `owner()` | `0x5fCE1A4989A8bE17CcECfFB08Ed01614769F91eE` |
| `donate()` | 同一测试账户捐赠 `0.001 ETH = 1000000000000000 wei`；`status = 0x1`，区块 `11630276` |
| Donation 事件 | 本合约发出；donor 为上述账户，amount 和 totalFromDonor 均为 `1000000000000000` |
| 捐赠后余额 | 区块 `11630276` 回读为 `1000000000000000 wei` |
| `getDonation(捐赠地址)` | `1000000000000000 wei`；提取后累计捐赠记录仍保留 |
| `withdraw()` | owner 调用，附带 value 为 `0`；`status = 0x1`，区块 `11630294` |
| Withdrawal 事件 | 本合约发出；owner 为上述账户，amount 为 `1000000000000000` |
| 提取后余额 | 区块 `11630294` 和复核时 `latest` 均为 `0 wei`；Etherscan 显示 `0 ETH` |
| `getTopDonors()` | 第一名为上述捐赠账户，金额 `1000000000000000 wei`；其余两名为零地址、金额为 `0` |
| `timeRestricted()` | `false` |

`getDonation`、`owner`、`getTopDonors` 为只读调用，不产生交易哈希。捐赠发生于 `2026-09-04 01:23:48 UTC`，提取发生于 `2026-09-04 01:27:24 UTC`。

MetaMask 的捐赠和提取使用包装调用，外层交易 `to` 为 `0xdb9B1e94B5b69Df7e401DDbedE43491141047dB3`；Etherscan 的 donate 页面同时显示 EIP-7702 委托。实际收款与事件发出地址均为本 BeggingContract，上述事件和区块余额已独立验证。查看这些测试应使用表中的完整交易链接及 Internal Transactions / Logs，不能仅按合约页的普通 Transactions 列表判断。

### 截图

原始 Chrome 截图位于 [screenshots/](./screenshots/)：

1. [01-remix-deployment-success.png](./screenshots/01-remix-deployment-success.png)：部署成功、合约地址、交易哈希、Sepolia 网络。
2. [02-remix-donate-success.png](./screenshots/02-remix-donate-success.png)：donate 金额、成功回执、合约余额 `0.001 ETH`。
3. [03-remix-getDonation-result.png](./screenshots/03-remix-getDonation-result.png)：`getDonation` 返回 `1000000000000000`。
4. [04-remix-withdraw-success-zero-balance.png](./screenshots/04-remix-withdraw-success-zero-balance.png)：withdraw 成功回执与余额 `0 ETH`。
5. [05-etherscan-contract-zero-balance.png](./screenshots/05-etherscan-contract-zero-balance.png)：Etherscan 合约地址及余额 `0 ETH`。
6. [06-etherscan-donate-success.png](./screenshots/06-etherscan-donate-success.png)：Etherscan donate 交易 Success，转入合约 `0.001 ETH`。
7. [07-etherscan-withdraw-success.png](./screenshots/07-etherscan-withdraw-success.png)：Etherscan withdraw 交易 Success。

### 本地验证

2026-09-04 复跑：

- 作业1/2：`cd solidity/homework-tests && npm test` → **11 passing**（Solidity `0.8.28`）。
- 作业3：`cd solidity/homework03 && npm test` → **18 passing**。作业3与 main 的版本化文件无差异。

本次仅补充作业2部署记录、链上证据、截图与导师提交稿。原作业1、作业2、作业3合约均未修改。
