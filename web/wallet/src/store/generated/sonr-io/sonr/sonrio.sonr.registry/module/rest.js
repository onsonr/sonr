/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */
/**
* ServiceProtocol are core modules that can be installed on custom services on the Sonr network.

 - SERVICE_PROTOCOL_UNSPECIFIED: SERVICE_PROTOCOL_UNSPECIFIED is the default value.
 - SERVICE_PROTOCOL_BUCKETS: SERVICE_PROTOCOL_BUCKETS is the module that provides the ability to store and retrieve data.
 - SERVICE_PROTOCOL_CHANNEL: SERVICE_PROTOCOL_CHANNEL is the module that provides the ability to communicate with other services.
 - SERVICE_PROTOCOL_OBJECTS: SERVICE_PROTOCOL_OBJECTS is the module that provides the ability to create new schemas for data on the network.
 - SERVICE_PROTOCOL_FUNCTIONS: SERVICE_PROTOCOL_FUNCTIONS is the module that provides the ability to create new functions for data on the network.
*/
export var RegistryServiceProtocol;
(function (RegistryServiceProtocol) {
    RegistryServiceProtocol["SERVICE_PROTOCOL_UNSPECIFIED"] = "SERVICE_PROTOCOL_UNSPECIFIED";
    RegistryServiceProtocol["SERVICE_PROTOCOL_BUCKETS"] = "SERVICE_PROTOCOL_BUCKETS";
    RegistryServiceProtocol["SERVICE_PROTOCOL_CHANNEL"] = "SERVICE_PROTOCOL_CHANNEL";
    RegistryServiceProtocol["SERVICE_PROTOCOL_OBJECTS"] = "SERVICE_PROTOCOL_OBJECTS";
    RegistryServiceProtocol["SERVICE_PROTOCOL_FUNCTIONS"] = "SERVICE_PROTOCOL_FUNCTIONS";
})(RegistryServiceProtocol || (RegistryServiceProtocol = {}));
/**
* ServiceType is the type of service that is being registered.

 - SERVICE_TYPE_UNSPECIFIED: SERVICE_TYPE_UNSPECIFIED is the default value.
 - SERVICE_TYPE_DID_COMM_MESSAGING: SERVICE_TYPE_APPLICATION is the type of service that is a DApp.
 - SERVICE_TYPE_LINKED_DOMAINS: SERVICE_TYPE_SERVICE is the type of service that is a service.
 - SERVICE_TYPE_SONR: SERVICE_TYPE_SONR is the type of service that is a DApp.
*/
export var RegistryServiceType;
(function (RegistryServiceType) {
    RegistryServiceType["SERVICE_TYPE_UNSPECIFIED"] = "SERVICE_TYPE_UNSPECIFIED";
    RegistryServiceType["SERVICE_TYPE_DID_COMM_MESSAGING"] = "SERVICE_TYPE_DID_COMM_MESSAGING";
    RegistryServiceType["SERVICE_TYPE_LINKED_DOMAINS"] = "SERVICE_TYPE_LINKED_DOMAINS";
    RegistryServiceType["SERVICE_TYPE_SONR"] = "SERVICE_TYPE_SONR";
})(RegistryServiceType || (RegistryServiceType = {}));
/**
*  - TYPE_UNSPECIFIED: TYPE_UNSPECIFIED is the default value.
 - TYPE_ECDSA_SECP256K1: TYPE_ECDSA_SECP256K1 represents the Ed25519VerificationKey2018 key type.
 - TYPE_X25519: TYPE_X25519 represents the X25519KeyAgreementKey2019 key type.
 - TYPE_ED25519: TYPE_ED25519 represents the Ed25519VerificationKey2018 key type.
 - TYPE_BLS_12381_G1: TYPE_BLS_12381_G1 represents the Bls12381G1Key2020 key type
 - TYPE_BLS_12381_G2: TYPE_BLS_12381_G2 represents the Bls12381G2Key2020 key type
 - TYPE_RSA: TYPE_RSA represents the RsaVerificationKey2018 key type.
 - TYPE_VERIFIABLE_CONDITION: TYPE_VERIFIABLE_CONDITION represents the VerifiableCondition2021 key type.
*/
export var RegistryVerificationMethodType;
(function (RegistryVerificationMethodType) {
    RegistryVerificationMethodType["TYPE_UNSPECIFIED"] = "TYPE_UNSPECIFIED";
    RegistryVerificationMethodType["TYPEECDSASECP256K1"] = "TYPE_ECDSA_SECP256K1";
    RegistryVerificationMethodType["TYPEX25519"] = "TYPE_X25519";
    RegistryVerificationMethodType["TYPEED25519"] = "TYPE_ED25519";
    RegistryVerificationMethodType["TYPEBLS12381G1"] = "TYPE_BLS_12381_G1";
    RegistryVerificationMethodType["TYPEBLS12381G2"] = "TYPE_BLS_12381_G2";
    RegistryVerificationMethodType["TYPE_RSA"] = "TYPE_RSA";
    RegistryVerificationMethodType["TYPE_VERIFIABLE_CONDITION"] = "TYPE_VERIFIABLE_CONDITION";
})(RegistryVerificationMethodType || (RegistryVerificationMethodType = {}));
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    constructor(apiConfig = {}) {
        this.baseUrl = "";
        this.securityData = null;
        this.securityWorker = null;
        this.abortControllers = new Map();
        this.baseApiParams = {
            credentials: "same-origin",
            headers: {},
            redirect: "follow",
            referrerPolicy: "no-referrer",
        };
        this.setSecurityData = (data) => {
            this.securityData = data;
        };
        this.contentFormatters = {
            [ContentType.Json]: (input) => input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
            [ContentType.FormData]: (input) => Object.keys(input || {}).reduce((data, key) => {
                data.append(key, input[key]);
                return data;
            }, new FormData()),
            [ContentType.UrlEncoded]: (input) => this.toQueryString(input),
        };
        this.createAbortSignal = (cancelToken) => {
            if (this.abortControllers.has(cancelToken)) {
                const abortController = this.abortControllers.get(cancelToken);
                if (abortController) {
                    return abortController.signal;
                }
                return void 0;
            }
            const abortController = new AbortController();
            this.abortControllers.set(cancelToken, abortController);
            return abortController.signal;
        };
        this.abortRequest = (cancelToken) => {
            const abortController = this.abortControllers.get(cancelToken);
            if (abortController) {
                abortController.abort();
                this.abortControllers.delete(cancelToken);
            }
        };
        this.request = ({ body, secure, path, type, query, format = "json", baseUrl, cancelToken, ...params }) => {
            const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
            const requestParams = this.mergeRequestParams(params, secureParams);
            const queryString = query && this.toQueryString(query);
            const payloadFormatter = this.contentFormatters[type || ContentType.Json];
            return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
                ...requestParams,
                headers: {
                    ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
                    ...(requestParams.headers || {}),
                },
                signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
                body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
            }).then(async (response) => {
                const r = response;
                r.data = null;
                r.error = null;
                const data = await response[format]()
                    .then((data) => {
                    if (r.ok) {
                        r.data = data;
                    }
                    else {
                        r.error = data;
                    }
                    return r;
                })
                    .catch((e) => {
                    r.error = e;
                    return r;
                });
                if (cancelToken) {
                    this.abortControllers.delete(cancelToken);
                }
                if (!response.ok)
                    throw data;
                return data;
            });
        };
        Object.assign(this, apiConfig);
    }
    addQueryParam(query, key) {
        const value = query[key];
        return (encodeURIComponent(key) +
            "=" +
            encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`));
    }
    toQueryString(rawQuery) {
        const query = rawQuery || {};
        const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
        return keys
            .map((key) => typeof query[key] === "object" && !Array.isArray(query[key])
            ? this.toQueryString(query[key])
            : this.addQueryParam(query, key))
            .join("&");
    }
    addQueryParams(rawQuery) {
        const queryString = this.toQueryString(rawQuery);
        return queryString ? `?${queryString}` : "";
    }
    mergeRequestParams(params1, params2) {
        return {
            ...this.baseApiParams,
            ...params1,
            ...(params2 || {}),
            headers: {
                ...(this.baseApiParams.headers || {}),
                ...(params1.headers || {}),
                ...((params2 && params2.headers) || {}),
            },
        };
    }
}
/**
 * @title registry/config.proto
 * @version version not set
 */
export class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryParams
         * @summary Parameters queries the parameters of the module.
         * @request GET:/sonrio/sonr/registry/params
         */
        this.queryParams = (params = {}) => this.request({
            path: `/sonrio/sonr/registry/params`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
