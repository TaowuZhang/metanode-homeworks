// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title IntegerToRoman
/// @notice 作业1 第3题：整数转罗马数字。对应 LeetCode 12。
/// @dev 合法范围 1..3999。题目原文把两道罗马题的链接写反了，这里按「整数 → 罗马」实现。
contract IntegerToRoman {
    error OutOfRange(uint256 number);

    function intToRoman(uint256 number) external pure returns (string memory) {
        if (number == 0 || number > 3999) {
            revert OutOfRange(number);
        }

        uint16[13] memory values = [uint16(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

        bytes memory result;
        uint256 remaining = number;

        for (uint256 i = 0; i < values.length; ++i) {
            while (remaining >= values[i]) {
                remaining -= values[i];
                result = bytes.concat(result, bytes(symbols[i]));
            }
        }

        return string(result);
    }
}
