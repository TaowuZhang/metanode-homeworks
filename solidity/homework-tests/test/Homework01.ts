import { expect } from "chai";
import hre from "hardhat";

describe("homework01", function () {
  let ethers: any;

  before(async function () {
    const connection = await hre.network.create();
    ({ ethers } = connection);
  });

  it("records votes, reads them, and resets with epoch", async function () {
    const voting = await ethers.deployContract("Voting");

    await (await voting.vote(1n)).wait();
    await (await voting.vote(1n)).wait();
    await (await voting.vote(2n)).wait();

    expect(await voting.getVotes(1n)).to.equal(2n);
    expect(await voting.getVotes(2n)).to.equal(1n);
    expect(await voting.getVotes(999n)).to.equal(0n);

    await (await voting.resetVotes()).wait();
    expect(await voting.getVotes(1n)).to.equal(0n);
    expect(await voting.getVotes(2n)).to.equal(0n);

    await (await voting.vote(1n)).wait();
    expect(await voting.getVotes(1n)).to.equal(1n);
  });

  it("reverses an ASCII string", async function () {
    const reverse = await ethers.deployContract("ReverseString");
    expect(await reverse.reverse("abcde")).to.equal("edcba");
    expect(await reverse.reverse("a")).to.equal("a");
    expect(await reverse.reverse("")).to.equal("");
  });

  it("converts integers to roman numerals", async function () {
    const converter = await ethers.deployContract("IntegerToRoman");
    expect(await converter.intToRoman(3749n)).to.equal("MMMDCCXLIX");
    expect(await converter.intToRoman(1994n)).to.equal("MCMXCIV");
    expect(await converter.intToRoman(58n)).to.equal("LVIII");
    await expect(converter.intToRoman(0n)).to.be.revertedWithCustomError(converter, "OutOfRange");
    await expect(converter.intToRoman(4000n)).to.be.revertedWithCustomError(converter, "OutOfRange");
  });

  it("converts roman numerals to integers", async function () {
    const converter = await ethers.deployContract("RomanToInteger");
    expect(await converter.romanToInt("MCMXCIV")).to.equal(1994n);
    expect(await converter.romanToInt("IV")).to.equal(4n);
    expect(await converter.romanToInt("LVIII")).to.equal(58n);
    await expect(converter.romanToInt("Z")).to.be.revertedWithCustomError(converter, "InvalidRoman");
  });

  it("merges two sorted arrays", async function () {
    const merger = await ethers.deployContract("MergeSortedArray");
    expect(await merger.mergeSorted([1n, 3n, 5n], [2n, 4n, 6n])).to.deep.equal([1n, 2n, 3n, 4n, 5n, 6n]);
    expect(await merger.mergeSorted([1n, 2n], [])).to.deep.equal([1n, 2n]);
    expect(await merger.mergeSorted([], [7n])).to.deep.equal([7n]);
  });

  it("binary-searches a sorted array", async function () {
    const searcher = await ethers.deployContract("BinarySearch");
    expect(await searcher.binarySearch([1n, 3n, 5n, 7n, 9n], 5n)).to.equal(2n);
    expect(await searcher.binarySearch([1n, 3n, 5n, 7n, 9n], 1n)).to.equal(0n);
    expect(await searcher.binarySearch([1n, 3n, 5n, 7n, 9n], 9n)).to.equal(4n);
    expect(await searcher.binarySearch([1n, 3n, 5n, 7n, 9n], 4n)).to.equal(-1n);
    expect(await searcher.binarySearch([], 1n)).to.equal(-1n);
  });
});
