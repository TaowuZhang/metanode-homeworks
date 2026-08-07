// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

contract NFTAuction is Initializable, UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuardTransient, ERC721Holder {
    using SafeERC20 for IERC20;

    address public constant NATIVE_TOKEN = address(0);
    uint256 public constant USD_SCALE = 1e18;

    struct Auction {
        address seller;
        address nft;
        uint256 tokenId;
        uint64 endTime;
        uint256 reservePriceUsd;
        address highestBidder;
        address highestPaymentToken;
        uint256 highestBidAmount;
        uint256 highestBidUsd;
        bool settled;
    }

    mapping(uint256 => Auction) public auctions;
    mapping(address => address) public priceFeeds;
    mapping(address => mapping(address => uint256)) public pendingReturns;
    uint256 public nextAuctionId;

    event PriceFeedSet(address indexed token, address indexed feed);
    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed seller,
        address indexed nft,
        uint256 tokenId,
        uint256 reservePriceUsd,
        uint64 endTime
    );
    event BidPlaced(
        uint256 indexed auctionId,
        address indexed bidder,
        address indexed paymentToken,
        uint256 amount,
        uint256 amountUsd
    );
    event RefundWithdrawn(address indexed bidder, address indexed paymentToken, uint256 amount);
    event AuctionSettled(
        uint256 indexed auctionId,
        address indexed seller,
        address indexed winner,
        address paymentToken,
        uint256 amount,
        uint256 amountUsd
    );

    error ZeroAddress();
    error InvalidDuration();
    error InvalidReservePrice();
    error InvalidAuction();
    error AuctionAlreadySettled();
    error AuctionEnded();
    error AuctionNotEnded();
    error SellerCannotBid();
    error InvalidBidAmount();
    error BidTooLow(uint256 bidUsd, uint256 requiredUsd);
    error PriceFeedNotSet(address token);
    error InvalidPrice(address feed);
    error UnsupportedDecimals(uint8 decimalsValue);
    error NoRefundAvailable();
    error EthTransferFailed();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) external initializer {
        if (initialOwner == address(0)) revert ZeroAddress();
        __Ownable_init(initialOwner);
        nextAuctionId = 1;
    }

    function setPriceFeed(address token, address feed) external onlyOwner {
        if (feed == address(0)) revert ZeroAddress();
        priceFeeds[token] = feed;
        emit PriceFeedSet(token, feed);
    }

    function createAuction(
        address nft,
        uint256 tokenId,
        uint256 reservePriceUsd,
        uint256 duration
    ) external nonReentrant returns (uint256 auctionId) {
        if (nft == address(0)) revert ZeroAddress();
        if (duration == 0 || block.timestamp + duration > type(uint64).max) revert InvalidDuration();
        if (reservePriceUsd == 0) revert InvalidReservePrice();

        auctionId = nextAuctionId++;
        uint64 endTime = uint64(block.timestamp + duration);

        auctions[auctionId] = Auction({
            seller: msg.sender,
            nft: nft,
            tokenId: tokenId,
            endTime: endTime,
            reservePriceUsd: reservePriceUsd,
            highestBidder: address(0),
            highestPaymentToken: address(0),
            highestBidAmount: 0,
            highestBidUsd: 0,
            settled: false
        });

        IERC721(nft).safeTransferFrom(msg.sender, address(this), tokenId);
        emit AuctionCreated(auctionId, msg.sender, nft, tokenId, reservePriceUsd, endTime);
    }

    function bidWithEth(uint256 auctionId) external payable nonReentrant {
        uint256 bidUsd = _validateBid(auctionId, NATIVE_TOKEN, msg.value, msg.sender);
        _recordBid(auctionId, NATIVE_TOKEN, msg.value, bidUsd, msg.sender);
    }

    function bidWithToken(uint256 auctionId, address token, uint256 amount) external nonReentrant {
        if (token == NATIVE_TOKEN) revert ZeroAddress();
        uint256 bidUsd = _validateBid(auctionId, token, amount, msg.sender);
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        _recordBid(auctionId, token, amount, bidUsd, msg.sender);
    }

    function withdrawRefund(address paymentToken) external nonReentrant {
        uint256 amount = pendingReturns[msg.sender][paymentToken];
        if (amount == 0) revert NoRefundAvailable();
        pendingReturns[msg.sender][paymentToken] = 0;
        _payout(paymentToken, msg.sender, amount);
        emit RefundWithdrawn(msg.sender, paymentToken, amount);
    }

    function settleAuction(uint256 auctionId) external nonReentrant {
        Auction storage auction = _getAuction(auctionId);
        if (auction.settled) revert AuctionAlreadySettled();
        if (block.timestamp < auction.endTime) revert AuctionNotEnded();

        auction.settled = true;
        address winner = auction.highestBidder;

        if (winner == address(0)) {
            IERC721(auction.nft).safeTransferFrom(address(this), auction.seller, auction.tokenId);
            emit AuctionSettled(auctionId, auction.seller, address(0), address(0), 0, 0);
            return;
        }

        IERC721(auction.nft).safeTransferFrom(address(this), winner, auction.tokenId);
        _payout(auction.highestPaymentToken, auction.seller, auction.highestBidAmount);
        emit AuctionSettled(
            auctionId,
            auction.seller,
            winner,
            auction.highestPaymentToken,
            auction.highestBidAmount,
            auction.highestBidUsd
        );
    }

    function quoteUsd(address token, uint256 amount) public view returns (uint256) {
        if (amount == 0) return 0;
        address feed = priceFeeds[token];
        if (feed == address(0)) revert PriceFeedNotSet(token);

        AggregatorV3Interface dataFeed = AggregatorV3Interface(feed);
        (uint80 roundId, int256 answer,, uint256 updatedAt, uint80 answeredInRound) = dataFeed.latestRoundData();
        if (answer <= 0 || updatedAt == 0 || answeredInRound < roundId) revert InvalidPrice(feed);

        uint8 tokenDecimals = token == NATIVE_TOKEN ? 18 : IERC20Metadata(token).decimals();
        uint8 feedDecimals = dataFeed.decimals();
        if (tokenDecimals > 36) revert UnsupportedDecimals(tokenDecimals);
        if (feedDecimals > 36) revert UnsupportedDecimals(feedDecimals);

        uint256 normalizedPrice = uint256(answer);
        if (feedDecimals < 18) {
            normalizedPrice *= 10 ** (18 - feedDecimals);
        } else if (feedDecimals > 18) {
            normalizedPrice /= 10 ** (feedDecimals - 18);
        }

        return Math.mulDiv(amount, normalizedPrice, 10 ** tokenDecimals);
    }

    function version() external pure virtual returns (string memory) {
        return "NFTAuction V1";
    }

    function _validateBid(
        uint256 auctionId,
        address paymentToken,
        uint256 amount,
        address bidder
    ) internal view returns (uint256 bidUsd) {
        Auction storage auction = _getAuction(auctionId);
        if (auction.settled) revert AuctionAlreadySettled();
        if (block.timestamp >= auction.endTime) revert AuctionEnded();
        if (bidder == auction.seller) revert SellerCannotBid();
        if (amount == 0) revert InvalidBidAmount();

        bidUsd = quoteUsd(paymentToken, amount);
        uint256 requiredUsd = auction.highestBidder == address(0)
            ? auction.reservePriceUsd
            : _minimumNextBidUsd(auction);
        if (bidUsd < requiredUsd) revert BidTooLow(bidUsd, requiredUsd);
    }

    function _recordBid(
        uint256 auctionId,
        address paymentToken,
        uint256 amount,
        uint256 bidUsd,
        address bidder
    ) internal {
        Auction storage auction = auctions[auctionId];
        if (auction.highestBidder != address(0)) {
            pendingReturns[auction.highestBidder][auction.highestPaymentToken] += auction.highestBidAmount;
        }

        auction.highestBidder = bidder;
        auction.highestPaymentToken = paymentToken;
        auction.highestBidAmount = amount;
        auction.highestBidUsd = bidUsd;

        emit BidPlaced(auctionId, bidder, paymentToken, amount, bidUsd);
    }

    function _minimumNextBidUsd(Auction storage auction) internal view virtual returns (uint256) {
        return auction.highestBidUsd + 1;
    }

    function _getAuction(uint256 auctionId) internal view returns (Auction storage auction) {
        auction = auctions[auctionId];
        if (auction.seller == address(0)) revert InvalidAuction();
    }

    function _payout(address paymentToken, address recipient, uint256 amount) internal {
        if (paymentToken == NATIVE_TOKEN) {
            (bool success,) = payable(recipient).call{value: amount}("");
            if (!success) revert EthTransferFailed();
        } else {
            IERC20(paymentToken).safeTransfer(recipient, amount);
        }
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
