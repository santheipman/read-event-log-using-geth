const { ethers } = require('hardhat');
const { expect } = require('chai');

describe("Market", function () {
    let admin;
    let seller;
    let buyer;

    let USD;
    let NFT;
    let market;

    beforeEach(async function () {
        [admin, seller, buyer] = await ethers.getSigners();

        const USDToken = await ethers.getContractFactory("USD", admin);
        USD = await USDToken.deploy();

        const ERC721Token = await ethers.getContractFactory("NFT", admin);
        NFT = await ERC721Token.deploy();

        const Market = await ethers.getContractFactory("Market", admin);
        market = await Market.deploy(USD.address, NFT.address);

        await USD.connect(admin).mint(buyer.address, 1000);
        await USD.connect(buyer).approve(market.address, 1000);

        await NFT.connect(admin).mint(seller.address);
        await NFT.connect(seller).approve(market.address, 1);
    });

    it("Emit event when user create sell order", async function () {
        await expect(market.connect(seller).createSellOrder(1, 150))
            .to.emit(market, 'CreateSellOrder')
            .withArgs(1, seller.address, 1, 150);
    });

    it("Buyer use seller's signature & order info to call buy function", async function () {
        await expect(market.connect(seller).createSellOrder(1, 150))
            .to.emit(market, 'CreateSellOrder')
            .withArgs(1, seller.address, 1, 150);

        const signature = await getSignatureUsingEthersJS(seller, 1, 1, 150);

        await market.connect(buyer).buy(1, 1, 150, signature);

        expect(await NFT.ownerOf(1)).to.equal(buyer.address);
        expect(await USD.balanceOf(seller.address)).to.equal(150);
    });

    it("Revert when buyer use his signature to call buy function", async function () {
        await expect(market.connect(seller).createSellOrder(1, 150))
            .to.emit(market, 'CreateSellOrder')
            .withArgs(1, seller.address, 1, 150);

        const signature = await getSignatureUsingEthersJS(buyer, 1, 1, 150);

        await expect(market.connect(buyer).buy(1, 1, 150, signature))
            .to.be.revertedWith('incorrect signature');
    });

    it("Revert when buyer use seller's signature and MODIFIED order info to call buy function", async function () {
        await expect(market.connect(seller).createSellOrder(1, 150))
            .to.emit(market, 'CreateSellOrder')
            .withArgs(1, seller.address, 1, 150);

        const signature = await getSignatureUsingEthersJS(seller, 1, 1, 150);

        await expect(market.connect(buyer).buy(1, 1, 10, signature))
            .to.be.revertedWith('incorrect signature');
    });

    it("Revert when user buy an order more than once", async function () {
        await expect(market.connect(seller).createSellOrder(1, 150))
            .to.emit(market, 'CreateSellOrder')
            .withArgs(1, seller.address, 1, 150);

        const signature = await getSignatureUsingEthersJS(seller, 1, 1, 150);

        await market.connect(buyer).buy(1, 1, 150, signature);

        // https://docs.ethers.io/v5/single-page/#/v5/migration/web3/-%23-migration-from-web3-js--contracts--overloaded-functions
        await NFT.connect(buyer)["safeTransferFrom(address,address,uint256)"](buyer.address, seller.address, 1);

        await expect(market.connect(buyer).buy(1, 1, 150, signature))
            .to.be.revertedWith('Order is not exist');
    });

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
});