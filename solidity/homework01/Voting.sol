// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title Voting
/// @notice 作业1 第1题：用 mapping 记录候选人得票，支持投票、查询和重置。
/// @dev 候选人用 uint256 编号。resetVotes 用 epoch 做 O(1) 清零，不必遍历 mapping。
contract Voting {
    mapping(uint256 => uint256) private _votes;
    mapping(uint256 => uint256) private _voteEpoch;
    uint256 public currentEpoch;

    event Voted(address indexed voter, uint256 indexed candidate, uint256 newTotal);
    event VotesReset(address indexed operator, uint256 newEpoch);

    function vote(uint256 candidate) external {
        if (_voteEpoch[candidate] != currentEpoch) {
            _voteEpoch[candidate] = currentEpoch;
            _votes[candidate] = 0;
        }
        _votes[candidate] += 1;
        emit Voted(msg.sender, candidate, _votes[candidate]);
    }

    function getVotes(uint256 candidate) external view returns (uint256) {
        if (_voteEpoch[candidate] != currentEpoch) {
            return 0;
        }
        return _votes[candidate];
    }

    function resetVotes() external {
        currentEpoch += 1;
        emit VotesReset(msg.sender, currentEpoch);
    }
}
