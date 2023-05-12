"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("./types/utils");
class DID {
    /**
     * This function fetches a list of DID documents from an API endpoint and returns them as an array
     * of DidDocument objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of `DidDocument`
     * objects. The `DidDocument` type is not defined in the code snippet provided, so it is unclear
     * what properties and methods it has.
     */
    async list() {
        try {
            const response = await (0, utils_1.getAxios)(false).get(`/id`);
            return response.data.documents;
        }
        catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
    /**
     * This function retrieves a DID document from an API endpoint using a provided DID string.
     * @param {string} did - The `did` parameter is a string representing the Decentralized Identifier
     * (DID) that we want to fetch the DID document for.
     * @returns The `get` method is returning a `Promise` that resolves to a `DidDocument` object.
     */
    async get(did) {
        try {
            const response = await (0, utils_1.getAxios)(false).get(`/id/${did}`);
            console.log(response.data);
            return response.data;
        }
        catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
    /**
     * This function retrieves a DID document by its alias using an API call and returns it as a Promise.
     * @param {string} alias - The `alias` parameter is a string representing the alias of a DID
     * (Decentralized Identifier) document. This method is used to fetch the DID document associated with
     * the given alias.
     * @returns a Promise that resolves to a DidDocument object. The DidDocument object is obtained by
     * making an API call to fetch a QueryDocumentResponse object, and then extracting the did_document
     * property from it. If there is an error during the API call, the function will log an error message
     * and re-throw the error.
     */
    async getByAlias(alias) {
        try {
            const response = await (0, utils_1.getAxios)(false).get(`/id/alias/${alias}`);
            console.log(response.data);
            return response.data;
        }
        catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
    /**
     * This function retrieves a DID document by owner address using an API call and returns it as a
     * Promise.
     * @param {string} address - The `address` parameter is a string representing the Ethereum address of
     * the owner of a DID (Decentralized Identifier). This function is used to fetch the DID document
     * associated with the given owner address.
     * @returns This function returns a Promise that resolves to a DidDocument object. The DidDocument is
     * obtained by making a GET request to the API endpoint `/id/owner/` and extracting the
     * `did_document` property from the response data. If an error occurs during the API call, the function
     * logs an error message and throws the error.
     */
    async getByOwner(address) {
        try {
            const response = await (0, utils_1.getAxios)(false).get(`/id/owner/${address}`);
            console.log(response.data);
            return response.data.did_document;
        }
        catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }
}
exports.default = DID;
