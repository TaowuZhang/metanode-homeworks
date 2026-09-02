# 作业2：BeggingContract（讨饭合约）

允许任何人向合约捐赠 ETH，记录每个地址的累计金额，只有 owner 能提取全部资金。

## 必做功能

| 函数 | 说明 |
|---|---|
| `donate()` | `payable`，累计 `msg.value` |
| `withdraw()` | `onlyOwner`，用 `address.transfer` 把余额转给 owner |
| `getDonation(address)` | 查询某地址累计捐赠 |
| `receive()` | 普通转账 ETH 也会记入捐赠，方便 MetaMask 直接打币 |

## 额外挑战

1. `Donation` / `Withdrawal` 事件
2. `getTopDonors()`：捐赠金额最高的 3 个地址
3. `setDonationWindow(start, end, enabled)`：owner 可打开时间窗，窗外交捐会 revert

时间窗默认关闭，Remix / 本地测试不用先设时间。

## Remix 部署与测试（作业要求的方式）

1. 打开 https://remix.ethereum.org
2. 复制 `BeggingContract.sol`
3. 编译器 `0.8.28`，Compile
4. Environment 改成 `Injected Provider - MetaMask`，网络切到 **Sepolia**
5. Deploy。记下合约地址
6. 在 Remix 里用另一个账户调用 `donate`，value 填 `0.001 ether`
7. `getDonation(你的地址)` 应等于刚捐的数量
8. 切回 owner 账户调用 `withdraw`，合约余额应变 0
9. 到 Sepolia Etherscan 打开合约地址，截 `donate` / `withdraw` 交易

Goerli 已停用，用 Sepolia。

## 本地测试

```bash
cd solidity/homework-tests
npm install
npm test
```

## 测试网地址

部署成功后写在 [DEPLOYMENTS.md](./DEPLOYMENTS.md)。若本次环境没有 Sepolia 私钥，请用上面的 Remix + MetaMask 步骤自己部署，再把地址补进该文件。
