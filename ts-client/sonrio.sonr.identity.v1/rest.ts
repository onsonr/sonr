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

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface V1DidDocument {
  /** optional */
  context?: string[];

  /** optional */
  creator?: string;
  iD?: string;

  /** optional */
  controller?: string[];

  /** optional */
  verification_method?: V1VerificationMethods;

  /** optional */
  authentication?: V1VerificationRelationships;

  /** optional */
  assertion_method?: V1VerificationRelationships;

  /** optional */
  capability_invocation?: V1VerificationRelationships;

  /** optional */
  capability_delegation?: V1VerificationRelationships;

  /** optional */
  key_agreement?: V1VerificationRelationships;

  /** optional */
  service?: V1Services;

  /** optional */
  also_known_as?: string[];
}

/**
* KeyType is the type of key used to sign a DID document.

 - KeyType_UNSPECIFIED: No key type specified
 - KeyType_JSON_WEB_KEY_2020: JsonWebKey2020 is a VerificationMethod type. https://w3c-ccg.github.io/lds-jws2020/
 - KeyType_ED25519_VERIFICATION_KEY_2018: ED25519VerificationKey2018 is the Ed25519VerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/
 - KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019: ECDSASECP256K1VerificationKey2019 is the EcdsaSecp256k1VerificationKey2019 verification key type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
 - KeyType_RSA_VERIFICATION_KEY_2018: RSAVerificationKey2018 is the RsaVerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-rsa2018/
*/
export enum V1KeyType {
  KeyTypeUNSPECIFIED = "KeyType_UNSPECIFIED",
  KeyTypeJSONWEBKEY2020 = "KeyType_JSON_WEB_KEY_2020",
  KeyTypeED25519VERIFICATIONKEY2018 = "KeyType_ED25519_VERIFICATION_KEY_2018",
  KeyTypeECDSASECP256K1VERIFICATIONKEY2019 = "KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019",
  KeyTypeRSAVERIFICATIONKEY2018 = "KeyType_RSA_VERIFICATION_KEY_2018",
}

export type V1MsgCreateDidDocumentResponse = object;

export type V1MsgDeleteDidDocumentResponse = object;

export type V1MsgUpdateDidDocumentResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface V1Params {
  did_base_context?: string;
  did_implementation_context?: string;
  ipfs_gateway?: string;
  ipfs_api?: string;
}

export interface V1QueryAllDidResponse {
  didDocument?: V1DidDocument[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1QueryByAlsoKnownAsResponse {
  didDocument?: V1DidDocument;
}

export interface V1QueryByKeyIDResponse {
  didDocument?: V1DidDocument;
}

export interface V1QueryByMethodResponse {
  didDocument?: V1DidDocument[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1QueryByServiceResponse {
  didDocument?: V1DidDocument;
}

export interface V1QueryGetDidResponse {
  didDocument?: V1DidDocument;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface V1QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: V1Params;
}

export interface V1Service {
  iD?: string;
  type?: string;
  service_endpoint?: string;

  /** optional */
  service_endpoints?: Record<string, string>;
}

export interface V1Services {
  data?: V1Service[];
}

export interface V1VerificationMethod {
  iD?: string;

  /**
   * KeyType is the type of key used to sign a DID document.
   *
   *  - KeyType_UNSPECIFIED: No key type specified
   *  - KeyType_JSON_WEB_KEY_2020: JsonWebKey2020 is a VerificationMethod type. https://w3c-ccg.github.io/lds-jws2020/
   *  - KeyType_ED25519_VERIFICATION_KEY_2018: ED25519VerificationKey2018 is the Ed25519VerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/
   *  - KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019: ECDSASECP256K1VerificationKey2019 is the EcdsaSecp256k1VerificationKey2019 verification key type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
   *  - KeyType_RSA_VERIFICATION_KEY_2018: RSAVerificationKey2018 is the RsaVerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-rsa2018/
   */
  type?: V1KeyType;
  controller?: string;
  public_key_jwk?: Record<string, string>;

  /** optional */
  public_key_multibase?: string;
}

export interface V1VerificationMethods {
  data?: V1VerificationMethod[];
}

export interface V1VerificationRelationship {
  verification_method?: V1VerificationMethod;
  reference?: string;
}

export interface V1VerificationRelationships {
  data?: V1VerificationRelationship[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title sonr/identity/v1/did.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryByAlsoKnownAs
   * @summary Queries a DIDDocument for the matching AlsoKnownAs
   * @request GET:/sonr-io/sonr/identity/aka/{aka_id}
   */
  queryQueryByAlsoKnownAs = (akaId: string, params: RequestParams = {}) =>
    this.request<V1QueryByAlsoKnownAsResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/aka/${akaId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidAll
   * @summary Queries a list of DidDocument items.
   * @request GET:/sonr-io/sonr/identity/did
   */
  queryDidAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1QueryAllDidResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/did`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDid
   * @summary Queries a DidDocument by index.
   * @request GET:/sonr-io/sonr/identity/did/{did}
   */
  queryDid = (did: string, params: RequestParams = {}) =>
    this.request<V1QueryGetDidResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/did/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryByKeyId
   * @summary Queries a DIDDocument for the matching key
   * @request GET:/sonr-io/sonr/identity/key/{key_id}
   */
  queryQueryByKeyId = (keyId: string, params: RequestParams = {}) =>
    this.request<V1QueryByKeyIDResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/key/${keyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryByMethod
   * @summary Queries a list of DIDDocument for the matching method
   * @request GET:/sonr-io/sonr/identity/method/{method_id}
   */
  queryQueryByMethod = (
    methodId: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1QueryByMethodResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/method/${methodId}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/sonr-io/sonr/identity/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<V1QueryParamsResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryByService
   * @summary Queries a DIDDocument for the matching service
   * @request GET:/sonr-io/sonr/identity/service/{service_id}
   */
  queryQueryByService = (serviceId: string, params: RequestParams = {}) =>
    this.request<V1QueryByServiceResponse, RpcStatus>({
      path: `/sonr-io/sonr/identity/service/${serviceId}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
