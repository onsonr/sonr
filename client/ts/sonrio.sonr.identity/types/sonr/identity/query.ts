/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { DidDocument } from "./did_document";
import { Params } from "./params";

export const protobufPackage = "sonrio.sonr.identity";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDidDocumentRequest {
  did: string;
}

export interface QueryGetDidDocumentResponse {
  didDocument: DidDocument | undefined;
}

export interface QueryAllDidDocumentRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDidDocumentResponse {
  didDocument: DidDocument[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDidDocumentRequest(): QueryGetDidDocumentRequest {
  return { did: "" };
}

export const QueryGetDidDocumentRequest = {
  encode(message: QueryGetDidDocumentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidDocumentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidDocumentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDidDocumentRequest {
    return { did: isSet(object.did) ? String(object.did) : "" };
  },

  toJSON(message: QueryGetDidDocumentRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidDocumentRequest>, I>>(object: I): QueryGetDidDocumentRequest {
    const message = createBaseQueryGetDidDocumentRequest();
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseQueryGetDidDocumentResponse(): QueryGetDidDocumentResponse {
  return { didDocument: undefined };
}

export const QueryGetDidDocumentResponse = {
  encode(message: QueryGetDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDidDocumentResponse {
    return { didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined };
  },

  toJSON(message: QueryGetDidDocumentResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidDocumentResponse>, I>>(object: I): QueryGetDidDocumentResponse {
    const message = createBaseQueryGetDidDocumentResponse();
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDidDocumentRequest(): QueryAllDidDocumentRequest {
  return { pagination: undefined };
}

export const QueryAllDidDocumentRequest = {
  encode(message: QueryAllDidDocumentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDidDocumentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDidDocumentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDidDocumentRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDidDocumentRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDidDocumentRequest>, I>>(object: I): QueryAllDidDocumentRequest {
    const message = createBaseQueryAllDidDocumentRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDidDocumentResponse(): QueryAllDidDocumentResponse {
  return { didDocument: [], pagination: undefined };
}

export const QueryAllDidDocumentResponse = {
  encode(message: QueryAllDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.didDocument) {
      DidDocument.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument.push(DidDocument.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDidDocumentResponse {
    return {
      didDocument: Array.isArray(object?.didDocument)
        ? object.didDocument.map((e: any) => DidDocument.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDidDocumentResponse): unknown {
    const obj: any = {};
    if (message.didDocument) {
      obj.didDocument = message.didDocument.map((e) => e ? DidDocument.toJSON(e) : undefined);
    } else {
      obj.didDocument = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDidDocumentResponse>, I>>(object: I): QueryAllDidDocumentResponse {
    const message = createBaseQueryAllDidDocumentResponse();
    message.didDocument = object.didDocument?.map((e) => DidDocument.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a DidDocument by index. */
  DidDocument(request: QueryGetDidDocumentRequest): Promise<QueryGetDidDocumentResponse>;
  /** Queries a list of DidDocument items. */
  DidDocumentAll(request: QueryAllDidDocumentRequest): Promise<QueryAllDidDocumentResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.DidDocument = this.DidDocument.bind(this);
    this.DidDocumentAll = this.DidDocumentAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  DidDocument(request: QueryGetDidDocumentRequest): Promise<QueryGetDidDocumentResponse> {
    const data = QueryGetDidDocumentRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.Query", "DidDocument", data);
    return promise.then((data) => QueryGetDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  DidDocumentAll(request: QueryAllDidDocumentRequest): Promise<QueryAllDidDocumentResponse> {
    const data = QueryAllDidDocumentRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.Query", "DidDocumentAll", data);
    return promise.then((data) => QueryAllDidDocumentResponse.decode(new _m0.Reader(data)));
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
