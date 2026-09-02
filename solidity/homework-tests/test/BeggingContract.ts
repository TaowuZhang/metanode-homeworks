import { expect } from "chai";
import hre from "hardhat";

describe("homework02 BeggingContract", function () {
  let ethers: any;
  let owner: any;
  let donorA: any;
  let donorB: any;
  let donorC: any;
  let stranger: any;

  before(async function () {
    const connection = await hre.network.create();
    ({ ethers } = connection);
  });

  async function deploy() {
    [owner, donorA, donorB, donorC, stranger] = await ethers.getSigners();
    const begging = await ethers.deployContract("BeggingContract");
    return begging;
  }

  it("records donations and lets the owner withdraw with transfer", async function () {
    const begging = await deploy();
    const donation = ethers.parseEther("0.01");

    await expect(begging.connect(donorA).donate({ value: donation }))
      .to.emit(begging, "Donation")
      .withArgs(donorA.address, donation, donation);

    expect(await begging.getDonation(donorA.address)).to.equal(donation);
    expect(await ethers.provider.getBalance(await begging.getAddress())).to.equal(donation);

    const ownerBefore = await ethers.provider.getBalance(owner.address);
    const tx = await begging.connect(owner).withdraw();
    const receipt = await tx.wait();

    expect(await begging.getDonation(donorA.address)).to.equal(donation);
    expect(await ethers.provider.getBalance(await begging.getAddress())).to.equal(0n);
    expect(await ethers.provider.getBalance(owner.address)).to.equal(ownerBefore + donation - receipt.fee);
  });

  it("counts plain ETH transfers through receive()", async function () {
    const begging = await deploy();
    const amount = ethers.parseEther("0.002");

    await donorB.sendTransaction({ to: await begging.getAddress(), value: amount });
    expect(await begging.getDonation(donorB.address)).to.equal(amount);
  });

  it("rejects zero-value donate and non-owner withdraw", async function () {
    const begging = await deploy();

    await expect(begging.connect(donorA).donate({ value: 0n })).to.be.revertedWithCustomError(
      begging,
      "ZeroDonation"
    );
    await expect(begging.connect(stranger).withdraw()).to.be.revertedWithCustomError(begging, "NotOwner");
    await expect(begging.connect(owner).withdraw()).to.be.revertedWithCustomError(begging, "NothingToWithdraw");
  });

  it("keeps the top 3 donors by cumulative amount", async function () {
    const begging = await deploy();

    await begging.connect(donorA).donate({ value: ethers.parseEther("1") });
    await begging.connect(donorB).donate({ value: ethers.parseEther("3") });
    await begging.connect(donorC).donate({ value: ethers.parseEther("2") });
    await begging.connect(donorA).donate({ value: ethers.parseEther("3") });

    const [donors, amounts] = await begging.getTopDonors();
    expect(donors[0]).to.equal(donorA.address);
    expect(amounts[0]).to.equal(ethers.parseEther("4"));
    expect(donors[1]).to.equal(donorB.address);
    expect(amounts[1]).to.equal(ethers.parseEther("3"));
    expect(donors[2]).to.equal(donorC.address);
    expect(amounts[2]).to.equal(ethers.parseEther("2"));
  });

  it("blocks donations outside an owner-configured window", async function () {
    const begging = await deploy();
    const latest = await ethers.provider.getBlock("latest");
    const now = BigInt(latest.timestamp);

    await begging.connect(owner).setDonationWindow(now + 100n, now + 200n, true);
    await expect(
      begging.connect(donorA).donate({ value: ethers.parseEther("0.001") })
    ).to.be.revertedWithCustomError(begging, "DonationClosed");

    await ethers.provider.send("evm_increaseTime", [150]);
    await ethers.provider.send("evm_mine", []);
    await begging.connect(donorA).donate({ value: ethers.parseEther("0.001") });
    expect(await begging.getDonation(donorA.address)).to.equal(ethers.parseEther("0.001"));
  });
});
