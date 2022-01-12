/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "wallet.v1";

/** SNID is the unique identifier of a user with Sonr Protocol */
export interface SNID {
  /** Domain Name of the node, e.g. prad.snr */
  domain: string;
  /** Public Key of the node, value in the Domains second TXT record */
  pubKey: Uint8Array;
  /** Peer ID of the node, calculated from the public key */
  peerId: string;
  /** DID of the node, calculated from the public key */
  did: string;
}

/** UUID is General Message ID with Signature, ID, and Timestamp */
export interface UUID {
  /** Signature of the message */
  signature: Uint8Array;
  /** ID of the message */
  value: string;
  /** Unix timestamp */
  timestamp: number;
}

function createBaseSNID(): SNID {
  return { domain: "", pubKey: new Uint8Array(), peerId: "", did: "" };
}

export const SNID = {
  encode(message: SNID, writer: Writer = Writer.create()): Writer {
    if (message.domain !== "") {
      writer.uint32(10).string(message.domain);
    }
    if (message.pubKey.length !== 0) {
      writer.uint32(18).bytes(message.pubKey);
    }
    if (message.peerId !== "") {
      writer.uint32(26).string(message.peerId);
    }
    if (message.did !== "") {
      writer.uint32(34).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SNID {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSNID();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domain = reader.string();
          break;
        case 2:
          message.pubKey = reader.bytes();
          break;
        case 3:
          message.peerId = reader.string();
          break;
        case 4:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SNID {
    return {
      domain: isSet(object.domain) ? String(object.domain) : "",
      pubKey: isSet(object.pubKey)
        ? bytesFromBase64(object.pubKey)
        : new Uint8Array(),
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      did: isSet(object.did) ? String(object.did) : "",
    };
  },

  toJSON(message: SNID): unknown {
    const obj: any = {};
    message.domain !== undefined && (obj.domain = message.domain);
    message.pubKey !== undefined &&
      (obj.pubKey = base64FromBytes(
        message.pubKey !== undefined ? message.pubKey : new Uint8Array()
      ));
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SNID>, I>>(object: I): SNID {
    const message = createBaseSNID();
    message.domain = object.domain ?? "";
    message.pubKey = object.pubKey ?? new Uint8Array();
    message.peerId = object.peerId ?? "";
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseUUID(): UUID {
  return { signature: new Uint8Array(), value: "", timestamp: 0 };
}

export const UUID = {
  encode(message: UUID, writer: Writer = Writer.create()): Writer {
    if (message.signature.length !== 0) {
      writer.uint32(10).bytes(message.signature);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    if (message.timestamp !== 0) {
      writer.uint32(24).int64(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UUID {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUUID();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.signature = reader.bytes();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UUID {
    return {
      signature: isSet(object.signature)
        ? bytesFromBase64(object.signature)
        : new Uint8Array(),
      value: isSet(object.value) ? String(object.value) : "",
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
    };
  },

  toJSON(message: UUID): unknown {
    const obj: any = {};
    message.signature !== undefined &&
      (obj.signature = base64FromBytes(
        message.signature !== undefined ? message.signature : new Uint8Array()
      ));
    message.value !== undefined && (obj.value = message.value);
    message.timestamp !== undefined &&
      (obj.timestamp = Math.round(message.timestamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UUID>, I>>(object: I): UUID {
    const message = createBaseUUID();
    message.signature = object.signature ?? new Uint8Array();
    message.value = object.value ?? "";
    message.timestamp = object.timestamp ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (const byte of arr) {
    bin.push(String.fromCharCode(byte));
  }
  return btoa(bin.join(""));
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
