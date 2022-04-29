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

export type RegistryMsgRegisterServiceResponse = object;

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
export type RegistryParams = object;

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
export enum RegistryServiceProtocol {
  SERVICE_PROTOCOL_UNSPECIFIED = "SERVICE_PROTOCOL_UNSPECIFIED",
  SERVICE_PROTOCOL_BUCKETS = "SERVICE_PROTOCOL_BUCKETS",
  SERVICE_PROTOCOL_CHANNEL = "SERVICE_PROTOCOL_CHANNEL",
  SERVICE_PROTOCOL_OBJECTS = "SERVICE_PROTOCOL_OBJECTS",
  SERVICE_PROTOCOL_FUNCTIONS = "SERVICE_PROTOCOL_FUNCTIONS",
}

/**
* ServiceType is the type of service that is being registered.

 - SERVICE_TYPE_UNSPECIFIED: SERVICE_TYPE_UNSPECIFIED is the default value.
 - SERVICE_TYPE_DID_COMM_MESSAGING: SERVICE_TYPE_APPLICATION is the type of service that is a DApp.
 - SERVICE_TYPE_LINKED_DOMAINS: SERVICE_TYPE_SERVICE is the type of service that is a service.
 - SERVICE_TYPE_SONR: SERVICE_TYPE_SONR is the type of service that is a DApp.
*/
export enum RegistryServiceType {
  SERVICE_TYPE_UNSPECIFIED = "SERVICE_TYPE_UNSPECIFIED",
  SERVICE_TYPE_DID_COMM_MESSAGING = "SERVICE_TYPE_DID_COMM_MESSAGING",
  SERVICE_TYPE_LINKED_DOMAINS = "SERVICE_TYPE_LINKED_DOMAINS",
  SERVICE_TYPE_SONR = "SERVICE_TYPE_SONR",
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
export enum RegistryVerificationMethodType {
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  TYPEECDSASECP256K1 = "TYPE_ECDSA_SECP256K1",
  TYPEX25519 = "TYPE_X25519",
  TYPEED25519 = "TYPE_ED25519",
  TYPEBLS12381G1 = "TYPE_BLS_12381_G1",
  TYPEBLS12381G2 = "TYPE_BLS_12381_G2",
  TYPE_RSA = "TYPE_RSA",
  TYPE_VERIFIABLE_CONDITION = "TYPE_VERIFIABLE_CONDITION",
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

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

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
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

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
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

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
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
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
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

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title registry/config.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/sonrio/sonr/registry/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<RegistryQueryParamsResponse, GooglerpcStatus>({
      path: `/sonrio/sonr/registry/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
