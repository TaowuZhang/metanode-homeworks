// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title RomanToInteger
/// @notice 作业1 第4题：罗马数字转整数。对应 LeetCode 13。
/// @dev 题目原文把两道罗马题的链接写反了，这里按「罗马 → 整数」实现。
contract RomanToInteger {
    error InvalidRoman(bytes1 character);

    function romanToInt(string memory roman) external pure returns (uint256) {
        bytes memory chars = bytes(roman);
        uint256 total = 0;
        uint256 previous = 0;

        for (uint256 i = 0; i < chars.length; ++i) {
            uint256 current = _valueOf(chars[i]);
            if (current > previous && previous != 0) {
                total += current - 2 * previous;
            } else {
                total += current;
            }
            previous = current;
        }

        return total;
    }

    function _valueOf(bytes1 character) private pure returns (uint256) {
        if (character == "I") return 1;
        if (character == "V") return 5;
        if (character == "X") return 10;
        if (character == "L") return 50;
        if (character == "C") return 100;
        if (character == "D") return 500;
        if (character == "M") return 1000;
        revert InvalidRoman(character);
    }
}
