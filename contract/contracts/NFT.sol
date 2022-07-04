//SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721, Ownable {
    uint256 public counter;

    constructor() ERC721("Non-fungible Token", "NFT") {
        counter = 0;
    }

    function mint(address to) external onlyOwner {
        counter += 1;
        _safeMint(to, counter);
    }
}