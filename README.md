# MetaNodeAcademy 作业总仓库

这是个人 MetaNodeAcademy 学习与作业的统一仓库。Go 后端基础作业与 Solidity 智能合约作业都在这里维护，后续 MetaNode 作业继续按课程/方向放入对应子目录，不再为每次作业单独创建仓库。

## 内容总览

### Go 后端基础

现有 Go 内容保持原有目录结构：

- `internal/homework01`：基础算法练习，包括只出现一次的数字、回文数、有效括号、最长公共前缀、加一、原地删除有序数组重复项、合并区间、两数之和。
- `internal/homework02`：指针、切片、goroutine、WaitGroup、任务调度器、接口与组合、channel、Mutex、atomic。
- `internal/gormpractice`：User/Post/Comment 模型、关联查询、Preload、统计查询与 GORM Hook。
- `internal/blog`：Gin + GORM + SQLite + JWT + bcrypt 博客后端，包含注册登录、profile、文章 CRUD、评论和作者权限校验。
- `cmd/`：各 Go 作业与博客项目的运行入口。
- `docs/`：API、测试、实现报告与自然语言到代码的学习记录。

### Solidity 智能合约基础

- [`solidity/homework03`](./solidity/homework03)：MetaNode Solidity 基础二「任务3」——可升级 NFT 拍卖市场。
  - ERC721 NFT mint / transfer / escrow；
  - ETH 与 ERC20 跨币种竞价；
  - Chainlink Data Feed 统一换算 USD；
  - pull-payment 退款与拍卖结算；
  - OpenZeppelin UUPS V1 → V2 升级；
  - Hardhat 3 + Mocha：18/18 tests passing；
  - 96.00% line / 96.30% statement coverage；
  - 已完成 Sepolia 实际部署、两个 Chainlink feed 配置与 V2 升级；
  - 真实合约地址与交易证据见 [`solidity/homework03/DEPLOYMENTS.md`](./solidity/homework03/DEPLOYMENTS.md)。

## 仓库结构

```text
cmd/                         # Go 可执行入口
internal/                    # Go 作业与博客实现
docs/                        # Go 文档与学习记录
solidity/
  homework03/                # Solidity NFT Auction 作业（独立 Hardhat 项目）
.github/workflows/ci.yml     # 现有 Go CI
```

## Go 环境与验证

技术栈：Go 1.26、Gin、GORM + SQLite、JWT、bcrypt、godotenv、httptest。

```bash
go mod tidy
go vet ./...
go test ./...
go test -race ./...
```

运行示例：

```bash
go run ./cmd/homework01
go run ./cmd/homework02
go run ./cmd/gorm_practice
go run ./cmd/blog
```

博客环境变量示例见仓库根目录 `.env.example`。

## Solidity homework03 环境与验证

```bash
cd solidity/homework03
npm install
npm run compile
npm test
npm run coverage
```

Sepolia 部署与升级方式、测试报告和链上地址分别见：

- [`solidity/homework03/README.md`](./solidity/homework03/README.md)
- [`solidity/homework03/TEST_REPORT.md`](./solidity/homework03/TEST_REPORT.md)
- [`solidity/homework03/DEPLOYMENTS.md`](./solidity/homework03/DEPLOYMENTS.md)

## 统一维护约定

MetaNodeAcademy 相关学习成果以本仓库为唯一个人 GitHub 母本。后续新增作业优先在现有课程目录下继续扩展，例如 `solidity/homework04/`，避免按单次任务拆分出平行仓库。
