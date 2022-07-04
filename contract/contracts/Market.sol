//SPDX-License-Identifier: MIT
pragma solidity 0.8.1;

import "./interfaces/INFT.sol";
import "./interfaces/IUSD.sol";

contract Market {
    event CreateSellOrder(uint256 indexed _orderId, address indexed _seller, uint256 _tokenId, uint256 _price);
    event CancelSell(uint256 indexed _orderId);
    event Buy(uint256 indexed _orderId, address indexed buyer);

    address public USDAddress;
    address public NFTAddress;

    uint256 public orderCounter;
    mapping(uint256 => address) ownerOfOrder;

    constructor(address _USDAddress, address _NFTAddress) {
        USDAddress = _USDAddress;
        NFTAddress = _NFTAddress;
    }

    function createSellOrder(uint256 tokenId, uint256 price) external {
        require(INFT(NFTAddress).ownerOf(tokenId) == msg.sender, "Seller isn't owner of NFT");

        orderCounter += 1;
        ownerOfOrder[orderCounter] = msg.sender;

        emit CreateSellOrder(orderCounter, msg.sender, tokenId, price);
    }

    function cancelSell(uint256 orderId) external {
        require(ownerOfOrder[orderId] == msg.sender, "Caller isn't order's owner");

        delete ownerOfOrder[orderId];

        emit CancelSell(orderId);
    }

    function buy(uint256 orderId, uint256 tokenId, uint256 price, bytes memory signature) external {
        // verify params

        address seller = ownerOfOrder[orderId];

        require(seller != address(0), 'Order is not exist');

        verify(seller, orderId, tokenId, price, signature);

        require(IUSD(USDAddress).balanceOf(msg.sender) >= price, "Buyer doesn't have enough money");

        // transfer
        IUSD(USDAddress).transferFrom(msg.sender, seller, price);
        INFT(NFTAddress).safeTransferFrom(seller, msg.sender, tokenId);

        delete ownerOfOrder[orderId];

        emit Buy(orderId, msg.sender);
    }

    // ---------------------Signature verification---------------------
    // ----------------------------------------------------------------
    function getMessageHash(uint256 orderId, uint256 tokenId, uint256 price) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                orderId,
                tokenId,
                price
            )
        );
    }

    function getEthSignedMessageHash(bytes32 _messageHash) internal pure returns (bytes32){
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", _messageHash));
    }

    function verify(address signer, uint256 orderId, uint256 tokenId, uint256 price, bytes memory signature) internal pure {
        bytes32 messageHash = getMessageHash(orderId, tokenId, price);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);

        require(recoverSigner(ethSignedMessageHash, signature) == signer, 'incorrect signature');
    }

    function recoverSigner(bytes32 _ethSignedMessageHash, bytes memory _signature) internal pure returns (address) {
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);

        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function splitSignature(bytes memory sig) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
        require(sig.length == 65, "invalid signature length");

        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
    }
}