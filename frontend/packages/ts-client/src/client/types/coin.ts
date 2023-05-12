enum CoinType {
    BITCOIN,
    ETHEREUM,
    LITECOIN,
    DOGE,
    SONR,
    COSMOS,
    FILECOIN,
    HNS,
    TESTNET,
    SOLANA,
    XRP,
}

/**
 * The function returns an array of all available coin types.
 * @returns The function `allCoinTypes()` returns an array of `CoinType` values, which includes
 * `BITCOIN`, `ETHEREUM`, `LITECOIN`, `DOGE`, `SONR`, `COSMOS`, `FILECOIN`, `HNS`, `TESTNET`, `SOLANA`,
 * and `XRP`.
 */
export function allCoinTypes(): CoinType[] {
    return [
        CoinType.BITCOIN,
        CoinType.ETHEREUM,
        CoinType.LITECOIN,
        CoinType.DOGE,
        CoinType.SONR,
        CoinType.COSMOS,
        CoinType.FILECOIN,
        CoinType.HNS,
        CoinType.TESTNET,
        CoinType.SOLANA,
        CoinType.XRP,
    ];
}

/**
 * This TypeScript function returns the coin type based on the address prefix string provided.
 * @param {string} str - A string that represents a cryptocurrency address.
 * @returns The function `coinTypeFromAddrPrefix` returns a `CoinType` enum value. If the input string
 * contains the address prefix of a known coin type, that coin type is returned. Otherwise, it returns
 * `CoinType.TESTNET` as a default value.
 */
export function coinTypeFromAddrPrefix(str: string): CoinType {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (str.includes(getAddrPrefix(coin))) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}

/**
 * This function returns the coin type based on the BIP path number.
 * @param {number} i - The parameter `i` is a number representing a BIP path.
 * @returns a `CoinType` enum value. If the `getBipPath` function for a `CoinType` matches the input
 * parameter `i`, that `CoinType` is returned. If no match is found, the function returns
 * `CoinType.TESTNET`.
 */
export function coinTypeFromBipPath(i: number): CoinType {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getBipPath(coin) === i) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}

/**
 * This TypeScript function returns the CoinType based on the given DID method string.
 * @param {string} str - The `str` parameter is a string representing a DID method.
 * @returns a value of type `CoinType`. If the input string matches a DID method for a supported coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
export function coinTypeFromDidMethod(str: string): CoinType {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getDidMethod(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}

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
export function coinTypeFromName(str: string): CoinType {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getBlockchainName(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}

/**
 * This TypeScript function takes a string argument representing a cryptocurrency ticker and returns
 * the corresponding CoinType enum value, or CoinType.TESTNET if no match is found.
 * @param {string} str - The `str` parameter is a string representing the ticker symbol of a
 * cryptocurrency.
 * @returns a value of type `CoinType`. If the input string matches the ticker symbol of a known coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
export function coinTypeFromTicker(str: string): CoinType {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getTicker(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}

/**
 * The function returns a prefix string based on the given CoinType.
 * @param {CoinType} ct - CoinType - an enum representing different types of cryptocurrencies.
 * @returns A string representing the address prefix for a given CoinType. The specific prefix returned
 * depends on the input CoinType, with a default value of "test" if the input is not recognized.
 */
export function getAddrPrefix(ct: CoinType): string {
    switch (ct) {
        case CoinType.BITCOIN:
            return "bc";
        case CoinType.ETHEREUM:
            return "0x";
        case CoinType.LITECOIN:
            return "ltc";
        case CoinType.DOGE:
            return "doge";
        case CoinType.SONR:
            return "idx";
        case CoinType.COSMOS:
            return "cosmos";
        case CoinType.FILECOIN:
            return "f";
        case CoinType.HNS:
            return "hs";
        case CoinType.TESTNET:
            return "test";
        case CoinType.SOLANA:
            return "sol";
        case CoinType.XRP:
            return "xrp";
        default:
            return "test";
    }
}

/**
 * The function "getBipPath" takes a parameter "ct" of type "CoinType" and returns a number.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as Bitcoin, Ethereum, Litecoin, etc. The function
 * `getBipPath` likely returns a BIP (Bitcoin Improvement Proposal) path for a given cryptocurrency or
 * blockchain network.
 */
export function getBipPath(ct: CoinType): number {
    switch (ct) {
        case CoinType.BITCOIN:
            return 0;
        case CoinType.ETHEREUM:
            return 60;
        case CoinType.LITECOIN:
            return 2;
        case CoinType.DOGE:
            return 3;
        case CoinType.SONR:
            return 703;
        case CoinType.COSMOS:
            return 118;
        case CoinType.FILECOIN:
            return 461;
        case CoinType.HNS:
            return 5353;
        case CoinType.TESTNET:
            return 1;
        case CoinType.SOLANA:
        case CoinType.SOLANA:
            return 501;
        case CoinType.XRP:
            return 144;
        default:
            return 1;
    }
}

