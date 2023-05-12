import { getAxios } from "./types/utils";
import { trytm } from "@bdsqqq/try";
import { ServiceRecord } from "./types/service";
import { LoginResponse, QueryAssertionResponse, QueryAttestionResponse, RegistrationResponse } from "./types";

export default class Services {
    private origin: string;

    constructor(origin: string) {
        this.origin = origin;
    }

    /**
     * This is an asynchronous function that retrieves a list of service records using Axios and returns
     * them as an array of ServiceRecord objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of
     * `ServiceRecord` objects. The `ServiceRecord` objects are obtained from the
     * `response.data.services` property, which is the array of services returned by the API endpoint
     * `/service`.
     */
    async list(): Promise<ServiceRecord[]> {
        try {
            const response = await getAxios(false).get(`/service`);
            return response.data.services;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This function retrieves a service record by name from an API and returns it as a Promise.
     * @param {string} name - A string representing the name of the service to be fetched.
     * @returns The `get` function is returning a `Promise` that resolves to a `ServiceRecord` object.
     * The `ServiceRecord` object is obtained from the `response.data.service` property.
     */
    async get(name: string): Promise<ServiceRecord> {
        try {
            const response = await getAxios(false).get(`/service/${name}`);
            console.log(response.data);
            return response.data.service;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This is an asynchronous function that starts the registration process for a user and returns a
     * Promise containing a QueryAttestionResponse object.
     * @param {string} alias - The `alias` parameter is a string that represents the username of
     * the user who is starting the registration process. This parameter is used to identify the user
     * and create a new registration request for them.
     * @returns The function `startRegistration` returns a Promise that resolves to a
     * `QueryAttestionResponse` object.
     */
    async startRegistration(alias: string): Promise<QueryAttestionResponse> {
        const [aData, aError] = await trytm(getAxios(false).get<QueryAttestionResponse>(`/service/${this.origin}/register/start`, {
            params: {
                alias: alias,
            },
        }));
        if (aError) {
            console.log(aError);
            throw aError;
        }
        console.log(aData.data);
        return aData.data;
    }

    /**
     * This is an async function that finishes the registration process by sending a POST request with
     * the user's username and credential response, and returns a promise with a RegistrationResponse
     * object.
     * @param {string} alias - The `alias` parameter is a string that represents the username of
     * the user who is finishing the registration process. It is used to identify the user and link
     * their account to the registration process.
     * @param {string} credential_response - The `credential_response` parameter is a string that
     * contains the response from a credential provider, such as a password or authentication code,
     * that is used to complete the registration process for a user.
     * @returns a Promise that resolves to a `RegistrationResponse` object.
     */
    async finishRegistration(alias: string, credential_response: string, challenge: string, ucwId: string): Promise<RegistrationResponse> {
        const [aData, aError] = await trytm(getAxios(false).post(`/service/${this.origin}/register/finish`, null, {
            params: {
                attestion: credential_response,
                alias: alias,
                challenge: challenge,
                ucw_id: ucwId,
            },
        }));
        if (aError) {
            console.log(aError);
            throw aError;
        }
        console.log(aData.data);
        return aData.data;
    }

    /**
     * This is an asynchronous function that starts a login process and returns a Promise containing a
     * QueryAssertionResponse object.
     * @param {string} address - The `address` parameter is a string that represents the underlying idx of
     * the user who is trying to log in. It is used as a parameter for the `startLogin` function.
     * @returns a Promise that resolves to a QueryAssertionResponse object.
     */
    async startLogin(alias: string): Promise<QueryAssertionResponse> {
        const [aData, aError] = await trytm(getAxios(false).get<QueryAssertionResponse>(`/service/${this.origin}/login/start`, {
            params: {
                alias: alias,
            },
        }));
        if (aError) {
            console.log(aError);
            throw aError;
        }
        console.log(aData.data);
        return aData.data;
    }

    /**
     * This is an async function that finishes the login process by sending a post request with account
     * address, assertion response, and origin, and returns a LoginResponse object.
     * @param {string} address - The `address` parameter is a string representing the account address
     * of the user who is trying to log in.
     * @param {string} assertion_response - `assertion_response` is a string parameter that represents
     * the response received from the authentication provider after the user has successfully
     * authenticated. This response typically contains information about the user's identity and
     * authentication status, which is then used to complete the login process.
     * @returns a Promise that resolves to a `LoginResponse` object.
     */
    async finishLogin(alias: string, assertion_response: string): Promise<LoginResponse> {
        const [aData, aError] = await trytm(getAxios(false).post(`/service/${this.origin}/login/finish`, null, {
            params: {
                alias: alias,
                assertion: assertion_response,
            }
        }));
        if (aError) {
            console.log(aError.message);
            throw aError;
        }
        console.log(aData.data);
        return aData.data
    }
}
