/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocument } from "../../../core/identity/did";

export const protobufPackage = "sonrhq.sonr.vault.v1";

/** RefreshSharesRequest is the request to refresh the keypair. */
export interface RefreshSharesRequest {
  credentialResponse: string;
  sessionId: string;
}

/** RefreshSharesResponse is the response to a Refresh request. */
export interface RefreshSharesResponse {
  id: Uint8Array;
  address: string;
  didDocument: DidDocument | undefined;
}

function createBaseRefreshSharesRequest(): RefreshSharesRequest {
  return { credentialResponse: "", sessionId: "" };
}

export const RefreshSharesRequest = {
  encode(message: RefreshSharesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.credentialResponse !== "") {
      writer.uint32(10).string(message.credentialResponse);
    }
    if (message.sessionId !== "") {
      writer.uint32(18).string(message.sessionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RefreshSharesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRefreshSharesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.credentialResponse = reader.string();
          break;
        case 2:
          message.sessionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RefreshSharesRequest {
    return {
      credentialResponse: isSet(object.credentialResponse) ? String(object.credentialResponse) : "",
      sessionId: isSet(object.sessionId) ? String(object.sessionId) : "",
    };
  },

  toJSON(message: RefreshSharesRequest): unknown {
    const obj: any = {};
    message.credentialResponse !== undefined && (obj.credentialResponse = message.credentialResponse);
    message.sessionId !== undefined && (obj.sessionId = message.sessionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RefreshSharesRequest>, I>>(object: I): RefreshSharesRequest {
    const message = createBaseRefreshSharesRequest();
    message.credentialResponse = object.credentialResponse ?? "";
    message.sessionId = object.sessionId ?? "";
    return message;
  },
};

function createBaseRefreshSharesResponse(): RefreshSharesResponse {
  return { id: new Uint8Array(), address: "", didDocument: undefined };
}

export const RefreshSharesResponse = {
  encode(message: RefreshSharesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id.length !== 0) {
      writer.uint32(10).bytes(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RefreshSharesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRefreshSharesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.bytes();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RefreshSharesResponse {
    return {
      id: isSet(object.id) ? bytesFromBase64(object.id) : new Uint8Array(),
      address: isSet(object.address) ? String(object.address) : "",
      didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined,
    };
  },

  toJSON(message: RefreshSharesResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = base64FromBytes(message.id !== undefined ? message.id : new Uint8Array()));
    message.address !== undefined && (obj.address = message.address);
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RefreshSharesResponse>, I>>(object: I): RefreshSharesResponse {
    const message = createBaseRefreshSharesResponse();
    message.id = object.id ?? new Uint8Array();
    message.address = object.address ?? "";
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

/** Vault is the service used for managing a node's keypair. */
export interface VaultStorage {
  /**
   * Refresh Shares
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
  RefreshShares(request: RefreshSharesRequest): Promise<RefreshSharesResponse>;
}

export class VaultStorageClientImpl implements VaultStorage {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RefreshShares = this.RefreshShares.bind(this);
  }
  RefreshShares(request: RefreshSharesRequest): Promise<RefreshSharesResponse> {
    const data = RefreshSharesRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultStorage", "RefreshShares", data);
    return promise.then((data) => RefreshSharesResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
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
