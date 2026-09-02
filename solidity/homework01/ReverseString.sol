// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title ReverseString
/// @notice 作业1 第2题：反转字符串。输入 "abcde"，输出 "edcba"。
/// @dev 按字节反转，适合 Remix 里测 ASCII；多字节 UTF-8 会按字节打乱。
contract ReverseString {
    function reverse(string memory input) external pure returns (string memory) {
        bytes memory source = bytes(input);
        uint256 length = source.length;
        bytes memory output = new bytes(length);

        for (uint256 i = 0; i < length; ++i) {
            output[i] = source[length - 1 - i];
        }

        return string(output);
    }
}
