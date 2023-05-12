"use strict";
/*
This file is a work taken from the following location: https://gist.github.com/enepomnyaschih/72c423f727d395eeaa09697058238727

It has been modified using the below algorithms to make it work for the standard web encoding.

MIT License

Copyright (c) 2020 Egor Nepomnyaschih

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.publicKeyCredentialAssertionToJson = exports.publicKeyCredentialAttestionToJson = exports.getAxios = exports.baseUrl = exports.arrayBufferDecode = exports.arrayBufferEncode = exports.getBytesFromBase64 = exports.getBase64WebEncodingFromBytes = exports.getBase64FromBytes = void 0;
const axios_1 = __importDefault(require("axios"));
/*
// This constant can also be computed with the following algorithm:
const base64Chars = [],
    A = "A".charCodeAt(0),
    a = "a".charCodeAt(0),
    n = "0".charCodeAt(0);
for (let i = 0; i < 26; ++i) {
    base64Chars.push(String.fromCharCode(A + i));
}
for (let i = 0; i < 26; ++i) {
    base64Chars.push(String.fromCharCode(a + i));
}
for (let i = 0; i < 10; ++i) {
    base64Chars.push(String.fromCharCode(n + i));
}
base64Chars.push("+");
base64Chars.push("/");
*/
const base64Chars = [
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z",
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
    "g",
    "h",
    "i",
    "j",
    "k",
    "l",
    "m",
    "n",
    "o",
    "p",
    "q",
    "r",
    "s",
    "t",
    "u",
    "v",
    "w",
    "x",
    "y",
    "z",
    "0",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "+",
    "/",
];
/*
// This constant can also be computed with the following algorithm:
const l = 256, base64codes = new Uint8Array(l);
for (let i = 0; i < l; ++i) {
    base64codes[i] = 255; // invalid character
}
base64Chars.forEach((char, index) => {
    base64codes[char.charCodeAt(0)] = index;
});
base64codes["=".charCodeAt(0)] = 0; // ignored anyway, so we just need to prevent an error
*/
const base64Codes = [
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 62, 255, 255,
    255, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 255, 255, 255, 0, 255, 255,
    255, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
    21, 22, 23, 24, 25, 255, 255, 255, 255, 255, 255, 26, 27, 28, 29, 30, 31, 32,
    33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
];
function getBase64Code(charCode) {
    if (charCode >= base64Codes.length) {
        throw new Error("Unable to parse base64 string.");
    }
    const code = base64Codes[charCode];
    if (code === 255) {
        throw new Error("Unable to parse base64 string.");
    }
    return code;
}
function getBase64FromBytes(bytes) {
    let result = "", i, l = bytes.length;
    for (i = 2; i < l; i += 3) {
        result += base64Chars[bytes[i - 2] >> 2];
        result += base64Chars[((bytes[i - 2] & 0x03) << 4) | (bytes[i - 1] >> 4)];
        result += base64Chars[((bytes[i - 1] & 0x0f) << 2) | (bytes[i] >> 6)];
        result += base64Chars[bytes[i] & 0x3f];
    }
    if (i === l + 1) {
        // 1 octet yet to write
        result += base64Chars[bytes[i - 2] >> 2];
        result += base64Chars[(bytes[i - 2] & 0x03) << 4];
        result += "==";
    }
    if (i === l) {
        // 2 octets yet to write
        result += base64Chars[bytes[i - 2] >> 2];
        result += base64Chars[((bytes[i - 2] & 0x03) << 4) | (bytes[i - 1] >> 4)];
        result += base64Chars[(bytes[i - 1] & 0x0f) << 2];
        result += "=";
    }
    return result;
}
exports.getBase64FromBytes = getBase64FromBytes;
function getBase64WebEncodingFromBytes(bytes) {
    return getBase64FromBytes(bytes)
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, "");
}
exports.getBase64WebEncodingFromBytes = getBase64WebEncodingFromBytes;
function getBytesFromBase64(str) {
    if (str.length % 4 !== 0) {
        throw new Error("Unable to parse base64 string.");
    }
    const index = str.indexOf("=");
    if (index !== -1 && index < str.length - 2) {
        throw new Error("Unable to parse base64 string.");
    }
    let missingOctets = str.endsWith("==") ? 2 : str.endsWith("=") ? 1 : 0, n = str.length, result = new Uint8Array(3 * (n / 4)), buffer;
    for (let i = 0, j = 0; i < n; i += 4, j += 3) {
        buffer =
            (getBase64Code(str.charCodeAt(i)) << 18) |
                (getBase64Code(str.charCodeAt(i + 1)) << 12) |
                (getBase64Code(str.charCodeAt(i + 2)) << 6) |
                getBase64Code(str.charCodeAt(i + 3));
        result[j] = buffer >> 16;
        result[j + 1] = (buffer >> 8) & 0xff;
        result[j + 2] = buffer & 0xff;
    }
    return result.subarray(0, result.length - missingOctets);
}
exports.getBytesFromBase64 = getBytesFromBase64;
function arrayBufferEncode(value) {
    return getBase64WebEncodingFromBytes(new Uint8Array(value));
}
exports.arrayBufferEncode = arrayBufferEncode;
function arrayBufferDecode(value) {
    if (value.length % 4 !== 0) {
        return Buffer.from(value, "base64");
    }
    return getBytesFromBase64(value);
}
exports.arrayBufferDecode = arrayBufferDecode;
function baseUrl() {
    if (process.env.NODE_ENV === "development") {
        return "http://0.0.0.0:8080";
    }
    return "https://highway.build";
}
exports.baseUrl = baseUrl;
function getAxios(authenticated) {
    if (authenticated) {
        return axios_1.default.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${localStorage.getItem('jwt')}` },
        });
    }
    return axios_1.default.create({
        baseURL: baseUrl(),
    });
}
exports.getAxios = getAxios;
function publicKeyCredentialAttestionToJson(credential) {
    let credResp = credential.response;
    let attestationObject = credResp.attestationObject;
    let clientDataJSON = credResp.clientDataJSON;
    let rawId = credential.rawId;
    let credRespString = JSON.stringify({
        id: credential.id,
        type: credential.type,
        rawId: arrayBufferEncode(rawId),
        clientExtensionResults: credential.getClientExtensionResults(),
        response: {
            attestationObject: arrayBufferEncode(attestationObject),
            clientDataJSON: arrayBufferEncode(clientDataJSON),
        },
        transports: credResp.getTransports(),
    });
    return credRespString;
}
exports.publicKeyCredentialAttestionToJson = publicKeyCredentialAttestionToJson;
function publicKeyCredentialAssertionToJson(credential) {
    let credResp = credential.response;
    let rawId = credential.rawId;
    let credRespString = JSON.stringify({
        id: credential.id,
        type: credential.type,
        rawId: arrayBufferEncode(rawId),
        clientExtensionResults: credential.getClientExtensionResults(),
        response: credResp,
    });
    return credRespString;
}
exports.publicKeyCredentialAssertionToJson = publicKeyCredentialAssertionToJson;
