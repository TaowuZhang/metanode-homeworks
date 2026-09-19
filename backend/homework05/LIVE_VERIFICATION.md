# Backend homework05 — Sepolia live verification

本文件记录 `backend/homework05` 已经在 Ethereum Sepolia 上真实执行并可公开回读的结果。记录对应已验证代码提交：

- commit: `764e58f48e4fd84e1ae3f07acfa4af35c4fe2ee5`
- network: Ethereum Sepolia
- chain ID: `11155111`
- verification date: `2026-09-16`
- independent readback RPC: `https://ethereum-sepolia-rpc.publicnode.com`
- GitHub Actions run: `35077129268` (`success`)

> 这里不记录私钥、助记词、RPC 密钥或 `.env.local` 内容。测试私钥只用于本地 Sepolia 写操作。

## 1. 区块查询

通过 homework05 CLI 查询 Sepolia 最新区块，得到：

```text
number: 11717359
hash: 0x14f34e1d4c7c0748454c76ecb8440966b13fd6b15bd816afd0e602487ac731eb
parent: 0x11d9d4e5210344f9838311fa90343b8f9c3966b1b11ce393df904b55a81669e4
timestamp: 2026-09-16T14:35:12Z
transactions: 150
```

随后通过公开 Sepolia JSON-RPC 独立回读相同区块，区块号、hash、parent hash 和交易数量一致。

## 2. ETH 转账

真实发送账户：

```text
0x2dbd02230965Ba9303644Bfa8197036A500eff11
```

接收地址：

```text
0xE0585859d25e72301A9F59a23e47f5a6cCaF5cb8
```

转账金额：`1 wei`

交易哈希：

```text
0x435b075c1676b219144b0148c354db4778ea6c639e4d1f10085c48905757b69b
```

独立 receipt 回读结果：

```text
status: 0x1
block: 11717360
from: 0x2dbd02230965ba9303644bfa8197036a500eff11
to: 0xe0585859d25e72301a9f59a23e47f5a6ccaf5cb8
```

该结果证明 homework05 的 Go 代码已经完成真实 Sepolia 交易构造、私钥签名和广播，而不是仅生成本地交易对象。

## 3. Counter 部署

使用 `bindings.DeployCounter` 在 Sepolia 部署 `contracts/Counter.sol`。

合约地址：

```text
0xFDFf79ff9D95c27BFa75041fC7c7A0953a02a4d4
```

部署交易：

```text
0xac1fa61a8c580ae95c108c76a585c9bc26650bc452a4d4468b786bec1973ae51
```

独立 receipt 回读结果：

```text
status: 0x1
block: 11717377
from: 0x2dbd02230965ba9303644bfa8197036a500eff11
contractAddress: 0xfdff79ff9d95c27bfa75041fc7c7a0953a02a4d4
```

部署后独立读取链上 runtime bytecode，长度为 `382 bytes`，证明该地址已实际存在合约代码。

使用 generated Go binding 读取初始状态：

```text
number: 0
```

## 4. Counter increment

通过 generated Go binding 调用 `increment()`。

交易哈希：

```text
0x6115c308e08305dabf9d804a2a359d373c811939a690660470122e9e0be74f0b
```

独立 receipt 回读结果：

```text
status: 0x1
block: 11717386
from: 0x2dbd02230965ba9303644bfa8197036a500eff11
to: 0xfdff79ff9d95c27bfa75041fc7c7a0953a02a4d4
```

receipt 中 `Incremented(uint256)` 日志的数据值为 `1`。

随后通过 homework05 CLI 和独立 `eth_call` 再次读取 `number()`，两者都得到：

```text
number: 1
```

这证明 generated binding 已经真实连接 Sepolia 合约、发送状态修改交易并读取修改后的链上状态。

## 5. Infura Sepolia endpoint

官方环境要求「注册 Infura 账户，获取 Sepolia 测试网络的 API Key」。该项在 2026-09-19 由操作者本机闭合：使用现有 homework05 CLI，对 `sepolia.infura.io` 做只读 `block latest`。API Key 只留在操作者本机环境，未进入聊天、仓库或 CI 日志。

CLI 命令：

```bash
go run ./cmd/homework05 block latest
```

操作者本机确认：

```text
rpc_host: sepolia.infura.io
```

CLI 输出：

```text
number: 11739627
hash: 0x7a1d7b2bba8b061d3c894c1550dccbe55979500c3cac315addbfcac5a63246f5
parent: 0xfd1d325b5ae9f9063e9ddc5a3536695cf9d2e250eb28d41ff9754e59f18f359f
timestamp: 2026-09-19T19:32:00Z
transactions: 102
```

随后通过公开 Sepolia JSON-RPC `https://ethereum-sepolia.publicnode.com` 对同一 `block hash` 做 `eth_getBlockByHash` 独立回读，区块号、hash、parent hash、UTC 时间戳和交易数量一致。Sepolia Etherscan 区块页 `11739627` 也给出相同 hash / parent hash / timestamp / 102 transactions。

该结果证明 Infura Sepolia HTTPS endpoint 已经可以驱动现有 `ethclient` 查询路径。不需要为此重放转账、部署或 `increment()`。

## 6. 对官方 homework05 要求的可证明状态

| 官方要求 | 当前证据 |
| --- | --- |
| 使用 `ethclient` 连接 Sepolia | 已通过真实区块查询、转账、部署和调用验证 |
| 查询区块 hash / timestamp / 交易数量并输出 | 已验证，见本文件第 1 节 |
| 构造、签名并发送 ETH 转账 | 已验证，见第 2 节真实 tx + receipt |
| 输出转账交易哈希 | 已验证 |
| Solidity Counter 合约、ABI、bytecode | 已保存在 `contracts/` |
| 使用 `abigen` 生成 Go binding | `bindings/counter.go` 已生成并纳入仓库；再生成命令见 `README.md` |
| generated binding 连接 Sepolia 合约 | 已验证，见第 3、4 节 |
| 调用 `increment()` 并输出结果 | 已验证，链上最终 `number = 1` |
| 注册 Infura 并获取 Sepolia API Key | 已验证：操作者本机通过 `sepolia.infura.io` 运行现有 CLI；公开回读见第 5 节。仓库不保存 API Key |

## 7. 安全说明

- 不在仓库中保存测试账户私钥。
- `.env.local` 由 Git ignore 规则排除，不属于提交内容。
- 本文件只记录公开链上数据和可复现实验结果。
- Infura API Key 只存在于操作者本机环境，不写入源码、聊天记录或本文件。
