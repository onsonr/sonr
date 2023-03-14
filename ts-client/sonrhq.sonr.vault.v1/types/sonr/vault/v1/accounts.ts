/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocument } from "../../../core/identity/did";
import { AccountInfo } from "../../common/info";

export const protobufPackage = "sonrhq.sonr.vault.v1";

/** CreateAccountRequest is the request to create an account. */
export interface CreateAccountRequest {
  sonrId: string;
  coinType: string;
}

/** CreateAccountResponse is the response to a CreateAccount request. */
export interface CreateAccountResponse {
  success: boolean;
  coinType: string;
  didDocument: DidDocument | undefined;
  accounts: AccountInfo[];
}

/** GetAccountRequest is the request to get an account. */
export interface GetAccountRequest {
  sonrId: string;
  coinType: string;
}

/** GetAccountResponse is the response to a GetAccount request. */
export interface GetAccountResponse {
  success: boolean;
  coinType: string;
  accounts: AccountInfo[];
}

/** ListAccountsRequest is the request to list the accounts. */
export interface ListAccountsRequest {
  sonrId: string;
}

/** ListAccountsResponse is the response to a ListAccounts request. */
export interface ListAccountsResponse {
  success: boolean;
  accounts: AccountInfo[];
}

/** DeleteAccountRequest is the request to delete an account. */
export interface DeleteAccountRequest {
  sonrId: string;
  targetDid: string;
}

/** DeleteAccountResponse is the response to a DeleteAccount request. */
export interface DeleteAccountResponse {
  success: boolean;
  didDocument: DidDocument | undefined;
  accounts: AccountInfo[];
}

function createBaseCreateAccountRequest(): CreateAccountRequest {
  return { sonrId: "", coinType: "" };
}

