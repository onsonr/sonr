"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getDidMethod = exports.isSonr = exports.isDogecoin = exports.isTestnet = exports.isRipple = exports.isSolana = exports.isLitecoin = exports.isHandshake = exports.isFilecoin = exports.isEthereum = exports.isCosmos = exports.isBitcoin = exports.getTicker = exports.getBlockchainName = exports.getBipPath = exports.getAddrPrefix = exports.coinTypeFromTicker = exports.coinTypeFromName = exports.coinTypeFromDidMethod = exports.coinTypeFromBipPath = exports.coinTypeFromAddrPrefix = exports.allCoinTypes = void 0;
var CoinType;
(function (CoinType) {
    CoinType[CoinType["BITCOIN"] = 0] = "BITCOIN";
    CoinType[CoinType["ETHEREUM"] = 1] = "ETHEREUM";
    CoinType[CoinType["LITECOIN"] = 2] = "LITECOIN";
    CoinType[CoinType["DOGE"] = 3] = "DOGE";
    CoinType[CoinType["SONR"] = 4] = "SONR";
    CoinType[CoinType["COSMOS"] = 5] = "COSMOS";
    CoinType[CoinType["FILECOIN"] = 6] = "FILECOIN";
    CoinType[CoinType["HNS"] = 7] = "HNS";
    CoinType[CoinType["TESTNET"] = 8] = "TESTNET";
    CoinType[CoinType["SOLANA"] = 9] = "SOLANA";
    CoinType[CoinType["XRP"] = 10] = "XRP";
})(CoinType || (CoinType = {}));
/**
 * The function returns an array of all available coin types.
 * @returns The function `allCoinTypes()` returns an array of `CoinType` values, which includes
 * `BITCOIN`, `ETHEREUM`, `LITECOIN`, `DOGE`, `SONR`, `COSMOS`, `FILECOIN`, `HNS`, `TESTNET`, `SOLANA`,
 * and `XRP`.
 */
function allCoinTypes() {
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
exports.allCoinTypes = allCoinTypes;
/**
 * This TypeScript function returns the coin type based on the address prefix string provided.
 * @param {string} str - A string that represents a cryptocurrency address.
 * @returns The function `coinTypeFromAddrPrefix` returns a `CoinType` enum value. If the input string
 * contains the address prefix of a known coin type, that coin type is returned. Otherwise, it returns
 * `CoinType.TESTNET` as a default value.
 */
function coinTypeFromAddrPrefix(str) {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (str.includes(getAddrPrefix(coin))) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}
exports.coinTypeFromAddrPrefix = coinTypeFromAddrPrefix;
/**
 * This function returns the coin type based on the BIP path number.
 * @param {number} i - The parameter `i` is a number representing a BIP path.
 * @returns a `CoinType` enum value. If the `getBipPath` function for a `CoinType` matches the input
 * parameter `i`, that `CoinType` is returned. If no match is found, the function returns
 * `CoinType.TESTNET`.
 */
function coinTypeFromBipPath(i) {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getBipPath(coin) === i) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}
exports.coinTypeFromBipPath = coinTypeFromBipPath;
/**
 * This TypeScript function returns the CoinType based on the given DID method string.
 * @param {string} str - The `str` parameter is a string representing a DID method.
 * @returns a value of type `CoinType`. If the input string matches a DID method for a supported coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
function coinTypeFromDidMethod(str) {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getDidMethod(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}
exports.coinTypeFromDidMethod = coinTypeFromDidMethod;
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
function coinTypeFromName(str) {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getBlockchainName(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}
exports.coinTypeFromName = coinTypeFromName;
/**
 * This TypeScript function takes a string argument representing a cryptocurrency ticker and returns
 * the corresponding CoinType enum value, or CoinType.TESTNET if no match is found.
 * @param {string} str - The `str` parameter is a string representing the ticker symbol of a
 * cryptocurrency.
 * @returns a value of type `CoinType`. If the input string matches the ticker symbol of a known coin
 * type, that coin type is returned. Otherwise, the function returns `CoinType.TESTNET`.
 */
function coinTypeFromTicker(str) {
    const coins = allCoinTypes();
    for (const coin of coins) {
        if (getTicker(coin).toLowerCase() === str.toLowerCase()) {
            return coin;
        }
    }
    return CoinType.TESTNET;
}
exports.coinTypeFromTicker = coinTypeFromTicker;
/**
 * The function returns a prefix string based on the given CoinType.
 * @param {CoinType} ct - CoinType - an enum representing different types of cryptocurrencies.
 * @returns A string representing the address prefix for a given CoinType. The specific prefix returned
 * depends on the input CoinType, with a default value of "test" if the input is not recognized.
 */
