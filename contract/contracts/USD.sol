//SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract USD is ERC20, Ownable {
    constructor() ERC20("USDollar", "USD"){}
    
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}