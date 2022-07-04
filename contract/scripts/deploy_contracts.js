async function main() {
    // ---------------------Deploy contracts---------------------
    // ---------------------------------------------------------
    const [admin, seller, buyer] = await ethers.getSigners();
  
    console.log("Deploying contracts with the account:", admin.address);

    const USDToken = await ethers.getContractFactory("USD", admin);
    const USD = await USDToken.deploy();

    const ERC721Token = await ethers.getContractFactory("NFT", admin);
    const NFT = await ERC721Token.deploy();

    const Market = await ethers.getContractFactory("Market", admin);
    const market = await Market.deploy(USD.address, NFT.address);
  
    console.log("USD address:", USD.address);
    console.log("NFT address:", NFT.address);
    console.log("Market address:", market.address);

    // ---------------------Mint tokens---------------------
    // -----------------------------------------------------
    await USD.connect(admin).mint(buyer.address, 1000);
    await USD.connect(buyer).approve(market.address, 1000);

    await NFT.connect(admin).mint(seller.address);
    await NFT.connect(seller).approve(market.address, 1);
  }
  
main()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
});