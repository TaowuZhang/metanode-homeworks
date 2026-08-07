import hre from "hardhat";
import { upgrades } from "@openzeppelin/hardhat-upgrades";

const DEFAULT_ETH_USD_FEED = "0x694AA1769357215DE4FAC081bf1f309aDC325306";
const DEFAULT_LINK_TOKEN = "0x779877A7B0D9E8603169DdbD7836e478b4624789";
const DEFAULT_LINK_USD_FEED = "0xc59E3633BAAC79493d908e63626716e204A45EdF";

async function main() {
  const connection = await hre.network.create();
  const { ethers } = connection;
  const upgradesApi = await upgrades(hre, connection);
  const [deployer] = await ethers.getSigners();

  const ethUsdFeed = process.env.SEPOLIA_ETH_USD_FEED ?? DEFAULT_ETH_USD_FEED;
  const linkToken = process.env.SEPOLIA_LINK_TOKEN ?? DEFAULT_LINK_TOKEN;
  const linkUsdFeed = process.env.SEPOLIA_LINK_USD_FEED ?? DEFAULT_LINK_USD_FEED;

  const NFT = await ethers.getContractFactory("MetaNodeNFT");
  const nft = await NFT.deploy();
  await nft.waitForDeployment();

  const Auction = await ethers.getContractFactory("NFTAuction");
  const auction = await upgradesApi.deployProxy(Auction, [deployer.address], { kind: "uups" });
  await auction.waitForDeployment();

  await (await auction.setPriceFeed(ethers.ZeroAddress, ethUsdFeed)).wait();
  await (await auction.setPriceFeed(linkToken, linkUsdFeed)).wait();

  const auctionAddress = await auction.getAddress();
  const implementationAddress = await upgradesApi.erc1967.getImplementationAddress(auctionAddress);
  const chainId = (await ethers.provider.getNetwork()).chainId;
  const network = chainId === 11155111n ? "sepolia" : `chain-${chainId}`;

  console.log(JSON.stringify({
    network,
    chainId: chainId.toString(),
    deployer: deployer.address,
    nft: await nft.getAddress(),
    auctionProxy: auctionAddress,
    auctionImplementation: implementationAddress,
    feeds: {
      ETH_USD: ethUsdFeed,
      LINK_USD: linkUsdFeed,
    },
    erc20: {
      LINK: linkToken,
    },
  }, null, 2));
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
