/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { DidDocument } from "./did";
import { Params } from "./params";

export const protobufPackage = "sonrio.sonr.identity.v1";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDidRequest {
  did: string;
}

export interface QueryGetDidResponse {
  didDocument: DidDocument | undefined;
}

export interface QueryByServiceRequest {
  serviceId: string;
}

export interface QueryByServiceResponse {
  didDocument: DidDocument | undefined;
}

export interface QueryByKeyIDRequest {
  keyId: string;
}

export interface QueryByKeyIDResponse {
  didDocument: DidDocument | undefined;
}

export interface QueryByAlsoKnownAsRequest {
  akaId: string;
}

export interface QueryByAlsoKnownAsResponse {
  didDocument: DidDocument | undefined;
}

export interface QueryByMethodRequest {
  methodId: string;
  pagination: PageRequest | undefined;
}

export interface QueryByMethodResponse {
  didDocument: DidDocument[];
  pagination: PageResponse | undefined;
}

export interface QueryAllDidRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDidResponse {
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

function createBaseQueryGetDidRequest(): QueryGetDidRequest {
  return { did: "" };
}

export const QueryGetDidRequest = {
  encode(message: QueryGetDidRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidRequest();
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

  fromJSON(object: any): QueryGetDidRequest {
    return { did: isSet(object.did) ? String(object.did) : "" };
  },

  toJSON(message: QueryGetDidRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidRequest>, I>>(object: I): QueryGetDidRequest {
    const message = createBaseQueryGetDidRequest();
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseQueryGetDidResponse(): QueryGetDidResponse {
  return { didDocument: undefined };
}

export const QueryGetDidResponse = {
  encode(message: QueryGetDidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidResponse();
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

  fromJSON(object: any): QueryGetDidResponse {
    return { didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined };
  },

  toJSON(message: QueryGetDidResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidResponse>, I>>(object: I): QueryGetDidResponse {
    const message = createBaseQueryGetDidResponse();
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

function createBaseQueryByServiceRequest(): QueryByServiceRequest {
  return { serviceId: "" };
}

export const QueryByServiceRequest = {
  encode(message: QueryByServiceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.serviceId !== "") {
      writer.uint32(10).string(message.serviceId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByServiceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByServiceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serviceId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByServiceRequest {
    return { serviceId: isSet(object.serviceId) ? String(object.serviceId) : "" };
  },

  toJSON(message: QueryByServiceRequest): unknown {
    const obj: any = {};
    message.serviceId !== undefined && (obj.serviceId = message.serviceId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByServiceRequest>, I>>(object: I): QueryByServiceRequest {
    const message = createBaseQueryByServiceRequest();
    message.serviceId = object.serviceId ?? "";
    return message;
  },
};

function createBaseQueryByServiceResponse(): QueryByServiceResponse {
  return { didDocument: undefined };
}

export const QueryByServiceResponse = {
  encode(message: QueryByServiceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByServiceResponse();
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

  fromJSON(object: any): QueryByServiceResponse {
    return { didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined };
  },

  toJSON(message: QueryByServiceResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByServiceResponse>, I>>(object: I): QueryByServiceResponse {
    const message = createBaseQueryByServiceResponse();
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

function createBaseQueryByKeyIDRequest(): QueryByKeyIDRequest {
  return { keyId: "" };
}

export const QueryByKeyIDRequest = {
  encode(message: QueryByKeyIDRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.keyId !== "") {
      writer.uint32(10).string(message.keyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByKeyIDRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByKeyIDRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByKeyIDRequest {
    return { keyId: isSet(object.keyId) ? String(object.keyId) : "" };
  },

  toJSON(message: QueryByKeyIDRequest): unknown {
    const obj: any = {};
    message.keyId !== undefined && (obj.keyId = message.keyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByKeyIDRequest>, I>>(object: I): QueryByKeyIDRequest {
    const message = createBaseQueryByKeyIDRequest();
    message.keyId = object.keyId ?? "";
    return message;
  },
};

function createBaseQueryByKeyIDResponse(): QueryByKeyIDResponse {
  return { didDocument: undefined };
}

export const QueryByKeyIDResponse = {
  encode(message: QueryByKeyIDResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByKeyIDResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByKeyIDResponse();
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

  fromJSON(object: any): QueryByKeyIDResponse {
    return { didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined };
  },

  toJSON(message: QueryByKeyIDResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByKeyIDResponse>, I>>(object: I): QueryByKeyIDResponse {
    const message = createBaseQueryByKeyIDResponse();
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

function createBaseQueryByAlsoKnownAsRequest(): QueryByAlsoKnownAsRequest {
  return { akaId: "" };
}

export const QueryByAlsoKnownAsRequest = {
  encode(message: QueryByAlsoKnownAsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.akaId !== "") {
      writer.uint32(10).string(message.akaId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByAlsoKnownAsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByAlsoKnownAsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.akaId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByAlsoKnownAsRequest {
    return { akaId: isSet(object.akaId) ? String(object.akaId) : "" };
  },

  toJSON(message: QueryByAlsoKnownAsRequest): unknown {
    const obj: any = {};
    message.akaId !== undefined && (obj.akaId = message.akaId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByAlsoKnownAsRequest>, I>>(object: I): QueryByAlsoKnownAsRequest {
    const message = createBaseQueryByAlsoKnownAsRequest();
    message.akaId = object.akaId ?? "";
    return message;
  },
};

function createBaseQueryByAlsoKnownAsResponse(): QueryByAlsoKnownAsResponse {
  return { didDocument: undefined };
}

export const QueryByAlsoKnownAsResponse = {
  encode(message: QueryByAlsoKnownAsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByAlsoKnownAsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByAlsoKnownAsResponse();
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

  fromJSON(object: any): QueryByAlsoKnownAsResponse {
    return { didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined };
  },

  toJSON(message: QueryByAlsoKnownAsResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByAlsoKnownAsResponse>, I>>(object: I): QueryByAlsoKnownAsResponse {
    const message = createBaseQueryByAlsoKnownAsResponse();
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    return message;
  },
};

function createBaseQueryByMethodRequest(): QueryByMethodRequest {
  return { methodId: "", pagination: undefined };
}

export const QueryByMethodRequest = {
  encode(message: QueryByMethodRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.methodId !== "") {
      writer.uint32(10).string(message.methodId);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByMethodRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByMethodRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.methodId = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByMethodRequest {
    return {
      methodId: isSet(object.methodId) ? String(object.methodId) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryByMethodRequest): unknown {
    const obj: any = {};
    message.methodId !== undefined && (obj.methodId = message.methodId);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByMethodRequest>, I>>(object: I): QueryByMethodRequest {
    const message = createBaseQueryByMethodRequest();
    message.methodId = object.methodId ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryByMethodResponse(): QueryByMethodResponse {
  return { didDocument: [], pagination: undefined };
}

export const QueryByMethodResponse = {
  encode(message: QueryByMethodResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.didDocument) {
      DidDocument.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByMethodResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByMethodResponse();
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

  fromJSON(object: any): QueryByMethodResponse {
    return {
      didDocument: Array.isArray(object?.didDocument)
        ? object.didDocument.map((e: any) => DidDocument.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryByMethodResponse): unknown {
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

  fromPartial<I extends Exact<DeepPartial<QueryByMethodResponse>, I>>(object: I): QueryByMethodResponse {
    const message = createBaseQueryByMethodResponse();
    message.didDocument = object.didDocument?.map((e) => DidDocument.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDidRequest(): QueryAllDidRequest {
  return { pagination: undefined };
}

export const QueryAllDidRequest = {
  encode(message: QueryAllDidRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDidRequest();
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

  fromJSON(object: any): QueryAllDidRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDidRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDidRequest>, I>>(object: I): QueryAllDidRequest {
    const message = createBaseQueryAllDidRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDidResponse(): QueryAllDidResponse {
  return { didDocument: [], pagination: undefined };
}

export const QueryAllDidResponse = {
  encode(message: QueryAllDidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.didDocument) {
      DidDocument.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDidResponse();
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

  fromJSON(object: any): QueryAllDidResponse {
    return {
      didDocument: Array.isArray(object?.didDocument)
        ? object.didDocument.map((e: any) => DidDocument.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDidResponse): unknown {
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

  fromPartial<I extends Exact<DeepPartial<QueryAllDidResponse>, I>>(object: I): QueryAllDidResponse {
    const message = createBaseQueryAllDidResponse();
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
  Did(request: QueryGetDidRequest): Promise<QueryGetDidResponse>;
  /** Queries a list of DidDocument items. */
  DidAll(request: QueryAllDidRequest): Promise<QueryAllDidResponse>;
  /** Queries a DIDDocument for the matching service */
  QueryByService(request: QueryByServiceRequest): Promise<QueryByServiceResponse>;
  /** Queries a DIDDocument for the matching key */
  QueryByKeyID(request: QueryByKeyIDRequest): Promise<QueryByKeyIDResponse>;
  /** Queries a DIDDocument for the matching AlsoKnownAs */
  QueryByAlsoKnownAs(request: QueryByAlsoKnownAsRequest): Promise<QueryByAlsoKnownAsResponse>;
  /** Queries a list of DIDDocument for the matching method */
  QueryByMethod(request: QueryByMethodRequest): Promise<QueryByMethodResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Did = this.Did.bind(this);
    this.DidAll = this.DidAll.bind(this);
    this.QueryByService = this.QueryByService.bind(this);
    this.QueryByKeyID = this.QueryByKeyID.bind(this);
    this.QueryByAlsoKnownAs = this.QueryByAlsoKnownAs.bind(this);
    this.QueryByMethod = this.QueryByMethod.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Did(request: QueryGetDidRequest): Promise<QueryGetDidResponse> {
    const data = QueryGetDidRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "Did", data);
    return promise.then((data) => QueryGetDidResponse.decode(new _m0.Reader(data)));
  }

  DidAll(request: QueryAllDidRequest): Promise<QueryAllDidResponse> {
    const data = QueryAllDidRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "DidAll", data);
    return promise.then((data) => QueryAllDidResponse.decode(new _m0.Reader(data)));
  }

  QueryByService(request: QueryByServiceRequest): Promise<QueryByServiceResponse> {
    const data = QueryByServiceRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "QueryByService", data);
    return promise.then((data) => QueryByServiceResponse.decode(new _m0.Reader(data)));
  }

  QueryByKeyID(request: QueryByKeyIDRequest): Promise<QueryByKeyIDResponse> {
    const data = QueryByKeyIDRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "QueryByKeyID", data);
    return promise.then((data) => QueryByKeyIDResponse.decode(new _m0.Reader(data)));
  }

  QueryByAlsoKnownAs(request: QueryByAlsoKnownAsRequest): Promise<QueryByAlsoKnownAsResponse> {
    const data = QueryByAlsoKnownAsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "QueryByAlsoKnownAs", data);
    return promise.then((data) => QueryByAlsoKnownAsResponse.decode(new _m0.Reader(data)));
  }

  QueryByMethod(request: QueryByMethodRequest): Promise<QueryByMethodResponse> {
    const data = QueryByMethodRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Query", "QueryByMethod", data);
    return promise.then((data) => QueryByMethodResponse.decode(new _m0.Reader(data)));
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
