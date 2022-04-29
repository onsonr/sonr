/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "sonrio.sonr.channel";

export interface MsgCreateChannel {
  creator: string;
  name: string;
  description: string;
  objectDid: string;
  ttl: number;
  maxSize: number;
}

export interface MsgCreateChannelResponse {}

export interface MsgReadChannel {
  creator: string;
  did: string;
}

export interface MsgReadChannelResponse {}

export interface MsgDeactivateChannel {
  creator: string;
  did: string;
  publicKey: string;
}

export interface MsgDeactivateChannelResponse {}

export interface MsgUpdateChannel {
  creator: string;
  did: string;
}

export interface MsgUpdateChannelResponse {}

const baseMsgCreateChannel: object = {
  creator: "",
  name: "",
  description: "",
  objectDid: "",
  ttl: 0,
  maxSize: 0,
};

export const MsgCreateChannel = {
  encode(message: MsgCreateChannel, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.objectDid !== "") {
      writer.uint32(34).string(message.objectDid);
    }
    if (message.ttl !== 0) {
      writer.uint32(40).int64(message.ttl);
    }
    if (message.maxSize !== 0) {
      writer.uint32(48).int64(message.maxSize);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChannel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateChannel } as MsgCreateChannel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.objectDid = reader.string();
          break;
        case 5:
          message.ttl = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.maxSize = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChannel {
    const message = { ...baseMsgCreateChannel } as MsgCreateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.objectDid !== undefined && object.objectDid !== null) {
      message.objectDid = String(object.objectDid);
    } else {
      message.objectDid = "";
    }
    if (object.ttl !== undefined && object.ttl !== null) {
      message.ttl = Number(object.ttl);
    } else {
      message.ttl = 0;
    }
    if (object.maxSize !== undefined && object.maxSize !== null) {
      message.maxSize = Number(object.maxSize);
    } else {
      message.maxSize = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateChannel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.objectDid !== undefined && (obj.objectDid = message.objectDid);
    message.ttl !== undefined && (obj.ttl = message.ttl);
    message.maxSize !== undefined && (obj.maxSize = message.maxSize);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateChannel>): MsgCreateChannel {
    const message = { ...baseMsgCreateChannel } as MsgCreateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.objectDid !== undefined && object.objectDid !== null) {
      message.objectDid = object.objectDid;
    } else {
      message.objectDid = "";
    }
    if (object.ttl !== undefined && object.ttl !== null) {
      message.ttl = object.ttl;
    } else {
      message.ttl = 0;
    }
    if (object.maxSize !== undefined && object.maxSize !== null) {
      message.maxSize = object.maxSize;
    } else {
      message.maxSize = 0;
    }
    return message;
  },
};

const baseMsgCreateChannelResponse: object = {};

export const MsgCreateChannelResponse = {
  encode(
    _: MsgCreateChannelResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChannelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChannelResponse,
    } as MsgCreateChannelResponse;
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

  fromJSON(_: any): MsgCreateChannelResponse {
    const message = {
      ...baseMsgCreateChannelResponse,
    } as MsgCreateChannelResponse;
    return message;
  },

  toJSON(_: MsgCreateChannelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateChannelResponse>
  ): MsgCreateChannelResponse {
    const message = {
      ...baseMsgCreateChannelResponse,
    } as MsgCreateChannelResponse;
    return message;
  },
};

const baseMsgReadChannel: object = { creator: "", did: "" };

