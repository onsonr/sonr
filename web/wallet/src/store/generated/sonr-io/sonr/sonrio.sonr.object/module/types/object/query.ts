/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../object/params";
import { Session } from "../registry/who_is";
import { WhatIs } from "../object/what_is";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sonrio.sonr.object";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryWhatIsRequest {
  did: string;
  session: Session | undefined;
}

export interface QueryWhatIsResponse {
  what_is: WhatIs | undefined;
}

export interface QueryAllWhatIsRequest {
  pagination: PageRequest | undefined;
  session: Session | undefined;
}

export interface QueryAllWhatIsResponse {
  what_is: WhatIs[];
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

const baseQueryWhatIsRequest: object = { did: "" };

export const QueryWhatIsRequest = {
  encode(
    message: QueryWhatIsRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryWhatIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhatIsRequest } as QueryWhatIsRequest;
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

  fromJSON(object: any): QueryWhatIsRequest {
    const message = { ...baseQueryWhatIsRequest } as QueryWhatIsRequest;
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

  toJSON(message: QueryWhatIsRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhatIsRequest>): QueryWhatIsRequest {
    const message = { ...baseQueryWhatIsRequest } as QueryWhatIsRequest;
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

const baseQueryWhatIsResponse: object = {};

export const QueryWhatIsResponse = {
  encode(
    message: QueryWhatIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.what_is !== undefined) {
      WhatIs.encode(message.what_is, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryWhatIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryWhatIsResponse } as QueryWhatIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.what_is = WhatIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryWhatIsResponse {
    const message = { ...baseQueryWhatIsResponse } as QueryWhatIsResponse;
    if (object.what_is !== undefined && object.what_is !== null) {
      message.what_is = WhatIs.fromJSON(object.what_is);
    } else {
      message.what_is = undefined;
    }
    return message;
  },

  toJSON(message: QueryWhatIsResponse): unknown {
    const obj: any = {};
    message.what_is !== undefined &&
      (obj.what_is = message.what_is
        ? WhatIs.toJSON(message.what_is)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryWhatIsResponse>): QueryWhatIsResponse {
    const message = { ...baseQueryWhatIsResponse } as QueryWhatIsResponse;
    if (object.what_is !== undefined && object.what_is !== null) {
      message.what_is = WhatIs.fromPartial(object.what_is);
    } else {
      message.what_is = undefined;
    }
    return message;
  },
};

const baseQueryAllWhatIsRequest: object = {};

export const QueryAllWhatIsRequest = {
  encode(
    message: QueryAllWhatIsRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhatIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhatIsRequest } as QueryAllWhatIsRequest;
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

  fromJSON(object: any): QueryAllWhatIsRequest {
    const message = { ...baseQueryAllWhatIsRequest } as QueryAllWhatIsRequest;
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

  toJSON(message: QueryAllWhatIsRequest): unknown {
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
    object: DeepPartial<QueryAllWhatIsRequest>
  ): QueryAllWhatIsRequest {
    const message = { ...baseQueryAllWhatIsRequest } as QueryAllWhatIsRequest;
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

const baseQueryAllWhatIsResponse: object = {};

export const QueryAllWhatIsResponse = {
  encode(
    message: QueryAllWhatIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.what_is) {
      WhatIs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhatIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhatIsResponse } as QueryAllWhatIsResponse;
    message.what_is = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.what_is.push(WhatIs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllWhatIsResponse {
    const message = { ...baseQueryAllWhatIsResponse } as QueryAllWhatIsResponse;
    message.what_is = [];
    if (object.what_is !== undefined && object.what_is !== null) {
      for (const e of object.what_is) {
        message.what_is.push(WhatIs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhatIsResponse): unknown {
    const obj: any = {};
    if (message.what_is) {
      obj.what_is = message.what_is.map((e) =>
        e ? WhatIs.toJSON(e) : undefined
      );
    } else {
      obj.what_is = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWhatIsResponse>
  ): QueryAllWhatIsResponse {
    const message = { ...baseQueryAllWhatIsResponse } as QueryAllWhatIsResponse;
    message.what_is = [];
    if (object.what_is !== undefined && object.what_is !== null) {
      for (const e of object.what_is) {
        message.what_is.push(WhatIs.fromPartial(e));
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
   * WhatIs
   *
   * Queries a WhatIs by DID.
   */
  WhatIs(request: QueryWhatIsRequest): Promise<QueryWhatIsResponse>;
  /**
   * WhatIsAll
   *
   * Queries a list of WhatIs items.
   */
  WhatIsAll(request: QueryAllWhatIsRequest): Promise<QueryAllWhatIsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  WhatIs(request: QueryWhatIsRequest): Promise<QueryWhatIsResponse> {
    const data = QueryWhatIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Query",
      "WhatIs",
      data
    );
    return promise.then((data) => QueryWhatIsResponse.decode(new Reader(data)));
  }

  WhatIsAll(request: QueryAllWhatIsRequest): Promise<QueryAllWhatIsResponse> {
    const data = QueryAllWhatIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Query",
      "WhatIsAll",
      data
    );
    return promise.then((data) =>
      QueryAllWhatIsResponse.decode(new Reader(data))
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
