// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title BeggingContract
/// @notice 作业2：讨饭合约。任何人可捐赠 ETH，owner 可提取全部资金。
/// @dev 必做：mapping / donate / withdraw / getDonation / onlyOwner。
///      额外挑战：Donation 事件、Top3 排行榜、捐赠时间窗。
contract BeggingContract {
    address public owner;
    mapping(address => uint256) private _donations;
    address[3] private _topDonors;

    uint256 public donationStart;
    uint256 public donationEnd;
    bool public timeRestricted;

    event Donation(address indexed donor, uint256 amount, uint256 totalFromDonor);
    event Withdrawal(address indexed owner, uint256 amount);
    event DonationWindowUpdated(uint256 start, uint256 end, bool timeRestricted);

    error NotOwner();
    error ZeroDonation();
    error DonationClosed();
    error InvalidWindow();
    error NothingToWithdraw();

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    /// @notice 直接往合约转 ETH 也记入捐赠，方便 MetaMask 普通转账测试。
    receive() external payable {
        _donate(msg.sender, msg.value);
    }

    function donate() external payable {
        _donate(msg.sender, msg.value);
    }

    function withdraw() external onlyOwner {
        uint256 amount = address(this).balance;
        if (amount == 0) revert NothingToWithdraw();
        payable(owner).transfer(amount);
        emit Withdrawal(owner, amount);
    }

    function getDonation(address donor) external view returns (uint256) {
        return _donations[donor];
    }

    function getTopDonors()
        external
        view
        returns (address[3] memory donors, uint256[3] memory amounts)
    {
        donors = _topDonors;
        for (uint256 i = 0; i < 3; ++i) {
            amounts[i] = _donations[_topDonors[i]];
        }
    }

    /// @notice 作业额外挑战：限制只能在 [start, end] 内捐赠。start==end==0 且 enabled=false 表示不限制。
    function setDonationWindow(uint256 start, uint256 end, bool enabled) external onlyOwner {
        if (enabled && start >= end) revert InvalidWindow();
        donationStart = start;
        donationEnd = end;
        timeRestricted = enabled;
        emit DonationWindowUpdated(start, end, enabled);
    }

    function _donate(address donor, uint256 amount) private {
        if (amount == 0) revert ZeroDonation();
        if (timeRestricted && (block.timestamp < donationStart || block.timestamp > donationEnd)) {
            revert DonationClosed();
        }

        _donations[donor] += amount;
        _updateTopDonors(donor);
        emit Donation(donor, amount, _donations[donor]);
    }

    function _updateTopDonors(address donor) private {
        uint256 amount = _donations[donor];

        for (uint256 i = 0; i < 3; ++i) {
            if (_topDonors[i] == donor) {
                _shiftLeftFrom(i);
                break;
            }
        }

        for (uint256 i = 0; i < 3; ++i) {
            address current = _topDonors[i];
            if (current == address(0) || amount > _donations[current]) {
                _shiftRightFrom(i);
                _topDonors[i] = donor;
                break;
            }
        }
    }

    function _shiftLeftFrom(uint256 index) private {
        for (uint256 i = index; i < 2; ++i) {
            _topDonors[i] = _topDonors[i + 1];
        }
        _topDonors[2] = address(0);
    }

    function _shiftRightFrom(uint256 index) private {
        for (uint256 i = 2; i > index; --i) {
            _topDonors[i] = _topDonors[i - 1];
        }
    }
}
