# Deployments

## Local Hardhat smoke

Deployment script was executed successfully on Hardhat chain `31337` on 2026-08-07. The local addresses are ephemeral and are not submission addresses.

The smoke proved that the script can deploy:

- `MetaNodeNFT`
- `NFTAuction` UUPS proxy
- `NFTAuction` implementation
- ETH/USD and LINK/USD feed configuration

## Sepolia

Status: **deployed and upgraded successfully on 2026-08-07** (`chainId = 11155111`).

Deployment and configuration were signed with the funded Sepolia test account through MetaMask. No private key or seed phrase was exported to the project.

### Contracts

| Component | Address | Deployment / upgrade transaction |
| --- | --- | --- |
| `MetaNodeNFT` | `0x9407360B3D01a82E8d70BBd584Ce2563F16447d4` | `0x135744ac049bf185a03963cfc856ed409cee0862cda255b4341a24717c78d604` |
| `NFTAuction` V1 implementation | `0xDC5D2eCB9D51B4a6D1b52cC5e6B3E7D5f9C02eA5` | `0x13ce063089ee3b587516862c9467812e875e1bbd56c8460cbb7831f4c9b598eb` |
| `NFTAuction` UUPS / ERC1967 proxy | `0xE1e2165313Ab655cfB3223083c7d706123fa2eDB` | `0xd265c0269e16592ec44eb00c175970ffb944b22a34c61e5fe12652c1e6efa26f` |
| `NFTAuctionV2` implementation | `0xF1e424623ef4A19C09B07008BFe762105856DcAf` | `0x5fafd1ae747cd8e1815dc183ea76d437fac636eeb32245a6480a460e2ec0790a` |
| Proxy upgrade V1 -> V2 | proxy unchanged | `0x98714c0ba10afd077390a5aeb12c7ce34550254f6a92dbeb2d4d433e7e100d1e` |

The upgrade transaction used `upgradeToAndCall` to switch the proxy to the V2 implementation and set `minimumBidIncrementBps = 500` (5%) atomically.

### Chainlink configuration

| Asset | Token | Feed | Configuration transaction |
| --- | --- | --- | --- |
| ETH / USD | native ETH (`address(0)`) | `0x694AA1769357215DE4FAC081bf1f309aDC325306` | `0x526df65a126c3e4892b96512115d53aeb8a8b74111253e67618a1424b1740051` |
| LINK / USD | `0x779877A7B0D9E8603169DdbD7836e478b4624789` | `0xc59E3633BAAC79493d908e63626716e204A45EdF` | `0xf9b4a50972331ed2ab7b89659213569af91c8f7d7c39887a9a3446bfa88987d7` |

### Post-deployment verification

Direct Sepolia reads after the upgrade verified all of the following:

- the proxy ERC1967 implementation slot points to `0xF1e424623ef4A19C09B07008BFe762105856DcAf`;
- `version()` returns `NFTAuction V2`;
- `minimumBidIncrementBps()` returns `500`;
- proxy ownership remained with the deploying Sepolia account;
- `priceFeeds(address(0))` still points to the ETH / USD feed;
- `priceFeeds(LINK)` still points to the LINK / USD feed;
- every deployment, configuration and upgrade receipt listed above returned `status = 0x1`.
