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
 * Account is used for storing all credentials and their locations to be encrypted.
 */
export interface CommonAccountInfo {
  /** Address is the associated Sonr address. */
  address?: string;

  /** Credentials is a list of all credentials associated with the account. */
  network?: string;

  /** Label is the label of the account. */
  label?: string;

  /**
   * Index is the index of the account.
   * @format int64
   */
  index?: number;

  /**
   * Balance is the balance of the account.
   * @format int32
   */
  balance?: number;
}

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
 * CreateAccountResponse is the response to a CreateAccount request.
 */
export interface V1CreateAccountResponse {
  success?: boolean;
  coin_type?: string;
  did_document?: IdentityDidDocument;
  accounts?: CommonAccountInfo[];
}

/**
 * DeleteAccountResponse is the response to a DeleteAccount request.
 */
export interface V1DeleteAccountResponse {
  success?: boolean;
  did_document?: IdentityDidDocument;
  accounts?: CommonAccountInfo[];
}

/**
 * GetAccountResponse is the response to a GetAccount request.
 */
export interface V1GetAccountResponse {
  success?: boolean;
  coin_type?: string;
  accounts?: CommonAccountInfo[];
}

/**
 * ListAccountsResponse is the response to a ListAccounts request.
 */
export interface V1ListAccountsResponse {
  success?: boolean;
  accounts?: CommonAccountInfo[];
}

/**
 * LoginFinishRequest is the request to login to an account.
 */
export interface V1LoginFinishRequest {
  /** Address of the account to login to. */
  account_address?: string;

  /** The signed credential response from the user. */
  credential_response?: string;

  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin?: string;
}

/**
 * LoginFinishResponse is the response to a Login request.
 */
export interface V1LoginFinishResponse {
  /** Success is true if the account exists. */
  success?: boolean;

  /** The account address for the user. */
  account_address?: string;

  /** Relaying party id for the request. */
  rp_id?: string;

  /** Relaying party name for the request. */
  rp_name?: string;

  /** The DID Document for the wallet. */
  did_document?: IdentityDidDocument;

  /** The account info for the wallet. */
  account_info?: CommonAccountInfo;

  /**
   * The UCAN token authorization header for subsequent requests.
   * @format byte
   */
  ucan_token_header?: string;
}

/**
 * LoginStartRequest is the request to login to an account.
 */
export interface V1LoginStartRequest {
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin?: string;

  /** The Sonr account address for the user. */
  account_address?: string;
}

/**
 * LoginStartResponse is the response to a Login request.
 */
export interface V1LoginStartResponse {
  /** Success is true if the account exists. */
  success?: boolean;

  /** The account address for the user. */
  account_address?: string;

  /** Json encoded WebAuthn credential options for the user to sign with. */
  credential_options?: string;

  /** Relaying party id for the request. */
  rp_id?: string;

  /** Relaying party name for the request. */
  rp_name?: string;
}

/**
 * RefreshSharesRequest is the request to refresh the keypair.
 */
export interface V1RefreshSharesRequest {
  credential_response?: string;
  session_id?: string;
}

/**
 * RefreshSharesResponse is the response to a Refresh request.
 */
export interface V1RefreshSharesResponse {
  /** @format byte */
  id?: string;
  address?: string;
  did_document?: IdentityDidDocument;
}

/**
 * RegisterFinishRequest is the request to CreateAccount a new key from the private key.
 */
export interface V1RegisterFinishRequest {
  /** The previously generated session id. */
  uuid?: string;

  /** The signed credential response from the user. */
  credential_response?: string;

  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin?: string;
}

/**
 * RegisterFinishResponse is the response to a CreateAccount request.
 */
export interface V1RegisterFinishResponse {
  /**
   * The id of the account.
   * @format byte
   */
  id?: string;

  /** The address of the account. */
  address?: string;

  /** Relaying party id for the request. */
  rp_id?: string;

  /** Relaying party name for the request. */
  rp_name?: string;

  /** The DID Document for the wallet. */
  did_document?: IdentityDidDocument;

  /** The account info for the wallet. */
  account_info?: CommonAccountInfo;

  /**
   * The UCAN token authorization header for subsequent requests.
   * @format byte
   */
  ucan_token_header?: string;
}

/**
 * RegisterStartRequest is the request to register a new account.
 */
export interface V1RegisterStartRequest {
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin?: string;

  /** The user defined label for the device. */
  device_label?: string;

  /**
   * The security threshold for the wallet account.
   * @format int32
   */
  security_threshold?: number;

