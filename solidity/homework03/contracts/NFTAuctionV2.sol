// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {NFTAuction} from "./NFTAuction.sol";

/// @custom:oz-upgrades-unsafe-allow missing-initializer
contract NFTAuctionV2 is NFTAuction {
    uint256 public minimumBidIncrementBps;

    event MinimumBidIncrementUpdated(uint256 minimumBidIncrementBps);

    error InvalidBidIncrement();

    function setMinimumBidIncrementBps(uint256 minimumBidIncrementBps_) external onlyOwner {
        _setMinimumBidIncrementBps(minimumBidIncrementBps_);
    }

    function version() external pure override returns (string memory) {
        return "NFTAuction V2";
    }

    function _minimumNextBidUsd(Auction storage auction) internal view override returns (uint256) {
        if (minimumBidIncrementBps == 0) {
            return super._minimumNextBidUsd(auction);
        }

        uint256 increment = (auction.highestBidUsd * minimumBidIncrementBps) / 10_000;
        if (increment == 0) increment = 1;
        return auction.highestBidUsd + increment;
    }

    function _setMinimumBidIncrementBps(uint256 minimumBidIncrementBps_) internal {
        if (minimumBidIncrementBps_ > 10_000) revert InvalidBidIncrement();
        minimumBidIncrementBps = minimumBidIncrementBps_;
        emit MinimumBidIncrementUpdated(minimumBidIncrementBps_);
    }
}
