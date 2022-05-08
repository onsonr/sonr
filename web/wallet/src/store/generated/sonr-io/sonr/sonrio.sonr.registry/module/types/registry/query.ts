/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../registry/params";
import { WhoIs } from "../registry/who_is";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sonrio.sonr.registry";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryWhoIsRequest {
  did: string;
}

export interface QueryWhoIsResponse {
  WhoIs: WhoIs | undefined;
}

export interface QueryAllWhoIsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllWhoIsResponse {
  WhoIs: WhoIs[];
  pagination: PageResponse | undefined;
}

export interface QueryWhoIsAliasRequest {
  alias: string;
}

export interface QueryWhoIsAliasResponse {}

export interface QueryWhoIsControllerRequest {
  controller: string;
}

export interface QueryWhoIsControllerResponse {}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryWhoIsRequest: object = { did: "" };

export const QueryWhoIsRequest = {
  encode(message: QueryWhoIsRequest, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhoIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhoIsRequest } as QueryWhoIsRequest;
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

  fromJSON(object: any): QueryWhoIsRequest {
    const message = { ...baseQueryWhoIsRequest } as QueryWhoIsRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: QueryWhoIsRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhoIsRequest>): QueryWhoIsRequest {
    const message = { ...baseQueryWhoIsRequest } as QueryWhoIsRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseQueryWhoIsResponse: object = {};

export const QueryWhoIsResponse = {
  encode(
    message: QueryWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.WhoIs !== undefined) {
      WhoIs.encode(message.WhoIs, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhoIsResponse } as QueryWhoIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.WhoIs = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhoIsResponse {
    const message = { ...baseQueryWhoIsResponse } as QueryWhoIsResponse;
    if (object.WhoIs !== undefined && object.WhoIs !== null) {
      message.WhoIs = WhoIs.fromJSON(object.WhoIs);
    } else {
      message.WhoIs = undefined;
    }
    return message;
  },

  toJSON(message: QueryWhoIsResponse): unknown {
    const obj: any = {};
    message.WhoIs !== undefined &&
      (obj.WhoIs = message.WhoIs ? WhoIs.toJSON(message.WhoIs) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhoIsResponse>): QueryWhoIsResponse {
    const message = { ...baseQueryWhoIsResponse } as QueryWhoIsResponse;
    if (object.WhoIs !== undefined && object.WhoIs !== null) {
      message.WhoIs = WhoIs.fromPartial(object.WhoIs);
    } else {
      message.WhoIs = undefined;
    }
    return message;
  },
};

const baseQueryAllWhoIsRequest: object = {};

export const QueryAllWhoIsRequest = {
  encode(
    message: QueryAllWhoIsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhoIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhoIsRequest } as QueryAllWhoIsRequest;
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

  fromJSON(object: any): QueryAllWhoIsRequest {
    const message = { ...baseQueryAllWhoIsRequest } as QueryAllWhoIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhoIsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllWhoIsRequest>): QueryAllWhoIsRequest {
    const message = { ...baseQueryAllWhoIsRequest } as QueryAllWhoIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllWhoIsResponse: object = {};

export const QueryAllWhoIsResponse = {
  encode(
    message: QueryAllWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.WhoIs) {
      WhoIs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhoIsResponse } as QueryAllWhoIsResponse;
    message.WhoIs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.WhoIs.push(WhoIs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllWhoIsResponse {
    const message = { ...baseQueryAllWhoIsResponse } as QueryAllWhoIsResponse;
    message.WhoIs = [];
    if (object.WhoIs !== undefined && object.WhoIs !== null) {
      for (const e of object.WhoIs) {
        message.WhoIs.push(WhoIs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhoIsResponse): unknown {
    const obj: any = {};
    if (message.WhoIs) {
      obj.WhoIs = message.WhoIs.map((e) => (e ? WhoIs.toJSON(e) : undefined));
    } else {
      obj.WhoIs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWhoIsResponse>
  ): QueryAllWhoIsResponse {
    const message = { ...baseQueryAllWhoIsResponse } as QueryAllWhoIsResponse;
    message.WhoIs = [];
    if (object.WhoIs !== undefined && object.WhoIs !== null) {
      for (const e of object.WhoIs) {
        message.WhoIs.push(WhoIs.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryWhoIsAliasRequest: object = { alias: "" };

export const QueryWhoIsAliasRequest = {
  encode(
    message: QueryWhoIsAliasRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhoIsAliasRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhoIsAliasRequest } as QueryWhoIsAliasRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.alias = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhoIsAliasRequest {
    const message = { ...baseQueryWhoIsAliasRequest } as QueryWhoIsAliasRequest;
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    return message;
  },

  toJSON(message: QueryWhoIsAliasRequest): unknown {
    const obj: any = {};
    message.alias !== undefined && (obj.alias = message.alias);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryWhoIsAliasRequest>
  ): QueryWhoIsAliasRequest {
    const message = { ...baseQueryWhoIsAliasRequest } as QueryWhoIsAliasRequest;
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    return message;
  },
};

const baseQueryWhoIsAliasResponse: object = {};

export const QueryWhoIsAliasResponse = {
  encode(_: QueryWhoIsAliasResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhoIsAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryWhoIsAliasResponse,
    } as QueryWhoIsAliasResponse;
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

  fromJSON(_: any): QueryWhoIsAliasResponse {
    const message = {
      ...baseQueryWhoIsAliasResponse,
    } as QueryWhoIsAliasResponse;
    return message;
  },

  toJSON(_: QueryWhoIsAliasResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryWhoIsAliasResponse>
  ): QueryWhoIsAliasResponse {
    const message = {
      ...baseQueryWhoIsAliasResponse,
    } as QueryWhoIsAliasResponse;
    return message;
  },
};

const baseQueryWhoIsControllerRequest: object = { controller: "" };

export const QueryWhoIsControllerRequest = {
  encode(
    message: QueryWhoIsControllerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.controller !== "") {
      writer.uint32(10).string(message.controller);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryWhoIsControllerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryWhoIsControllerRequest,
    } as QueryWhoIsControllerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.controller = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhoIsControllerRequest {
    const message = {
      ...baseQueryWhoIsControllerRequest,
    } as QueryWhoIsControllerRequest;
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = String(object.controller);
    } else {
      message.controller = "";
    }
    return message;
  },

  toJSON(message: QueryWhoIsControllerRequest): unknown {
    const obj: any = {};
    message.controller !== undefined && (obj.controller = message.controller);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryWhoIsControllerRequest>
  ): QueryWhoIsControllerRequest {
    const message = {
      ...baseQueryWhoIsControllerRequest,
    } as QueryWhoIsControllerRequest;
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = object.controller;
    } else {
      message.controller = "";
    }
    return message;
  },
};

const baseQueryWhoIsControllerResponse: object = {};

export const QueryWhoIsControllerResponse = {
  encode(
    _: QueryWhoIsControllerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryWhoIsControllerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryWhoIsControllerResponse,
    } as QueryWhoIsControllerResponse;
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

  fromJSON(_: any): QueryWhoIsControllerResponse {
    const message = {
      ...baseQueryWhoIsControllerResponse,
    } as QueryWhoIsControllerResponse;
    return message;
  },

  toJSON(_: QueryWhoIsControllerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryWhoIsControllerResponse>
  ): QueryWhoIsControllerResponse {
    const message = {
      ...baseQueryWhoIsControllerResponse,
    } as QueryWhoIsControllerResponse;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a WhoIs by id. */
  WhoIs(request: QueryWhoIsRequest): Promise<QueryWhoIsResponse>;
  /** Queries a list of WhoIs items. */
  WhoIsAll(request: QueryAllWhoIsRequest): Promise<QueryAllWhoIsResponse>;
  /** Queries a list of WhoIsAlias items. */
  WhoIsAlias(request: QueryWhoIsAliasRequest): Promise<QueryWhoIsAliasResponse>;
  /** Queries a list of WhoIsController items. */
  WhoIsController(
    request: QueryWhoIsControllerRequest
  ): Promise<QueryWhoIsControllerResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  WhoIs(request: QueryWhoIsRequest): Promise<QueryWhoIsResponse> {
    const data = QueryWhoIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Query",
      "WhoIs",
      data
    );
    return promise.then((data) => QueryWhoIsResponse.decode(new Reader(data)));
  }

  WhoIsAll(request: QueryAllWhoIsRequest): Promise<QueryAllWhoIsResponse> {
    const data = QueryAllWhoIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Query",
      "WhoIsAll",
      data
    );
    return promise.then((data) =>
      QueryAllWhoIsResponse.decode(new Reader(data))
    );
  }

  WhoIsAlias(
    request: QueryWhoIsAliasRequest
  ): Promise<QueryWhoIsAliasResponse> {
    const data = QueryWhoIsAliasRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Query",
      "WhoIsAlias",
      data
    );
    return promise.then((data) =>
      QueryWhoIsAliasResponse.decode(new Reader(data))
    );
  }

  WhoIsController(
    request: QueryWhoIsControllerRequest
  ): Promise<QueryWhoIsControllerResponse> {
    const data = QueryWhoIsControllerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Query",
      "WhoIsController",
      data
    );
    return promise.then((data) =>
      QueryWhoIsControllerResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
