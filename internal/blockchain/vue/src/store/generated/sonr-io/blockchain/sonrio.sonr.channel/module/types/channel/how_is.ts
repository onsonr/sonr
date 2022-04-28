/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { ChannelDoc } from "../channel/channel";

export const protobufPackage = "sonrio.sonr.channel";

export interface HowIs {
  /** Did is the DID of the channel */
  did: string;
  /** Document is the DID Document of the registered name and account encoded as JSON */
  creator: string;
  /** ChannelDoc is the structure of the channel encoded as JSON */
  channel: ChannelDoc | undefined;
  /** Timestamp is the time of the last update of the DID Document */
  timestamp: number;
  /** IsActive is the status of the DID Document */
  is_active: boolean;
}

const baseHowIs: object = {
  did: "",
  creator: "",
  timestamp: 0,
  is_active: false,
};

export const HowIs = {
  encode(message: HowIs, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.channel !== undefined) {
      ChannelDoc.encode(message.channel, writer.uint32(26).fork()).ldelim();
    }
    if (message.timestamp !== 0) {
      writer.uint32(32).int64(message.timestamp);
    }
    if (message.is_active === true) {
      writer.uint32(40).bool(message.is_active);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): HowIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHowIs } as HowIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.channel = ChannelDoc.decode(reader, reader.uint32());
          break;
        case 4:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.is_active = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HowIs {
    const message = { ...baseHowIs } as HowIs;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = ChannelDoc.fromJSON(object.channel);
    } else {
      message.channel = undefined;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    if (object.is_active !== undefined && object.is_active !== null) {
      message.is_active = Boolean(object.is_active);
    } else {
      message.is_active = false;
    }
    return message;
  },

  toJSON(message: HowIs): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.creator !== undefined && (obj.creator = message.creator);
    message.channel !== undefined &&
      (obj.channel = message.channel
        ? ChannelDoc.toJSON(message.channel)
        : undefined);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.is_active !== undefined && (obj.is_active = message.is_active);
    return obj;
  },

  fromPartial(object: DeepPartial<HowIs>): HowIs {
    const message = { ...baseHowIs } as HowIs;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = ChannelDoc.fromPartial(object.channel);
    } else {
      message.channel = undefined;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = 0;
    }
    if (object.is_active !== undefined && object.is_active !== null) {
      message.is_active = object.is_active;
    } else {
      message.is_active = false;
    }
    return message;
  },
};

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
