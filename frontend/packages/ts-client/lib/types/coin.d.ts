declare enum CoinType {
    BITCOIN = 0,
    ETHEREUM = 1,
    LITECOIN = 2,
    DOGE = 3,
    SONR = 4,
    COSMOS = 5,
    FILECOIN = 6,
    HNS = 7,
    TESTNET = 8,
    SOLANA = 9,
    XRP = 10
}
/**
 * The function returns an array of all available coin types.
 * @returns The function `allCoinTypes()` returns an array of `CoinType` values, which includes
 * `BITCOIN`, `ETHEREUM`, `LITECOIN`, `DOGE`, `SONR`, `COSMOS`, `FILECOIN`, `HNS`, `TESTNET`, `SOLANA`,
 * and `XRP`.
 */
export declare function allCoinTypes(): CoinType[];
/**
 * This TypeScript function returns the coin type based on the address prefix string provided.
 * @param {string} str - A string that represents a cryptocurrency address.
 * @returns The function `coinTypeFromAddrPrefix` returns a `CoinType` enum value. If the input string
 * contains the address prefix of a known coin type, that coin type is returned. Otherwise, it returns
 * `CoinType.TESTNET` as a default value.
 */
export declare function coinTypeFromAddrPrefix(str: string): CoinType;
/**
 * This function returns the coin type based on the BIP path number.
 * @param {number} i - The parameter `i` is a number representing a BIP path.
 * @returns a `CoinType` enum value. If the `getBipPath` function for a `CoinType` matches the input
 * parameter `i`, that `CoinType` is returned. If no match is found, the function returns
 * `CoinType.TESTNET`.
 */
export declare function coinTypeFromBipPath(i: number): CoinType;
/**
 * This TypeScript function returns the CoinType based on the given DID method string.
 * @param {string} str - The `str` parameter is a string representing a DID method.
 * @returns a value of type `CoinType`. If the input string matches a DID method for a supported coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
export declare function coinTypeFromDidMethod(str: string): CoinType;
/**
 * This TypeScript function returns the CoinType based on the input string representing the blockchain
 * name.
 * @param {string} str - The `str` parameter is a string that represents the name of a blockchain. This
 * function searches through a list of all available `CoinType` values and returns the `CoinType` that
 * matches the given blockchain name. If no match is found, it returns `CoinType.TESTNET`.
 * @returns a `CoinType` value. If the input string matches the name of a blockchain in the
 * `allCoinTypes()` list, the corresponding `CoinType` is returned. If there is no match, the function
 * returns `CoinType.TESTNET`.
 */
export declare function coinTypeFromName(str: string): CoinType;
/**
 * This TypeScript function takes a string argument representing a cryptocurrency ticker and returns
 * the corresponding CoinType enum value, or CoinType.TESTNET if no match is found.
 * @param {string} str - The `str` parameter is a string representing the ticker symbol of a
 * cryptocurrency.
 * @returns a value of type `CoinType`. If the input string matches the ticker symbol of a known coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
export declare function coinTypeFromTicker(str: string): CoinType;
/**
 * The function returns a prefix string based on the given CoinType.
 * @param {CoinType} ct - CoinType - an enum representing different types of cryptocurrencies.
 * @returns A string representing the address prefix for a given CoinType. The specific prefix returned
 * depends on the input CoinType, with a default value of "test" if the input is not recognized.
 */
export declare function getAddrPrefix(ct: CoinType): string;
/**
 * The function "getBipPath" takes a parameter "ct" of type "CoinType" and returns a number.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as Bitcoin, Ethereum, Litecoin, etc. The function
 * `getBipPath` likely returns a BIP (Bitcoin Improvement Proposal) path for a given cryptocurrency or
 * blockchain network.
 */
export declare function getBipPath(ct: CoinType): number;
/**
 * The function returns the name of a blockchain based on the given coin type.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or blockchain networks. Without more context or information about the codebase,
 * it's difficult to say for sure.
 */
export declare function getBlockchainName(ct: CoinType): string;
/**
 * The function "getTicker" takes in a parameter "ct" of type "CoinType" and returns a string.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or tokens. The function `getTicker` probably takes in a `CoinType` parameter and
 * returns a string representing the ticker symbol of that cryptocurrency or token.
 */
export declare function getTicker(ct: CoinType): string;
/**
 * The function checks if a given coin type is Bitcoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isBitcoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Bitcoin or not.
 */
export declare function isBitcoin(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Cosmos.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `c` can take on.
 */
export declare function isCosmos(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Ethereum.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `CoinType` can take on.
 */
export declare function isEthereum(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Filecoin.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * that represents a cryptocurrency. The function `isFilecoin` takes in this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Filecoin or not.
 */
export declare function isFilecoin(c: CoinType): boolean;
/**
 * The function checks if a given coin type is a handshake coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isHandshake" likely
 * checks if the given coin type is Handshake, and returns a boolean value of true if it is, and false
 * otherwise
 */
export declare function isHandshake(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Litecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function is checking whether the input
 * coin type is Litecoin or not, and it returns a boolean value accordingly.
 */
export declare function isLitecoin(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Solana and returns a boolean value.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a string
 * type that represents a cryptocurrency. The function `isSolana` takes this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Solana or not.
 */
export declare function isSolana(c: CoinType): boolean;
/**
 * The function checks if a given coin type is a Ripple coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing a cryptocurrency.
 */
export declare function isRipple(c: CoinType): boolean;
/**
 * The function checks if a given coin type is a testnet.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency. It
 * could be used to determine whether a given cryptocurrency is a testnet or not.
 */
export declare function isTestnet(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Dogecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isDogecoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Dogecoin or
 * not.
 */
export declare function isDogecoin(c: CoinType): boolean;
/**
 * The function checks if a given coin type is Sonr or not and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isSonr" likely checks
 * if the given coin is a specific type of coin, possibly called "SONR".
 */
export declare function isSonr(c: CoinType): boolean;
/**
 * The function returns a string representing the DID method based on the given coin type.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as "Bitcoin", "Ethereum", "Litecoin", "Ripple",
 * etc.
 */
export declare function getDidMethod(c: CoinType): string;
export {};
