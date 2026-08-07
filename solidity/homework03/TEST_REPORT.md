# Test Report

Date: 2026-08-07

## Commands

```bash
npm install
npx hardhat compile
npx hardhat test
npx hardhat test --coverage
```

## Result

- Hardhat tests: **18 passing / 0 failing**
- Total line coverage: **96.00%**
- Total statement coverage: **96.30%**

### File coverage

| File | Line | Statement |
|---|---:|---:|
| `contracts/NFTAuction.sol` | 95.92% | 96.19% |
| `contracts/NFTAuctionV2.sol` | 90.00% | 92.31% |
| `contracts/MetaNodeNFT.sol` | 100.00% | 100.00% |
| `contracts/mocks/MockERC20.sol` | 100.00% | 100.00% |
| `contracts/mocks/MockV3Aggregator.sol` | 100.00% | 100.00% |

## Covered behavior

- ERC721 minting, custody and settlement transfer.
- Auction creation and invalid input rejection.
- ETH bids and ERC20 bids in the same auction.
- Chainlink-style `AggregatorV3Interface` price normalization into 18-decimal USD values.
- Cross-currency bid comparison.
- Pull-payment refunds for displaced ETH and ERC20 bidders.
- ETH and ERC20 seller settlement.
- No-bid NFT return.
- End-time and repeated-settlement guards.
- Missing/invalid price feed and decimal boundary handling.
- Owner-only feed management.
- UUPS upgrade authorization.
- V1 -> V2 storage preservation.
- V2 minimum bid increment behavior and configuration bounds.

The remaining uncovered Solidity lines are defensive branches that are difficult to trigger without deliberately malicious helper contracts (for example a recipient that always rejects ETH transfers) or direct implementation initialization attempts.