/**
 * The function returns the name of a blockchain based on the given coin type.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or blockchain networks. Without more context or information about the codebase,
 * it's difficult to say for sure.
 */
export function getBlockchainName(ct: CoinType): string {
    switch (ct) {
        case CoinType.BITCOIN:
            return "Bitcoin";
        case CoinType.ETHEREUM:
            return "Ethereum";
        case CoinType.LITECOIN:
            return "Litecoin";
        case CoinType.DOGE:
            return "Dogecoin";
        case CoinType.SONR:
            return "Sonr";
        case CoinType.COSMOS:
            return "Cosmos";
        case CoinType.FILECOIN:
            return "Filecoin";
        case CoinType.HNS:
            return "Handshake";
        case CoinType.SOLANA:
            return "Solana";
        case CoinType.XRP:
            return "Ripple";
        default:
            return "Testnet";
    }
}

/**
 * The function "getTicker" takes in a parameter "ct" of type "CoinType" and returns a string.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or tokens. The function `getTicker` probably takes in a `CoinType` parameter and
 * returns a string representing the ticker symbol of that cryptocurrency or token.
 */
export function getTicker(ct: CoinType): string {
    switch (ct) {
        case CoinType.BITCOIN:
            return "BTC";
        case CoinType.ETHEREUM:
            return "ETH";
        case CoinType.LITECOIN:
            return "LTC";
        case CoinType.DOGE:
            return "DOGE";
        case CoinType.SONR:
            return "SNR";
        case CoinType.COSMOS:
            return "ATOM";
        case CoinType.FILECOIN:
            return "FIL";
        case CoinType.HNS:
            return "HNS";
        case CoinType.SOLANA:
            return "SOL";
        case CoinType.XRP:
            return "XRP";
        default:
            return "TESTNET";
    }
}

/**
 * The function checks if a given coin type is Bitcoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isBitcoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Bitcoin or not.
 */
export function isBitcoin(c: CoinType): boolean {
    return c === CoinType.BITCOIN;
}

/**
 * The function checks if a given coin type is Cosmos.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `c` can take on.
 */
export function isCosmos(c: CoinType): boolean {
    return c === CoinType.COSMOS;
}

/**
 * The function checks if a given coin type is Ethereum.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `CoinType` can take on.
 */
export function isEthereum(c: CoinType): boolean {
    return c === CoinType.ETHEREUM;
}

/**
 * The function checks if a given coin type is Filecoin.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * that represents a cryptocurrency. The function `isFilecoin` takes in this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Filecoin or not.
 */
export function isFilecoin(c: CoinType): boolean {
    return c === CoinType.FILECOIN;
}

/**
 * The function checks if a given coin type is a handshake coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isHandshake" likely
 * checks if the given coin type is Handshake, and returns a boolean value of true if it is, and false
 * otherwise
 */
export function isHandshake(c: CoinType): boolean {
    return c === CoinType.HNS;
}

/**
 * The function checks if a given coin type is Litecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function is checking whether the input
 * coin type is Litecoin or not, and it returns a boolean value accordingly.
 */
export function isLitecoin(c: CoinType): boolean {
    return c === CoinType.LITECOIN;
}

/**
 * The function checks if a given coin type is Solana and returns a boolean value.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a string
 * type that represents a cryptocurrency. The function `isSolana` takes this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Solana or not.
 */
export function isSolana(c: CoinType): boolean {
    return c === CoinType.SOLANA;
}

/**
 * The function checks if a given coin type is a Ripple coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing a cryptocurrency.
 */
export function isRipple(c: CoinType): boolean {
    return c === CoinType.XRP;
}

/**
 * The function checks if a given coin type is a testnet.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency. It
 * could be used to determine whether a given cryptocurrency is a testnet or not.
 */
export function isTestnet(c: CoinType): boolean {
    return c === CoinType.TESTNET;
}

/**
 * The function checks if a given coin type is Dogecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isDogecoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Dogecoin or
 * not.
 */
export function isDogecoin(c: CoinType): boolean {
    return c === CoinType.DOGE;
}

/**
 * The function checks if a given coin type is Sonr or not and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isSonr" likely checks
 * if the given coin is a specific type of coin, possibly called "SONR".
 */
export function isSonr(c: CoinType): boolean {
    return c === CoinType.SONR;
}

/**
 * The function returns a string representing the DID method based on the given coin type.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as "Bitcoin", "Ethereum", "Litecoin", "Ripple",
 * etc.
 */
export function getDidMethod(c: CoinType): string {
    if (isBitcoin(c)) {
        return "btcr";
    }
    if (isEthereum(c)) {
        return "ethr";
    }
    if (isSonr(c)) {
        return "sonr";
    }
    return getTicker(c).toLowerCase();
}
