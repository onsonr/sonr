import DID from './did';
import Accounts from './accounts';
import Services from './services';
import Webauthn from "./webauthn";
import { BlockResponse, LoginResponse, RegistrationResponse } from './types/api';
import { DidDocument } from './types';
import Mailbox from './mailbox';
import { SonrLoginProps, SonrRegisterProps } from './types/props';
import Staking from './staking';
/**
 * The `ApiClient` class is the main class of the client library. It is used to interact with the
 * Sonr Blockchain.
 **/
export default class SonrClient {
    did: DID;
    services: Services;
    webauthn: Webauthn;
    accounts: Accounts;
    mailbox: Mailbox;
    staking: Staking;
    private _address;
    private _primaryDoc;
    private _jwt;
    /**
     * This is a constructor function that initializes a DID and Services object with a given origin.
     * @param {string} origin - The `origin` parameter is a string that represents the origin of the
     * DID (Decentralized Identifier) and the services associated with it. It is used to initialize the
     * `DID` and `Services` objects in the constructor.
     */
    constructor(origin: string);
    /**
     * The function checks if the user is authenticated by verifying if their account information is
     * defined.
     * @returns A boolean value is being returned. The method `isAuthenticated()` checks if the `accounts`
     * property is defined and returns `true` if it is defined, and `false` otherwise.
     */
    isAuthenticated(): boolean;
    /**
     * The getAddress function returns the address as a string.
     * @returns A string representing the address.
     */
    getAddress(): string;
    /**
     * This is an asynchronous function that retrieves a block response from a specified URL using axios in
     * TypeScript.
     * @returns The `getBlock()` function is returning a `Promise` that resolves to a `BlockResponse`
     * object. The `BlockResponse` object is obtained by making a GET request to the
     * "https://rpc.sonr.ws/block" endpoint using the `axios` library. The `resp.data` property of the
     * response object is returned as the result of the `getBlock()` function.
     */
    getBlock(): Promise<BlockResponse>;
    /**
     * This function returns the primary document of type DidDocument.
     * @returns The `getPrimaryDoc()` method is returning the `_primaryDoc` property, which is of type
     * `DidDocument`.
     */
    getPrimaryDoc(): DidDocument;
    /**
     * This function registers a user by generating a web authentication credential and sending it to the
     * server for verification.
     * @param {string} username - A string representing the username of the user who is registering.
     * @returns The `register` function returns a `Promise` that resolves to a `KeygenResponse` object.
     */
    register({ alias, onCredentialSet, onRegisterComplete }: SonrRegisterProps): Promise<RegistrationResponse>;
    /**
     * This is an async function that logs in a user by starting and finishing a web authentication
     * process.
     * @param {string} alias - The `alias` parameter is a string that represents the user's alias or
     * username used for authentication.
     * @returns The `login` function returns a Promise that resolves to a `LoginResponse` object.
     */
    login({ alias, onCredentialSet, onLoginComplete }: SonrLoginProps): Promise<LoginResponse>;
    _initJwt(jwt: string): void;
}