function getAddrPrefix(ct) {
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
exports.getAddrPrefix = getAddrPrefix;
/**
 * The function "getBipPath" takes a parameter "ct" of type "CoinType" and returns a number.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as Bitcoin, Ethereum, Litecoin, etc. The function
 * `getBipPath` likely returns a BIP (Bitcoin Improvement Proposal) path for a given cryptocurrency or
 * blockchain network.
 */
function getBipPath(ct) {
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
exports.getBipPath = getBipPath;
/**
 * The function returns the name of a blockchain based on the given coin type.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or blockchain networks. Without more context or information about the codebase,
 * it's difficult to say for sure.
 */
function getBlockchainName(ct) {
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
exports.getBlockchainName = getBlockchainName;
/**
 * The function "getTicker" takes in a parameter "ct" of type "CoinType" and returns a string.
 * @param {CoinType} ct - CoinType is likely an enum or a type that represents different types of
 * cryptocurrencies or tokens. The function `getTicker` probably takes in a `CoinType` parameter and
 * returns a string representing the ticker symbol of that cryptocurrency or token.
 */
function getTicker(ct) {
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
exports.getTicker = getTicker;
/**
 * The function checks if a given coin type is Bitcoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isBitcoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Bitcoin or not.
 */
function isBitcoin(c) {
    return c === CoinType.BITCOIN;
}
exports.isBitcoin = isBitcoin;
/**
 * The function checks if a given coin type is Cosmos.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `c` can take on.
 */
function isCosmos(c) {
    return c === CoinType.COSMOS;
}
exports.isCosmos = isCosmos;
/**
 * The function checks if a given coin type is Ethereum.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * alias defined elsewhere in the codebase. Without more information, it's difficult to say exactly
 * what values `CoinType` can take on.
 */
function isEthereum(c) {
    return c === CoinType.ETHEREUM;
}
exports.isEthereum = isEthereum;
/**
 * The function checks if a given coin type is Filecoin.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a type
 * that represents a cryptocurrency. The function `isFilecoin` takes in this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Filecoin or not.
 */
function isFilecoin(c) {
    return c === CoinType.FILECOIN;
}
exports.isFilecoin = isFilecoin;
/**
 * The function checks if a given coin type is a handshake coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isHandshake" likely
 * checks if the given coin type is Handshake, and returns a boolean value of true if it is, and false
 * otherwise
 */
function isHandshake(c) {
    return c === CoinType.HNS;
}
exports.isHandshake = isHandshake;
/**
 * The function checks if a given coin type is Litecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function is checking whether the input
 * coin type is Litecoin or not, and it returns a boolean value accordingly.
 */
function isLitecoin(c) {
    return c === CoinType.LITECOIN;
}
exports.isLitecoin = isLitecoin;
/**
 * The function checks if a given coin type is Solana and returns a boolean value.
 * @param {CoinType} c - The parameter `c` is of type `CoinType`, which is likely an enum or a string
 * type that represents a cryptocurrency. The function `isSolana` takes this parameter and returns a
 * boolean value indicating whether the given cryptocurrency is Solana or not.
 */
function isSolana(c) {
    return c === CoinType.SOLANA;
}
exports.isSolana = isSolana;
/**
 * The function checks if a given coin type is a Ripple coin.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing a cryptocurrency.
 */
function isRipple(c) {
    return c === CoinType.XRP;
}
exports.isRipple = isRipple;
/**
 * The function checks if a given coin type is a testnet.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency. It
 * could be used to determine whether a given cryptocurrency is a testnet or not.
 */
function isTestnet(c) {
    return c === CoinType.TESTNET;
}
exports.isTestnet = isTestnet;
/**
 * The function checks if a given coin type is Dogecoin and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type that represents different types of cryptocurrencies. The function "isDogecoin" takes in this
 * parameter and returns a boolean value indicating whether the given cryptocurrency is Dogecoin or
 * not.
 */
function isDogecoin(c) {
    return c === CoinType.DOGE;
}
exports.isDogecoin = isDogecoin;
/**
 * The function checks if a given coin type is Sonr or not and returns a boolean value.
 * @param {CoinType} c - The parameter "c" is of type CoinType, which is likely an enum or a custom
 * type representing different types of cryptocurrencies or coins. The function "isSonr" likely checks
 * if the given coin is a specific type of coin, possibly called "SONR".
 */
function isSonr(c) {
    return c === CoinType.SONR;
}
exports.isSonr = isSonr;
/**
 * The function returns a string representing the DID method based on the given coin type.
 * @param {CoinType} c - CoinType is likely an enum or a type that represents a cryptocurrency or
 * blockchain network. It could include values such as "Bitcoin", "Ethereum", "Litecoin", "Ripple",
 * etc.
 */
function getDidMethod(c) {
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
exports.getDidMethod = getDidMethod;
