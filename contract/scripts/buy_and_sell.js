async function main() {
    // ---------------------getContracts---------------------
    // ------------------------------------------------------
    const USDAddress = '0x5FbDB2315678afecb367f032d93F642f64180aa3'
    const NFTAddress = '0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512'
    const marketAddress = '0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0'

    const [admin, seller, buyer] = await ethers.getSigners();

    const USD = await ethers.getContractAt("USD", USDAddress)
    const NFT = await ethers.getContractAt("NFT", NFTAddress)
    const market = await ethers.getContractAt("Market", marketAddress)

    // ---------------------Buy, sign, sell---------------------
    // ---------------------------------------------------------
    await market.connect(seller).createSellOrder(1, 150);
    const signature = getSignatureUsingEthersJS(seller, 1, 1, 150);
    await market.connect(buyer).buy(1, 1, 150, signature);
}

async function getSignatureUsingEthersJS(signer, orderId, tokenId, price) {
    // hash the message
    const messageHash = ethers.utils.solidityKeccak256(
        ['uint256', 'uint256', 'uint256'],
        [orderId.toString(), tokenId.toString(), price.toString()]
    );

    // get signature produced by etherjs library
    const signature = await signer.signMessage(ethers.utils.arrayify(messageHash));

    return signature;
}
  
main()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
});