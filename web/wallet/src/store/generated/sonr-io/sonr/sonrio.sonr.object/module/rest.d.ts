export declare type ObjectMsgCreateObjectResponse = object;
export declare type ObjectMsgDeactivateObjectResponse = object;
export declare type ObjectMsgReadObjectResponse = object;
export declare type ObjectMsgUpdateObjectResponse = object;
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
export declare enum ObjectObjectFieldType {
    OBJECT_FIELD_TYPE_UNSPECIFIED = "OBJECT_FIELD_TYPE_UNSPECIFIED",
    OBJECT_FIELD_TYPE_STRING = "OBJECT_FIELD_TYPE_STRING",
    OBJECT_FIELD_TYPE_NUMBER = "OBJECT_FIELD_TYPE_NUMBER",
    OBJECT_FIELD_TYPE_BOOL = "OBJECT_FIELD_TYPE_BOOL",
    OBJECT_FIELD_TYPE_ARRAY = "OBJECT_FIELD_TYPE_ARRAY",
    OBJECT_FIELD_TYPE_TIMESTAMP = "OBJECT_FIELD_TYPE_TIMESTAMP",
    OBJECT_FIELD_TYPE_GEOPOINT = "OBJECT_FIELD_TYPE_GEOPOINT",
    OBJECT_FIELD_TYPE_BLOB = "OBJECT_FIELD_TYPE_BLOB",
    OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS = "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS"
}
/**
 * Params defines the parameters for the module.
 */
export declare type ObjectParams = object;
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
 * @title object/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/sonrio/sonr/object/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<ObjectQueryParamsResponse, RpcStatus>>;
}
export {};