  /** The recovery passcode for the wallet account. */
  passcode?: string;

  /** The Unique Identifier for the client device. Typically in a cookie. */
  uuid?: string;
}

/**
 * RegisterStartResponse is the response to a Register request.
 */
export interface V1RegisterStartResponse {
  /** Credential options for the user to sign with WebAuthn. */
  creation_options?: string;

  /** Relaying party id for the request. */
  rp_id?: string;

  /** Relaying party name for the request. */
  rp_name?: string;
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
 * @title sonr/vault/v1/accounts.proto
 * @version version not set
 *
 * Package Motor is used for defining a Motor node and its properties.
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAccounts
   * @name VaultAccountsListAccounts
   * @summary List the accounts
   * @request GET:/sonr/vault/accounts/{sonr_id}
   */
  vaultAccountsListAccounts = (sonrId: string, params: RequestParams = {}) =>
    this.request<V1ListAccountsResponse, RpcStatus>({
      path: `/sonr/vault/accounts/${sonrId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAccounts
   * @name VaultAccountsCreateAccount
   * @summary Create a new account
   * @request POST:/sonr/vault/accounts/{sonr_id}/create
   */
  vaultAccountsCreateAccount = (sonrId: string, body: { coin_type?: string }, params: RequestParams = {}) =>
    this.request<V1CreateAccountResponse, RpcStatus>({
      path: `/sonr/vault/accounts/${sonrId}/create`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAccounts
   * @name VaultAccountsGetAccount
   * @summary Get Account
   * @request GET:/sonr/vault/accounts/{sonr_id}/{coin_type}
   */
  vaultAccountsGetAccount = (sonrId: string, coinType: string, params: RequestParams = {}) =>
    this.request<V1GetAccountResponse, RpcStatus>({
      path: `/sonr/vault/accounts/${sonrId}/${coinType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAccounts
   * @name VaultAccountsDeleteAccount
   * @summary Delete Account
   * @request POST:/sonr/vault/accounts/{target_did}/delete
   */
  vaultAccountsDeleteAccount = (targetDid: string, body: { sonr_id?: string }, params: RequestParams = {}) =>
    this.request<V1DeleteAccountResponse, RpcStatus>({
      path: `/sonr/vault/accounts/${targetDid}/delete`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAuthentication
   * @name VaultAuthenticationLoginFinish
   * @summary Login Finish
   * @request POST:/sonr/vault/auth/login/finish
   */
  vaultAuthenticationLoginFinish = (body: V1LoginFinishRequest, params: RequestParams = {}) =>
    this.request<V1LoginFinishResponse, RpcStatus>({
      path: `/sonr/vault/auth/login/finish`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAuthentication
   * @name VaultAuthenticationLoginStart
   * @summary Login Start
   * @request POST:/sonr/vault/auth/login/start
   */
  vaultAuthenticationLoginStart = (body: V1LoginStartRequest, params: RequestParams = {}) =>
    this.request<V1LoginStartResponse, RpcStatus>({
      path: `/sonr/vault/auth/login/start`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAuthentication
   * @name VaultAuthenticationRegisterFinish
   * @summary Register Finish
   * @request POST:/sonr/vault/auth/register/finish
   */
  vaultAuthenticationRegisterFinish = (body: V1RegisterFinishRequest, params: RequestParams = {}) =>
    this.request<V1RegisterFinishResponse, RpcStatus>({
      path: `/sonr/vault/auth/register/finish`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultAuthentication
   * @name VaultAuthenticationRegisterStart
   * @summary Register Start
   * @request POST:/sonr/vault/auth/register/start
   */
  vaultAuthenticationRegisterStart = (body: V1RegisterStartRequest, params: RequestParams = {}) =>
    this.request<V1RegisterStartResponse, RpcStatus>({
      path: `/sonr/vault/auth/register/start`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });

  /**
   * @description {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service. It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}". #### {{.RequestType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .RequestType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}} #### {{.ResponseType.Name}} | Name | Type | Description | | ---- | ---- | ----------- |{{range .ResponseType.Fields}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * @tags VaultStorage
   * @name VaultStorageRefreshShares
   * @summary Refresh Shares
   * @request POST:/sonr/vault/storage/refresh
   */
  vaultStorageRefreshShares = (body: V1RefreshSharesRequest, params: RequestParams = {}) =>
    this.request<V1RefreshSharesResponse, RpcStatus>({
      path: `/sonr/vault/storage/refresh`,
      method: "POST",
      body: body,
      type: ContentType.Json,
      format: "json",
      ...params,
    });
}
