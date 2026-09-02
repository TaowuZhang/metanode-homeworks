// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title BinarySearch
/// @notice 作业1 第6题：在有序数组中查找目标值，找到返回下标，找不到返回 -1。
contract BinarySearch {
    function binarySearch(uint256[] memory data, uint256 target) external pure returns (int256) {
        if (data.length == 0) {
            return -1;
        }

        uint256 left = 0;
        uint256 right = data.length - 1;

        while (left <= right) {
            uint256 mid = left + (right - left) / 2;
            if (data[mid] == target) {
                return int256(mid);
            }
            if (data[mid] < target) {
                left = mid + 1;
            } else {
                if (mid == 0) {
                    return -1;
                }
                right = mid - 1;
            }
        }

        return -1;
    }
}
