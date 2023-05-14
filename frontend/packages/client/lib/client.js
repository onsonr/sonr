"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const did_1 = __importDefault(require("./did"));
const accounts_1 = __importDefault(require("./accounts"));
const services_1 = __importDefault(require("./services"));
const webauthn_1 = __importDefault(require("./webauthn"));
const utils_1 = require("./types/utils");
const try_1 = require("@bdsqqq/try");
const mailbox_1 = __importDefault(require("./mailbox"));
const staking_1 = __importDefault(require("./staking"));
const axios_1 = __importDefault(require("axios"));
/**
 * The `ApiClient` class is the main class of the client library. It is used to interact with the
 * Sonr Blockchain.
 **/
class SonrClient {
    /**
     * This is a constructor function that initializes a DID and Services object with a given origin.
     * @param {string} origin - The `origin` parameter is a string that represents the origin of the
     * DID (Decentralized Identifier) and the services associated with it. It is used to initialize the
     * `DID` and `Services` objects in the constructor.
     */
    constructor(origin) {
        this.did = new did_1.default();
        this.webauthn = new webauthn_1.default(origin);
        this.services = new services_1.default(origin);
    }
    /**
     * The function checks if the user is authenticated by verifying if their account information is
     * defined.
     * @returns A boolean value is being returned. The method `isAuthenticated()` checks if the `accounts`
     * property is defined and returns `true` if it is defined, and `false` otherwise.
     */
    isAuthenticated() {
        return this.accounts !== undefined && this._primaryDoc !== null && this._address !== undefined && this._address !== null && this._jwt !== undefined;
    }
    /**
     * The getAddress function returns the address as a string.
     * @returns A string representing the address.
     */
    getAddress() {
        return this._address;
    }
    /**
     * This is an asynchronous function that retrieves a block response from a specified URL using axios in
     * TypeScript.
     * @returns The `getBlock()` function is returning a `Promise` that resolves to a `BlockResponse`
     * object. The `BlockResponse` object is obtained by making a GET request to the
     * "https://rpc.sonr.ws/block" endpoint using the `axios` library. The `resp.data` property of the
     * response object is returned as the result of the `getBlock()` function.
     */
    async getBlock() {
        let resp = await axios_1.default.get("https://rpc.sonr.ws/block");
        return resp.data;
    }
    /**
     * This function returns the primary document of type DidDocument.
     * @returns The `getPrimaryDoc()` method is returning the `_primaryDoc` property, which is of type
     * `DidDocument`.
     */
    getPrimaryDoc() {
        return this._primaryDoc;
    }
    /**
     * This function registers a user by generating a web authentication credential and sending it to the
     * server for verification.
     * @param {string} username - A string representing the username of the user who is registering.
     * @returns The `register` function returns a `Promise` that resolves to a `KeygenResponse` object.
     */
    async register({ alias, onCredentialSet, onRegisterComplete }) {
        const [aData, aError] = await (0, try_1.trytm)(this.services.startRegistration(alias));
        if (aError) {
            throw aError;
        }
        // Generate WebAuthn Credential
        const credential = await this.webauthn.generateWebAuthnCredential(aData.attestion_options);
        onCredentialSet(credential);
        let pubKeystr = (0, utils_1.publicKeyCredentialAttestionToJson)(credential);
        const [bData, bError] = await (0, try_1.trytm)(this.services.finishRegistration(alias, pubKeystr, aData.challenge, aData.ucw_id));
        if (bError) {
            throw bError;
        }
        // Setup Accounts Endpoint
        onRegisterComplete(bData.did, bData.primary, bData.jwt);
        // Setup address
        this._address = bData.address;
        this._primaryDoc = bData.primary;
        this._initJwt(bData.jwt);
        return bData;
    }
    /**
     * This is an async function that logs in a user by starting and finishing a web authentication
     * process.
     * @param {string} alias - The `alias` parameter is a string that represents the user's alias or
     * username used for authentication.
     * @returns The `login` function returns a Promise that resolves to a `LoginResponse` object.
     */
    async login({ alias, onCredentialSet, onLoginComplete }) {
        const [bData, bError] = await (0, try_1.trytm)(this.services.startLogin(alias));
        if (bError) {
            throw bError;
        }
        const credential = await this.webauthn.authenticateWebAuthnCredential(bData.assertion_options);
        onCredentialSet(credential);
        let pubKeystr = JSON.stringify(credential);
        const [cData, cError] = await (0, try_1.trytm)(this.services.finishLogin(alias, pubKeystr));
        if (cError) {
            throw cError;
        }
        this._address = bData.address;
        this._primaryDoc = cData.did_document;
        onLoginComplete(cData.did_document.id, cData.did_document, cData.jwt);
        this._initJwt(cData.jwt);
        return cData;
    }
    /* The `_initJwt(jwt: string)` function initializes the `jwt` property of the `ApiClient` class
    with the provided `jwt` string. It also creates new instances of the `Accounts` and `Mailbox`
    classes with the `jwt` string as a parameter, which allows the `ApiClient` to interact with the
    Sonr Blockchain using the provided `jwt` token for authentication. */
    _initJwt(jwt) {
        this._jwt = jwt;
        this.accounts = new accounts_1.default(jwt);
        this.mailbox = new mailbox_1.default(jwt);
        this.staking = new staking_1.default(jwt);
    }
}
exports.default = SonrClient;
