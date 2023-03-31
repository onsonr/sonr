import {mnemonic} from "./wallet.js";
import {NftTestEngine} from "./nft_test.js";
import {main_addr} from "./constants.js";

let NFT721 = new NftTestEngine(
    mnemonic,
    main_addr,
    "TEST",
    "TEST"
)

await NFT721.setup()
console.log(NFT721.contract_address)
// console.log("Re-resending Same TX to crash and regenerate the error ( tx already exists in cache )")
// await NFT721.setup()
