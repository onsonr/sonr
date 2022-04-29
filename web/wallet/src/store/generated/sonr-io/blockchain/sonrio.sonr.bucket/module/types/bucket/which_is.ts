/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { BucketDoc } from "../bucket/bucket";

export const protobufPackage = "sonrio.sonr.bucket";

export interface WhichIs {
  /** DID is the DID of the bucket. */
  did: string;
  /** Creator is the Account that created the bucket. */
  creator: string;
  /** Bucket is the document of the bucket. */
  bucket: BucketDoc | undefined;
  /** Timestamp is the time of the last update of the DID Document */
  timestamp: number;
  /** IsActive is the status of the DID Document */
  is_active: boolean;
}

const baseWhichIs: object = {
  did: "",
  creator: "",
  timestamp: 0,
  is_active: false,
};

export const WhichIs = {
  encode(message: WhichIs, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    if (message.bucket !== undefined) {
      BucketDoc.encode(message.bucket, writer.uint32(34).fork()).ldelim();
    }
    if (message.timestamp !== 0) {
      writer.uint32(40).int64(message.timestamp);
    }
    if (message.is_active === true) {
      writer.uint32(48).bool(message.is_active);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): WhichIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhichIs } as WhichIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        case 4:
          message.bucket = BucketDoc.decode(reader, reader.uint32());
          break;
        case 5:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.is_active = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WhichIs {
    const message = { ...baseWhichIs } as WhichIs;
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
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromJSON(object.bucket);
    } else {
      message.bucket = undefined;
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

  toJSON(message: WhichIs): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.creator !== undefined && (obj.creator = message.creator);
    message.bucket !== undefined &&
      (obj.bucket = message.bucket
        ? BucketDoc.toJSON(message.bucket)
        : undefined);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.is_active !== undefined && (obj.is_active = message.is_active);
    return obj;
  },

  fromPartial(object: DeepPartial<WhichIs>): WhichIs {
    const message = { ...baseWhichIs } as WhichIs;
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
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromPartial(object.bucket);
    } else {
      message.bucket = undefined;
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
