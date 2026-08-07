import hre from "hardhat";
import { upgrades } from "@openzeppelin/hardhat-upgrades";

async function main() {
  const proxyAddress = process.env.AUCTION_PROXY_ADDRESS;
  if (!proxyAddress) {
    throw new Error("AUCTION_PROXY_ADDRESS is required");
  }

  const minimumBidIncrementBps = BigInt(process.env.MIN_BID_INCREMENT_BPS ?? "500");

  const connection = await hre.network.create();
  const { ethers } = connection;
  const upgradesApi = await upgrades(hre, connection);

  const AuctionV2 = await ethers.getContractFactory("NFTAuctionV2");
  const upgraded = await upgradesApi.upgradeProxy(proxyAddress, AuctionV2, { kind: "uups" });
  await upgraded.waitForDeployment();
  await (await upgraded.setMinimumBidIncrementBps(minimumBidIncrementBps)).wait();

  const implementationAddress = await upgradesApi.erc1967.getImplementationAddress(proxyAddress);

  console.log(JSON.stringify({
    network: "sepolia",
    auctionProxy: proxyAddress,
    auctionImplementation: implementationAddress,
    version: await upgraded.version(),
    minimumBidIncrementBps: (await upgraded.minimumBidIncrementBps()).toString(),
  }, null, 2));
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
