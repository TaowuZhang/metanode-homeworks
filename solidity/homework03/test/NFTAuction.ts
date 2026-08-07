import { expect } from "chai";
import hre from "hardhat";
import { upgrades } from "@openzeppelin/hardhat-upgrades";

const USD = 10n ** 18n;

describe("NFTAuction homework03", function () {
  let connection: any;
  let ethers: any;
  let upgradesApi: any;
  let owner: any;
  let seller: any;
  let ethBidder: any;
  let tokenBidder: any;
  let auction: any;
  let nft: any;
  let usdc: any;
  let ethFeed: any;
  let usdcFeed: any;

  before(async function () {
    connection = await hre.network.create();
    ({ ethers } = connection);
    upgradesApi = await upgrades(hre, connection);
  });

  beforeEach(async function () {
    [owner, seller, ethBidder, tokenBidder] = await ethers.getSigners();

    const Auction = await ethers.getContractFactory("NFTAuction");
    auction = await upgradesApi.deployProxy(Auction, [owner.address], { kind: "uups" });
    await auction.waitForDeployment();

    const NFT = await ethers.getContractFactory("MetaNodeNFT");
    nft = await NFT.deploy();
    await nft.waitForDeployment();

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    usdc = await MockERC20.deploy("Mock USDC", "mUSDC", 6);
    await usdc.waitForDeployment();

    const MockFeed = await ethers.getContractFactory("MockV3Aggregator");
    ethFeed = await MockFeed.deploy(8, 3000n * 10n ** 8n);
    usdcFeed = await MockFeed.deploy(8, 1n * 10n ** 8n);
    await ethFeed.waitForDeployment();
    await usdcFeed.waitForDeployment();

    await auction.connect(owner).setPriceFeed(ethers.ZeroAddress, await ethFeed.getAddress());
    await auction.connect(owner).setPriceFeed(await usdc.getAddress(), await usdcFeed.getAddress());

    await nft.mint(seller.address);
    await nft.connect(seller).approve(await auction.getAddress(), 1n);
    await usdc.mint(tokenBidder.address, ethers.parseUnits("10000", 6));
    await usdc.connect(tokenBidder).approve(await auction.getAddress(), ethers.MaxUint256);
  });

  async function createAuction(reserveUsd = 1000n * USD, duration = 3600n) {
    await auction.connect(seller).createAuction(await nft.getAddress(), 1n, reserveUsd, duration);
    return 1n;
  }

  it("normalizes ETH and ERC20 Chainlink feed prices to 18-decimal USD", async function () {
    expect(await auction.quoteUsd(ethers.ZeroAddress, ethers.parseEther("1"))).to.equal(3000n * USD);
    expect(await auction.quoteUsd(await usdc.getAddress(), ethers.parseUnits("250", 6))).to.equal(250n * USD);
  });

  it("escrows an ERC721 when the seller creates an auction", async function () {
    const auctionId = await createAuction();
    const data = await auction.auctions(auctionId);

    expect(data.seller).to.equal(seller.address);
    expect(data.reservePriceUsd).to.equal(1000n * USD);
    expect(await nft.ownerOf(1n)).to.equal(await auction.getAddress());
  });

  it("compares ETH and ERC20 bids in USD and credits the displaced bidder refund", async function () {
    const auctionId = await createAuction();

    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.5") }); // $1,500
    await auction
      .connect(tokenBidder)
      .bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1600", 6)); // $1,600

    const data = await auction.auctions(auctionId);
    expect(data.highestBidder).to.equal(tokenBidder.address);
    expect(data.highestPaymentToken).to.equal(await usdc.getAddress());
    expect(data.highestBidUsd).to.equal(1600n * USD);
    expect(await auction.pendingReturns(ethBidder.address, ethers.ZeroAddress)).to.equal(ethers.parseEther("0.5"));
  });

  it("lets an outbid user pull their refund safely", async function () {
    const auctionId = await createAuction();
    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.5") });
    await auction
      .connect(tokenBidder)
      .bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1600", 6));

    const contractBefore = await ethers.provider.getBalance(await auction.getAddress());
    await auction.connect(ethBidder).withdrawRefund(ethers.ZeroAddress);
    const contractAfter = await ethers.provider.getBalance(await auction.getAddress());

    expect(contractBefore - contractAfter).to.equal(ethers.parseEther("0.5"));
    expect(await auction.pendingReturns(ethBidder.address, ethers.ZeroAddress)).to.equal(0n);
  });

  it("rejects seller self-bids and bids below the current USD value", async function () {
    const auctionId = await createAuction();

    await expect(
      auction.connect(seller).bidWithEth(auctionId, { value: ethers.parseEther("0.5") })
    ).to.be.revertedWithCustomError(auction, "SellerCannotBid");

    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.5") });
    await expect(
      auction.connect(tokenBidder).bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1499", 6))
    ).to.be.revertedWithCustomError(auction, "BidTooLow");
  });

  it("settles an ERC20-winning auction by transferring NFT and funds", async function () {
    const auctionId = await createAuction();
    const bidAmount = ethers.parseUnits("1600", 6);
    await auction.connect(tokenBidder).bidWithToken(auctionId, await usdc.getAddress(), bidAmount);

    await ethers.provider.send("evm_increaseTime", [3601]);
    await ethers.provider.send("evm_mine", []);

    const sellerBefore = await usdc.balanceOf(seller.address);
    await auction.settleAuction(auctionId);

    expect(await nft.ownerOf(1n)).to.equal(tokenBidder.address);
    expect((await usdc.balanceOf(seller.address)) - sellerBefore).to.equal(bidAmount);
    expect((await auction.auctions(auctionId)).settled).to.equal(true);
  });

  it("returns the NFT to the seller if an auction ends without bids", async function () {
    const auctionId = await createAuction();
    await ethers.provider.send("evm_increaseTime", [3601]);
    await ethers.provider.send("evm_mine", []);

    await auction.settleAuction(auctionId);
    expect(await nft.ownerOf(1n)).to.equal(seller.address);
  });

  it("preserves state across a UUPS upgrade and activates the V2 minimum increment", async function () {
    const auctionId = await createAuction();
    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.5") }); // $1,500

    const AuctionV2 = await ethers.getContractFactory("NFTAuctionV2");
    const upgraded = await upgradesApi.upgradeProxy(await auction.getAddress(), AuctionV2, {
      kind: "uups",
    });
    await upgraded.waitForDeployment();
    await upgraded.connect(owner).setMinimumBidIncrementBps(500);

    expect(await upgraded.version()).to.equal("NFTAuction V2");
    expect(await upgraded.minimumBidIncrementBps()).to.equal(500n);
    expect((await upgraded.auctions(auctionId)).highestBidder).to.equal(ethBidder.address);

    await expect(
      upgraded.connect(tokenBidder).bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1574", 6))
    ).to.be.revertedWithCustomError(upgraded, "BidTooLow");

    await upgraded
      .connect(tokenBidder)
      .bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1575", 6));
    expect((await upgraded.auctions(auctionId)).highestBidder).to.equal(tokenBidder.address);
  });

  it("prevents a non-owner from authorizing a UUPS upgrade", async function () {
    const AuctionV2 = await ethers.getContractFactory("NFTAuctionV2");
    const implementationV2 = await AuctionV2.deploy();
    await implementationV2.waitForDeployment();

    await expect(
      auction.connect(ethBidder).upgradeToAndCall(await implementationV2.getAddress(), "0x")
    ).to.be.revertedWithCustomError(auction, "OwnableUnauthorizedAccount");
  });

  it("restricts feed configuration to the owner", async function () {
    await expect(
      auction.connect(ethBidder).setPriceFeed(ethers.ZeroAddress, await ethFeed.getAddress())
    ).to.be.revertedWithCustomError(auction, "OwnableUnauthorizedAccount");
    await expect(
      auction.connect(owner).setPriceFeed(ethers.ZeroAddress, ethers.ZeroAddress)
    ).to.be.revertedWithCustomError(auction, "ZeroAddress");
  });

  it("rejects malformed auction creation and invalid auction ids", async function () {
    await expect(
      auction.connect(seller).createAuction(ethers.ZeroAddress, 1n, 1000n * USD, 3600n)
    ).to.be.revertedWithCustomError(auction, "ZeroAddress");
    await expect(
      auction.connect(seller).createAuction(await nft.getAddress(), 1n, 1000n * USD, 0n)
    ).to.be.revertedWithCustomError(auction, "InvalidDuration");
    await expect(
      auction.connect(seller).createAuction(await nft.getAddress(), 1n, 0n, 3600n)
    ).to.be.revertedWithCustomError(auction, "InvalidReservePrice");
    await expect(
      auction.connect(ethBidder).bidWithEth(999n, { value: ethers.parseEther("1") })
    ).to.be.revertedWithCustomError(auction, "InvalidAuction");
  });

  it("rejects zero-value, ended, and zero-address token bids", async function () {
    const auctionId = await createAuction();
    await expect(
      auction.connect(ethBidder).bidWithEth(auctionId, { value: 0n })
    ).to.be.revertedWithCustomError(auction, "InvalidBidAmount");
    await expect(
      auction.connect(tokenBidder).bidWithToken(auctionId, ethers.ZeroAddress, 1n)
    ).to.be.revertedWithCustomError(auction, "ZeroAddress");

    await ethers.provider.send("evm_increaseTime", [3601]);
    await ethers.provider.send("evm_mine", []);
    await expect(
      auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("1") })
    ).to.be.revertedWithCustomError(auction, "AuctionEnded");
  });

  it("rejects early or repeated settlement", async function () {
    const auctionId = await createAuction();
    await expect(auction.settleAuction(auctionId)).to.be.revertedWithCustomError(auction, "AuctionNotEnded");

    await ethers.provider.send("evm_increaseTime", [3601]);
    await ethers.provider.send("evm_mine", []);
    await auction.settleAuction(auctionId);
    await expect(auction.settleAuction(auctionId)).to.be.revertedWithCustomError(auction, "AuctionAlreadySettled");
  });

  it("settles an ETH-winning auction and pays the seller", async function () {
    const auctionId = await createAuction();
    const bid = ethers.parseEther("0.5");
    await auction.connect(ethBidder).bidWithEth(auctionId, { value: bid });

    await ethers.provider.send("evm_increaseTime", [3601]);
    await ethers.provider.send("evm_mine", []);
    const sellerBefore = await ethers.provider.getBalance(seller.address);
    const tx = await auction.connect(tokenBidder).settleAuction(auctionId);
    await tx.wait();
    const sellerAfter = await ethers.provider.getBalance(seller.address);

    expect(sellerAfter - sellerBefore).to.equal(bid);
    expect(await nft.ownerOf(1n)).to.equal(ethBidder.address);
  });

  it("supports ERC20 refund withdrawal after a higher ETH bid", async function () {
    const auctionId = await createAuction();
    const tokenBid = ethers.parseUnits("1500", 6);
    await auction.connect(tokenBidder).bidWithToken(auctionId, await usdc.getAddress(), tokenBid);
    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.6") }); // $1,800

    const before = await usdc.balanceOf(tokenBidder.address);
    await auction.connect(tokenBidder).withdrawRefund(await usdc.getAddress());
    const after = await usdc.balanceOf(tokenBidder.address);

    expect(after - before).to.equal(tokenBid);
    expect(await auction.pendingReturns(tokenBidder.address, await usdc.getAddress())).to.equal(0n);
    await expect(
      auction.connect(tokenBidder).withdrawRefund(await usdc.getAddress())
    ).to.be.revertedWithCustomError(auction, "NoRefundAvailable");
  });

  it("validates Chainlink feed state and decimal normalization boundaries", async function () {
    expect(await auction.quoteUsd(ethers.ZeroAddress, 0n)).to.equal(0n);

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const unsupported = await MockERC20.deploy("Unsupported", "BAD", 37);
    await unsupported.waitForDeployment();
    await auction.connect(owner).setPriceFeed(await unsupported.getAddress(), await usdcFeed.getAddress());
    await expect(
      auction.quoteUsd(await unsupported.getAddress(), 1n)
    ).to.be.revertedWithCustomError(auction, "UnsupportedDecimals");

    const MockFeed = await ethers.getContractFactory("MockV3Aggregator");
    const highPrecisionFeed = await MockFeed.deploy(20, 3000n * 10n ** 20n);
    await highPrecisionFeed.waitForDeployment();
    await auction.connect(owner).setPriceFeed(ethers.ZeroAddress, await highPrecisionFeed.getAddress());
    expect(await auction.quoteUsd(ethers.ZeroAddress, ethers.parseEther("1"))).to.equal(3000n * USD);

    const unsupportedFeed = await MockFeed.deploy(37, 1n);
    await unsupportedFeed.waitForDeployment();
    await auction.connect(owner).setPriceFeed(ethers.ZeroAddress, await unsupportedFeed.getAddress());
    await expect(
      auction.quoteUsd(ethers.ZeroAddress, ethers.parseEther("1"))
    ).to.be.revertedWithCustomError(auction, "UnsupportedDecimals");

    const invalidFeed = await MockFeed.deploy(8, 1n);
    await invalidFeed.waitForDeployment();
    await invalidFeed.updateAnswer(0n);
    await auction.connect(owner).setPriceFeed(ethers.ZeroAddress, await invalidFeed.getAddress());
    await expect(
      auction.quoteUsd(ethers.ZeroAddress, ethers.parseEther("1"))
    ).to.be.revertedWithCustomError(auction, "InvalidPrice");

    await expect(
      auction.quoteUsd(ethBidder.address, 1n)
    ).to.be.revertedWithCustomError(auction, "PriceFeedNotSet");
  });

  it("covers NFT metadata helpers and the Chainlink mock interface", async function () {
    expect(await nft.nextTokenId()).to.equal(2n);
    await expect(nft.connect(seller).mint(seller.address)).to.be.revertedWithCustomError(nft, "OwnableUnauthorizedAccount");
    expect(await ethFeed.description()).to.equal("Mock V3 Aggregator");
    expect(await ethFeed.version()).to.equal(1n);
    const round = await ethFeed.getRoundData(42n);
    expect(round[0]).to.equal(42n);
    await ethFeed.updateAnswer(3100n * 10n ** 8n);
    expect((await ethFeed.latestRoundData())[1]).to.equal(3100n * 10n ** 8n);
  });

  it("enforces V2 bid increment configuration bounds and zero-config fallback", async function () {
    const auctionId = await createAuction();
    await auction.connect(ethBidder).bidWithEth(auctionId, { value: ethers.parseEther("0.5") });

    const AuctionV2 = await ethers.getContractFactory("NFTAuctionV2");
    const upgraded = await upgradesApi.upgradeProxy(await auction.getAddress(), AuctionV2, { kind: "uups" });
    await upgraded.waitForDeployment();

    await upgraded
      .connect(tokenBidder)
      .bidWithToken(auctionId, await usdc.getAddress(), ethers.parseUnits("1501", 6));
    await expect(
      upgraded.connect(owner).setMinimumBidIncrementBps(10_001)
    ).to.be.revertedWithCustomError(upgraded, "InvalidBidIncrement");
  });
});
