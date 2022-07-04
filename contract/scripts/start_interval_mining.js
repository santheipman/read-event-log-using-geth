const { ethers } = require('hardhat');

async function main() {
    const provider = ethers.provider;
    await provider.send("evm_setAutomine", [false]);
    await provider.send("evm_setIntervalMining", [5000]);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
    console.error(error);
    process.exit(1);
});