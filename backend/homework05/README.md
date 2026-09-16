# Backend homework05 — Go Ethereum / Sepolia

## 作业内容对应

### 1. Ethereum RPC

`ethrpc/client.go` 使用 `ethclient` 实现：

- 按区块号查询区块；不传区块号时查询最新区块；
- 按交易哈希查询交易，输出 sender / receiver / value / gas / pending 状态；
- 使用 EIP-1559 `DynamicFeeTx`（`maxPriorityFeePerGas` + `maxFeePerGas`）签名并广播 Sepolia ETH 转账；

CLI：

```bash
go run ./cmd/homework05 block latest
go run ./cmd/homework05 block 1234567
go run ./cmd/homework05 tx 0x<transaction-hash>
go run ./cmd/homework05 transfer 0x<receiver> 0.001
```

### 2. Solidity Counter + abigen

`contracts/Counter.sol` 提供：

- `uint256 public number`；
- `increment()`；
- `Incremented(uint256)` 事件。

本目录保留 Solidity `0.8.30` 编译产出的 ABI / bytecode，并检入与其一致的 Go binding。binding 使用 go-ethereum v1 `bind` API；可用下方固定命令通过 `abigen` 重新生成并复核。

CLI：

```bash
go run ./cmd/homework05 counter-deploy
go run ./cmd/homework05 counter-get 0x<counter-address>
go run ./cmd/homework05 counter-inc 0x<counter-address>
```

写操作会等待交易被打包，并检查 receipt 状态；部署成功会输出合约地址、交易哈希和区块号。



## 环境变量

`.env.example` 仅作为变量模板；程序直接读取 shell 环境变量，并不会自动加载 `.env`。**不要把真实私钥提交到 Git。**

```bash
export SEPOLIA_RPC_URL='https://...'
export SEPOLIA_PRIVATE_KEY='0x...'
```


## 安装与验证

本作业使用独立 Go module，避免 `go-ethereum` 的依赖影响仓库里此前的 Go 基础作业。

```bash
cd backend/homework05
go mod tidy
gofmt -w ./bindings/*.go ./counter/*.go ./ethrpc/*.go ./cmd/homework05/*.go
go vet ./...
go test ./...
go build -o /tmp/metanode-homework05 ./cmd/homework05
```


## 重新生成合约 binding

若修改 `Counter.sol`，可按下面方式重新生成。生成后再次执行 `gofmt`、`go vet` 和 `go test`。

```bash
solc --abi --bin contracts/Counter.sol -o contracts --overwrite
abigen \
  --abi contracts/Counter.abi \
  --bin contracts/Counter.bin \
  --pkg bindings \
  --out bindings/counter.go
```

`bindings/counter.go` 只承载合约 ABI / bytecode 与调用包装，不放业务逻辑；业务交互封装在 `counter/counter.go`。若修改合约，使用与模块一致的 go-ethereum `abigen v1.17.5` 重新生成并确认 diff。

## 目录结构

```text
backend/homework05/
├── .env.example
├── README.md
├── go.mod
├── bindings/
│   ├── counter.go
│   └── counter_test.go
├── cmd/homework05/
│   ├── main.go
│   ├── main_test.go
├── contracts/
│   ├── Counter.sol
│   ├── Counter.abi
│   └── Counter.bin
├── counter/
│   └── counter.go
└── ethrpc/
    ├── client.go
    └── client_test.go
```

## 安全边界

- 私钥只从 `SEPOLIA_PRIVATE_KEY` 读取，不写入源码或日志；
- `transfer` 使用精确十进制转换，不经过 `float64`；
- ETH 转账使用 EIP-1559 动态费用：tip 由节点建议，fee cap 取 `2 × baseFee + tip`；
- 写操作会主动校验 chain ID 必须为 Sepolia `11155111`，防止误把测试私钥/金额发到其他网络；
- 仓库提交的是可复现代码与测试。真实链上写交易只有在运行环境提供 RPC、测试私钥和测试 ETH 时才会发生。

## 课程资料

- 课件：<https://github.com/MetaNodeAcademy/golang-learning/tree/main/lesson-04>
- 作业：<https://github.com/MetaNodeAcademy/LearningRoadmap/blob/main/backend/homework05.md>
- Go Ethereum Book：<https://goethereumbook.org/zh/>
- Ethereum JSON-RPC：<https://ethereum.org/zh/developers/docs/apis/json-rpc/>
