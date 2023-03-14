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

export interface IdentityDidDocument {
  context?: string[];
  id?: string;

  /** optional */
  controller?: string[];

  /** optional */
  verification_method?: IdentityVerificationMethod[];

  /** optional */
  authentication?: IdentityVerificationRelationship[];

  /** optional */
  assertion_method?: IdentityVerificationRelationship[];

  /** optional */
  capability_invocation?: IdentityVerificationRelationship[];

  /** optional */
  capability_delegation?: IdentityVerificationRelationship[];

  /** optional */
  key_agreement?: IdentityVerificationRelationship[];

  /** optional */
  service?: IdentityService[];

  /** optional */
  also_known_as?: string[];

  /** optional */
  metadata?: IdentityKeyValuePair[];
}

export interface IdentityKeyValuePair {
  key?: string;
  value?: string;
}

export type IdentityMsgCreateDidDocumentResponse = object;

export type IdentityMsgDeactivateServiceResponse = object;

export type IdentityMsgDeleteDidDocumentResponse = object;

export type IdentityMsgRegisterServiceResponse = object;

export interface IdentityMsgUpdateDidDocumentResponse {
  creator?: string;
}

export type IdentityMsgUpdateServiceResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface IdentityParams {
  did_base_context?: string;
  did_method_context?: string;
  did_method_name?: string;
  did_method_version?: string;
  did_network?: string;
  ipfs_gateway?: string;
  ipfs_api?: string;
  webauthn_attestion_preference?: string;
  webauthn_authenticator_attachment?: string;

  /** @format int32 */
  webauthn_timeout?: number;
}

export interface IdentityQueryAllDidResponse {
  didDocument?: IdentityDidDocument[];

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

export interface IdentityQueryAllServiceResponse {
  services?: IdentityService[];

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

export interface IdentityQueryDidByAlsoKnownAsResponse {
  /** A DIDDocument is a JSON-LD object that contains a set of public keys */
  didDocument?: IdentityDidDocument;
}

export interface IdentityQueryDidByKeyIDResponse {
  /** A DIDDocument is a JSON-LD object that contains a set of public keys */
  didDocument?: IdentityDidDocument;
}

export type IdentityQueryDidByPubKeyResponse = object;

export interface IdentityQueryGetDidResponse {
  /** A DIDDocument is a JSON-LD object that contains a set of public keys */
  didDocument?: IdentityDidDocument;
}

export interface IdentityQueryGetServiceResponse {
  /** A Service is a JSON-LD object that contains relaying information to authenticate a client */
  service?: IdentityService;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface IdentityQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: IdentityParams;
}

export interface IdentityService {
  id?: string;
  controller?: string;
  type?: string;
  origin?: string;
  name?: string;

  /** optional */
  service_endpoints?: IdentityKeyValuePair[];

  /** optional */
  metadata?: IdentityKeyValuePair[];
}

export interface IdentityVerificationMethod {
  id?: string;
  type?: string;
  controller?: string;

  /** optional */
  public_key_jwk?: IdentityKeyValuePair[];

  /** optional */
  public_key_multibase?: string;

  /** optional */
  blockchain_account_id?: string;
  metadata?: IdentityKeyValuePair[];
}

export interface IdentityVerificationRelationship {
  verification_method?: IdentityVerificationMethod;
  reference?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
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
 * @title Sonr API
 * @version version not set
 *
 * Sonr is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, user-friendly way to manage their digital identity and assets.
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
 * No description
 * 
 * @tags Query
 * @name QueryDidAll
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/did
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
    this.request<IdentityQueryAllDidResponse, RpcStatus>({
      path: `/sonr/core/identity/did`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryDidByAlsoKnownAs
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/did/aka/{aka_id}
 */
  queryDidByAlsoKnownAs = (akaId: string, params: RequestParams = {}) =>
    this.request<IdentityQueryDidByAlsoKnownAsResponse, RpcStatus>({
      path: `/sonr/core/identity/did/aka/${akaId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryDidByKeyId
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/did/key/{key_id}
 */
  queryDidByKeyID = (keyId: string, params: RequestParams = {}) =>
    this.request<IdentityQueryDidByKeyIDResponse, RpcStatus>({
      path: `/sonr/core/identity/did/key/${keyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryDid
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/did/{did}
 */
  queryDid = (did: string, params: RequestParams = {}) =>
    this.request<IdentityQueryGetDidResponse, RpcStatus>({
      path: `/sonr/core/identity/did/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryParams
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/params
 */
  queryParams = (params: RequestParams = {}) =>
    this.request<IdentityQueryParamsResponse, RpcStatus>({
      path: `/sonr/core/identity/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryServiceAll
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/service
 */
  queryServiceAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<IdentityQueryAllServiceResponse, RpcStatus>({
      path: `/sonr/core/identity/service`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryService
 * @summary #### {{.ResponseType.Name}}
| Name | Type | Description |
| ---- | ---- | ----------- |{{range .ResponseType.Fields}}
| {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
 * @request GET:/sonr/core/identity/service/{origin}
 */
  queryService = (origin: string, params: RequestParams = {}) =>
    this.request<IdentityQueryGetServiceResponse, RpcStatus>({
      path: `/sonr/core/identity/service/${origin}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidByPubKey
   * @summary Queries a list of DidByPubKey items.
   * @request GET:/sonrhq/core/identity/did_by_pub_key
   */
  queryDidByPubKey = (params: RequestParams = {}) =>
    this.request<IdentityQueryDidByPubKeyResponse, RpcStatus>({
      path: `/sonrhq/core/identity/did_by_pub_key`,
      method: "GET",
      format: "json",
      ...params,
    });
}
