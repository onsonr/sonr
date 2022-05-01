/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../channel/params";
import { Session } from "../registry/who_is";
import { HowIs } from "../channel/how_is";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sonrio.sonr.channel";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryHowIsRequest {
  did: string;
  session: Session | undefined;
}

export interface QueryHowIsResponse {
  how_is: HowIs | undefined;
}

export interface QueryAllHowIsRequest {
  pagination: PageRequest | undefined;
  session: Session | undefined;
}

export interface QueryAllHowIsResponse {
  how_is: HowIs[];
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

const baseQueryHowIsRequest: object = { did: "" };

export const QueryHowIsRequest = {
  encode(message: QueryHowIsRequest, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryHowIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryHowIsRequest } as QueryHowIsRequest;
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

  fromJSON(object: any): QueryHowIsRequest {
    const message = { ...baseQueryHowIsRequest } as QueryHowIsRequest;
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

  toJSON(message: QueryHowIsRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryHowIsRequest>): QueryHowIsRequest {
    const message = { ...baseQueryHowIsRequest } as QueryHowIsRequest;
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

const baseQueryHowIsResponse: object = {};

export const QueryHowIsResponse = {
  encode(
    message: QueryHowIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.how_is !== undefined) {
      HowIs.encode(message.how_is, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryHowIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryHowIsResponse } as QueryHowIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.how_is = HowIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryHowIsResponse {
    const message = { ...baseQueryHowIsResponse } as QueryHowIsResponse;
    if (object.how_is !== undefined && object.how_is !== null) {
      message.how_is = HowIs.fromJSON(object.how_is);
    } else {
      message.how_is = undefined;
    }
    return message;
  },

  toJSON(message: QueryHowIsResponse): unknown {
    const obj: any = {};
    message.how_is !== undefined &&
      (obj.how_is = message.how_is ? HowIs.toJSON(message.how_is) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryHowIsResponse>): QueryHowIsResponse {
    const message = { ...baseQueryHowIsResponse } as QueryHowIsResponse;
    if (object.how_is !== undefined && object.how_is !== null) {
      message.how_is = HowIs.fromPartial(object.how_is);
    } else {
      message.how_is = undefined;
    }
    return message;
  },
};

const baseQueryAllHowIsRequest: object = {};

export const QueryAllHowIsRequest = {
  encode(
    message: QueryAllHowIsRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryAllHowIsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllHowIsRequest } as QueryAllHowIsRequest;
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

  fromJSON(object: any): QueryAllHowIsRequest {
    const message = { ...baseQueryAllHowIsRequest } as QueryAllHowIsRequest;
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

  toJSON(message: QueryAllHowIsRequest): unknown {
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

  fromPartial(object: DeepPartial<QueryAllHowIsRequest>): QueryAllHowIsRequest {
    const message = { ...baseQueryAllHowIsRequest } as QueryAllHowIsRequest;
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

const baseQueryAllHowIsResponse: object = {};

export const QueryAllHowIsResponse = {
  encode(
    message: QueryAllHowIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.how_is) {
      HowIs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllHowIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllHowIsResponse } as QueryAllHowIsResponse;
    message.how_is = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.how_is.push(HowIs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllHowIsResponse {
    const message = { ...baseQueryAllHowIsResponse } as QueryAllHowIsResponse;
    message.how_is = [];
    if (object.how_is !== undefined && object.how_is !== null) {
      for (const e of object.how_is) {
        message.how_is.push(HowIs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllHowIsResponse): unknown {
    const obj: any = {};
    if (message.how_is) {
      obj.how_is = message.how_is.map((e) => (e ? HowIs.toJSON(e) : undefined));
    } else {
      obj.how_is = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllHowIsResponse>
  ): QueryAllHowIsResponse {
    const message = { ...baseQueryAllHowIsResponse } as QueryAllHowIsResponse;
    message.how_is = [];
    if (object.how_is !== undefined && object.how_is !== null) {
      for (const e of object.how_is) {
        message.how_is.push(HowIs.fromPartial(e));
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
   * HowIs
   *
   * Queries a HowIs by did.
   */
  HowIs(request: QueryHowIsRequest): Promise<QueryHowIsResponse>;
  /**
   * HowIsAll
   *
   * Queries a list of HowIs items.
   */
  HowIsAll(request: QueryAllHowIsRequest): Promise<QueryAllHowIsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  HowIs(request: QueryHowIsRequest): Promise<QueryHowIsResponse> {
    const data = QueryHowIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Query",
      "HowIs",
      data
    );
    return promise.then((data) => QueryHowIsResponse.decode(new Reader(data)));
  }

  HowIsAll(request: QueryAllHowIsRequest): Promise<QueryAllHowIsResponse> {
    const data = QueryAllHowIsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Query",
      "HowIsAll",
      data
    );
    return promise.then((data) =>
      QueryAllHowIsResponse.decode(new Reader(data))
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
