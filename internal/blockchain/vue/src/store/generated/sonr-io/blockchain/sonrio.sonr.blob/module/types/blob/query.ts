/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../blob/params";
import { ThereIs } from "../blob/there_is";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sonrio.sonr.blob";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetThereIsRequest {
  index: string;
}

export interface QueryGetThereIsResponse {
  thereIs: ThereIs | undefined;
}

export interface QueryAllThereIsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllThereIsResponse {
  thereIs: ThereIs[];
  pagination: PageResponse | undefined;
}

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

const baseQueryGetThereIsRequest: object = { index: "" };

export const QueryGetThereIsRequest = {
  encode(
    message: QueryGetThereIsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetThereIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetThereIsRequest } as QueryGetThereIsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetThereIsRequest {
    const message = { ...baseQueryGetThereIsRequest } as QueryGetThereIsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetThereIsRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetThereIsRequest>
  ): QueryGetThereIsRequest {
    const message = { ...baseQueryGetThereIsRequest } as QueryGetThereIsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetThereIsResponse: object = {};

export const QueryGetThereIsResponse = {
  encode(
    message: QueryGetThereIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.thereIs !== undefined) {
      ThereIs.encode(message.thereIs, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetThereIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetThereIsResponse,
    } as QueryGetThereIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.thereIs = ThereIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetThereIsResponse {
    const message = {
      ...baseQueryGetThereIsResponse,
    } as QueryGetThereIsResponse;
    if (object.thereIs !== undefined && object.thereIs !== null) {
      message.thereIs = ThereIs.fromJSON(object.thereIs);
    } else {
      message.thereIs = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetThereIsResponse): unknown {
    const obj: any = {};
    message.thereIs !== undefined &&
      (obj.thereIs = message.thereIs
        ? ThereIs.toJSON(message.thereIs)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetThereIsResponse>
  ): QueryGetThereIsResponse {
    const message = {
      ...baseQueryGetThereIsResponse,
    } as QueryGetThereIsResponse;
    if (object.thereIs !== undefined && object.thereIs !== null) {
      message.thereIs = ThereIs.fromPartial(object.thereIs);
    } else {
      message.thereIs = undefined;
    }
    return message;
  },
};

const baseQueryAllThereIsRequest: object = {};

export const QueryAllThereIsRequest = {
  encode(
    message: QueryAllThereIsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllThereIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllThereIsRequest } as QueryAllThereIsRequest;
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

  fromJSON(object: any): QueryAllThereIsRequest {
    const message = { ...baseQueryAllThereIsRequest } as QueryAllThereIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllThereIsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllThereIsRequest>
  ): QueryAllThereIsRequest {
    const message = { ...baseQueryAllThereIsRequest } as QueryAllThereIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllThereIsResponse: object = {};

export const QueryAllThereIsResponse = {
  encode(
    message: QueryAllThereIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.thereIs) {
      ThereIs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllThereIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllThereIsResponse,
    } as QueryAllThereIsResponse;
    message.thereIs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.thereIs.push(ThereIs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllThereIsResponse {
    const message = {
      ...baseQueryAllThereIsResponse,
    } as QueryAllThereIsResponse;
    message.thereIs = [];
    if (object.thereIs !== undefined && object.thereIs !== null) {
      for (const e of object.thereIs) {
        message.thereIs.push(ThereIs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllThereIsResponse): unknown {
    const obj: any = {};
    if (message.thereIs) {
      obj.thereIs = message.thereIs.map((e) =>
        e ? ThereIs.toJSON(e) : undefined
      );
    } else {
      obj.thereIs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllThereIsResponse>
  ): QueryAllThereIsResponse {
    const message = {
      ...baseQueryAllThereIsResponse,
    } as QueryAllThereIsResponse;
    message.thereIs = [];
    if (object.thereIs !== undefined && object.thereIs !== null) {
      for (const e of object.thereIs) {
        message.thereIs.push(ThereIs.fromPartial(e));
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a ThereIs by index. */
  ThereIs(request: QueryGetThereIsRequest): Promise<QueryGetThereIsResponse>;
  /** Queries a list of ThereIs items. */
  ThereIsAll(request: QueryAllThereIsRequest): Promise<QueryAllThereIsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.blob.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  ThereIs(request: QueryGetThereIsRequest): Promise<QueryGetThereIsResponse> {
    const data = QueryGetThereIsRequest.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.blob.Query", "ThereIs", data);
    return promise.then((data) =>
      QueryGetThereIsResponse.decode(new Reader(data))
    );
  }

  ThereIsAll(
    request: QueryAllThereIsRequest
  ): Promise<QueryAllThereIsResponse> {
    const data = QueryAllThereIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Query",
      "ThereIsAll",
      data
    );
    return promise.then((data) =>
      QueryAllThereIsResponse.decode(new Reader(data))
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
