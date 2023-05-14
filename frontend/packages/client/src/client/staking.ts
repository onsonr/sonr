import { baseUrl, getAxios } from "./types/utils";
import { trytm } from "@bdsqqq/try";
import { ServiceRecord } from "./types/service";
import { QueryAttestionResponse } from "./types";
import axios from "axios";

export default class Staking {
    private token: string;

    constructor(jwt: string) {
        this.token = jwt;
    }

    /**
     * This is an asynchronous function that retrieves a list of service records using Axios and returns
     * them as an array of ServiceRecord objects.
     * @returns The `list()` function is returning a Promise that resolves to an array of
     * `ServiceRecord` objects. The `ServiceRecord` objects are obtained from the
     * `response.data.services` property, which is the array of services returned by the API endpoint
     * `/service`.
     */
    async listValidators(): Promise<ServiceRecord[]> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.get(`/staking/validators`);
            return response.data.services;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This function retrieves a service record by name from an API and returns it as a Promise.
     * @param {string} valAddress - A string representing the name of the service to be fetched.
     * @returns The `get` function is returning a `Promise` that resolves to a `ServiceRecord` object.
     * The `ServiceRecord` object is obtained from the `response.data.service` property.
     */
    async listDelegators(valAddress: string): Promise<ServiceRecord> {
        try {
            const response = await getAxios(false).get(`/staking/validators/${valAddress}`);
            console.log(response.data);
            return response.data.service;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
   * This is an asynchronous function that retrieves a list of service records using Axios and returns
   * them as an array of ServiceRecord objects.
   * @returns The `list()` function is returning a Promise that resolves to an array of
   * `ServiceRecord` objects. The `ServiceRecord` objects are obtained from the
   * `response.data.services` property, which is the array of services returned by the API endpoint
   * `/service`.
   */
    async delegate(valAddress: string, amount: number): Promise<ServiceRecord[]> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.post(`/staking/validators/${valAddress}/delegate`, null, { params: { amount: amount } });
            return response.data.services;
        } catch (error) {
            console.error(`Error fetching DID: ${error}`);
            throw error;
        }
    }

    /**
     * This function retrieves a service record by name from an API and returns it as a Promise.
     * @param {string} valAddress - A string representing the name of the service to be fetched.
     * @returns The `get` function is returning a `Promise` that resolves to a `ServiceRecord` object.
     * The `ServiceRecord` object is obtained from the `response.data.service` property.
     */
    async undelegate(valAddress: string, amount: number): Promise<ServiceRecord> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        try {
            const response = await api.post(`/staking/validators/${valAddress}/undelegate`, null, { params: { amount: amount } });
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
    async cancelDelegation(valAddress: string, amount: number, creationHeight: number): Promise<QueryAttestionResponse> {
        let api = axios.create({
            baseURL: baseUrl(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        })
        const [aData, aError] = await trytm(api.post(`/staking/validators/${valAddress}/cancel`, null, {
            params: {
                amount: amount,
                creationHeight: creationHeight,
            },
        }));
        if (aError) {
            throw aError;
        }
        console.log(aData.data);
        return aData.data;
    }
}
