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

export type ObjectMsgCreateObjectResponse = object;

export type ObjectMsgDeactivateObjectResponse = object;

export type ObjectMsgReadObjectResponse = object;

export type ObjectMsgUpdateObjectResponse = object;

/**
 * ObjectField is a field of an Object.
 */
export interface ObjectObjectField {
  label?: string;

  /**
   * - OBJECT_FIELD_TYPE_UNSPECIFIED: ObjectFieldTypeUnspecified is the default value
   *  - OBJECT_FIELD_TYPE_STRING: ObjectFieldTypeString is a string or text field
   *  - OBJECT_FIELD_TYPE_NUMBER: ObjectFieldTypeInt is an integer
   *  - OBJECT_FIELD_TYPE_BOOL: ObjectFieldTypeBool is a boolean
   *  - OBJECT_FIELD_TYPE_ARRAY: ObjectFieldTypeArray is a list of values
   *  - OBJECT_FIELD_TYPE_TIMESTAMP: ObjectFieldTypeDateTime is a datetime
   *  - OBJECT_FIELD_TYPE_GEOPOINT: ObjectFieldTypeGeopoint is a geopoint
   *  - OBJECT_FIELD_TYPE_BLOB: ObjectFieldTypeBlob is a blob of data
   *  - OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS: ObjectFieldTypeETU is a pointer to an Ethereum account address.
   */
  type?: ObjectObjectFieldType;

  /** Did is the identifier of the field. */
  did?: string;

  /** ObjectFieldText is a text field of an Object. */
  stringValue?: ObjectObjectFieldText;

  /** ObjectFieldNumber is a number field of an Object. */
  numberValue?: ObjectObjectFieldNumber;

  /** ObjectFieldBool is a boolean field of an Object. */
  boolValue?: ObjectObjectFieldBool;

  /** ObjectFieldArray is an array of ObjectFields to be stored in the graph object. */
  arrayValue?: ObjectObjectFieldArray;

  /** Time is defined by milliseconds since epoch. */
  timeStampValue?: ObjectObjectFieldTime;

  /** ObjectFieldGeopoint is a field of an Object for geopoints. */
  geopointValue?: ObjectObjectFieldGeopoint;

  /** ObjectFieldBlob is a field of an Object for blobs. */
  blobValue?: ObjectObjectFieldBlob;

  /** ObjectFieldBlockchainAddress is a field of an Object for blockchain addresses. */
  blockchainAddrValue?: ObjectObjectFieldBlockchainAddress;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldArray is an array of ObjectFields to be stored in the graph object.
 */
export interface ObjectObjectFieldArray {
  label?: string;

  /**
   * - OBJECT_FIELD_TYPE_UNSPECIFIED: ObjectFieldTypeUnspecified is the default value
   *  - OBJECT_FIELD_TYPE_STRING: ObjectFieldTypeString is a string or text field
   *  - OBJECT_FIELD_TYPE_NUMBER: ObjectFieldTypeInt is an integer
   *  - OBJECT_FIELD_TYPE_BOOL: ObjectFieldTypeBool is a boolean
   *  - OBJECT_FIELD_TYPE_ARRAY: ObjectFieldTypeArray is a list of values
   *  - OBJECT_FIELD_TYPE_TIMESTAMP: ObjectFieldTypeDateTime is a datetime
   *  - OBJECT_FIELD_TYPE_GEOPOINT: ObjectFieldTypeGeopoint is a geopoint
   *  - OBJECT_FIELD_TYPE_BLOB: ObjectFieldTypeBlob is a blob of data
   *  - OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS: ObjectFieldTypeETU is a pointer to an Ethereum account address.
   */
  type?: ObjectObjectFieldType;

  /** Did is the identifier of the field. */
  did?: string;
  items?: ObjectObjectField[];
}

/**
 * ObjectFieldBlob is a field of an Object for blobs.
 */
export interface ObjectObjectFieldBlob {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;

  /** @format byte */
  value?: string;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldBlockchainAddress is a field of an Object for blockchain addresses.
 */
export interface ObjectObjectFieldBlockchainAddress {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;
  value?: string;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldBool is a boolean field of an Object.
 */
export interface ObjectObjectFieldBool {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;
  value?: boolean;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldGeopoint is a field of an Object for geopoints.
 */
export interface ObjectObjectFieldGeopoint {
  label?: string;

