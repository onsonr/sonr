import { ServiceRecord } from "./types/service";
import { QueryAttestionResponse } from "./types";
export default class Staking {
    private token;
    constructor(jwt: string);
    /**
     * This is an asynchronous function that retrieves a list of service records using Axios and returns
     * them as an array of ServiceRecord objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of
     * `ServiceRecord` objects. The `ServiceRecord` objects are obtained from the
     * `response.data.services` property, which is the array of services returned by the API endpoint
     * `/service`.
     */
    listValidators(): Promise<ServiceRecord[]>;
    /**
     * This function retrieves a service record by name from an API and returns it as a Promise.
     * @param {string} valAddress - A string representing the name of the service to be fetched.
     * @returns The `get` function is returning a `Promise` that resolves to a `ServiceRecord` object.
     * The `ServiceRecord` object is obtained from the `response.data.service` property.
     */
    listDelegators(valAddress: string): Promise<ServiceRecord>;
    /**
   * This is an asynchronous function that retrieves a list of service records using Axios and returns
   * them as an array of ServiceRecord objects.
   * @returns The `list()` function is returning a Promise that resolves to an array of
   * `ServiceRecord` objects. The `ServiceRecord` objects are obtained from the
   * `response.data.services` property, which is the array of services returned by the API endpoint
   * `/service`.
   */
    delegate(valAddress: string, amount: number): Promise<ServiceRecord[]>;
    /**
     * This function retrieves a service record by name from an API and returns it as a Promise.
     * @param {string} valAddress - A string representing the name of the service to be fetched.
     * @returns The `get` function is returning a `Promise` that resolves to a `ServiceRecord` object.
     * The `ServiceRecord` object is obtained from the `response.data.service` property.
     */
    undelegate(valAddress: string, amount: number): Promise<ServiceRecord>;
    /**
     * This is an asynchronous function that starts the registration process for a user and returns a
     * Promise containing a QueryAttestionResponse object.
     * @param {string} alias - The `alias` parameter is a string that represents the username of
     * the user who is starting the registration process. This parameter is used to identify the user
     * and create a new registration request for them.
     * @returns The function `startRegistration` returns a Promise that resolves to a
     * `QueryAttestionResponse` object.
     */
    cancelDelegation(valAddress: string, amount: number, creationHeight: number): Promise<QueryAttestionResponse>;
}
