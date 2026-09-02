// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title MergeSortedArray
/// @notice 作业1 第5题：合并两个有序数组，得到一个新的有序数组。
contract MergeSortedArray {
    function mergeSorted(uint256[] memory left, uint256[] memory right)
        external
        pure
        returns (uint256[] memory merged)
    {
        merged = new uint256[](left.length + right.length);
        uint256 i;
        uint256 j;
        uint256 k;

        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                merged[k++] = left[i++];
            } else {
                merged[k++] = right[j++];
            }
        }

        while (i < left.length) {
            merged[k++] = left[i++];
        }
        while (j < right.length) {
            merged[k++] = right[j++];
        }
    }
}