  /**
   * - OBJECT_FIELD_TYPE_UNSPECIFIED: ObjectFieldTypeUnspecified is the default value
   *  - OBJECT_FIELD_TYPE_STRING: ObjectFieldTypeString is a string or text field
   *  - OBJECT_FIELD_TYPE_NUMBER: ObjectFieldTypeInt is an integer
   *  - OBJECT_FIELD_TYPE_BOOL: ObjectFieldTypeBool is a boolean
   *  - OBJECT_FIELD_TYPE_ARRAY: ObjectFieldTypeArray is a list of values
   *  - OBJECT_FIELD_TYPE_TIMESTAMP: ObjectFieldTypeDateTime is a datetime
   *  - OBJECT_FIELD_TYPE_GEOPOINT: ObjectFieldTypeGeopoint is a geopoint
   *  - OBJECT_FIELD_TYPE_BLOB: ObjectFieldTypeBlob is a blob of data
   *  - OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS: ObjectFieldTypeETU is a pointer to an Ethereum account address.
   */
  type?: ObjectObjectFieldType;

  /** Did is the identifier of the field. */
  did?: string;

  /**
   * Latitude is the geo-latitude of the point.
   * @format double
   */
  latitude?: number;

  /** @format double */
  longitude?: number;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldNumber is a number field of an Object.
 */
export interface ObjectObjectFieldNumber {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;

  /** @format double */
  value?: number;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldText is a text field of an Object.
 */
export interface ObjectObjectFieldText {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;
  value?: string;
  metadata?: Record<string, string>;
}

/**
 * ObjectFieldTime is a time field of an Object.
 */
export interface ObjectObjectFieldTime {
  label?: string;

  /** Did is the identifier of the field. */
  did?: string;

  /** @format int64 */
  value?: string;
  metadata?: Record<string, string>;
}

/**
* - OBJECT_FIELD_TYPE_UNSPECIFIED: ObjectFieldTypeUnspecified is the default value
 - OBJECT_FIELD_TYPE_STRING: ObjectFieldTypeString is a string or text field
 - OBJECT_FIELD_TYPE_NUMBER: ObjectFieldTypeInt is an integer
 - OBJECT_FIELD_TYPE_BOOL: ObjectFieldTypeBool is a boolean
 - OBJECT_FIELD_TYPE_ARRAY: ObjectFieldTypeArray is a list of values
 - OBJECT_FIELD_TYPE_TIMESTAMP: ObjectFieldTypeDateTime is a datetime
 - OBJECT_FIELD_TYPE_GEOPOINT: ObjectFieldTypeGeopoint is a geopoint
 - OBJECT_FIELD_TYPE_BLOB: ObjectFieldTypeBlob is a blob of data
 - OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS: ObjectFieldTypeETU is a pointer to an Ethereum account address.
*/
export enum ObjectObjectFieldType {
  OBJECT_FIELD_TYPE_UNSPECIFIED = "OBJECT_FIELD_TYPE_UNSPECIFIED",
  OBJECT_FIELD_TYPE_STRING = "OBJECT_FIELD_TYPE_STRING",
  OBJECT_FIELD_TYPE_NUMBER = "OBJECT_FIELD_TYPE_NUMBER",
  OBJECT_FIELD_TYPE_BOOL = "OBJECT_FIELD_TYPE_BOOL",
  OBJECT_FIELD_TYPE_ARRAY = "OBJECT_FIELD_TYPE_ARRAY",
  OBJECT_FIELD_TYPE_TIMESTAMP = "OBJECT_FIELD_TYPE_TIMESTAMP",
  OBJECT_FIELD_TYPE_GEOPOINT = "OBJECT_FIELD_TYPE_GEOPOINT",
  OBJECT_FIELD_TYPE_BLOB = "OBJECT_FIELD_TYPE_BLOB",
  OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS = "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS",
}

/**
 * Params defines the parameters for the module.
 */
export type ObjectParams = object;

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface ObjectQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: ObjectParams;
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
 * @title object/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/sonrio/sonr/object/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<ObjectQueryParamsResponse, RpcStatus>({
      path: `/sonrio/sonr/object/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
