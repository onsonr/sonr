export interface GooglerpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export interface ProtobufAny {
    "@type"?: string;
}
/**
* Did represents a string that has been parsed and validated as a DID. The parts are stored
in the individual fields.
*/
export interface RegistryDid {
    /** Method is the method used to create the DID. For the Sonr network it is "sonr". */
    method?: string;
    /** Network is the network the DID is on. For testnet it is "testnet". i.e "did:sonr:testnet:". */
    network?: string;
    id?: string;
    /** Paths is a list of paths that the DID is valid for. This is used to identify the Service. */
    paths?: string[];
    /** Query is the query string that was used to create the DID. This is followed by a '?'. */
    query?: string;
    /** Fragment is the fragment string that was used to create the DID. This is followed by a '#'. */
    fragment?: string;
}
/**
 * DidDocument is the document that describes a DID. This document is stored on the blockchain.
 */
export interface RegistryDidDocument {
    /** Context is the context of the DID document. This is used to identify the Service. */
    context?: string[];
    /** Id is the DID of the document. */
    id?: string;
    /** Controller is the DID of the controller of the document. This will be the individual user devices and mailboxes. */
    controller?: string[];
    /** VerificationMethod is the list of verification methods for the user. */
    verificationMethod?: RegistryVerificationMethod[];
    /** Authentication is the list of authentication methods for the user. */
    authentication?: string[];
    /** AssertionMethod is the list of assertion methods for the user. */
    assertionMethod?: string[];
    /** CapabilityInvocation is the list of capability invocation methods for the user. */
    capabilityInvocation?: string[];
    /** CapabilityDelegation is the list of capability delegation methods for the user. */
    capabilityDelegation?: string[];
    /** KeyAgreement is the list of key agreement methods for the user. */
    keyAgreement?: string[];
    /** Service is the list of services or DApps that the user has access to. */
    service?: RegistryService[];
    /** AlsoKnownAs is the list of ".snr" aliases for the user. */
    alsoKnownAs?: string[];
    /** Metadata is the metadata of the service. */
    metadata?: Record<string, string>;
}
export interface RegistryMsgAccessNameResponse {
    name?: string;
    publicKey?: string;
    peerId?: string;
}
export interface RegistryMsgAccessServiceResponse {
    /** @format int32 */
    code?: number;
    message?: string;
    metadata?: Record<string, string>;
}
export interface RegistryMsgRegisterNameResponse {
    isSuccess?: boolean;
    /**
     * Did represents a string that has been parsed and validated as a DID. The parts are stored
     * in the individual fields.
     */
    did?: RegistryDid;
    /** DidDocument is the document that describes a DID. This document is stored on the blockchain. */
    didDocument?: RegistryDidDocument;
}
export declare type RegistryMsgRegisterServiceResponse = object;
export interface RegistryMsgUpdateNameResponse {
    /** DidDocument is the document that describes a DID. This document is stored on the blockchain. */
    didDocument?: RegistryDidDocument;
    metadata?: Record<string, string>;
}
export interface RegistryMsgUpdateServiceResponse {
    /** DidDocument is the document that describes a DID. This document is stored on the blockchain. */
    didDocument?: RegistryDidDocument;
    configuration?: Record<string, string>;
    metadata?: Record<string, string>;
}
/**
 * Params defines the parameters for the module.
 */
export declare type RegistryParams = object;
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RegistryQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: RegistryParams;
}
/**
 * Service is a Application that runs on the Sonr network.
 */
export interface RegistryService {
    /** ID is the DID of the service. */
    id?: string;
    /** Type is the type of the service. */
    type?: RegistryServiceType;
    /** ServiceEndpoint is the endpoint of the service. */
    serviceEndpoint?: RegistryServiceEndpoint;
    /** Metadata is the metadata of the service. */
    metadata?: Record<string, string>;
}
/**
 * ServiceEndpoint is the endpoint of the service.
 */
