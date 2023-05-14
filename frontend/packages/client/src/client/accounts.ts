import { baseUrl } from "./types/utils";
import { Account } from "./types/user";
import axios from "axios";
import { ListAccountsResponse, GetAccountResponse, CreateAccountResponse, ListAccountsResult } from "./types";


export default class Accounts {
    private token: string;
    constructor(token: string) {
        this.token = token;
    }

    /**
     * This is a TypeScript function that asynchronously retrieves a list of accounts from an API and
     * returns them as an array of Account objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of `Account`
     * objects. The `Account` objects are obtained from the `accounts` property of the
     * `ListAccountsResponse` object returned by the API.
     */
    public async list(): Promise<ListAccountsResult> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.get<ListAccountsResult>(`/accounts`);
            console.log(response.data);
            return response.data;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This function retrieves an account object from an API endpoint using an address as a parameter.
     * @param {string} address - The `address` parameter is a string representing the address of the
     * account to be fetched.
     * @returns The `get` method is returning a `Promise` that resolves to an `Account` object. The
     * `Account` object is obtained by extracting the `account` property from the `data` property of the
     * `GetAccountResponse` object returned by the API. If there is an error, the method logs an error
     * message and throws the error.
     */
    public async get(address: string): Promise<Account> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.get<GetAccountResponse>(`/accounts/${address}`);
            console.log(response.data);
            return response.data.account;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
    /**
     * This function creates a new account with a given name and coin type using an API call and returns
     * the newly created account.
     * @param {string} name - A string representing the name of the account to be created.
     * @param {string} coin_type - The type of cryptocurrency for which the account is being created.
     * @returns The `create` method is returning a Promise that resolves to an `Account` object. The
     * `Account` object is obtained from the `new_account` property of the `CreateAccountResponse` object
     * returned by the API call.
     */
    public async create(name: string, coin_type: string): Promise<Account> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.get<CreateAccountResponse>(`/accounts/create/${coin_type}/${name}`);
            console.log(response.data);
            return response.data.new_account;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This is an asynchronous function that signs a message using a specified address and returns the
     * signature.
     * @param {string} address - The Ethereum address of the account that will be used to sign the
     * message.
     * @param {string} message - The message parameter is a string that represents the message to be
     * signed. The sign() function takes this message as input and returns a signature string after
     * signing the message using the private key associated with the given address.
     * @returns a Promise that resolves to a string, which is the signature of the message that was
     * signed using the private key associated with the specified address.
     */
    public async sign(address: string, message: string): Promise<string> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.post(`/accounts/${address}/sign`, { message: message });
            console.log(response.data);
            return response.data.signature;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This function verifies a message signature for a given address using an API call and returns a
     * boolean indicating whether the signature is valid.
     * @param {string} address - The Ethereum address of the account that signed the message.
     * @param {string} message - The message that was signed by the address owner.
     * @param {string} signature - The `signature` parameter is a string representing the digital
     * signature of the `message` parameter, which is signed by the private key associated with the
     * `address` parameter. The `verify` function uses this signature to verify the authenticity of the
     * message and returns a boolean value indicating whether the signature is valid
     * @returns a boolean value indicating whether the provided message and signature are verified for
     * the given address.
     */
    public async verify(address: string, message: string, signature: string): Promise<boolean> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.post(`/accounts/${address}/verify`, { message: message, signature: signature });
            console.log(response.data);
            return response.data.verified;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
}