export const MsgReadChannel = {
  encode(message: MsgReadChannel, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReadChannel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReadChannel } as MsgReadChannel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReadChannel {
    const message = { ...baseMsgReadChannel } as MsgReadChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: MsgReadChannel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReadChannel>): MsgReadChannel {
    const message = { ...baseMsgReadChannel } as MsgReadChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseMsgReadChannelResponse: object = {};

export const MsgReadChannelResponse = {
  encode(_: MsgReadChannelResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReadChannelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReadChannelResponse } as MsgReadChannelResponse;
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

  fromJSON(_: any): MsgReadChannelResponse {
    const message = { ...baseMsgReadChannelResponse } as MsgReadChannelResponse;
    return message;
  },

  toJSON(_: MsgReadChannelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgReadChannelResponse>): MsgReadChannelResponse {
    const message = { ...baseMsgReadChannelResponse } as MsgReadChannelResponse;
    return message;
  },
};

const baseMsgDeactivateChannel: object = { creator: "", did: "", publicKey: "" };

export const MsgDeactivateChannel = {
  encode(message: MsgDeactivateChannel, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateChannel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeactivateChannel } as MsgDeactivateChannel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateChannel {
    const message = { ...baseMsgDeactivateChannel } as MsgDeactivateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = String(object.publicKey);
    } else {
      message.publicKey = "";
    }
    return message;
  },

  toJSON(message: MsgDeactivateChannel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateChannel>): MsgDeactivateChannel {
    const message = { ...baseMsgDeactivateChannel } as MsgDeactivateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = "";
    }
    return message;
  },
};

const baseMsgDeactivateChannelResponse: object = {};

export const MsgDeactivateChannelResponse = {
  encode(
    _: MsgDeactivateChannelResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeactivateChannelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeactivateChannelResponse,
    } as MsgDeactivateChannelResponse;
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

  fromJSON(_: any): MsgDeactivateChannelResponse {
    const message = {
      ...baseMsgDeactivateChannelResponse,
    } as MsgDeactivateChannelResponse;
    return message;
  },

  toJSON(_: MsgDeactivateChannelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeactivateChannelResponse>
  ): MsgDeactivateChannelResponse {
    const message = {
      ...baseMsgDeactivateChannelResponse,
    } as MsgDeactivateChannelResponse;
    return message;
  },
};

const baseMsgUpdateChannel: object = { creator: "", did: "" };

export const MsgUpdateChannel = {
  encode(message: MsgUpdateChannel, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateChannel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateChannel } as MsgUpdateChannel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateChannel {
    const message = { ...baseMsgUpdateChannel } as MsgUpdateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateChannel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateChannel>): MsgUpdateChannel {
    const message = { ...baseMsgUpdateChannel } as MsgUpdateChannel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseMsgUpdateChannelResponse: object = {};

export const MsgUpdateChannelResponse = {
  encode(
    _: MsgUpdateChannelResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateChannelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateChannelResponse,
    } as MsgUpdateChannelResponse;
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

  fromJSON(_: any): MsgUpdateChannelResponse {
    const message = {
      ...baseMsgUpdateChannelResponse,
    } as MsgUpdateChannelResponse;
    return message;
  },

  toJSON(_: MsgUpdateChannelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateChannelResponse>
  ): MsgUpdateChannelResponse {
    const message = {
      ...baseMsgUpdateChannelResponse,
    } as MsgUpdateChannelResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateChannel(request: MsgCreateChannel): Promise<MsgCreateChannelResponse>;
  ReadChannel(request: MsgReadChannel): Promise<MsgReadChannelResponse>;
  DeleteChannel(request: MsgDeactivateChannel): Promise<MsgDeactivateChannelResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateChannel(request: MsgUpdateChannel): Promise<MsgUpdateChannelResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateChannel(request: MsgCreateChannel): Promise<MsgCreateChannelResponse> {
    const data = MsgCreateChannel.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Msg",
      "CreateChannel",
      data
    );
    return promise.then((data) =>
      MsgCreateChannelResponse.decode(new Reader(data))
    );
  }

  ReadChannel(request: MsgReadChannel): Promise<MsgReadChannelResponse> {
    const data = MsgReadChannel.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Msg",
      "ReadChannel",
      data
    );
    return promise.then((data) =>
      MsgReadChannelResponse.decode(new Reader(data))
    );
  }

  DeleteChannel(request: MsgDeactivateChannel): Promise<MsgDeactivateChannelResponse> {
    const data = MsgDeactivateChannel.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Msg",
      "DeleteChannel",
      data
    );
    return promise.then((data) =>
      MsgDeactivateChannelResponse.decode(new Reader(data))
    );
  }

  UpdateChannel(request: MsgUpdateChannel): Promise<MsgUpdateChannelResponse> {
    const data = MsgUpdateChannel.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.channel.Msg",
      "UpdateChannel",
      data
    );
    return promise.then((data) =>
      MsgUpdateChannelResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
