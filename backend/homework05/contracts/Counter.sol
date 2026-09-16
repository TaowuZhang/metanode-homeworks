// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Counter {
    uint256 public number;

    event Incremented(uint256 newValue);

    function increment() external {
        number += 1;
        emit Incremented(number);
    }
}
