老师您好，Solidity 作业1、2、3已整理完成，作业2本次已补齐 Sepolia 部署、链上功能测试和截图，请您查收。

仓库：https://github.com/TaowuZhang/metanode-homeworks
本次完整提交位于分支 `cursor/solidity-homework-01-02-a18d`：
https://github.com/TaowuZhang/metanode-homeworks/tree/cursor/solidity-homework-01-02-a18d

1. 作业1：`solidity/homework01/`。
2. 作业2：`solidity/homework02/`。使用原始 BeggingContract.sol，通过 Remix + MetaMask、Solidity 0.8.28 部署到 Sepolia（11155111）。
   - 合约地址：`0x25c98d2940eDc3f8E34e19e49753bF101C68E762`
   - Etherscan：https://sepolia.etherscan.io/address/0x25c98d2940eDc3f8E34e19e49753bF101C68E762
   - 已完成 donate 0.001 ETH，getDonation 返回 1000000000000000 wei，owner 执行 withdraw 后合约余额为 0。
   - 部署交易：https://sepolia.etherscan.io/tx/0x0c8f15371849324282fad5767c22f01fc5d07ba7c6e1c5ee02d1ec58c3a776c8
   - donate 交易：https://sepolia.etherscan.io/tx/0xd9384b55c427d47bb1304ae246178e910c4d9e5fd5cb7a790e5efcd648ad802f
   - withdraw 交易：https://sepolia.etherscan.io/tx/0xdbd87c65e07f2ee9adda50dbce547762c2b00d967cb45da13e7832cdeef0b152
   - 地址、交易哈希、链上回读见 `solidity/homework02/DEPLOYMENTS.md`；7张截图位于 `solidity/homework02/screenshots/`，独立 RPC 证据位于 `solidity/homework02/evidence/sepolia-verification.json`。
3. 作业3：`solidity/homework03/`，已完成 Sepolia 部署和升级，既有地址如下：
   - NFT：`0x9407360B3D01a82E8d70BBd584Ce2563F16447d4`
   - 拍卖代理：`0xE1e2165313Ab655cfB3223083c7d706123fa2eDB`
   - 升级交易：`0x98714c0ba10afd077390a5aeb12c7ce34550254f6a92dbeb2d4d433e7e100d1e`
   - 升级交易链接：https://sepolia.etherscan.io/tx/0x98714c0ba10afd077390a5aeb12c7ce34550254f6a92dbeb2d4d433e7e100d1e
   - 详细记录见 `solidity/homework03/DEPLOYMENTS.md`。

本地验证命令与本次复跑结果：
- 作业1/2：`cd solidity/homework-tests && npm test`，11 passing。
- 作业3：`cd solidity/homework03 && npm test`，18 passing。

本次仅补齐作业2交件证据，作业1和作业3合约未改动，BeggingContract 的 withdraw 保持 address.transfer。