export const CreateAccountRequest = {
  encode(message: CreateAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sonrId !== "") {
      writer.uint32(10).string(message.sonrId);
    }
    if (message.coinType !== "") {
      writer.uint32(18).string(message.coinType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sonrId = reader.string();
          break;
        case 2:
          message.coinType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateAccountRequest {
    return {
      sonrId: isSet(object.sonrId) ? String(object.sonrId) : "",
      coinType: isSet(object.coinType) ? String(object.coinType) : "",
    };
  },

  toJSON(message: CreateAccountRequest): unknown {
    const obj: any = {};
    message.sonrId !== undefined && (obj.sonrId = message.sonrId);
    message.coinType !== undefined && (obj.coinType = message.coinType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateAccountRequest>, I>>(object: I): CreateAccountRequest {
    const message = createBaseCreateAccountRequest();
    message.sonrId = object.sonrId ?? "";
    message.coinType = object.coinType ?? "";
    return message;
  },
};

function createBaseCreateAccountResponse(): CreateAccountResponse {
  return { success: false, coinType: "", didDocument: undefined, accounts: [] };
}

export const CreateAccountResponse = {
  encode(message: CreateAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.coinType !== "") {
      writer.uint32(18).string(message.coinType);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.accounts) {
      AccountInfo.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.coinType = reader.string();
          break;
        case 3:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 4:
          message.accounts.push(AccountInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateAccountResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      coinType: isSet(object.coinType) ? String(object.coinType) : "",
      didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined,
      accounts: Array.isArray(object?.accounts) ? object.accounts.map((e: any) => AccountInfo.fromJSON(e)) : [],
    };
  },

  toJSON(message: CreateAccountResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.coinType !== undefined && (obj.coinType = message.coinType);
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    if (message.accounts) {
      obj.accounts = message.accounts.map((e) => e ? AccountInfo.toJSON(e) : undefined);
    } else {
      obj.accounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateAccountResponse>, I>>(object: I): CreateAccountResponse {
    const message = createBaseCreateAccountResponse();
    message.success = object.success ?? false;
    message.coinType = object.coinType ?? "";
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    message.accounts = object.accounts?.map((e) => AccountInfo.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetAccountRequest(): GetAccountRequest {
  return { sonrId: "", coinType: "" };
}

export const GetAccountRequest = {
  encode(message: GetAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sonrId !== "") {
      writer.uint32(10).string(message.sonrId);
    }
    if (message.coinType !== "") {
      writer.uint32(18).string(message.coinType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sonrId = reader.string();
          break;
        case 2:
          message.coinType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAccountRequest {
    return {
      sonrId: isSet(object.sonrId) ? String(object.sonrId) : "",
      coinType: isSet(object.coinType) ? String(object.coinType) : "",
    };
  },

  toJSON(message: GetAccountRequest): unknown {
    const obj: any = {};
    message.sonrId !== undefined && (obj.sonrId = message.sonrId);
    message.coinType !== undefined && (obj.coinType = message.coinType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAccountRequest>, I>>(object: I): GetAccountRequest {
    const message = createBaseGetAccountRequest();
    message.sonrId = object.sonrId ?? "";
    message.coinType = object.coinType ?? "";
    return message;
  },
};

function createBaseGetAccountResponse(): GetAccountResponse {
  return { success: false, coinType: "", accounts: [] };
}

export const GetAccountResponse = {
  encode(message: GetAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.coinType !== "") {
      writer.uint32(18).string(message.coinType);
    }
    for (const v of message.accounts) {
      AccountInfo.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.coinType = reader.string();
          break;
        case 4:
          message.accounts.push(AccountInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAccountResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      coinType: isSet(object.coinType) ? String(object.coinType) : "",
      accounts: Array.isArray(object?.accounts) ? object.accounts.map((e: any) => AccountInfo.fromJSON(e)) : [],
    };
  },

  toJSON(message: GetAccountResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.coinType !== undefined && (obj.coinType = message.coinType);
    if (message.accounts) {
      obj.accounts = message.accounts.map((e) => e ? AccountInfo.toJSON(e) : undefined);
    } else {
      obj.accounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAccountResponse>, I>>(object: I): GetAccountResponse {
    const message = createBaseGetAccountResponse();
    message.success = object.success ?? false;
    message.coinType = object.coinType ?? "";
    message.accounts = object.accounts?.map((e) => AccountInfo.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListAccountsRequest(): ListAccountsRequest {
  return { sonrId: "" };
}

export const ListAccountsRequest = {
  encode(message: ListAccountsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sonrId !== "") {
      writer.uint32(10).string(message.sonrId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAccountsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAccountsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sonrId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListAccountsRequest {
    return { sonrId: isSet(object.sonrId) ? String(object.sonrId) : "" };
  },

  toJSON(message: ListAccountsRequest): unknown {
    const obj: any = {};
    message.sonrId !== undefined && (obj.sonrId = message.sonrId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAccountsRequest>, I>>(object: I): ListAccountsRequest {
    const message = createBaseListAccountsRequest();
    message.sonrId = object.sonrId ?? "";
    return message;
  },
};

function createBaseListAccountsResponse(): ListAccountsResponse {
  return { success: false, accounts: [] };
}

export const ListAccountsResponse = {
  encode(message: ListAccountsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    for (const v of message.accounts) {
      AccountInfo.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAccountsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAccountsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 3:
          message.accounts.push(AccountInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListAccountsResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      accounts: Array.isArray(object?.accounts) ? object.accounts.map((e: any) => AccountInfo.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListAccountsResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    if (message.accounts) {
      obj.accounts = message.accounts.map((e) => e ? AccountInfo.toJSON(e) : undefined);
    } else {
      obj.accounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAccountsResponse>, I>>(object: I): ListAccountsResponse {
    const message = createBaseListAccountsResponse();
    message.success = object.success ?? false;
    message.accounts = object.accounts?.map((e) => AccountInfo.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteAccountRequest(): DeleteAccountRequest {
  return { sonrId: "", targetDid: "" };
}

export const DeleteAccountRequest = {
  encode(message: DeleteAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sonrId !== "") {
      writer.uint32(10).string(message.sonrId);
    }
    if (message.targetDid !== "") {
      writer.uint32(18).string(message.targetDid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sonrId = reader.string();
          break;
        case 2:
          message.targetDid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteAccountRequest {
    return {
      sonrId: isSet(object.sonrId) ? String(object.sonrId) : "",
      targetDid: isSet(object.targetDid) ? String(object.targetDid) : "",
    };
  },

  toJSON(message: DeleteAccountRequest): unknown {
    const obj: any = {};
    message.sonrId !== undefined && (obj.sonrId = message.sonrId);
    message.targetDid !== undefined && (obj.targetDid = message.targetDid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteAccountRequest>, I>>(object: I): DeleteAccountRequest {
    const message = createBaseDeleteAccountRequest();
    message.sonrId = object.sonrId ?? "";
    message.targetDid = object.targetDid ?? "";
    return message;
  },
};

function createBaseDeleteAccountResponse(): DeleteAccountResponse {
  return { success: false, didDocument: undefined, accounts: [] };
}

export const DeleteAccountResponse = {
  encode(message: DeleteAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.accounts) {
      AccountInfo.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 3:
          message.accounts.push(AccountInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteAccountResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined,
      accounts: Array.isArray(object?.accounts) ? object.accounts.map((e: any) => AccountInfo.fromJSON(e)) : [],
    };
  },

  toJSON(message: DeleteAccountResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    if (message.accounts) {
      obj.accounts = message.accounts.map((e) => e ? AccountInfo.toJSON(e) : undefined);
    } else {
      obj.accounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteAccountResponse>, I>>(object: I): DeleteAccountResponse {
    const message = createBaseDeleteAccountResponse();
    message.success = object.success ?? false;
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    message.accounts = object.accounts?.map((e) => AccountInfo.fromPartial(e)) || [];
    return message;
  },
};

/** Vault is the service used for managing a node's keypair. */
export interface VaultAccounts {
  /**
   * Create a new account
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  CreateAccount(request: CreateAccountRequest): Promise<CreateAccountResponse>;
  /**
   * List the accounts
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  ListAccounts(request: ListAccountsRequest): Promise<ListAccountsResponse>;
  /**
   * Get Account
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  GetAccount(request: GetAccountRequest): Promise<GetAccountResponse>;
  /**
   * Delete Account
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  DeleteAccount(request: DeleteAccountRequest): Promise<DeleteAccountResponse>;
}

export class VaultAccountsClientImpl implements VaultAccounts {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAccount = this.CreateAccount.bind(this);
    this.ListAccounts = this.ListAccounts.bind(this);
    this.GetAccount = this.GetAccount.bind(this);
    this.DeleteAccount = this.DeleteAccount.bind(this);
  }
  CreateAccount(request: CreateAccountRequest): Promise<CreateAccountResponse> {
    const data = CreateAccountRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAccounts", "CreateAccount", data);
    return promise.then((data) => CreateAccountResponse.decode(new _m0.Reader(data)));
  }

  ListAccounts(request: ListAccountsRequest): Promise<ListAccountsResponse> {
    const data = ListAccountsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAccounts", "ListAccounts", data);
    return promise.then((data) => ListAccountsResponse.decode(new _m0.Reader(data)));
  }

  GetAccount(request: GetAccountRequest): Promise<GetAccountResponse> {
    const data = GetAccountRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAccounts", "GetAccount", data);
    return promise.then((data) => GetAccountResponse.decode(new _m0.Reader(data)));
  }

  DeleteAccount(request: DeleteAccountRequest): Promise<DeleteAccountResponse> {
    const data = DeleteAccountRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAccounts", "DeleteAccount", data);
    return promise.then((data) => DeleteAccountResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
