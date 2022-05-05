/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../bucket/params";
import { Session } from "../registry/who_is";
import { WhichIs } from "../bucket/which_is";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sonrio.sonr.bucket";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryWhichIsRequest {
  did: string;
  session: Session | undefined;
}

export interface QueryWhichIsResponse {
  which_is: WhichIs | undefined;
}

export interface QueryAllWhichIsRequest {
  pagination: PageRequest | undefined;
  session: Session | undefined;
}

export interface QueryAllWhichIsResponse {
  which_is: WhichIs[];
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

const baseQueryWhichIsRequest: object = { did: "" };

export const QueryWhichIsRequest = {
  encode(
    message: QueryWhichIsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhichIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhichIsRequest } as QueryWhichIsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhichIsRequest {
    const message = { ...baseQueryWhichIsRequest } as QueryWhichIsRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: QueryWhichIsRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhichIsRequest>): QueryWhichIsRequest {
    const message = { ...baseQueryWhichIsRequest } as QueryWhichIsRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseQueryWhichIsResponse: object = {};

export const QueryWhichIsResponse = {
  encode(
    message: QueryWhichIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.which_is !== undefined) {
      WhichIs.encode(message.which_is, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhichIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhichIsResponse } as QueryWhichIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.which_is = WhichIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhichIsResponse {
    const message = { ...baseQueryWhichIsResponse } as QueryWhichIsResponse;
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromJSON(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },

  toJSON(message: QueryWhichIsResponse): unknown {
    const obj: any = {};
    message.which_is !== undefined &&
      (obj.which_is = message.which_is
        ? WhichIs.toJSON(message.which_is)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhichIsResponse>): QueryWhichIsResponse {
    const message = { ...baseQueryWhichIsResponse } as QueryWhichIsResponse;
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromPartial(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },
};

const baseQueryAllWhichIsRequest: object = {};

export const QueryAllWhichIsRequest = {
  encode(
    message: QueryAllWhichIsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhichIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhichIsRequest } as QueryAllWhichIsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        case 2:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllWhichIsRequest {
    const message = { ...baseQueryAllWhichIsRequest } as QueryAllWhichIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhichIsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWhichIsRequest>
  ): QueryAllWhichIsRequest {
    const message = { ...baseQueryAllWhichIsRequest } as QueryAllWhichIsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseQueryAllWhichIsResponse: object = {};

export const QueryAllWhichIsResponse = {
  encode(
    message: QueryAllWhichIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.which_is) {
      WhichIs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhichIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllWhichIsResponse,
    } as QueryAllWhichIsResponse;
    message.which_is = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.which_is.push(WhichIs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllWhichIsResponse {
    const message = {
      ...baseQueryAllWhichIsResponse,
    } as QueryAllWhichIsResponse;
    message.which_is = [];
    if (object.which_is !== undefined && object.which_is !== null) {
      for (const e of object.which_is) {
        message.which_is.push(WhichIs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhichIsResponse): unknown {
    const obj: any = {};
    if (message.which_is) {
      obj.which_is = message.which_is.map((e) =>
        e ? WhichIs.toJSON(e) : undefined
      );
    } else {
      obj.which_is = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWhichIsResponse>
  ): QueryAllWhichIsResponse {
    const message = {
      ...baseQueryAllWhichIsResponse,
    } as QueryAllWhichIsResponse;
    message.which_is = [];
    if (object.which_is !== undefined && object.which_is !== null) {
      for (const e of object.which_is) {
        message.which_is.push(WhichIs.fromPartial(e));
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
  /**
   * Params
   *
   * Parameters queries the parameters of the module.
   */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /**
   * WhichIs
   *
   * Queries a WhichIs by did.
   */
  WhichIs(request: QueryWhichIsRequest): Promise<QueryWhichIsResponse>;
  /**
   * WhichIsAll
   *
   * Queries a list of WhichIs items.
   */
  WhichIsAll(request: QueryAllWhichIsRequest): Promise<QueryAllWhichIsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  WhichIs(request: QueryWhichIsRequest): Promise<QueryWhichIsResponse> {
    const data = QueryWhichIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Query",
      "WhichIs",
      data
    );
    return promise.then((data) =>
      QueryWhichIsResponse.decode(new Reader(data))
    );
  }

  WhichIsAll(
    request: QueryAllWhichIsRequest
  ): Promise<QueryAllWhichIsResponse> {
    const data = QueryAllWhichIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Query",
      "WhichIsAll",
      data
    );
    return promise.then((data) =>
      QueryAllWhichIsResponse.decode(new Reader(data))
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
