import { DidDocument } from "./types/did";
import { QueryAliasResponse } from "./types";
export default class DID {
    /**
     * This function fetches a list of DID documents from an API endpoint and returns them as an array
     * of DidDocument objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of `DidDocument`
     * objects. The `DidDocument` type is not defined in the code snippet provided, so it is unclear
     * what properties and methods it has.
     */
    list(): Promise<DidDocument[]>;
    /**
     * This function retrieves a DID document from an API endpoint using a provided DID string.
     * @param {string} did - The `did` parameter is a string representing the Decentralized Identifier
     * (DID) that we want to fetch the DID document for.
     * @returns The `get` method is returning a `Promise` that resolves to a `DidDocument` object.
     */
    get(did: string): Promise<DidDocument>;
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
    getByAlias(alias: string): Promise<QueryAliasResponse>;
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
    getByOwner(address: string): Promise<DidDocument>;
}