export interface RegistryServiceEndpoint {
    /** TransportType is the type of transport used to connect to the service. */
    transportType?: string;
    /** Network is the network the service is on. */
    network?: string;
    supportedProtocols?: RegistryServiceProtocol[];
}
/**
* ServiceProtocol are core modules that can be installed on custom services on the Sonr network.

 - SERVICE_PROTOCOL_UNSPECIFIED: SERVICE_PROTOCOL_UNSPECIFIED is the default value.
 - SERVICE_PROTOCOL_BUCKETS: SERVICE_PROTOCOL_BUCKETS is the module that provides the ability to store and retrieve data.
 - SERVICE_PROTOCOL_CHANNEL: SERVICE_PROTOCOL_CHANNEL is the module that provides the ability to communicate with other services.
 - SERVICE_PROTOCOL_OBJECTS: SERVICE_PROTOCOL_OBJECTS is the module that provides the ability to create new schemas for data on the network.
 - SERVICE_PROTOCOL_FUNCTIONS: SERVICE_PROTOCOL_FUNCTIONS is the module that provides the ability to create new functions for data on the network.
*/
export declare enum RegistryServiceProtocol {
    SERVICE_PROTOCOL_UNSPECIFIED = "SERVICE_PROTOCOL_UNSPECIFIED",
    SERVICE_PROTOCOL_BUCKETS = "SERVICE_PROTOCOL_BUCKETS",
    SERVICE_PROTOCOL_CHANNEL = "SERVICE_PROTOCOL_CHANNEL",
    SERVICE_PROTOCOL_OBJECTS = "SERVICE_PROTOCOL_OBJECTS",
    SERVICE_PROTOCOL_FUNCTIONS = "SERVICE_PROTOCOL_FUNCTIONS"
}
/**
* ServiceType is the type of service that is being registered.

 - SERVICE_TYPE_UNSPECIFIED: SERVICE_TYPE_UNSPECIFIED is the default value.
 - SERVICE_TYPE_DID_COMM_MESSAGING: SERVICE_TYPE_APPLICATION is the type of service that is a DApp.
 - SERVICE_TYPE_LINKED_DOMAINS: SERVICE_TYPE_SERVICE is the type of service that is a service.
 - SERVICE_TYPE_SONR: SERVICE_TYPE_SONR is the type of service that is a DApp.
*/
export declare enum RegistryServiceType {
    SERVICE_TYPE_UNSPECIFIED = "SERVICE_TYPE_UNSPECIFIED",
    SERVICE_TYPE_DID_COMM_MESSAGING = "SERVICE_TYPE_DID_COMM_MESSAGING",
    SERVICE_TYPE_LINKED_DOMAINS = "SERVICE_TYPE_LINKED_DOMAINS",
    SERVICE_TYPE_SONR = "SERVICE_TYPE_SONR"
}
/**
 * VerificationMethod is a method that can be used to verify the DID.
 */
export interface RegistryVerificationMethod {
    /** ID is the DID of the verification method. */
    id?: string;
    /** Type is the type of the verification method. */
    type?: RegistryVerificationMethodType;
    /** Controller is the DID of the controller of the verification method. */
    controller?: string;
    /** PublicKeyHex is the public key of the verification method in hexidecimal. */
    publicKeyHex?: string;
    /** PublicKeyBase58 is the public key of the verification method in base58. */
    publicKeyBase58?: string;
    /** BlockchainAccountId is the blockchain account id of the verification method. */
    blockchainAccountId?: string;
}
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
export declare enum RegistryVerificationMethodType {
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
    TYPEECDSASECP256K1 = "TYPE_ECDSA_SECP256K1",
    TYPEX25519 = "TYPE_X25519",
    TYPEED25519 = "TYPE_ED25519",
    TYPEBLS12381G1 = "TYPE_BLS_12381_G1",
    TYPEBLS12381G2 = "TYPE_BLS_12381_G2",
    TYPE_RSA = "TYPE_RSA",
    TYPE_VERIFIABLE_CONDITION = "TYPE_VERIFIABLE_CONDITION"
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title registry/config.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/sonrio/sonr/registry/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<RegistryQueryParamsResponse, GooglerpcStatus>>;
}
export {};
